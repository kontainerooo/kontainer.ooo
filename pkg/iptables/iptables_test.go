package iptables_test

import (
	"os"
	"os/exec"

	"github.com/kontainerooo/kontainer.ooo/pkg/iptables"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

var _ = Describe("Iptables", func() {
	iptables.ExecCommand = fakeExecCommand
	Describe("New Service", func() {

		It("Should create a new service", func() {
			ipts, err := iptables.NewService("iptables")

			Ω(err).ShouldNot(HaveOccurred())
			Ω(ipts).ShouldNot(BeNil())
		})
	})
})
