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

// sslhpkekey is a named resource with add/delete/get CRUD and NO update -
// every schema attribute is RequiresReplace, so an attribute change forces
// recreation. There is therefore no step2 "update" path; the basic test is a
// single create+verify.
//
// The `file` attribute names an HPKE key file that must already exist under
// /nsconfig/ssl/ on the appliance for `add ssl hpkekey` to succeed. The test's
// PreCheck (doSslhpkekeyPreChecks) stages testdata/hpke_key.der there.
// IMPORTANT: the appliance accepts only the X25519 PKCS#8 *DER* form; a
// PEM-armored key is rejected with "Invalid HPKEKey" (verified on the
// appliance), and there is no `create ssl hpkekey` generator, so the key must be
// pre-staged as DER.
const testAccSslhpkekey_basic = `
resource "citrixadc_sslhpkekey" "tf_sslhpkekey" {
  hpkekeyname = "tf_sslhpkekey"
  dhkem       = "X_25519"
  file = "hpke_key.der"
}
`

func TestAccSslhpkekey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslhpkekeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslhpkekeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslhpkekey_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslhpkekeyExist("citrixadc_sslhpkekey.tf_sslhpkekey", nil),
					resource.TestCheckResourceAttr("citrixadc_sslhpkekey.tf_sslhpkekey", "hpkekeyname", "tf_sslhpkekey"),
					resource.TestCheckResourceAttr("citrixadc_sslhpkekey.tf_sslhpkekey", "dhkem", "X_25519"),
				),
			},
		},
	})
}

func TestAccSslhpkekey_import(t *testing.T) {
	const resAddr = "citrixadc_sslhpkekey.tf_sslhpkekey"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslhpkekeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslhpkekeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslhpkekey_basic,
			},
			{
				Config:                  testAccSslhpkekey_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSslhpkekeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslhpkekey name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Sslhpkekey.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslhpkekey %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslhpkekeyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslhpkekey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslhpkekey.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslhpkekey %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSslhpkekeyDataSource_basic = `
resource "citrixadc_sslhpkekey" "tf_sslhpkekey" {
  hpkekeyname = "tf_sslhpkekey"
  dhkem       = "X_25519"
  // Staged under /nsconfig/ssl/ by doSslhpkekeyPreChecks (X25519 PKCS#8 DER).
  file = "hpke_key.der"
}

data "citrixadc_sslhpkekey" "tf_sslhpkekey" {
  hpkekeyname = citrixadc_sslhpkekey.tf_sslhpkekey.hpkekeyname
  depends_on  = [citrixadc_sslhpkekey.tf_sslhpkekey]
}
`

func TestAccSslhpkekeyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslhpkekeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslhpkekeyDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslhpkekey.tf_sslhpkekey", "hpkekeyname", "tf_sslhpkekey"),
					resource.TestCheckResourceAttr("data.citrixadc_sslhpkekey.tf_sslhpkekey", "dhkem", "X_25519"),
					resource.TestCheckResourceAttrSet("data.citrixadc_sslhpkekey.tf_sslhpkekey", "id"),
				),
			},
		},
	})
}
