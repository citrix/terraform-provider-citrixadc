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

const testAccAaaparameter_basic = `

	resource "citrixadc_aaaparameter" "tf_aaaparameter" {
		enablestaticpagecaching    = "NO"
		enableenhancedauthfeedback = "YES"
		defaultauthtype            = "LOCAL"
		maxaaausers                = 3
		maxloginattempts           = 5
		failedlogintimeout         = 15
	}
  
`
const testAccAaaparameter_update = `

	resource "citrixadc_aaaparameter" "tf_aaaparameter" {
		enablestaticpagecaching    = "YES"
		enableenhancedauthfeedback = "NO"
		defaultauthtype            = "LOCAL"
		maxaaausers                = 4
		maxloginattempts           = 10
		failedlogintimeout         = 20
	}
  
`

func TestAccAaaparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaparameterExist("citrixadc_aaaparameter.tf_aaaparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enablestaticpagecaching", "NO"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enableenhancedauthfeedback", "YES"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "defaultauthtype", "LOCAL"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "maxaaausers", "3"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "maxloginattempts", "5"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "failedlogintimeout", "15"),
				),
			},
			{
				Config: testAccAaaparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaparameterExist("citrixadc_aaaparameter.tf_aaaparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enablestaticpagecaching", "YES"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enableenhancedauthfeedback", "NO"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "defaultauthtype", "LOCAL"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "maxaaausers", "4"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "maxloginattempts", "10"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "failedlogintimeout", "20"),
				),
			},
		},
	})
}

func testAccCheckAaaparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaaparameter name is set")
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
		data, err := client.FindResource(service.Aaaparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaaparameter %s not found", n)
		}

		return nil
	}
}

const testAccAaaparameterDataSource_basic = `

	resource "citrixadc_aaaparameter" "tf_aaaparameter" {
		enablestaticpagecaching    = "NO"
		enableenhancedauthfeedback = "YES"
		defaultauthtype            = "LOCAL"
		maxloginattempts           = 5
		failedlogintimeout         = 15
	}
	
	data "citrixadc_aaaparameter" "tf_aaaparameter" {
		depends_on = [citrixadc_aaaparameter.tf_aaaparameter]
	}
`

func TestAccAaaparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaaparameter.tf_aaaparameter", "enablestaticpagecaching", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaparameter.tf_aaaparameter", "enableenhancedauthfeedback", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaparameter.tf_aaaparameter", "defaultauthtype", "LOCAL"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaparameter.tf_aaaparameter", "maxloginattempts", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaparameter.tf_aaaparameter", "failedlogintimeout", "15"),
				),
			},
		},
	})
}
