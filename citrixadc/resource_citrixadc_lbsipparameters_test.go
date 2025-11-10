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

const testAccLbsipparameters_basic = `
	resource "citrixadc_lbsipparameters" "tf_lbsipparameters" {
		addrportvip = "ENABLED"
		retrydur = 100
		rnatdstport = 80
		rnatsecuredstport = 81
		rnatsecuresrcport = 82
		rnatsrcport = 83
		sip503ratethreshold = 15
	}
`

const testAccLbsipparameters_basic_update = `
	resource "citrixadc_lbsipparameters" "tf_lbsipparameters" {
		addrportvip = "DISABLED"
		retrydur = 120
		rnatdstport = 1
		rnatsecuredstport = 2
		rnatsecuresrcport = 3
		rnatsrcport = 4
		sip503ratethreshold = 100
	}
`

func TestAccLbsipparameters_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		// lbsipparameters resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLbsipparameters_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbsipparametersExist("citrixadc_lbsipparameters.tf_lbsipparameters", nil),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "addrportvip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "retrydur", "100"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatdstport", "80"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsecuredstport", "81"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsecuresrcport", "82"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsrcport", "83"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "sip503ratethreshold", "15"),
				),
			},
			{
				Config: testAccLbsipparameters_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbsipparametersExist("citrixadc_lbsipparameters.tf_lbsipparameters", nil),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "addrportvip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "retrydur", "120"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatdstport", "1"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsecuredstport", "2"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsecuresrcport", "3"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsrcport", "4"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "sip503ratethreshold", "100"),
				),
			},
		},
	})
}

func testAccCheckLbsipparametersExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbsipparameters name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lbsipparameters.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lbsipparameters %s not found", n)
		}

		return nil
	}
}
