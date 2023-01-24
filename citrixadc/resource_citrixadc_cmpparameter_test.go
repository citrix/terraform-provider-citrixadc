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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccCmpparameter_basic = `


resource "citrixadc_cmpparameter" "tf_cmpparameter" {
	cmplevel    = "optimal"
	quantumsize = 20
	servercmp   = "OFF"
  }
`
const testAccCmpparameter_update = `


	resource "citrixadc_cmpparameter" "tf_cmpparameter" {
		cmplevel    = "bestspeed"
		quantumsize = 30
		servercmp   = "ON"
	}
`

func TestAccCmpparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCmpparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpparameterExist("citrixadc_cmpparameter.tf_cmpparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "cmplevel", "optimal"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "quantumsize", "20"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "servercmp", "OFF"),
				),
			},
			resource.TestStep{
				Config: testAccCmpparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpparameterExist("citrixadc_cmpparameter.tf_cmpparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "cmplevel", "bestspeed"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "quantumsize", "30"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "servercmp", "ON"),
				),
			},
		},
	})
}

func testAccCheckCmpparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cmpparameter name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Cmpparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cmpparameter %s not found", n)
		}

		return nil
	}
}