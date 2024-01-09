package conformance_e2e

import (
	"testing"

	// "example.com/conformance/testcases"
	"github.com/katanomi/pkg/testing/framework"
	"github.com/katanomi/pkg/testing/framework/conformance"
)

var fmw = framework.New("conformance-e2e")

func TestMain(m *testing.M) {
	fmw.Config(conformance.Configure).MRun(m)
}

func TestConformance(t *testing.T) {
	fmw.Run(t)
}
