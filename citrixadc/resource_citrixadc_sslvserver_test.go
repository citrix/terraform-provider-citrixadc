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

const testAccSslvserver_basic = `
	resource "citrixadc_sslvserver" "tf_sslvserver" {
		cipherredirect = "ENABLED"
		cipherurl = "http://www.citrix.com"
		cleartextport = "80"
		clientauth = "ENABLED"
		clientcert = "Optional"
		hsts = "ENABLED"
		includesubdomains = "YES"
		maxage = "100"
		ocspstapling = "ENABLED"
		preload = "YES"
		sendclosenotify = "YES"
		sessreuse = "ENABLED"
		sesstimeout = "180"
		snienable = "ENABLED"
		sslredirect = "ENABLED"
		strictsigdigestcheck = "ENABLED"
		tls1 = "ENABLED"
		tls11 = "ENABLED"
		tls12 = "ENABLED"
		tls13 = "ENABLED"
		tls13sessionticketsperauthcontext = "7"
		zerorttearlydata = "ENABLED"
		vservername = citrixadc_lbvserver.tf_lbvserver.name
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_vserver"
		servicetype = "SSL"
	}
`

func TestAccSslvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslvserverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslvserver_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserverExist("citrixadc_sslvserver.tf_sslvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "cipherredirect", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "cipherurl", "http://www.citrix.com"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "cleartextport", "80"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "clientauth", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "clientcert", "Optional"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "hsts", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "includesubdomains", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "maxage", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "ocspstapling", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "preload", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "sendclosenotify", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "sessreuse", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "sesstimeout", "180"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "snienable", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "sslredirect", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "strictsigdigestcheck", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "tls1", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "tls11", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "tls12", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "tls13", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "tls13sessionticketsperauthcontext", "7"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "zerorttearlydata", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "vservername", "tf_vserver"),
				),
			},
		},
	})
}

func testAccCheckSslvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslvserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Sslvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslvserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslvserverDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Sslvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslvserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
