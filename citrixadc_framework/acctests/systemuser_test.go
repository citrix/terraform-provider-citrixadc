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

func TestAccSystemuser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemuser_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_user", nil),
				),
			},
			{
				Config: testAccSystemuser_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_user", nil),
				),
			},
		},
	})
}

func testAccCheckSystemuserExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
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
		data, err := client.FindResource(service.Systemuser.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemuserDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemuser" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Systemuser.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystemuser_basic_step1 = `
resource "citrixadc_systemuser" "tf_user" {
    username = "tf_user"
    password = "tf_password"
    timeout = 900

    cmdpolicybinding {
        policyname = "superuser"
        priority = 100
	}

    cmdpolicybinding {
        policyname = "network"
        priority = 200
	}
}
`

const testAccSystemuser_basic_step2 = `
resource "citrixadc_systemuser" "tf_user" {
    username = "tf_user"
    password = "tf_password"
    timeout = 200

}
`

const testAccSystemuserDataSource_basic = `
resource "citrixadc_systemuser" "tf_user" {
    username = "tf_user"
    password = "tf_password"
    timeout = 900

    cmdpolicybinding {
        policyname = "superuser"
        priority = 100
	}

    cmdpolicybinding {
        policyname = "network"
        priority = 200
	}
}

data "citrixadc_systemuser" "tf_user" {
    username = citrixadc_systemuser.tf_user.username
}
`

func TestAccSystemuserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemuserDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemuser.tf_user", "username", "tf_user"),
					resource.TestCheckResourceAttr("data.citrixadc_systemuser.tf_user", "timeout", "900"),
				),
			},
		},
	})
}
