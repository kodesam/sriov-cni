package providers

import (
	"testing"

	"github.com/k8snetworkplumbingwg/sriov-cni/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Providers Suite")
}

var _ = BeforeSuite(func() {
	// create test sys tree
	err := utils.CreateTmpSysFs()
	check(err)
})

var _ = AfterSuite(func() {
	var err error
	err = utils.RemoveTmpSysFs()
	check(err)
})