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

const testAccCrvserver_add = `

	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver"
		servicetype = "HTTP"
		arp = "OFF"
	}
`

const testAccCrvserver_update = `

	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver"
		servicetype = "HTTP"
		arp = "ON"
	}
`

func TestAccCrvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserverExist("citrixadc_crvserver.crvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "name", "my_vserver"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "arp", "OFF"),
				),
			},
			{
				Config: testAccCrvserver_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserverExist("citrixadc_crvserver.crvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "name", "my_vserver"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "arp", "ON"),
				),
			},
		},
	})
}

func testAccCheckCrvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No crvserver name is set")
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
		data, err := client.FindResource(service.Crvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("crvserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckCrvserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_crvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Crvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("crvserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCrvserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_crvserver.tf_crvserver_ds", "name", "my_vserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_crvserver.tf_crvserver_ds", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("data.citrixadc_crvserver.tf_crvserver_ds", "arp", "OFF"),
				),
			},
		},
	})
}

const testAccCrvserverDataSource_basic = `

resource "citrixadc_crvserver" "tf_crvserver_ds" {
	name        = "my_vserver_ds"
	servicetype = "HTTP"
	arp         = "OFF"
}

data "citrixadc_crvserver" "tf_crvserver_ds" {
	name       = citrixadc_crvserver.tf_crvserver_ds.name
	depends_on = [citrixadc_crvserver.tf_crvserver_ds]
}

`
