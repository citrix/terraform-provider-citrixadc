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

const testAccCrvserver_policymap_binding_basic = `

resource "citrixadc_crvserver" "crvserver" {
	name        = "my_vserver"
	servicetype = "HTTP"
	ipv46       = "10.102.80.55"
	port        = 8090
	cachetype   = "REVERSE"
	}
  resource "citrixadc_policymap" "tf_policymap" {
	mappolicyname = "ia_mappol123"
	sd            = "amazon.com"
	td            = "apple.com"
	}
  resource "citrixadc_lbvserver" "foo_lbvserver" {
	name        = "test_lbvserver"
	servicetype = "HTTP"
	ipv46       = "192.122.3.31"
	port        = 8000
	comment     = "hello"
	}
  resource "citrixadc_service" "tf_service" {
	lbvserver   = citrixadc_lbvserver.foo_lbvserver.name
	name        = "tf_service"
	port        = 8081
	ip          = "10.33.4.5"
	servicetype = "HTTP"
	cachetype   = "TRANSPARENT"
	}
  resource "citrixadc_crvserver_policymap_binding" "crvserver_policymap_binding" {
	name          = citrixadc_crvserver.crvserver.name
	policyname    = citrixadc_policymap.tf_policymap.mappolicyname
	targetvserver = citrixadc_lbvserver.foo_lbvserver.name
	depends_on = [
	  citrixadc_service.tf_service
	]
	}
  
`

const testAccCrvserver_policymap_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_crvserver" "crvserver" {
		name        = "my_vserver"
		servicetype = "HTTP"
		ipv46       = "10.102.80.55"
		port        = 8090
		cachetype   = "REVERSE"
	}
	  resource "citrixadc_policymap" "tf_policymap" {
		mappolicyname = "ia_mappol123"
		sd            = "amazon.com"
		td            = "apple.com"
	}
	  resource "citrixadc_lbvserver" "foo_lbvserver" {
		name        = "test_lbvserver"
		servicetype = "HTTP"
		ipv46       = "192.122.3.31"
		port        = 8000
		comment     = "hello"
	}
	  resource "citrixadc_service" "tf_service" {
		lbvserver   = citrixadc_lbvserver.foo_lbvserver.name
		name        = "tf_service"
		port        = 8081
		ip          = "10.33.4.5"
		servicetype = "HTTP"
		cachetype   = "TRANSPARENT"
	}
`

func TestAccCrvserver_policymap_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCrvserver_policymap_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_policymap_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_policymap_bindingExist("citrixadc_crvserver_policymap_binding.crvserver_policymap_binding", nil),
				),
			},
			{
				Config: testAccCrvserver_policymap_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_policymap_bindingNotExist("citrixadc_crvserver_policymap_binding.crvserver_policymap_binding", "my_vserver,ia_mappol123"),
				),
			},
		},
	})
}

func testAccCheckCrvserver_policymap_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No crvserver_policymap_binding id is set")
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
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "crvserver_policymap_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("crvserver_policymap_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_policymap_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "crvserver_policymap_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("crvserver_policymap_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_policymap_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_crvserver_policymap_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Crvserver_policymap_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("crvserver_policymap_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
