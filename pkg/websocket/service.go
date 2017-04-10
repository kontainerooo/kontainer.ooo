package websocket

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

var (
	// ErrNoName is returned, if a service or endpoint is created with no name
	ErrNoName = errors.New("no name provided")

	// ErrInvalidProtoID is returned, if the protoID is not valid
	ErrInvalidProtoID = errors.New("invalid protoID")

	// ErrNoEndpoint is returned, if no endpoint function is provided
	ErrNoEndpoint = errors.New("no endpoint provided")
)

// EndpointHandler is a function with calls an Endpoint with the decoded message
// and returns the encoded response or an error
type EndpointHandler func(message interface{}) (response interface{}, err error)

var stdDencode = func(_ context.Context, i interface{}) (interface{}, error) {
	return i, nil
}

// ServiceEndpoint is a struct type containing every value/function needed for an Endpoint in a Service
type ServiceEndpoint struct {
	name         string
	protocolName ProtoID
	e            endpoint.Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
}

// NewServiceEndpoint returns a pointer to a ServiceEdnpoint instance
func NewServiceEndpoint(
	name string,
	protocolName ProtoID,
	e endpoint.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
) (*ServiceEndpoint, error) {
	if name == "" {
		return nil, ErrNoName
	}

	if protocolName[0] == 0 && protocolName[1] == 0 && protocolName[2] == 0 {
		return nil, ErrInvalidProtoID
	}

	if e == nil {
		return nil, ErrNoEndpoint
	}

	if dec == nil {
		dec = stdDencode
	}

	if enc == nil {
		enc = stdDencode
	}

	return &ServiceEndpoint{
		name:         name,
		protocolName: protocolName,
		e:            e,
		dec:          dec,
		enc:          enc,
	}, nil
}

// ServiceDescription is a struct type containing every value needed for a Service in a Websocket Server
type ServiceDescription struct {
	name         string
	protocolName ProtoID
	endpoints    map[ProtoID]*ServiceEndpoint
}

// AddEndpoint takes a ServiceEndpoint and adds it to the ServiceDescription's endpoint map
func (s *ServiceDescription) AddEndpoint(se *ServiceEndpoint, err ...error) error {
	if len(err) != 0 {
		return err[0]
	}

	_, exist := s.endpoints[se.protocolName]
	if exist {
		return fmt.Errorf("Service Endpoint %s already exists", se.protocolName)
	}

	s.endpoints[se.protocolName] = se
	return nil
}

// GetEndpointHandler returns an EndpointHandler if an endpoint with name name exists, if not an error is returned
func (s *ServiceDescription) GetEndpointHandler(name ProtoID) (EndpointHandler, error) {
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
	protocolName ProtoID,
) (*ServiceDescription, error) {
	if name == "" {
		return nil, ErrNoName
	}

	if protocolName[0] == 0 && protocolName[1] == 0 && protocolName[2] == 0 {
		return nil, ErrInvalidProtoID
	}

	return &ServiceDescription{
		name:         name,
		protocolName: protocolName,
		endpoints:    make(map[ProtoID]*ServiceEndpoint),
	}, nil
}
