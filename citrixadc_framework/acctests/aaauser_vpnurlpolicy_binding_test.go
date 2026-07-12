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

const testAccAaauser_vpnurlpolicy_binding_basic = `

resource "citrixadc_aaauser_vpnurlpolicy_binding" "tf_aaauser_vpnurlpolicy_binding" {
	username = citrixadc_aaauser.tf_aaauser.username
	policy    = citrixadc_vpnurlpolicy.tf_vpnurlpolicy.name
	type     = "REQUEST"
	priority  = 100
	}
  
  

  resource "citrixadc_aaauser" "tf_aaauser" {
	username = "user1"
	password = "my_pass"
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

const testAccAaauser_vpnurlpolicy_binding_basic_step2 = `

resource "citrixadc_aaauser" "tf_aaauser" {
  username = "user1"
  password = "my_pass"
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

func TestAccAaauser_vpnurlpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaauser_vpnurlpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaauser_vpnurlpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_vpnurlpolicy_bindingExist("citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding", nil),
				),
			},
			{
				Config: testAccAaauser_vpnurlpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_vpnurlpolicy_bindingNotExist("citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding", "user1,new_policy"),
				),
			},
		},
	})
}

func testAccCheckAaauser_vpnurlpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaauser_vpnurlpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"username", "policy"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		username := idMap["username"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "aaauser_vpnurlpolicy_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("aaauser_vpnurlpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaauser_vpnurlpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		username := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaauser_vpnurlpolicy_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("aaauser_vpnurlpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaauser_vpnurlpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaauser_vpnurlpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("aaauser_vpnurlpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaauser_vpnurlpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAaauser_vpnurlpolicy_bindingDataSource_basic = `

resource "citrixadc_aaauser_vpnurlpolicy_binding" "tf_aaauser_vpnurlpolicy_binding" {
	username = citrixadc_aaauser.tf_aaauser.username
	policy    = citrixadc_vpnurlpolicy.tf_vpnurlpolicy.name
	type     = "REQUEST"
	priority  = 100
}

resource "citrixadc_aaauser" "tf_aaauser" {
	username = "user1"
	password = "my_pass"
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

data "citrixadc_aaauser_vpnurlpolicy_binding" "tf_aaauser_vpnurlpolicy_binding" {
	username = citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding.username
	policy   = citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding.policy
}
`

func TestAccAaauser_vpnurlpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAaauser_vpnurlpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding", "username", "user1"),
					resource.TestCheckResourceAttr("data.citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding", "policy", "new_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding", "priority", "100"),
				),
			},
		},
	})
}

// testAccAaauser_vpnurlpolicy_binding_upgrade_basic reuses the same resource block +
// prerequisites as testAccAaauser_vpnurlpolicy_binding_basic. It is valid under BOTH the
// SDK v2 2.2.0 schema and the current Framework schema (the migration restored the SDK v2
// attribute names), so the same config can create the resource with the last SDK v2 release
// and then be refreshed/applied through the current Framework provider.
const testAccAaauser_vpnurlpolicy_binding_upgrade_basic = `

resource "citrixadc_aaauser_vpnurlpolicy_binding" "tf_aaauser_vpnurlpolicy_binding" {
	username = citrixadc_aaauser.tf_aaauser.username
	policy    = citrixadc_vpnurlpolicy.tf_vpnurlpolicy.name
	type     = "REQUEST"
	priority  = 100
	}



  resource "citrixadc_aaauser" "tf_aaauser" {
	username = "user1"
	password = "my_pass"
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

// TestAccAaauser_vpnurlpolicy_binding_sdkv2StateUpgrade verifies that state written by the
// last SDK v2 release (legacy comma-separated id) is correctly upgraded when refreshed and
// applied through the current Framework provider (which re-derives the canonical key:value id
// on Read via SetAttrFromGet).
func TestAccAaauser_vpnurlpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAaauser_vpnurlpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release (2.2.0), writing state with the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAaauser_vpnurlpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_vpnurlpolicy_bindingExist("citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding", "id", "user1,new_policy"),
				),
			},
			// Step 2: refresh the legacy-id state through the current Framework provider.
			// Read exercises ParseIdString on the legacy id and re-derives the new-format id.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAaauser_vpnurlpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_vpnurlpolicy_bindingExist("citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding", "id", "policy:new_policy,username:user1"),
				),
			},
		},
	})
}

func TestAccAaauser_vpnurlpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaauser_vpnurlpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAaauser_vpnurlpolicy_binding_basic},
			{Config: testAccAaauser_vpnurlpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"type"}},
		},
	})
}
