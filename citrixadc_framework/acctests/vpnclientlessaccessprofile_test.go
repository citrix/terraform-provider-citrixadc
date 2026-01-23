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

const testAccVpnclientlessaccessprofile_basic = `

	resource "citrixadc_vpnclientlessaccessprofile" "tf_vpnclientlessaccessprofile" {
		profilename = "tf_vpnclientlessaccessprofile"
		requirepersistentcookie = "ON"
	}
`

const testAccVpnclientlessaccessprofile_basic_update = `

	resource "citrixadc_vpnclientlessaccessprofile" "tf_vpnclientlessaccessprofile" {
		profilename = "tf_vpnclientlessaccessprofile"
		requirepersistentcookie = "OFF"
		regexforfindingurlinjavascript = citrixadc_policypatset.tf_patset.name
		regexforfindingurlincss = citrixadc_policypatset.tf_patset.name
		regexforfindingurlinxcomponent = citrixadc_policypatset.tf_patset.name
		regexforfindingurlinxml = citrixadc_policypatset.tf_patset.name
		regexforfindingcustomurls = citrixadc_policypatset.tf_patset.name
	}

	resource "citrixadc_policypatset" "tf_patset" {
		name = "tf_patset"
	}
`

func TestAccVpnclientlessaccessprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnclientlessaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnclientlessaccessprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnclientlessaccessprofileExist("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", "profilename", "tf_vpnclientlessaccessprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", "requirepersistentcookie", "ON"),
				),
			},
			{
				Config: testAccVpnclientlessaccessprofile_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnclientlessaccessprofileExist("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", "requirepersistentcookie", "OFF"),
				),
			},
		},
	})
}

func testAccCheckVpnclientlessaccessprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnclientlessaccessprofile name is set")
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
		data, err := client.FindResource(service.Vpnclientlessaccessprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnclientlessaccessprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnclientlessaccessprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnclientlessaccessprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnclientlessaccessprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnclientlessaccessprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
