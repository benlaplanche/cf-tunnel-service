package main_test

import (
	"os/exec"

	"github.com/cloudfoundry/cli/testhelpers/rpc_server"
	fake_rpc_handlers "github.com/cloudfoundry/cli/testhelpers/rpc_server/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

const validPluginPath = "./service-tunnel"

var _ = Describe("TunnelServiceCmd", func() {
	var (
		rpcHandlers *fake_rpc_handlers.FakeHandlers
		ts          *test_rpc_server.TestServer
		err         error
	)

	BeforeEach(func() {
		rpcHandlers = &fake_rpc_handlers.FakeHandlers{}
		ts, err = test_rpc_server.NewTestRpcServer(rpcHandlers)
		Expect(err).NotTo(HaveOccurred())

		err = ts.Start()
		Expect(err).NotTo(HaveOccurred())

		rpcHandlers.CallCoreCommandStub = func(_ []string, retVal *bool) error {
			*retVal = true
			return nil
		}

		rpcHandlers.GetOutputAndResetStub = func(_ bool, retVal *[]string) error {
			*retVal = []string{"{}"}
			return nil
		}

	})

	AfterEach(func() {
		ts.Stop()
	})

	Describe("tunnel-service", func() {
		Context("Option flags", func() {
			It("accepts service-instance-name and remote-port as valid mandatory flags", func() {
				args := []string{ts.Port(), "tunnel-service", "my-data-service", "8080"}
				session, err := gexec.Start(exec.Command(validPluginPath, args...), GinkgoWriter, GinkgoWriter)
				session.Wait()

				Expect(err).NotTo(HaveOccurred())
				Expect(session).To(gbytes.Say("hello from tunnel-service command"))
			})
		})
	})
})
