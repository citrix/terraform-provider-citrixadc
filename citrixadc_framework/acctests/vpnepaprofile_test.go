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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Note: vpnepaprofile has no in-place update support; name, filename and data are
// all RequiresReplace. The NITRO "vpnepaprofile" command is deprecated but still
// functional. The basic test therefore performs a single create + verify step.

const testAccVpnepaprofile_basic_step1 = `
resource "citrixadc_vpnepaprofile" "tf_vpnepaprofile" {
  name     = "tf_vpnepaprofile"
  filename = "tf_vpnepaprofile.xml"
}

`

func TestAccVpnepaprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnepaprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnepaprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnepaprofileExist("citrixadc_vpnepaprofile.tf_vpnepaprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnepaprofile.tf_vpnepaprofile", "name", "tf_vpnepaprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnepaprofile.tf_vpnepaprofile", "filename", "tf_vpnepaprofile.xml"),
				),
			},
		},
	})
}

func testAccCheckVpnepaprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnepaprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnepaprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnepaprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnepaprofileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnepaprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnepaprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnepaprofile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccVpnepaprofileDataSource_basic = `

resource "citrixadc_vpnepaprofile" "tf_vpnepaprofile" {
  name     = "tf_vpnepaprofile"
  filename = "tf_vpnepaprofile.xml"
}

data "citrixadc_vpnepaprofile" "tf_vpnepaprofile" {
  name       = citrixadc_vpnepaprofile.tf_vpnepaprofile.name
  depends_on = [citrixadc_vpnepaprofile.tf_vpnepaprofile]
}
`

func TestAccVpnepaprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnepaprofileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnepaprofile.tf_vpnepaprofile", "name", "tf_vpnepaprofile"),
				),
			},
		},
	})
}
