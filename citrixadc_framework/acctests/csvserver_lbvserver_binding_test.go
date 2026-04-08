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

const testAccCsvserver_lbvserver_binding_basic = `
	resource "citrixadc_csvserver_lbvserver_binding" "tf_csvserver_lbvserver_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
        lbvserver = citrixadc_lbvserver.tf_lbvserver.name
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		servicetype = "HTTP"
		ipv46 = "10.10.10.10"
		port = 80
		lifecycle {
			ignore_changes = [lbvserverbinding]
		}
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name = "tf_lbvserver"
		servicetype = "HTTP"
	}
`

const testAccCsvserver_lbvserver_binding_basic_step2 = `
	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		servicetype = "HTTP"
		ipv46 = "10.10.10.10"
		port = 80
		lifecycle {
			ignore_changes = [lbvserverbinding]
		}
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name = "tf_lbvserver"
		servicetype = "HTTP"
	}
`

func TestAccCsvserver_lbvserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_lbvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_lbvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_lbvserver_bindingExist("citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding", "name", "tf_csvserver"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding", "lbvserver", "tf_lbvserver"),
				),
			},
			{
				Config: testAccCsvserver_lbvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_lbvserver_bindingNotExist("citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding", "name:tf_csvserver,lbvserver:tf_lbvserver"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_lbvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_lbvserver_binding id is set")
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

		// Parse the ID to extract the ipaddress
		bindingId := rs.Primary.ID
		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		lbvserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_lbvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right lbvserver name
		found := false
		for _, v := range dataArr {
			if v["lbvserver"].(string) == lbvserver {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("csvserver_lbvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_lbvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		// Parse the ID to extract the ipaddress
		bindingId := id
		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		lbvserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_lbvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right lbvserver name
		found := false
		for _, v := range dataArr {
			if v["lbvserver"].(string) == lbvserver {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("csvserver_lbvserver_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_lbvserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_lbvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		// Parse the ID to extract the ipaddress
		bindingId := rs.Primary.ID
		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]

		_, err = client.FindResource(service.Csvserver_lbvserver_binding.Type(), name)
		if err == nil {
			return fmt.Errorf("csvserver_lbvserver_binding still exists")
		}

	}

	return nil
}

const testAccCsvserver_lbvserver_bindingDataSource_basic = `
	resource "citrixadc_csvserver" "tf_csvserver_ds" {
		name = "tf_csvserver_ds"
		servicetype = "HTTP"
		ipv46 = "10.10.10.10"
		port = 80
		lifecycle {
			ignore_changes = [lbvserverbinding]
		}
	}

	resource "citrixadc_lbvserver" "tf_lbvserver_ds" {
		name = "tf_lbvserver_ds"
		servicetype = "HTTP"
	}

	resource "citrixadc_csvserver_lbvserver_binding" "tf_csvserver_lbvserver_binding_ds" {
		name = citrixadc_csvserver.tf_csvserver_ds.name
		lbvserver = citrixadc_lbvserver.tf_lbvserver_ds.name
	}

	data "citrixadc_csvserver_lbvserver_binding" "tf_csvserver_lbvserver_binding_ds_read" {
		name = citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding_ds.name
		lbvserver = citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding_ds.lbvserver
		depends_on = [citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding_ds]
	}
`

func TestAcccsvserver_lbvserver_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_lbvserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding_ds_read", "name", "tf_csvserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding_ds_read", "lbvserver", "tf_lbvserver_ds"),
					resource.TestCheckResourceAttrSet("data.citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding_ds_read", "id"),
				),
			},
		},
	})
}
