package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPractice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Practice Suite")
}
