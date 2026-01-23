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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccSnmpalarm_basic = `
resource "citrixadc_snmpalarm" "tf_snmpalarm" {
	trapname = "CPU-USAGE"
	thresholdvalue = 15
	normalvalue    = 10
	state          = "DISABLED"
	severity       = "Minor"
	}
  
`
const testAccSnmpalarm_update = `
resource "citrixadc_snmpalarm" "tf_snmpalarm" {
	trapname = "CPU-USAGE"
	thresholdvalue = 20
	normalvalue    = 15
	state          = "ENABLED"
	severity       = "Major"
	}
  
`

func TestAccSnmpalarm_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpalarm_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpalarmExist("citrixadc_snmpalarm.tf_snmpalarm", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "thresholdvalue", "15"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "normalvalue", "10"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "state", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "severity", "Minor"),
				),
			},
			{
				Config: testAccSnmpalarm_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpalarmExist("citrixadc_snmpalarm.tf_snmpalarm", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "thresholdvalue", "20"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "normalvalue", "15"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "severity", "Major"),
				),
			},
		},
	})
}

func testAccCheckSnmpalarmExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpalarm name is set")
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
		data, err := client.FindResource(service.Snmpalarm.Type(), rs.Primary.Attributes["trapclass"])

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("snmpalarm %s not found", n)
		}

		return nil
	}
}
