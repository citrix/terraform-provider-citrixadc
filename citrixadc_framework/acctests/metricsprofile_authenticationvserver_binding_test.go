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

// Participating-entity config below is lifted from the existing acceptance tests:
//   - metricsprofile        -> metricsprofile_test.go (testAccMetricsprofile_basic_step1)
//   - authenticationvserver -> authenticationvserver_test.go (testAccAuthenticationvserver_add)
//
// This binding has no NITRO update endpoint (all attributes are RequiresReplace),
// so the basic test uses a single create+verify step plus a destroy step that
// removes the binding while keeping the participating entities.

const testAccMetricsprofile_authenticationvserver_binding_basic_step1 = `

resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "Hello"
  authentication = "ON"
  state          = "ENABLED"
}

resource "citrixadc_metricsprofile_authenticationvserver_binding" "tf_metricsprofile_authenticationvserver_binding" {
  name       = citrixadc_metricsprofile.tf_metricsprofile.name
  entityname = citrixadc_authenticationvserver.tf_authenticationvserver.name
  entitytype = "authvserver"
  depends_on = [
    citrixadc_metricsprofile.tf_metricsprofile,
    citrixadc_authenticationvserver.tf_authenticationvserver,
  ]
}

`

// Step 2 keeps the participating entities but removes the binding to verify clean deletion.
const testAccMetricsprofile_authenticationvserver_binding_basic_step2 = `

resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "Hello"
  authentication = "ON"
  state          = "ENABLED"
}

`

func TestAccMetricsprofile_authenticationvserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofile_authenticationvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_authenticationvserver_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_authenticationvserver_bindingExist("citrixadc_metricsprofile_authenticationvserver_binding.tf_metricsprofile_authenticationvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_authenticationvserver_binding.tf_metricsprofile_authenticationvserver_binding", "name", "tf_metricsprofile"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_authenticationvserver_binding.tf_metricsprofile_authenticationvserver_binding", "entityname", "tf_authenticationvserver"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_authenticationvserver_binding.tf_metricsprofile_authenticationvserver_binding", "entitytype", "authvserver"),
				),
			},
			{
				Config: testAccMetricsprofile_authenticationvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_authenticationvserver_bindingNotExist("tf_metricsprofile", "tf_authenticationvserver"),
				),
			},
		},
	})
}

func testAccCheckMetricsprofile_authenticationvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No metricsprofile_authenticationvserver_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

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

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_authenticationvserver_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching entityname
		found := false
		for _, v := range dataArr {
			if val, ok := v["entityname"].(string); ok && val == entityname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("metricsprofile_authenticationvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_authenticationvserver_bindingNotExist(name string, entityname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_authenticationvserver_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching entityname
		found := false
		for _, v := range dataArr {
			if val, ok := v["entityname"].(string); ok && val == entityname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("metricsprofile_authenticationvserver_binding (%s,%s) was found, but it should have been destroyed", name, entityname)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_authenticationvserver_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_metricsprofile_authenticationvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		entityname := idMap["entityname"]

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_authenticationvserver_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent metricsprofile gone or no bindings - binding is destroyed
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["entityname"].(string); ok && val == entityname {
				return fmt.Errorf("metricsprofile_authenticationvserver_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccMetricsprofile_authenticationvserver_bindingDataSource_basic = `

resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile_ds"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver_ds"
  servicetype    = "SSL"
  comment        = "DataSource Test"
  authentication = "ON"
  state          = "ENABLED"
}

resource "citrixadc_metricsprofile_authenticationvserver_binding" "tf_metricsprofile_authenticationvserver_binding" {
  name       = citrixadc_metricsprofile.tf_metricsprofile.name
  entityname = citrixadc_authenticationvserver.tf_authenticationvserver.name
  entitytype = "authvserver"
  depends_on = [
    citrixadc_metricsprofile.tf_metricsprofile,
    citrixadc_authenticationvserver.tf_authenticationvserver,
  ]
}

data "citrixadc_metricsprofile_authenticationvserver_binding" "tf_metricsprofile_authenticationvserver_binding" {
  name       = citrixadc_metricsprofile_authenticationvserver_binding.tf_metricsprofile_authenticationvserver_binding.name
  entityname = citrixadc_metricsprofile_authenticationvserver_binding.tf_metricsprofile_authenticationvserver_binding.entityname
  entitytype = citrixadc_metricsprofile_authenticationvserver_binding.tf_metricsprofile_authenticationvserver_binding.entitytype
  depends_on = [citrixadc_metricsprofile_authenticationvserver_binding.tf_metricsprofile_authenticationvserver_binding]
}
`

func TestAccMetricsprofile_authenticationvserver_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_authenticationvserver_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_authenticationvserver_binding.tf_metricsprofile_authenticationvserver_binding", "name", "tf_metricsprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_authenticationvserver_binding.tf_metricsprofile_authenticationvserver_binding", "entityname", "tf_authenticationvserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_authenticationvserver_binding.tf_metricsprofile_authenticationvserver_binding", "entitytype", "authvserver"),
				),
			},
		},
	})
}
