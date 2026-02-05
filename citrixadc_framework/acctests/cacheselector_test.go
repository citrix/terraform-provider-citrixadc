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

const testAccCacheselector_basic = `
	resource "citrixadc_cacheselector" "tf_cacheselector" {
		selectorname = "my_cacheselector"
		rule = ["true"]
	}
`
const testAccCacheselector_update = `
	resource "citrixadc_cacheselector" "tf_cacheselector" {
		selectorname = "my_cacheselector2"
		rule = ["false"]
	}
`

func TestAccCacheselector_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCacheselectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCacheselector_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCacheselectorExist("citrixadc_cacheselector.tf_cacheselector", nil),
					resource.TestCheckResourceAttr("citrixadc_cacheselector.tf_cacheselector", "selectorname", "my_cacheselector"),
				),
			},
			{
				Config: testAccCacheselector_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCacheselectorExist("citrixadc_cacheselector.tf_cacheselector", nil),
					resource.TestCheckResourceAttr("citrixadc_cacheselector.tf_cacheselector", "selectorname", "my_cacheselector2"),
				),
			},
		},
	})
}

func testAccCheckCacheselectorExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cacheselector name is set")
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
		data, err := client.FindResource(service.Cacheselector.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cacheselector %s not found", n)
		}

		return nil
	}
}

func testAccCheckCacheselectorDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cacheselector" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cacheselector.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cacheselector %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCacheselectorDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCacheselectorDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cacheselector.tf_cacheselector_ds", "selectorname", "tf_cacheselector_ds"),
				),
			},
		},
	})
}

const testAccCacheselectorDataSource_basic = `

resource "citrixadc_cacheselector" "tf_cacheselector_ds" {
    selectorname = "tf_cacheselector_ds"
    rule         = ["true"]
}

data "citrixadc_cacheselector" "tf_cacheselector_ds" {
    selectorname = citrixadc_cacheselector.tf_cacheselector_ds.selectorname
    depends_on   = [citrixadc_cacheselector.tf_cacheselector_ds]
}

`
