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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLsnstatic_basic = `

	resource "citrixadc_lsnclient" "tf_lsnclient" {
		clientname = "my_lsnclient"
	}

	resource "citrixadc_lsnclient_network_binding" "tf_lsnclient_network_binding" {
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
	network    = "10.222.74.128"
	}

	resource "citrixadc_lsngroup" "tf_lsngroup_ds" {
	groupname     = "my_lsngroup_ds"
	clientname    = citrixadc_lsnclient.tf_lsnclient.clientname
	logging       = "DISABLED"
	nattype       = "DYNAMIC"
	snmptraplimit = 50
}

resource "citrixadc_lsnstatic" "tf_lsnstatic" {
	name              = "my_lsn_static"
	transportprotocol = "TCP"
	subscrip          = "10.222.74.128"
	subscrport        = 3000
	depends_on = [citrixadc_lsngroup.tf_lsngroup_ds,citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding]
	}
  
`
const testAccLsnstatic_update = `

	resource "citrixadc_lsnclient" "tf_lsnclient" {
		clientname = "my_lsnclient"
	}

	resource "citrixadc_lsnclient_network_binding" "tf_lsnclient_network_binding" {
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
	network    = "10.222.74.128"
	}

	resource "citrixadc_lsngroup" "tf_lsngroup_ds" {
	groupname     = "my_lsngroup_ds"
	clientname    = citrixadc_lsnclient.tf_lsnclient.clientname
	logging       = "DISABLED"
	nattype       = "DYNAMIC"
	snmptraplimit = 50
}

resource "citrixadc_lsnstatic" "tf_lsnstatic" {
	name              = "my_lsn_static"
	transportprotocol = "UDP"
	subscrip          = "10.222.74.128"
	subscrport        = 4000
	depends_on = [citrixadc_lsngroup.tf_lsngroup_ds,citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding]
	}
  
`

func TestAccLsnstatic_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnstaticDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnstatic_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnstaticExist("citrixadc_lsnstatic.tf_lsnstatic", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnstatic.tf_lsnstatic", "name", "my_lsn_static"),
					resource.TestCheckResourceAttr("citrixadc_lsnstatic.tf_lsnstatic", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsnstatic.tf_lsnstatic", "subscrip", "10.222.74.128"),
					resource.TestCheckResourceAttr("citrixadc_lsnstatic.tf_lsnstatic", "subscrport", "3000"),
				),
			},
			{
				Config: testAccLsnstatic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnstaticExist("citrixadc_lsnstatic.tf_lsnstatic", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnstatic.tf_lsnstatic", "name", "my_lsn_static"),
					resource.TestCheckResourceAttr("citrixadc_lsnstatic.tf_lsnstatic", "transportprotocol", "UDP"),
					resource.TestCheckResourceAttr("citrixadc_lsnstatic.tf_lsnstatic", "subscrip", "10.222.74.128"),
					resource.TestCheckResourceAttr("citrixadc_lsnstatic.tf_lsnstatic", "subscrport", "4000"),
				),
			},
		},
	})
}

func testAccCheckLsnstaticExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnstatic name is set")
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
		data, err := client.FindResource("lsnstatic", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnstatic %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnstaticDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnstatic" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnstatic", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnstatic %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsnstaticDataSource_basic = `

	resource "citrixadc_lsnclient" "tf_lsnclient" {
		clientname = "my_lsnclient"
	}

	resource "citrixadc_lsnclient_network_binding" "tf_lsnclient_network_binding" {
		clientname = citrixadc_lsnclient.tf_lsnclient.clientname
		network    = "10.222.74.128"
	}

	resource "citrixadc_lsngroup" "tf_lsngroup_ds" {
		groupname     = "my_lsngroup_ds"
		clientname    = citrixadc_lsnclient.tf_lsnclient.clientname
		logging       = "DISABLED"
		nattype       = "DYNAMIC"
		snmptraplimit = 50
	}

	resource "citrixadc_lsnstatic" "tf_lsnstatic_ds" {
		name              = "my_lsn_static_ds"
		transportprotocol = "TCP"
		subscrip          = "10.222.74.128"
		subscrport        = 3000
		depends_on = [citrixadc_lsngroup.tf_lsngroup_ds,citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding]
	}

	data "citrixadc_lsnstatic" "tf_lsnstatic_ds" {
		name    = citrixadc_lsnstatic.tf_lsnstatic_ds.name
		depends_on = [citrixadc_lsnstatic.tf_lsnstatic_ds]
	}
`

func TestAccLsnstaticDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnstaticDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnstatic.tf_lsnstatic_ds", "name", "my_lsn_static_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnstatic.tf_lsnstatic_ds", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnstatic.tf_lsnstatic_ds", "subscrip", "10.222.74.128"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnstatic.tf_lsnstatic_ds", "subscrport", "3000"),
				),
			},
		},
	})
}
