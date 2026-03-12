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

func TestAccSystemgroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemgroup_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemgroupExist("citrixadc_systemgroup.tf_systemgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_systemgroup.tf_systemgroup", "warnpriorndays", "10"),
					resource.TestCheckResourceAttr("citrixadc_systemgroup.tf_systemgroup", "daystoexpire", "45"),
				),
			},
			{
				Config: testAccSystemgroup_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemgroupExist("citrixadc_systemgroup.tf_systemgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_systemgroup.tf_systemgroup", "warnpriorndays", "15"),
					resource.TestCheckResourceAttr("citrixadc_systemgroup.tf_systemgroup", "daystoexpire", "60"),
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Systemgroup.Type(), rs.Primary.ID)

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemgroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Systemgroup.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystemgroup_basic_step1 = `
resource "citrixadc_systemparameter" "tf_systemparameter" {
    passwordhistorycontrol = "ENABLED"
}

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
	warnpriorndays = 10
	daystoexpire = 45

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
	warnpriorndays = 15
	daystoexpire = 60

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
resource "citrixadc_systemparameter" "tf_systemparameter" {
    passwordhistorycontrol = "DISABLED"
}


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

const testAccSystemgroupDataSource_basic = `
resource "citrixadc_systemparameter" "tf_systemparameter_ds" {
    passwordhistorycontrol = "ENABLED"
}

resource "citrixadc_systemgroup" "tf_systemgroup_ds" {
    groupname = "tf_datasource_group"
    timeout = 999
    promptstring = "test>"
    warnpriorndays = 10
    daystoexpire = 45
    depends_on = [citrixadc_systemparameter.tf_systemparameter_ds]
}

data "citrixadc_systemgroup" "tf_systemgroup_datasource" {
    groupname = citrixadc_systemgroup.tf_systemgroup_ds.groupname
}
`

func TestAccSystemgroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemgroupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemgroup.tf_systemgroup_datasource", "groupname", "tf_datasource_group"),
					resource.TestCheckResourceAttr("data.citrixadc_systemgroup.tf_systemgroup_datasource", "timeout", "999"),
					resource.TestCheckResourceAttr("data.citrixadc_systemgroup.tf_systemgroup_datasource", "promptstring", "test>"),
					resource.TestCheckResourceAttr("data.citrixadc_systemgroup.tf_systemgroup_datasource", "warnpriorndays", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_systemgroup.tf_systemgroup_datasource", "daystoexpire", "45"),
					resource.TestCheckResourceAttrSet("data.citrixadc_systemgroup.tf_systemgroup_datasource", "id"),
				),
			},
		},
	})
}
