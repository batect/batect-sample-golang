//go:build unitTests
// +build unitTests

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("The application", func() {
	Describe("when the ping endpoint is called", func() {
		var response *http.Response

		BeforeEach(func() {
			recorder := httptest.NewRecorder()

			handlePing(recorder, nil)

			response = recorder.Result()
		})

		It("should return 'pong' as the response", func() {
			body, _ := ioutil.ReadAll(response.Body)

			Expect(string(body)).To(Equal("pong"))
		})

		It("should return a HTTP 200 status code", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})
	})
})
