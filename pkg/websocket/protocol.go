package websocket

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
)

// ProtoID is a byte array with a length of 3 used for identification in a protocol
type ProtoID [3]byte

func (p ProtoID) String() string {
	return fmt.Sprintf("%s%s%s", string(p[0]), string(p[1]), string(p[2]))
}

// ProtoIDFromString creates a ProtoID from a string
func ProtoIDFromString(id string) ProtoID {
	return ProtoID{id[0], id[1], id[2]}
}

// ProtocolHandler is an interface defining the needed functionality to Decode and Encode
type ProtocolHandler interface {
	// Decode is function which extracts a service and method identifier as well as a data instance out of a message
	// If the message is malformatted an error should be returned
	Decode(message []byte) (service *ProtoID, method *ProtoID, data interface{}, err error)

	// Encode is a function which converts a service and a method identifier as well as a data instance into a message
	// If the message components are malformatted an error should be returned
	Encode(service *ProtoID, method *ProtoID, data interface{}) (message []byte, err error)
}

// BasicHandler is a basic protocolHandler which encodes the responses uses protobuf
type BasicHandler struct{}

// Decode implements the ProtocolHandler Decode function
func (h BasicHandler) Decode(message []byte) (*ProtoID, *ProtoID, interface{}, error) {
	var (
		service ProtoID
		method  ProtoID
		request interface{}
	)

	if len(message) < 6 {
		return nil, nil, nil, errors.New("unaccepted message format")
	}

	service = ProtoIDFromString(string(message[0:3]))
	method = ProtoIDFromString(string(message[3:6]))
	request = message[6:]

	return &service, &method, request, nil
}

// Encode implements the ProtocolHandler Encode function
func (h BasicHandler) Encode(service *ProtoID, method *ProtoID, data interface{}) ([]byte, error) {
	var message []byte

	pb, err := proto.Marshal(data.(proto.Message))
	if err != nil {
		return nil, err
	}

	message = append(message, []byte(service.String())...)
	message = append(message, []byte(method.String())...)
	message = append(message, pb...)

	return message, nil
}
