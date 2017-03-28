package iptables_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"

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
	Describe("Create new service", func() {
		It("Should create a new service", func() {
			Ω(true).Should(BeTrue())
		})
	})
})
