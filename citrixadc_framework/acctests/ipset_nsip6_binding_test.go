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

const testAccIpset_nsip6_binding_basic = `

	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_ipset"
	}
	resource "citrixadc_nsip6" "tf_nsip6" {
		ipv6address = "2003:db8:100::fb/64"
		type 		= "VIP"
	}

	resource "citrixadc_ipset_nsip6_binding" "tf_ipset_nsip6_binding" {
		name      = citrixadc_ipset.tf_ipset.name
		ipaddress = citrixadc_nsip6.tf_nsip6.ipv6address
	}
    
`

const testAccIpset_nsip6_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_ipset"
	}
	resource "citrixadc_nsip6" "tf_nsip6" {
		ipv6address = "2003:db8:100::fb/64"
		type 		= "VIP"
	}

`

const testAccIpset_nsip6_bindingDataSource_basic = `

	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_ipset"
	}
	resource "citrixadc_nsip6" "tf_nsip6" {
		ipv6address = "2003:db8:100::fb/64"
		type 		= "VIP"
	}

	resource "citrixadc_ipset_nsip6_binding" "tf_ipset_nsip6_binding" {
		name      = citrixadc_ipset.tf_ipset.name
		ipaddress = citrixadc_nsip6.tf_nsip6.ipv6address
	}

	data "citrixadc_ipset_nsip6_binding" "tf_ipset_nsip6_binding" {
		name       = citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding.name
		ipaddress  = citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding.ipaddress
		depends_on = [citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding]
	}
    
`

func TestAccIpset_nsip6_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpset_nsip6_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpset_nsip6_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpset_nsip6_bindingExist("citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding", nil),
				),
			},
			{
				Config: testAccIpset_nsip6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpset_nsip6_bindingNotExist("citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding", "tf_ipset,2003:db8:100::fb/64"),
				),
			},
		},
	})
}

func testAccCheckIpset_nsip6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ipset_nsip6_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "ipaddress"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %v: %v", bindingId, err)
		}
		name := idMap["name"]
		ipaddress := idMap["ipaddress"]

		findParams := service.FindParams{
			ResourceType:             "ipset_nsip6_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ipaddress
		found := false
		for _, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("ipset_nsip6_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckIpset_nsip6_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		name := idSlice[0]
		ipaddress := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "ipset_nsip6_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ipaddress
		found := false
		for _, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("ipset_nsip6_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckIpset_nsip6_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ipset_nsip6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Ipset_nsip6_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("ipset_nsip6_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccIpset_nsip6_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIpset_nsip6_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding", "name", "tf_ipset"),
					resource.TestCheckResourceAttr("data.citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding", "ipaddress", "2003:db8:100::fb/64"),
				),
			},
		},
	})
}

const testAccIpset_nsip6_binding_upgrade_basic = `

	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_ipset"
	}
	resource "citrixadc_nsip6" "tf_nsip6" {
		ipv6address = "2003:db8:100::fb/64"
		type 		= "VIP"
	}

	resource "citrixadc_ipset_nsip6_binding" "tf_ipset_nsip6_binding" {
		name      = citrixadc_ipset.tf_ipset.name
		ipaddress = citrixadc_nsip6.tf_nsip6.ipv6address
	}

`

// TestAccIpset_nsip6_binding_sdkv2StateUpgrade verifies that state written by the
// last SDK v2 release (legacy comma-separated ID) is correctly upgraded when the
// same config is subsequently managed by the current Framework provider. Step 1
// creates the binding with citrix/citrixadc 2.2.0 (writes the legacy id
// "tf_ipset,2003:db8:100::fb/64"). Step 2 refreshes/plans/applies the same config
// through the Framework provider, exercising ParseIdString on the legacy id; because
// the Framework recomputes the id on Read (SetAttrFromGet re-derives data.Id), the
// id upgrades to the new "key:value" form.
func TestAccIpset_nsip6_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckIpset_nsip6_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccIpset_nsip6_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpset_nsip6_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_ipset,2003:db8:100::fb/64"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccIpset_nsip6_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpset_nsip6_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "ipaddress:2003%3Adb8%3A100%3A%3Afb%2F64,name:tf_ipset"),
				),
			},
		},
	})
}

func TestAccIpset_nsip6_binding_import(t *testing.T) {
	const resAddr = "citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpset_nsip6_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccIpset_nsip6_binding_basic},
			{Config: testAccIpset_nsip6_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
