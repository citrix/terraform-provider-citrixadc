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
	"strings"
	"testing"
)

const testAccAuthenticationvserver_authenticationsamlidppolicy_binding_basic = `
	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey = "tf_sslcertkey"
		cert    = "/var/tmp/certificate1.crt"
		key     = "/var/tmp/key1.pem"
	}
	resource "citrixadc_authenticationsamlidpprofile" "tf_samlidpprofile" {
		name                        = "tf_samlidpprofile"
		samlspcertname              = citrixadc_sslcertkey.tf_sslcertkey.certkey
		assertionconsumerserviceurl = "http://www.example.com"
		sendpassword                = "OFF"
		samlissuername              = "new_user"
		rejectunsignedrequests      = "ON"
		signaturealg                = "RSA-SHA1"
		digestmethod                = "SHA1"
		nameidformat                = "Unspecified"
	}
	resource "citrixadc_authenticationsamlidppolicy" "tf_samlidppolicy" {
		name    = "tf_samlidppolicy"
		rule    = "false"
		action  = citrixadc_authenticationsamlidpprofile.tf_samlidpprofile.name
		comment = "aSimpleTesting"
	}
	resource "citrixadc_authenticationvserver_authenticationsamlidppolicy_binding" "tf_bind" {
		name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
		policy    = citrixadc_authenticationsamlidppolicy.tf_samlidppolicy.name
		bindpoint = "REQUEST"
		priority  = 88
		secondary = "false"
	}
`

const testAccAuthenticationvserver_authenticationsamlidppolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey = "tf_sslcertkey"
		cert    = "/var/tmp/certificate1.crt"
		key     = "/var/tmp/key1.pem"
	}
	resource "citrixadc_authenticationsamlidpprofile" "tf_samlidpprofile" {
		name                        = "tf_samlidpprofile"
		samlspcertname              = citrixadc_sslcertkey.tf_sslcertkey.certkey
		assertionconsumerserviceurl = "http://www.example.com"
		sendpassword                = "OFF"
		samlissuername              = "new_user"
		rejectunsignedrequests      = "ON"
		signaturealg                = "RSA-SHA1"
		digestmethod                = "SHA1"
		nameidformat                = "Unspecified"
	}
	resource "citrixadc_authenticationsamlidppolicy" "tf_samlidppolicy" {
		name    = "tf_samlidppolicy"
		rule    = "false"
		action  = citrixadc_authenticationsamlidpprofile.tf_samlidpprofile.name
		comment = "aSimpleTesting"
	}
`

func TestAccAuthenticationvserver_authenticationsamlidppolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAuthenticationvserver_authenticationsamlidppolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationvserver_authenticationsamlidppolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserver_authenticationsamlidppolicy_bindingExist("citrixadc_authenticationvserver_authenticationsamlidppolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccAuthenticationvserver_authenticationsamlidppolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserver_authenticationsamlidppolicy_bindingNotExist("citrixadc_authenticationvserver_authenticationsamlidppolicy_binding.tf_bind", "tf_authenticationvserver,tf_samlidppolicy"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationvserver_authenticationsamlidppolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationvserver_authenticationsamlidppolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "authenticationvserver_authenticationsamlidppolicy_binding",
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
			return fmt.Errorf("authenticationvserver_authenticationsamlidppolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationvserver_authenticationsamlidppolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "authenticationvserver_authenticationsamlidppolicy_binding",
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
			return fmt.Errorf("authenticationvserver_authenticationsamlidppolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationvserver_authenticationsamlidppolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationvserver_authenticationsamlidppolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationvserver_authenticationsamlidppolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationvserver_authenticationsamlidppolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
