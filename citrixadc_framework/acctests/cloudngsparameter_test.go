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

const testAccCloudngsparameter_basic_step1 = `

	resource "citrixadc_cloudngsparameter" "tf_cloudngsparameter" {
		blockonallowedngstktprof   = "YES"
		allowedudtversion          = "V5"
		csvserverticketingdecouple = "NO"
		allowdtls12                = "NO"
	}
`

const testAccCloudngsparameter_basic_step2 = `

	resource "citrixadc_cloudngsparameter" "tf_cloudngsparameter" {
		blockonallowedngstktprof   = "NO"
		allowedudtversion          = "V6"
		csvserverticketingdecouple = "YES"
		allowdtls12                = "YES"
	}
`

func TestAccCloudngsparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: it always exists on the ADC and is never deleted,
		// so no CheckDestroy is used.
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudngsparameter_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudngsparameterExist("citrixadc_cloudngsparameter.tf_cloudngsparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudngsparameter.tf_cloudngsparameter", "blockonallowedngstktprof", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cloudngsparameter.tf_cloudngsparameter", "allowedudtversion", "V5"),
					resource.TestCheckResourceAttr("citrixadc_cloudngsparameter.tf_cloudngsparameter", "csvserverticketingdecouple", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cloudngsparameter.tf_cloudngsparameter", "allowdtls12", "NO"),
				),
			},
			{
				Config: testAccCloudngsparameter_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudngsparameterExist("citrixadc_cloudngsparameter.tf_cloudngsparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudngsparameter.tf_cloudngsparameter", "blockonallowedngstktprof", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cloudngsparameter.tf_cloudngsparameter", "allowedudtversion", "V6"),
					resource.TestCheckResourceAttr("citrixadc_cloudngsparameter.tf_cloudngsparameter", "csvserverticketingdecouple", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cloudngsparameter.tf_cloudngsparameter", "allowdtls12", "YES"),
				),
			},
		},
	})
}

func testAccCheckCloudngsparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cloudngsparameter name is set")
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
		// Singleton set-get resource: read with an empty name.
		data, err := client.FindResource(service.Cloudngsparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cloudngsparameter %s not found", n)
		}

		return nil
	}
}

const testAccCloudngsparameterDataSource_basic = `
	resource "citrixadc_cloudngsparameter" "tf_cloudngsparameter" {
		blockonallowedngstktprof   = "YES"
		allowedudtversion          = "V5"
		csvserverticketingdecouple = "NO"
		allowdtls12                = "NO"
	}

	data "citrixadc_cloudngsparameter" "tf_cloudngsparameter" {
		depends_on = [citrixadc_cloudngsparameter.tf_cloudngsparameter]
	}
`

func TestAccCloudngsparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudngsparameterDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cloudngsparameter.tf_cloudngsparameter", "blockonallowedngstktprof", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudngsparameter.tf_cloudngsparameter", "allowedudtversion", "V5"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudngsparameter.tf_cloudngsparameter", "csvserverticketingdecouple", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudngsparameter.tf_cloudngsparameter", "allowdtls12", "NO"),
				),
			},
		},
	})
}
