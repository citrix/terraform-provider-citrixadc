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

const testAccIcaaccessprofile_basic = `


	resource "citrixadc_icaaccessprofile" "tf_icaaccessprofile" {
		name                   = "my_ica_accessprofile"
		connectclientlptports  = "DEFAULT"
		localremotedatasharing = "DEFAULT"
		wiaredirection		 = "DISABLED"
		smartcardredirection	= "DISABLED"
		fido2redirection		 = "DISABLED"
		draganddrop			 = "DISABLED"
		clienttwaindeviceredirection = "DISABLED"
	}
	
`
const testAccIcaaccessprofile_update = `


	resource "citrixadc_icaaccessprofile" "tf_icaaccessprofile" {
		name                   = "my_ica_accessprofile"
		connectclientlptports  = "DISABLED"
		localremotedatasharing = "DISABLED"
		wiaredirection		 = "DEFAULT"
		smartcardredirection	= "DEFAULT"
		fido2redirection		 = "DEFAULT"
		draganddrop			 = "DEFAULT"
		clienttwaindeviceredirection = "DEFAULT"
	}
	
`

func TestAccIcaaccessprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcaaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaaccessprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaaccessprofileExist("citrixadc_icaaccessprofile.tf_icaaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "name", "my_ica_accessprofile"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "connectclientlptports", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "localremotedatasharing", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "wiaredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "smartcardredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "fido2redirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "draganddrop", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "clienttwaindeviceredirection", "DISABLED"),
				),
			},
			{
				Config: testAccIcaaccessprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaaccessprofileExist("citrixadc_icaaccessprofile.tf_icaaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "name", "my_ica_accessprofile"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "connectclientlptports", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "localremotedatasharing", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "wiaredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "smartcardredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "fido2redirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "draganddrop", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "clienttwaindeviceredirection", "DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckIcaaccessprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No icaaccessprofile name is set")
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
		data, err := client.FindResource("icaaccessprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("icaaccessprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckIcaaccessprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_icaaccessprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("icaaccessprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("icaaccessprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
