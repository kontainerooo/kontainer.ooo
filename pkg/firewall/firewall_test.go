package firewall_test

import (
	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/firewall"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Firewall", func() {
	Describe("Create Service", func() {
		It("Should create a new service", func() {
			mockIpt := testutils.NewMockIPTClient()
			mockIptEndpoints := testutils.NewMockIPTEndpoints(log.NewNopLogger(), *mockIpt)

			fws := firewall.NewService(mockIptEndpoints)

			Ω(fws).ShouldNot(BeNil())
		})
	})

	Describe("Init bridge", func() {
		It("Should set up rules for bridge initialisation", func() {
			mockIpt := testutils.NewMockIPTClient()
			mockIptEndpoints := testutils.NewMockIPTEndpoints(log.NewNopLogger(), *mockIpt)
			fws := firewall.NewService(mockIptEndpoints)

			ip, _ := abstraction.NewInet("0.0.0.0/0")
			err := fws.InitBridge(ip, "br-ca8015388757")

			Ω(err).ShouldNot(HaveOccurred())
			mockIpt.ListRuleStrings()
		})
	})
})
