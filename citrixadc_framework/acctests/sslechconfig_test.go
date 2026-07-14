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

// sslechconfig is a standard add/delete named resource (no update; every
// attribute is RequiresReplace). It references an HPKE key by hpkekeyname, so
// the test first creates a citrixadc_sslhpkekey (participating entity) and
// wires hpkekeyname to it via reference + depends_on.
//
// The HPKE key requires a `file` (X25519 PKCS#8 DER) that already exists under
// /nsconfig/ssl on the appliance and dhkem = X_25519. The test PreCheck
// (doSslhpkekeyPreChecks) stages testdata/hpke_key.der there.
//
// echcipher must be one of the NITRO "Possible values" for sslechconfig.echcipher
// (AES128/256-GCM-HKDFSHA256/384/512); this test uses AES128-GCM-HKDFSHA256.

const testAccSslechconfig_basic_step1 = `
resource "citrixadc_sslhpkekey" "tf_sslhpkekey" {
  hpkekeyname = "tf_sslhpkekey"
  dhkem       = "X_25519"
  file        = "hpke_key.der"
}

resource "citrixadc_sslechconfig" "tf_sslechconfig" {
  echconfigname = "tf_sslechconfig"
  echcipher     = "AES128-GCM-HKDFSHA256"
  hpkekeyname   = citrixadc_sslhpkekey.tf_sslhpkekey.hpkekeyname
  echpublicname = "public.example.com"
  echconfigid   = 1
  version       = 65037

  depends_on = [citrixadc_sslhpkekey.tf_sslhpkekey]
}

`

func TestAccSslechconfig_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslhpkekeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslechconfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslechconfig_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslechconfigExist("citrixadc_sslechconfig.tf_sslechconfig", nil),
					resource.TestCheckResourceAttr("citrixadc_sslechconfig.tf_sslechconfig", "echconfigname", "tf_sslechconfig"),
					resource.TestCheckResourceAttr("citrixadc_sslechconfig.tf_sslechconfig", "echcipher", "AES128-GCM-HKDFSHA256"),
					resource.TestCheckResourceAttr("citrixadc_sslechconfig.tf_sslechconfig", "echpublicname", "public.example.com"),
					resource.TestCheckResourceAttr("citrixadc_sslechconfig.tf_sslechconfig", "echconfigid", "1"),
				),
			},
		},
	})
}

func testAccCheckSslechconfigExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslechconfig name is set")
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
		data, err := client.FindResource(service.Sslechconfig.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslechconfig %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslechconfigDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslechconfig" {
			continue
		}

		_, err := client.FindResource(service.Sslechconfig.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslechconfig %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccSslechconfigDataSource_basic = `

resource "citrixadc_sslhpkekey" "tf_sslhpkekey" {
  hpkekeyname = "tf_sslhpkekey"
  dhkem       = "X_25519"
  file        = "hpke_key.der"
}

resource "citrixadc_sslechconfig" "tf_sslechconfig" {
  echconfigname = "tf_sslechconfig"
  echcipher     = "AES128-GCM-HKDFSHA256"
  hpkekeyname   = citrixadc_sslhpkekey.tf_sslhpkekey.hpkekeyname
  echpublicname = "public.example.com"
  echconfigid   = 1
  version       = 65037

  depends_on = [citrixadc_sslhpkekey.tf_sslhpkekey]
}

data "citrixadc_sslechconfig" "tf_sslechconfig" {
  echconfigname = citrixadc_sslechconfig.tf_sslechconfig.echconfigname
  depends_on    = [citrixadc_sslechconfig.tf_sslechconfig]
}
`

func TestAccSslechconfigDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslhpkekeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslechconfigDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslechconfig.tf_sslechconfig", "echconfigname", "tf_sslechconfig"),
					resource.TestCheckResourceAttr("data.citrixadc_sslechconfig.tf_sslechconfig", "echcipher", "AES128-GCM-HKDFSHA256"),
					resource.TestCheckResourceAttr("data.citrixadc_sslechconfig.tf_sslechconfig", "echpublicname", "public.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_sslechconfig.tf_sslechconfig", "echconfigid", "1"),
				),
			},
		},
	})
}
