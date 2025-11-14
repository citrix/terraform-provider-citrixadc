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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
    ocspcheck = "Mandatory"
	snicert = false
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
	snicert = false
	ca = false
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
	ca = false
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
	ocspcheck = "Mandatory"
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
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckSslvserver_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_sslcertkey_binding_lb_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "ocspcheck", "Mandatory"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "snicert", "false"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "ca", "true"),
				),
			},
			{
				Config: testAccSslvserver_sslcertkey_binding_lb_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "snicert", "false"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "ca", "false"),
				),
			},
			{
				Config: testAccSslvserver_sslcertkey_binding_lb_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "snicert", "true"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "ca", "false"),
				),
			},
		},
	})
}

func TestAccSslvserver_sslcertkey_binding_cs(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckSslvserver_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_sslcertkey_binding_cs_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "ocspcheck", "Mandatory"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "snicert", "false"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "ca", "true"),
				),
			},
			{
				Config: testAccSslvserver_sslcertkey_binding_cs_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "snicert", "false"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "ca", "false"),
				),
			},
			{
				Config: testAccSslvserver_sslcertkey_binding_cs_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcertkey_bindingExist("citrixadc_sslvserver_sslcertkey_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "snicert", "true"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcertkey_binding.tf_binding", "ca", "false"),
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID
		idSlice := strings.Split(bindingId, ",")

		vservername := idSlice[0]
		certkeyname := idSlice[1]
		snicert := false
		ca := false
		if val, ok := rs.Primary.Attributes["ca"]; ok {
			ca = val == "true"
		}
		if val, ok := rs.Primary.Attributes["snicert"]; ok {
			snicert = val == "true"
		}

		findParams := service.FindParams{
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
			if v["certkeyname"].(string) == certkeyname && v["snicert"].(bool) == snicert && v["ca"].(bool) == ca {
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
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslvserver_sslcertkey_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslvserver_sslcertkey_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslvserver_sslcertkey_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
