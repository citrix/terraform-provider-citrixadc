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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccOnlinkipv6prefix_basic = `


	resource "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix" {
		ipv6prefix      = "8000::/64"
		onlinkprefix    = "YES"
		autonomusprefix = "NO"
	}
`
const testAccOnlinkipv6prefix_update = `


	resource "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix" {
		ipv6prefix      = "8000::/64"
		onlinkprefix    = "NO"
		autonomusprefix = "YES"
	}
`

func TestAccOnlinkipv6prefix_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOnlinkipv6prefixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOnlinkipv6prefix_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnlinkipv6prefixExist("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", nil),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", "onlinkprefix", "YES"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", "autonomusprefix", "NO"),
				),
			},
			{
				Config: testAccOnlinkipv6prefix_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnlinkipv6prefixExist("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", nil),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", "onlinkprefix", "NO"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", "autonomusprefix", "YES"),
				),
			},
		},
	})
}

func testAccCheckOnlinkipv6prefixExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No onlinkipv6prefix name is set")
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
		data, err := client.FindResource(service.Onlinkipv6prefix.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("onlinkipv6prefix %s not found", n)
		}

		return nil
	}
}

func testAccCheckOnlinkipv6prefixDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_onlinkipv6prefix" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Onlinkipv6prefix.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("onlinkipv6prefix %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccOnlinkipv6prefixDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOnlinkipv6prefixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOnlinkipv6prefixDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix_ds", "ipv6prefix", "9000::/64"),
					resource.TestCheckResourceAttr("data.citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix_ds", "onlinkprefix", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix_ds", "autonomusprefix", "NO"),
				),
			},
		},
	})
}

const testAccOnlinkipv6prefixDataSource_basic = `

resource "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix_ds" {
	ipv6prefix      = "9000::/64"
	onlinkprefix    = "YES"
	autonomusprefix = "NO"
}

data "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix_ds" {
	ipv6prefix = citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix_ds.ipv6prefix
	depends_on = [citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix_ds]
}
`
