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

// sslpkcs8 is an ACTION-ONLY resource: NITRO exposes only ?action=convert with
// NO add/get/delete endpoints. The convert operation is NON-IDEMPOTENT - it
// requires the source key file to exist on the appliance filesystem and writes
// the converted PKCS#8 output file under /nsconfig/ssl/.
//
// Because there is no GET endpoint:
//   - the Exist check is STATE-ONLY (verifies a non-empty Terraform ID),
//   - there is NO CheckDestroy (no DELETE endpoint; the output file persists),
//   - there is NO datasource test (no GET).
//
// The test performs a single apply. Before running, supply a real input key file
// that exists on the testbed appliance (see TODO_PLACEHOLDER). `password` is the
// (optional) pass phrase for an encrypted PEM key, passed via TF_VAR_* and never
// asserted.
const testAccSslpkcs8_basic = `

	variable "sslpkcs8_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslpkcs8" "tf_sslpkcs8" {
		// TODO_PLACEHOLDER: existing PEM/DER private-key file under /nsconfig/ssl/
		// on the testbed appliance, e.g. "tf_key.pem".
		keyfile = "TODO_PLACEHOLDER"

		// Output PKCS#8 file written under /nsconfig/ssl/ by the convert action.
		pkcs8file = "tf_sslpkcs8.pk8"

		// Format the input key is stored in. Possible values: DER, PEM (default PEM).
		keyform = "PEM"

		// password applies only when the PEM key is encrypted. Optional.
		password = var.sslpkcs8_password
	}
`

func TestAccSslpkcs8_basic(t *testing.T) {
	// TODO_PLACEHOLDER: pass phrase for the encrypted PEM key, if any. Leave a
	// real value if the input key is encrypted; otherwise an empty string is fine.
	t.Setenv("TF_VAR_sslpkcs8_password", "TODO_PLACEHOLDER")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: action-only resource has no DELETE endpoint.
		Steps: []resource.TestStep{
			{
				Config: testAccSslpkcs8_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpkcs8Exist("citrixadc_sslpkcs8.tf_sslpkcs8", nil),
					// Only assert HCL-set / echoed non-secret attributes. The secret
					// (password) is never asserted.
					resource.TestCheckResourceAttr("citrixadc_sslpkcs8.tf_sslpkcs8", "pkcs8file", "tf_sslpkcs8.pk8"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs8.tf_sslpkcs8", "keyform", "PEM"),
				),
			},
		},
	})
}

// testAccCheckSslpkcs8Exist is STATE-ONLY: sslpkcs8 has no NITRO GET endpoint,
// so there is no FindResource call. It only verifies the Terraform ID is set.
func testAccCheckSslpkcs8Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslpkcs8 ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No FindResource call: sslpkcs8 is action-only (?action=convert) with no
		// GET endpoint to read back the converted result.
		return nil
	}
}
