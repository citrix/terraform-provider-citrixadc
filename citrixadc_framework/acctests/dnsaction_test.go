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

const testAccDnsaction_add = `

	resource "citrixadc_dnsprofile" "dnsprofile" {
		dnsprofilename         = "tf_profile1"
		dnsquerylogging        = "DISABLED"
		dnsanswerseclogging    = "DISABLED"
		dnsextendedlogging     = "DISABLED"
		dnserrorlogging        = "DISABLED"
		cacherecords           = "ENABLED"
		cachenegativeresponses = "ENABLED"
		dropmultiqueryrequest  = "DISABLED"
		cacheecsresponses      = "DISABLED"
	}

	resource "citrixadc_dnsaction" "dnsaction" {
		actionname       = "tf_action1"
		actiontype       = "Rewrite_Response"
		ipaddress        = ["192.0.2.20","192.0.2.56","198.51.130.10"]
		dnsprofilename   = citrixadc_dnsprofile.dnsprofile.dnsprofilename
	}
`

const testAccDnsactionDataSource_basic = `

	resource "citrixadc_dnsprofile" "dnsprofile" {
		dnsprofilename         = "tf_profile1"
		dnsquerylogging        = "DISABLED"
		dnsanswerseclogging    = "DISABLED"
		dnsextendedlogging     = "DISABLED"
		dnserrorlogging        = "DISABLED"
		cacherecords           = "ENABLED"
		cachenegativeresponses = "ENABLED"
		dropmultiqueryrequest  = "DISABLED"
		cacheecsresponses      = "DISABLED"
	}

	resource "citrixadc_dnsaction" "dnsaction" {
		actionname       = "tf_action1"
		actiontype       = "Rewrite_Response"
		ipaddress        = ["192.0.2.20","192.0.2.56","198.51.130.10"]
		dnsprofilename   = citrixadc_dnsprofile.dnsprofile.dnsprofilename
	}

	data "citrixadc_dnsaction" "dnsaction_data" {
		actionname = citrixadc_dnsaction.dnsaction.actionname
	}
`

func TestAccDnsaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsactionExist("citrixadc_dnsaction.dnsaction", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaction.dnsaction", "actionname", "tf_action1"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction.dnsaction", "actiontype", "Rewrite_Response"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction.dnsaction", "dnsprofilename", "tf_profile1"),
				),
			},
		},
	})
}

func TestAccDnsactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnsaction.dnsaction_data", "actionname", "tf_action1"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsaction.dnsaction_data", "actiontype", "Rewrite_Response"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsaction.dnsaction_data", "dnsprofilename", "tf_profile1"),
				),
			},
		},
	})
}

func testAccCheckDnsactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsaction name is set")
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
		data, err := client.FindResource(service.Dnsaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnsaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
