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

	resource "citrixadc_sslpkcs8" "tf_sslpkcs8" {
		keyfile = "/nsconfig/ssl/servercert1.key"
		pkcs8file = "tf_sslpkcs8.pk8"
		keyform = "PEM"
	}
`

func TestAccSslpkcs8_basic(t *testing.T) {
	// TODO_PLACEHOLDER: pass phrase for the encrypted PEM key, if any. Leave a
	// real value if the input key is encrypted; otherwise an empty string is fine.
	t.Skip("Requires clean up of pkcs8file from ADC file system or unique pkcs8file has to be provided.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
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

// -----------------------------------------------------------------------------
// Backward-compatibility test: the legacy plaintext `password` attribute must
// keep working after the write-only `password_wo` variant was introduced.
// Mirrors TestAccSslcertkey_passplain in sslcertkey_test.go.
//
// Uses servercert3.key, an encrypted PEM key staged on the appliance by
// doSslcertkeyPreChecks (pass phrase "1234567", the same value the sslcertkey
// tests use). The convert action decrypts the input key with `password` and
// writes the PKCS#8 output. The secret is sensitive and is never asserted; only
// the echoed non-secret attributes are checked.
const testAccSslpkcs8_password_basic = `

	variable "sslpkcs8_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslpkcs8" "tf_sslpkcs8_pw" {
		keyfile   = "/nsconfig/ssl/servercert3.key"
		pkcs8file = "tf_sslpkcs8_pw.pk8"
		keyform   = "PEM"
		password = var.sslpkcs8_password
	}
`

func TestAccSslpkcs8_password(t *testing.T) {
	t.Skip("Requires clean up of pkcs8file from ADC file system or unique pkcs8file has to be provided.")
	// Pass phrase of the encrypted input key servercert3.key.
	t.Setenv("TF_VAR_sslpkcs8_password", "1234567")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: action-only resource has no DELETE endpoint.
		Steps: []resource.TestStep{
			{
				Config: testAccSslpkcs8_password_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpkcs8Exist("citrixadc_sslpkcs8.tf_sslpkcs8_pw", nil),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs8.tf_sslpkcs8_pw", "keyfile", "/nsconfig/ssl/servercert3.key"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs8.tf_sslpkcs8_pw", "pkcs8file", "tf_sslpkcs8_pw.pk8"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs8.tf_sslpkcs8_pw", "keyform", "PEM"),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Write-only (_wo) test: exercises `password_wo` + `password_wo_version`, plus a
// version bump across steps to signal a secret rotation. Mirrors the write-only
// flow in TestAccSslcertkey_basic (passplain_wo / passplain_wo_version).
//
// Every sslpkcs8 attribute is RequiresReplace (action-only convert), so each
// step re-runs the convert. Step 1 converts the encrypted servercert2.key (pass
// phrase "123456"); step 2 rotates to servercert3.key (pass phrase "1234567")
// and bumps password_wo_version. The write-only secret is never stored in state
// or asserted.
const testAccSslpkcs8_wo_step1 = `

	variable "sslpkcs8_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslpkcs8" "tf_sslpkcs8_wo" {
		keyfile             = "/nsconfig/ssl/servercert2.key"
		pkcs8file           = "tf_sslpkcs8_wo.pk8"
		keyform             = "PEM"
		password_wo         = var.sslpkcs8_password_wo
		password_wo_version = 1
	}
`

const testAccSslpkcs8_wo_step2 = `

	variable "sslpkcs8_password_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslpkcs8" "tf_sslpkcs8_wo" {
		keyfile             = "/nsconfig/ssl/servercert3.key"
		pkcs8file           = "tf_sslpkcs8_wo_2.pk8"
		keyform             = "PEM"
		password_wo         = var.sslpkcs8_password_wo_2
		password_wo_version = 2
	}
`

func TestAccSslpkcs8_wo(t *testing.T) {
	t.Skip("Requires clean up of pkcs8file from ADC file system or unique pkcs8file has to be provided.")
	// Pass phrases of the encrypted input keys (staged by doSslcertkeyPreChecks).
	t.Setenv("TF_VAR_sslpkcs8_password_wo", "123456")
	t.Setenv("TF_VAR_sslpkcs8_password_wo_2", "1234567")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: action-only resource has no DELETE endpoint.
		Steps: []resource.TestStep{
			{
				Config: testAccSslpkcs8_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpkcs8Exist("citrixadc_sslpkcs8.tf_sslpkcs8_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs8.tf_sslpkcs8_wo", "keyfile", "/nsconfig/ssl/servercert2.key"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs8.tf_sslpkcs8_wo", "pkcs8file", "tf_sslpkcs8_wo.pk8"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs8.tf_sslpkcs8_wo", "password_wo_version", "1"),
				),
			},
			{
				Config: testAccSslpkcs8_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpkcs8Exist("citrixadc_sslpkcs8.tf_sslpkcs8_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs8.tf_sslpkcs8_wo", "keyfile", "/nsconfig/ssl/servercert3.key"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs8.tf_sslpkcs8_wo", "pkcs8file", "tf_sslpkcs8_wo_2.pk8"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs8.tf_sslpkcs8_wo", "password_wo_version", "2"),
				),
			},
		},
	})
}
