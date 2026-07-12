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

const testAccLbgroup_lbvserver_binding_basic = `
	resource "citrixadc_lbgroup_lbvserver_binding" "tf_lbvserverbinding" {
		name = citrixadc_lbgroup.tf_lbgroup.name
		vservername = citrixadc_lbvserver.tf_lbvserver.name
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name 		= "tf_lbvserver"
		ipv46       = "1.1.1.8"
		port        = "80"
		servicetype = "HTTP"
	}
	
	resource "citrixadc_lbgroup" "tf_lbgroup" {
		name = "tf_lbgroup"
	}
`

const testAccLbgroup_lbvserver_binding_basic_step2 = `
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name = "tf_lbvserver"
		ipv46       = "1.1.1.8"
		port        = "80"
		servicetype = "HTTP"
	}
	
	resource "citrixadc_lbgroup" "tf_lbgroup" {
		name = "tf_lbgroup"
	}
`

func TestAccLbgroup_lbvserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbgroup_lbvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbgroup_lbvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroup_lbvserver_bindingExist("citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding", nil),
				),
			},
			{
				Config: testAccLbgroup_lbvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroup_lbvserver_bindingNotExist("citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding", "tf_lbgroup,tf_lbvserver"),
				),
			},
		},
	})
}

func TestAccLbgroup_lbvserver_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbgroup_lbvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbgroup_lbvserver_binding_basic},
			{Config: testAccLbgroup_lbvserver_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckLbgroup_lbvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbgroup_lbvserver_binding name is set")
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
		// ParseIdString handles both the new key:value ID format and the legacy
		// comma-separated SDK v2 format.
		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "vservername"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		lbgroupName := idMap["name"]
		lbvserverName := idMap["vservername"]

		findParams := service.FindParams{
			ResourceType:             service.Lbgroup_lbvserver_binding.Type(),
			ResourceName:             lbgroupName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if v["vservername"].(string) == lbvserverName {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find lbgroup_lbvserver_binding ID %v", bindingId)
		}

		return nil
	}
}

func testAccCheckLbgroup_lbvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		lbgroupName := idSlice[0]
		lbvserverName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lbgroup_lbvserver_binding",
			ResourceName:             lbgroupName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		found := false
		for _, v := range dataArr {
			if v["vservername"].(string) == lbvserverName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("Lbgroup_lbvserver_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbgroup_lbvserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbgroup_lbvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbgroup_lbvserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbgroup_lbvserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbgroup_lbvserver_bindingDataSource_basic = `
	resource "citrixadc_lbgroup_lbvserver_binding" "tf_lbvserverbinding" {
		name = citrixadc_lbgroup.tf_lbgroup.name
		vservername = citrixadc_lbvserver.tf_lbvserver.name
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name 		= "tf_lbvserver"
		ipv46       = "1.1.1.8"
		port        = "80"
		servicetype = "HTTP"
	}
	
	resource "citrixadc_lbgroup" "tf_lbgroup" {
		name = "tf_lbgroup"
	}

	data "citrixadc_lbgroup_lbvserver_binding" "tf_lbvserverbinding" {
		name = citrixadc_lbgroup.tf_lbgroup.name
		vservername = citrixadc_lbvserver.tf_lbvserver.name
		depends_on = [citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding]
	}
`

func TestAccLbgroup_lbvserver_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbgroup_lbvserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding", "name", "tf_lbgroup"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding", "vservername", "tf_lbvserver"),
				),
			},
		},
	})
}

const testAcclbgroup_lbvserver_binding_upgrade_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver"
	ipv46       = "1.1.1.8"
	port        = "80"
	servicetype = "HTTP"
}

resource "citrixadc_lbgroup" "tf_lbgroup" {
	name = "tf_lbgroup"
}

resource "citrixadc_lbgroup_lbvserver_binding" "tf_lbvserverbinding" {
	name        = citrixadc_lbgroup.tf_lbgroup.name
	vservername = citrixadc_lbvserver.tf_lbvserver.name
}
`

// TestAccLbgroup_lbvserver_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release (legacy comma-separated ID) is correctly upgraded when
// the same config is subsequently managed by the current Framework provider. Step 1
// creates the binding with citrix/citrixadc 2.2.0 (writes the legacy id
// "tf_lbgroup,tf_lbvserver"). Step 2 refreshes/plans/applies the same config through
// the Framework provider, exercising ParseIdString on the legacy id; because the
// Framework recomputes the id on Read (SetAttrFromGet), the id upgrades to the new
// "key:value" form "name:tf_lbgroup,vservername:tf_lbvserver".
func TestAccLbgroup_lbvserver_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbgroup_lbvserver_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAcclbgroup_lbvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroup_lbvserver_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_lbgroup,tf_lbvserver"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAcclbgroup_lbvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroup_lbvserver_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "name:tf_lbgroup,vservername:tf_lbvserver"),
				),
			},
		},
	})
}
