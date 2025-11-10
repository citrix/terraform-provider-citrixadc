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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccSubscriberparam_basic = `

resource "citrixadc_subscriberparam" "tf_subscriberparam" {
	keytype       = "IP"
	interfacetype = "None"
	idlettl       = 40
	idleaction    = "ccrTerminate"
	ipv6prefixlookuplist = [64]
	}
  
`
const testAccSubscriberparam_update = `

resource "citrixadc_subscriberparam" "tf_subscriberparam" {
	keytype       = "IP"
	interfacetype = "RadiusOnly"
	idlettl       = 50
	idleaction    = "ccrTerminate"
	ipv6prefixlookuplist = [64]
	}
  
`

func TestAccSubscriberparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriberparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberparamExist("citrixadc_subscriberparam.tf_subscriberparam", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "keytype", "IP"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "interfacetype", "None"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "idlettl", "40"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "idleaction", "ccrTerminate"),
				),
			},
			{
				Config: testAccSubscriberparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberparamExist("citrixadc_subscriberparam.tf_subscriberparam", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "keytype", "IP"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "interfacetype", "RadiusOnly"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "idlettl", "50"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "idleaction", "ccrTerminate"),
				),
			},
		},
	})
}

func testAccCheckSubscriberparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No subscriberparam name is set")
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
		data, err := client.FindResource("subscriberparam", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("subscriberparam %s not found", n)
		}

		return nil
	}
}
