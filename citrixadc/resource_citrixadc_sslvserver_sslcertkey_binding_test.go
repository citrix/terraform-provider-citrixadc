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
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

const testAccSslvserver_sslcertkey_binding_lb_step1 = `
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/var/tmp/certificate2.crt"
  key = "/var/tmp/key2.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
}

resource "citrixadc_sslcertkey" "tf_cacertkey" {
  certkey = "tf_cacertkey"
  cert = "/var/tmp/ca.crt"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.44"
  name        = "tf_lbvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
    vservername = citrixadc_lbvserver.tf_lbvserver.name
	certkeyname = citrixadc_sslcertkey.tf_cacertkey.certkey
    ca = true
}
`
const testAccSslvserver_sslcertkey_binding_lb_step2 = `
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/var/tmp/certificate2.crt"
  key = "/var/tmp/key2.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
}

resource "citrixadc_sslcertkey" "tf_cacertkey" {
  certkey = "tf_cacertkey"
  cert = "/var/tmp/ca.crt"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.44"
  name        = "tf_lbvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
    vservername = citrixadc_lbvserver.tf_lbvserver.name
	certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
}
`

const testAccSslvserver_sslcertkey_binding_lb_step3 = `
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/var/tmp/certificate2.crt"
  key = "/var/tmp/key2.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
}

resource "citrixadc_sslcertkey" "tf_cacertkey" {
  certkey = "tf_cacertkey"
  cert = "/var/tmp/ca.crt"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.44"
  name        = "tf_lbvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
    vservername = citrixadc_lbvserver.tf_lbvserver.name
	certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
	snicert = true
}
`

const testAccSslvserver_sslcertkey_binding_cs_step1 = `
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/var/tmp/certificate2.crt"
  key = "/var/tmp/key2.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
}

resource "citrixadc_sslcertkey" "tf_cacertkey" {
  certkey = "tf_cacertkey"
  cert = "/var/tmp/ca.crt"
}


resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.55"
  name        = "tf_csvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
    vservername = citrixadc_csvserver.tf_csvserver.name
	certkeyname = citrixadc_sslcertkey.tf_cacertkey.certkey
    ca = true
}
`

const testAccSslvserver_sslcertkey_binding_cs_step2 = `
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/var/tmp/certificate2.crt"
  key = "/var/tmp/key2.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
}

resource "citrixadc_sslcertkey" "tf_cacertkey" {
  certkey = "tf_cacertkey"
  cert = "/var/tmp/ca.crt"
}


resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.55"
  name        = "tf_csvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
    vservername = citrixadc_csvserver.tf_csvserver.name
	certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
}
`

const testAccSslvserver_sslcertkey_binding_cs_step3 = `
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/var/tmp/certificate2.crt"
  key = "/var/tmp/key2.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
}

resource "citrixadc_sslcertkey" "tf_cacertkey" {
  certkey = "tf_cacertkey"
  cert = "/var/tmp/ca.crt"
}


resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.55"
  name        = "tf_csvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
    vservername = citrixadc_csvserver.tf_csvserver.name
	certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
	snicert = true
}
`

func TestAccSslvserver_sslcertkey_binding_lb(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslvserver_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslvserver_sslcertkey_binding_lb_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccSslvserver_sslcertkey_binding_lb_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccSslvserver_sslcertkey_binding_lb_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
				),
			},
		},
	})
}

func TestAccSslvserver_sslcertkey_binding_cs(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslvserver_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslvserver_sslcertkey_binding_cs_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccSslvserver_sslcertkey_binding_cs_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccSslvserver_sslcertkey_binding_cs_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
				),
			},
		},
	})
}

func testAccCheckSslvserver_sslcertkey_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslvserver_sslcertkey_binding name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID
		idSlice := strings.SplitN(bindingId, ",", 2)

		vservername := idSlice[0]
		certkeyname := idSlice[1]

		findParams := netscaler.FindParams{
			ResourceType:             "sslvserver_sslcertkey_binding",
			ResourceName:             vservername,
			ResourceMissingErrorCode: 461,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right certkeyname
		foundIndex := -1
		for i, v := range dataArr {
			if v["certkeyname"].(string) == certkeyname {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find binding %v", bindingId)
		}

		return nil
	}
}

func testAccCheckSslvserver_sslcertkey_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslvserver_sslcertkey_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Sslvserver_sslcertkey_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslvserver_sslcertkey_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
