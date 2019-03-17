package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"

	"github.com/onsi/ginkgo/reporters"
	"testing"
)

func TestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "App Suite", createReporters())
}

func createReporters() []Reporter {
	path := os.Getenv("REPORT_PATH")

	if path == "" {
		return []Reporter{}
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		panic(err)
	}

	junitReporter := reporters.NewJUnitReporter(path + "/report.xml")

	return []Reporter{junitReporter}
}
