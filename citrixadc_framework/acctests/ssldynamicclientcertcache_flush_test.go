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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// NOTE on the ssldynamicclientcertcache_flush resource:
//   - Models the NITRO POST /ssldynamicclientcertcache?action=flush endpoint.
//     This is a ZERO-ATTRIBUTE, ACTION-ONLY resource: Create performs the flush
//     action, and Read/Update/Delete are no-ops. Because the object is
//     action-only, the TF resource is named with the action suffix
//     (citrixadc_ssldynamicclientcertcache_flush), mirroring
//     citrixadc_appfwarchive_export / citrixadc_nsmemrecovery_start.
//   - The backing service/config/ssl.Ssldynamicclientcertcache{} struct is empty
//     and NITRO exposes only the flush operation (no add/set/delete/get), so
//     there is no datasource for this object.
//   - Flushing the dynamic client-certificate cache is low-risk, so the test
//     runs by default.
//   - The Exist check below only verifies the resource landed in Terraform state
//     with its synthetic ID ("ssldynamicclientcertcache").

const testAccSsldynamicclientcertcacheFlush_basic = `
resource "citrixadc_ssldynamicclientcertcache_flush" "tf_ssldynamicclientcertcache_flush" {
}
`

func TestAccSsldynamicclientcertcacheFlush_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the flush action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccSsldynamicclientcertcacheFlush_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSsldynamicclientcertcacheFlushExist("citrixadc_ssldynamicclientcertcache_flush.tf_ssldynamicclientcertcache_flush", nil),
					resource.TestCheckResourceAttrSet("citrixadc_ssldynamicclientcertcache_flush.tf_ssldynamicclientcertcache_flush", "id"),
				),
			},
		},
	})
}

// testAccCheckSsldynamicclientcertcacheFlushExist is a state-only existence
// check. The flush action has no NITRO GET endpoint, so it only verifies the
// resource landed in Terraform state with its synthetic ID.
func testAccCheckSsldynamicclientcertcacheFlushExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ssldynamicclientcertcache_flush ID is set")
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
