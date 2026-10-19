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

const testAccSslcertkeyUpdate_basic = `

	resource "citrixadc_sslcertkey" "tf_cert" {
		certkey = "sample_ssl_cert"
		cert    = "/nsconfig/ssl/servercert1.cert"
		key     = "/nsconfig/ssl/servercert1.key"
	}

	resource "citrixadc_sslcertkey_update" "tf_cert" {
		certkey = citrixadc_sslcertkey.tf_cert.certkey
		cert    = "/nsconfig/ssl/servercert1.cert"
		key     = "/nsconfig/ssl/servercert1.key"
	}
`

func TestAccSslcertkeyUpdate_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertkeyUpdate_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyUpdateExist("citrixadc_sslcertkey_update.tf_cert", nil),
				),
			},
		},
	})
}

const testAccSslcertkeyUpdate_passplain_wo = `

	variable "sslcertkey_update_passplain_wo" {
		type      = string
		sensitive = true
	}

	# Both the base certkey and the update target the SAME encrypted cert/key so the two
	# resources agree on the appliance's final state (no cross-resource cert drift). The base
	# loads servercert2 on create; the update resource re-asserts it via the write-only
	# pass phrase, exercising sslcertkey_update's passplain_wo path (GH #1480).
	resource "citrixadc_sslcertkey" "tf_upd_wo" {
		certkey              = "tf_upd_wo_certkey"
		cert                 = "/nsconfig/ssl/servercert2.cert"
		key                  = "/nsconfig/ssl/servercert2.key"
		passplain_wo         = var.sslcertkey_update_passplain_wo
		passplain_wo_version = 1
	}

	resource "citrixadc_sslcertkey_update" "tf_upd_wo" {
		certkey              = citrixadc_sslcertkey.tf_upd_wo.certkey
		cert                 = "/nsconfig/ssl/servercert2.cert"
		key                  = "/nsconfig/ssl/servercert2.key"
		passplain_wo         = var.sslcertkey_update_passplain_wo
		passplain_wo_version = 1
	}
`

// TestAccSslcertkeyUpdate_passplain_wo is the GH #1480 feature test: drive the certkey with
// an ENCRYPTED cert/key via the write-only passplain_wo. servercert2.key is an encrypted PEM
// key whose pass phrase is 123456, so the update action must re-read it WITH the pass phrase;
// if the provider dropped the write-only value (e.g. read it from the plan instead of config,
// the GH #1479 trap), NITRO would reject the still-encrypted key with errorcode 1537.
func TestAccSslcertkeyUpdate_passplain_wo(t *testing.T) {
	t.Setenv("TF_VAR_sslcertkey_update_passplain_wo", "123456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertkeyUpdate_passplain_wo,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyUpdateExist("citrixadc_sslcertkey_update.tf_upd_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcertkey_update.tf_upd_wo", "certkey", "tf_upd_wo_certkey"),
					resource.TestCheckResourceAttr("citrixadc_sslcertkey_update.tf_upd_wo", "passplain_wo_version", "1"),
				),
			},
		},
	})
}

func testAccCheckSslcertkeyUpdateExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No NsConfigClear is set")
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
