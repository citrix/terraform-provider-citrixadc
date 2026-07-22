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

const testAccMetricsprofile_vpnvserver_binding_basic_step1 = `
resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_mp_vpnvserver_bind"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserver_mpbind"
  servicetype = "SSL"
}

resource "citrixadc_metricsprofile_vpnvserver_binding" "tf_bind" {
  name       = citrixadc_metricsprofile.tf_metricsprofile.name
  entityname = citrixadc_vpnvserver.tf_vpnvserver.name
  entitytype = "vpnvserver"

  depends_on = [
    citrixadc_metricsprofile.tf_metricsprofile,
    citrixadc_vpnvserver.tf_vpnvserver,
  ]
}
`

const testAccMetricsprofile_vpnvserver_binding_basic_step2 = `
# Keep the participating entities but drop the binding to verify proper deletion
resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_mp_vpnvserver_bind"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserver_mpbind"
  servicetype = "SSL"
}
`

func TestAccMetricsprofile_vpnvserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofile_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_vpnvserver_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_vpnvserver_bindingExist("citrixadc_metricsprofile_vpnvserver_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_vpnvserver_binding.tf_bind", "name", "tf_mp_vpnvserver_bind"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_vpnvserver_binding.tf_bind", "entityname", "tf_vpnvserver_mpbind"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile_vpnvserver_binding.tf_bind", "entitytype", "vpnvserver"),
				),
			},
			{
				Config: testAccMetricsprofile_vpnvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofile_vpnvserver_bindingNotExist("tf_mp_vpnvserver_bind", "tf_vpnvserver_mpbind", "vpnvserver"),
				),
			},
		},
	})
}

func TestAccMetricsprofile_vpnvserver_binding_import(t *testing.T) {
	const resAddr = "citrixadc_metricsprofile_vpnvserver_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofile_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_vpnvserver_binding_basic_step1,
			},
			{
				Config:                  testAccMetricsprofile_vpnvserver_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckMetricsprofile_vpnvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No metricsprofile_vpnvserver_binding id is set")
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
		entitytype := idMap["entitytype"]

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_vpnvserver_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["entityname"] == entityname && v["entitytype"] == entitytype {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("metricsprofile_vpnvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_vpnvserver_bindingNotExist(name, entityname, entitytype string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_vpnvserver_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["entityname"] == entityname && v["entitytype"] == entitytype {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("metricsprofile_vpnvserver_binding (name=%s, entityname=%s, entitytype=%s) was found, but it should have been destroyed", name, entityname, entitytype)
		}

		return nil
	}
}

func testAccCheckMetricsprofile_vpnvserver_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_metricsprofile_vpnvserver_binding" {
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
		entitytype := idMap["entitytype"]

		findParams := service.FindParams{
			ResourceType:             service.Metricsprofile_vpnvserver_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone or no bindings - treat as destroyed
			continue
		}

		for _, v := range dataArr {
			if v["entityname"] == entityname && v["entitytype"] == entitytype {
				return fmt.Errorf("metricsprofile_vpnvserver_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccMetricsprofile_vpnvserver_bindingDataSource_basic = `

resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_mp_vpnvserver_bind_ds"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserver_mpbind_ds"
  servicetype = "SSL"
}

resource "citrixadc_metricsprofile_vpnvserver_binding" "tf_bind" {
  name       = citrixadc_metricsprofile.tf_metricsprofile.name
  entityname = citrixadc_vpnvserver.tf_vpnvserver.name
  entitytype = "vpnvserver"

  depends_on = [
    citrixadc_metricsprofile.tf_metricsprofile,
    citrixadc_vpnvserver.tf_vpnvserver,
  ]
}

data "citrixadc_metricsprofile_vpnvserver_binding" "tf_bind" {
  name       = citrixadc_metricsprofile_vpnvserver_binding.tf_bind.name
  entityname = citrixadc_metricsprofile_vpnvserver_binding.tf_bind.entityname
  entitytype = citrixadc_metricsprofile_vpnvserver_binding.tf_bind.entitytype
  depends_on = [citrixadc_metricsprofile_vpnvserver_binding.tf_bind]
}
`

func TestAccMetricsprofile_vpnvserver_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofile_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_vpnvserver_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_vpnvserver_binding.tf_bind", "name", "tf_mp_vpnvserver_bind_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_vpnvserver_binding.tf_bind", "entityname", "tf_vpnvserver_mpbind_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile_vpnvserver_binding.tf_bind", "entitytype", "vpnvserver"),
				),
			},
		},
	})
}
