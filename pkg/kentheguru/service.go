// Package kentheguru is the service which provides access to kontainerooo via a websocket server
package kentheguru

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/kontainerooo/kontainer.ooo/pkg/bart"
	"github.com/kontainerooo/kontainer.ooo/pkg/container"
	"github.com/kontainerooo/kontainer.ooo/pkg/kentheguru/pb"
	"github.com/kontainerooo/kontainer.ooo/pkg/module"

	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/kontainerooo/kontainer.ooo/pkg/user"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

// Service is the interface describing the KenTheGuru.Service used for communication with the frontend
type Service interface {
	StartWebsocketTransport(errorChannel chan error, logger log.Logger, wsAddr string)
}

type service struct {
	ProtocolMap        ws.ProtocolMap
	WebsocketUpgrader  websocket.Upgrader
	TokenAuth          ws.Authenticator
	BartBus            bart.Bus
	SSLConfig          ws.SSLConfig
	UserEndpoints      user.Endpoints
	KMIEndpoints       kmi.Endpoints
	ContainerEndpoints container.Endpoints
	RoutingEndpoints   routing.Endpoints
	ModuleEndpoints    module.Endpoints
}

func (s *service) StartWebsocketTransport(errc chan error, logger log.Logger, wsAddr string) {
	logger = log.With(logger, "transport", "ws")
	wss := ws.NewServer(s.ProtocolMap, logger, s.WebsocketUpgrader, s.TokenAuth, s.SSLConfig, s.ErrorHandler, ws.Before(s.BartBus.LostAndFound), ws.Before(s.BartBus.GetOff), ws.After(s.BartBus.GetOn))

	userService := user.MakeWebsocketService(s.UserEndpoints)
	wss.RegisterService(userService)

	kmiService := kmi.MakeWebsocketService(s.KMIEndpoints)
	wss.RegisterService(kmiService)

	csServer := container.MakeWebsocketService(s.ContainerEndpoints)
	wss.RegisterService(csServer)

	routingServer := routing.MakeWebsocketService(s.RoutingEndpoints)
	wss.RegisterService(routingServer)

	moduleServer := module.MakeWebsocketService(s.ModuleEndpoints)
	wss.RegisterService(moduleServer)

	logger.Log("addr", wsAddr)
	errc <- wss.Serve(wsAddr)
}

func (s *service) DecodeFunc(_ context.Context, data interface{}) (interface{}, error) {
	request := &pb.AuthenticationRequest{}
	err := proto.Unmarshal(data.([]byte), request)
	if err != nil {
		return nil, err
	}

	return user.CheckLoginCredentialsRequest{
		Username: request.Username,
		Password: request.Password,
	}, nil
}

func (s *service) EncodeFunc(_ context.Context, res interface{}) (interface{}, error) {
	token := res.(string)
	return &pb.AuthenticationResponse{
		Token: token,
	}, nil
}

func (s *service) MakeEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(user.CheckLoginCredentialsRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}

		res, err := s.UserEndpoints.CheckLoginCredentialsEndpoint(ctx, request)
		if err != nil {
			return nil, err
		}
		response := res.(user.CheckLoginCredentialsResponse)
		if response.ID == 0 {
			return nil, errors.New("not authenticated")
		}

		return bart.Claims{
			Username: req.Username,
			ID:       response.ID,
		}, nil
	}
}

func (s *service) ErrorHandler(srv, me *ws.ProtoID, err error, ph ws.ProtocolHandler) []byte {
	var (
		ktgID               = ws.ProtoIDFromString("KTG")
		errID               = ws.ProtoIDFromString("ERR")
		srvString, meString string
	)

	if srv != nil {
		srvString = srv.String()
	}
	if me != nil {
		meString = me.String()
	}

	res := &pb.ErrorResponse{
		Error:   err.Error(),
		Service: srvString,
		Method:  meString,
	}

	data, _ := ph.Encode(&ktgID, &errID, res)
	return data
}

// NewService returns a new KenTheGuru.Service instance
func NewService(
	authenticationKey, encryptionKey, signingKey string,
	upgrader websocket.Upgrader,
	sslConfig ws.SSLConfig,
	ue user.Endpoints,
	ke kmi.Endpoints,
	ce container.Endpoints,
	re routing.Endpoints,
	me module.Endpoints,
) Service {
	s := &service{
		ProtocolMap: ws.ProtocolMap{
			"v1": ws.BasicHandler{},
		},
		WebsocketUpgrader:  upgrader,
		BartBus:            bart.NewBus(signingKey, ue),
		SSLConfig:          sslConfig,
		UserEndpoints:      ue,
		KMIEndpoints:       ke,
		ContainerEndpoints: ce,
		RoutingEndpoints:   re,
		ModuleEndpoints:    me,
	}

	s.TokenAuth = ws.NewTokenAuth(
		ws.ProtoIDFromString("KTG"), ws.ProtoIDFromString("AUT"),
		authenticationKey, encryptionKey,
		signingKey,
		"KenTheGuru",
		s.DecodeFunc,
		s.EncodeFunc,
		s.MakeEndpoint(),
	)

	return s
}
