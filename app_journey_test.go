// +build journeyTests

package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
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
