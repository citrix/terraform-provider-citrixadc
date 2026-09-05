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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// DANGER: citrixadc_nsconfig_update issues `set ns config`, which changes the
// appliance NSIP/netmask. This test is SAFE ONLY because ipaddress/netmask are
// set to the running box's OWN current NSIP/netmask, making the `set ns config`
// an effective no-op that CANNOT disconnect the box. It must be run ONLY against
// the designated disposable box whose current NSIP is 10.101.132.152
// (NS_URL=http://10.101.132.152/). nsvlan/ifnum/tagged are intentionally omitted
// so the management data path is never disturbed.
const testAccNsconfigUpdate_basic = `
	resource "citrixadc_nsconfig_update" "foo" {
		ipaddress = "10.101.132.152"
		netmask   = "255.255.255.0"
	}
`

func TestAccNsconfigUpdate_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsconfigUpdate_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsconfigUpdateExist("citrixadc_nsconfig_update.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "ipaddress", "10.101.132.152"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "netmask", "255.255.255.0"),
				),
			},
		},
	})
}

func testAccCheckNsconfigUpdateExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No NsConfigUpdate is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Read the live nsconfig back from the ADC and confirm the settable
		// params we applied match what the appliance now reports.
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Error creating client for nsconfig read-back: %s", err.Error())
		}
		data, err := client.FindResource(service.Nsconfig.Type(), "")
		if err != nil {
			return fmt.Errorf("Error reading nsconfig from ADC: %s", err.Error())
		}

		wantIP := rs.Primary.Attributes["ipaddress"]
		if got, ok := data["ipaddress"]; ok && wantIP != "" {
			if fmt.Sprintf("%v", got) != wantIP {
				return fmt.Errorf("nsconfig ipaddress mismatch: state=%s adc=%v", wantIP, got)
			}
		}
		wantNetmask := rs.Primary.Attributes["netmask"]
		if got, ok := data["netmask"]; ok && wantNetmask != "" {
			if fmt.Sprintf("%v", got) != wantNetmask {
				return fmt.Errorf("nsconfig netmask mismatch: state=%s adc=%v", wantNetmask, got)
			}
		}
		return nil
	}
}
