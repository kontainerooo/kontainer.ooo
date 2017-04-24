package websocket

import (
	"net/http"
)

// The Authenticator may be implemented for use with the websocket Server
type Authenticator interface {
	// Mux is basicaly a http.Handler but with an interface and a boolean return parameter
	// the interface may contain session data
	// the boolean should be set to true if a response was send and the conection should not be upgraded
	Mux(http.ResponseWriter, *http.Request) (interface{}, bool)

	// GetID is used to set the Service ProtoID for the Authentication
	GetID() ProtoID

	// GetEndpoint is used to add the Logical Endpoint to the Authentication Service
	GetEndpoint() *ServiceEndpoint
}

