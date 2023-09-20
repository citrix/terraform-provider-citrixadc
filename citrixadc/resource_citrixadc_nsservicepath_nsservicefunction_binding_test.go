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

const testAccNsservicepath_nsservicefunction_binding_basic = `
	resource "citrixadc_nsservicepath" "tf_servicepath" {
		servicepathname = "tf_servicepath"
	}
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 20
		aliasname = "Management VLAN"
	}
	resource "citrixadc_nsservicefunction" "tf_servicefunc" {
		servicefunctionname = "tf_servicefunc"
		ingressvlan         = citrixadc_vlan.tf_vlan.vlanid
	}
	resource "citrixadc_nsservicepath_nsservicefunction_binding" "tf_binding" {
		servicepathname = citrixadc_nsservicepath.tf_servicepath.servicepathname
		servicefunction = citrixadc_nsservicefunction.tf_servicefunc.servicefunctionname
		index           = 2
	}
`

const testAccNsservicepath_nsservicefunction_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_nsservicepath" "tf_servicepath" {
		servicepathname = "tf_servicepath"
	}
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 20
		aliasname = "Management VLAN"
	}
	resource "citrixadc_nsservicefunction" "tf_servicefunc" {
		servicefunctionname = "tf_servicefunc"
		ingressvlan         = citrixadc_vlan.tf_vlan.vlanid
	}
`

func TestAccNsservicepath_nsservicefunction_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsservicepath_nsservicefunction_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsservicepath_nsservicefunction_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsservicepath_nsservicefunction_bindingExist("citrixadc_nsservicepath_nsservicefunction_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_nsservicepath_nsservicefunction_binding.tf_binding", "servicepathname", "tf_servicepath"),
					resource.TestCheckResourceAttr("citrixadc_nsservicepath_nsservicefunction_binding.tf_binding", "servicefunction", "tf_servicefunc"),
				),
			},
			{
				Config: testAccNsservicepath_nsservicefunction_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsservicepath_nsservicefunction_bindingNotExist("citrixadc_nsservicepath_nsservicefunction_binding.tf_binding", "tf_servicepath,tf_servicefunc"),
				),
			},
		},
	})
}

func testAccCheckNsservicepath_nsservicefunction_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsservicepath_nsservicefunction_binding id is set")
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

		servicepathname := idSlice[0]
		servicefunction := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "nsservicepath_nsservicefunction_binding",
			ResourceName:             servicepathname,
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
			if v["servicefunction"].(string) == servicefunction {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("nsservicepath_nsservicefunction_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsservicepath_nsservicefunction_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		servicepathname := idSlice[0]
		servicefunction := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "nsservicepath_nsservicefunction_binding",
			ResourceName:             servicepathname,
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
			if v["servicefunction"].(string) == servicefunction {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("nsservicepath_nsservicefunction_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckNsservicepath_nsservicefunction_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsservicepath_nsservicefunction_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nsservicepath_nsservicefunction_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsservicepath_nsservicefunction_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
