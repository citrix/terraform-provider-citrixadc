/*
Copyright 2021 Citrix Systems, Inc

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

const testAccLbgroup_basic = `
# The cookiedomain, rule and usevserverpersistency properties variabled cannot
# be updated and so they were deliberately left out of the test suite.

resource "citrixadc_lbgroup" "tf_lbgroup" {
	name = "tf_lbgroup"
	persistencetype = "RULE"
	persistencebackup = "SOURCEIP"
	backuppersistencetimeout = 10.0
	persistmask = "255.255.255.0"
	v6persistmasklen = 64
	timeout = 10.0
}
`

const testAccLbgroup_update_properties = `
resource "citrixadc_lbgroup" "tf_lbgroup" {
	name = "tf_lbgroup"
	persistencetype = "COOKIEINSERT"
	persistencebackup = "SOURCEIP"
	backuppersistencetimeout = 15.0
	persistmask = "255.255.254.0"
	cookiename = "tf_cookie_1"
	v6persistmasklen = 96
	timeout = 15.0
}
`

const testAccLbgroupDataSource_basic = `
resource "citrixadc_lbgroup" "tf_lbgroup" {
	name = "tf_lbgroup_ds"
	persistencetype = "COOKIEINSERT"
	persistencebackup = "SOURCEIP"
	backuppersistencetimeout = 10
	persistmask = "255.255.255.0"
	cookiename = "test_cookie"
	v6persistmasklen = 64
	timeout = 10
}

data "citrixadc_lbgroup" "tf_lbgroup" {
	name = citrixadc_lbgroup.tf_lbgroup.name
}
`

func TestAccLbgroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbgroupDestroy,
		Steps: []resource.TestStep{
			// create Lbgroup
			{
				Config: testAccLbgroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroupExist("citrixadc_lbgroup.tf_lbgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistencetype", "RULE"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistencebackup", "SOURCEIP"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "backuppersistencetimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "v6persistmasklen", "64"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "timeout", "10"),
					testAccCheckUserAgent(),
				),
			},
			// update Lbgroup properties
			{
				Config: testAccLbgroup_update_properties,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroupExist("citrixadc_lbgroup.tf_lbgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistencetype", "COOKIEINSERT"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistencebackup", "SOURCEIP"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "backuppersistencetimeout", "15"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistmask", "255.255.254.0"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "cookiename", "tf_cookie_1"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "v6persistmasklen", "96"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "timeout", "15"),
					testAccCheckUserAgent(),
				),
			},
		},
	})
}

func TestAccLbgroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbgroupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "name", "tf_lbgroup_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "persistencetype", "COOKIEINSERT"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "persistencebackup", "SOURCEIP"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "backuppersistencetimeout", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "persistmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "cookiename", "test_cookie"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "v6persistmasklen", "64"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "timeout", "10"),
				),
			},
		},
	})
}

func testAccCheckLbgroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Lbgroup name is set")
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
		data, err := client.FindResource("lbgroup", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lbgroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbgroupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbgroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbgroup", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbgroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
