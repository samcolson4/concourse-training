package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"flag"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var fixturesPath = flag.String("fixtures", "", "path of test fixtures")

var _ = Describe("yml2env", func() {
	var cliPath string
	usage := "yml2env <YAML file> <command>"

	BeforeSuite(func() {
		var err error
		cliPath, err = Build("github.com/EngineerBetter/yml2env")
		Ω(err).ShouldNot(HaveOccurred())

		rand.Seed(time.Now().UnixNano())
		_, wasSet := os.LookupEnv("FLAKE")
		if wasSet && rand.Float64() > 0.5 {
			Fail("bad luck, try again next time")
		}
	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
	})

	It("requires a YAML file argument", func() {
		command := exec.Command(cliPath)
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(Exit(1))
		Ω(session.Err).Should(Say(usage))
	})

	It("requires a the YAML file to exist", func() {
		command := exec.Command(cliPath, "no/such/file.yml", "echo foo")
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(Exit(1))
		Ω(session.Err).Should(Say("no/such/file.yml does not exist"))
		Ω(session.Err).ShouldNot(Say("foo"))
	})

	It("requires a command to invoke", func() {
		command := exec.Command(cliPath, getFixturePath("vars.yml"))
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(Exit(1))
		Ω(session.Err).Should(Say(usage))
	})

	It("invokes the given command passing env vars from the YAML file", func() {
		command := exec.Command(cliPath, getFixturePath("vars.yml"), getFixturePath("script.sh"))
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(Exit(0))
		Ω(session).Should(Say("value from yaml"))
	})

	It("invokes the given command passing boolean env vars from the YAML file", func() {
		command := exec.Command(cliPath, getFixturePath("boolean.yml"), getFixturePath("script.sh"))
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(Exit(0))
		Ω(session).Should(Say("true"))
	})

	It("invokes the given command passing integer env vars from the YAML file", func() {
		command := exec.Command(cliPath, getFixturePath("integer.yml"), getFixturePath("script.sh"))
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(Exit(0))
		Ω(session).Should(Say("42"))
	})

})

func getFixturePath(relativePath string) string {
	return filepath.Join(*fixturesPath, relativePath)
}
