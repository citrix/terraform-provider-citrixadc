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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccLbvserver_botpolicy_binding_basic_step1 = `

resource citrixadc_lbvserver_botpolicy_binding demo_lbvserver_botpolicy_binding {
name                   = citrixadc_lbvserver.demo_lb1.name
policyname             = citrixadc_botpolicy.demo_botpolicy1.name
labeltype              = "reqvserver" # Possible values = reqvserver, resvserver, policylabel
labelname              = citrixadc_lbvserver.demo_lb1.name
priority               = 100
gotopriorityexpression = "END"
invoke                 = true         # boolean
}

resource "citrixadc_lbvserver" "demo_lb1" {
name        = "demo_lb1"
servicetype = "HTTP"
}

resource "citrixadc_botpolicy" "demo_botpolicy1" {
name        = "demo_botpolicy1"
profilename = citrixadc_botprofile.tf_botprofile1.name
rule        = "true"
comment     = "COMMENT FOR BOTPOLICY"
}

resource "citrixadc_botprofile" "tf_botprofile1" {
	name = "tf_botprofile1"
	errorurl = "http://www.citrix.com"
	trapurl = "/http://www.citrix.com"
	comment = "tf_botprofile comment"
	bot_enable_white_list = "ON"
	bot_enable_black_list = "ON"
	bot_enable_rate_limit = "ON"
	devicefingerprint = "ON"
	devicefingerprintaction = ["LOG", "RESET"]
	bot_enable_ip_reputation = "ON"
	trap = "ON"
	trapaction = ["LOG", "RESET"]
	bot_enable_tps = "ON"
}
`
const testAccLbvserver_botpolicy_binding_basic_step2 = `	
resource "citrixadc_lbvserver" "demo_lb1" {
	name        = "demo_lb1"
	servicetype = "HTTP"
}
	
resource "citrixadc_botpolicy" "demo_botpolicy1" {
	name        = "demo_botpolicy1"
	profilename = citrixadc_botprofile.tf_botprofile1.name
	rule        = "true"
	comment     = "COMMENT FOR BOTPOLICY"
	}
	
resource "citrixadc_botprofile" "tf_botprofile1" {
	name = "tf_botprofile1"
	errorurl = "http://www.citrix.com"
	trapurl = "/http://www.citrix.com"
	comment = "tf_botprofile comment"
	bot_enable_white_list = "ON"
	bot_enable_black_list = "ON"
	bot_enable_rate_limit = "ON"
	devicefingerprint = "ON"
	devicefingerprintaction = ["LOG", "RESET"]
	bot_enable_ip_reputation = "ON"
	trap = "ON"
	trapaction = ["LOG", "RESET"]
	bot_enable_tps = "ON"
}

`

func TestAccLbvserver_botpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLbvserver_botpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_botpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_botpolicy_bindingExist("citrixadc_lbvserver_botpolicy_binding.demo_lbvserver_botpolicy_binding", nil),
				),
			},
			{
				Config: testAccLbvserver_botpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_botpolicy_bindingNotExist("citrixadc_lbvserver_botpolicy_binding.demo_lbvserver_botpolicy_binding", "demo_lb1,demo_botpolicy1"),
				),
			},
		},
	})
}

func testAccCheckLbvserver_botpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_botpolicy_binding id is set")
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
			ResourceType:             "lbvserver_botpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lbvserver_botpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_botpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}

		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_botpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lbvserver_botpolicy_binding %s not deleted", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_botpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_botpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("lbvserverbotpolicybinding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_botpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
