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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccClusternode_basic_nogroup_config = `


resource "citrixadc_clusternode" "tf_clusternode" {
	nodeid             = 1
	ipaddress          = "10.222.74.150"
	state              = "PASSIVE"
  }
  
`
const testAccClusternode_update_nogroup_config = `


resource "citrixadc_clusternode" "tf_clusternode" {
	nodeid             = 1
	ipaddress          = "10.222.74.150"
	state              = "ACTIVE"  
  }
  
`

func TestAccClusternode_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClusternodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternode_basic_nogroup_config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.222.74.150"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "PASSIVE"),
				),
			},
			{
				Config: testAccClusternode_update_nogroup_config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.222.74.150"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "ACTIVE"),
				),
			},
		},
	})
}

const testAccClusternode_basic_group_config_yes = `


	resource "citrixadc_clusternode" "tf_clusternode" {
		nodeid               = 1
		ipaddress            = "10.222.74.150"
		state                = "PASSIVE"
		clearnodegroupconfig = "YES"
	}
`
const testAccClusternode_update_group_config_yes = `


	resource "citrixadc_clusternode" "tf_clusternode" {
		nodeid               = 1
		ipaddress            = "10.222.74.150"
		state                = "ACTIVE"
		clearnodegroupconfig = "YES"
	}
  
`

func TestAccClusternode_group_config_yes(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClusternodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternode_basic_group_config_yes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.222.74.150"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "PASSIVE"),
				),
			},
			{
				Config: testAccClusternode_update_group_config_yes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.222.74.150"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "ACTIVE"),
				),
			},
		},
	})
}

const testAccClusternode_basic_group_config_no = `


	resource "citrixadc_clusternode" "tf_clusternode" {
		nodeid               = 1
		ipaddress            = "10.222.74.150"
		state                = "PASSIVE"
		clearnodegroupconfig = "NO"
	}
`
const testAccClusternode_update_group_config_no = `


	resource "citrixadc_clusternode" "tf_clusternode" {
		nodeid               = 1
		ipaddress            = "10.222.74.150"
		state                = "ACTIVE"
		clearnodegroupconfig = "NO"
	}
  
`

func TestAccClusternode_group_config_no(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClusternodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternode_basic_group_config_no,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.222.74.150"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "PASSIVE"),
				),
			},
			{
				Config: testAccClusternode_update_group_config_no,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.222.74.150"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "ACTIVE"),
				),
			},
		},
	})
}

func testAccCheckClusternodeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusternode name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Clusternode.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("clusternode %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusternodeDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusternode" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Clusternode.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusternode %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
