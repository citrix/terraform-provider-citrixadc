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

const testAccLbvserver_servicegroup_binding_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_lbvserver2" {
  ipv46       = "10.10.10.34"
  name        = "tf_lbvserver2"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
    servicegroupname = "tf_servicegroup"
    servicetype  = "HTTP"
}

resource "citrixadc_servicegroup" "tf_servicegroup2" {
    servicegroupname = "tf_servicegroup2"
    servicetype  = "HTTP"
}

resource "citrixadc_lbvserver_servicegroup_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}

resource "citrixadc_lbvserver_servicegroup_binding" "tf_binding2" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicegroupname = citrixadc_servicegroup.tf_servicegroup2.servicegroupname
  order = 4
}

resource "citrixadc_lbvserver_servicegroup_binding" "tf_binding3" {
  name = citrixadc_lbvserver.tf_lbvserver2.name
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}

resource "citrixadc_lbvserver_servicegroup_binding" "tf_binding4" {
  name = citrixadc_lbvserver.tf_lbvserver2.name
  servicegroupname = citrixadc_servicegroup.tf_servicegroup2.servicegroupname
  order = 4
}

`

func TestAccLbvserver_servicegroup_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_servicegroup_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_servicegroup_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_servicegroup_bindingExist("citrixadc_lbvserver_servicegroup_binding.tf_binding", nil, map[string]interface{}{}),
					testAccCheckLbvserver_servicegroup_bindingExist("citrixadc_lbvserver_servicegroup_binding.tf_binding2", nil, map[string]interface{}{"order": 4}),
					testAccCheckLbvserver_servicegroup_bindingExist("citrixadc_lbvserver_servicegroup_binding.tf_binding3", nil, map[string]interface{}{}),
					testAccCheckLbvserver_servicegroup_bindingExist("citrixadc_lbvserver_servicegroup_binding.tf_binding4", nil, map[string]interface{}{"order": 4}),
				),
			},
		},
	})
}

func TestAccLbvserver_servicegroup_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_servicegroup_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_servicegroup_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbvserver_servicegroup_binding_basic},
			{Config: testAccLbvserver_servicegroup_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckLbvserver_servicegroup_bindingExist(n string, id *string, expectedValues map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_servicegroup_binding name is set")
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

		// ID-parse helper line updated for the new key:value ID format (migration adopts
		// the Framework ID format; ParseIdString also decodes the legacy "name,servicegroupname" form).
		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "servicegroupname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %q: %v", bindingId, err)
		}
		name := idMap["name"]
		servicegroupname := idMap["servicegroupname"]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_servicegroup_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		foundIndex := -1
		data := map[string]interface{}{}
		for i, v := range dataArr {
			if v["servicegroupname"].(string) == servicegroupname {
				foundIndex = i
				data = v
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("lbvserver_servicegroup_binding %s not found", n)
		}

		// Iterate through all expected values and validate them
		for key, expectedValue := range expectedValues {
			if actualValue, exists := data[key]; !exists {
				return fmt.Errorf("Expected key %q not found in retrieved data", key)
			} else if !compareValues(expectedValue, actualValue) {
				return fmt.Errorf("Expected value for %q differs. Expected: %v, Retrieved: %v",
					key, expectedValue, actualValue)
			}
		}

		return nil
	}
}

func testAccCheckLbvserver_servicegroup_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_servicegroup_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		// ID-parse helper line updated for the new key:value ID format (ParseIdString also
		// decodes the legacy "name,servicegroupname" form). FindResource call shape unchanged.
		idMap, _, perr := utils.ParseIdString(rs.Primary.ID, []string{"name", "servicegroupname"}, nil)
		if perr != nil {
			return fmt.Errorf("Error parsing ID %q: %v", rs.Primary.ID, perr)
		}
		_, err := client.FindResource(service.Lbvserver_servicegroup_binding.Type(), idMap["name"])
		if err == nil {
			return fmt.Errorf("lbvserver_servicegroup_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbvserver_servicegroup_bindingDataSource_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
    servicegroupname = "tf_servicegroup"
    servicetype  = "HTTP"
}

resource "citrixadc_lbvserver_servicegroup_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}

data "citrixadc_lbvserver_servicegroup_binding" "tf_binding" {
    name = citrixadc_lbvserver_servicegroup_binding.tf_binding.name
    servicegroupname = citrixadc_lbvserver_servicegroup_binding.tf_binding.servicegroupname
    depends_on = [citrixadc_lbvserver_servicegroup_binding.tf_binding]
}
`

func TestAccLbvserver_servicegroup_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_servicegroup_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_servicegroup_binding.tf_binding", "name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_servicegroup_binding.tf_binding", "servicegroupname", "tf_servicegroup"),
				),
			},
		},
	})
}

const testAccLbvserver_servicegroup_binding_upgrade_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
    servicegroupname = "tf_servicegroup"
    servicetype  = "HTTP"
}

resource "citrixadc_lbvserver_servicegroup_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}
`

// TestAccLbvserver_servicegroup_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release (legacy comma-separated ID) is correctly
// upgraded when the same config is subsequently managed by the current Framework
// provider. Step 1 creates the binding with citrix/citrixadc 2.2.0 (writes the
// legacy id "tf_lbvserver,tf_servicegroup"). Step 2 refreshes/plans/applies the
// same config through the Framework provider, exercising ParseIdString on the
// legacy id; because the Framework recomputes the id on Read (SetAttrFromGet),
// the id upgrades to the new "key:value" form.
func TestAccLbvserver_servicegroup_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_lbvserver_servicegroup_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbvserver_servicegroup_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLbvserver_servicegroup_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_servicegroup_bindingExist(resourceAddr, nil, map[string]interface{}{}),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_lbvserver,tf_servicegroup"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbvserver_servicegroup_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_servicegroup_bindingExist(resourceAddr, nil, map[string]interface{}{}),
					resource.TestCheckResourceAttr(resourceAddr, "id", "name:tf_lbvserver,servicegroupname:tf_servicegroup"),
				),
			},
		},
	})
}
