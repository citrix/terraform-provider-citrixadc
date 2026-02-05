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

const testAccAppfwfieldtype_add = `
resource "citrixadc_appfwfieldtype" "tfAcc_appfwfieldtype" {
	name = "tfAcc_appfwfieldtype"
	regex = "test_.*regex"
	priority = "100"
}
`
const testAccAppfwfieldtype_update = `
resource "citrixadc_appfwfieldtype" "tfAcc_appfwfieldtype" {
	name = "tfAcc_appfwfieldtype"
	regex = "test_.*regex"
	priority = "30"
}
`

func TestAccAppfwfieldtype_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwfieldtypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwfieldtype_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwfieldtypeExist("citrixadc_appfwfieldtype.tfAcc_appfwfieldtype", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwfieldtype.tfAcc_appfwfieldtype", "name", "tfAcc_appfwfieldtype"),
					resource.TestCheckResourceAttr("citrixadc_appfwfieldtype.tfAcc_appfwfieldtype", "priority", "100"),
				),
			},
			{
				Config: testAccAppfwfieldtype_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwfieldtypeExist("citrixadc_appfwfieldtype.tfAcc_appfwfieldtype", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwfieldtype.tfAcc_appfwfieldtype", "name", "tfAcc_appfwfieldtype"),
					resource.TestCheckResourceAttr("citrixadc_appfwfieldtype.tfAcc_appfwfieldtype", "priority", "30"),
				),
			},
		},
	})
}

func testAccCheckAppfwfieldtypeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
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
		data, err := client.FindResource(service.Appfwfieldtype.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwfieldtypeDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwfieldtype" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwfieldtype.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwfieldtypeDataSource_basic = `
resource "citrixadc_appfwfieldtype" "tfAcc_appfwfieldtype" {
	name = "tfAcc_appfwfieldtype_ds"
	regex = "test_.*regex_ds"
	priority = "100"
	comment = "Test datasource comment"
}

data "citrixadc_appfwfieldtype" "tfAcc_appfwfieldtype" {
	name = citrixadc_appfwfieldtype.tfAcc_appfwfieldtype.name
}
`

func TestAccAppfwfieldtypeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwfieldtypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwfieldtypeDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwfieldtype.tfAcc_appfwfieldtype", "name", "tfAcc_appfwfieldtype_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwfieldtype.tfAcc_appfwfieldtype", "regex", "test_.*regex_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwfieldtype.tfAcc_appfwfieldtype", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwfieldtype.tfAcc_appfwfieldtype", "comment", "Test datasource comment"),
				),
			},
		},
	})
}
