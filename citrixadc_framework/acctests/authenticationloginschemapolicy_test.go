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

const testAccAuthenticationloginschemapolicy_add = `
	resource "citrixadc_authenticationloginschema" "tf_loginschema" {
		name                    = "tf_loginschema"
		authenticationschema    = "LoginSchema/SingleAuth.xml"
		ssocredentials          = "YES"
		authenticationstrength  = "30"
		passwordcredentialindex = "10"
	}
	resource "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy" {
		name      = "tf_loginschemapolicy"
		rule      = "true"
		action    = citrixadc_authenticationloginschema.tf_loginschema.name
		comment   = "sample_testing"
	}
`
const testAccAuthenticationloginschemapolicy_update = `
	resource "citrixadc_authenticationloginschema" "tf_loginschema" {
		name                    = "tf_loginschema"
		authenticationschema    = "LoginSchema/SingleAuth.xml"
		ssocredentials          = "YES"
		authenticationstrength  = "30"
		passwordcredentialindex = "10"
	}
	resource "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy" {
		name      = "tf_loginschemapolicy"
		rule      = "false"
		action    = citrixadc_authenticationloginschema.tf_loginschema.name
		comment   = "samplenew_testing"
	}
`

func TestAccAuthenticationloginschemapolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationloginschemapolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationloginschemapolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemapolicyExist("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "name", "tf_loginschemapolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "comment", "sample_testing"),
				),
			},
			{
				Config: testAccAuthenticationloginschemapolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemapolicyExist("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "name", "tf_loginschemapolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "comment", "samplenew_testing"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationloginschemapolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationloginschemapolicy name is set")
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
		data, err := client.FindResource(service.Authenticationloginschemapolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationloginschemapolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationloginschemapolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationloginschemapolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationloginschemapolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationloginschemapolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationloginschemapolicyDataSource_basic = `

resource "citrixadc_authenticationloginschema" "tf_loginschema_ds" {
	name                    = "tf_loginschema_ds"
	authenticationschema    = "LoginSchema/SingleAuth.xml"
	ssocredentials          = "YES"
	authenticationstrength  = "30"
	passwordcredentialindex = "10"
}

resource "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy_ds" {
	name      = "tf_loginschemapolicy_ds"
	rule      = "true"
	action    = citrixadc_authenticationloginschema.tf_loginschema_ds.name
	comment   = "datasource_test"
}

data "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy_ds" {
	name = citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy_ds.name
}
`

func TestAccAuthenticationloginschemapolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationloginschemapolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationloginschemapolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy_ds", "name", "tf_loginschemapolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy_ds", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy_ds", "comment", "datasource_test"),
				),
			},
		},
	})
}
