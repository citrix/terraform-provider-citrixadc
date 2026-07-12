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

const testAccVpnvserver_vpnurlpolicy_binding_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_example_vserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
		name             = "tf_vpnurlaction"
		linkname         = "new_link"
		actualurl        = "http://www.citrix.com"
		applicationtype  = "CVPN"
		clientlessaccess = "OFF"
		comment          = "Testing"
		ssotype          = "unifiedgateway"
		vservername      = "vserver1"
	}
	resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
		name   = "new_policy"
		rule   = "true"
		action = citrixadc_vpnurlaction.tf_vpnurlaction.name
	}
	resource "citrixadc_vpnvserver_vpnurlpolicy_binding" "tf_bind" {
		name                   = citrixadc_vpnvserver.tf_vpnvserver.name
		policy                 = citrixadc_vpnurlpolicy.tf_vpnurlpolicy.name
		priority               = 20
		bindpoint              = "REQUEST"
		gotopriorityexpression = "next"
	}
`

const testAccVpnvserver_vpnurlpolicy_bindingDataSource_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_example_vserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
		name             = "tf_vpnurlaction"
		linkname         = "new_link"
		actualurl        = "http://www.citrix.com"
		applicationtype  = "CVPN"
		clientlessaccess = "OFF"
		comment          = "Testing"
		ssotype          = "unifiedgateway"
		vservername      = "vserver1"
	}
	resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
		name   = "new_policy"
		rule   = "true"
		action = citrixadc_vpnurlaction.tf_vpnurlaction.name
	}
	resource "citrixadc_vpnvserver_vpnurlpolicy_binding" "tf_bind" {
		name                   = citrixadc_vpnvserver.tf_vpnvserver.name
		policy                 = citrixadc_vpnurlpolicy.tf_vpnurlpolicy.name
		priority               = 20
		bindpoint              = "REQUEST"
		gotopriorityexpression = "next"
	}

	data "citrixadc_vpnvserver_vpnurlpolicy_binding" "tf_bind" {
		name   = citrixadc_vpnvserver_vpnurlpolicy_binding.tf_bind.name
		policy = citrixadc_vpnvserver_vpnurlpolicy_binding.tf_bind.policy
	}
`

const testAccVpnvserver_vpnurlpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_example_vserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
		name             = "tf_vpnurlaction"
		linkname         = "new_link"
		actualurl        = "http://www.citrix.com"
		applicationtype  = "CVPN"
		clientlessaccess = "OFF"
		comment          = "Testing"
		ssotype          = "unifiedgateway"
		vservername      = "vserver1"
	}
	resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
		name   = "new_policy"
		rule   = "true"
		action = citrixadc_vpnurlaction.tf_vpnurlaction.name
	}
`

const testAccVpnvserver_vpnurlpolicy_binding_upgrade_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_example_vserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
		name             = "tf_vpnurlaction"
		linkname         = "new_link"
		actualurl        = "http://www.citrix.com"
		applicationtype  = "CVPN"
		clientlessaccess = "OFF"
		comment          = "Testing"
		ssotype          = "unifiedgateway"
		vservername      = "vserver1"
	}
	resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
		name   = "new_policy"
		rule   = "true"
		action = citrixadc_vpnurlaction.tf_vpnurlaction.name
	}
	resource "citrixadc_vpnvserver_vpnurlpolicy_binding" "tf_bind" {
		name                   = citrixadc_vpnvserver.tf_vpnvserver.name
		policy                 = citrixadc_vpnurlpolicy.tf_vpnurlpolicy.name
		priority               = 20
		bindpoint              = "REQUEST"
		gotopriorityexpression = "next"
	}
`

// TestAccVpnvserver_vpnurlpolicy_binding_sdkv2StateUpgrade verifies that a
// binding created with the last SDK v2 provider release (2.2.0, legacy comma-joined ID)
// is refreshed and upgraded cleanly by the current framework provider, which recomputes
// the ID into the new key:value format on Read.
func TestAccVpnvserver_vpnurlpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_vpnvserver_vpnurlpolicy_binding.tf_bind"
	legacyId := "tf_example_vserver,new_policy"
	newId := "name:tf_example_vserver,policy:new_policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnvserver_vpnurlpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release, writing legacy-id state.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccVpnvserver_vpnurlpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpnurlpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", legacyId),
				),
			},
			// Step 2: refresh/plan/apply the legacy-id state through the current framework
			// provider. Read exercises ParseIdString on the legacy id and recomputes the
			// canonical new-format id.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVpnvserver_vpnurlpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpnurlpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", newId),
				),
			},
		},
	})
}

func TestAccVpnvserver_vpnurlpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_vpnurlpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_vpnurlpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpnurlpolicy_bindingExist("citrixadc_vpnvserver_vpnurlpolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccVpnvserver_vpnurlpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpnurlpolicy_bindingNotExist("citrixadc_vpnvserver_vpnurlpolicy_binding.tf_bind", "tf_example_vserver,new_policy"),
				),
			},
		},
	})
}

func TestAccVpnvserver_vpnurlpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver_vpnurlpolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_vpnurlpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnvserver_vpnurlpolicy_binding_basic},
			{Config: testAccVpnvserver_vpnurlpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"bindpoint"}},
		},
	})
}

func testAccCheckVpnvserver_vpnurlpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_vpnurlpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "policy"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %q: %v", bindingId, err)
		}
		name := idMap["name"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_vpnurlpolicy_binding",
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
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnvserver_vpnurlpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_vpnurlpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"name", "policy"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %q: %v", id, err)
		}
		name := idMap["name"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_vpnurlpolicy_binding",
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
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnvserver_vpnurlpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_vpnurlpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_vpnurlpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnvserver_vpnurlpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver_vpnurlpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnvserver_vpnurlpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_vpnurlpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_vpnurlpolicy_binding.tf_bind", "name", "tf_example_vserver"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_vpnurlpolicy_binding.tf_bind", "policy", "new_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_vpnurlpolicy_binding.tf_bind", "priority", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_vpnurlpolicy_binding.tf_bind", "gotopriorityexpression", "next"),
				),
			},
		},
	})
}
