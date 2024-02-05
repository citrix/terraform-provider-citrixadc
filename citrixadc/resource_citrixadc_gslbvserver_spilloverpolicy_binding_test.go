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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strings"
	"testing"
)

const testAccGslbvserver_spilloverpolicy_binding_basic = `

	resource "citrixadc_spilloveraction" "tf_spilloveraction" {
		name   = "my_spilloveraction"
		action = "SPILLOVER"
	}
	resource "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
		name    = "tf_spilloverpolicy"
		rule    = "true"
		action  = citrixadc_spilloveraction.tf_spilloveraction.name
		comment = "This is example of spilloverpolicy"
	}

resource "citrixadc_gslbvserver_spilloverpolicy_binding" "tf_gslbvserver_spilloverpolicy_binding" {
	name       = citrixadc_gslbvserver.tf_gslbvserver.name
	policyname = citrixadc_spilloverpolicy.tf_spilloverpolicy.name
	priority   = 100
  
  }
  
  resource "citrixadc_gslbvserver" "tf_gslbvserver" {
	dnsrecordtype = "A"
	name          = "gslb_vserver"
	servicetype   = "HTTP"
	domain {
	  domainname = "www.fooco.co"
	  ttl        = "60"
	}
	domain {
	  domainname = "www.barco.com"
	  ttl        = "65"
	}
  }
`

const testAccGslbvserver_spilloverpolicy_binding_basic_step2 = `

	resource "citrixadc_spilloveraction" "tf_spilloveraction" {
		name   = "my_spilloveraction"
		action = "SPILLOVER"
	}
	resource "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
		name    = "tf_spilloverpolicy"
		rule    = "true"
		action  = citrixadc_spilloveraction.tf_spilloveraction.name
		comment = "This is example of spilloverpolicy"
	}

  resource "citrixadc_gslbvserver" "tf_gslbvserver" {
	dnsrecordtype = "A"
	name          = "gslb_vserver"
	servicetype   = "HTTP"
	domain {
	  domainname = "www.fooco.co"
	  ttl        = "60"
	}
	domain {
	  domainname = "www.barco.com"
	  ttl        = "65"
	}
  }
`

func TestAccGslbvserver_spilloverpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGslbvserver_spilloverpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbvserver_spilloverpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserver_spilloverpolicy_bindingExist("citrixadc_gslbvserver_spilloverpolicy_binding.tf_gslbvserver_spilloverpolicy_binding", nil),
				),
			},
			{
				Config: testAccGslbvserver_spilloverpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserver_spilloverpolicy_bindingNotExist("citrixadc_gslbvserver_spilloverpolicy_binding.tf_gslbvserver_spilloverpolicy_binding", "gslb_vserver,tf_spilloverpolicy"),
				),
			},
		},
	})
}

func testAccCheckGslbvserver_spilloverpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbvserver_spilloverpolicy_binding id is set")
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
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "gslbvserver_spilloverpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("gslbvserver_spilloverpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbvserver_spilloverpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "gslbvserver_spilloverpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("gslbvserver_spilloverpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckGslbvserver_spilloverpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbvserver_spilloverpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Gslbvserver_spilloverpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbvserver_spilloverpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
