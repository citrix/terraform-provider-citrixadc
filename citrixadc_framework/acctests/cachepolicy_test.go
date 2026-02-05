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

const testAccCachepolicy_basic = `


resource "citrixadc_cachepolicy" "tf_cachepolicy" {
	policyname  = "my_cachepolicy"
	rule        = "true"
	action      = "CACHE"
	undefaction = "NOCACHE"
	}
  
`
const testAccCachepolicy_update = `

	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "my_cachepolicy"
		rule        = "true"
		action      = "MAY_CACHE"
		undefaction = "RESET"
	}
  
`

func TestAccCachepolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCachepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCachepolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicyExist("citrixadc_cachepolicy.tf_cachepolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "policyname", "my_cachepolicy"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "action", "CACHE"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "undefaction", "NOCACHE"),
				),
			},
			{
				Config: testAccCachepolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicyExist("citrixadc_cachepolicy.tf_cachepolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "policyname", "my_cachepolicy"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "action", "MAY_CACHE"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "undefaction", "RESET"),
				),
			},
		},
	})
}

func testAccCheckCachepolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cachepolicy name is set")
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
		data, err := client.FindResource(service.Cachepolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cachepolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckCachepolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cachepolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cachepolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cachepolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCachepolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCachepolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cachepolicy.tf_cachepolicy_ds", "policyname", "tf_cachepolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_cachepolicy.tf_cachepolicy_ds", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_cachepolicy.tf_cachepolicy_ds", "action", "CACHE"),
					resource.TestCheckResourceAttr("data.citrixadc_cachepolicy.tf_cachepolicy_ds", "undefaction", "NOCACHE"),
				),
			},
		},
	})
}

const testAccCachepolicyDataSource_basic = `

resource "citrixadc_cachepolicy" "tf_cachepolicy_ds" {
    policyname  = "tf_cachepolicy_ds"
    rule        = "true"
    action      = "CACHE"
    undefaction = "NOCACHE"
}

data "citrixadc_cachepolicy" "tf_cachepolicy_ds" {
    policyname = citrixadc_cachepolicy.tf_cachepolicy_ds.policyname
    depends_on = [citrixadc_cachepolicy.tf_cachepolicy_ds]
}

`
