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

const testAccGslbservicegroup_add = `


resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "test_gslbvservicegroup"
	servicetype      = "HTTP"
	cip              = "DISABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
  }
  resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
  }
  
`

const testAccGslbservicegroup_update = `


resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "test_gslbvservicegroup"
	servicetype      = "HTTP"
	cip              = "ENABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
  }
  resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
  }
  
`

func TestAccGslbservicegroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGslbservicegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroupExist("citrixadc_gslbservicegroup.tf_gslbservicegroup", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "servicegroupname", "test_gslbvservicegroup"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "cip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "sitename", "Site-Local"),
				),
			},
			{
				Config: testAccGslbservicegroup_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroupExist("citrixadc_gslbservicegroup.tf_gslbservicegroup", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "servicegroupname", "test_gslbvservicegroup"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "cip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "sitename", "Site-Local"),
				),
			},
		},
	})
}

func testAccCheckGslbservicegroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbservicegroup name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("gslbservicegroup", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("gslbservicegroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbservicegroupDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbservicegroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("gslbservicegroup", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbservicegroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
