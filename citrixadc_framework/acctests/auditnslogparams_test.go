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
	"testing"
)

const testAccAuditnslogparams_basic = `

	resource "citrixadc_auditnslogparams" "tf_auditnslogparams" {
		dateformat = "DDMMYYYY"
		loglevel   = ["EMERGENCY"]
		tcp        = "ALL"
		protocolviolations = "NONE"
	}
`
const testAccAuditnslogparams_update = `

	resource "citrixadc_auditnslogparams" "tf_auditnslogparams" {
		dateformat = "MMDDYYYY"
		loglevel   = ["EMERGENCY"]
		tcp        = "NONE"
		protocolviolations = "ALL"
	}
`

func TestAccAuditnslogparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditnslogparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogparamsExist("citrixadc_auditnslogparams.tf_auditnslogparams", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "dateformat", "DDMMYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "tcp", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "protocolviolations", "NONE"),
				),
			},
			{
				Config: testAccAuditnslogparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogparamsExist("citrixadc_auditnslogparams.tf_auditnslogparams", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "dateformat", "MMDDYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "tcp", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "protocolviolations", "ALL"),
				),
			},
		},
	})
}

func testAccCheckAuditnslogparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No auditnslogparams name is set")
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
		data, err := client.FindResource(service.Auditnslogparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("auditnslogparams %s not found", n)
		}

		return nil
	}
}
