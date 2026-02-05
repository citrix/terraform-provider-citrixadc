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

const testAccSslpolicy_add = `
	resource "citrixadc_sslaction" "foo" {
	name                   = "tf_sslaction"
	clientauth             = "DOCLIENTAUTH"
	clientcertverification = "Mandatory"
	}

	resource "citrixadc_sslpolicy" "foo" {
	name   = "tf_sslpolicy"
	rule   = "false"
	action = citrixadc_sslaction.foo.name
	}
`

const testAccSslpolicy_update = `
	resource "citrixadc_sslaction" "foo" {
	name                   = "tf_sslaction"
	clientauth             = "DOCLIENTAUTH"
	clientcertverification = "Mandatory"
	}

	resource "citrixadc_sslpolicy" "foo" {
	name   = "tf_sslpolicy"
	rule   = "true"
	action = citrixadc_sslaction.foo.name
	}
`

const testAccSslpolicyDataSource_basic = `
	resource "citrixadc_sslaction" "foo" {
	name                   = "tf_sslaction"
	clientauth             = "DOCLIENTAUTH"
	clientcertverification = "Mandatory"
	}

	resource "citrixadc_sslpolicy" "foo" {
	name   = "tf_sslpolicy"
	rule   = "false"
	action = citrixadc_sslaction.foo.name
	}

	data "citrixadc_sslpolicy" "foo" {
		name = citrixadc_sslpolicy.foo.name
	}
`

func TestAccSslpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpolicyExist("citrixadc_sslpolicy.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "name", "tf_sslpolicy"),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "action", "tf_sslaction"),
				),
			},
			{
				Config: testAccSslpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpolicyExist("citrixadc_sslpolicy.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "name", "tf_sslpolicy"),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "action", "tf_sslaction"),
				),
			},
		},
	})
}

func testAccCheckSslpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SSL Policy name is set")
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
		data, err := client.FindResource(service.Sslpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("SSL Policy %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("SSL Policy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslpolicy.foo", "name", "tf_sslpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_sslpolicy.foo", "rule", "false"),
					resource.TestCheckResourceAttr("data.citrixadc_sslpolicy.foo", "action", "tf_sslaction"),
				),
			},
		},
	})
}
