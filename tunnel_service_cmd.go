package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type TunnelService struct {
	ServiceInstanceName string
	ServiceInstancePort string
	ServiceName         string
	ServicePlan         string
}

func (t *TunnelService) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Tunnel Service",
		Commands: []plugin.Command{
			{
				Name:     "tunnel-service",
				HelpText: "Forwards a remote port from a bound service instance to your local machine",
				UsageDetails: plugin.Usage{
					Usage: "cf tunnel-service service-instance-name remote-port",
					Options: map[string]string{
						"service-instance-name": "Name of the already created service instance that you want to forward the ports for",
						"remote-port":           "Port number (integer) of the remote port to be forwarded",
					},
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(TunnelService))
}

func (t *TunnelService) Run(cliConnection plugin.CliConnection, args []string) {
	if len(args) < 3 {
		fmt.Println("Incorrect usage")
		fmt.Println(t.GetMetadata().Commands[0].UsageDetails.Usage)
	} else {

		t.SetProperties(args)
		t.FetchServiceDetails(cliConnection)

	}
}

func (t *TunnelService) SetProperties(args []string) {
	t.ServiceInstanceName = args[1]
	t.ServiceInstancePort = args[2]
}

func (t *TunnelService) FetchServiceDetails(cliConnection plugin.CliConnection) {

	returnedService, err := cliConnection.GetService(t.ServiceInstanceName)

	if err != nil {
		fmt.Printf("Service instance %v not found", t.ServiceInstanceName)
		fmt.Printf("error %v", err)
	} else {
		t.ServiceName = returnedService.ServiceOffering.Name
		t.ServicePlan = returnedService.ServicePlan.Name

		fmt.Printf("Found service instance %v", t.ServiceInstanceName)
		fmt.Printf("Service Name: %v", t.ServiceName)
		fmt.Printf("Service Plan: %v", t.ServicePlan)

	}

}
