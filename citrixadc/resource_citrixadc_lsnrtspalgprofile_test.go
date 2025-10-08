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

const testAccLsnrtspalgprofile_basic = `

	resource "citrixadc_lsnrtspalgprofile" "tf_lsnrtspalgprofile" {
		rtspalgprofilename = "my_lsn_rtspalgprofile"
		rtspportrange      = 4200
		rtspidletimeout    = 150
	}
`
const testAccLsnrtspalgprofile_update = `

	resource "citrixadc_lsnrtspalgprofile" "tf_lsnrtspalgprofile" {
		rtspalgprofilename = "my_lsn_rtspalgprofile"
		rtspportrange      = 4500
		rtspidletimeout    = 100
	}
`

func TestAccLsnrtspalgprofile_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this LSN resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLsnrtspalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnrtspalgprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnrtspalgprofileExist("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspalgprofilename", "my_lsn_rtspalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspportrange", "4200"),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspidletimeout", "150"),
				),
			},
			{
				Config: testAccLsnrtspalgprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnrtspalgprofileExist("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspalgprofilename", "my_lsn_rtspalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspportrange", "4500"),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspidletimeout", "100"),
				),
			},
		},
	})
}

func testAccCheckLsnrtspalgprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnrtspalgprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource("lsnrtspalgprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnrtspalgprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnrtspalgprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnrtspalgprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnrtspalgprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnrtspalgprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
