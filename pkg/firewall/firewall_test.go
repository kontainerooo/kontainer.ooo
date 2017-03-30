package firewall_test

import (
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/firewall"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Firewall", func() {
	Describe("Create Service", func() {
		It("Should create a new service", func() {
			mockIpt := testutils.NewMockIPTService()

			fws, err := firewall.NewService(mockIpt)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(fws).ShouldNot(BeNil())
		})
	})

	Describe("Init bridge", func() {
		It("Should set up rules for bridge initialisation", func() {
			mockIpt := testutils.NewMockIPTService()

			fws, _ := firewall.NewService(mockIpt)

			ip, _ := abstraction.NewInet("0.0.0.0/0")
			err := fws.InitBridge(ip, "br-084d60eeada1")

			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
