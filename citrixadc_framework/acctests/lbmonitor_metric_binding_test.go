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
		idSlice := strings.SplitN(bindingId, ",", 2)
		lbmonitorName := idSlice[0]
		metricName := idSlice[1]

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
