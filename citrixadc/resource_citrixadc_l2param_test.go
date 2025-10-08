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

const testAccL2param_basic = `


	resource "citrixadc_l2param" "tf_l2param" {
		mbfpeermacupdate   = 20
		maxbridgecollision = 30
		bdggrpproxyarp     = "DISABLED"
	}
`

const testAccL2param_update = `


	resource "citrixadc_l2param" "tf_l2param" {
		mbfpeermacupdate   = 30
		maxbridgecollision = 40
		bdggrpproxyarp     = "ENABLED"
	}
`

func TestAccL2param_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testAccL2param_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL2paramExist("citrixadc_l2param.tf_l2param", nil),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "mbfpeermacupdate", "20"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "maxbridgecollision", "30"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "bdggrpproxyarp", "DISABLED"),
				),
			},
			{
				Config: testAccL2param_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL2paramExist("citrixadc_l2param.tf_l2param", nil),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "mbfpeermacupdate", "30"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "maxbridgecollision", "40"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "bdggrpproxyarp", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckL2paramExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No l2param name is set")
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
		data, err := client.FindResource(service.L2param.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("l2param %s not found", n)
		}

		return nil
	}
}
