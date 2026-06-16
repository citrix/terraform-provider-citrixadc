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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLsnappsprofile_port_binding_basic = `

	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile" {
		appsprofilename   = "my_lsn_appsprofile"
		transportprotocol = "TCP"
		mapping           = "ENDPOINT-INDEPENDENT"
	}

resource "citrixadc_lsnappsprofile_port_binding" "tf_lsnappsprofile_port_binding" {
	appsprofilename = citrixadc_lsnappsprofile.tf_lsnappsprofile.appsprofilename
	lsnport         = "80"
}
  
`

const testAccLsnappsprofile_port_binding_basic_step2 = `
	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile" {
		appsprofilename   = "my_lsn_appsprofile"
		transportprotocol = "TCP"
		mapping           = "ENDPOINT-INDEPENDENT"
	}


`

func TestAccLsnappsprofile_port_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsprofile_port_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsprofile_port_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsprofile_port_bindingExist("citrixadc_lsnappsprofile_port_binding.tf_lsnappsprofile_port_binding", nil),
				),
			},
			{
				Config: testAccLsnappsprofile_port_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsprofile_port_bindingNotExist("citrixadc_lsnappsprofile_port_binding.tf_lsnappsprofile_port_binding", "my_lsn_profile,80"),
				),
			},
		},
	})
}

// getLsnappsprofilePortBindingsForTest reads lsnappsprofile_port_binding members via
// the aggregate lsnappsprofile_binding endpoint (this ADC firmware has no direct GET
// for lsnappsprofile_port_binding) and returns the nested member array. Mirrors the
// resource's getLsnappsprofilePortBindings helper.
func getLsnappsprofilePortBindingsForTest(client *service.NitroClient, appsprofilename string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType:             "lsnappsprofile_binding",
		ResourceName:             appsprofilename,
		ResourceMissingErrorCode: 258,
	}
	aggArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}
	result := make([]map[string]interface{}, 0)
	for _, agg := range aggArr {
		nested, ok := agg["lsnappsprofile_port_binding"]
		if !ok || nested == nil {
			continue
		}
		nestedArr, ok := nested.([]interface{})
		if !ok {
			continue
		}
		for _, nn := range nestedArr {
			if m, ok := nn.(map[string]interface{}); ok {
				result = append(result, m)
			}
		}
	}
	return result, nil
}

func testAccCheckLsnappsprofile_port_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnappsprofile_port_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"appsprofilename", "lsnport"}, nil)
		if err != nil {
			return err
		}
		appsprofilename := idMap["appsprofilename"]
		lsnport := idMap["lsnport"]

		// No direct GET for lsnappsprofile_port_binding on this firmware; read via the
		// aggregate lsnappsprofile_binding endpoint (mirrors the resource getter).
		dataArr, err := getLsnappsprofilePortBindingsForTest(client, appsprofilename)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching lsnport
		found := false
		for _, v := range dataArr {
			if portVal, ok := v["lsnport"].(string); ok && portVal == lsnport {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsnappsprofile_port_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnappsprofile_port_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"appsprofilename", "lsnport"}, nil)
		if err != nil {
			return err
		}
		appsprofilename := idMap["appsprofilename"]
		lsnport := idMap["lsnport"]

		// No direct GET for lsnappsprofile_port_binding on this firmware; read via the
		// aggregate lsnappsprofile_binding endpoint (mirrors the resource getter).
		dataArr, err := getLsnappsprofilePortBindingsForTest(client, appsprofilename)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching lsnport
		found := false
		for _, v := range dataArr {
			if portVal, ok := v["lsnport"].(string); ok && portVal == lsnport {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsnappsprofile_port_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsnappsprofile_port_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnappsprofile_port_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"appsprofilename", "lsnport"}, nil)
		if err != nil {
			return err
		}
		appsprofilename := idMap["appsprofilename"]
		lsnport := idMap["lsnport"]

		// No direct GET for lsnappsprofile_port_binding; query the aggregate endpoint.
		// If the parent appsprofile is gone the query errors -> the binding is destroyed.
		dataArr, err := getLsnappsprofilePortBindingsForTest(client, appsprofilename)
		if err != nil {
			continue
		}
		for _, v := range dataArr {
			if portVal, ok := v["lsnport"].(string); ok && portVal == lsnport {
				return fmt.Errorf("lsnappsprofile_port_binding %s still exists", rs.Primary.ID)
			}
		}

	}

	return nil
}

const testAccLsnappsprofile_port_bindingDataSource_basic = `

	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile" {
		appsprofilename   = "my_lsn_appsprofile"
		transportprotocol = "TCP"
		mapping           = "ENDPOINT-INDEPENDENT"
	}

resource "citrixadc_lsnappsprofile_port_binding" "tf_lsnappsprofile_port_binding" {
	appsprofilename = citrixadc_lsnappsprofile.tf_lsnappsprofile.appsprofilename
	lsnport         = "80"
}

data "citrixadc_lsnappsprofile_port_binding" "tf_lsnappsprofile_port_binding" {
	appsprofilename = citrixadc_lsnappsprofile_port_binding.tf_lsnappsprofile_port_binding.appsprofilename
	lsnport         = citrixadc_lsnappsprofile_port_binding.tf_lsnappsprofile_port_binding.lsnport
}
`

func TestAccLsnappsprofile_port_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsprofile_port_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile_port_binding.tf_lsnappsprofile_port_binding", "appsprofilename", "my_lsn_appsprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile_port_binding.tf_lsnappsprofile_port_binding", "lsnport", "80"),
					resource.TestCheckResourceAttrSet("data.citrixadc_lsnappsprofile_port_binding.tf_lsnappsprofile_port_binding", "id"),
				),
			},
		},
	})
}
