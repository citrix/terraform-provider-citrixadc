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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Step 1 creates the metricsprofile and the csvserver (participating entities, lifted from
// metricsprofile_test.go and csvserver_test.go) and binds them together.
// entitytype is the constant "csvserver"; entityname references the csvserver name; name
// references the metricsprofile name.
const testAccMetricsprofile_csvserver_binding_basic_step1 = `
	resource "citrixadc_metricsprofile" "tf_metricsprofile" {
		name                   = "tf_metricsprofile_csvbind"
		outputmode             = "avro"
		metrics                = "ENABLED"
		servemode              = "Push"
		metricsexportfrequency = 30
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name        = "tf_csvserver_metricsbind"
		ipv46       = "10.202.11.11"
		port        = 8080
		servicetype = "HTTP"
	}

	resource "citrixadc_metricsprofile_csvserver_binding" "tf_metricsprofile_csvserver_binding" {
		name       = citrixadc_metricsprofile.tf_metricsprofile.name
		entityname = citrixadc_csvserver.tf_csvserver.name
		entitytype = "csvserver"

		depends_on = [
			citrixadc_metricsprofile.tf_metricsprofile,
			citrixadc_csvserver.tf_csvserver,
		]
	}
`

// Step 2 drops the binding but keeps the participating entities, so CheckDestroy / NotExist can
// confirm the binding was removed from the ADC.
const testAccMetricsprofile_csvserver_binding_basic_step2 = `
	resource "citrixadc_metricsprofile" "tf_metricsprofile" {
		name                   = "tf_metricsprofile_csvbind"
		outputmode             = "avro"
		metrics                = "ENABLED"
		servemode              = "Push"
		metricsexportfrequency = 30
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name        = "tf_csvserver_metricsbind"
		ipv46       = "10.202.11.11"
		port        = 8080
		servicetype = "HTTP"
	}
`

func TestAccMetricsprofile_csvserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofile_csvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_csvserver_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_csvserver_bindingExist("citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding", "name", "tf_metricsprofile_csvbind"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding", "entityname", "tf_csvserver_metricsbind"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding", "entitytype", "csvserver"),
				),
			},
			{
				Config: testAccMetricsprofile_csvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_csvserver_bindingNotExist("citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding", "entityname:tf_csvserver_metricsbind,entitytype:csvserver,name:tf_metricsprofile_csvbind"),
				),
			},
		},
	})
}

func TestAccMetricsprofile_csvserver_binding_import(t *testing.T) {
	const resAddr = "citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofile_csvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_csvserver_binding_basic_step1,
			},
			{
				Config:                  testAccMetricsprofile_csvserver_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckMetricsprofile_csvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No metricsprofile_csvserver_binding id is set")
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

		// ID is comma-separated key:UrlEncode(value) pairs: entityname:<v>,entitytype:<v>,name:<v>
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		entityname := idMap["entityname"]
		entitytype := idMap["entitytype"]

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_csvserver_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one matching entityname + entitytype
		found := false
		for _, v := range dataArr {
			if v["entityname"] == entityname && v["entitytype"] == entitytype {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("metricsprofile_csvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_csvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// ID is comma-separated key:UrlEncode(value) pairs: entityname:<v>,entitytype:<v>,name:<v>
		idMap, _, err := utils.ParseIdString(id, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}

		name := idMap["name"]
		entityname := idMap["entityname"]
		entitytype := idMap["entitytype"]

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_csvserver_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one matching entityname + entitytype
		found := false
		for _, v := range dataArr {
			if v["entityname"] == entityname && v["entitytype"] == entitytype {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("metricsprofile_csvserver_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_csvserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_metricsprofile_csvserver_binding" {
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

		_, err = client.FindResource(service.Metricsprofile_csvserver_binding.Type(), name)
		if err == nil {
			return fmt.Errorf("metricsprofile_csvserver_binding still exists")
		}
	}

	return nil
}

const testAccMetricsprofile_csvserver_bindingDataSource_basic = `
	resource "citrixadc_metricsprofile" "tf_metricsprofile_ds" {
		name                   = "tf_metricsprofile_csvbind_ds"
		outputmode             = "avro"
		metrics                = "ENABLED"
		servemode              = "Push"
		metricsexportfrequency = 30
	}

	resource "citrixadc_csvserver" "tf_csvserver_ds" {
		name        = "tf_csvserver_metricsbind_ds"
		ipv46       = "10.202.11.12"
		port        = 8080
		servicetype = "HTTP"
	}

	resource "citrixadc_metricsprofile_csvserver_binding" "tf_metricsprofile_csvserver_binding_ds" {
		name       = citrixadc_metricsprofile.tf_metricsprofile_ds.name
		entityname = citrixadc_csvserver.tf_csvserver_ds.name
		entitytype = "csvserver"

		depends_on = [
			citrixadc_metricsprofile.tf_metricsprofile_ds,
			citrixadc_csvserver.tf_csvserver_ds,
		]
	}

	data "citrixadc_metricsprofile_csvserver_binding" "tf_metricsprofile_csvserver_binding_ds_read" {
		name       = citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding_ds.name
		entityname = citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding_ds.entityname
		entitytype = citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding_ds.entitytype
		depends_on = [citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding_ds]
	}
`

func TestAccMetricsprofile_csvserver_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_csvserver_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding_ds_read", "name", "tf_metricsprofile_csvbind_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding_ds_read", "entityname", "tf_csvserver_metricsbind_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding_ds_read", "entitytype", "csvserver"),
					resource.TestCheckResourceAttrSet("data.citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding_ds_read", "id"),
				),
			},
		},
	})
}
