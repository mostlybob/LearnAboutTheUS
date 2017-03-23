package quiz_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLearnAboutTheUS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LearnAboutTheUS Suite")
}
