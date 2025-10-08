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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAuditsyslogaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAuditsyslogactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditsyslogaction_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditsyslogactionExist("citrixadc_auditsyslogaction.tf_syslogaction", nil, map[string]interface{}{"name": "tf_syslogaction", "serverip": "10.78.60.33", "serverport": 514, "transport": "TCP", "loglevel": []string{"ERROR", "NOTICE"}}),
				),
			},
			{
				Config: testAccAuditsyslogaction_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditsyslogactionExist("citrixadc_auditsyslogaction.tf_syslogaction", nil, map[string]interface{}{"name": "tf_syslogaction", "serverip": "10.78.60.34", "serverport": 514, "transport": "TCP", "loglevel": []string{"ALL"}}),
				),
			},
			{
				Config: testAccAuditsyslogaction_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditsyslogactionExist("citrixadc_auditsyslogaction.tf_syslogaction", nil, map[string]interface{}{"name": "tf_syslogaction", "serverip": "10.78.60.34", "serverport": 514, "transport": "UDP", "loglevel": []string{"NONE"}}),
				),
			},
			{
				Config: testAccAuditsyslogaction_basic_step4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditsyslogactionExist("citrixadc_auditsyslogaction.tf_syslogaction", nil, map[string]interface{}{"name": "tf_syslogaction", "serverip": "10.78.60.34", "serverport": 514, "transport": "UDP", "loglevel": []string{"ALL"}, "managementlog": []string{"ALL"}, "mgmtloglevel": []string{"ALL"}, "syslogcompliance": "RFC5424", "streamanalytics": "ENABLED"}),
				),
			},
		},
	})
}

func testAccCheckAuditsyslogactionExist(n string, id *string, expectedValues map[string]interface{}) resource.TestCheckFunc {
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Auditsyslogaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("%s not found", n)
		}

		// Iterate through all expected values and validate them
		for key, expectedValue := range expectedValues {
			if actualValue, exists := data[key]; !exists {
				return fmt.Errorf("Expected key %q not found in retrieved data", key)
			} else if !compareValues(expectedValue, actualValue) {
				return fmt.Errorf("Expected value for %q differs. Expected: %v, Retrieved: %v",
					key, expectedValue, actualValue)
			}
		}

		return nil
	}
}

func testAccCheckAuditsyslogactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_auditsyslogaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Auditsyslogaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("%s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuditsyslogaction_basic_step1 = `

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
    name = "tf_syslogaction"
    serverip = "10.78.60.33"
    serverport = 514
    loglevel = [
        "ERROR",
        "NOTICE",
    ]
	transport = "TCP"
}
`

const testAccAuditsyslogaction_basic_step2 = `

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
    name = "tf_syslogaction"
    serverip = "10.78.60.34"
    serverport = 514
    loglevel = [
        "ALL",
    ]
	transport = "TCP"
}
`

const testAccAuditsyslogaction_basic_step3 = `

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
    name = "tf_syslogaction"
    serverip = "10.78.60.34"
    serverport = 514
    loglevel = [
        "NONE",
    ]
	transport = "UDP"
}
`

const testAccAuditsyslogaction_basic_step4 = `

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
    name = "tf_syslogaction"
    serverip = "10.78.60.34"
    serverport = 514
	transport = "UDP"
	loglevel = [ "ALL" ]
	managementlog = [ "ALL" ]
	mgmtloglevel = [ "ALL" ]
	syslogcompliance = "RFC5424"
	streamanalytics = "ENABLED"
}
`
