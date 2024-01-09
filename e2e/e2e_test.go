package e2e

import (
	"testing"

	"example.com/conformance/testcases"
	"github.com/katanomi/pkg/testing/framework"
)

func TestE2E(t *testing.T) {
	testcases.CaseAuthentication.Do()

	// start step to run e2e
	framework.New("core-e2e").Run(t)
}
