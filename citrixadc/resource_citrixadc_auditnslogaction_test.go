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
	"testing"
)

const testAccAuditnslogaction_basic = `


resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
	name     = "my_auditnslogaction"
	serverip = "10.222.74.180"
	loglevel = ["ALERT", "CRITICAL"]
	tcp      = "ALL"
	acl      = "ENABLED"
  }
  
`
const testAccAuditnslogaction_update= `


resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
	name     = "my_auditnslogaction"
	serverip = "10.222.74.180"
	loglevel = ["ALERT", "CRITICAL"]
	tcp      = "NONE"
	acl      = "DISABLED"
  }
  
`

func TestAccAuditnslogaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuditnslogactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuditnslogaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogactionExist("citrixadc_auditnslogaction.tf_auditnslogaction", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "name", "my_auditnslogaction"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "serverip", "10.222.74.180"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "tcp", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "acl", "ENABLED"),
				),
			},
			resource.TestStep{
				Config: testAccAuditnslogaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogactionExist("citrixadc_auditnslogaction.tf_auditnslogaction", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "name", "my_auditnslogaction"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "serverip", "10.222.74.180"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "tcp", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "acl", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckAuditnslogactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No auditnslogaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Auditnslogaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("auditnslogaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuditnslogactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_auditnslogaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Auditnslogaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("auditnslogaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
