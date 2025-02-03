package sm_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sm Suite")
}
