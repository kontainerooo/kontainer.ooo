package websocket_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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

var _ = Describe("Websocket", func() {
	var service *ws.ServiceDescription
	var endpoint *ws.ServiceEndpoint
	Context("Services", func() {
		Context("Create", func() {
			It("Should create a new ServiceDescription", func() {
				service = ws.NewServiceDescription("Test Service", ws.ProtoID{'T', 'S', 'T'})
				Ω(service).ShouldNot(BeNil())
			})

			It("Should create a new ServiceEndpoint", func() {
				e := makeTestEndpoint()
				endpoint = ws.NewServiceEndpoint("Test Endpoint", ws.ProtoID{'T', 'S', 'T'}, e, decodeTest, encodeTest)
			})
		})

		Context("Add Endpoint to Service", func() {
			It("Should add a ServiceEndpoint to a ServiceDescription", func() {
				err := service.AddEndpoint(endpoint)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Should return an Error if a Endpoint already exists in a Service", func() {
				err := service.AddEndpoint(endpoint)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("Get Endpoint Handler", func() {
			It("Should return an error if the requested endpoint does not exist", func() {
				_, err := service.GetEndpointHandler(ws.ProtoID{'N', 'O', 'T'})
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an EndpointHandler for an existing Endpoint", func() {
				handler, err := service.GetEndpointHandler(ws.ProtoID{'T', 'S', 'T'})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(handler).ShouldNot(BeNil())
			})
		})

		Context("Endpoint Handler", func() {
			It("Should return the encoded response", func() {
				req := "test"
				handler, _ := service.GetEndpointHandler(ws.ProtoID{'T', 'S', 'T'})
				res, err := handler(request{
					req: req,
				})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(res.(response).res).Should(BeEquivalentTo(req))
			})

			It("Should return a decode error", func() {
				handler, _ := service.GetEndpointHandler(ws.ProtoID{'T', 'S', 'T'})
				_, err := handler(request{
					req: true,
				})
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an endpoint error", func() {
				handler, _ := service.GetEndpointHandler(ws.ProtoID{'T', 'S', 'T'})
				_, err := handler(request{
					req: errors.New("error"),
				})
				Ω(err).Should(HaveOccurred())
			})

			It("Should return an encode error", func() {
				handler, _ := service.GetEndpointHandler(ws.ProtoID{'T', 'S', 'T'})
				_, err := handler(request{
					req: uint64(1),
				})
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	var wsserver *ws.Server
	Context("Servers", func() {
		Context("Create", func() {
			It("Should create a new Server", func() {
				wsserver = ws.NewServer(protocol{}, log.NewLogfmtLogger(os.Stdout))
			})
		})

		Context("Register Service", func() {
			It("Should register a new service", func() {
				err := wsserver.RegisterService(service)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Should return an error if a service is already registered", func() {
				err := wsserver.RegisterService(service)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("Serve HTTP and Connection Handling", func() {
			var httpserver *httptest.Server
			var conn *websocket.Conn
			It("Should use the websocketserver to serve http", func() {
				httpserver = httptest.NewServer(wsserver)
			})

			It("Should accept connections", func() {
				dialer := websocket.Dialer{}
				url := fmt.Sprintf("ws://%s", strings.Split(httpserver.URL, "//")[1])
				c, _, err := dialer.Dial(url, http.Header{})
				conn = c
				Ω(err).ShouldNot(HaveOccurred())
			})

			byteMessage := []byte("TST TST test")
			It("Should accept messages", func() {
				err := conn.WriteMessage(websocket.TextMessage, byteMessage)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Should return answers", func() {
				_, message, err := conn.ReadMessage()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(message).Should(BeEquivalentTo(byteMessage))
			})

			It("Should return an error if decoding the message fails", func() {
				conn.WriteMessage(websocket.TextMessage, []byte("fail"))
				_, message, _ := conn.ReadMessage()
				Ω(message).Should(BeEquivalentTo("service fail doesnt exist"))
			})
		})
	})

	Context("BasicHandler", func() {
		h := ws.BasicHandler{}
		It("Should Decode messages", func() {
			message := []byte("TSTTST")
			path := "test"
			p, _ := proto.Marshal(&pb.AddKMIRequest{
				Path: path,
			})
			message = append(message, p...)

			_, _, data, err := h.Decode(message)
			Ω(err).ShouldNot(HaveOccurred())
			q := &pb.AddKMIRequest{}
			proto.Unmarshal(data.([]byte), q)
			Ω(q.Path).Should(BeEquivalentTo(path))
		})

		It("Should return an error if message is malformatted", func() {
			_, _, _, err := h.Decode([]byte{})
			Ω(err).Should(HaveOccurred())
		})

		It("Should encode messages", func() {
			tst := ws.ProtoIDFromString("TST")
			message, err := h.Encode(&tst, &tst, &pb.AddKMIRequest{
				Path: "path",
			})
			Ω(err).ShouldNot(HaveOccurred())
			srv, me, data, _ := h.Decode(message)
			Ω(srv.String()).Should(BeEquivalentTo(tst.String()))
			Ω(me.String()).Should(BeEquivalentTo(tst.String()))
			p, _ := proto.Marshal(&pb.AddKMIRequest{
				Path: "path",
			})
			Ω(data).Should(BeEquivalentTo(p))
		})
	})
})
