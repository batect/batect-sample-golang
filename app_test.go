// +build unitTests

package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
