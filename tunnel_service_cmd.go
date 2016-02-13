package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type TunnelService struct{}

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

		switch args[0] {
		case "tunnel-service":
			fmt.Printf("hello from tunnel-service command")
		}

	}
}
