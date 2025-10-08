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
	"strings"
	"testing"
)

const testAccCsvserver_vpnvserver_binding_basic = `

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
	name           = "tf_vpnvserver"
	servicetype    = "SSL"
}

resource "citrixadc_csvserver_vpnvserver_binding" "tf_csvserver_vpnvserver_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	vserver = citrixadc_vpnvserver.tf_vpnvserver.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "SSL"
	sslprofile = citrixadc_sslprofile.tf_sslprofile.name
}

resource "citrixadc_sslprofile" "tf_sslprofile" {
	name = "tf_sslprofile"
	ecccurvebindings = []
}
`

const testAccCsvserver_vpnvserver_binding_basic_step2 = `

	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name           = "tf_vpnvserver"
		servicetype    = "SSL"
	}
	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "SSL"
		sslprofile = citrixadc_sslprofile.tf_sslprofile.name
	}

	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name = "tf_sslprofile"
		ecccurvebindings = []
	}
`

func TestAccCsvserver_vpnvserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCsvserver_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_vpnvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_vpnvserver_bindingExist("citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_vpnvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_vpnvserver_bindingNotExist("citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding", "tf_csvserver,tf_vpnvserver"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_vpnvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_vpnvserver_binding id is set")
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

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		vserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_vpnvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["vserver"].(string) == vserver {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("csvserver_vpnvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_vpnvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		vserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_vpnvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right vpn vserver name
		found := false
		for _, v := range dataArr {
			if v["vserver"].(string) == vserver {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("csvserver_vpnvserver_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_vpnvserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_vpnvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Csvserver_vpnvserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_vpnvserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
