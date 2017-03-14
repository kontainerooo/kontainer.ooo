package websocket

import (
	"context"
)

// DecodeRequestFunc extracts a user-domain request object from a socket request.
// It's designed to be used in websocket servers, for server-side endpoints.
type DecodeRequestFunc func(context.Context, interface{}) (request interface{}, err error)

// EncodeResponseFunc encodes the passed response object to the socket response
// message. It's designed to be used in websocket servers, for server-side endpoints.
type EncodeResponseFunc func(context.Context, interface{}) (response interface{}, err error)
