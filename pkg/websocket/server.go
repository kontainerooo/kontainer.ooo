package websocket

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/websocket"
)

// MiddlewareFunc is a function type used in the websocket package
// Its parameters are the service and method id in a message as well as its data
// furthermore the session information is added
// Its return value may be an error
type MiddlewareFunc func(ProtoID, ProtoID, interface{}, interface{}) error

type position uint8

const (
	execBefore position = iota
	execAfter  position = iota
)

// Middleware is a type, which combines a MiddlewareFunc with a position
type Middleware struct {
	mid MiddlewareFunc
	pos position
}

// Before returns a middleware, which will be executed in the websocket loop
// before executing an endpoint, but after using the protocol handler to decode a message
func Before(m MiddlewareFunc) *Middleware {
	return &Middleware{m, execBefore}
}

// After returns a middleware, which will be executed in the websocket loop
// after executing an endpoint, but before using the protocol handler to encode a message
func After(m MiddlewareFunc) *Middleware {
	return &Middleware{m, execAfter}
}

// SSLConfig is a type containing the path to a certificate and its keyfile
type SSLConfig struct {
	// Path to the Certificate, which should be used to serve via ssl
	Certificate string

	// Path to the key, which belongs to the certificate
	Key string

	// Addr is the address on which the ssl server should listen
	Addr string

	// Only specfies if the transport should be ssl only
	Only bool
}

// ProtocolMap maps a protocol name to its handler
type ProtocolMap map[string]ProtocolHandler

// Server is a struct type containing every value needed for a websocket server
type Server struct {
	// Protocols is a ProtocolMap including the proctols supported by the websocket server
	Protocols ProtocolMap

	// Logger is the log.Logger instance used to log websocket related output
	Logger log.Logger

	// Upgrader is the websocket.Upgrader instance used for the websocket server
	// There is no need to define Subprotocols, since this will be filled with the help of the ProtocolMap
	Upgrader websocket.Upgrader

	auth     Authenticator
	ssl      SSLConfig
	services map[ProtoID]*ServiceDescription
	before   []*Middleware
	after    []*Middleware
}

// RegisterService adds the given ServiceDescription to the Server's map of services
func (s *Server) RegisterService(sd *ServiceDescription) error {
	_, exist := s.services[sd.ProtocolName]
	if exist {
		return fmt.Errorf("Service Endpoint %s already exists", sd.ProtocolName)
	}

	s.services[sd.ProtocolName] = sd
	return nil
}

// GetService returns a ServiceDescription given its ProtoID or an error
func (s *Server) GetService(name ProtoID) (*ServiceDescription, error) {
	sd, exist := s.services[name]
	if !exist {
		return nil, fmt.Errorf("Service Description %s does not exist", name)
	}

	return sd, nil
}

// Serve starts the http(s) transport for the websocket, listening on addr
func (s *Server) Serve(addr string) error {
	var serving bool

	if !s.ssl.Only {
		err := http.ListenAndServe(addr, s)
		if err != nil {
			return err
		}
		serving = true
	}

	if s.ssl.Certificate != "" && s.ssl.Key != "" && s.ssl.Addr != "" {
		return http.ListenAndServeTLS(s.ssl.Addr, s.ssl.Certificate, s.ssl.Key, s)
	}

	if !serving {
		return errors.New("incomplete configuration, no server started")
	}

	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		session interface{}
		abort   bool
	)

	if s.auth != nil {
		session, abort = s.auth.Mux(w, r)
		if abort {
			return
		}
	}

	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Logger.Log("err", err)
		return
	}

	s.Logger.Log("conn", conn.RemoteAddr())
	go s.handleConnection(conn, session)
}

func (s *Server) handleConnection(conn *websocket.Conn, session interface{}) {
	defer conn.Close()

	protocolName := conn.Subprotocol()
	if protocolName == "" {
		protocolName = "default"
	}
	protocolHandler, ok := s.Protocols[protocolName]

	if !ok || protocolHandler == nil {
		conn.WriteMessage(websocket.TextMessage, []byte("requested protocol not available"))
		return
	}

	for {
		messageType, request, err := conn.ReadMessage()
		if err != nil {
			return
		}

		srv, me, data, err := protocolHandler.Decode(request)
		if err != nil {
			err = conn.WriteMessage(messageType, []byte(err.Error()))
			if err != nil {
				s.Logger.Log("err", err)
				return
			}
			continue
		}

		for _, middleware := range s.before {
			err = middleware.mid(*srv, *me, &data, &session)
			if err != nil {
				err = conn.WriteMessage(messageType, []byte(err.Error()))
				if err != nil {
					s.Logger.Log("err", err)
					return
				}
				continue
			}
		}

		service, err := s.GetService(*srv)
		if err != nil {
			err = conn.WriteMessage(messageType, []byte(err.Error()))
			if err != nil {
				s.Logger.Log("err", err)
				return
			}
			continue
		}

		handler, err := service.GetEndpointHandler(*me)
		if err != nil {
			err = conn.WriteMessage(messageType, []byte(err.Error()))
			if err != nil {
				s.Logger.Log("err", err)
				return
			}
			continue
		}

		res, err := handler(data)
		if err != nil {
			err = conn.WriteMessage(messageType, []byte(err.Error()))
			if err != nil {
				s.Logger.Log("err", err)
				return
			}
			continue
		}

		for _, middleware := range s.after {
			err = middleware.mid(*srv, *me, &res, &session)
			if err != nil {
				err = conn.WriteMessage(messageType, []byte(err.Error()))
				if err != nil {
					s.Logger.Log("err", err)
					return
				}
				continue
			}
		}

		response, err := protocolHandler.Encode(srv, me, res)
		if err != nil {
			err = conn.WriteMessage(messageType, []byte(err.Error()))
			if err != nil {
				s.Logger.Log("err", err)
				return
			}
			continue
		}

		err = conn.WriteMessage(messageType, response)
		if err != nil {
			s.Logger.Log("err", err)
			return
		}
	}
}

// NewServer returns a pointer to a Server instance, given its dependencies
func NewServer(
	pm ProtocolMap,
	logger log.Logger,
	upgrader websocket.Upgrader,
	auth Authenticator,
	ssl SSLConfig,
	m ...*Middleware,
) *Server {
	var (
		before   = make([]*Middleware, 0)
		after    = make([]*Middleware, 0)
		services = make(map[ProtoID]*ServiceDescription)
		server   *Server
	)

	if upgrader.ReadBufferSize == 0 {
		if upgrader.WriteBufferSize != 0 {
			upgrader.ReadBufferSize = upgrader.WriteBufferSize
		} else {
			upgrader.ReadBufferSize = 1024
		}
	}

	if upgrader.WriteBufferSize == 0 {
		upgrader.WriteBufferSize = upgrader.ReadBufferSize
	}

	if upgrader.CheckOrigin == nil {
		logger.Log("caution", "no CheckOrigin function provided, every connection will be accepted")
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	}

	for name := range pm {
		upgrader.Subprotocols = append(upgrader.Subprotocols, name)
	}

	for _, w := range m {
		if w.pos == execBefore {
			before = append(before, w)
		} else if w.pos == execAfter {
			after = append(after, w)
		}
	}

	server = &Server{
		Protocols: pm,
		Logger:    logger,
		Upgrader:  upgrader,
		ssl:       ssl,
		auth:      auth,
		services:  services,
		before:    before,
		after:     after,
	}

	if auth != nil {
		authService, _ := NewServiceDescription("Authentification", auth.GetID())
		authService.AddEndpoint(auth.GetEndpoint())
		server.RegisterService(authService)
	}

	return server
}
