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
//   - metricsprofile -> metricsprofile_test.go (testAccMetricsprofile_basic_step1)
//   - crvserver      -> crvserver_test.go (testAccCrvserver_add)
//
// This binding has no NITRO update endpoint (all attributes are RequiresReplace),
// so the basic test uses a single create+verify step plus a destroy step that
// removes the binding while keeping the participating entities.

const testAccMetricsprofile_crvserver_binding_basic_step1 = `

resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_crvserver" "tf_crvserver" {
  name        = "tf_crvserver"
  servicetype = "HTTP"
  arp         = "OFF"
}

resource "citrixadc_metricsprofile_crvserver_binding" "tf_metricsprofile_crvserver_binding" {
  name       = citrixadc_metricsprofile.tf_metricsprofile.name
  entityname = citrixadc_crvserver.tf_crvserver.name
  entitytype = "crvserver"
  depends_on = [
    citrixadc_metricsprofile.tf_metricsprofile,
    citrixadc_crvserver.tf_crvserver,
  ]
}

`

// Step 2 keeps the participating entities but removes the binding to verify clean deletion.
const testAccMetricsprofile_crvserver_binding_basic_step2 = `

resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_crvserver" "tf_crvserver" {
  name        = "tf_crvserver"
  servicetype = "HTTP"
  arp         = "OFF"
}

`

func TestAccMetricsprofile_crvserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofile_crvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_crvserver_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_crvserver_bindingExist("citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding", "name", "tf_metricsprofile"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding", "entityname", "tf_crvserver"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding", "entitytype", "crvserver"),
				),
			},
			{
				Config: testAccMetricsprofile_crvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_crvserver_bindingNotExist("tf_metricsprofile", "tf_crvserver"),
				),
			},
		},
	})
}

func testAccCheckMetricsprofile_crvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No metricsprofile_crvserver_binding id is set")
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
			ResourceType:             service.Metricsprofile_crvserver_binding.Type(),
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
			return fmt.Errorf("metricsprofile_crvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_crvserver_bindingNotExist(name string, entityname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_crvserver_binding.Type(),
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
			return fmt.Errorf("metricsprofile_crvserver_binding (%s,%s) was found, but it should have been destroyed", name, entityname)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_crvserver_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_metricsprofile_crvserver_binding" {
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
			ResourceType:             service.Metricsprofile_crvserver_binding.Type(),
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
				return fmt.Errorf("metricsprofile_crvserver_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccMetricsprofile_crvserver_bindingDataSource_basic = `

resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile_ds"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_crvserver" "tf_crvserver" {
  name        = "tf_crvserver_ds"
  servicetype = "HTTP"
  arp         = "OFF"
}

resource "citrixadc_metricsprofile_crvserver_binding" "tf_metricsprofile_crvserver_binding" {
  name       = citrixadc_metricsprofile.tf_metricsprofile.name
  entityname = citrixadc_crvserver.tf_crvserver.name
  entitytype = "crvserver"
  depends_on = [
    citrixadc_metricsprofile.tf_metricsprofile,
    citrixadc_crvserver.tf_crvserver,
  ]
}

data "citrixadc_metricsprofile_crvserver_binding" "tf_metricsprofile_crvserver_binding" {
  name       = citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding.name
  entityname = citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding.entityname
  entitytype = citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding.entitytype
  depends_on = [citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding]
}
`

func TestAccMetricsprofile_crvserver_binding_import(t *testing.T) {
	const resAddr = "citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofile_crvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_crvserver_binding_basic_step1,
			},
			{
				Config:                  testAccMetricsprofile_crvserver_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccMetricsprofile_crvserver_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_crvserver_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding", "name", "tf_metricsprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding", "entityname", "tf_crvserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding", "entitytype", "crvserver"),
				),
			},
		},
	})
}
