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

const testAccRdpserverprofile_basic = `

resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile" {
	name           = "my_rdpserverprofile"
	psk            = "key"
	rdpredirection = "ENABLE"
	rdpport        = 4000
	}
  
`

const testAccRdpserverprofile_update = `


resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile" {
	name           = "my_rdpserverprofile"
	psk            = "key"
	rdpredirection = "DISABLE"
	rdpport        = 4100
	}
  
`

func TestAccRdpserverprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRdpserverprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRdpserverprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpserverprofileExist("citrixadc_rdpserverprofile.tf_rdpserverprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "name", "my_rdpserverprofile"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "psk", "key"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "rdpredirection", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "rdpport", "4000"),
				),
			},
			{
				Config: testAccRdpserverprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpserverprofileExist("citrixadc_rdpserverprofile.tf_rdpserverprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "name", "my_rdpserverprofile"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "psk", "key"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "rdpredirection", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "rdpport", "4100"),
				),
			},
		},
	})
}

func testAccCheckRdpserverprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rdpserverprofile name is set")
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
		data, err := client.FindResource("rdpserverprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("rdpserverprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckRdpserverprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rdpserverprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("rdpserverprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("rdpserverprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
