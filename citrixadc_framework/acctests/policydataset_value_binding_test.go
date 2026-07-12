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

func TestAccPolicydataset_value_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicydataset_value_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicydataset_value_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value1"),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value2"),
				),
			},
			{
				Config: testAccPolicydataset_value_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value1"),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value3"),
				),
			},
			{
				Config: testAccPolicydataset_value_binding_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value3"),
				),
			},
		},
	})
}

func testAccCheckPolicydatasetValue(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No binding id")
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"name", "value"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		value := idMap["value"]

		findParams := service.FindParams{
			ResourceType:             "policydataset_value_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 2823,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return fmt.Errorf("Error during FindResourceArrayWithParams %s", err.Error())
		}

		// Resource is missing
		if len(dataArr) == 0 {
			return fmt.Errorf("FindResourceArrayWithParams returned empty array")
		}

		// Iterate through results to find the one with the right monitor name
		foundIndex := -1
		for i, v := range dataArr {
			if v["value"].(string) == value {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("FindResourceArrayWithParams monitor name not found in array")
		}

		return nil
	}
}

func testAccCheckPolicydataset_value_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policydataset_value_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Policydataset_value_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccPolicydataset_value_binding_basic_step1 = `
resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value1" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 100
  index    = 111
  endrange = 150
}

resource "citrixadc_policydataset_value_binding" "tf_value2" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 200
  endrange = 250
}
`

const testAccPolicydataset_value_binding_basic_step2 = `
resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value1" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 100
  index    = 111
  endrange = 160
}

resource "citrixadc_policydataset_value_binding" "tf_value3" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 300
}
`

const testAccPolicydataset_value_binding_basic_step3 = `
resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value3" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 300
  index    = 333
  endrange = 360
}
`

const testAccPolicydataset_value_bindingDataSource_basic = `
resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value1" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 100
  index    = 111
  endrange = 150
}

data "citrixadc_policydataset_value_binding" "tf_value1" {
  name     = citrixadc_policydataset_value_binding.tf_value1.name
  value    = citrixadc_policydataset_value_binding.tf_value1.value
  endrange = citrixadc_policydataset_value_binding.tf_value1.endrange
  depends_on = [citrixadc_policydataset_value_binding.tf_value1]
}
`

func TestAccPolicydataset_value_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicydataset_value_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_policydataset_value_binding.tf_value1", "name", "tf_dataset"),
					resource.TestCheckResourceAttr("data.citrixadc_policydataset_value_binding.tf_value1", "value", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_policydataset_value_binding.tf_value1", "index", "111"),
					resource.TestCheckResourceAttr("data.citrixadc_policydataset_value_binding.tf_value1", "endrange", "150"),
				),
			},
		},
	})
}

// testAccPolicydataset_value_binding_upgrade_basic mirrors the tf_value1 entry from
// _basic_step1 (a policydataset parent plus a single value binding). It is valid
// under BOTH the SDK v2 2.2.0 schema and the current framework schema, so it can be
// created with the old provider in step 1 and re-planned with the new provider in
// step 2 of the state-upgrade test below.
const testAccPolicydataset_value_binding_upgrade_basic = `
resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value1" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 100
  index    = 111
  endrange = 150
}
`

func TestAccPolicydataset_value_binding_import(t *testing.T) {
	const resAddr = "citrixadc_policydataset_value_binding.tf_value1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicydataset_value_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccPolicydataset_value_binding_basic_step1},
			{Config: testAccPolicydataset_value_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

// TestAccPolicydataset_value_binding_sdkv2StateUpgrade verifies that a binding
// created by the LAST SDK v2 release (2.2.0) — which writes the legacy comma-joined
// id "name,value" — is refreshed and re-applied correctly by the CURRENT framework
// provider. Step 2 exercises ParseIdString on the legacy id during the framework
// Read; the framework SetAttrFromGet then recomputes data.Id into the canonical new
// "key:value" format, so the id upgrades in place.
func TestAccPolicydataset_value_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckPolicydataset_value_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release from the registry. This
			// writes state carrying the LEGACY comma-joined id "name,value".
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccPolicydataset_value_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value1"),
					resource.TestCheckResourceAttr("citrixadc_policydataset_value_binding.tf_value1", "id", "tf_dataset,100"),
				),
			},
			// Step 2: same config through the CURRENT framework provider. Terraform
			// refreshes the legacy-id state through the framework Read (exercising
			// ParseIdString on the legacy id), then plans/applies. SetAttrFromGet
			// recomputes the id to the new "key:value" format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccPolicydataset_value_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value1"),
					resource.TestCheckResourceAttr("citrixadc_policydataset_value_binding.tf_value1", "id", "endrange:150,name:tf_dataset,value:100"),
				),
			},
		},
	})
}
