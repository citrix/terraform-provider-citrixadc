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

// Need the following cli commands since no resource yet exists
// add lb metricTable tab1
// bind metrictable tab1 metric1 1.3.6.1.4.1.5951.4.1.1.8.0

const testAccLbmonitor_metric_binding_basic = `

resource "citrixadc_lbmetrictable" "tab1" {
	metrictable = "tab1"
	}

resource "citrixadc_lbmetrictable_metric_binding" "tf_bind" {
	metric      = "metric1"
	metrictable = citrixadc_lbmetrictable.tab1.metrictable
	snmpoid     = "1.3.6.1.4.1.5951.4.1.1.8.0"
	}
resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "LOAD"
  metrictable = citrixadc_lbmetrictable.tab1.metrictable
}

resource citrixadc_lbmonitor_metric_binding tf_acclbmonitor_metric_binding {
	monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
	metric = citrixadc_lbmetrictable_metric_binding.tf_bind.metric
	metricthreshold = 100
	}
`

const testAccLbmonitor_metric_bindingDataSource_basic = `

resource "citrixadc_lbmetrictable" "tab1" {
	metrictable = "tab1"
	}

resource "citrixadc_lbmetrictable_metric_binding" "tf_bind" {
	metric      = "metric1"
	metrictable = citrixadc_lbmetrictable.tab1.metrictable
	snmpoid     = "1.3.6.1.4.1.5951.4.1.1.8.0"
	}
resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "LOAD"
  metrictable = citrixadc_lbmetrictable.tab1.metrictable
}

resource citrixadc_lbmonitor_metric_binding tf_acclbmonitor_metric_binding {
	monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
	metric = citrixadc_lbmetrictable_metric_binding.tf_bind.metric
	metricthreshold = 100
	}

data "citrixadc_lbmonitor_metric_binding" "tf_acclbmonitor_metric_binding" {
	monitorname = citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding.monitorname
	metric = citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding.metric
	depends_on = [citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding]
}
`

func TestAccLbmonitor_metric_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitor_metric_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_metric_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitor_metric_bindingExist("citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding", "monitorname", "tf-monitor1"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding", "metric", "metric1"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding", "metricthreshold", "100"),
				),
			},
		},
	})
}

func testAccCheckLbmonitor_metric_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb monitor metric binding name is set")
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
		idMap, _, err := utils.ParseIdString(bindingId, []string{"monitorname", "metric"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		lbmonitorName := idMap["monitorname"]
		metricName := idMap["metric"]

		findParams := service.FindParams{
			ResourceType:             "lbmonitor_metric_binding",
			ResourceName:             lbmonitorName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right metric name
		foundIndex := -1
		for i, v := range dataArr {
			if v["metric"].(string) == metricName {
				foundIndex = i
				break
			}
		}
		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find lbmonitor_metric_binding ID %v", bindingId)
		}

		return nil
	}
}

func testAccCheckLbmonitor_metric_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbmonitor_metric_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbmonitor_metric_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB monitor metric binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLbmonitor_metric_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_metric_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding", "monitorname", "tf-monitor1"),
					resource.TestCheckResourceAttr("data.citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding", "metric", "metric1"),
					resource.TestCheckResourceAttr("data.citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding", "metricthreshold", "100"),
				),
			},
		},
	})
}

// testAccLbmonitor_metric_binding_upgrade_basic reuses the working _basic config values
// (prerequisite lbmetrictable + lbmetrictable_metric_binding + lbmonitor plus the binding).
// It must be valid under BOTH the SDK v2 2.2.0 schema (step 1) and the current provider
// schema (step 2), so it uses the SDK v2 attribute names.
const testAccLbmonitor_metric_binding_upgrade_basic = `

resource "citrixadc_lbmetrictable" "tab1" {
	metrictable = "tab1"
	}

resource "citrixadc_lbmetrictable_metric_binding" "tf_bind" {
	metric      = "metric1"
	metrictable = citrixadc_lbmetrictable.tab1.metrictable
	snmpoid     = "1.3.6.1.4.1.5951.4.1.1.8.0"
	}
resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "LOAD"
  metrictable = citrixadc_lbmetrictable.tab1.metrictable
}

resource citrixadc_lbmonitor_metric_binding tf_acclbmonitor_metric_binding {
	monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
	metric = citrixadc_lbmetrictable_metric_binding.tf_bind.metric
	metricthreshold = 100
	}
`

// TestAccLbmonitor_metric_binding_sdkv2StateUpgrade creates the binding with the last SDK v2
// provider release (2.2.0) — which writes state with the legacy comma-joined id
// (monitorname,metric) — then refreshes/applies that legacy-id state through the current
// framework provider. Step 2's Read exercises ParseIdString on the legacy id and
// SetAttrFromGet recomputes the id into the new key:value form (metric:...,monitorname:...).
func TestAccLbmonitor_metric_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbmonitor_metric_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy comma-joined id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLbmonitor_metric_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitor_metric_bindingExist("citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding", "id", "tf-monitor1,metric1"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current framework
				// provider. Read exercises ParseIdString on the legacy id and
				// SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbmonitor_metric_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitor_metric_bindingExist("citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding", "id", "metric:metric1,monitorname:tf-monitor1"),
				),
			},
		},
	})
}

func TestAccLbmonitor_metric_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbmonitor_metric_binding.tf_acclbmonitor_metric_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitor_metric_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbmonitor_metric_binding_basic},
			{Config: testAccLbmonitor_metric_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
