package iptables_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/iptables"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"

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

func simpleNewInet(s string) abstraction.Inet {
	v, _ := abstraction.NewInet(s)
	return v
}

var _ = Describe("Iptables", func() {

	Describe("Redirect rule string", func() {
		It("Should create REDIRECT rule", func() {
			r := iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
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
				Destination:     simpleNewInet("127.0.0.2"),
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
				Destination:     simpleNewInet("127.0.0.2"),
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
				Destination: simpleNewInet("127.0.0.2"),
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

		It("Should error on wrong protocol", func() {
			r := iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "mia",
				SourcePort:      8080,
				Destination:     simpleNewInet("127.0.0.2"),
				DestinationPort: 80,
			}

			str, err := r.ToString()

			Ω(err).Should(HaveOccurred())
			Ω(err).Should(Equal(iptables.ErrWrongProtocol))
			Ω(str).Should(BeEmpty())
		})
	})

	Describe("Accept/Drop rule string", func() {
		It("Should create an accept rule", func() {
			r := iptables.Rule{
				Target:          "ACCEPT",
				Chain:           "INPUT",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				DestinationPort: 80,
			}

			str, err := r.ToString()

			Ω(err).ShouldNot(HaveOccurred())
			Ω(str).Should(Equal("-A INPUT -j ACCEPT --dst 127.0.0.2 -p tcp --dport 80"))
		})

		It("Should create an accept rule with input interface", func() {
			r := iptables.Rule{
				Target:          "ACCEPT",
				Chain:           "INPUT",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				DestinationPort: 80,
				In:              "eth0",
			}

			str, err := r.ToString()

			Ω(err).ShouldNot(HaveOccurred())
			Ω(str).Should(Equal("-A INPUT -j ACCEPT --dst 127.0.0.2 -p tcp --dport 80 -i eth0"))
		})

		It("Should create an accept rule with output interface", func() {
			r := iptables.Rule{
				Target:          "ACCEPT",
				Chain:           "INPUT",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				DestinationPort: 80,
				Out:             "eth0",
			}

			str, err := r.ToString()

			Ω(err).ShouldNot(HaveOccurred())
			Ω(str).Should(Equal("-A INPUT -j ACCEPT --dst 127.0.0.2 -p tcp --dport 80 -o eth0"))
		})

		It("Should create an accept rule with state", func() {
			r := iptables.Rule{
				Target:          "ACCEPT",
				Chain:           "INPUT",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				DestinationPort: 80,
				State:           "RELATED,ESTABLISHED",
			}

			str, err := r.ToString()

			Ω(err).ShouldNot(HaveOccurred())
			Ω(str).Should(Equal("-A INPUT -j ACCEPT --dst 127.0.0.2 -p tcp --dport 80 -m state --state RELATED,ESTABLISHED"))
		})

		It("Should error on wrong chain", func() {
			r := iptables.Rule{
				Target:          "ACCEPT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				DestinationPort: 80,
			}

			str, err := r.ToString()

			Ω(err).Should(HaveOccurred())
			Ω(str).Should(BeEmpty())
		})

		It("Should error on missing destination", func() {
			r := iptables.Rule{
				Target:          "ACCEPT",
				Chain:           "INPUT",
				Protocol:        "tcp",
				DestinationPort: 80,
			}

			str, err := r.ToString()

			Ω(err).Should(HaveOccurred())
			Ω(str).Should(BeEmpty())
		})
	})

	Describe("Unknown rule string", func() {
		It("Should error on unknown target", func() {
			r := iptables.Rule{
				Target:          "Test",
				Chain:           "INPUT",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				DestinationPort: 80,
			}

			str, err := r.ToString()

			Ω(err).Should(HaveOccurred())
			Ω(str).Should(BeEmpty())
		})
	})

	Describe("Rule hash", func() {
		It("Should return hash of rule", func() {
			r := iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				SourcePort:      8080,
				DestinationPort: 80,
			}

			hash := r.GetHash()
			Ω(hash).Should(Equal("dc5ae82240b7db6c9458c34cd81373755046d490ea6e58dec35fba9e2b63f520"))
		})

		It("Should not return hash on invalid rule", func() {
			r := iptables.Rule{
				Target:          "fail",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				SourcePort:      8080,
				DestinationPort: 80,
			}

			hash := r.GetHash()
			Ω(hash).Should(BeEmpty())
		})
	})

	iptables.ExecCommand = fakeExecCommand
	Describe("New Service", func() {
		It("Should create a new service", func() {
			ipts, err := iptables.NewService("iptables", testutils.NewMockDB())

			Ω(err).ShouldNot(HaveOccurred())
			Ω(ipts).ShouldNot(BeNil())
		})

		It("Should error when database connection fails", func() {
			db := testutils.NewMockDB()
			db.SetError(1)
			ipts, err := iptables.NewService("iptables", db)

			Ω(err).Should(HaveOccurred())
			Ω(ipts).Should(BeNil())
		})

		It("Should error when iptables cannot be executed", func() {
			iptablesIsPresent = 0
			ipts, err := iptables.NewService("iptables", testutils.NewMockDB())

			Ω(err).Should(HaveOccurred())
			Ω(ipts).Should(BeNil())
			iptablesIsPresent = 1
		})
	})

	Describe("Add Rules", func() {
		It("Should add a rule", func() {
			ipts, _ := iptables.NewService("iptables", testutils.NewMockDB())

			err := ipts.AddRule(123, iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				SourcePort:      8080,
				DestinationPort: 80,
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error on invalid rule", func() {
			ipts, _ := iptables.NewService("iptables", testutils.NewMockDB())

			err := ipts.AddRule(123, iptables.Rule{
				Target:          "FAIL",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				SourcePort:      8080,
				DestinationPort: 80,
			})

			Ω(err).Should(HaveOccurred())
		})

		It("Should error on existing rule", func() {
			ipts, _ := iptables.NewService("iptables", testutils.NewMockDB())

			ipts.AddRule(123, iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				SourcePort:      8080,
				DestinationPort: 80,
			})

			err := ipts.AddRule(123, iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				SourcePort:      8080,
				DestinationPort: 80,
			})

			Ω(err).Should(HaveOccurred())
		})

		It("Should error on iptables execution", func() {
			ipts, _ := iptables.NewService("iptables", testutils.NewMockDB())

			iptablesIsPresent = 0

			err := ipts.AddRule(123, iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				SourcePort:      8080,
				DestinationPort: 80,
			})

			Ω(err).Should(HaveOccurred())

			iptablesIsPresent = 1
		})
	})

	Describe("Delete Rules", func() {
		hash := ""
		It("Should delete a rule", func() {
			ipts, _ := iptables.NewService("iptables", testutils.NewMockDB())
			r := iptables.Rule{
				Target:          "REDIRECT",
				Chain:           "PREROUTING",
				Protocol:        "tcp",
				Destination:     simpleNewInet("127.0.0.2"),
				SourcePort:      8080,
				DestinationPort: 80,
			}

			ipts.AddRule(123, r)

			hash = r.GetHash()
			err := ipts.RemoveRule(hash)

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error on non-existing rule", func() {
			ipts, _ := iptables.NewService("iptables", testutils.NewMockDB())
			err := ipts.RemoveRule(hash)

			Ω(err).Should(HaveOccurred())
		})
	})
})
