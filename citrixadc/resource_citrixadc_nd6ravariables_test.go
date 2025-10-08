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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccNd6ravariables_basic = `


	resource "citrixadc_nd6ravariables" "tf_nd6ravariables" {
		vlan                     = 1
		ceaserouteradv           = "NO"
		onlyunicastrtadvresponse = "NO"
		srclinklayeraddroption   = "NO"
	}
`
const testAccNd6ravariables_update = `


	resource "citrixadc_nd6ravariables" "tf_nd6ravariables" {
		vlan                     = 1
		ceaserouteradv           = "YES"
		onlyunicastrtadvresponse = "YES"
		srclinklayeraddroption   = "YES"
	}
`

func TestAccNd6ravariables_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNd6ravariables_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNd6ravariablesExist("citrixadc_nd6ravariables.tf_nd6ravariables", nil),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "ceaserouteradv", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "onlyunicastrtadvresponse", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "srclinklayeraddroption", "NO"),
				),
			},
			{
				Config: testAccNd6ravariables_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNd6ravariablesExist("citrixadc_nd6ravariables.tf_nd6ravariables", nil),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "ceaserouteradv", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "onlyunicastrtadvresponse", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "srclinklayeraddroption", "YES"),
				),
			},
		},
	})
}

func testAccCheckNd6ravariablesExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nd6ravariables name is set")
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
		data, err := client.FindResource(service.Nd6ravariables.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nd6ravariables %s not found", n)
		}

		return nil
	}
}
