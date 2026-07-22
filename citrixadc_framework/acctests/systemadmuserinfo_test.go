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

// systemadmuserinfo is an update-only, no-GET resource:
//   - Create/Update = NITRO `update` (PUT) of system/systemadmuserinfo, no `add`
//   - Read         = no-op (NITRO exposes no GET endpoint, so no drift detection)
//   - Delete       = state-only removal (no NITRO delete endpoint)
//
// Because there is no GET endpoint, the Exist check below is STATE-ONLY: it
// verifies the resource is present in Terraform state with a non-empty ID and
// does NOT call client.FindResource (there is nothing to read back). For the
// same reason there is no CheckDestroy (nothing to delete on the ADC).
// There is also no datasource for this resource, so no datasource test exists.

const testAccSystemadmuserinfo_basic_step1 = `
resource "citrixadc_systemadmuserinfo" "tf_systemadmuserinfo" {
  username = "nsroot"
}
`

const testAccSystemadmuserinfo_basic_step2 = `
resource "citrixadc_systemadmuserinfo" "tf_systemadmuserinfo" {
  username = "tf_admuser"
}
`

func TestAccSystemadmuserinfo_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: systemadmuserinfo has no NITRO delete endpoint;
		// Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccSystemadmuserinfo_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemadmuserinfoExist("citrixadc_systemadmuserinfo.tf_systemadmuserinfo", nil),
					resource.TestCheckResourceAttr("citrixadc_systemadmuserinfo.tf_systemadmuserinfo", "username", "nsroot"),
				),
			},
			{
				Config: testAccSystemadmuserinfo_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemadmuserinfoExist("citrixadc_systemadmuserinfo.tf_systemadmuserinfo", nil),
					resource.TestCheckResourceAttr("citrixadc_systemadmuserinfo.tf_systemadmuserinfo", "username", "tf_admuser"),
				),
			},
		},
	})
}

// testAccCheckSystemadmuserinfoExist is STATE-ONLY by design.
// systemadmuserinfo has no NITRO GET endpoint, so there is nothing to read
// back from the ADC; we only verify the resource exists in Terraform state
// with a non-empty ID.
func testAccCheckSystemadmuserinfoExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemadmuserinfo ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}
			*id = rs.Primary.ID
		}

		// No NITRO GET endpoint for systemadmuserinfo - state-only check.
		return nil
	}
}
