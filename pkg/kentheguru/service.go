package kentheguru

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/websocket"
	"github.com/kontainerooo/kontainer.ooo/pkg/containerlifecycle"
	"github.com/kontainerooo/kontainer.ooo/pkg/customercontainer"
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
	ProtocolMap                 ws.ProtocolMap
	WebsocketUpgrader           websocket.Upgrader
	TokenAuth                   ws.Authenticator
	SSLConfig                   ws.SSLConfig
	UserEndpoints               user.Endpoints
	KMIEndpoints                kmi.Endpoints
	ContainerLifecycleEndpoints containerlifecycle.Endpoints
	CustomerContainerEndpoints  customercontainer.Endpoints
	RoutingEndpoints            routing.Endpoints
}

func (s *service) StartWebsocketTransport(errc chan error, logger log.Logger, wsAddr string) {
	logger = log.With(logger, "transport", "ws")
	wss := ws.NewServer(s.ProtocolMap, logger, s.WebsocketUpgrader, s.TokenAuth, s.SSLConfig)

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
	return &service{
		ProtocolMap: ws.ProtocolMap{
			"v1": ws.BasicHandler{},
		},
		WebsocketUpgrader: upgrader,
		TokenAuth: ws.NewTokenAuth(
			ws.ProtoIDFromString("AUT"), ws.ProtoIDFromString("AUT"),
			authenticationKey, encryptionKey,
			signingKey,
			"KenTheGuru",
			user.DecodeWSCheckLoginCredentialsRequest,
			user.EncodeGRPCCheckLoginCredentialsResponse,
			ue.CheckLoginCredentialsEndpoint,
		),
		SSLConfig:                   sslConfig,
		UserEndpoints:               ue,
		KMIEndpoints:                ke,
		ContainerLifecycleEndpoints: cle,
		CustomerContainerEndpoints:  cce,
		RoutingEndpoints:            re,
	}
}
