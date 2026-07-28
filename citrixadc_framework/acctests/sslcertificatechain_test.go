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

// sslcertificatechain is create-only (AddResource); its NITRO Delete is a no-op
// (state-only removal) - the chain remains on the appliance after destroy. Therefore
// CheckDestroy below intentionally TOLERATES the chain still existing and only verifies
// the participating sslcertkey was destroyed.
//
// Participating entity: citrixadc_sslcertkey (config lifted from sslcertkey_test.go).
// The cert/key files referenced by the sslcertkey must ALREADY EXIST on the appliance
// under /nsconfig/ssl/. The filenames below are TODO_PLACEHOLDERs - replace them with
// real cert/key files present on your testbed.
const testAccSslcertificatechain_basic_step1 = `
resource "citrixadc_sslcertkey" "tf_chain_certkey" {
  certkey            = "tf_chain_certkey"
  cert               = "/nsconfig/ssl/servercert1.cert" // TODO_PLACEHOLDER: cert file must pre-exist on the appliance
  key                = "/nsconfig/ssl/servercert1.key"  // TODO_PLACEHOLDER: key file must pre-exist on the appliance
  notificationperiod = 40
  expirymonitor      = "ENABLED"
}

resource "citrixadc_sslcertificatechain" "tf_sslcertificatechain" {
  certkeyname = citrixadc_sslcertkey.tf_chain_certkey.certkey
  depends_on  = [citrixadc_sslcertkey.tf_chain_certkey]
}

`

func TestAccSslcertificatechain_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertificatechainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertificatechain_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertificatechainExist("citrixadc_sslcertificatechain.tf_sslcertificatechain", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcertificatechain.tf_sslcertificatechain", "certkeyname", "tf_chain_certkey"),
				),
			},
		},
	})
}

func TestAccSslcertificatechain_import(t *testing.T) {
	const resAddr = "citrixadc_sslcertificatechain.tf_sslcertificatechain"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertificatechainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertificatechain_basic_step1,
			},
			{
				Config:                  testAccSslcertificatechain_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSslcertificatechainExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcertificatechain name is set")
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
		data, err := client.FindResource(service.Sslcertificatechain.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslcertificatechain %s not found", n)
		}

		return nil
	}
}

// sslcertificatechain has no NITRO delete endpoint - the chain is removed from Terraform
// state only and remains on the appliance. We therefore only verify that the participating
// sslcertkey is destroyed; the chain remaining is expected and not treated as a failure.
func testAccCheckSslcertificatechainDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcertkey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslcertkey.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslcertkey %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccSslcertificatechainDataSource_basic = `

resource "citrixadc_sslcertkey" "tf_chain_certkey" {
  certkey            = "tf_chain_certkey"
  cert               = "/nsconfig/ssl/servercert1.cert" // TODO_PLACEHOLDER: cert file must pre-exist on the appliance
  key                = "/nsconfig/ssl/servercert1.key"  // TODO_PLACEHOLDER: key file must pre-exist on the appliance
  notificationperiod = 40
  expirymonitor      = "ENABLED"
}

resource "citrixadc_sslcertificatechain" "tf_sslcertificatechain" {
  certkeyname = citrixadc_sslcertkey.tf_chain_certkey.certkey
  depends_on  = [citrixadc_sslcertkey.tf_chain_certkey]
}

data "citrixadc_sslcertificatechain" "tf_sslcertificatechain" {
  certkeyname = citrixadc_sslcertificatechain.tf_sslcertificatechain.certkeyname
  depends_on  = [citrixadc_sslcertificatechain.tf_sslcertificatechain]
}
`

func TestAccSslcertificatechainDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertificatechainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertificatechainDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslcertificatechain.tf_sslcertificatechain", "certkeyname", "tf_chain_certkey"),
				),
			},
		},
	})
}
