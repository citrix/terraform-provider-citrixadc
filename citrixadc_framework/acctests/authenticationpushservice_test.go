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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccAuthenticationpushservice_add = `
	resource "citrixadc_authenticationpushservice" "tf_pushservice" {
		name            = "tf_pushservice"
		clientid        = "cliId"
		clientsecret    = "secret"
		customerid      = "cusID"
		refreshinterval = 50
	}
`
const testAccAuthenticationpushservice_update = `
	resource "citrixadc_authenticationpushservice" "tf_pushservice" {
		name            = "tf_pushservice"
		clientid        = "cliId"
		clientsecret    = "secret"
		customerid      = "cusID1"
		refreshinterval = 80
	}
`

func TestAccAuthenticationpushservice_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpushserviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationpushservice_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.tf_pushservice", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.tf_pushservice", "name", "tf_pushservice"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.tf_pushservice", "refreshinterval", "50"),
				),
			},
			{
				Config: testAccAuthenticationpushservice_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.tf_pushservice", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.tf_pushservice", "name", "tf_pushservice"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.tf_pushservice", "refreshinterval", "80"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationpushserviceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationpushservice name is set")
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
		data, err := client.FindResource("authenticationpushservice", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationpushservice %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationpushserviceDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationpushservice" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationpushservice", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationpushservice %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
