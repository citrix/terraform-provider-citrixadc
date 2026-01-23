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

const testAccNshostname_basic = `

resource "citrixadc_nshostname" "tf_nshostname" {
	hostname = "mycitrix_adc"
	}
`
const testAccNshostname_update = `

resource "citrixadc_nshostname" "tf_nshostname" {
	hostname = "mycitrix_adc_rename"
	}
 `

func TestAccNshostname_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNshostname_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshostnameExist("citrixadc_nshostname.tf_nshostname", nil),
					resource.TestCheckResourceAttr("citrixadc_nshostname.tf_nshostname", "hostname", "mycitrix_adc"),
				),
			},
			{
				Config: testAccNshostname_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshostnameExist("citrixadc_nshostname.tf_nshostname", nil),
					resource.TestCheckResourceAttr("citrixadc_nshostname.tf_nshostname", "hostname", "mycitrix_adc_rename"),
				),
			},
		},
	})
}

func testAccCheckNshostnameExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nshostname name is set")
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
		data, err := client.FindResource(service.Nshostname.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nshostname %s not found", n)
		}

		return nil
	}
}
