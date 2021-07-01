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
package citrixadc

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCluster_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("clustering not supported in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClusterDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCluster_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExist("citrixadc_cluster.tf_cluster", nil),
				),
			},
			resource.TestStep{
				Config: testAccCluster_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExist("citrixadc_cluster.tf_cluster", nil),
				),
			},
			resource.TestStep{
				Config: testAccCluster_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExist("citrixadc_cluster.tf_cluster", nil),
				),
			},
			resource.TestStep{
				Config: testAccCluster_step4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExist("citrixadc_cluster.tf_cluster", nil),
				),
			},
		},
	})
}

func testAccCheckClusterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Clusterinstance.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusterDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusterinstance" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Clusterinstance.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCluster_step1 = `

resource "citrixadc_cluster" "tf_cluster" {
    clid = 1
    clip = "10.78.60.15"
	hellointerval = 200

    clusternode { 
        nodeid = 0
        delay = 0
        priority = 30
        endpoint = "http://10.78.60.10"
        backplane = "0/1/1"
        ipaddress = "10.78.60.10"
        tunnelmode = "NONE"
        nodegroup = "DEFAULT_NG"

        state = "ACTIVE"
    }

    clusternode { 
        nodeid = 1
        delay = 0
        priority = 31
        endpoint = "http://10.78.60.11"
        ipaddress = "10.78.60.11"
        backplane = "1/1/1"
        tunnelmode = "NONE"
        nodegroup = "DEFAULT_NG"

        state = "ACTIVE"
    }
}
`

const testAccCluster_step2 = `

resource "citrixadc_cluster" "tf_cluster" {
    clid = 1
    clip = "10.78.60.15"
	hellointerval = 400

    clusternode { 
        nodeid = 0
        delay = 0
        priority = 30
        endpoint = "http://10.78.60.10"
        backplane = "0/1/1"
        ipaddress = "10.78.60.10"
        tunnelmode = "NONE"
        nodegroup = "DEFAULT_NG"

        state = "ACTIVE"
    }

    clusternode { 
        nodeid = 1
        delay = 0
        priority = 20
        endpoint = "http://10.78.60.11"
        ipaddress = "10.78.60.11"
        backplane = "1/1/1"
        tunnelmode = "NONE"
        nodegroup = "DEFAULT_NG"

        state = "ACTIVE"
    }
}
`

const testAccCluster_step3 = `

resource "citrixadc_cluster" "tf_cluster" {
    clid = 1
    clip = "10.78.60.15"
	hellointerval = 400

    clusternode { 
        nodeid = 1
        delay = 0
        priority = 20
        endpoint = "http://10.78.60.11"
        ipaddress = "10.78.60.11"
        backplane = "1/1/1"
        tunnelmode = "NONE"
        nodegroup = "DEFAULT_NG"

        state = "ACTIVE"
    }
}
`

const testAccCluster_step4 = `

resource "citrixadc_cluster" "tf_cluster" {
    clid = 1
    clip = "10.78.60.15"
	hellointerval = 400

    clusternode { 
        nodeid = 0
        delay = 0
        priority = 30
        endpoint = "http://10.78.60.10"
        backplane = "0/1/1"
        ipaddress = "10.78.60.10"
        tunnelmode = "NONE"
        nodegroup = "DEFAULT_NG"

        state = "ACTIVE"
    }

    clusternode { 
        nodeid = 1
        delay = 0
        priority = 31
        endpoint = "http://10.78.60.11"
        ipaddress = "10.78.60.11"
        backplane = "1/1/1"
        tunnelmode = "NONE"
        nodegroup = "DEFAULT_NG"

        state = "ACTIVE"
    }
}
`
