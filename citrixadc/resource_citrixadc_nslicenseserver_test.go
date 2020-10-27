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

func TestAccNslicenseserver_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("Feature not supported in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNslicenseserverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNslicenseserver_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslicenseserverExist("citrixadc_nslicenseserver.tf_licenseserver", nil),
				),
			},
		},
	})
}

func testAccCheckNslicenseserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id has been set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		findParams := netscaler.FindParams{
			ResourceType: "nslicenseserver",
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		licenseServers, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		if len(licenseServers) == 0 {
			return fmt.Errorf("Could not find %s license server", rs.Primary.ID)
		}
		if licenseServers[0]["servername"] != rs.Primary.ID {
			return fmt.Errorf("Wrong servername %s. Expected %s", licenseServers[0]["servername"], rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckNslicenseserverDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nslicenseserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("nslicenseserver", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNslicenseserver_basic = `
resource "citrixadc_nslicenseserver" "tf_licenseserver" {
    servername = "10.78.60.200"
    port = 27000
}
`
