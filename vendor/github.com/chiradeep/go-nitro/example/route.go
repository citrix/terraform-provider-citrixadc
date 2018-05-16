package main

import (
	"fmt"

	"github.com/chiradeep/go-nitro/config/network"
	"github.com/chiradeep/go-nitro/netscaler"
)

func main() {
	client, _ := netscaler.NewNitroClientFromEnv()
	route := network.Route{
		Network: "192.168.15.0",
		Netmask: "255.255.255.0",
		Gateway: "172.17.0.2",
	}
	// adding route
	_, err := client.AddResource(netscaler.Route.Type(), "anyRoute", &route)
	if err == nil {
		client.SaveConfig()
	}
	// deleting route
	//var argsBundle = []string{"network:192.168.15.0", "netmask:255.255.255.0", "gateway:172.17.0.2"}
	var argsBundle = map[string]string{"network": "192.168.15.0", "netmask": "255.255.255.0", "gateway": "172.17.0.2"}
	//err2 := client.DeleteResourceWithArgs(netscaler.Route.Type(), "", argsBundle)
	err2 := client.DeleteResourceWithArgsMap(netscaler.Route.Type(), "", argsBundle)
	if err2 != nil {
		fmt.Println(err2)
	}
}
