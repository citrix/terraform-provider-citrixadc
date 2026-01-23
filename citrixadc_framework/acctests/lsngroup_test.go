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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccLsngroup_basic = `

	resource "citrixadc_lsngroup" "tf_lsngroup" {
		groupname     = "my_lsngroup"
		clientname    = "my_lsnclient"
		logging       = "DISABLED"
		nattype       = "DYNAMIC"
		snmptraplimit = 50
	}
`
const testAccLsngroup_update = `

	resource "citrixadc_lsngroup" "tf_lsngroup" {
		groupname     = "my_lsngroup"
		clientname    = "my_lsnclient"
		nattype       = "DYNAMIC"
		snmptraplimit = 100
	}
`

func TestAccLsngroup_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this LSN resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroupExist("citrixadc_lsngroup.tf_lsngroup", nil),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "groupname", "my_lsngroup"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "clientname", "my_lsnclient"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "logging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "nattype", "DYNAMIC"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "snmptraplimit", "50"),
				),
			},
			{
				Config: testAccLsngroup_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroupExist("citrixadc_lsngroup.tf_lsngroup", nil),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "groupname", "my_lsngroup"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "clientname", "my_lsnclient"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "nattype", "DYNAMIC"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "snmptraplimit", "100"),
				),
			},
		},
	})
}

func testAccCheckLsngroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsngroup name is set")
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
		data, err := client.FindResource("lsngroup", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsngroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsngroupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsngroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsngroup", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsngroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
