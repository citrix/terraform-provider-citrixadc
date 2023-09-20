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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccCachecontentgroup_basic = `

	resource "citrixadc_cachecontentgroup" "tf_cachecontentgroup" {
		name                 = "my_cachecontentgroup"
		heurexpiryparam      = 30
		prefetch             = "YES"
		quickabortsize       = 40
		ignorereqcachinghdrs = "YES"
	}
`
const testAccCachecontentgroup_update = `

	resource "citrixadc_cachecontentgroup" "tf_cachecontentgroup" {
		name                 = "my_cachecontentgroup"
		heurexpiryparam      = 50
		prefetch             = "NO"
		quickabortsize       = 50
		ignorereqcachinghdrs = "NO"
	}
`

func TestAccCachecontentgroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCachecontentgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCachecontentgroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachecontentgroupExist("citrixadc_cachecontentgroup.tf_cachecontentgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "name", "my_cachecontentgroup"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "heurexpiryparam", "30"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "prefetch", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "quickabortsize", "40"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "ignorereqcachinghdrs", "YES"),
				),
			},
			{
				Config: testAccCachecontentgroup_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachecontentgroupExist("citrixadc_cachecontentgroup.tf_cachecontentgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "name", "my_cachecontentgroup"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "heurexpiryparam", "50"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "prefetch", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "quickabortsize", "50"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "ignorereqcachinghdrs", "NO"),
				),
			},
		},
	})
}

func testAccCheckCachecontentgroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cachecontentgroup name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Cachecontentgroup.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cachecontentgroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckCachecontentgroupDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cachecontentgroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Cachecontentgroup.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cachecontentgroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
