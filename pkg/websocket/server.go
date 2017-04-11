package websocket

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/websocket"
)

// SSLConfig is a type containing the path to a certificate and its keyfile
type SSLConfig struct {
	certificate string
	key         string
	addr        string
	only        bool
}

// ProtocolMap maps a protocol name to its handler
type ProtocolMap map[string]ProtocolHandler

// Server is a struct type containing every value needed for a websocket server
type Server struct {
	protocols ProtocolMap
	logger    log.Logger
	services  map[ProtoID]*ServiceDescription
	upgrader  websocket.Upgrader
	ssl       SSLConfig
}

// RegisterService adds the given ServiceDescription to the Server's services map
func (s *Server) RegisterService(sd *ServiceDescription) error {
	_, exist := s.services[sd.protocolName]
	if exist {
		return fmt.Errorf("Service Endpoint %s already exists", sd.protocolName)
	}

	s.services[sd.protocolName] = sd
	return nil
}

// GetService returns a ServiceDescription given its ProtoID or an error
func (s *Server) GetService(name ProtoID) (*ServiceDescription, error) {
	sd, exist := s.services[name]
	if !exist {
		return nil, fmt.Errorf("Service Description %s does not exists", name)
	}

	return sd, nil
}

// Serve starts the http transport for the websocket, listening on addr
func (s *Server) Serve(addr string) error {
	if !s.ssl.only {
		err := http.ListenAndServe(addr, s)
		if err != nil {
			return err
		}
	}

	if s.ssl.certificate != "" && s.ssl.key != "" && s.ssl.addr != "" {
		return http.ListenAndServeTLS(s.ssl.addr, s.ssl.certificate, s.ssl.key, s)
	}

	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Log("err", err)
		return
	}

	s.logger.Log("conn", conn.RemoteAddr())
	go s.handleConnection(conn)
}

func (s *Server) handleConnection(conn *websocket.Conn) {
	defer conn.Close()

	protocolName := conn.Subprotocol()
	if protocolName == "" {
		protocolName = "default"
	}
	protocolHandler, ok := s.protocols[protocolName]

	if !ok {
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
				s.logger.Log("err", err)
				return
			}
			continue
		}

		service, err := s.GetService(*srv)
		if err != nil {
			err = conn.WriteMessage(messageType, []byte(err.Error()))
			if err != nil {
				s.logger.Log("err", err)
				return
			}
			continue
		}

		handler, err := service.GetEndpointHandler(*me)
		if err != nil {
			err = conn.WriteMessage(messageType, []byte(err.Error()))
			if err != nil {
				s.logger.Log("err", err)
				return
			}
			continue
		}

		res, err := handler(data)
		if err != nil {
			err = conn.WriteMessage(messageType, []byte(err.Error()))
			if err != nil {
				s.logger.Log("err", err)
				return
			}
			continue
		}

		response, err := protocolHandler.Encode(srv, me, res)
		if err != nil {
			err = conn.WriteMessage(messageType, []byte(err.Error()))
			if err != nil {
				s.logger.Log("err", err)
				return
			}
			continue
		}

		err = conn.WriteMessage(messageType, response)
		if err != nil {
			s.logger.Log("err", err)
			return
		}
	}
}

// NewServer returns a pointer to a Server instance
func NewServer(
	pm ProtocolMap,
	logger log.Logger,
	upgrader websocket.Upgrader,
	ssl SSLConfig,
) *Server {
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
		logger.Log("caution", "no upgrader provided, every connection will be accepted")
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	}

	for name := range pm {
		upgrader.Subprotocols = append(upgrader.Subprotocols, name)
	}

	return &Server{
		protocols: pm,
		logger:    logger,
		upgrader:  upgrader,
		ssl:       ssl,
		services:  make(map[ProtoID]*ServiceDescription),
	}
}
