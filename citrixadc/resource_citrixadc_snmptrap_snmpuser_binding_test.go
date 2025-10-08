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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccSnmptrap_snmpuser_binding_basic = `
	resource "citrixadc_snmpuser" "tf_snmpuser" {
	name       = "tf_snmpuser"
	group      = "all_group"
	authtype   = "SHA"
	authpasswd = "secretpassword"
	privtype   = "AES"
	privpasswd = "secretpassword"
	}
	resource "citrixadc_snmptrap" "tf_snmptrap" {
	trapclass       = "generic"
	trapdestination = "10.50.50.10"
	version         = "V3"
	}
	resource "citrixadc_snmptrap_snmpuser_binding" "tf_binding" {
	trapclass       = citrixadc_snmptrap.tf_snmptrap.trapclass
	trapdestination = citrixadc_snmptrap.tf_snmptrap.trapdestination
	username        = citrixadc_snmpuser.tf_snmpuser.name
	securitylevel   = "authPriv"
	}
`

// const testAccSnmptrap_snmpuser_binding_basic_step2 = `
// 	# Keep the above bound resources without the actual binding to check proper deletion
// 	resource "citrixadc_snmpuser" "tf_snmpuser" {
// 		name       = "tf_snmpuser"
// 		group      = "all_group"
// 		authtype   = "SHA"
// 		authpasswd = "secretpassword"
// 		privtype   = "AES"
// 		privpasswd = "secretpassword"
// 	}
// 	resource "citrixadc_snmptrap" "tf_snmptrap" {
// 		trapclass       = "generic"
// 		trapdestination = "10.50.50.10"
// 		version         = "V3"
// 		td = 0
// 	}
// `

func TestAccSnmptrap_snmpuser_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckSnmptrap_snmpuser_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmptrap_snmpuser_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmptrap_snmpuser_bindingExist("citrixadc_snmptrap_snmpuser_binding.tf_binding", nil),
				),
			},
			// resource.TestStep{
			// 	Config: testAccSnmptrap_snmpuser_binding_basic_step2,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckSnmptrap_snmpuser_bindingNotExist("citrixadc_snmptrap_snmpuser_binding.tf_binding", "generic,10.50.50.10,tf_snmpuser"),
			// 	),
			// },
		},
	})
}

func testAccCheckSnmptrap_snmpuser_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmptrap_snmpuser_binding id is set")
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

		idSlice := strings.SplitN(bindingId, ",", 3)

		trapclass := idSlice[0]
		trapdestination := idSlice[1]
		username := idSlice[2]

		args := make(map[string]string, 0)
		args["trapclass"] = trapclass
		args["trapdestination"] = trapdestination
		args["version"] = rs.Primary.Attributes["version"]
		args["td"] = rs.Primary.Attributes["td"]

		findParams := service.FindParams{
			ResourceType:             "snmptrap_snmpuser_binding",
			ArgsMap:                  args,
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
			if v["username"].(string) == username {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("snmptrap_snmpuser_binding %s not found", n)
		}

		return nil
	}
}

// FIXME: Know how to access the other attributes of the resource other than attributes that are included in the ID of the resource
// func testAccCheckSnmptrap_snmpuser_bindingNotExist(n string, id string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		Use the shared utility function to get a configured client
// 		client, err := testAccGetClient()
// 		if err != nil {
// 			return fmt.Errorf("Failed to get test client: %v", err)
// 		}
// 		rs, _ := s.RootModule().Resources[n]
// 		if !strings.Contains(id, ",") {
// 			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
// 		}

// 		bindingId := rs.Primary.ID

// 		idSlice := strings.SplitN(bindingId, ",", 3)

// 		trapclass := idSlice[0]
// 		trapdestination := idSlice[1]
// 		username := idSlice[2]

// 		args := make(map[string]string, 0)
// 		args["trapclass"] = trapclass
// 		args["trapdestination"] = trapdestination
// 		args["version"] = rs.Primary.Attributes["version"]
// 		args["td"] = rs.Primary.Attributes["td"]

// 		findParams := service.FindParams{
// 			ResourceType:             "snmptrap_snmpuser_binding",
// 			ArgsMap:                  args,
// 			ResourceMissingErrorCode: 258,
// 		}
// 		dataArr, err := client.FindResourceArrayWithParams(findParams)

// 		// Unexpected error
// 		if err != nil {
// 			return err
// 		}

// 		// Iterate through results to hopefully not find the one with the matching secondIdComponent
// 		found := false
// 		for _, v := range dataArr {
// 			if v["username"].(string) == username {
// 				found = true
// 				break
// 			}
// 		}

// 		if found {
// 			return fmt.Errorf("snmptrap_snmpuser_binding %s was found, but it should have been destroyed", n)
// 		}

// 		return nil
// 	}
// }

func testAccCheckSnmptrap_snmpuser_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_snmptrap_snmpuser_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Snmptrap_snmpuser_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("snmptrap_snmpuser_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
