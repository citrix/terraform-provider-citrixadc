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

const testAccLsnip6profile_basic = `

	resource "citrixadc_nsip6" "tf_nsip6" {
		ipv6address = "2003::1/64"
		type       = "SNIP"
	}

	resource "citrixadc_lsnip6profile" "tf_lsnaip6profile" {
		name     = "my_lsn_ip6profile"
		type     = "DS-Lite"
		network6 = "2003::1/64"
		depends_on = [citrixadc_nsip6.tf_nsip6]
	}
  
`
const testAccLsnip6profile_update = `

	resource "citrixadc_nsip6" "tf_nsip6" {
		ipv6address = "1001::/64"
		type       = "SNIP"
	}

	resource "citrixadc_lsnip6profile" "tf_lsnaip6profile" {
		name     = "my_lsn_ip6profile"
		type     = "DS-Lite"
		network6 = "1001::/64"
		depends_on = [citrixadc_nsip6.tf_nsip6]
	}
  
`

func TestAccLsnip6profile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnip6profileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnip6profile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnip6profileExist("citrixadc_lsnip6profile.tf_lsnaip6profile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnip6profile.tf_lsnaip6profile", "name", "my_lsn_ip6profile"),
					resource.TestCheckResourceAttr("citrixadc_lsnip6profile.tf_lsnaip6profile", "type", "DS-Lite"),
					resource.TestCheckResourceAttr("citrixadc_lsnip6profile.tf_lsnaip6profile", "network6", "2003::1/64"),
				),
			},
			{
				Config: testAccLsnip6profile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnip6profileExist("citrixadc_lsnip6profile.tf_lsnaip6profile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnip6profile.tf_lsnaip6profile", "name", "my_lsn_ip6profile"),
					resource.TestCheckResourceAttr("citrixadc_lsnip6profile.tf_lsnaip6profile", "type", "DS-Lite"),
					resource.TestCheckResourceAttr("citrixadc_lsnip6profile.tf_lsnaip6profile", "network6", "1001::/64"),
				),
			},
		},
	})
}

func testAccCheckLsnip6profileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnip6profile name is set")
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
		data, err := client.FindResource("lsnip6profile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnip6profile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnip6profileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnip6profile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnip6profile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnip6profile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsnip6profileDataSource_basic = `

	resource "citrixadc_nsip6" "tf_nsip6" {
		ipv6address = "2003::1/64"
		type       = "SNIP"
	}

	resource "citrixadc_lsnip6profile" "tf_lsnaip6profile_ds" {
		name     = "my_lsn_ip6profile_ds"
		type     = "DS-Lite"
		network6 = "2003::1/64"
		depends_on = [citrixadc_nsip6.tf_nsip6]
	}

	data "citrixadc_lsnip6profile" "tf_lsnaip6profile_ds" {
		name = citrixadc_lsnip6profile.tf_lsnaip6profile_ds.name
	}
`

func TestAccLsnip6profileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnip6profileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnip6profile.tf_lsnaip6profile_ds", "name", "my_lsn_ip6profile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnip6profile.tf_lsnaip6profile_ds", "type", "DS-Lite"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnip6profile.tf_lsnaip6profile_ds", "network6", "2003::1/64"),
				),
			},
		},
	})
}
