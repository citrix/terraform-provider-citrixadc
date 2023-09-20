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

const testAccPcpprofile_basic = `


resource "citrixadc_pcpprofile" "tf_pcpprofile" {
	name               = "my_pcpprofile"
	mapping            = "ENABLED"
	peer               = "ENABLED"
  }
  
`
const testAccPcpprofile_update = `


resource "citrixadc_pcpprofile" "tf_pcpprofile" {
	name               = "my_pcpprofile"
	mapping            = "DISABLED"
	peer               = "DISABLED"
  }
  
`

func TestAccPcpprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPcpprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPcpprofileExist("citrixadc_pcpprofile.tf_pcpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "name", "my_pcpprofile"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "mapping", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "peer", "ENABLED"),
				),
			},
			{
				Config: testAccPcpprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPcpprofileExist("citrixadc_pcpprofile.tf_pcpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "name", "my_pcpprofile"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "mapping", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "peer", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckPcpprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No pcpprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("pcpprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("pcpprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckPcpprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_pcpprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("pcpprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("pcpprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
