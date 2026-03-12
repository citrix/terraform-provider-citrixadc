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

const testAccCspolicylabel_basic = `
	resource "citrixadc_cspolicylabel" "tf_policylabel" {
		cspolicylabeltype = "HTTP"
		labelname = "tf_policylabel"
	}
`

func TestAccCspolicylabel_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCspolicylabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCspolicylabel_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspolicylabelExist("citrixadc_cspolicylabel.tf_policylabel", nil),
				),
			},
		},
	})
}

func testAccCheckCspolicylabelExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cspolicylabel name is set")
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
		data, err := client.FindResource(service.Cspolicylabel.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cspolicylabel %s not found", n)
		}

		return nil
	}
}

func testAccCheckCspolicylabelDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cspolicylabel" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cspolicylabel.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cspolicylabel %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCspolicylabelDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCspolicylabelDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cspolicylabel.tf_policylabel_ds", "labelname", "tf_policylabel_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_cspolicylabel.tf_policylabel_ds", "cspolicylabeltype", "HTTP"),
				),
			},
		},
	})
}

const testAccCspolicylabelDataSource_basic = `

resource "citrixadc_cspolicylabel" "tf_policylabel_ds" {
	cspolicylabeltype = "HTTP"
	labelname = "tf_policylabel_ds"
}

data "citrixadc_cspolicylabel" "tf_policylabel_ds" {
	labelname = citrixadc_cspolicylabel.tf_policylabel_ds.labelname
	depends_on = [citrixadc_cspolicylabel.tf_policylabel_ds]
}

`
