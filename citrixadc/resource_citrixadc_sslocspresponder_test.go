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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccSslocspresponder_basic = `
	resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = "tf_sslocspresponder"
		url = "http://www.citrix.com"
		batchingdelay = 5
		batchingdepth = 2
		cache = "ENABLED"
		cachetimeout = 1
		httpmethod = "GET"
		insertclientcert = "YES"
		ocspurlresolvetimeout = 100
		producedattimeskew = 300
		resptimeout = 100
		trustresponder = false
		usenonce = "NO"
	}
`

const testAccSslocspresponder_basic_update1 = `
	resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = "tf_sslocspresponder"
		url = "http://www.google.com"
		batchingdelay = 6
		batchingdepth = 3
		cache = "DISABLED"
		httpmethod = "POST"
		insertclientcert = "NO"
		ocspurlresolvetimeout = 101
		producedattimeskew = 301
		resptimeout = 101
		trustresponder = true
		usenonce = "YES"
	}
`

const testAccSslocspresponder_basic_update2 = `
	resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = "tf_sslocspresponder"
		url = "http://www.google.com"
		batchingdelay = 6
		batchingdepth = 3
		cache = "DISABLED"
		httpmethod = "POST"
		insertclientcert = "NO"
		ocspurlresolvetimeout = 101
		producedattimeskew = 301
		respondercert = "ns-server-certificate"
		resptimeout = 101
		signingcert = "ns-server-certificate"
		trustresponder = false
		usenonce = "YES"
	}
`

func TestAccSslocspresponder_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslocspresponderDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslocspresponder_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslocspresponderExist("citrixadc_sslocspresponder.tf_sslocspresponder", nil),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "url", "http://www.citrix.com"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "batchingdelay", "5"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "batchingdepth", "2"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "cache", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "cachetimeout", "1"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "httpmethod", "GET"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "insertclientcert", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "ocspurlresolvetimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "producedattimeskew", "300"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "resptimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "trustresponder", "false"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "usenonce", "NO"),
				),
			},
			resource.TestStep{
				Config: testAccSslocspresponder_basic_update1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslocspresponderExist("citrixadc_sslocspresponder.tf_sslocspresponder", nil),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "url", "http://www.google.com"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "batchingdelay", "6"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "batchingdepth", "3"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "cache", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "cachetimeout", "0"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "httpmethod", "POST"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "insertclientcert", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "ocspurlresolvetimeout", "101"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "producedattimeskew", "301"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "resptimeout", "101"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "trustresponder", "true"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "usenonce", "YES"),
				),
			},
			resource.TestStep{
				Config: testAccSslocspresponder_basic_update2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslocspresponderExist("citrixadc_sslocspresponder.tf_sslocspresponder", nil),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "respondercert", "ns-server-certificate"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "signingcert", "ns-server-certificate"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "trustresponder", "false"),
				),
			},
		},
	})
}

func testAccCheckSslocspresponderExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslocspresponder name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Sslocspresponder.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslocspresponder %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslocspresponderDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslocspresponder" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Sslocspresponder.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslocspresponder %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
