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

// NOTE on the cacheobject resource:
//   - cacheobject is an ACTION-ONLY runtime object of the NetScaler integrated
//     cache (IC). NITRO exposes ONLY get(all), count, and the POST actions
//     expire | flush | save. There is NO add, NO update/set, and NO delete.
//     Cached objects are created by the traffic engine, not the config API.
//   - The RESOURCE therefore fires the chosen action (expire|flush|save) on
//     Create and treats Read/Update/Delete as no-ops. Because there is no
//     GET-by-id endpoint, the resource CANNOT be verified by reading it back
//     from the ADC. The Exist check below only asserts the synthetic Terraform
//     state ID is set (mirrors gslbconfig_test.go / clusterfiles_test.go).
//   - There is NO CheckDestroy: the action has no inverse on NITRO and there is
//     no GET-by-id to confirm absence; Delete is a state-only removal.
//   - ValidateConfig enforces: for action expire|flush you must supply either
//     "locator" XOR ("url" + "host"), BUT flush/expire with NO args means
//     "flush/expire ALL" and is accepted. The tests below use the flush-all /
//     save-all form so the apply succeeds with no prerequisite cached objects.
//
// TODO_PLACEHOLDER (prereq): The integrated caching feature must be licensed and
//   enabled on the target ADC before these tests can pass:
//       enable ns feature IC
//   If IC is not enabled, the expire/flush/save actions will error. This cannot
//   be asserted from the test; enable it out-of-band on the testbed.
//
// TODO_PLACEHOLDER (variant not covered): expire/flush with a SPECIFIC locator
//   or url+host requires a LIVE cached object to already exist on the appliance
//   (created by real traffic through a caching content group). That value is
//   testbed-specific and cannot be hard-coded here, so the targeted-expire
//   variant is intentionally omitted. To exercise it manually, drive traffic
//   through an IC content group, read the object's locator from the datasource,
//   then apply:
//       resource "citrixadc_cacheobject" "expire_one" {
//         action  = "expire"
//         locator = <TODO_PLACEHOLDER: live locator id, e.g. 8320>
//       }

// flush-all: action=flush with no locator/url/host. This is the safest testable
// path — it succeeds without any prerequisite cached objects on the appliance.
const testAccCacheobject_basic_step1 = `
resource "citrixadc_cacheobject" "tf_cacheobject" {
  action = "flush"
}

`

// save-all: action=save with no args. Also safe with no prerequisite objects.
// "action" is RequiresReplace, so step2 recreates the resource with a new action.
const testAccCacheobject_basic_step2 = `
resource "citrixadc_cacheobject" "tf_cacheobject" {
  action = "save"
}

`

func TestAccCacheobject_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: cacheobject has no delete/GET-by-id endpoint (action-only).
		Steps: []resource.TestStep{
			{
				Config: testAccCacheobject_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCacheobjectExist("citrixadc_cacheobject.tf_cacheobject", nil),
					// Synthetic ID for a flush-all (no locator/url) is just "flush".
					resource.TestCheckResourceAttrSet("citrixadc_cacheobject.tf_cacheobject", "id"),
					resource.TestCheckResourceAttr("citrixadc_cacheobject.tf_cacheobject", "action", "flush"),
				),
			},
			{
				Config: testAccCacheobject_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCacheobjectExist("citrixadc_cacheobject.tf_cacheobject", nil),
					resource.TestCheckResourceAttrSet("citrixadc_cacheobject.tf_cacheobject", "id"),
					resource.TestCheckResourceAttr("citrixadc_cacheobject.tf_cacheobject", "action", "save"),
				),
			},
		},
	})
}

// testAccCheckCacheobjectExist is a state-only existence check.
//
// cacheobject is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the fired action via NITRO. We only
// assert that Terraform recorded the resource in state with a non-empty ID
// (which equals the synthetic "flush"/"save"/... after a successful POST
// ?action=<verb>). This mirrors testAccCheckGslbconfigExist.
func testAccCheckCacheobjectExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cacheobject ID is set")
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

// Datasource test: reads the cache object list via get(all) and filters locally.
//
// TODO_PLACEHOLDER (prereq): The datasource Read FAILS if the cache object list
//
//	is EMPTY — the implementation returns an error ("cacheobject returned empty
//	array") when NITRO get(all) yields no objects. On a fresh appliance with no
//	traffic through an IC content group, the list IS empty, so this test will
//	NOT pass as-is. Before running it you must:
//	  1. enable ns feature IC   (integrated caching must be licensed/enabled)
//	  2. configure an IC content group + cache policy and drive HTTP traffic
//	     through a vserver so at least one object is cached.
//	Then re-run. Because the cached object's attribute values (locator, url,
//	host, httpstatus, ...) are testbed-specific and non-deterministic, the only
//	assertion made below is that the datasource id is set (read succeeded).
//	Add stricter attribute assertions once a known object is guaranteed present.
const testAccCacheobjectDataSource_basic = `
data "citrixadc_cacheobject" "tf_cacheobject" {
  // No filters: match the first cached object in the get(all) list.
  // Requires at least one cached object to exist (see TODO_PLACEHOLDER above).
}

`

func TestAccCacheobjectDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCacheobjectDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Read succeeded and an ID was composed (locator:/url:/"cacheobject").
					resource.TestCheckResourceAttrSet("data.citrixadc_cacheobject.tf_cacheobject", "id"),
					// TODO_PLACEHOLDER: add attribute assertions once a known cached
					// object is guaranteed on the testbed, e.g.:
					//   resource.TestCheckResourceAttr("data.citrixadc_cacheobject.tf_cacheobject", "httpmethod", "GET"),
					//   resource.TestCheckResourceAttrSet("data.citrixadc_cacheobject.tf_cacheobject", "locator"),
				),
			},
		},
	})
}
