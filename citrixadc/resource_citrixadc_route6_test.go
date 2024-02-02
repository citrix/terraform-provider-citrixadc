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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
	//"strconv"
)

const testAccRoute6_basic = `

resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 2
}
resource "citrixadc_route6" "tf_route6" {
	network  = "2001:db8:85a3::/64"
	vlan     = citrixadc_vlan.tf_vlan.vlanid
	weight   = 5
	distance = 3
  }
  
`
const testAccRoute6_update = `

resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 2
}
resource "citrixadc_route6" "tf_route6" {
	network  = "2001:db8:85a3::/64"
	vlan     = citrixadc_vlan.tf_vlan.vlanid
	weight   = 6
	distance = 4
  }
  
`

func TestAccRoute6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRoute6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRoute6_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoute6Exist("citrixadc_route6.tf_route6", nil),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "vlan", "2"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "weight", "5"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "distance", "3"),
				),
			},
			{
				Config: testAccRoute6_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoute6Exist("citrixadc_route6.tf_route6", nil),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "vlan", "2"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "weight", "6"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "distance", "4"),
				),
			},
		},
	})
}

func testAccCheckRoute6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No route6 name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		route6Network := rs.Primary.ID
		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

		dataArr, err := nsClient.FindAllResources(service.Route6.Type())
		found := false
		for _, v := range dataArr {
			if v["network"] == route6Network &&
				v["vlan"] == rs.Primary.Attributes["vlan"] {
				found = true
				break
			}
		}

		if err != nil {
			return err
		}

		if !found {
			return fmt.Errorf("route6 %s not found", n)
		}

		return nil
	}
}

func testAccCheckRoute6Destroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_route6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		route6Network := rs.Primary.ID
		dataArr, err := nsClient.FindAllResources(service.Route6.Type())
		if err != nil {
			return err
		}
		found := false
		for _, v := range dataArr {
			if v["network"] == route6Network &&
				v["vlan"] == rs.Primary.Attributes["vlan"] {
				found = true
				break
			}
		}
		if found {
			return fmt.Errorf("route6 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
