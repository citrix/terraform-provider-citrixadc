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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccNslicense_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("ssh does not work correctly with CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNslicense_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslicenseExist("citrixadc_nslicense.tf_license", nil),
				),
			},
		},
	})
}

func testAccCheckNslicenseExist(n string, id *string) resource.TestCheckFunc {
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

const testAccNslicense_basic = `
resource "citrixadc_nslicense" "tf_license" {

    license_file = "CNS_V10000_SERVER_PLT_Retail.lic"
    ssh_host_pubkey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDaA2H70ONYk1JDPHmqKNoOYzLZeR8jNu252P63OsI+N1k4hHQUPeysV20vzeDqgtDOoOkb90By9ryRTjGDOzxers04B23+BM+gaTFp0ONNr8uCLNt5mtZXK6dp2JjYpysl3qmpDDZ4qYhoDikliL05+bO/3dEpK6kOo25DjwjHsJDK8HovAiLdHg7v6Y6PTbJseT/+pae+0P0/gBFY901cEeB/DJqzyH7Qd1lUuUroy9buROTVhkF5VdaaPQJK8YX2oH8ocoqQOHxrSfh3U0+OuboQSyle5MnFjO88yRJrRwpT1ooJGse3xWf/0Zd5/gbuZTzswqPen2x0JN3iIvpekKItcTEegy9JlVFPEtcLeO738uYJxJuSen2HECmtl9LFjtFkLRkC5/t7qZK3SCvkKaEF/ol2K53aOPd5P9K6mYtc9xJvgtX1gntuDMuxNZBoZCeX/+5dxL0SAro9bBY0ArwpnhAo7xYgdY7F7RsXvNBJuZZiZQvFJNqnFtteKbk="
}
`
