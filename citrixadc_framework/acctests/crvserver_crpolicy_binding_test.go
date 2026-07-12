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

const testAccCrvserver_crpolicy_binding_basic = `

resource "citrixadc_crpolicy" "crpolicy" {
    policyname = "crpolicy1"
    rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
    action = "ORIGIN"
}
resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_crvserver_crpolicy_binding" "crvserver_crpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_crpolicy.crpolicy.policyname
    priority = 10 
}
`

const testAccCrvserver_crpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_crpolicy" "crpolicy" {
		policyname = "crpolicy1"
		rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
		action = "ORIGIN"
	}
	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver"
		servicetype = "HTTP"
		arp = "OFF"
	}
`

func TestAccCrvserver_crpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserver_crpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_crpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_crpolicy_bindingExist("citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding", nil),
				),
			},
			{
				Config: testAccCrvserver_crpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_crpolicy_bindingNotExist("citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding", "my_vserver,crpolicy1"),
				),
			},
		},
	})
}

// TestAccCrvserver_crpolicy_binding_import verifies `terraform import` for the
// binding. The binding implements ImportStatePassthroughID(id), so Read must
// reconstruct the full state from the id alone. Two import shapes are exercised:
//   1. the canonical new-format id that the resource stores (default: prior state's id)
//   2. the legacy SDK v2 positional id "name,policyname" (via ImportStateIdFunc)
// ImportStateVerify compares the imported instance attribute-by-attribute against
// the applied state; attributes the ADC does not echo back on the binding GET are
// listed in ImportStateVerifyIgnore.
func TestAccCrvserver_crpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserver_crpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_crpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_crpolicy_bindingExist(resAddr, nil),
				),
			},
			// (1) import using the resource's stored (new-format) id
			{
				Config:                  testAccCrvserver_crpolicy_binding_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
			// (2) import using the legacy SDK v2 positional id "name,policyname"
			{
				Config:                  testAccCrvserver_crpolicy_binding_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           "my_vserver,crpolicy1",
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckCrvserver_crpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No crvserver_crpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "crvserver_crpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("crvserver_crpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_crpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "crvserver_crpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("crvserver_crpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_crpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_crvserver_crpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Crvserver_crpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("crvserver_crpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCrvserver_crpolicy_bindingDataSource_basic = `

	resource "citrixadc_crpolicy" "tf_crpolicy" {
		policyname = "my_crpolicy_ds"
		rule       = "true"
		action     = "ORIGIN"
	}
	resource "citrixadc_crvserver" "crvserver" {
		name        = "my_vserver_ds"
		servicetype = "HTTP"
		arp         = "OFF"
	}
	resource "citrixadc_crvserver_crpolicy_binding" "crvserver_crpolicy_binding" {
		name       = citrixadc_crvserver.crvserver.name
		policyname = citrixadc_crpolicy.tf_crpolicy.policyname
		priority   = 10
	}

	data "citrixadc_crvserver_crpolicy_binding" "crvserver_crpolicy_binding" {
		name       = citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding.name
		policyname = citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding.policyname
		depends_on = [citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding]
	}
`

func TestAcccrvserver_crpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_crpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding", "name", "my_vserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding", "policyname", "my_crpolicy_ds"),
				),
			},
		},
	})
}

// testAcccrvserver_crpolicy_binding_upgrade_basic is the config used by the
// sdkv2 -> framework state-upgrade test. It reuses the same values and resource
// labels as testAccCrvserver_crpolicy_binding_basic so it is valid under BOTH
// the SDK v2 2.2.0 schema and the current framework schema.
const testAcccrvserver_crpolicy_binding_upgrade_basic = `

resource "citrixadc_crpolicy" "crpolicy" {
    policyname = "crpolicy1"
    rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
    action = "ORIGIN"
}
resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_crvserver_crpolicy_binding" "crvserver_crpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_crpolicy.crpolicy.policyname
    priority = 10
}
`

// TestAccCrvserver_crpolicy_binding_sdkv2StateUpgrade verifies that a binding
// created with the last SDK v2 release (2.2.0, legacy comma-separated ID) is
// correctly refreshed/planned/applied by the current framework provider.
func TestAccCrvserver_crpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCrvserver_crpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create the binding with the last SDK v2 release.
			// State is written with the LEGACY comma-separated id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAcccrvserver_crpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_crpolicy_bindingExist("citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding", "id", "my_vserver,crpolicy1"),
				),
			},
			// Step 2: same config, current (framework) provider. Terraform
			// refreshes the legacy-id state through the framework Read
			// (exercising ParseIdString on the legacy id) then plans/applies.
			// The framework recomputes the id on read to the new key:value form.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAcccrvserver_crpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_crpolicy_bindingExist("citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding", "id", "name:my_vserver,policyname:crpolicy1"),
				),
			},
		},
	})
}
