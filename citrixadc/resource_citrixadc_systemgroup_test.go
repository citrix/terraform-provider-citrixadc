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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSystemgroup_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSystemgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemgroup_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemgroupExist("citrixadc_systemgroup.tf_systemgroup", nil),
				),
			},
			{
				Config: testAccSystemgroup_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemgroupExist("citrixadc_systemgroup.tf_systemgroup", nil),
				),
			},
			{
				Config: testAccSystemgroup_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemgroupExist("citrixadc_systemgroup.tf_systemgroup", nil),
				),
			},
		},
	})
}

func testAccCheckSystemgroupExist(n string, id *string) resource.TestCheckFunc {
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

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Systemgroup.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemgroupDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemgroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Systemgroup.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystemgroup_basic_step1 = `

resource "citrixadc_systemuser" "tf_user1" {
    username = "tf_user1"
    password = "tf_password"
    timeout = 900

}

resource "citrixadc_systemuser" "tf_user2" {
    username = "tf_user2"
    password = "tf_password"
    timeout = 900

}

resource "citrixadc_systemgroup" "tf_systemgroup" {
    groupname = "testgroupname"
    timeout = 999
    promptstring = "bye>"

    cmdpolicybinding { 
        policyname = "superuser"
        priority = 100
    }

    cmdpolicybinding { 
        policyname = "network"
        priority = 200
    }

    systemusers = [
        citrixadc_systemuser.tf_user1.username,
		citrixadc_systemuser.tf_user2.username,
    ]
}


`

const testAccSystemgroup_basic_step2 = `

resource "citrixadc_systemuser" "tf_user1" {
    username = "tf_user1"
    password = "tf_password"
    timeout = 900

}

resource "citrixadc_systemuser" "tf_user2" {
    username = "tf_user2"
    password = "tf_password"
    timeout = 900

}

resource "citrixadc_systemgroup" "tf_systemgroup" {
    groupname = "testgroupname"
    timeout = 999
    promptstring = "bye>"

    cmdpolicybinding { 
        policyname = "superuser"
        priority = 200
    }

    cmdpolicybinding { 
        policyname = "network"
        priority = 100
    }

    systemusers = [
		citrixadc_systemuser.tf_user2.username
    ]
}


`
const testAccSystemgroup_basic_step3 = `

resource "citrixadc_systemuser" "tf_user1" {
    username = "tf_user1"
    password = "tf_password"
    timeout = 900

}

resource "citrixadc_systemuser" "tf_user2" {
    username = "tf_user2"
    password = "tf_password"
    timeout = 900

}

resource "citrixadc_systemgroup" "tf_systemgroup" {
    groupname = "testgroupname"
    timeout = 1000
    promptstring = "hello>"

    systemusers = [
		citrixadc_systemuser.tf_user1.username,
		citrixadc_systemuser.tf_user2.username,
    ]
}


`
