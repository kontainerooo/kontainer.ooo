package kmi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKmi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kmi Suite")
}
