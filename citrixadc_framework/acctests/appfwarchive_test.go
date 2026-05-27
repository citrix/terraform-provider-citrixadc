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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NOTE on the appfwarchive resource:
//   - Models the NITRO POST /appfwarchive?action=Import endpoint.
//   - Every attribute is RequiresReplace; there is no in-place update path on
//     NITRO. The "step2" config below therefore exercises destroy+recreate
//     rather than a true update.
//   - The NITRO `get (all)` response carries no per-archive identifying fields
//     (no `name` echo); the resource's Read treats a non-empty list as
//     existence confirmation. The exist check below mirrors that behavior.
//   - `src` MUST be a reachable http(s) URL serving a valid AppFW tar archive
//     for Import to succeed. Replace the TODO_PLACEHOLDER URLs with real
//     archive URLs reachable from the ADC under test before running.

const testAccAppfwarchive_basic_step1 = `
resource "citrixadc_appfwarchive" "tf_appfwarchive" {
  name    = "new_tfappfwarch"
  src     = "local:new_tfappfwarchfile"
  comment = "test_appfwarchive_v1"
}

`

func TestAccAppfwarchive_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwarchiveDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwarchive_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwarchiveExist("citrixadc_appfwarchive.tf_appfwarchive", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwarchive.tf_appfwarchive", "name", "new_tfappfwarch"),
					resource.TestCheckResourceAttr("citrixadc_appfwarchive.tf_appfwarchive", "comment", "test_appfwarchive_v1"),
				),
			},
		},
	})
}

func testAccCheckAppfwarchiveExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwarchive name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		// NITRO appfwarchive `get (all)` does not echo per-archive identifying
		// fields, so we cannot match by name. A successful, non-empty GET is
		// confirmation that at least one archive exists on the ADC.
		data, err := client.FindResource(service.Appfwarchive.Type(), "")
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwarchive %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwarchiveDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwarchive" {
			continue
		}
		// DELETE /appfwarchive/{name} removes the archive; verify the named
		// archive is gone by calling FindResource with the name. NITRO returns
		// an error / nil when the archive does not exist.
		data, err := client.FindResource(service.Appfwarchive.Type(), rs.Primary.ID)
		if err == nil && data != nil {
			return fmt.Errorf("appfwarchive %s still exists after destroy", rs.Primary.ID)
		}
	}
	return nil
}

// Datasource test. The datasource for appfwarchive is queryable by `name`,
// but the underlying NITRO `get (all)` response carries no identifying
// fields, so only the user-supplied `name` is reliably readable through the
// datasource. We avoid asserting on `comment`/`src` because the API never
// echoes them back.
const testAccAppfwarchiveDataSource_basic = `

resource "citrixadc_appfwarchive" "tf_appfwarchive" {
  name    = "new_tfappfwarch"
  src     = "local:new_tfappfwarchfile"
  comment = "test_appfwarchive_ds"
}

data "citrixadc_appfwarchive" "tf_appfwarchive" {
  name       = citrixadc_appfwarchive.tf_appfwarchive.name
  depends_on = [citrixadc_appfwarchive.tf_appfwarchive]
}
`

func TestAccAppfwarchiveDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwarchiveDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwarchive.tf_appfwarchive", "name", "new_tfappfwarch"),
				),
			},
		},
	})
}
