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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAuthenticationepaaction_add = `

	resource "citrixadc_authenticationepaaction" "tf_epaaction" {
		name            = "tf_epaaction"
		csecexpr        = "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"
		defaultepagroup = "new_group"
		deletefiles     = "old_files"
		killprocess     = "old_process"
	}

	resource "citrixadc_authenticationepaaction" "tf_epaaction2" {
		name            = "tf_epaaction2"
		deviceposture	 = "DISABLED"
		defaultepagroup = "new_group"
	}
`
const testAccAuthenticationepaaction_update = `

	resource "citrixadc_authenticationepaaction" "tf_epaaction" {
		name            = "tf_epaaction"
		csecexpr        = "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"
		defaultepagroup = "new_group"
		deletefiles     = "new_files"
		killprocess     = "new_process"
	}

	resource "citrixadc_authenticationepaaction" "tf_epaaction2" {
		name            = "tf_epaaction2"
		deviceposture	 = "ENABLED"
		defaultepagroup = "new_group"
	}
`

func TestAccAuthenticationepaaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationepaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationepaaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_epaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "name", "tf_epaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "deletefiles", "old_files"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "killprocess", "old_process"),
					testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_epaaction2", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction2", "name", "tf_epaaction2"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction2", "deviceposture", "DISABLED"),
				),
			},
			{
				Config: testAccAuthenticationepaaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_epaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "name", "tf_epaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "deletefiles", "new_files"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "killprocess", "new_process"),
					testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_epaaction2", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction2", "name", "tf_epaaction2"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction2", "deviceposture", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationepaactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationepaaction name is set")
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
		data, err := client.FindResource("authenticationepaaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationepaaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationepaactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationepaaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("aAuthenticationepaaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationepaaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationepaactionDataSource_basic = `

	resource "citrixadc_authenticationepaaction" "tf_epaaction_ds" {
		name            = "tf_epaaction_ds"
		csecexpr        = "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"
		defaultepagroup = "new_group"
		deletefiles     = "old_files"
		killprocess     = "old_process"
	}

	data "citrixadc_authenticationepaaction" "tf_epaaction_ds" {
		name = citrixadc_authenticationepaaction.tf_epaaction_ds.name
	}
`

func TestAccAuthenticationepaactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationepaactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationepaaction.tf_epaaction_ds", "name", "tf_epaaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationepaaction.tf_epaaction_ds", "csecexpr", "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationepaaction.tf_epaaction_ds", "defaultepagroup", "new_group"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationepaaction.tf_epaaction_ds", "deletefiles", "old_files"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationepaaction.tf_epaaction_ds", "killprocess", "old_process"),
				),
			},
		},
	})
}
