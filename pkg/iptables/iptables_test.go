package iptables_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/kontainerooo/kontainer.ooo/pkg/iptables"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var iptablesIsPresent = 1

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", fmt.Sprintf("GO_IPT_IS_PRESENT=%d", iptablesIsPresent)}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	if os.Getenv("GO_IPT_IS_PRESENT") == "1" {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

var _ = Describe("Iptables", func() {

	Describe("Redirect rule string", func() {
		It("Should create REDIRECT rule", func() {
			r := iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     "127.0.0.2",
				SourcePort:      8080,
				DestinationPort: 80,
			}

			str, err := r.ToString()

			Ω(err).ShouldNot(HaveOccurred())
			Ω(str).Should(Equal("-t nat -A PREROUTING --dst 127.0.0.2 -p tcp --dport 8080 -j REDIRECT --to-port 80"))
		})

		It("Should error on wrong chain", func() {
			r := iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "INPUT",
				Protocol:        "tcp",
				Destination:     "127.0.0.2",
				SourcePort:      8080,
				DestinationPort: 80,
			}

			str, err := r.ToString()

			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(iptables.ErrWrongChain))
			Ω(str).Should(BeEmpty())
		})

		It("Should error on no source port", func() {
			r := iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     "127.0.0.2",
				DestinationPort: 80,
			}

			str, err := r.ToString()

			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(iptables.ErrNoPorts))
			Ω(str).Should(BeEmpty())
		})

		It("Should error on no dest port", func() {
			r := iptables.Rule{
				Target:      "REDIRECT",
				Chain:       "PREROUTING",
				Protocol:    "tcp",
				Destination: "127.0.0.2",
				SourcePort:  8080,
			}

			str, err := r.ToString()

			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(iptables.ErrNoPorts))
			Ω(str).Should(BeEmpty())
		})

		It("Should error on missing destionation", func() {
			r := iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				SourcePort:      8080,
				DestinationPort: 80,
			}

			str, err := r.ToString()

			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(iptables.ErrNoDestination))
			Ω(str).Should(BeEmpty())
		})

		It("Should error when destionation is not an ip address", func() {
			r := iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				SourcePort:      8080,
				Destination:     "google.com",
				DestinationPort: 80,
			}

			str, err := r.ToString()

			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(iptables.ErrIPWrongFormat))
			Ω(str).Should(BeEmpty())
		})

		It("Should error on wrong protocol", func() {
			r := iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "mia",
				SourcePort:      8080,
				Destination:     "127.0.0.2",
				DestinationPort: 80,
			}

			str, err := r.ToString()

			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(iptables.ErrWrongProtocol))
			Ω(str).Should(BeEmpty())
		})
	})

	iptables.ExecCommand = fakeExecCommand
	Describe("New Service", func() {
		It("Should create a new service", func() {
			ipts, err := iptables.NewService("iptables")

			Ω(err).ShouldNot(HaveOccurred())
			Ω(ipts).ShouldNot(BeNil())
		})

		It("Should error when iptables cannot be executed", func() {
			iptablesIsPresent = 0
			ipts, err := iptables.NewService("iptables")

			Ω(err).Should(HaveOccurred())
			Ω(ipts).Should(BeNil())
		})
	})
})
