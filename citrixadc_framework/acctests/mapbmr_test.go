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

const testAccMapbmr_basic = `

	resource "citrixadc_mapbmr" "tf_mapbmr" {
		name           = "tf_mapbmr"
		ruleipv6prefix = "2001:db8:abcd:12::/64"
		psidoffset     = 6
		eabitlength    = 16
		psidlength     = 8
	}
`

func TestAccMapbmr_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMapbmrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMapbmr_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMapbmrExist("citrixadc_mapbmr.tf_mapbmr", nil),
					resource.TestCheckResourceAttr("citrixadc_mapbmr.tf_mapbmr", "name", "tf_mapbmr"),
					resource.TestCheckResourceAttr("citrixadc_mapbmr.tf_mapbmr", "ruleipv6prefix", "2001:db8:abcd:12::/64"),
					resource.TestCheckResourceAttr("citrixadc_mapbmr.tf_mapbmr", "psidoffset", "6"),
				),
			},
		},
	})
}

func testAccCheckMapbmrExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No mapbmr name is set")
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
		data, err := client.FindResource("mapbmr", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("mapbmr %s not found", n)
		}

		return nil
	}
}

func testAccCheckMapbmrDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_mapbmr" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("mapbmr", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("mapbmr %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
