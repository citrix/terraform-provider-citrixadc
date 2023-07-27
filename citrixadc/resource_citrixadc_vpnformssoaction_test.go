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

const testAccVpnformssoaction_basic = `

resource "citrixadc_vpnformssoaction" "tf_vpnformssoaction" {
	name = "tf_vpnformssoaction"
	actionurl = "/home"
	userfield = "username"
	passwdfield = "password"
	ssosuccessrule = "true"
	namevaluepair = "name1=value1&name2=value2"
	nvtype = "STATIC"
	responsesize = "150"
	submitmethod = "POST"
}
`

const testAccVpnformssoaction_basic_update_mandatory_attributes = `

	resource "citrixadc_vpnformssoaction" "tf_vpnformssoaction" {
		name = "tf_vpnformssoaction"
		actionurl = "/contact"
		userfield = "username1"
		passwdfield = "password1"
		ssosuccessrule = "false"
	}
`

const testAccVpnformssoaction_basic_update_non_mandatory_attributes = `

	resource "citrixadc_vpnformssoaction" "tf_vpnformssoaction" {
		name = "tf_vpnformssoaction"
		actionurl = "/contact"
		userfield = "username1"
		passwdfield = "password1"
		ssosuccessrule = "false"
		namevaluepair = "name3=value3"
		nvtype = "DYNAMIC"
		responsesize = "151"
		submitmethod = "GET"
	}
`

func TestAccVpnformssoaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnformssoactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnformssoaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnformssoactionExist("citrixadc_vpnformssoaction.tf_vpnformssoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "name", "tf_vpnformssoaction"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "actionurl", "/home"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "userfield", "username"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "passwdfield", "password"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "ssosuccessrule", "true"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "namevaluepair", "name1=value1&name2=value2"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "nvtype", "STATIC"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "responsesize", "150"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "submitmethod", "POST"),
				),
			},
			{
				Config: testAccVpnformssoaction_basic_update_mandatory_attributes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnformssoactionExist("citrixadc_vpnformssoaction.tf_vpnformssoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "actionurl", "/contact"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "userfield", "username1"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "passwdfield", "password1"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "ssosuccessrule", "false"),
				),
			},
			{
				Config: testAccVpnformssoaction_basic_update_non_mandatory_attributes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnformssoactionExist("citrixadc_vpnformssoaction.tf_vpnformssoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "namevaluepair", "name3=value3"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "nvtype", "DYNAMIC"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "responsesize", "151"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "submitmethod", "GET"),
				),
			},
		},
	})
}

func testAccCheckVpnformssoactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnformssoaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Vpnformssoaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnformssoaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnformssoactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnformssoaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vpnformssoaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnformssoaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
