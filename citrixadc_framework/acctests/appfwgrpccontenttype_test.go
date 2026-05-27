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

const testAccAppfwgrpccontenttype_basic_step1 = `
resource "citrixadc_appfwgrpccontenttype" "tf_appfwgrpccontenttype" {
  grpccontenttypevalue = "tf_acc_grpc_test"
  isregex              = "NOTREGEX"
}

`

const testAccAppfwgrpccontenttype_basic_step2 = `
resource "citrixadc_appfwgrpccontenttype" "tf_appfwgrpccontenttype" {
  grpccontenttypevalue = "tf_acc_grpc.*test_updated"
  isregex              = "REGEX"
}

`

func TestAccAppfwgrpccontenttype_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwgrpccontenttypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwgrpccontenttype_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwgrpccontenttypeExist("citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype", "grpccontenttypevalue", "tf_acc_grpc_test"),
					resource.TestCheckResourceAttr("citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype", "isregex", "NOTREGEX"),
				),
			},
			{
				Config: testAccAppfwgrpccontenttype_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwgrpccontenttypeExist("citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype", "grpccontenttypevalue", "tf_acc_grpc.*test_updated"),
					resource.TestCheckResourceAttr("citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype", "isregex", "REGEX"),
				),
			},
		},
	})
}

func testAccCheckAppfwgrpccontenttypeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwgrpccontenttype name is set")
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
		data, err := client.FindResource(service.Appfwgrpccontenttype.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwgrpccontenttype %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwgrpccontenttypeDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwgrpccontenttype" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwgrpccontenttype.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwgrpccontenttype %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwgrpccontenttypeDataSource_basic = `

resource "citrixadc_appfwgrpccontenttype" "tf_appfwgrpccontenttype" {
  grpccontenttypevalue = "tf_acc_grpc_test"
  isregex              = "NOTREGEX"
}

data "citrixadc_appfwgrpccontenttype" "tf_appfwgrpccontenttype" {
  grpccontenttypevalue = citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype.grpccontenttypevalue
  depends_on           = [citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype]
}
`

func TestAccAppfwgrpccontenttypeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwgrpccontenttypeDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype", "grpccontenttypevalue", "tf_acc_grpc_test"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwgrpccontenttype.tf_appfwgrpccontenttype", "isregex", "NOTREGEX"),
				),
			},
		},
	})
}
