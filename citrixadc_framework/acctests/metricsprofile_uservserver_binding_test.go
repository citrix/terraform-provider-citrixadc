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

// Step 1 creates the metricsprofile and the uservserver (participating entities, lifted from
// metricsprofile_test.go and uservserver_test.go) and binds them together.
// entitytype is the constant "uservserver"; entityname references the uservserver name; name
// references the metricsprofile name.
//
// NOTE: The uservserver participating entity requires a userprotocol backed by an NS Lua
// extension to be loaded on the ADC (see uservserver_test.go, which is skipped for the same
// reason). The userprotocol/lbvserver/uservserver blocks below are lifted verbatim from
// testAccUservserver_basic. Because of that prerequisite, these tests are skipped by default;
// remove the t.Skip once the "mqtt_code" extension is available on the testbed.
const testAccMetricsprofile_uservserver_binding_basic_step1 = `
	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "MQTT"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}

	resource "citrixadc_lbvserver" "tf_defaultlb" {
		name        = "tf_defaultlb"
		servicetype = "USER_TCP"
	}

	resource "citrixadc_uservserver" "tf_uservserver" {
		name         = "my_user_vserver"
		userprotocol = "MQTT"
		ipaddress    = "10.222.74.180"
		port         = 3200
		defaultlb    = citrixadc_lbvserver.tf_defaultlb.name
		depends_on   = [citrixadc_userprotocol.tf_userprotocol, citrixadc_lbvserver.tf_defaultlb]
	}

	resource "citrixadc_metricsprofile" "tf_metricsprofile" {
		name                   = "tf_metricsprofile_uservbind"
		outputmode             = "avro"
		metrics                = "ENABLED"
		servemode              = "Push"
		metricsexportfrequency = 30
	}

	resource "citrixadc_metricsprofile_uservserver_binding" "tf_metricsprofile_uservserver_binding" {
		name       = citrixadc_metricsprofile.tf_metricsprofile.name
		entityname = citrixadc_uservserver.tf_uservserver.name
		entitytype = "uservserver"

		depends_on = [
			citrixadc_metricsprofile.tf_metricsprofile,
			citrixadc_uservserver.tf_uservserver,
		]
	}
`

// Step 2 drops the binding but keeps the participating entities, so CheckDestroy / NotExist can
// confirm the binding was removed from the ADC.
const testAccMetricsprofile_uservserver_binding_basic_step2 = `
	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "MQTT"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}

	resource "citrixadc_lbvserver" "tf_defaultlb" {
		name        = "tf_defaultlb"
		servicetype = "USER_TCP"
	}

	resource "citrixadc_uservserver" "tf_uservserver" {
		name         = "my_user_vserver"
		userprotocol = "MQTT"
		ipaddress    = "10.222.74.180"
		port         = 3200
		defaultlb    = citrixadc_lbvserver.tf_defaultlb.name
		depends_on   = [citrixadc_userprotocol.tf_userprotocol, citrixadc_lbvserver.tf_defaultlb]
	}

	resource "citrixadc_metricsprofile" "tf_metricsprofile" {
		name                   = "tf_metricsprofile_uservbind"
		outputmode             = "avro"
		metrics                = "ENABLED"
		servemode              = "Push"
		metricsexportfrequency = 30
	}
`

func TestAccMetricsprofile_uservserver_binding_basic(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension for the uservserver participating entity. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofile_uservserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_uservserver_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_uservserver_bindingExist("citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding", "name", "tf_metricsprofile_uservbind"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding", "entityname", "my_user_vserver"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding", "entitytype", "uservserver"),
				),
			},
			{
				Config: testAccMetricsprofile_uservserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_uservserver_bindingNotExist("citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding", "entityname:my_user_vserver,entitytype:uservserver,name:tf_metricsprofile_uservbind"),
				),
			},
		},
	})
}

func testAccCheckMetricsprofile_uservserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No metricsprofile_uservserver_binding id is set")
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
			ResourceType:             service.Metricsprofile_uservserver_binding.Type(),
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
			return fmt.Errorf("metricsprofile_uservserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_uservserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
			ResourceType:             service.Metricsprofile_uservserver_binding.Type(),
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
			return fmt.Errorf("metricsprofile_uservserver_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_uservserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_metricsprofile_uservserver_binding" {
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

		_, err = client.FindResource(service.Metricsprofile_uservserver_binding.Type(), name)
		if err == nil {
			return fmt.Errorf("metricsprofile_uservserver_binding still exists")
		}
	}

	return nil
}

const testAccMetricsprofile_uservserver_bindingDataSource_basic = `
	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "MQTT"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}

	resource "citrixadc_lbvserver" "tf_defaultlb" {
		name        = "tf_defaultlb"
		servicetype = "USER_TCP"
	}

	resource "citrixadc_uservserver" "tf_uservserver_ds" {
		name         = "my_user_vserver_ds"
		userprotocol = "MQTT"
		ipaddress    = "10.222.74.181"
		port         = 3200
		defaultlb    = citrixadc_lbvserver.tf_defaultlb.name
		depends_on   = [citrixadc_userprotocol.tf_userprotocol, citrixadc_lbvserver.tf_defaultlb]
	}

	resource "citrixadc_metricsprofile" "tf_metricsprofile_ds" {
		name                   = "tf_metricsprofile_uservbind_ds"
		outputmode             = "avro"
		metrics                = "ENABLED"
		servemode              = "Push"
		metricsexportfrequency = 30
	}

	resource "citrixadc_metricsprofile_uservserver_binding" "tf_metricsprofile_uservserver_binding_ds" {
		name       = citrixadc_metricsprofile.tf_metricsprofile_ds.name
		entityname = citrixadc_uservserver.tf_uservserver_ds.name
		entitytype = "uservserver"

		depends_on = [
			citrixadc_metricsprofile.tf_metricsprofile_ds,
			citrixadc_uservserver.tf_uservserver_ds,
		]
	}

	data "citrixadc_metricsprofile_uservserver_binding" "tf_metricsprofile_uservserver_binding_ds_read" {
		name       = citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding_ds.name
		entityname = citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding_ds.entityname
		entitytype = citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding_ds.entitytype
		depends_on = [citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding_ds]
	}
`

func TestAccMetricsprofile_uservserver_bindingDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension for the uservserver participating entity. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_uservserver_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding_ds_read", "name", "tf_metricsprofile_uservbind_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding_ds_read", "entityname", "my_user_vserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding_ds_read", "entitytype", "uservserver"),
					resource.TestCheckResourceAttrSet("data.citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding_ds_read", "id"),
				),
			},
		},
	})
}
