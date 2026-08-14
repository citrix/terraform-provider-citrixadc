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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// opoption is an alias of the snmpoption unnamed singleton. All attributes are
// TypeString ENABLED/DISABLED style options; there is no NITRO ADD/DELETE.
const testAccOpoption_basic = `

	resource "citrixadc_opoption" "tf_opoption" {
		snmpset              = "ENABLED"
		snmptraplogging      = "ENABLED"
		partitionnameintrap  = "ENABLED"
		snmptraplogginglevel = "WARNING"
		severityinfointrap   = "ENABLED"
	}
`

func TestAccOpoption_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil, // unnamed singleton: no NITRO DELETE, nothing to assert destroyed
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccOpoption_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOpoptionExist("citrixadc_opoption.tf_opoption", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccOpoption_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOpoptionExist("citrixadc_opoption.tf_opoption", nil),
				),
			},
		},
	})
}

func testAccCheckOpoptionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No opoption name is set")
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
		// opoption maps to the unnamed snmpoption singleton (key is "").
		data, err := client.FindResource(service.Snmpoption.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("opoption %s not found", n)
		}

		return nil
	}
}

// testAccCheckOpoptionDestroy is provided for parity with other custom_resources
// tests. opoption is an unnamed singleton with no NITRO DELETE, so the object
// always remains on the appliance; there is nothing to assert as destroyed.
func testAccCheckOpoptionDestroy(s *terraform.State) error {
	return nil
}
