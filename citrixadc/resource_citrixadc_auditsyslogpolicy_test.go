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

func TestAccAuditsyslogpolicy_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("global binding causes issues with ADC version 12.0")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAuditsyslogpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditsyslogpolicy_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditsyslogpolicyExist("citrixadc_auditsyslogpolicy.tf_syslogpolicy", nil),
				),
			},
			{
				Config: testAccAuditsyslogpolicy_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditsyslogpolicyExist("citrixadc_auditsyslogpolicy.tf_syslogpolicy", nil),
				),
			},
			{
				Config: testAccAuditsyslogpolicy_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditsyslogpolicyExist("citrixadc_auditsyslogpolicy.tf_syslogpolicy", nil),
				),
			},
		},
	})
}

func testAccCheckAuditsyslogpolicyExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := client.FindResource(service.Auditsyslogpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("%s not found", n)
		}

		return nil
	}
}

func testAccCheckAuditsyslogpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_auditsyslogpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Auditsyslogpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("%s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuditsyslogpolicy_basic_step1 = `

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
    name = "tf_syslogaction"
    serverip = "10.78.60.33"
    serverport = 514
    loglevel = [
        "ERROR",
        "NOTICE",
    ]
}

resource "citrixadc_auditsyslogpolicy" "tf_syslogpolicy" {
    name = "tf_auditsyslogpolicy"
    rule = "ns_true"
    action = citrixadc_auditsyslogaction.tf_syslogaction.name

}

`

const testAccAuditsyslogpolicy_basic_step2 = `

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
    name = "tf_syslogaction"
    serverip = "10.78.60.33"
    serverport = 514
    loglevel = [
        "ERROR",
        "NOTICE",
    ]
}

resource "citrixadc_auditsyslogpolicy" "tf_syslogpolicy" {
    name = "tf_auditsyslogpolicy"
    rule = "ns_true"
    action = citrixadc_auditsyslogaction.tf_syslogaction.name

    globalbinding {
        priority = 120
        feature = "SYSTEM"
        globalbindtype = "SYSTEM_GLOBAL"
	}
}

`

const testAccAuditsyslogpolicy_basic_step3 = `

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
    name = "tf_syslogaction"
    serverip = "10.78.60.33"
    serverport = 514
    loglevel = [
        "ERROR",
        "NOTICE",
    ]
}

resource "citrixadc_auditsyslogpolicy" "tf_syslogpolicy" {
    name = "tf_auditsyslogpolicy"
    rule = "ns_false"
    action = citrixadc_auditsyslogaction.tf_syslogaction.name

    globalbinding {
        priority = 110
        feature = "SYSTEM"
        globalbindtype = "SYSTEM_GLOBAL"
	}
}

`
