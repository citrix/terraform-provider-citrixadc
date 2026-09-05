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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccAppfwmultipartformcontenttype_basic = `
	resource "citrixadc_appfwmultipartformcontenttype" "tf_multipartform" {
		multipartformcontenttypevalue = "date/tf_multipartform"
		isregex                       = "REGEX"
	}
`

const testAccAppfwmultipartformcontenttypeDataSource_basic = `
	resource "citrixadc_appfwmultipartformcontenttype" "tf_multipartform" {
		multipartformcontenttypevalue = "date/tf_multipartform"
		isregex                       = "REGEX"
	}

	data "citrixadc_appfwmultipartformcontenttype" "tf_multipartform" {
		multipartformcontenttypevalue = citrixadc_appfwmultipartformcontenttype.tf_multipartform.multipartformcontenttypevalue
		depends_on = [citrixadc_appfwmultipartformcontenttype.tf_multipartform]
	}
`

func TestAccAppfwmultipartformcontenttype_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwmultipartformcontenttypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwmultipartformcontenttype_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwmultipartformcontenttypeExist("citrixadc_appfwmultipartformcontenttype.tf_multipartform", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwmultipartformcontenttype.tf_multipartform", "multipartformcontenttypevalue", "date/tf_multipartform"),
					resource.TestCheckResourceAttr("citrixadc_appfwmultipartformcontenttype.tf_multipartform", "isregex", "REGEX"),
				),
			},
		},
	})
}

func testAccCheckAppfwmultipartformcontenttypeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwmultipartformcontenttype name is set")
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
		appfwmultipartformcontenttypeName := rs.Primary.ID
		data, err := client.FindResource("appfwmultipartformcontenttype", appfwmultipartformcontenttypeName)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwmultipartformcontenttype %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwmultipartformcontenttypeDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwmultipartformcontenttype" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		appfwmultipartformcontenttypeName := rs.Primary.ID
		_, err := client.FindResource("appfwmultipartformcontenttype", appfwmultipartformcontenttypeName)

		if err == nil {
			return fmt.Errorf("appfwmultipartformcontenttype %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAppfwmultipartformcontenttype_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_appfwmultipartformcontenttype.tf_multipartform"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwmultipartformcontenttypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwmultipartformcontenttype_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwmultipartformcontenttypeExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Appfwmultipartformcontenttype.Type(), "date/tf_multipartform"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAppfwmultipartformcontenttype_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwmultipartformcontenttypeExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAppfwmultipartformcontenttype_import(t *testing.T) {
	const resAddr = "citrixadc_appfwmultipartformcontenttype.tf_multipartform"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwmultipartformcontenttypeDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwmultipartformcontenttype_basic},
			{
				Config:                  testAccAppfwmultipartformcontenttype_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAppfwmultipartformcontenttype_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwmultipartformcontenttypeDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAppfwmultipartformcontenttype_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwmultipartformcontenttypeExist("citrixadc_appfwmultipartformcontenttype.tf_multipartform", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAppfwmultipartformcontenttype_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwmultipartformcontenttypeExist("citrixadc_appfwmultipartformcontenttype.tf_multipartform", nil)),
			},
		},
	})
}

func TestAccAppfwmultipartformcontenttypeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwmultipartformcontenttypeDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwmultipartformcontenttype.tf_multipartform", "multipartformcontenttypevalue", "date/tf_multipartform"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwmultipartformcontenttype.tf_multipartform", "isregex", "REGEX"),
					// id is the universal runtime-binding proof for the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_appfwmultipartformcontenttype.tf_multipartform", "id"),
				),
			},
		},
	})
}
