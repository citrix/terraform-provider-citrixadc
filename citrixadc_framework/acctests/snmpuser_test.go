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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccSnmpuser_basic = `
	# adc CLI command add snmpgroup test_group noAuthNoPriv -readViewName test_readviewname
	resource "citrixadc_snmpuser" "tf_snmpuser" {
		name    = "test_user"
		group = "test_group"
		authtype = "SHA"
		authpasswd = "this_is_my_password"
		privtype   = "DES"
		privpasswd = "this_is_my_password2"
	}
	
	# resource "citrixadc_snmpgroup" "tf_snmpgroup" {
		#   name    = "test_group"
		#   securitylevel = "noAuthNoPriv"
		#   readviewname = "test_name"
		# }
`
const testAccSnmpuser_update = `
	# adc CLI command add snmpgroup test2_group authNoPriv -readViewName test2_readviewname
	
	resource "citrixadc_snmpuser" "tf_snmpuser" {
		name    = "test_user"
		group = "test2_group"
		authtype = "SHA"
		authpasswd = "this_is_my_second_password"
		privtype   = "AES"
		privpasswd = "this_is_my_password"
	}
	
	# resource "citrixadc_snmpgroup" "tf_snmpgroup" {
		#   name    = "test2_group"
		#   securitylevel = "test2_group"
		#   readviewname = "test_name"
		# }
`

func TestAccSnmpuser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpuser_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExist("citrixadc_snmpuser.tf_snmpuser", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "group", "test_group"),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "authtype", "SHA"),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "authpasswd", "this_is_my_password"),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "privtype", "DES"),
				),
			},
			{
				Config: testAccSnmpuser_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExist("citrixadc_snmpuser.tf_snmpuser", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "group", "test2_group"),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "authtype", "SHA"),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "privtype", "AES"),
				),
			},
		},
	})
}

func testAccCheckSnmpuserExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpuser name is set")
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
		data, err := client.FindResource(service.Snmpuser.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("snmpuser %s not found", n)
		}

		return nil
	}
}

func testAccCheckSnmpuserDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_snmpuser" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Snmpuser.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("snmpuser %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
