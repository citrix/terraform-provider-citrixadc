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

const testAccLsnlogprofile_basic = `

	resource "citrixadc_lsnlogprofile" "tf_lsnlogprofile" {
		logprofilename = "my_lsn_logprofile"
		logsubscrinfo   = "ENABLED"
		logcompact      = "ENABLED"
		logipfix        = "ENABLED"
	}
	
`
const testAccLsnlogprofile_update = `

	resource "citrixadc_lsnlogprofile" "tf_lsnlogprofile" {
		logprofilename = "my_lsn_logprofile"
		logsubscrinfo   = "DISABLED"
		logcompact      = "DISABLED"
		logipfix        = "DISABLED"
	}
	
`

func TestAccLsnlogprofile_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this LSN resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnlogprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnlogprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnlogprofileExist("citrixadc_lsnlogprofile.tf_lsnlogprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logprofilename", "my_lsn_logprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logsubscrinfo", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logcompact", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logipfix", "ENABLED"),
				),
			},
			{
				Config: testAccLsnlogprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnlogprofileExist("citrixadc_lsnlogprofile.tf_lsnlogprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logprofilename", "my_lsn_logprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logsubscrinfo", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logcompact", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logipfix", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckLsnlogprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnlogprofile name is set")
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
		data, err := client.FindResource("lsnlogprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnlogprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnlogprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnlogprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnlogprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnlogprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
