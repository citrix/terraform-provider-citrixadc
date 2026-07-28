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

// !!! DANGER -- DO NOT RUN IN AUTOMATION !!!
//
// nslaslicense_apply is an APPLY-ONLY resource whose Create performs the NITRO
// `apply` action (POST ?action=apply). This action is DISRUPTIVE and
// NON-IDEMPOTENT: it ACTIVATES a LAS license and ALTERS the licensed capacity
// of the appliance. There is no NITRO get/add/delete/update endpoint.
//
// Running this test requires a real `.lic` file already STAGED on the ADC at
// `filelocation` (e.g. /nsconfig/license). Applying it changes the appliance's
// licensed capacity. For that reason the test is t.Skip-gated. Only un-skip and
// run it manually, against a disposable appliance, with real values
// substituted for the TODO_PLACEHOLDER fields.

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Required: filename, filelocation. The .lic file named `filename` must already
// exist at `filelocation` on the ADC before apply.
const testAccNslaslicenseApply_basic_step1 = `
resource "citrixadc_nslaslicense_apply" "tf_nslaslicense" {
  filename     = "offline_token_10.101.132.151_activation.blob.tgz"
  filelocation = "/nsconfig/license"
  fixedbandwidth = "true"
}

`

func TestAccNslaslicenseApply_basic(t *testing.T) {
	t.Skip("Requires valid license blob file staged at /nsconfig/license")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Apply-only resource: no CheckDestroy (NITRO exposes no delete endpoint).
		Steps: []resource.TestStep{
			{
				Config: testAccNslaslicenseApply_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslaslicenseApplyExist("citrixadc_nslaslicense_apply.tf_nslaslicense", nil),
					resource.TestCheckResourceAttrSet("citrixadc_nslaslicense_apply.tf_nslaslicense", "id"),
					resource.TestCheckResourceAttr("citrixadc_nslaslicense_apply.tf_nslaslicense", "filename", "offline_token_10.101.132.151_activation.blob.tgz"),
					resource.TestCheckResourceAttr("citrixadc_nslaslicense_apply.tf_nslaslicense", "filelocation", "/nsconfig/license"),
					resource.TestCheckResourceAttr("citrixadc_nslaslicense_apply.tf_nslaslicense", "fixedbandwidth", "true"),
				),
			},
		},
	})
}

// testAccCheckNslaslicenseApplyExist is a STATE-ONLY check. nslaslicense_apply is
// apply-only with no NITRO GET endpoint, so there is nothing to read back; we
// only confirm the resource has a non-empty synthetic ID.
func testAccCheckNslaslicenseApplyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nslaslicense_apply ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO FindResource call: nslaslicense_apply exposes no GET endpoint.
		return nil
	}
}
