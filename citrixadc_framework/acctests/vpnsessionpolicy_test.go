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

const testAccVpnsessionpolicy_add = `

	resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction" {
		name                       = "newsession"
		sesstimeout                = "10"
		defaultauthorizationaction = "ALLOW"
	}
	resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction_update" {
		name                       = "newsession_update"
		sesstimeout                = "10"
		defaultauthorizationaction = "ALLOW"
	}
	
	resource "citrixadc_vpnsessionpolicy" "foo" {
		
		name   = "tf_vpnsessionpolicy"
  		rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
  		action = citrixadc_vpnsessionaction.tf_vpnsessionaction.name	
	}
`
const testAccVpnsessionpolicy_update = `

	resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction" {
		name                       = "newsession"
		sesstimeout                = "10"
		defaultauthorizationaction = "ALLOW"
	}
	resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction_update" {
		name                       = "newsession_update"
		sesstimeout                = "10"
		defaultauthorizationaction = "ALLOW"
	}

	resource "citrixadc_vpnsessionpolicy" "foo" {
		
		name   = "tf_vpnsessionpolicy"
  		rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\")"
  		action = citrixadc_vpnsessionaction.tf_vpnsessionaction_update.name
	}
`

const testAccVpnsessionpolicyDataSource_basic = `

	resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction" {
		name                       = "newsession"
		sesstimeout                = "10"
		defaultauthorizationaction = "ALLOW"
	}
	resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction_update" {
		name                       = "newsession_update"
		sesstimeout                = "10"
		defaultauthorizationaction = "ALLOW"
	}
	
	resource "citrixadc_vpnsessionpolicy" "foo" {
		
		name   = "tf_vpnsessionpolicy"
  		rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
  		action = citrixadc_vpnsessionaction.tf_vpnsessionaction.name	
	}

	data "citrixadc_vpnsessionpolicy" "foo" {
		name = citrixadc_vpnsessionpolicy.foo.name
	}
`

func TestAccVpnsessionpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnsessionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnsessionpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsessionpolicyExist("citrixadc_vpnsessionpolicy.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionpolicy.foo", "name", "tf_vpnsessionpolicy"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionpolicy.foo", "rule", "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionpolicy.foo", "action", "newsession"),
				),
			},
			{
				Config: testAccVpnsessionpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsessionpolicyExist("citrixadc_vpnsessionpolicy.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionpolicy.foo", "name", "tf_vpnsessionpolicy"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionpolicy.foo", "rule", "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\")"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionpolicy.foo", "action", "newsession_update"),
				),
			},
		},
	})
}

func testAccCheckVpnsessionpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnsessionpolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnsessionpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnsessionpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnsessionpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnsessionpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnsessionpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnsessionpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnsessionpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnsessionpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnsessionpolicy.foo", "name", "tf_vpnsessionpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnsessionpolicy.foo", "rule", "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnsessionpolicy.foo", "action", "newsession"),
				),
			},
		},
	})
}
