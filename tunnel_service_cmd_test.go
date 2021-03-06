package main_test

import (
	"os/exec"

	"errors"

	"github.com/cloudfoundry/cli/plugin/models"
	"github.com/cloudfoundry/cli/testhelpers/rpc_server"
	fake_rpc_handlers "github.com/cloudfoundry/cli/testhelpers/rpc_server/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

const validPluginPath = "./tunnel_service_cmd.exe"

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
			})

			It("raises an error when not enough arguments are supplied", func() {
				args := []string{ts.Port(), "tunnel-service", "my-data-service"}
				session, err := gexec.Start(exec.Command(validPluginPath, args...), GinkgoWriter, GinkgoWriter)
				session.Wait()

				Expect(err).NotTo(HaveOccurred())
				Expect(session).To(gbytes.Say("cf tunnel-service service-instance-name remote-port"))
			})
		})

		Context("Finding a service instance", func() {
			It("raises an error when a service with the provided service name doesn't exist", func() {
				rpcHandlers.GetServiceStub = func(_ string, retVal *plugin_models.GetService_Model) error {
					//*retVal = plugin_models.GetService_Model{}
					return errors.New("Service instance not found")
				}

				args := []string{ts.Port(), "tunnel-service", "my-data-service", "8080"}
				session, err := gexec.Start(exec.Command(validPluginPath, args...), GinkgoWriter, GinkgoWriter)
				session.Wait()

				Expect(err).NotTo(HaveOccurred())
				Expect(session).To(gbytes.Say("Service instance my-data-service not found"))
			})

			It("returns the service and plan name if the service is successfully found", func() {
				rpcHandlers.GetServiceStub = func(_ string, retVal *plugin_models.GetService_Model) error {
					*retVal = plugin_models.GetService_Model{
						Guid:           "abcde-12345",
						Name:           "my-data-service",
						DashboardUrl:   "https://my-dashboard.com",
						IsUserProvided: false,
						ServiceOffering: plugin_models.GetService_ServiceFields{
							Name:             "my-service",
							DocumentationUrl: "https://docs.my-service.com",
						},
						ServicePlan: plugin_models.GetService_ServicePlan{
							Name: "my-plan",
							Guid: "7890-defg",
						},
						LastOperation: plugin_models.GetService_LastOperation{
							Type:        "type",
							State:       "state",
							Description: "description",
							CreatedAt:   "created at",
							UpdatedAt:   "updated at",
						},
					}
					return nil
				}

				args := []string{ts.Port(), "tunnel-service", "my-data-service", "8080"}
				session, err := gexec.Start(exec.Command(validPluginPath, args...), GinkgoWriter, GinkgoWriter)
				session.Wait()

				Expect(err).NotTo(HaveOccurred())
				Expect(session).To(gbytes.Say("Found service instance my-data-service"))
				Expect(session).To(gbytes.Say("Service Name: my-service"))
				Expect(session).To(gbytes.Say("Service Plan: my-plan"))
			})
		})
	})
})
