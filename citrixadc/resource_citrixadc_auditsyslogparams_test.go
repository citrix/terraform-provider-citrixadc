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

const testAccauditsyslogparams_basic = `

	resource "citrixadc_auditsyslogparams" "tf_auditsyslogparams" {
		dateformat = "DDMMYYYY"
		loglevel   = ["EMERGENCY"]
		tcp        = "ALL"
		protocolviolations = "NONE"
		streamanalytics = "DISABLED"
	}
`
const testAccauditsyslogparams_update = `

	resource "citrixadc_auditsyslogparams" "tf_auditsyslogparams" {
		dateformat = "MMDDYYYY"
		loglevel   = ["EMERGENCY"]
		tcp        = "NONE"
		protocolviolations = "ALL"
		streamanalytics = "ENABLED"
	}
`

func TestAccauditsyslogparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testAccauditsyslogparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckauditsyslogparamsExist("citrixadc_auditsyslogparams.tf_auditsyslogparams", nil),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "dateformat", "DDMMYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "tcp", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "protocolviolations", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "streamanalytics", "DISABLED"),
				),
			},
			{
				Config: testAccauditsyslogparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckauditsyslogparamsExist("citrixadc_auditsyslogparams.tf_auditsyslogparams", nil),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "dateformat", "MMDDYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "tcp", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "protocolviolations", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "streamanalytics", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckauditsyslogparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No auditsyslogparams name is set")
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
		data, err := client.FindResource(service.Auditsyslogparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("auditsyslogparams %s not found", n)
		}

		return nil
	}
}
