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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccFeoparameter_basic = `
	
resource "citrixadc_feoparameter" "tf_feoparameter" {
		jpegqualitypercent = 10
		cssinlinethressize = 100
		jsinlinethressize  = 50
		imginlinethressize = 1
	}
  
`

const testAccFeoparameter_update = `
	resource "citrixadc_feoparameter" "tf_feoparameter" {
		jpegqualitypercent = 0
		cssinlinethressize = 50
		jsinlinethressize  = 100
		imginlinethressize = 20
	}
  
`

func TestAccFeoparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccFeoparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoparameterExist("citrixadc_feoparameter.tf_feoparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "jpegqualitypercent", "10"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "cssinlinethressize", "100"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "jsinlinethressize", "50"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "imginlinethressize", "1"),
				),
			},
			resource.TestStep{
				Config: testAccFeoparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoparameterExist("citrixadc_feoparameter.tf_feoparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "jpegqualitypercent", "0"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "cssinlinethressize", "50"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "jsinlinethressize", "100"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "imginlinethressize", "20"),
				),
			},
		},
	})
}

func testAccCheckFeoparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No feoparameter name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("feoparameter", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("feoparameter %s not found", n)
		}

		return nil
	}
}