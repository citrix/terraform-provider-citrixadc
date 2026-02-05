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

const testAccSnmpcommunity_basic = `

resource "citrixadc_snmpcommunity" "tf_snmpcommunity" {
	communityname = "test_community"
	permissions   = "GET"
	}
`
const testAccSnmpcommunity_update = `

resource "citrixadc_snmpcommunity" "tf_snmpcommunity" {
	communityname = "test_community"
	permissions   = "GET_BULK"
	}
`

func TestAccSnmpcommunity_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpcommunityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpcommunity_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpcommunityExist("citrixadc_snmpcommunity.tf_snmpcommunity", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpcommunity.tf_snmpcommunity", "communityname", "test_community"),
					resource.TestCheckResourceAttr("citrixadc_snmpcommunity.tf_snmpcommunity", "permissions", "GET"),
				),
			},
			{
				Config: testAccSnmpcommunity_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpcommunityExist("citrixadc_snmpcommunity.tf_snmpcommunity", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpcommunity.tf_snmpcommunity", "communityname", "test_community"),
					resource.TestCheckResourceAttr("citrixadc_snmpcommunity.tf_snmpcommunity", "permissions", "GET_BULK"),
				),
			},
		},
	})
}

func testAccCheckSnmpcommunityExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpcommunity name is set")
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
		data, err := client.FindResource(service.Snmpcommunity.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("snmpcommunity %s not found", n)
		}

		return nil
	}
}

func testAccCheckSnmpcommunityDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_snmpcommunity" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Snmpcommunity.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("snmpcommunity %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSnmpcommunityDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpcommunityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpcommunityDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_snmpcommunity.tf_snmpcommunity_ds", "communityname", "tf_community_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_snmpcommunity.tf_snmpcommunity_ds", "permissions", "GET"),
				),
			},
		},
	})
}

const testAccSnmpcommunityDataSource_basic = `

resource "citrixadc_snmpcommunity" "tf_snmpcommunity_ds" {
	communityname = "tf_community_ds"
	permissions   = "GET"
}

data "citrixadc_snmpcommunity" "tf_snmpcommunity_ds" {
	communityname = citrixadc_snmpcommunity.tf_snmpcommunity_ds.communityname
	depends_on = [citrixadc_snmpcommunity.tf_snmpcommunity_ds]
}
`
