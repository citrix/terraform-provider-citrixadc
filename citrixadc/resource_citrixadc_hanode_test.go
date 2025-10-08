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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccHanodeLocal_basic = `
  
resource "citrixadc_hanode" "local_node" {
	hanode_id     = 0
	hellointerval = 200
	deadinterval = 5
	}
   
`
const testAccHanodeLocal_update = `
	resource "citrixadc_hanode" "local_node" {
		hanode_id     = 0
		hellointerval = 400
		deadinterval = 30
	}
	
`

func TestAccHanodeLocal_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testAccHanodeLocal_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.local_node", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "hanode_id", "0"),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "hellointerval", "200"),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "deadinterval", "5"),
				),
			},
			{
				Config: testAccHanodeLocal_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.local_node", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "hanode_id", "0"),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "hellointerval", "400"),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "deadinterval", "30"),
				),
			},
		},
	})
}

const testAccHanodeRemote_basic = `
  
resource "citrixadc_hanode" "remote_node" {
	hanode_id = 2
	ipaddress = "10.222.74.145"
	}
  
   
`
const testAccHanodeRemote_update = `
	resource "citrixadc_hanode" "remote_node" {
		hanode_id = 3
		ipaddress = "10.222.74.145"
	}
	
`

func TestAccHanodeRemote_basic(t *testing.T) {
	if adcTestbed != "HA" {
		t.Skipf("ADC testbed is %s. Expected HA.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckHanodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHanodeRemote_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.remote_node", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.remote_node", "hanode_id", "2"),
					resource.TestCheckResourceAttr("citrixadc_hanode.remote_node", "ipaddress", "10.222.74.145"),
				),
			},
			{
				Config: testAccHanodeRemote_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.remote_node", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.remote_node", "hanode_id", "3"),
					resource.TestCheckResourceAttr("citrixadc_hanode.remote_node", "ipaddress", "10.222.74.145"),
				),
			},
		},
	})
}

func testAccCheckHanodeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No hanode name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Hanode.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("hanode %s not found", n)
		}

		return nil
	}
}

func testAccCheckHanodeDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_hanode" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Hanode.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("hanode %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
