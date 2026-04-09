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

// Test backward-compatible path: using authpasswd (Sensitive attribute)
const testAccSnmpuser_authpasswd_step1 = `

	variable "snmpuser_authpasswd" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_snmpuser" "tf_snmpuser" {
		name       = "test_user"
		group      = "test_group"
		authtype   = "SHA"
		authpasswd = var.snmpuser_authpasswd
		privtype   = "DES"
		privpasswd = "this_is_my_password2"
	}
`

const testAccSnmpuser_authpasswd_step2 = `

	variable "snmpuser_authpasswd_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_snmpuser" "tf_snmpuser" {
		name       = "test_user"
		group      = "test_group"
		authtype   = "SHA"
		authpasswd = var.snmpuser_authpasswd_2
		privtype   = "DES"
		privpasswd = "this_is_my_password2"
	}
`

func TestAccSnmpuser_authpasswd_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_snmpuser_authpasswd", "oldauthpass123")
	t.Setenv("TF_VAR_snmpuser_authpasswd_2", "newauthpass456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpuser_authpasswd_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExist("citrixadc_snmpuser.tf_snmpuser", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "authtype", "SHA"),
				),
			},
			{
				Config: testAccSnmpuser_authpasswd_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExist("citrixadc_snmpuser.tf_snmpuser", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "authtype", "SHA"),
				),
			},
		},
	})
}

// Test ephemeral path: using authpasswd_wo (WriteOnly attribute) with version tracker
const testAccSnmpuser_authpasswd_wo_step1 = `

	variable "snmpuser_authpasswd_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_snmpuser" "tf_snmpuser" {
		name                 = "test_user"
		group                = "test_group"
		authtype             = "SHA"
		authpasswd_wo        = var.snmpuser_authpasswd_wo
		authpasswd_wo_version = 1
		privtype             = "DES"
		privpasswd           = "this_is_my_password2"
	}
`

const testAccSnmpuser_authpasswd_wo_step2 = `

	variable "snmpuser_authpasswd_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_snmpuser" "tf_snmpuser" {
		name                 = "test_user"
		group                = "test_group"
		authtype             = "SHA"
		authpasswd_wo        = var.snmpuser_authpasswd_wo_2
		authpasswd_wo_version = 2
		privtype             = "DES"
		privpasswd           = "this_is_my_password2"
	}
`

func TestAccSnmpuser_authpasswd_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_snmpuser_authpasswd_wo", "ephemeral_auth1")
	t.Setenv("TF_VAR_snmpuser_authpasswd_wo_2", "ephemeral_auth2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpuser_authpasswd_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExist("citrixadc_snmpuser.tf_snmpuser", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "authpasswd_wo_version", "1"),
				),
			},
			{
				Config: testAccSnmpuser_authpasswd_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExist("citrixadc_snmpuser.tf_snmpuser", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "authpasswd_wo_version", "2"),
				),
			},
		},
	})
}

// Test backward-compatible path: using privpasswd (Sensitive attribute)
const testAccSnmpuser_privpasswd_step1 = `

	variable "snmpuser_privpasswd" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_snmpuser" "tf_snmpuser" {
		name       = "test_user"
		group      = "test_group"
		authtype   = "SHA"
		authpasswd = "this_is_my_password"
		privtype   = "DES"
		privpasswd = var.snmpuser_privpasswd
	}
`

const testAccSnmpuser_privpasswd_step2 = `

	variable "snmpuser_privpasswd_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_snmpuser" "tf_snmpuser" {
		name       = "test_user"
		group      = "test_group"
		authtype   = "SHA"
		authpasswd = "this_is_my_password"
		privtype   = "DES"
		privpasswd = var.snmpuser_privpasswd_2
	}
`

func TestAccSnmpuser_privpasswd_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_snmpuser_privpasswd", "oldprivpass123")
	t.Setenv("TF_VAR_snmpuser_privpasswd_2", "newprivpass456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpuser_privpasswd_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExist("citrixadc_snmpuser.tf_snmpuser", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "privtype", "DES"),
				),
			},
			{
				Config: testAccSnmpuser_privpasswd_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExist("citrixadc_snmpuser.tf_snmpuser", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "privtype", "DES"),
				),
			},
		},
	})
}

// Test ephemeral path: using privpasswd_wo (WriteOnly attribute) with version tracker
const testAccSnmpuser_privpasswd_wo_step1 = `

	variable "snmpuser_privpasswd_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_snmpuser" "tf_snmpuser" {
		name                 = "test_user"
		group                = "test_group"
		authtype             = "SHA"
		authpasswd           = "this_is_my_password"
		privtype             = "DES"
		privpasswd_wo        = var.snmpuser_privpasswd_wo
		privpasswd_wo_version = 1
	}
`

const testAccSnmpuser_privpasswd_wo_step2 = `

	variable "snmpuser_privpasswd_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_snmpuser" "tf_snmpuser" {
		name                 = "test_user"
		group                = "test_group"
		authtype             = "SHA"
		authpasswd           = "this_is_my_password"
		privtype             = "DES"
		privpasswd_wo        = var.snmpuser_privpasswd_wo_2
		privpasswd_wo_version = 2
	}
`

func TestAccSnmpuser_privpasswd_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_snmpuser_privpasswd_wo", "ephemeral_priv1")
	t.Setenv("TF_VAR_snmpuser_privpasswd_wo_2", "ephemeral_priv2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpuser_privpasswd_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExist("citrixadc_snmpuser.tf_snmpuser", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "privpasswd_wo_version", "1"),
				),
			},
			{
				Config: testAccSnmpuser_privpasswd_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExist("citrixadc_snmpuser.tf_snmpuser", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpuser.tf_snmpuser", "privpasswd_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccSnmpuserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpuserDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_snmpuser.tf_snmpuser_ds", "name", "test_user_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_snmpuser.tf_snmpuser_ds", "group", "test_group"),
					resource.TestCheckResourceAttr("data.citrixadc_snmpuser.tf_snmpuser_ds", "authtype", "SHA"),
				),
			},
		},
	})
}

const testAccSnmpuserDataSource_basic = `

resource "citrixadc_snmpuser" "tf_snmpuser_ds" {
	name       = "test_user_ds"
	group      = "test_group"
	authtype   = "SHA"
	authpasswd = "this_is_my_password"
	privtype   = "DES"
	privpasswd = "this_is_my_password2"
}

data "citrixadc_snmpuser" "tf_snmpuser_ds" {
	name       = citrixadc_snmpuser.tf_snmpuser_ds.name
	depends_on = [citrixadc_snmpuser.tf_snmpuser_ds]
}
`
