package websocket

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

// EndpointHandler is a function with calls an Endpoint with the decoded message
// and returns the encoded response or an error
type EndpointHandler func(message interface{}) (response interface{}, err error)

// ServiceEndpoint is a struct type containing every value/function needed for an Endpoint in a Service
type ServiceEndpoint struct {
	name         string
	protocolName [3]byte
	e            endpoint.Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
}

// NewServiceEndpoint returns a pointer to a ServiceEdnpoint instance
func NewServiceEndpoint(
	name string,
	protocolName [3]byte,
	e endpoint.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
) *ServiceEndpoint {
	return &ServiceEndpoint{
		name:         name,
		protocolName: protocolName,
		e:            e,
		dec:          dec,
		enc:          enc,
	}
}

// ServiceDescription is a struct type containing every value needed for a Service in a Websocket Server
type ServiceDescription struct {
	name         string
	protocolName [3]byte
	endpoints    map[[3]byte]*ServiceEndpoint
}

// AddEndpoint takes a ServiceEndpoint and adds it to the ServiceDescription's endpoint map
func (s *ServiceDescription) AddEndpoint(se *ServiceEndpoint) error {
	_, exist := s.endpoints[se.protocolName]
	if exist {
		return fmt.Errorf("Service Endpoint %s already exists", se.protocolName)
	}

	s.endpoints[se.protocolName] = se
	return nil
}

// EndpointHandler returns an EndpointHandler if an endpoint with name name exists, if not an error is returned
func (s *ServiceDescription) EndpointHandler(name [3]byte) (EndpointHandler, error) {
	e, exist := s.endpoints[name]
	if !exist {
		return nil, fmt.Errorf("Service Endpoint %s does not exists", name)
	}

	return func(message interface{}) (interface{}, error) {
		ctx := context.Background()
		req, err := e.dec(ctx, message)
		if err != nil {
			return nil, err
		}

		res, err := e.e(ctx, req)
		if err != nil {
			return nil, err
		}

		return e.enc(ctx, res)
	}, nil
}

// NewServiceDescription returns a pointer to a ServiceDescription instance
func NewServiceDescription(
	name string,
	protocolName [3]byte,
) *ServiceDescription {
	return &ServiceDescription{
		name:         name,
		protocolName: protocolName,
	}
}
