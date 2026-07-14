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
// NO add/get/delete endpoints. The convert operation is NON-IDEMPOTENT - it
// requires the source certificate/key files to exist on the appliance filesystem
// and writes the converted output file(s) under /nsconfig/ssl/.
//
// Because there is no GET endpoint:
//   - the Exist check is STATE-ONLY (verifies a non-empty Terraform ID),
//   - there is NO CheckDestroy (no DELETE endpoint; the output files persist),
//   - there is NO datasource test (no GET).
//
// All tests below use the EXPORT direction (PEM cert + key -> PKCS#12). The input
// cert/key files are staged on the appliance by doSslcertkeyPreChecks:
//   servercert1.{cert,key} - UNENCRYPTED key   (no pempassphrase needed)
//   servercert2.{cert,key} - key encrypted with pass phrase "123456"
//   servercert3.{cert,key} - key encrypted with pass phrase "1234567"
//
// `password` protects the PKCS#12 output; `pempassphrase` unlocks an encrypted
// input key. Secrets are passed via TF_VAR_* and are never asserted.
//
// Each test is t.Skip-gated: the convert action writes pkcs12file + outfile that
// have no DELETE endpoint, so a re-run collides unless unique output names are
// used or the files are cleaned up on the appliance (mirrors sslpkcs8_test.go).

// -----------------------------------------------------------------------------
// basic: export an UNENCRYPTED key -> no pempassphrase supplied. Exercises the
// "without pempassphrase" path.
const testAccSslpkcs12_basic = `

	variable "sslpkcs12_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslpkcs12" "tf_sslpkcs12" {
		export = true
		certfile = "/nsconfig/ssl/servercert1.cert"
		keyfile  = "/nsconfig/ssl/servercert1.key"
		outfile    = "tf_sslpkcs12_out.pem"
		password = var.sslpkcs12_password
	}
`

func TestAccSslpkcs12_basic(t *testing.T) {
	t.Skip("Requires clean up of pkcs12file/outfile from ADC file system or unique pkcs12file/outfile has to be provided.")
	t.Setenv("TF_VAR_sslpkcs12_password", "p12password")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: action-only resource has no DELETE endpoint.
		Steps: []resource.TestStep{
			{
				Config: testAccSslpkcs12_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpkcs12Exist("citrixadc_sslpkcs12.tf_sslpkcs12", nil),
					// Only assert HCL-set / echoed non-secret attributes.
					resource.TestCheckResourceAttr("citrixadc_sslpkcs12.tf_sslpkcs12", "export", "true"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs12.tf_sslpkcs12", "keyfile", "/nsconfig/ssl/servercert1.key"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs12.tf_sslpkcs12", "outfile", "tf_sslpkcs12_out.pem"),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Backward-compatibility test: export an ENCRYPTED key using the LEGACY plaintext
// secret attributes `password` + `pempassphrase`. These must keep working after
// the write-only variants were introduced. Mirrors TestAccSslcertkey_passplain /
// TestAccSslpkcs8_password. Uses servercert3.key (pass phrase "1234567").
const testAccSslpkcs12_password_basic = `

	variable "sslpkcs12_password" {
	  type      = string
	  sensitive = true
	}

	variable "sslpkcs12_pempassphrase" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslpkcs12" "tf_sslpkcs12_pw" {
		export = true
		certfile = "/nsconfig/ssl/servercert3.cert"
		keyfile  = "/nsconfig/ssl/servercert3.key"
		outfile    = "tf_sslpkcs12_pw_out.pem"
		password      = var.sslpkcs12_password
		pempassphrase = var.sslpkcs12_pempassphrase
	}
`

func TestAccSslpkcs12_password(t *testing.T) {
	t.Skip("Requires clean up of pkcs12file/outfile from ADC file system or unique pkcs12file/outfile has to be provided.")
	t.Setenv("TF_VAR_sslpkcs12_password", "p12password")
	// Pass phrase of the encrypted input key servercert3.key.
	t.Setenv("TF_VAR_sslpkcs12_pempassphrase", "1234567")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: action-only resource has no DELETE endpoint.
		Steps: []resource.TestStep{
			{
				Config: testAccSslpkcs12_password_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpkcs12Exist("citrixadc_sslpkcs12.tf_sslpkcs12_pw", nil),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs12.tf_sslpkcs12_pw", "keyfile", "/nsconfig/ssl/servercert3.key"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs12.tf_sslpkcs12_pw", "outfile", "tf_sslpkcs12_pw_out.pem"),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Write-only (_wo) test: export ENCRYPTED keys via `password_wo` /
// `pempassphrase_wo` and bump the *_wo_version across steps to signal secret
// rotation. Mirrors the write-only flow in TestAccSslcertkey_basic /
// TestAccSslpkcs8_wo. Every sslpkcs12 attribute is RequiresReplace (action-only
// convert), so each step re-runs the convert. Step 1 uses servercert2.key (pass
// phrase "123456"); step 2 rotates to servercert3.key (pass phrase "1234567").
// The write-only secrets are never stored in state or asserted.
const testAccSslpkcs12_wo_step1 = `

	variable "sslpkcs12_password_wo" {
	  type      = string
	  sensitive = true
	}

	variable "sslpkcs12_pempassphrase_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslpkcs12" "tf_sslpkcs12_wo" {
		export = true
		certfile = "/nsconfig/ssl/servercert2.cert"
		keyfile  = "/nsconfig/ssl/servercert2.key"
		outfile    = "tf_sslpkcs12_wo_out.pem"
		password_wo              = var.sslpkcs12_password_wo
		password_wo_version      = 1
		pempassphrase_wo         = var.sslpkcs12_pempassphrase_wo
		pempassphrase_wo_version = 1
	}
`

func TestAccSslpkcs12_wo(t *testing.T) {
	t.Skip("Requires clean up of pkcs12file/outfile from ADC file system or unique pkcs12file/outfile has to be provided.")
	t.Setenv("TF_VAR_sslpkcs12_password_wo", "p12password")
	// Pass phrase of the encrypted input key servercert2.key (staged by PreCheck).
	t.Setenv("TF_VAR_sslpkcs12_pempassphrase_wo", "123456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: action-only resource has no DELETE endpoint.
		Steps: []resource.TestStep{
			{
				Config: testAccSslpkcs12_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpkcs12Exist("citrixadc_sslpkcs12.tf_sslpkcs12_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs12.tf_sslpkcs12_wo", "keyfile", "/nsconfig/ssl/servercert2.key"),
					resource.TestCheckResourceAttr("citrixadc_sslpkcs12.tf_sslpkcs12_wo", "pempassphrase_wo_version", "1"),
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
