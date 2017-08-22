package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"flag"
	"github.com/onsi/ginkgo/reporters"
	"testing"
)

var junitPath = flag.String("junit", "", "filename to write JUnit XML too")

func TestGoto(t *testing.T) {
	RegisterFailHandler(Fail)
	if *junitPath == "" {
		RunSpecs(t, "")
	} else {
		junitReporter := reporters.NewJUnitReporter(*junitPath)
		RunSpecsWithDefaultAndCustomReporters(t, "yml2env Suite", []Reporter{junitReporter})
	}
}
