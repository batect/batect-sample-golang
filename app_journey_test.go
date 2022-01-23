//go:build journeyTests
// +build journeyTests

package main

import (
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("The application", func() {
	Describe("pinging", func() {
		It("should return the expected response", func() {
			resp, err := http.Get("http://app:8080/ping")

			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			body, _ := ioutil.ReadAll(resp.Body)

			Expect(string(body)).To(Equal("pong"))
		})
	})
})
