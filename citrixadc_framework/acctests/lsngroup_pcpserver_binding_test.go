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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLsngroup_pcpserver_binding_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsnclient"
}

resource "citrixadc_lsngroup" "tf_lsngroup" {
	groupname  = "my_lsn_group"
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
}

resource "citrixadc_nsip" "tf_nsip" {
	ipaddress = "10.222.74.185"
	netmask   = "255.255.255.0"
	type      = "SNIP"
}

resource "citrixadc_pcpserver" "tf_pcpserver" {
	name       = "my_pcpserver"
	ipaddress  = "10.222.74.185"
	port       = 5351
	depends_on = [citrixadc_nsip.tf_nsip]
}

resource "citrixadc_lsngroup_pcpserver_binding" "tf_lsngroup_pcpserver_binding" {
	groupname = citrixadc_lsngroup.tf_lsngroup.groupname
	pcpserver = citrixadc_pcpserver.tf_pcpserver.name
}
  
`

const testAccLsngroup_pcpserver_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccLsngroup_pcpserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_pcpserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_pcpserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_pcpserver_bindingExist("citrixadc_lsngroup_pcpserver_binding.tf_lsngroup_pcpserver_binding", nil),
				),
			},
			{
				Config: testAccLsngroup_pcpserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_pcpserver_bindingNotExist("citrixadc_lsngroup_pcpserver_binding.tf_lsngroup_pcpserver_binding", "my_lsn_group,my_pcpserver"),
				),
			},
		},
	})
}

func testAccCheckLsngroup_pcpserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsngroup_pcpserver_binding id is set")
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

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		groupname := idSlice[0]
		pcpserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lsngroup_pcpserver_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching pcpserver
		found := false
		for _, v := range dataArr {
			if v["pcpserver"].(string) == pcpserver {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsngroup_pcpserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_pcpserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		groupname := idSlice[0]
		pcpserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lsngroup_pcpserver_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching pcpserver
		found := false
		for _, v := range dataArr {
			if v["pcpserver"].(string) == pcpserver {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsngroup_pcpserver_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_pcpserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsngroup_pcpserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsngroup_pcpserver_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsngroup_pcpserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsngroup_pcpserver_bindingDataSource_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsnclient"
}

resource "citrixadc_lsngroup" "tf_lsngroup" {
	groupname  = "my_lsn_group"
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
}

resource "citrixadc_nsip" "tf_nsip" {
	ipaddress = "10.222.74.185"
	netmask   = "255.255.255.0"
	type      = "SNIP"
}

resource "citrixadc_pcpserver" "tf_pcpserver" {
	name       = "my_pcpserver"
	ipaddress  = "10.222.74.185"
	port       = 5351
	depends_on = [citrixadc_nsip.tf_nsip]
}

resource "citrixadc_lsngroup_pcpserver_binding" "tf_lsngroup_pcpserver_binding" {
	groupname = citrixadc_lsngroup.tf_lsngroup.groupname
	pcpserver = citrixadc_pcpserver.tf_pcpserver.name
}

data "citrixadc_lsngroup_pcpserver_binding" "tf_lsngroup_pcpserver_binding" {
	groupname = citrixadc_lsngroup_pcpserver_binding.tf_lsngroup_pcpserver_binding.groupname
	pcpserver = citrixadc_lsngroup_pcpserver_binding.tf_lsngroup_pcpserver_binding.pcpserver
}
`

func TestAccLsngroup_pcpserver_bindingDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_pcpserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_pcpserver_binding.tf_lsngroup_pcpserver_binding", "groupname", "my_lsn_group"),
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_pcpserver_binding.tf_lsngroup_pcpserver_binding", "pcpserver", "my_pcpserver"),
					resource.TestCheckResourceAttrSet("data.citrixadc_lsngroup_pcpserver_binding.tf_lsngroup_pcpserver_binding", "id"),
				),
			},
		},
	})
}
