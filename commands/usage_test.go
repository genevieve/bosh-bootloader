package commands_test

import (
	"bytes"
	"strings"

	"github.com/pivotal-cf-experimental/bosh-bootloader/commands"
	"github.com/pivotal-cf-experimental/bosh-bootloader/storage"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Usage", func() {
	var (
		usage  commands.Usage
		stdout *bytes.Buffer
	)

	BeforeEach(func() {
		stdout = bytes.NewBuffer([]byte{})
		usage = commands.NewUsage(stdout)
	})

	Describe("Execute", func() {
		It("prints out the usage information", func() {
			_, err := usage.Execute(commands.GlobalFlags{}, []string{}, storage.State{})
			Expect(err).NotTo(HaveOccurred())
			Expect(stdout.String()).To(Equal(strings.TrimSpace(`
Usage:
  bbl [GLOBAL OPTIONS] COMMAND [OPTIONS]

Global Options:
  --help    [-h] "print usage"
  --version [-v] "print version"

  --aws-access-key-id     "AWS AccessKeyID value"
  --aws-secret-access-key "AWS SecretAccessKey value"
  --aws-region            "AWS Region"
  --state-dir             "Directory that stores the state.json"

Commands:
  destroy [--no-confirm]                                                        "tears down a BOSH Director environment on AWS"
  director-address                                                              "print the BOSH director address"
  director-password                                                             "print the BOSH director password"
  director-username                                                             "print the BOSH director username"
  help                                                                          "print usage"
  ssh-key                                                                       "print the ssh private key"
  unsupported-create-lbs [--type=<concourse,cf>] [--cert=<path>] [--key=<path>] "attaches a load balancer with the supplied certificate and key"
  unsupported-deploy-bosh-on-aws-for-concourse [--lb-type=concourse,cf,none]    "deploys a BOSH Director on AWS"
  version                                                                       "print version"
`)))
		})

		It("returns the given state unmodified", func() {
			state, err := usage.Execute(commands.GlobalFlags{}, []string{}, storage.State{
				Version: 2,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(state).To(Equal(storage.State{
				Version: 2,
			}))
		})
	})
})
