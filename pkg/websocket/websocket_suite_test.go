package websocket_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-kit/kit/endpoint"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var (
	errEndpoint       = errors.New("endpoint error")
	errDecode         = errors.New("decode error")
	errEncode         = errors.New("encode error")
	errService        = errors.New("service error")
	errMethod         = errors.New("method error")
	errProtocolEncode = errors.New("protocol encode error")
	protocolMap       = ws.ProtocolMap{
		"default": protocol{},
	}
)

func TestWebsocket(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Websocket Suite")
}

type request struct {
	req interface{}
}

type domainRequest struct {
	req interface{}
}

type response struct {
	res interface{}
}

type testAuth struct{}

func (testAuth) Mux(http.ResponseWriter, *http.Request) (interface{}, bool) {
	return nil, false
}

func (testAuth) GetID() ws.ProtoID {
	return ws.ProtoIDFromString("AUT")
}

func (testAuth) GetEndpoint() *ws.ServiceEndpoint {
	return &ws.ServiceEndpoint{}
}

type protocol struct{}

func (p protocol) Decode(message []byte) (*ws.ProtoID, *ws.ProtoID, interface{}, error) {
	m := string(message)
	mparts := strings.Split(m, " ")

	services := regexp.MustCompile("TST|NYI")
	methods := services

	if services.FindString(mparts[0]) == "" {
		return nil, nil, nil, errService
	}
	service := ws.ProtoIDFromString(mparts[0])

	if methods.FindString(mparts[1]) == "" {
		return nil, nil, nil, errMethod
	}
	method := ws.ProtoIDFromString(mparts[1])

	data := request{
		req: mparts[2],
	}

	number, err := strconv.ParseUint(mparts[2], 10, 64)
	if err == nil {
		data.req = number
	}

	return &service, &method, data, nil
}

func (p protocol) Encode(service *ws.ProtoID, method *ws.ProtoID, data interface{}) ([]byte, error) {
	res := data.(response).res
	if res.(string) == "error" {
		return nil, errProtocolEncode
	}

	return []byte(fmt.Sprintf("%s %s %s", *service, *method, res)), nil
}

func makeTestEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		reqStruct := request.(domainRequest)
		switch reqStruct.req.(type) {
		case error:
			return nil, errEndpoint
		default:
			return reqStruct.req, nil
		}
	}
}

func decodeTest(ctx context.Context, req interface{}) (interface{}, error) {
	reqStruct := req.(request)
	switch reqStruct.req.(type) {
	case bool:
		return nil, errDecode
	default:
		return domainRequest{
			req: reqStruct.req,
		}, nil
	}
}

func encodeTest(ctx context.Context, res interface{}) (interface{}, error) {
	switch res.(type) {
	case uint64:
		return nil, errEncode
	default:
		return response{
			res: res,
		}, nil
	}
}
