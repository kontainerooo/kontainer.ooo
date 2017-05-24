package firewall_test

import (
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/firewall"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = XDescribe("Firewall", func() {
	Describe("Create Service", func() {
		It("Should create a new service", func() {
			mockIpt, _ := testutils.NewMockIPTService()

			fws, err := firewall.NewService(mockIpt)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(fws).ShouldNot(BeNil())
		})
	})

	Describe("Init bridge", func() {
		It("Should set up rules for bridge initialisation", func() {
			mockIpt, _ := testutils.NewMockIPTService()

			fws, _ := firewall.NewService(mockIpt)

			ip, _ := abstraction.NewInet("172.18.0.0/16")
			err := fws.InitBridge(ip, "br-084d60eeada1")

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("Allow connection", func() {
		It("Should allow connection", func() {
			mockIpt, _ := testutils.NewMockIPTService()

			fws, _ := firewall.NewService(mockIpt)

			ip1, _ := abstraction.NewInet("172.18.0.0/16")
			ip2, _ := abstraction.NewInet("172.18.0.0/16")

			err := fws.AllowConnection(ip1, "br-0815", ip2, "br-0815")

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("Block connection", func() {
		It("Should block a connection", func() {
			mockIpt, _ := testutils.NewMockIPTService()

			fws, _ := firewall.NewService(mockIpt)

			ip1, _ := abstraction.NewInet("172.18.0.0/16")
			ip2, _ := abstraction.NewInet("172.18.0.0/16")

			fws.AllowConnection(ip1, "br-0815", ip2, "br-0815")

			err := fws.BlockConnection(ip1, "br-0815", ip2, "br-0815")

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("Allow port", func() {
		It("Should allow a port", func() {
			mockIpt, _ := testutils.NewMockIPTService()

			fws, _ := firewall.NewService(mockIpt)

			ip1, _ := abstraction.NewInet("172.18.0.0/16")
			ip2, _ := abstraction.NewInet("172.18.0.0/16")

			err := fws.AllowPort(ip1, "br-0815", ip2, "br-0815", 8080, "tcp")

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error on invalid protocol", func() {
			mockIpt, _ := testutils.NewMockIPTService()

			fws, _ := firewall.NewService(mockIpt)

			ip1, _ := abstraction.NewInet("172.18.0.0/16")
			ip2, _ := abstraction.NewInet("172.18.0.0/16")

			err := fws.AllowPort(ip1, "br-0815", ip2, "br-0815", 8080, "DPD")

			Ω(err).Should(HaveOccurred())
		})

		It("Should error on existing rule", func() {
			mockIpt, _ := testutils.NewMockIPTService()

			fws, _ := firewall.NewService(mockIpt)

			ip1, _ := abstraction.NewInet("172.18.0.0/16")
			ip2, _ := abstraction.NewInet("172.18.0.0/16")

			fws.AllowPort(ip1, "br-0815", ip2, "br-0815", 8080, "udp")
			err := fws.AllowPort(ip1, "br-0815", ip2, "br-0815", 8080, "udp")

			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Block port", func() {
		It("Should block a port", func() {
			mockIpt, _ := testutils.NewMockIPTService()

			fws, _ := firewall.NewService(mockIpt)

			ip1, _ := abstraction.NewInet("172.18.0.0/16")
			ip2, _ := abstraction.NewInet("172.18.0.0/16")

			fws.AllowPort(ip1, "br-0815", ip2, "br-0815", 8080, "tcp")

			err := fws.BlockPort(ip1, "br-0815", ip2, "br-0815", 8080, "tcp")

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error on invalid protocol", func() {
			mockIpt, _ := testutils.NewMockIPTService()

			fws, _ := firewall.NewService(mockIpt)

			ip1, _ := abstraction.NewInet("172.18.0.0/16")
			ip2, _ := abstraction.NewInet("172.18.0.0/16")

			fws.AllowPort(ip1, "br-0815", ip2, "br-0815", 8080, "tcp")

			err := fws.BlockPort(ip1, "br-0815", ip2, "br-0815", 8080, "DPD")

			Ω(err).Should(HaveOccurred())
		})

		It("Should error non-existing rule", func() {
			mockIpt, _ := testutils.NewMockIPTService()

			fws, _ := firewall.NewService(mockIpt)

			ip1, _ := abstraction.NewInet("172.18.0.0/16")
			ip2, _ := abstraction.NewInet("172.18.0.0/16")

			err := fws.BlockPort(ip1, "br-0815", ip2, "br-0815", 8080, "DPD")

			Ω(err).Should(HaveOccurred())
		})
	})
})
