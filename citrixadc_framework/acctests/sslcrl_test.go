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

const testAccSslcrl_basic = `
	resource "citrixadc_sslcrl" "tf_sslcrl" {
		crlname = "tf_sslcrl"
		crlpath = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert  = "rootrsa_cert1"
	}
`

const testAccSslcrlDataSource_basic = `
	resource "citrixadc_sslcrl" "tf_sslcrl" {
		crlname = "tf_sslcrl"
		crlpath = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert  = "rootrsa_cert1"
	}

	data "citrixadc_sslcrl" "tf_sslcrl" {
		crlname = citrixadc_sslcrl.tf_sslcrl.crlname
	}
`

func TestAccSslcrl_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcrlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrl_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlExist("citrixadc_sslcrl.tf_sslcrl", nil),
				),
			},
		},
	})
}

func testAccCheckSslcrlExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcrl name is set")
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
		data, err := client.FindResource(service.Sslcrl.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslcrl %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcrlDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcrl" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslcrl.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslcrl %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslcrlDataSource_basic(t *testing.T) {
	t.Skipf("Find  a way to upload a CRL file to the ADC instance before running this test")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrlDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslcrl.tf_sslcrl", "crlname", "tf_sslcrl"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcrl.tf_sslcrl", "crlpath", "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcrl.tf_sslcrl", "cacert", "rootrsa_cert1"),
				),
			},
		},
	})
}
