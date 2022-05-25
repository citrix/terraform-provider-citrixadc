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

const testAccMapdomain_add = `
	resource "citrixadc_mapdmr" "tf_mapdmr" {
		name         = "tf_mapdmr"
		bripv6prefix = "2002:db8::/64"
	}
	resource "citrixadc_mapdomain" "tf_mapdomain" {
		name       = "tf_mapdomain"
		mapdmrname = citrixadc_mapdmr.tf_mapdmr.name
	}
`

func TestAccMapdomain_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMapdomainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccMapdomain_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMapdomainExist("citrixadc_mapdomain.tf_mapdomain", nil),
					resource.TestCheckResourceAttr("citrixadc_mapdomain.tf_mapdomain", "name", "tf_mapdomain"),
					resource.TestCheckResourceAttr("citrixadc_mapdomain.tf_mapdomain", "mapdmrname", "tf_mapdmr"),
				),
			},
		},
	})
}

func testAccCheckMapdomainExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No mapdomain name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("mapdomain", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("mapdomain %s not found", n)
		}

		return nil
	}
}

func testAccCheckMapdomainDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_mapdomain" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("mapdomain", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("mapdomain %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
