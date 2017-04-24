package websocket_test

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-kit/kit/endpoint"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
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

type protocol struct{}

func (p protocol) Decode(message []byte) (*ws.ProtoID, *ws.ProtoID, interface{}, error) {
	var (
		service ws.ProtoID
		method  ws.ProtoID
		data    request
	)
	m := string(message)
	mparts := strings.Split(m, " ")
	switch mparts[0] {
	case "TST":
		service = ws.ProtoIDFromString(mparts[0])
		switch mparts[1] {
		case "TST":
			method = ws.ProtoIDFromString(mparts[1])
			data.req = mparts[2]
		default:
			return nil, nil, nil, errors.New("method")
		}
	default:
		return nil, nil, nil, fmt.Errorf("service %s doesnt exist", mparts[0])
	}
	return &service, &method, data, nil
}

func (p protocol) Encode(service *ws.ProtoID, method *ws.ProtoID, data interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%s %s %s", *service, *method, data.(response).res)), nil
}

func makeTestEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		reqStruct := request.(domainRequest)
		switch reqStruct.req.(type) {
		case error:
			return nil, errors.New("endpoint error")
		default:
			return reqStruct.req, nil
		}
	}
}

func decodeTest(ctx context.Context, req interface{}) (interface{}, error) {
	reqStruct := req.(request)
	switch reqStruct.req.(type) {
	case bool:
		return nil, errors.New("decode error")
	default:
		return domainRequest{
			req: reqStruct.req,
		}, nil
	}
}

func encodeTest(ctx context.Context, res interface{}) (interface{}, error) {
	switch res.(type) {
	case uint64:
		return nil, errors.New("encode error")
	default:
		return response{
			res: res,
		}, nil
	}
}
