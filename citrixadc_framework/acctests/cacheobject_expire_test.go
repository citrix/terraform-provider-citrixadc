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

// NOTE on the cacheobject_expire resource:
//   - cacheobject is an ACTION-ONLY runtime object of the NetScaler integrated
//     cache (IC). NITRO exposes ONLY get(all), count, and the POST actions
//     expire | flush | save. There is NO add, NO update/set, and NO delete.
//     Cached objects are created by the traffic engine, not the config API.
//   - citrixadc_cacheobject_expire fires the `?action=expire` POST on Create and
//     treats Read/Update/Delete as no-ops. Because there is no GET-by-id
//     endpoint, the resource CANNOT be verified by reading it back from the ADC.
//     testAccCheckCacheobjectExpireExist only asserts the synthetic Terraform
//     state ID ("cacheobject_expire") is set (mirrors gslbconfig_test.go /
//     clusterfiles_test.go).
//   - There is NO CheckDestroy: the action has no inverse on NITRO and there is
//     no GET-by-id to confirm absence; Delete is a state-only removal.
//   - ValidateConfig enforces (locator) XOR (url + host); the config below uses
//     the url+host form.

// Targeted expire: url + host identify the object to expire. This satisfies the
// resource's ValidateConfig (locator XOR url+host).
const testAccCacheobjectExpire_basic = `
resource "citrixadc_cacheobject_expire" "tf_cacheobject_expire" {
  url  = "/image.gif"
  host = "www.example.com"
}

`

func TestAccCacheobjectExpire_basic(t *testing.T) {
	t.Skip("TODO: Requires review - requires Integrated Caching enabled and live cached objects; argument-less flush is rejected (NITRO errorcode 1095)")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: cacheobject_expire has no delete/GET-by-id endpoint (action-only).
		Steps: []resource.TestStep{
			{
				Config: testAccCacheobjectExpire_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCacheobjectExpireExist("citrixadc_cacheobject_expire.tf_cacheobject_expire", nil),
					resource.TestCheckResourceAttrSet("citrixadc_cacheobject_expire.tf_cacheobject_expire", "id"),
					resource.TestCheckResourceAttr("citrixadc_cacheobject_expire.tf_cacheobject_expire", "id", "cacheobject_expire"),
					resource.TestCheckResourceAttr("citrixadc_cacheobject_expire.tf_cacheobject_expire", "url", "/image.gif"),
					resource.TestCheckResourceAttr("citrixadc_cacheobject_expire.tf_cacheobject_expire", "host", "www.example.com"),
				),
			},
		},
	})
}

// testAccCheckCacheobjectExpireExist is a state-only existence check.
//
// cacheobject_expire is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the fired action via NITRO. We only
// assert that Terraform recorded the resource in state with a non-empty ID
// (which equals the synthetic "cacheobject_expire" after a successful POST
// ?action=expire). This mirrors testAccCheckGslbconfigExist.
func testAccCheckCacheobjectExpireExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cacheobject_expire ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against; presence of the synthetic state
		// ID is the only confirmation we can make for an action-only resource.
		return nil
	}
}
