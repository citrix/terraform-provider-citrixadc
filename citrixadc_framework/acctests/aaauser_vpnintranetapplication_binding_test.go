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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAaauser_vpnintranetapplication_binding_basic = `

	resource "citrixadc_aaauser_vpnintranetapplication_binding" "tf_aaauser_vpnintranetapplication_binding" {
		username            = citrixadc_aaauser.tf_aaauser.username
		intranetapplication = citrixadc_vpnintranetapplication.tf_vpnintranetapplication.intranetapplication
	}
	
	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
		intranetapplication = "tf_vpnintranetapplication"
		protocol            = "UDP"
		destip              = "2.3.6.5"
		interception        = "TRANSPARENT"
	}
  
`

const testAccAaauser_vpnintranetapplication_bindingDataSource_basic = `

	resource "citrixadc_aaauser_vpnintranetapplication_binding" "tf_aaauser_vpnintranetapplication_binding" {
		username            = citrixadc_aaauser.tf_aaauser.username
		intranetapplication = citrixadc_vpnintranetapplication.tf_vpnintranetapplication.intranetapplication
	}
	
	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
		intranetapplication = "tf_vpnintranetapplication"
		protocol            = "UDP"
		destip              = "2.3.6.5"
		interception        = "TRANSPARENT"
	}

	data "citrixadc_aaauser_vpnintranetapplication_binding" "tf_aaauser_vpnintranetapplication_binding" {
		username            = citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding.username
		intranetapplication = citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding.intranetapplication
		depends_on          = [citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding]
	}
  
`

const testAccAaauser_vpnintranetapplication_binding_basic_step2 = `
	 
	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
		intranetapplication = "tf_vpnintranetapplication"
		protocol            = "UDP"
		destip              = "2.3.6.5"
		interception        = "TRANSPARENT"
	}
`

func TestAccAaauser_vpnintranetapplication_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaauser_vpnintranetapplication_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaauser_vpnintranetapplication_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_vpnintranetapplication_bindingExist("citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding", nil),
				),
			},
			{
				Config: testAccAaauser_vpnintranetapplication_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_vpnintranetapplication_bindingNotExist("citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding", "user1,tf_vpnintranetapplication"),
				),
			},
		},
	})
}

func testAccCheckAaauser_vpnintranetapplication_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaauser_vpnintranetapplication_binding id is set")
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

		bindingId := rs.Primary.ID

		idMap, _, err := utils.ParseIdString(bindingId, []string{"username", "intranetapplication"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		username := idMap["username"]
		intranetapplication := idMap["intranetapplication"]

		findParams := service.FindParams{
			ResourceType:             "aaauser_vpnintranetapplication_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching intranetapplication
		found := false
		for _, v := range dataArr {
			if v["intranetapplication"].(string) == intranetapplication {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("aaauser_vpnintranetapplication_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaauser_vpnintranetapplication_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"username", "intranetapplication"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		username := idMap["username"]
		intranetapplication := idMap["intranetapplication"]

		findParams := service.FindParams{
			ResourceType:             "aaauser_vpnintranetapplication_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching intranetapplication
		found := false
		for _, v := range dataArr {
			if v["intranetapplication"].(string) == intranetapplication {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("aaauser_vpnintranetapplication_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaauser_vpnintranetapplication_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaauser_vpnintranetapplication_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Aaauser_vpnintranetapplication_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaauser_vpnintranetapplication_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAaauser_vpnintranetapplication_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAaauser_vpnintranetapplication_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding", "username", "user1"),
					resource.TestCheckResourceAttr("data.citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding", "intranetapplication", "tf_vpnintranetapplication"),
				),
			},
		},
	})
}

const testAccAaauser_vpnintranetapplication_binding_upgrade_basic = `

	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
		intranetapplication = "tf_vpnintranetapplication"
		protocol            = "UDP"
		destip              = "2.3.6.5"
		interception        = "TRANSPARENT"
	}

	resource "citrixadc_aaauser_vpnintranetapplication_binding" "tf_aaauser_vpnintranetapplication_binding" {
		username            = citrixadc_aaauser.tf_aaauser.username
		intranetapplication = citrixadc_vpnintranetapplication.tf_vpnintranetapplication.intranetapplication
	}
`

func TestAccAaauser_vpnintranetapplication_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAaauser_vpnintranetapplication_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy comma-joined id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAaauser_vpnintranetapplication_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_vpnintranetapplication_bindingExist("citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding", "id", "user1,tf_vpnintranetapplication"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAaauser_vpnintranetapplication_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_vpnintranetapplication_bindingExist("citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding", "id", "intranetapplication:tf_vpnintranetapplication,username:user1"),
				),
			},
		},
	})
}

func TestAccAaauser_vpnintranetapplication_binding_import(t *testing.T) {
	const resAddr = "citrixadc_aaauser_vpnintranetapplication_binding.tf_aaauser_vpnintranetapplication_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaauser_vpnintranetapplication_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAaauser_vpnintranetapplication_binding_basic},
			{Config: testAccAaauser_vpnintranetapplication_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
