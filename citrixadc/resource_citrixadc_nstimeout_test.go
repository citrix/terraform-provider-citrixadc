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

const testAccNstimeout_basic = `


resource "citrixadc_nstimeout" "tf_nstimeout" {
	zombie     = 60
	client     = 2000
	server     = 2000
	httpclient = 2000
	reducedrsttimeout = 10
	}
  
`
const testAccNstimeout_update = `


resource "citrixadc_nstimeout" "tf_nstimeout" {
	zombie     = 70
	client     = 2300
	server     = 2400
	httpclient = 2500
	reducedrsttimeout = 15
	}
  
`

func TestAccNstimeout_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNstimeout_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstimeoutExist("citrixadc_nstimeout.tf_nstimeout", nil),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "zombie", "60"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "client", "2000"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "server", "2000"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "httpclient", "2000"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "reducedrsttimeout", "10"),
				),
			},
			{
				Config: testAccNstimeout_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstimeoutExist("citrixadc_nstimeout.tf_nstimeout", nil),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "zombie", "70"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "client", "2300"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "server", "2400"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "httpclient", "2500"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "reducedrsttimeout", "15"),
				),
			},
		},
	})
}

func testAccCheckNstimeoutExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nstimeout name is set")
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
		data, err := client.FindResource(service.Nstimeout.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nstimeout %s not found", n)
		}

		return nil
	}
}
