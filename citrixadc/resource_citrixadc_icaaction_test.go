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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccIcaaction_basic = `

	resource "citrixadc_icaaction" "tf_icaaction" {
		name              = "my_ica_action"
		accessprofilename = "default_ica_accessprofile"
	}
  
`

const testAccIcaaction_update = `

	resource "citrixadc_icaaccessprofile" "tf_icaaccessprofile" {
		name                   = "my_profile"
		connectclientlptports  = "DEFAULT"
		localremotedatasharing = "DEFAULT"
	}

	resource "citrixadc_icaaction" "tf_icaaction" {
		name              = "my_ica_action"
		accessprofilename = citrixadc_icaaccessprofile.tf_icaaccessprofile.name
	}
  
`

func TestAccIcaaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIcaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaactionExist("citrixadc_icaaction.tf_icaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_icaaction.tf_icaaction", "name", "my_ica_action"),
					resource.TestCheckResourceAttr("citrixadc_icaaction.tf_icaaction", "accessprofilename", "default_ica_accessprofile"),
				),
			},
			{
				Config: testAccIcaaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaactionExist("citrixadc_icaaction.tf_icaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_icaaction.tf_icaaction", "name", "my_ica_action"),
					resource.TestCheckResourceAttr("citrixadc_icaaction.tf_icaaction", "accessprofilename", "my_profile"),
				),
			},
		},
	})
}

func testAccCheckIcaactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No icaaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("icaaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("icaaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckIcaactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_icaaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("icaaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("icaaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
