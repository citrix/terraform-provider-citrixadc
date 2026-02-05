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

const testAccSslparameter_basic = `
	resource "citrixadc_sslparameter" "default" {
		denysslreneg   = "NONSECURE"
		defaultprofile = "ENABLED"
		operationqueuelimit = 4096
	}
`
const testAccSslparameter_basic_update = `
	resource "citrixadc_sslparameter" "default" {
		denysslreneg   = "ALL"
		defaultprofile = "ENABLED"
		operationqueuelimit = 4088
	}
`

const testAccSslparameterDataSource_basic = `
	resource "citrixadc_sslparameter" "default" {
		denysslreneg   = "NONSECURE"
		defaultprofile = "ENABLED"
		operationqueuelimit = 4096
	}

	data "citrixadc_sslparameter" "default" {
		depends_on = [citrixadc_sslparameter.default]
	}
`

func TestAccSslparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// sslparameter resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSslparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslparameterExist("citrixadc_sslparameter.default", nil),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "denysslreneg", "NONSECURE"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "defaultprofile", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "operationqueuelimit", "4096"),
				),
			},
			{
				Config: testAccSslparameter_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslparameterExist("citrixadc_sslparameter.default", nil),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "denysslreneg", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "defaultprofile", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "operationqueuelimit", "4088"),
				),
			},
		},
	})
}

func testAccCheckSslparameterExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := client.FindResource(service.Sslparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("SSL Parameter %s not found", n)
		}

		return nil
	}
}

func TestAccSslparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSslparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslparameter.default", "denysslreneg", "NONSECURE"),
					resource.TestCheckResourceAttr("data.citrixadc_sslparameter.default", "defaultprofile", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslparameter.default", "operationqueuelimit", "4096"),
				),
			},
		},
	})
}
