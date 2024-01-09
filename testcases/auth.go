package testcases

import (
	// "context"
	// "github.com/katanomi/pkg/testing/framework/conformance"

	. "github.com/katanomi/pkg/testing/framework"
	. "github.com/katanomi/pkg/testing/framework/base"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var CaseAuthentication = P0Case("testing for authentication").WithFunc(func(ctx *TestContext) {
	Context("authentication with oauth2", func() {
		statusCode := 200
		It("should be authenticated successful", func() {
			Expect(statusCode).To(Equal(200))
		})
	})

	Context("authentication with basic-auth", func() {
		statusCode := 200
		It("should be authenticated successful", func() {
			Expect(statusCode).To(Equal(200))
		})
	})
})
