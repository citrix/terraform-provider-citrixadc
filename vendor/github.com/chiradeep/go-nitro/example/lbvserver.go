/*
Copyright 2016 Citrix Systems, Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"log"

	"github.com/chiradeep/go-nitro/config/basic"
	"github.com/chiradeep/go-nitro/config/lb"
	"github.com/chiradeep/go-nitro/netscaler"
)

func main() {
	client, err := netscaler.NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}
	log.Printf("Client is %+v\n", *client)
	lb1 := lb.Lbvserver{
		Name:        "sample_lb",
		Ipv46:       "10.71.136.50",
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	client.AddResource(netscaler.Lbvserver.Type(), "sample_lb", &lb1)

	lb1 = lb.Lbvserver{
		Name:     "sample_lb",
		Lbmethod: "LEASTCONNECTION",
	}
	client.UpdateResource(netscaler.Lbvserver.Type(), "sample_lb", &lb1)

	service1 := basic.Service{
		Name:        "sample_svc_1",
		Ip:          "172.22.33.4",
		Port:        80,
		Servicetype: "HTTP",
	}

	client.AddResource(netscaler.Service.Type(), "sample_svc_1", &service1)

	binding := lb.Lbvserverservicebinding{
		Name:        "sample_lb",
		Servicename: "sample_svc_1",
	}

	client.BindResource(netscaler.Lbvserver.Type(), "sample_lb", netscaler.Service.Type(), "sample_svc_1", &binding)
	client.SaveConfig()

	client.UnbindResource(netscaler.Lbvserver.Type(), "sample_lb", netscaler.Service.Type(), "sample_svc_1", "servicename")

	client.DeleteResource(netscaler.Lbvserver.Type(), "sample_lb")
	client.DeleteResource(netscaler.Service.Type(), "sample_svc_1")

	client.EnableFeatures([]string{"CS"})
	client.SaveConfig()

}
