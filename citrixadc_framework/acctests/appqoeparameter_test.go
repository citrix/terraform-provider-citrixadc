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

const testAccAppqoeparameter_basic = `

	resource "citrixadc_appqoeparameter" "tf_appqoeparameter" {
		sessionlife         = 300
		avgwaitingclient    = 400
		maxaltrespbandwidth = 50
		dosattackthresh     = 100
	}
`
const testAccAppqoeparameter_update = `

	resource "citrixadc_appqoeparameter" "tf_appqoeparameter" {
		sessionlife         = 400
		avgwaitingclient    = 300
		maxaltrespbandwidth = 100
		dosattackthresh     = 50
	}
`

func TestAccAppqoeparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppqoeparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeparameterExist("citrixadc_appqoeparameter.tf_appqoeparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "sessionlife", "300"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "avgwaitingclient", "400"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "maxaltrespbandwidth", "50"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "dosattackthresh", "100"),
				),
			},
			{
				Config: testAccAppqoeparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeparameterExist("citrixadc_appqoeparameter.tf_appqoeparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "sessionlife", "400"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "avgwaitingclient", "300"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "maxaltrespbandwidth", "100"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "dosattackthresh", "50"),
				),
			},
		},
	})
}

func testAccCheckAppqoeparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appqoeparameter name is set")
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
		data, err := client.FindResource(service.Appqoeparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appqoeparameter %s not found", n)
		}

		return nil
	}
}

const testAccAppqoeparameterDataSource_basic = `

	resource "citrixadc_appqoeparameter" "tf_appqoeparameter" {
		sessionlife         = 300
		avgwaitingclient    = 400
		maxaltrespbandwidth = 50
		dosattackthresh     = 100
	}

	data "citrixadc_appqoeparameter" "tf_appqoeparameter" {
		depends_on = [citrixadc_appqoeparameter.tf_appqoeparameter]
	}
`

func TestAccAppqoeparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppqoeparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appqoeparameter.tf_appqoeparameter", "sessionlife", "300"),
					resource.TestCheckResourceAttr("data.citrixadc_appqoeparameter.tf_appqoeparameter", "avgwaitingclient", "400"),
					resource.TestCheckResourceAttr("data.citrixadc_appqoeparameter.tf_appqoeparameter", "maxaltrespbandwidth", "50"),
					resource.TestCheckResourceAttr("data.citrixadc_appqoeparameter.tf_appqoeparameter", "dosattackthresh", "100"),
				),
			},
		},
	})
}
