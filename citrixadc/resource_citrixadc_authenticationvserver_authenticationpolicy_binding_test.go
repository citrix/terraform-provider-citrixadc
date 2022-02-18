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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

const testAccAuthenticationvserver_authenticationpolicy_binding_basic = `

		resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
			name           = "tf_authenticationvserver"
			servicetype    = "SSL"
			comment        = "new"
			authentication = "ON"
			state          = "DISABLED"
		}
		resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
			name          = "ldapaction"
			serverip      = "1.2.3.4"
			serverport    = 8080
			authtimeout   = 1
			ldaploginname = "username"
		}
		resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
			name   = "tf_authenticationpolicy"
			rule   = "true"
			action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		}
		resource "citrixadc_authenticationvserver_authenticationpolicy_binding" "tf_bind" {
			name     = citrixadc_authenticationvserver.tf_authenticationvserver.name
			policy   = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
			priority = 30
		}
`

const testAccAuthenticationvserver_authenticationpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name   = "tf_authenticationpolicy"
		rule   = "true"
		action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
`

func TestAccAuthenticationvserver_authenticationpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationvserver_authenticationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationvserver_authenticationpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserver_authenticationpolicy_bindingExist("citrixadc_authenticationvserver_authenticationpolicy_binding.tf_bind", nil),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationvserver_authenticationpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserver_authenticationpolicy_bindingNotExist("citrixadc_authenticationvserver_authenticationpolicy_binding.tf_bind", "tf_authenticationvserver,tf_authenticationpolicy"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationvserver_authenticationpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationvserver_authenticationpolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "authenticationvserver_authenticationpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("authenticationvserver_authenticationpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationvserver_authenticationpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "authenticationvserver_authenticationpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("authenticationvserver_authenticationpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationvserver_authenticationpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationvserver_authenticationpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationvserver_authenticationpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationvserver_authenticationpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
