package e2e_tests

import (
	"testing"

	"github.com/joho/godotenv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "e2e test Suite")
}

func init() {
	_ = godotenv.Load("../../.env")

}
