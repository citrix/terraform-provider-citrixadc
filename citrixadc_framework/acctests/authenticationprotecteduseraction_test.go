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

const testAccAuthenticationprotecteduseraction_basic_step1 = `
resource "citrixadc_authenticationprotecteduseraction" "tf_authenticationprotecteduseraction" {
  name                = "tf_authenticationprotecteduseraction"
  realmstr            = "krealm1"
  maxconcurrentusers  = 8
}

`

const testAccAuthenticationprotecteduseraction_basic_step2 = `
resource "citrixadc_authenticationprotecteduseraction" "tf_authenticationprotecteduseraction" {
  name                = "tf_authenticationprotecteduseraction"
  realmstr            = "krealm2"
  maxconcurrentusers  = 10
}

`

func TestAccAuthenticationprotecteduseraction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationprotecteduseractionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationprotecteduseraction_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationprotecteduseractionExist("citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction", "name", "tf_authenticationprotecteduseraction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction", "realmstr", "krealm1"),
					resource.TestCheckResourceAttr("citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction", "maxconcurrentusers", "8"),
				),
			},
			{
				Config: testAccAuthenticationprotecteduseraction_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationprotecteduseractionExist("citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction", "name", "tf_authenticationprotecteduseraction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction", "realmstr", "krealm2"),
					resource.TestCheckResourceAttr("citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction", "maxconcurrentusers", "10"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationprotecteduseractionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationprotecteduseraction name is set")
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
		data, err := client.FindResource(service.Authenticationprotecteduseraction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationprotecteduseraction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationprotecteduseractionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationprotecteduseraction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationprotecteduseraction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationprotecteduseraction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationprotecteduseractionDataSource_basic = `

resource "citrixadc_authenticationprotecteduseraction" "tf_authenticationprotecteduseraction" {
  name                = "tf_authenticationprotecteduseraction"
  realmstr            = "krealm1"
  maxconcurrentusers  = 8
}

data "citrixadc_authenticationprotecteduseraction" "tf_authenticationprotecteduseraction" {
  name       = citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction.name
  depends_on = [citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction]
}
`

func TestAccAuthenticationprotecteduseractionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationprotecteduseractionDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction", "name", "tf_authenticationprotecteduseraction"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction", "realmstr", "krealm1"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationprotecteduseraction.tf_authenticationprotecteduseraction", "maxconcurrentusers", "8"),
				),
			},
		},
	})
}
