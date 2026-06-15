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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLbmonitor_sslcertkey_binding_basic = `
	resource "citrixadc_lbmonitor_sslcertkey_binding" "tf_lbmonitor_sslcertkey_binding" {
		monitorname = citrixadc_lbmonitor.tf_monitor.monitorname
		certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
	}

	resource "citrixadc_lbmonitor" "tf_monitor" {
		monitorname = "tf_monitor"
		type = "HTTP"
		sslprofile = "ns_default_ssl_profile_backend"
	}

	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey = "tf_sslcertkey"
		cert = "/var/tmp/certificate1.crt"
		key = "/var/tmp/key1.pem"
	}
`

const testAccLbmonitor_sslcertkey_binding_basic_step2 = `
	resource "citrixadc_lbmonitor" "tf_monitor" {
		monitorname = "tf_monitor"
		type = "HTTP"
		sslprofile = "ns_default_ssl_profile_backend"
	}

	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey = "tf_sslcertkey"
		cert = "/var/tmp/certificate1.crt"
		key = "/var/tmp/key1.pem"
	}
`

const testAccLbmonitor_sslcertkey_bindingDataSource_basic = `
	resource "citrixadc_lbmonitor" "tf_monitor" {
		monitorname = "tf_monitor"
		type = "HTTP"
		sslprofile = "ns_default_ssl_profile_backend"
	}

	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey = "tf_sslcertkey"
		cert = "/var/tmp/certificate1.crt"
		key = "/var/tmp/key1.pem"
	}

	resource "citrixadc_lbmonitor_sslcertkey_binding" "tf_lbmonitor_sslcertkey_binding" {
		monitorname = citrixadc_lbmonitor.tf_monitor.monitorname
		certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
	}

	data "citrixadc_lbmonitor_sslcertkey_binding" "tf_lbmonitor_sslcertkey_binding" {
		monitorname = citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding.monitorname
		certkeyname = citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding.certkeyname
		ca          = false
		depends_on  = [citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding]
	}
`

const testAccLbmonitor_sslcertkey_binding_ca = `
	resource "citrixadc_lbmonitor_sslcertkey_binding" "tf_lbmonitor_sslcertkey_binding" {
		ca = true
		monitorname = citrixadc_lbmonitor.tf_monitor.monitorname
		certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
	}

	resource "citrixadc_lbmonitor" "tf_monitor" {
		monitorname = "tf_monitor_ca"
		type = "HTTP"
		sslprofile = "ns_default_ssl_profile_backend"
	}

	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey = "tf_sslcertkey_ca"
		cert = "/nsconfig/ssl/rootcert2.cert"
		key = "/nsconfig/ssl/rootcert2.key"
	}
`

const testAccLbmonitor_sslcertkey_binding_ca_step2 = `
	resource "citrixadc_lbmonitor" "tf_monitor" {
		monitorname = "tf_monitor_ca"
		type = "HTTP"
		sslprofile = "ns_default_ssl_profile_backend"
	}

	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey = "tf_sslcertkey_ca"
		cert = "/nsconfig/ssl/rootcert2.cert"
		key = "/nsconfig/ssl/rootcert2.key"
	}
`

func TestAccLbmonitor_sslcertkey_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitor_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_sslcertkey_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitor_sslcertkey_bindingExist("citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding", nil),
				),
			},
			{
				Config: testAccLbmonitor_sslcertkey_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitor_sslcertkey_bindingNotExist("citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding", "tf_monitor,tf_sslcertkey"),
				),
			},
		},
	})
}

func TestAccLbmonitor_sslcertkey_binding_ca(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitor_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_sslcertkey_binding_ca,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitor_sslcertkey_bindingExist("citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding", "ca", "true"),
				),
			},
			{
				Config: testAccLbmonitor_sslcertkey_binding_ca_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitor_sslcertkey_bindingNotExist("citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding", "tf_monitor_ca,tf_sslcertkey_ca"),
				),
			},
		},
	})
}

func testAccCheckLbmonitor_sslcertkey_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbmonitor_sslcertkey_binding id is set")
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

		bindingId := rs.Primary.ID

		idMap, _, err := utils.ParseIdString(bindingId, []string{"monitorname", "certkeyname"}, nil)
		if err != nil {
			return err
		}
		monitorName := idMap["monitorname"]
		sslcertkeyName := idMap["certkeyname"]

		findParams := service.FindParams{
			ResourceType:             "lbmonitor_sslcertkey_binding",
			ResourceName:             monitorName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["certkeyname"].(string) == sslcertkeyName {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lbmonitor_sslcertkey_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbmonitor_sslcertkey_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"monitorname", "certkeyname"}, nil)
		if err != nil {
			return err
		}
		lbmonitorName := idMap["monitorname"]
		sslcertkeyName := idMap["certkeyname"]

		findParams := service.FindParams{
			ResourceType:             "lbmonitor_sslcertkey_binding",
			ResourceName:             lbmonitorName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		found := false
		for _, v := range dataArr {
			if v["certkeyname"].(string) == sslcertkeyName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lbmonitor_sslcertkey_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbmonitor_sslcertkey_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbmonitor_sslcertkey_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbmonitor_sslcertkey_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbmonitor_sslcertkey_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLbmonitor_sslcertkey_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_sslcertkey_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding", "monitorname", "tf_monitor"),
					resource.TestCheckResourceAttr("data.citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding", "certkeyname", "tf_sslcertkey"),
					resource.TestCheckResourceAttr("data.citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding", "ca", "false"),
				),
			},
		},
	})
}
