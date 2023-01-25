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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccNsrpcnode_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	if isCpxRun {
		t.Skip("Operation not permitted under CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsrpcnode_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
				),
			},
			resource.TestStep{
				Config: testAccNsrpcnode_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
				),
			},
		},
	})
}

func testAccCheckNsrpcnodeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Nsrpcnode.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("RPC node %s not found", n)
		}

		return nil
	}
}

const testAccNsrpcnode_basic_step1 = `

resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "10.222.74.146"
    password = "notnsroot"
    secure = "ON"
    srcip = "10.222.74.146"
}
`

const testAccNsrpcnode_basic_step2 = `

resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "10.222.74.146"
    password = "notnsroot"
    secure = "OFF"
    srcip = "10.222.74.146"
}
`
