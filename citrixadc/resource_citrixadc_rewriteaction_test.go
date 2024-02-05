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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccRewriteaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRewriteactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewriteaction_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteactionExist("citrixadc_rewriteaction.tf_rewrite_action", nil),
				),
			},
			{
				Config: testAccRewriteaction_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteactionExist("citrixadc_rewriteaction.tf_rewrite_action", nil),
				),
			},
		},
	})
}

func testAccCheckRewriteactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Rewriteaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckRewriteactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rewriteaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Rewriteaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccRewriteaction_step1 = `

resource "citrixadc_rewriteaction" "tf_rewrite_action" {
    name = "tf_rewrite_action"
    target = "HTTP.REQ.HOSTNAME"
    type = "delete"
}
`

const testAccRewriteaction_step2 = `

resource "citrixadc_rewriteaction" "tf_rewrite_action" {
    name = "tf_rewrite_action"
    target = "HTTP.REQ.COOKIE"
    type = "delete"
}
`
