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

const testAccLbmetrictable_metric_binding_basic = `

	resource "citrixadc_lbmetrictable" "Table" {
		metrictable = "Table-Custom"
	}
	resource "citrixadc_lbmetrictable_metric_binding" "tf_bind" {
		metric      = "2.3.6.4.5"
		metrictable = citrixadc_lbmetrictable.Table.metrictable
		snmpoid     = "1.2.3.6.5"
	}
`

const testAccLbmetrictable_metric_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_lbmetrictable" "Table" {
		metrictable = "Table-Custom"
	}
`

func TestAccLbmetrictable_metric_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmetrictable_metric_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmetrictable_metric_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmetrictable_metric_bindingExist("citrixadc_lbmetrictable_metric_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccLbmetrictable_metric_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmetrictable_metric_bindingNotExist("citrixadc_lbmetrictable_metric_binding.tf_bind", "Table-Custom,2.3.6.4.5"),
				),
			},
		},
	})
}

func TestAccLbmetrictable_metric_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbmetrictable_metric_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmetrictable_metric_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbmetrictable_metric_binding_basic},
			{Config: testAccLbmetrictable_metric_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckLbmetrictable_metric_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbmetrictable_metric_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"metrictable", "metric"}, nil)
		if err != nil {
			return err
		}
		metrictable := idMap["metrictable"]
		metric := idMap["metric"]

		findParams := service.FindParams{
			ResourceType:             "lbmetrictable_metric_binding",
			ResourceName:             metrictable,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["metric"].(string) == metric {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lbmetrictable_metric_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbmetrictable_metric_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		metrictable := idSlice[0]
		metric := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lbmetrictable_metric_binding",
			ResourceName:             metrictable,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["metric"].(string) == metric {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lbmetrictable_metric_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbmetrictable_metric_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbmetrictable_metric_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbmetrictable_metric_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbmetrictable_metric_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbmetrictable_metric_bindingDataSource_basic = `

resource "citrixadc_lbmetrictable" "Table" {
	metrictable = "Table-Custom"
}
resource "citrixadc_lbmetrictable_metric_binding" "tf_bind" {
	metric      = "2.3.6.4.5"
	metrictable = citrixadc_lbmetrictable.Table.metrictable
	snmpoid     = "1.2.3.6.5"
}

data "citrixadc_lbmetrictable_metric_binding" "tf_bind" {
	metric      = citrixadc_lbmetrictable_metric_binding.tf_bind.metric
	metrictable = citrixadc_lbmetrictable_metric_binding.tf_bind.metrictable
	depends_on  = [citrixadc_lbmetrictable_metric_binding.tf_bind]
}
`

func TestAccLbmetrictable_metric_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmetrictable_metric_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbmetrictable_metric_binding.tf_bind", "metric", "2.3.6.4.5"),
					resource.TestCheckResourceAttr("data.citrixadc_lbmetrictable_metric_binding.tf_bind", "metrictable", "Table-Custom"),
					resource.TestCheckResourceAttr("data.citrixadc_lbmetrictable_metric_binding.tf_bind", "snmpoid", "1.2.3.6.5"),
				),
			},
		},
	})
}

const testAccLbmetrictable_metric_binding_upgrade_basic = `

	resource "citrixadc_lbmetrictable" "Table" {
		metrictable = "Table-Custom"
	}
	resource "citrixadc_lbmetrictable_metric_binding" "tf_bind" {
		metric      = "2.3.6.4.5"
		metrictable = citrixadc_lbmetrictable.Table.metrictable
		snmpoid     = "1.2.3.6.5"
	}
`

func TestAccLbmetrictable_metric_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbmetrictable_metric_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create the resource with the last SDK v2 release (2.2.0),
			// which writes state with the legacy comma-separated ID.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLbmetrictable_metric_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmetrictable_metric_bindingExist("citrixadc_lbmetrictable_metric_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmetrictable_metric_binding.tf_bind", "id", "Table-Custom,2.3.6.4.5"),
				),
			},
			// Step 2: Refresh/plan/apply the legacy-ID state through the current
			// (framework) provider. The framework Read exercises ParseIdString on
			// the legacy ID and recomputes the ID to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbmetrictable_metric_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmetrictable_metric_bindingExist("citrixadc_lbmetrictable_metric_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmetrictable_metric_binding.tf_bind", "id", "metric:2.3.6.4.5,metrictable:Table-Custom"),
				),
			},
		},
	})
}
