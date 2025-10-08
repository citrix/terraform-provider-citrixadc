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

func TestAccReboot_basic(t *testing.T) {
	t.Skip("Cluster does not support reboot operation")
	// if isCluster {
	// 	t.Skip("Cluster does not support reboot operation")
	// }
	if isCpxRun {
		t.Skip("CPX does not support reboot operation")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRebootDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRebooter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRebooterExist("citrixadc_rebooter.tf_rebooter", nil),
				),
			},
		},
	})
}

func testAccCheckRebooterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		return nil
	}
}

func testAccCheckRebootDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_reboot" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Reboot.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccRebooter_basic = `

resource "citrixadc_rebooter" "tf_rebooter" {
  timestamp            = "somethingrandom"
  warm                 = true
  wait_until_reachable = true


  reachable_timeout = "10m"
  reachable_poll_delay = "60s"
  reachable_poll_interval = "60s"
  reachable_poll_timeout = "20s"
}
`
