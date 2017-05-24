package websocket_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/websocket"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	Context("Middleware", func() {
		It("Should return a Before Middleware", func() {
			bm := ws.Before(func(ws.ProtoID, ws.ProtoID, interface{}, interface{}) error { return nil })
			Ω(bm).ShouldNot(BeNil())
		})

		It("Should return an After Middleware", func() {
			bm := ws.After(func(ws.ProtoID, ws.ProtoID, interface{}, interface{}) error { return nil })
			Ω(bm).ShouldNot(BeNil())
		})
	})

	Context("Server", func() {
		Describe("New Server", func() {
			It("Should return a new server", func() {
				server := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{}, testAuth{}, ws.SSLConfig{})
				Ω(server).ShouldNot(BeNil())
			})

			It("Should accept middleware", func() {
				server := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{}, testAuth{}, ws.SSLConfig{},
					ws.Before(func(ws.ProtoID, ws.ProtoID, interface{}, interface{}) error { return nil }),
					ws.After(func(ws.ProtoID, ws.ProtoID, interface{}, interface{}) error { return nil }))
				Ω(server).ShouldNot(BeNil())
			})

			Context("Error Handling", func() {
				Describe("Upgrader", func() {
					It("Should set a ReadBuffersize if none is provided", func() {
						server := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{}, testAuth{}, ws.SSLConfig{})
						Ω(server.Upgrader.ReadBufferSize).ShouldNot(BeNil())
					})

					It("Should set the ReadBuffersize to the value of WriteBufferSize if that is not 0", func() {
						size := 123
						server := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{
							WriteBufferSize: size,
						}, testAuth{}, ws.SSLConfig{})
						Ω(server.Upgrader.ReadBufferSize).Should(BeEquivalentTo(size))
					})

					It("Should set the WriteBufferSize to the value of ReadBufferSize", func() {
						size := 123
						server := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{
							ReadBufferSize: size,
						}, testAuth{}, ws.SSLConfig{})
						Ω(server.Upgrader.WriteBufferSize).Should(BeEquivalentTo(size))
					})

					It("Should provide a default CheckOrigin function", func() {
						server := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{}, testAuth{}, ws.SSLConfig{})
						Ω(server.Upgrader.CheckOrigin(nil)).Should(BeTrue())
					})
				})
			})
		})

		Describe("Register Service", func() {
			It("Should register a new Service", func() {
				server := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{}, testAuth{}, ws.SSLConfig{})
				sd, _ := ws.NewServiceDescription("name", ws.ProtoIDFromString("TST"))

				err := server.RegisterService(sd)
				Ω(err).ShouldNot(HaveOccurred())
			})

			Context("Error Handling", func() {
				It("Should return an error if a service id already exists", func() {
					server := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{}, testAuth{}, ws.SSLConfig{})
					sd, _ := ws.NewServiceDescription("name", ws.ProtoIDFromString("TST"))

					server.RegisterService(sd)
					err := server.RegisterService(sd)
					Ω(err).Should(HaveOccurred())
				})
			})
		})

		Describe("Get Service", func() {
			It("Should return a Service", func() {
				id := ws.ProtoIDFromString("TST")
				server := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{}, testAuth{}, ws.SSLConfig{})
				sd, _ := ws.NewServiceDescription("name", id)
				server.RegisterService(sd)

				sd, err := server.GetService(id)
				Ω(sd).ShouldNot(BeNil())
				Ω(err).ShouldNot(HaveOccurred())
			})

			Context("Error Handling", func() {
				It("Should return an error if a service does not exist", func() {
					id := ws.ProtoIDFromString("TST")
					server := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{}, testAuth{}, ws.SSLConfig{})

					_, err := server.GetService(id)
					Ω(err).Should(HaveOccurred())
				})
			})
		})

		Describe("Serve", func() {
			Context("Error Handling", func() {
				It("Should return an error if no server has been started", func() {
					server := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{}, testAuth{}, ws.SSLConfig{
						Only: true,
					})
					err := server.Serve("")
					Ω(err).Should(HaveOccurred())
				})
			})
		})

		Describe("Serve HTTP", func() {

			It("Should accept connections", func() {
				wsServer := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{}, testAuth{}, ws.SSLConfig{})
				httpServer := httptest.NewServer(wsServer)

				dialer := websocket.Dialer{}
				url := fmt.Sprintf("ws://%s", strings.Split(httpServer.URL, "//")[1])
				c, _, err := dialer.Dial(url, http.Header{})
				Ω(c).ShouldNot(BeNil())
				Ω(err).ShouldNot(HaveOccurred())
			})

			Context("Connection Handling", func() {
				var (
					testMsg    = []byte("TST TST test")
					connection *websocket.Conn
					httpServer *httptest.Server
				)

				BeforeEach(func() {
					wsServer := ws.NewServer(protocolMap, log.NewNopLogger(), websocket.Upgrader{}, testAuth{}, ws.SSLConfig{})

					sd, _ := ws.NewServiceDescription("test", ws.ProtoIDFromString("TST"))
					sd.AddEndpoint(ws.NewServiceEndpoint("test", ws.ProtoIDFromString("TST"), makeTestEndpoint(), decodeTest, encodeTest))
					wsServer.RegisterService(sd)

					httpServer = httptest.NewServer(wsServer)

					dialer := websocket.Dialer{}
					url := fmt.Sprintf("ws://%s", strings.Split(httpServer.URL, "//")[1])
					connection, _, _ = dialer.Dial(url, http.Header{})
				})

				AfterEach(func() {
					connection.Close()
					httpServer.Close()
					httpServer, connection = nil, nil
				})

				It("Should accept messages", func() {
					err := connection.WriteMessage(websocket.TextMessage, testMsg)
					Ω(err).ShouldNot(HaveOccurred())
				})

				It("Should return answers", func() {
					connection.WriteMessage(websocket.TextMessage, testMsg)
					_, msg, err := connection.ReadMessage()
					Ω(err).ShouldNot(HaveOccurred())
					Ω(msg).Should(BeEquivalentTo(testMsg))
				})

				Context("Error Handling", func() {
					It("Should return an error message if a Service does not exist", func() {
						connection.WriteMessage(websocket.TextMessage, []byte("BLA BLA bla"))
						_, msg, _ := connection.ReadMessage()
						Ω(msg).Should(BeEquivalentTo([]byte(errService.Error())))
					})

					It("Should return an error message if a Method does not exist", func() {
						connection.WriteMessage(websocket.TextMessage, []byte("TST BLA bla"))
						_, msg, _ := connection.ReadMessage()
						Ω(msg).Should(BeEquivalentTo([]byte(errMethod.Error())))
					})

					It("Should return an error if a Service is not registered", func() {
						connection.WriteMessage(websocket.TextMessage, []byte("NYI NYI bla"))
						_, msg, _ := connection.ReadMessage()
						Ω(msg).Should(BeEquivalentTo([]byte("Service Description NYI does not exist")))
					})

					It("Should return an error if a Method is not registered", func() {
						connection.WriteMessage(websocket.TextMessage, []byte("TST NYI bla"))
						_, msg, _ := connection.ReadMessage()
						Ω(msg).Should(BeEquivalentTo([]byte("Service Endpoint NYI does not exist")))
					})

					It("Should return an error if the Endpoint handler returns one", func() {
						connection.WriteMessage(websocket.TextMessage, []byte("TST TST 42"))
						_, msg, _ := connection.ReadMessage()
						Ω(msg).Should(BeEquivalentTo([]byte(errEncode.Error())))
					})

					It("Should return an error if the Protocol Encoder returns one", func() {
						connection.WriteMessage(websocket.TextMessage, []byte("TST TST error"))
						_, msg, _ := connection.ReadMessage()
						Ω(msg).Should(BeEquivalentTo([]byte(errProtocolEncode.Error())))
					})

					XContext("Middleware", func() {})
				})
			})

			Context("Error Handling", func() {
				XIt("Should return an error if the requested protocol does not exist", func() {
				})
			})
		})
	})
})
