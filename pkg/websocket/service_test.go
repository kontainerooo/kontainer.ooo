package websocket_test

import (
	"context"

	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	Context("Service Endpoint", func() {
		It("Should create a new Endpoint", func() {
			e, err := ws.NewServiceEndpoint("name", ws.ProtoIDFromString("TST"), makeTestEndpoint(), decodeTest, encodeTest)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(e).ShouldNot(BeNil())
		})

		Context("Helper", func() {
			It("StDencode should return its interface parameter", func() {
				s := "test"
				i, err := ws.StdDencode(context.TODO(), s)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(i).Should(BeEquivalentTo(interface{}(s)))
			})
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
})
