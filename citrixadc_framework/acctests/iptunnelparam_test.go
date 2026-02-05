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

const testAccIptunnelparam_add = `
	resource "citrixadc_iptunnelparam" "tf_iptunnelparam" {
		dropfrag             = "YES"
		dropfragcputhreshold = 1
		srciproundrobin      = "NO"
		enablestrictrx       = "YES"
		enablestricttx       = "YES"
		useclientsourceip    = "YES"
	}
`
const testAccIptunnelparam_update = `
	resource "citrixadc_iptunnelparam" "tf_iptunnelparam" {
		dropfrag             = "NO"
		dropfragcputhreshold = 1
		srciproundrobin      = "NO"
		enablestrictrx       = "NO"
		enablestricttx       = "NO"
		useclientsourceip    = "NO"
	}
`

func TestAccIptunnelparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccIptunnelparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelparamExist("citrixadc_iptunnelparam.tf_iptunnelparam", nil),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_iptunnelparam", "dropfrag", "YES"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_iptunnelparam", "enablestrictrx", "YES"),
				),
			},
			{
				Config: testAccIptunnelparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelparamExist("citrixadc_iptunnelparam.tf_iptunnelparam", nil),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_iptunnelparam", "dropfrag", "NO"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_iptunnelparam", "enablestrictrx", "NO"),
				),
			},
		},
	})
}

func testAccCheckIptunnelparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No iptunnelparam name is set")
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
		data, err := client.FindResource(service.Iptunnelparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("iptunnelparam %s not found", n)
		}

		return nil
	}
}

const testAccIptunnelparamDataSource_basic = `

	resource "citrixadc_iptunnelparam" "tf_iptunnelparam" {
		dropfrag             = "YES"
		dropfragcputhreshold = 1
		srciproundrobin      = "NO"
		enablestrictrx       = "YES"
		enablestricttx       = "YES"
		useclientsourceip    = "YES"
	}

	data "citrixadc_iptunnelparam" "tf_iptunnelparam" {
		depends_on = [citrixadc_iptunnelparam.tf_iptunnelparam]
	}
`

func TestAccIptunnelparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccIptunnelparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "dropfrag", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "dropfragcputhreshold", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "srciproundrobin", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "enablestrictrx", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "enablestricttx", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "useclientsourceip", "YES"),
				),
			},
		},
	})
}
