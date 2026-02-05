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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLbmetrictable_add = `
resource "citrixadc_lbmetrictable" "tfAcc_lbmetrictable" {
        metrictable = "tf_lbmetrictable"
}
`

func TestAccLbmetrictable_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmetrictableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmetrictable_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmetrictableExist("citrixadc_lbmetrictable.tfAcc_lbmetrictable", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmetrictable.tfAcc_lbmetrictable", "metrictable", "tf_lbmetrictable"),
				),
			},
		},
	})
}

func testAccCheckLbmetrictableExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb metrictable name is set")
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
		data, err := client.FindResource("lbmetrictable", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB metrictable %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbmetrictableDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbmetrictable" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbmetrictable", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB metrictable %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLbmetrictableDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmetrictableDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbmetrictable.tf_lbmetrictable_ds", "metrictable", "tf_lbmetrictable_ds"),
				),
			},
		},
	})
}

const testAccLbmetrictableDataSource_basic = `

resource "citrixadc_lbmetrictable" "tf_lbmetrictable_ds" {
  metrictable = "tf_lbmetrictable_ds"
}

data "citrixadc_lbmetrictable" "tf_lbmetrictable_ds" {
  metrictable = citrixadc_lbmetrictable.tf_lbmetrictable_ds.metrictable
  depends_on  = [citrixadc_lbmetrictable.tf_lbmetrictable_ds]
}

`
