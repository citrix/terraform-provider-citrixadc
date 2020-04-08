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
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccAuditmessageaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuditmessageactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuditmessageaction_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditmessageactionExist("citrixadc_auditmessageaction.tf_msgaction", nil),
				),
			},
			resource.TestStep{
				Config: testAccAuditmessageaction_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditmessageactionExist("citrixadc_auditmessageaction.tf_msgaction", nil),
				),
			},
			resource.TestStep{
				Config: testAccAuditmessageaction_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditmessageactionExist("citrixadc_auditmessageaction.tf_msgaction", nil),
				),
			},
		},
	})
}

func testAccCheckAuditmessageactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(netscaler.Auditmessageaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Audit message action %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuditmessageactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_auditmessageaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Auditmessageaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuditmessageaction_basic_step1 = `

resource "citrixadc_auditmessageaction" "tf_msgaction" {
    name = "tf_msgaction"
    loglevel = "NOTICE"
    stringbuilderexpr = "\"hello\""
    logtonewnslog = "YES"
}

`

const testAccAuditmessageaction_basic_step2 = `

resource "citrixadc_auditmessageaction" "tf_msgaction" {
    name = "tf_msgaction"
    loglevel = "DEBUG"
    stringbuilderexpr = "\"hello and bye\""
    logtonewnslog = "NO"
}

`

const testAccAuditmessageaction_basic_step3 = `

resource "citrixadc_auditmessageaction" "tf_msgaction" {
    name = "tf_msgaction2"
    loglevel = "NOTICE"
    stringbuilderexpr = "\"hello\""
    logtonewnslog = "YES"
}

`
