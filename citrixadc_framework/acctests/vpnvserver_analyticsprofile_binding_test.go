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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccVpnvserver_analyticsprofile_binding_basic = `

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name = "new_profile"
		type = "tcpinsight"
	}
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name           = "tf_vserver"
		servicetype    = "SSL"
		ipv46          = "3.3.3.3"
		port           = 443
	}
	resource "citrixadc_vpnvserver_analyticsprofile_binding" "tf_bind" {
		name 			 = citrixadc_vpnvserver.tf_vpnvserver.name
		analyticsprofile = citrixadc_analyticsprofile.tf_analyticsprofile.name
	}
`

const testAccVpnvserver_analyticsprofile_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	
	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name = "new_profile"
		type = "tcpinsight"
	}
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name           = "tf_vserver"
		servicetype    = "SSL"
		ipv46          = "3.3.3.3"
		port           = 443
	}
`

func TestAccVpnvserver_analyticsprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_analyticsprofile_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_analyticsprofile_bindingExist("citrixadc_vpnvserver_analyticsprofile_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccVpnvserver_analyticsprofile_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_analyticsprofile_bindingNotExist("citrixadc_vpnvserver_analyticsprofile_binding.tf_bind", "tf_vserver,new_profile"),
				),
			},
		},
	})
}

func testAccCheckVpnvserver_analyticsprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_analyticsprofile_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "analyticsprofile"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		analyticsprofile := idMap["analyticsprofile"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_analyticsprofile_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["analyticsprofile"].(string) == analyticsprofile {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnvserver_analyticsprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_analyticsprofile_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"name", "analyticsprofile"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		name := idMap["name"]
		analyticsprofile := idMap["analyticsprofile"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_analyticsprofile_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["analyticsprofile"].(string) == analyticsprofile {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnvserver_analyticsprofile_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_analyticsprofile_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_analyticsprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnvserver_analyticsprofile_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver_analyticsprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnvserver_analyticsprofile_bindingDataSource_basic = `

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name = "new_profile"
		type = "tcpinsight"
	}
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name           = "tf_vserver"
		servicetype    = "SSL"
		ipv46          = "3.3.3.3"
		port           = 443
	}
	resource "citrixadc_vpnvserver_analyticsprofile_binding" "tf_bind" {
		name 			 = citrixadc_vpnvserver.tf_vpnvserver.name
		analyticsprofile = citrixadc_analyticsprofile.tf_analyticsprofile.name
	}

	data "citrixadc_vpnvserver_analyticsprofile_binding" "tf_bind" {
		name             = citrixadc_vpnvserver_analyticsprofile_binding.tf_bind.name
		analyticsprofile = citrixadc_vpnvserver_analyticsprofile_binding.tf_bind.analyticsprofile
		depends_on       = [citrixadc_vpnvserver_analyticsprofile_binding.tf_bind]
	}
`

// Config for the SDK v2 -> Framework state-upgrade test. Reuses the _basic
// resource block and prerequisites; valid under both the SDK v2 2.2.0 schema
// and the current Framework schema (same attribute names).
const testAccVpnvserver_analyticsprofile_binding_upgrade_basic = `

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name = "new_profile"
		type = "tcpinsight"
	}
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name           = "tf_vserver"
		servicetype    = "SSL"
		ipv46          = "3.3.3.3"
		port           = 443
	}
	resource "citrixadc_vpnvserver_analyticsprofile_binding" "tf_bind" {
		name 			 = citrixadc_vpnvserver.tf_vpnvserver.name
		analyticsprofile = citrixadc_analyticsprofile.tf_analyticsprofile.name
	}
`

// TestAccVpnvserver_analyticsprofile_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release (legacy comma-joined id "name,analyticsprofile")
// is transparently upgraded by the current Framework provider. Step 1 creates the
// binding with citrix/citrixadc 2.2.0; step 2 refreshes/plans the same config through
// the current Framework provider, whose Read parses the legacy id and recomputes it
// to the new "analyticsprofile:<v>,name:<v>" canonical format (SetAttrFromGet).
func TestAccVpnvserver_analyticsprofile_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnvserver_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release, writing the legacy id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccVpnvserver_analyticsprofile_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_analyticsprofile_bindingExist("citrixadc_vpnvserver_analyticsprofile_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_analyticsprofile_binding.tf_bind", "id", "tf_vserver,new_profile"),
				),
			},
			{
				// Step 2: refresh/apply the same config through the current Framework
				// provider. Read exercises ParseIdString on the legacy id, then
				// recomputes the id to the new key:value canonical format.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVpnvserver_analyticsprofile_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_analyticsprofile_bindingExist("citrixadc_vpnvserver_analyticsprofile_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_analyticsprofile_binding.tf_bind", "id", "analyticsprofile:new_profile,name:tf_vserver"),
				),
			},
		},
	})
}

func TestAccVpnvserver_analyticsprofile_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_analyticsprofile_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_analyticsprofile_binding.tf_bind", "name", "tf_vserver"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_analyticsprofile_binding.tf_bind", "analyticsprofile", "new_profile"),
				),
			},
		},
	})
}

func TestAccVpnvserver_analyticsprofile_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver_analyticsprofile_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnvserver_analyticsprofile_binding_basic},
			{Config: testAccVpnvserver_analyticsprofile_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
