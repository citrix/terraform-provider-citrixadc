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

func TestAccLbprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLbprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbprofileExist("citrixadc_lbprofile.tf_lbprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbprofilename", "tf_lbprofile"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "dbslb", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "processlocal", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "httponlycookieflag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbhashfingers", "258"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbhashalgorithm", "PRAC"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "storemqttclientidandusername", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "proximityfromself", "NO"),
					testAccCheckUserAgent(),
				),
			},
			{
				Config: testAccLbprofile_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbprofileExist("citrixadc_lbprofile.tf_lbprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbprofilename", "tf_lbprofile"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "dbslb", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "processlocal", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "httponlycookieflag", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbhashfingers", "255"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbhashalgorithm", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "storemqttclientidandusername", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "proximityfromself", "YES"),
					testAccCheckUserAgent(),
				),
			},
		},
	})
}

func testAccCheckLbprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbprofile name is set")
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
		data, err := client.FindResource("lbprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Lbprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_Lbprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Lbprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbprofile_basic = `
resource "citrixadc_lbprofile" "tf_lbprofile" {
    lbprofilename = "tf_lbprofile"
    dbslb = "ENABLED"
	processlocal = "DISABLED"
	httponlycookieflag = "ENABLED"
	lbhashfingers = 258
	lbhashalgorithm = "PRAC"
	storemqttclientidandusername = "YES"
	proximityfromself = "NO"
}

`

const testAccLbprofile_basic_update = `

resource "citrixadc_lbprofile" "tf_lbprofile" {
    lbprofilename = "tf_lbprofile"
	dbslb = "DISABLED"
	processlocal = "ENABLED"
	httponlycookieflag = "DISABLED"
	lbhashfingers = 255
	lbhashalgorithm = "DEFAULT"
	storemqttclientidandusername = "NO"
	proximityfromself = "YES"
    
}

`
