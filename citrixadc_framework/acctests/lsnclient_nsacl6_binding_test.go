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

const testAccLsnclient_nsacl6_binding_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsn_client"
}

resource "citrixadc_nsacl6" "tf_nsacl6" {
	acl6name   = "my_acl6"
	acl6action = "ALLOW"
}

resource "citrixadc_lsnclient_nsacl6_binding" "tf_lsnclient_nsacl6_binding" {
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
	acl6name   = citrixadc_nsacl6.tf_nsacl6.acl6name
}
`

const testAccLsnclient_nsacl6_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccLsnclient_nsacl6_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnclient_nsacl6_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnclient_nsacl6_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_nsacl6_bindingExist("citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding", nil),
				),
			},
			{
				Config: testAccLsnclient_nsacl6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_nsacl6_bindingNotExist("citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding", "my_lsn_client,my_acl6"),
				),
			},
		},
	})
}

func testAccCheckLsnclient_nsacl6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnclient_nsacl6_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"clientname", "acl6name"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		clientname := idMap["clientname"]
		acl6name := idMap["acl6name"]

		findParams := service.FindParams{
			ResourceType:             "lsnclient_nsacl6_binding",
			ResourceName:             clientname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching acl6name
		found := false
		for _, v := range dataArr {
			if v["acl6name"].(string) == acl6name {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsnclient_nsacl6_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnclient_nsacl6_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"clientname", "acl6name"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		clientname := idMap["clientname"]
		acl6name := idMap["acl6name"]

		findParams := service.FindParams{
			ResourceType:             "lsnclient_nsacl6_binding",
			ResourceName:             clientname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching acl6name
		found := false
		for _, v := range dataArr {
			if v["acl6name"].(string) == acl6name {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsnclient_nsacl6_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsnclient_nsacl6_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnclient_nsacl6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnclient_nsacl6_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnclient_nsacl6_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsnclient_nsacl6_bindingDataSource_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsn_client"
}

resource "citrixadc_nsacl6" "tf_nsacl6" {
	acl6name   = "my_acl6"
	acl6action = "ALLOW"
}

resource "citrixadc_lsnclient_nsacl6_binding" "tf_lsnclient_nsacl6_binding" {
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
	acl6name   = citrixadc_nsacl6.tf_nsacl6.acl6name
}

data "citrixadc_lsnclient_nsacl6_binding" "tf_lsnclient_nsacl6_binding" {
	clientname = citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding.clientname
	acl6name   = citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding.acl6name
}
`

func TestAccLsnclient_nsacl6_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnclient_nsacl6_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding", "clientname", "my_lsn_client"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding", "acl6name", "my_acl6"),
				),
			},
		},
	})
}

const testAccLsnclient_nsacl6_binding_upgrade_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsn_client"
}

resource "citrixadc_nsacl6" "tf_nsacl6" {
	acl6name   = "my_acl6"
	acl6action = "ALLOW"
}

resource "citrixadc_lsnclient_nsacl6_binding" "tf_lsnclient_nsacl6_binding" {
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
	acl6name   = citrixadc_nsacl6.tf_nsacl6.acl6name
}
`

func TestAccLsnclient_nsacl6_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsnclient_nsacl6_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create the resource with the last SDK v2 release (2.2.0),
			// which writes state using the legacy comma-separated ID format.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLsnclient_nsacl6_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_nsacl6_bindingExist("citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding", "id", "my_lsn_client,my_acl6"),
				),
			},
			// Step 2: refresh the legacy-ID state through the current (Framework)
			// provider. Read parses the legacy ID via ParseIdString and recomputes
			// the ID into the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLsnclient_nsacl6_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_nsacl6_bindingExist("citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding", "id", "acl6name:my_acl6,clientname:my_lsn_client"),
				),
			},
		},
	})
}

func TestAccLsnclient_nsacl6_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnclient_nsacl6_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsnclient_nsacl6_binding_basic},
			{Config: testAccLsnclient_nsacl6_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
