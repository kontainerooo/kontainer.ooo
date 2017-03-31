package firewall_test

import (
	"os"
	"testing"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/firewall"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TestHelperProcess Mocks the iptables binary
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	os.Exit(0)
}

var _ = Describe("Firewall", func() {
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
})
