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

const testAccAppfwgrpcwebtextcontenttype_basic_step1 = `
resource "citrixadc_appfwgrpcwebtextcontenttype" "tf_appfwgrpcwebtextcontenttype" {
  grpcwebtextcontenttypevalue = "application/grpc-web-text-test"
  isregex                     = "NOTREGEX"
}

`

const testAccAppfwgrpcwebtextcontenttype_basic_step2 = `
resource "citrixadc_appfwgrpcwebtextcontenttype" "tf_appfwgrpcwebtextcontenttype" {
  grpcwebtextcontenttypevalue = "application/grpc-web-text-test-updated"
  isregex                     = "NOTREGEX"
}

`

func TestAccAppfwgrpcwebtextcontenttype_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwgrpcwebtextcontenttypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwgrpcwebtextcontenttype_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwgrpcwebtextcontenttypeExist("citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype", "grpcwebtextcontenttypevalue", "application/grpc-web-text-test"),
					resource.TestCheckResourceAttr("citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype", "isregex", "NOTREGEX"),
				),
			},
			{
				Config: testAccAppfwgrpcwebtextcontenttype_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwgrpcwebtextcontenttypeExist("citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype", "grpcwebtextcontenttypevalue", "application/grpc-web-text-test-updated"),
					resource.TestCheckResourceAttr("citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype", "isregex", "NOTREGEX"),
				),
			},
		},
	})
}

func TestAccAppfwgrpcwebtextcontenttype_import(t *testing.T) {
	const resAddr = "citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwgrpcwebtextcontenttypeDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwgrpcwebtextcontenttype_basic_step1},
			{
				Config:                  testAccAppfwgrpcwebtextcontenttype_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckAppfwgrpcwebtextcontenttypeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwgrpcwebtextcontenttype name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Appfwgrpcwebtextcontenttype.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwgrpcwebtextcontenttype %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwgrpcwebtextcontenttypeDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwgrpcwebtextcontenttype" {
			continue
		}
		data, err := client.FindResource(service.Appfwgrpcwebtextcontenttype.Type(), rs.Primary.ID)
		if err == nil && data != nil {
			return fmt.Errorf("appfwgrpcwebtextcontenttype %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

const testAccAppfwgrpcwebtextcontenttypeDataSource_basic = `

resource "citrixadc_appfwgrpcwebtextcontenttype" "tf_appfwgrpcwebtextcontenttype" {
  grpcwebtextcontenttypevalue = "application/grpc-web-text-ds-test"
  isregex                     = "NOTREGEX"
}

data "citrixadc_appfwgrpcwebtextcontenttype" "tf_appfwgrpcwebtextcontenttype" {
  grpcwebtextcontenttypevalue = citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype.grpcwebtextcontenttypevalue
  depends_on                  = [citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype]
}
`

func TestAccAppfwgrpcwebtextcontenttypeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwgrpcwebtextcontenttypeDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype", "grpcwebtextcontenttypevalue", "application/grpc-web-text-ds-test"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwgrpcwebtextcontenttype.tf_appfwgrpcwebtextcontenttype", "isregex", "NOTREGEX"),
				),
			},
		},
	})
}
