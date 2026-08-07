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

const testAccAppfwurlencodedformcontenttype_basic = `
	resource "citrixadc_appfwurlencodedformcontenttype" "tf_urlencodedformcontenttype" {
		urlencodedformcontenttypevalue = "tf_urlencodedformcontenttype"
		isregex                        = "NOTREGEX"
	}
`

func TestAccAppfwurlencodedformcontenttype_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwurlencodedformcontenttypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwurlencodedformcontenttype_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwurlencodedformcontenttypeExist("citrixadc_appfwurlencodedformcontenttype.tf_urlencodedformcontenttype", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwurlencodedformcontenttype.tf_urlencodedformcontenttype", "urlencodedformcontenttypevalue", "tf_urlencodedformcontenttype"),
					resource.TestCheckResourceAttr("citrixadc_appfwurlencodedformcontenttype.tf_urlencodedformcontenttype", "isregex", "NOTREGEX"),
				),
			},
		},
	})
}

func testAccCheckAppfwurlencodedformcontenttypeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwurlencodedformcontenttype name is set")
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
		data, err := client.FindResource("appfwurlencodedformcontenttype", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwurlencodedformcontenttype %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwurlencodedformcontenttypeDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwurlencodedformcontenttype" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("appfwurlencodedformcontenttype", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwurlencodedformcontenttype %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAppfwurlencodedformcontenttype_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_appfwurlencodedformcontenttype.tf_urlencodedformcontenttype"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwurlencodedformcontenttypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwurlencodedformcontenttype_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwurlencodedformcontenttypeExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Appfwurlencodedformcontenttype.Type(), "tf_urlencodedformcontenttype"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAppfwurlencodedformcontenttype_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwurlencodedformcontenttypeExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAppfwurlencodedformcontenttype_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwurlencodedformcontenttypeDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAppfwurlencodedformcontenttype_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwurlencodedformcontenttypeExist("citrixadc_appfwurlencodedformcontenttype.tf_urlencodedformcontenttype", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppfwurlencodedformcontenttype_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckAppfwurlencodedformcontenttypeExist("citrixadc_appfwurlencodedformcontenttype.tf_urlencodedformcontenttype", nil)),
			},
		},
	})
}

func TestAccAppfwurlencodedformcontenttype_import(t *testing.T) {
	const resAddr = "citrixadc_appfwurlencodedformcontenttype.tf_urlencodedformcontenttype"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwurlencodedformcontenttypeDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwurlencodedformcontenttype_basic},
			{
				Config:                  testAccAppfwurlencodedformcontenttype_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccAppfwurlencodedformcontenttypeDataSource_basic = `
	resource "citrixadc_appfwurlencodedformcontenttype" "tf_urlencodedformcontenttype" {
		urlencodedformcontenttypevalue = "tf_urlencodedformcontenttype"
		isregex                        = "NOTREGEX"
	}

	data "citrixadc_appfwurlencodedformcontenttype" "tf_urlencodedformcontenttype" {
		urlencodedformcontenttypevalue = citrixadc_appfwurlencodedformcontenttype.tf_urlencodedformcontenttype.urlencodedformcontenttypevalue
	}
`

func TestAccAppfwurlencodedformcontenttypeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwurlencodedformcontenttypeDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwurlencodedformcontenttype.tf_urlencodedformcontenttype", "urlencodedformcontenttypevalue", "tf_urlencodedformcontenttype"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwurlencodedformcontenttype.tf_urlencodedformcontenttype", "isregex", "NOTREGEX"),
				),
			},
		},
	})
}
