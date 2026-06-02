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

// Participating entities reused from existing acceptance tests:
//   - metricsprofile : citrixadc_framework/acctests/metricsprofile_test.go (_basic_step1 block)
//   - servicegroup   : citrixadc_framework/acctests/servicegroup_test.go (servicegroup "foo" block)
//
// The binding has no NITRO update endpoint and all attributes are RequiresReplace,
// so the basic test creates the binding in step1 and drops it in step2 to verify deletion.

const testAccMetricsprofile_servicegroup_binding_basic_step1 = `
resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
}

resource "citrixadc_metricsprofile_servicegroup_binding" "tf_metricsprofile_servicegroup_binding" {
  name       = citrixadc_metricsprofile.tf_metricsprofile.name
  entityname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
  entitytype = "servicegroup"
  depends_on = [
    citrixadc_metricsprofile.tf_metricsprofile,
    citrixadc_servicegroup.tf_servicegroup,
  ]
}

`

const testAccMetricsprofile_servicegroup_binding_basic_step2 = `
# Keep the participating entities without the actual binding to verify proper deletion
resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
}

`

func TestAccMetricsprofile_servicegroup_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofile_servicegroup_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_servicegroup_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_servicegroup_bindingExist("citrixadc_metricsprofile_servicegroup_binding.tf_metricsprofile_servicegroup_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_servicegroup_binding.tf_metricsprofile_servicegroup_binding", "name", "tf_metricsprofile"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_servicegroup_binding.tf_metricsprofile_servicegroup_binding", "entityname", "tf_servicegroup"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_servicegroup_binding.tf_metricsprofile_servicegroup_binding", "entitytype", "servicegroup"),
				),
			},
			{
				Config: testAccMetricsprofile_servicegroup_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_servicegroup_bindingNotExist("citrixadc_metricsprofile_servicegroup_binding.tf_metricsprofile_servicegroup_binding", "entityname:tf_servicegroup,entitytype:servicegroup,name:tf_metricsprofile"),
				),
			},
		},
	})
}

func testAccCheckMetricsprofile_servicegroup_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No metricsprofile_servicegroup_binding id is set")
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

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		entityname := idMap["entityname"]
		entitytype := idMap["entitytype"]

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_servicegroup_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching entityname/entitytype
		found := false
		for _, v := range dataArr {
			if en, ok := v["entityname"].(string); ok && en == entityname {
				if et, ok := v["entitytype"].(string); ok && et == entitytype {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("metricsprofile_servicegroup_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_servicegroup_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}

		name := idMap["name"]
		entityname := idMap["entityname"]
		entitytype := idMap["entitytype"]

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_servicegroup_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the matching entityname/entitytype
		found := false
		for _, v := range dataArr {
			if en, ok := v["entityname"].(string); ok && en == entityname {
				if et, ok := v["entitytype"].(string); ok && et == entitytype {
					found = true
					break
				}
			}
		}

		if found {
			return fmt.Errorf("metricsprofile_servicegroup_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_servicegroup_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_metricsprofile_servicegroup_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		entityname := idMap["entityname"]
		entitytype := idMap["entitytype"]

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_servicegroup_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent metricsprofile or binding no longer exists - destroyed as expected
			continue
		}

		found := false
		for _, v := range dataArr {
			if en, ok := v["entityname"].(string); ok && en == entityname {
				if et, ok := v["entitytype"].(string); ok && et == entitytype {
					found = true
					break
				}
			}
		}

		if found {
			return fmt.Errorf("metricsprofile_servicegroup_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccMetricsprofile_servicegroup_binding_DataSource_basic = `
resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
}

resource "citrixadc_metricsprofile_servicegroup_binding" "tf_metricsprofile_servicegroup_binding" {
  name       = citrixadc_metricsprofile.tf_metricsprofile.name
  entityname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
  entitytype = "servicegroup"
  depends_on = [
    citrixadc_metricsprofile.tf_metricsprofile,
    citrixadc_servicegroup.tf_servicegroup,
  ]
}

data "citrixadc_metricsprofile_servicegroup_binding" "tf_metricsprofile_servicegroup_binding" {
  name       = citrixadc_metricsprofile.tf_metricsprofile.name
  entityname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
  entitytype = "servicegroup"
  depends_on = [citrixadc_metricsprofile_servicegroup_binding.tf_metricsprofile_servicegroup_binding]
}
`

func TestAccMetricsprofile_servicegroup_binding_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofile_servicegroup_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_servicegroup_binding_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_servicegroup_binding.tf_metricsprofile_servicegroup_binding", "name", "tf_metricsprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_servicegroup_binding.tf_metricsprofile_servicegroup_binding", "entityname", "tf_servicegroup"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_servicegroup_binding.tf_metricsprofile_servicegroup_binding", "entitytype", "servicegroup"),
				),
			},
		},
	})
}
