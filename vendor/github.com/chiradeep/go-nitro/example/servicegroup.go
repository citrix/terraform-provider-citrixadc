/*
Copyright 2018 Citrix Systems, Inc

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
	"github.com/chiradeep/go-nitro/netscaler"
)

func example2() {
	client, err := netscaler.NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}

	createSvcGrp := basic.Servicegroup{
		Servicegroupname: "test-svcgroup",
		Servicetype:      "TCP",
	}

	client.AddResource(netscaler.Servicegroup.Type(), "test-svcgroup", &createSvcGrp)

	createServer := basic.Server{
		Ipaddress: "192.168.1.101",
		Name:      "test-srvr",
	}

	client.AddResource(netscaler.Server.Type(), "test-server", &createServer)

	bindSvcGrpToServer := basic.Servicegroupservicegroupmemberbinding{
		Servicegroupname: "test-svcgroup",
		Servername:       "test-srvr",
		Port:             22,
	}

	client.AddResource(netscaler.Servicegroup_servicegroupmember_binding.Type(), "test-svcgroup", &bindSvcGrpToServer)

	bindSvcGrpToServer2 := basic.Servicegroupservicegroupmemberbinding{
		Servicegroupname: "test-svcgroup",
		Ip:               "192.168.1.102",
		Port:             22,
	}
	client.AddResource(netscaler.Servicegroup_servicegroupmember_binding.Type(), "test-svcgroup", &bindSvcGrpToServer2)
}
