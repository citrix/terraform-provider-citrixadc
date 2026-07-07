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

// NOTE on the nssourceroutecachetable resource:
//   - Models the NITRO POST /nssourceroutecachetable?action=flush endpoint plus
//     a keyless get(all)/count. This is a ZERO-ATTRIBUTE, ACTION-ONLY resource:
//     Create performs the flush action, and Read/Update/Delete are no-ops.
//   - Flushing the source-route cache table is low-risk, so the test runs by
//     default.
//   - The Exist check below only verifies the resource landed in Terraform state
//     with its synthetic ID ("nssourceroutecachetable-config").

const testAccNssourceroutecachetable_basic = `
resource "citrixadc_nssourceroutecachetable" "tf_nssourceroutecachetable" {
}
`

func TestAccNssourceroutecachetable_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the flush action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccNssourceroutecachetable_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNssourceroutecachetableExist("citrixadc_nssourceroutecachetable.tf_nssourceroutecachetable", nil),
					resource.TestCheckResourceAttrSet("citrixadc_nssourceroutecachetable.tf_nssourceroutecachetable", "id"),
				),
			},
		},
	})
}

// testAccCheckNssourceroutecachetableExist is a state-only existence check.
func testAccCheckNssourceroutecachetableExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nssourceroutecachetable ID is set")
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

// Datasource: nssourceroutecachetable get(all) returns a TABLE (list) of cache
// entries. On an idle testbed the table is empty and the get(all) errors
// ("no resource found"), so this test is skipped by default. Remove the t.Skip
// only when the source-route cache is known to be populated.
const testAccNssourceroutecachetableDataSource_basic = `
data "citrixadc_nssourceroutecachetable" "tf_nssourceroutecachetable" {
}
`

func TestAccNssourceroutecachetableDataSource_basic(t *testing.T) {
	t.Skip("nssourceroutecachetable get(all) returns a cache table that is empty on an idle testbed; the datasource Read errors on an empty result.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNssourceroutecachetableDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_nssourceroutecachetable.tf_nssourceroutecachetable", "id"),
				),
			},
		},
	})
}
