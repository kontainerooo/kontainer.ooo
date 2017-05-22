package websocket_test

import (
	"context"
	"errors"

	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	Context("Service Endpoint", func() {
		Describe("New Service Endpoint", func() {
			It("Should create a new Endpoint", func() {
				e, err := ws.NewServiceEndpoint("name", ws.ProtoIDFromString("TST"), makeTestEndpoint(), decodeTest, encodeTest)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(e).ShouldNot(BeNil())
			})

			Context("Error Handling", func() {
				It("Should return an error if no name was provided", func() {
					_, err := ws.NewServiceEndpoint("", ws.ProtoIDFromString("TST"), makeTestEndpoint(), decodeTest, encodeTest)
					Ω(err).Should(BeEquivalentTo(ws.ErrNoName))
				})

				It("Should return an error if the protocol id is invalid", func() {
					_, err := ws.NewServiceEndpoint("name", ws.ProtoIDFromString(""), makeTestEndpoint(), decodeTest, encodeTest)
					Ω(err).Should(BeEquivalentTo(ws.ErrInvalidProtoID))
				})

				It("Should return an error if no endpoint was provided", func() {
					_, err := ws.NewServiceEndpoint("name", ws.ProtoIDFromString("TST"), nil, decodeTest, encodeTest)
					Ω(err).Should(BeEquivalentTo(ws.ErrNoEndpoint))
				})

				It("Should set a standard decode function if no one was provided", func() {
					e, err := ws.NewServiceEndpoint("name", ws.ProtoIDFromString("TST"), makeTestEndpoint(), nil, encodeTest)
					Ω(err).ShouldNot(HaveOccurred())
					Ω(e.Dec).ShouldNot(BeNil())
				})

				It("Should set a standard encode function if no one was provided", func() {
					e, err := ws.NewServiceEndpoint("name", ws.ProtoIDFromString("TST"), makeTestEndpoint(), decodeTest, nil)
					Ω(err).ShouldNot(HaveOccurred())
					Ω(e.Enc).ShouldNot(BeNil())
				})
			})
		})

		Context("Helper", func() {
			It("StDencode should return its interface parameter", func() {
				s := "test"
				i, err := ws.StdDencode(context.TODO(), s)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(i).Should(BeEquivalentTo(interface{}(s)))
			})
		})
	})

	Context("Service Description", func() {
		Describe("New Service Description", func() {
			It("Should return a new Description", func() {
				sd, err := ws.NewServiceDescription("name", ws.ProtoIDFromString("TST"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(sd).ShouldNot(BeNil())
			})

			Context("Error Handling", func() {
				It("Should return an error if no name was provided", func() {
					_, err := ws.NewServiceDescription("", ws.ProtoIDFromString("TST"))
					Ω(err).Should(BeEquivalentTo(ws.ErrNoName))
				})

				It("Should return an error if the protocol id is invalid", func() {
					_, err := ws.NewServiceDescription("name", ws.ProtoIDFromString(""))
					Ω(err).Should(BeEquivalentTo(ws.ErrInvalidProtoID))
				})
			})
		})

		Describe("Add Endpoint", func() {
			It("Should add a new Endpoint to the Service description", func() {
				sd, _ := ws.NewServiceDescription("name", ws.ProtoIDFromString("TST"))
				e, _ := ws.NewServiceEndpoint("name", ws.ProtoIDFromString("TST"), makeTestEndpoint(), decodeTest, encodeTest)
				err := sd.AddEndpoint(e)
				Ω(err).ShouldNot(HaveOccurred())
			})

			Context("Error Handling", func() {
				It("Should return an error if an endpoint id already exists", func() {
					sd, _ := ws.NewServiceDescription("name", ws.ProtoIDFromString("TST"))
					e, _ := ws.NewServiceEndpoint("name", ws.ProtoIDFromString("TST"), makeTestEndpoint(), decodeTest, encodeTest)

					sd.AddEndpoint(e)

					err := sd.AddEndpoint(e)
					Ω(err).Should(HaveOccurred())
				})

				It("Should return an error if it is called with one", func() {
					testErr := errors.New("test")
					sd, _ := ws.NewServiceDescription("name", ws.ProtoIDFromString("TST"))

					err := sd.AddEndpoint(nil, testErr)
					Ω(err).Should(BeEquivalentTo(testErr))
				})
			})
		})

		Describe("Get Endpoint Handler", func() {
			It("Should return an Endpoint Handler", func() {
				protoID := ws.ProtoIDFromString("TST")
				sd, _ := ws.NewServiceDescription("name", protoID)
				e, _ := ws.NewServiceEndpoint("name", protoID, makeTestEndpoint(), decodeTest, encodeTest)
				sd.AddEndpoint(e)

				eh, err := sd.GetEndpointHandler(protoID, nil, nil)
				Ω(eh).ShouldNot(BeNil())
				Ω(err).ShouldNot(HaveOccurred())
			})

			Context("Endpoint Handler", func() {
				It("Should", func() {
					protoID := ws.ProtoIDFromString("TST")
					sd, _ := ws.NewServiceDescription("name", protoID)
					e, _ := ws.NewServiceEndpoint("name", protoID, makeTestEndpoint(), decodeTest, encodeTest)
					sd.AddEndpoint(e)
					eh, _ := sd.GetEndpointHandler(protoID, nil, nil)

					val := 0
					req := request{val}
					res, err := eh(req)

					Ω(err).ShouldNot(HaveOccurred())
					Ω(res.(response).res).Should(BeEquivalentTo(val))
				})

				It("Should", func() {
					protoID := ws.ProtoIDFromString("TST")
					sd, _ := ws.NewServiceDescription("name", protoID)
					e, _ := ws.NewServiceEndpoint("name", protoID, makeTestEndpoint(), decodeTest, encodeTest)
					sd.AddEndpoint(e)
					eh, _ := sd.GetEndpointHandler(protoID, nil, nil)

					val := true
					req := request{val}
					_, err := eh(req)

					Ω(err).Should(BeEquivalentTo(errDecode))
				})

				It("Should", func() {
					protoID := ws.ProtoIDFromString("TST")
					sd, _ := ws.NewServiceDescription("name", protoID)
					e, _ := ws.NewServiceEndpoint("name", protoID, makeTestEndpoint(), decodeTest, encodeTest)
					sd.AddEndpoint(e)
					eh, _ := sd.GetEndpointHandler(protoID, nil, nil)

					val := errEndpoint
					req := request{val}
					_, err := eh(req)

					Ω(err).Should(BeEquivalentTo(errEndpoint))
				})

				It("Should", func() {
					protoID := ws.ProtoIDFromString("TST")
					sd, _ := ws.NewServiceDescription("name", protoID)
					e, _ := ws.NewServiceEndpoint("name", protoID, makeTestEndpoint(), decodeTest, encodeTest)
					sd.AddEndpoint(e)
					eh, _ := sd.GetEndpointHandler(protoID, nil, nil)

					val := uint64(1)
					req := request{val}
					_, err := eh(req)

					Ω(err).Should(BeEquivalentTo(errEncode))
				})
			})

			Context("Error Handling", func() {
				It("Should return an error if the requested endpoint does not exist", func() {
					sd, _ := ws.NewServiceDescription("name", ws.ProtoIDFromString("TST"))
					_, err := sd.GetEndpointHandler(ws.ProtoIDFromString("TST"), nil, nil)
					Ω(err).Should(HaveOccurred())
				})
			})
		})
	})
})
