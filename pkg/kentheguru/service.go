package kentheguru

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/kontainerooo/kontainer.ooo/pkg/bart"
	"github.com/kontainerooo/kontainer.ooo/pkg/containerlifecycle"
	"github.com/kontainerooo/kontainer.ooo/pkg/customercontainer"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/kontainerooo/kontainer.ooo/pkg/user"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

// Service is the interface describing the KenTheGuru.Service used for communication with the frontend
type Service interface {
	StartWebsocketTransport(errorChannel chan error, logger log.Logger, wsAddr string)
}

type service struct {
	ProtocolMap                 ws.ProtocolMap
	WebsocketUpgrader           websocket.Upgrader
	TokenAuth                   ws.Authenticator
	BartBus                     bart.Bus
	SSLConfig                   ws.SSLConfig
	UserEndpoints               user.Endpoints
	KMIEndpoints                kmi.Endpoints
	ContainerLifecycleEndpoints containerlifecycle.Endpoints
	CustomerContainerEndpoints  customercontainer.Endpoints
	RoutingEndpoints            routing.Endpoints
}

func (s *service) StartWebsocketTransport(errc chan error, logger log.Logger, wsAddr string) {
	logger = log.With(logger, "transport", "ws")
	wss := ws.NewServer(s.ProtocolMap, logger, s.WebsocketUpgrader, s.TokenAuth, s.SSLConfig, ws.Before(s.BartBus.GetOff))

	userService := user.MakeWebsocketService(s.UserEndpoints)
	wss.RegisterService(userService)

	kmiService := kmi.MakeWebsocketService(s.KMIEndpoints)
	wss.RegisterService(kmiService)

	clsServer := containerlifecycle.MakeWebsocketService(s.ContainerLifecycleEndpoints)
	wss.RegisterService(clsServer)

	ccsServer := customercontainer.MakeWebsocketService(s.CustomerContainerEndpoints)
	wss.RegisterService(ccsServer)

	routingServer := routing.MakeWebsocketService(s.RoutingEndpoints)
	wss.RegisterService(routingServer)

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
		res, err := s.UserEndpoints.CheckLoginCredentialsEndpoint(ctx, request)
		if err != nil {
			return nil, err
		}
		response := res.(user.CheckLoginCredentialsResponse)
		if response.ID == 0 {
			return nil, errors.New("not authenticated")
		}

		return Claims{
			Username: request.(user.CheckLoginCredentialsRequest).Username,
			ID:       response.ID,
		}, nil
	}
}

// NewService returns a new KenTheGuru.Service instance
func NewService(
	authenticationKey, encryptionKey, signingKey string,
	upgrader websocket.Upgrader,
	sslConfig ws.SSLConfig,
	ue user.Endpoints,
	ke kmi.Endpoints,
	cle containerlifecycle.Endpoints,
	cce customercontainer.Endpoints,
	re routing.Endpoints,
) Service {
	s := &service{
		ProtocolMap: ws.ProtocolMap{
			"v1": ws.BasicHandler{},
		},
		WebsocketUpgrader:           upgrader,
		BartBus:                     bart.NewBus(ue),
		SSLConfig:                   sslConfig,
		UserEndpoints:               ue,
		KMIEndpoints:                ke,
		ContainerLifecycleEndpoints: cle,
		CustomerContainerEndpoints:  cce,
		RoutingEndpoints:            re,
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
