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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLsnhttphdrlogprofile_basic = `


resource "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile" {
	httphdrlogprofilename = "my_lsn_httphdrlogprofile"
	logurl                = "DISABLED"
	logversion            = "DISABLED"
	loghost               = "DISABLED"
	}
  
`
const testAccLsnhttphdrlogprofile_update = `


resource "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile" {
	httphdrlogprofilename = "my_lsn_httphdrlogprofile"
	logurl                = "ENABLED"
	logversion            = "ENABLED"
	loghost               = "ENABLED"
	}
  
`

func TestAccLsnhttphdrlogprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnhttphdrlogprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnhttphdrlogprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnhttphdrlogprofileExist("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "httphdrlogprofilename", "my_lsn_httphdrlogprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "logurl", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "logversion", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "loghost", "DISABLED"),
				),
			},
			{
				Config: testAccLsnhttphdrlogprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnhttphdrlogprofileExist("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "httphdrlogprofilename", "my_lsn_httphdrlogprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "logurl", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "logversion", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "loghost", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckLsnhttphdrlogprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnhttphdrlogprofile name is set")
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
		data, err := client.FindResource("lsnhttphdrlogprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnhttphdrlogprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnhttphdrlogprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnhttphdrlogprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnhttphdrlogprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnhttphdrlogprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsnhttphdrlogprofileDataSource_basic = `

resource "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile_ds" {
	httphdrlogprofilename = "my_lsn_httphdrlogprofile_ds"
	logurl                = "DISABLED"
	logversion            = "DISABLED"
	loghost               = "DISABLED"
}

data "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile_ds" {
	httphdrlogprofilename = citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile_ds.httphdrlogprofilename
}
`

func TestAccLsnhttphdrlogprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnhttphdrlogprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile_ds", "httphdrlogprofilename", "my_lsn_httphdrlogprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile_ds", "logurl", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile_ds", "logversion", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile_ds", "loghost", "DISABLED"),
				),
			},
		},
	})
}
