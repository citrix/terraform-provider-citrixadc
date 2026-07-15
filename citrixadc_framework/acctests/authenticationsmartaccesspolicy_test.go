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

const testAccAuthenticationsmartaccesspolicy_basic_step1 = `
resource "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name    = "tf_authenticationsmartaccessprofile"
  tags    = "test_tag"
  comment = "test_comment"
}

resource "citrixadc_authenticationsmartaccesspolicy" "tf_authenticationsmartaccesspolicy" {
  name    = "tf_authenticationsmartaccesspolicy"
  action  = citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile.name
  rule    = "TRUE"
  comment = "test_comment"
}

`

const testAccAuthenticationsmartaccesspolicy_basic_step2 = `
resource "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name    = "tf_authenticationsmartaccessprofile"
  tags    = "test_tag"
  comment = "test_comment"
}
resource "citrixadc_authenticationsmartaccesspolicy" "tf_authenticationsmartaccesspolicy" {
  name    = "tf_authenticationsmartaccesspolicy"
  action  = citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile.name
  rule    = "FALSE"
  comment = "test_comment_updated"
}

`

func TestAccAuthenticationsmartaccesspolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationsmartaccesspolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsmartaccesspolicy_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsmartaccesspolicyExist("citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "name", "tf_authenticationsmartaccesspolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "action", "tf_authenticationsmartaccessprofile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "rule", "TRUE"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "comment", "test_comment"),
				),
			},
			{
				Config: testAccAuthenticationsmartaccesspolicy_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsmartaccesspolicyExist("citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "name", "tf_authenticationsmartaccesspolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "action", "tf_authenticationsmartaccessprofile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "rule", "FALSE"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "comment", "test_comment_updated"),
				),
			},
		},
	})
}

func TestAccAuthenticationsmartaccesspolicy_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationsmartaccesspolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsmartaccesspolicy_basic_step1,
			},
			{
				Config:                  testAccAuthenticationsmartaccesspolicy_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckAuthenticationsmartaccesspolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationsmartaccesspolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationsmartaccesspolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationsmartaccesspolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationsmartaccesspolicyDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationsmartaccesspolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationsmartaccesspolicy name is set")
		}

		_, err := client.FindResource(service.Authenticationsmartaccesspolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationsmartaccesspolicy %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccAuthenticationsmartaccesspolicyDataSource_basic = `

resource "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name    = "tf_authenticationsmartaccessprofile"
  tags    = "test_tag"
  comment = "test_comment"
}

resource "citrixadc_authenticationsmartaccesspolicy" "tf_authenticationsmartaccesspolicy" {
  name    = "tf_authenticationsmartaccesspolicy"
  action  = citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile.name
  rule    = "TRUE"
  comment = "test_comment"
}

data "citrixadc_authenticationsmartaccesspolicy" "tf_authenticationsmartaccesspolicy" {
  name       = citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy.name
  depends_on = [citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy]
}
`

func TestAccAuthenticationsmartaccesspolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsmartaccesspolicyDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "name", "tf_authenticationsmartaccesspolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "action", "tf_authenticationsmartaccessprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "rule", "TRUE"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy", "comment", "test_comment"),
				),
			},
		},
	})
}
