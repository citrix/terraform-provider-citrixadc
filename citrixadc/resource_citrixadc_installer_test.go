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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccInstaller_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("Install not available in CPX")
	}
	// if isCluster {
	// 	t.Skip("Install not available in Cluster")
	// }
	if adcTestbed != "INSTALLER" {
		t.Skipf("ADC testbed is %s. Expected INSTALLER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccInstaller_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInstallExist("citrixadc_installer.tf_installer", nil),
				),
			},
		},
	})
}

func testAccCheckInstallExist(n string, id *string) resource.TestCheckFunc {
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

		return nil
	}
}

const testAccInstaller_basic = `
resource "citrixadc_installer" "tf_installer" {
#url =  "file:///var/tmp/build_mana_52_24_nc_64.tgz"
	url =  "file:///var/tmp/build_mana_47_24_nc_64.tgz"
    y = true
    l = false
    wait_until_reachable = true
}
`
