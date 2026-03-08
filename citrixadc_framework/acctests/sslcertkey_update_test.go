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
package acctests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
