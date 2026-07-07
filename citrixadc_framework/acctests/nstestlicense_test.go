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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NOTE on the nstestlicense resource:
//   - Models the NITRO POST /nstestlicense?action=apply endpoint plus a keyless
//     get(all). This is a ZERO-ATTRIBUTE, ACTION-ONLY resource: Create performs
//     the apply action, and Read/Update/Delete are no-ops.
//   - The Exist check below only verifies the resource landed in Terraform state
//     with its synthetic ID ("nstestlicense-config").
//
// !!! DESTRUCTIVE !!!
// Applying this resource APPLIES A TEST/EVAL LICENSE to the appliance, which
// changes the licensed feature set and can disrupt the running configuration.
// The resource test is therefore SKIPPED by default. Remove the t.Skip line only
// when you intend to apply a test license against a disposable ADC on purpose.

const testAccNstestlicense_basic = `
resource "citrixadc_nstestlicense" "tf_nstestlicense" {
}
`

func TestAccNstestlicense_basic(t *testing.T) {
	t.Skip("DESTRUCTIVE: applying citrixadc_nstestlicense applies a test/eval license and changes the licensed feature set. Remove this skip only to run intentionally against a disposable ADC.")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNstestlicense_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstestlicenseExist("citrixadc_nstestlicense.tf_nstestlicense", nil),
					resource.TestCheckResourceAttrSet("citrixadc_nstestlicense.tf_nstestlicense", "id"),
				),
			},
		},
	})
}

// testAccCheckNstestlicenseExist is a state-only existence check.
func testAccCheckNstestlicenseExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nstestlicense ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		return nil
	}
}

// Datasource: nstestlicense has a keyless get(all) that returns the license
// feature-flag object. Only the synthetic "id" is stably assertable. The
// datasource is read-only and non-destructive (it does NOT apply a license).
const testAccNstestlicenseDataSource_basic = `
data "citrixadc_nstestlicense" "tf_nstestlicense" {
}
`

func TestAccNstestlicenseDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNstestlicenseDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_nstestlicense.tf_nstestlicense", "id"),
				),
			},
		},
	})
}
