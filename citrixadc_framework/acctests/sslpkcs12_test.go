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

// sslpkcs12 is an ACTION-ONLY resource: NITRO exposes only ?action=convert with
// NO add/get/delete endpoints. The convert operation is DISRUPTIVE and
// NON-IDEMPOTENT - it requires the source certificate/key files to exist on the
// appliance filesystem and writes the converted output file under /nsconfig/ssl/.
//
// Because there is no GET endpoint:
//   - the Exist check is STATE-ONLY (verifies a non-empty Terraform ID; no
//     FindResource call is possible),
//   - there is NO CheckDestroy (no DELETE endpoint; the output file persists),
//   - there is NO datasource test (the datasource was removed - no GET).
//
// The test performs a single apply. Before running, supply real input files that
// exist on the testbed appliance (see TODO_PLACEHOLDER below) and the secret
// pass phrases via TF_VAR_* (set with t.Setenv). secrets are NEVER asserted.
//
// This example uses the EXPORT direction (PEM -> PKCS#12): certfile + keyfile ->
// pkcs12file (outfile is required by the schema regardless of direction).
const testAccSslpkcs12_basic = `

	variable "sslpkcs12_password" {
	  type      = string
	  sensitive = true
	}

	variable "sslpkcs12_pempassphrase" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslpkcs12" "tf_sslpkcs12" {
		export = true

		// TODO_PLACEHOLDER: existing PEM certificate file under /nsconfig/ssl/ on
		// the testbed appliance, e.g. "tf_cert.pem".
		certfile = "TODO_PLACEHOLDER"
		// TODO_PLACEHOLDER: existing PEM private-key file under /nsconfig/ssl/ on
		// the testbed appliance, e.g. "tf_key.pem".
		keyfile = "TODO_PLACEHOLDER"

		// Output PKCS#12 file written under /nsconfig/ssl/ by the convert action.
		pkcs12file = "tf_sslpkcs12.p12"
		// outfile is Required by the schema for both directions.
		outfile = "tf_sslpkcs12_out.pem"

		// password protects the PKCS#12 output; pempassphrase is the pass phrase
		// for the encrypted PEM key. Both are mandatory (ValidateConfig).
		password      = var.sslpkcs12_password
		pempassphrase = var.sslpkcs12_pempassphrase
	}
`

func TestAccSslpkcs12_basic(t *testing.T) {
	// TODO_PLACEHOLDER: replace these with the real pass phrases for the input
	// files on the testbed. Secrets are passed via env and never asserted.
	t.Setenv("TF_VAR_sslpkcs12_password", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslpkcs12_pempassphrase", "TODO_PLACEHOLDER")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: action-only resource has no DELETE endpoint.
		Steps: []resource.TestStep{
			{
				Config: testAccSslpkcs12_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpkcs12Exist("citrixadc_sslpkcs12.tf_sslpkcs12", nil),
					// Only assert HCL-set / echoed non-secret attributes. Secrets
					// (password, pempassphrase) are never asserted.
					resource.TestCheckResourceAttr("citrixadc_sslpkcs12.tf_sslpkcs12", "export", "true"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs12.tf_sslpkcs12", "pkcs12file", "tf_sslpkcs12.p12"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs12.tf_sslpkcs12", "outfile", "tf_sslpkcs12_out.pem"),
				),
			},
		},
	})
}

// testAccCheckSslpkcs12Exist is STATE-ONLY: sslpkcs12 has no NITRO GET endpoint,
// so there is no FindResource call. It only verifies the Terraform ID is set
// (the convert action ran and produced state).
func testAccCheckSslpkcs12Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslpkcs12 ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No FindResource call: sslpkcs12 is action-only (?action=convert) with no
		// GET endpoint to read back the converted result.
		return nil
	}
}
