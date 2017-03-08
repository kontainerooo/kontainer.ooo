package customercontainer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCustomercontainer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Customercontainer Suite")
}
