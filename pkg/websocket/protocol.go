package websocket

import "fmt"

// ProtoID is a byte array with length 3 used for identification in a protocol
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
	Decode(message []byte) (service *ProtoID, method *ProtoID, data interface{}, err error)
	Encode(service *ProtoID, method *ProtoID, data interface{}) (message []byte, err error)
}
