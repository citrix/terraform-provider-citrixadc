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

const testAccAppfwxmlcontenttype_basic = `
	resource "citrixadc_appfwxmlcontenttype" "tf_Acc_appfwxmlcontenttype" {
		xmlcontenttypevalue = "tf_Acc.*test"
		isregex = "REGEX"
	}
`

func TestAccAppfwxmlcontenttype_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwxmlcontenttypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwxmlcontenttype_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwxmlcontenttypeExist("citrixadc_appfwxmlcontenttype.tf_Acc_appfwxmlcontenttype", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwxmlcontenttype.tf_Acc_appfwxmlcontenttype", "xmlcontenttypevalue", "tf_Acc.*test"),
				),
			},
		},
	})
}

func testAccCheckAppfwxmlcontenttypeExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := client.FindResource(service.Appfwxmlcontenttype.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwxmlcontenttypeDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwxmlcontenttype" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwxmlcontenttype.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwxmlcontenttypeDataSource_basic = `

	resource "citrixadc_appfwxmlcontenttype" "tf_Acc_appfwxmlcontenttype" {
		xmlcontenttypevalue = "tf_Acc.*test"
		isregex = "REGEX"
	}

	data "citrixadc_appfwxmlcontenttype" "tf_Acc_appfwxmlcontenttype" {
		xmlcontenttypevalue = citrixadc_appfwxmlcontenttype.tf_Acc_appfwxmlcontenttype.xmlcontenttypevalue
		depends_on = [citrixadc_appfwxmlcontenttype.tf_Acc_appfwxmlcontenttype]
	}
`

func TestAccAppfwxmlcontenttypeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwxmlcontenttypeDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwxmlcontenttype.tf_Acc_appfwxmlcontenttype", "xmlcontenttypevalue", "tf_Acc.*test"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwxmlcontenttype.tf_Acc_appfwxmlcontenttype", "isregex", "REGEX"),
				),
			},
		},
	})
}
