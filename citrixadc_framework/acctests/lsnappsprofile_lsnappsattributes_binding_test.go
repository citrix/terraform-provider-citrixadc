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

const testAccLsnappsprofile_lsnappsattributes_binding_basic = `

	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile" {
		appsprofilename   = "my_lsn_profile"
		transportprotocol = "TCP"
		mapping           = "ENDPOINT-INDEPENDENT"
	}

	# Prerequisite: the appsprofile must have a port binding whose range covers the
	# appsattributes port (90) before an appsattributes binding is accepted (NITRO errorcode 257).
	resource "citrixadc_lsnappsprofile_port_binding" "tf_lsnappsprofile_port_binding" {
		appsprofilename = citrixadc_lsnappsprofile.tf_lsnappsprofile.appsprofilename
		lsnport         = "80-100"
	}

	resource "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
		name              = "my_lsn_appattributes"
		transportprotocol = "TCP"
		port              = 90
		sessiontimeout    = 40
	}
resource "citrixadc_lsnappsprofile_lsnappsattributes_binding" "tf_lsnappsprofile_lsnappsattributes_binding" {
	appsprofilename    = citrixadc_lsnappsprofile.tf_lsnappsprofile.appsprofilename
	appsattributesname = citrixadc_lsnappsattributes.tf_lsnappsattributes.name
	depends_on         = [citrixadc_lsnappsprofile_port_binding.tf_lsnappsprofile_port_binding]
}
  
`

const testAccLsnappsprofile_lsnappsattributes_binding_basic_step2 = `
	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile" {
		appsprofilename   = "my_lsn_profile"
		transportprotocol = "TCP"
		mapping           = "ENDPOINT-INDEPENDENT"
	}

	# Prerequisite: the appsprofile must have a port binding whose range covers the
	# appsattributes port (90) before an appsattributes binding is accepted (NITRO errorcode 257).
	resource "citrixadc_lsnappsprofile_port_binding" "tf_lsnappsprofile_port_binding" {
		appsprofilename = citrixadc_lsnappsprofile.tf_lsnappsprofile.appsprofilename
		lsnport         = "80-100"
	}

	resource "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
		name              = "my_lsn_appattributes"
		transportprotocol = "TCP"
		port              = 90
		sessiontimeout    = 40
	}

`

const testAccLsnappsprofile_lsnappsattributes_bindingDataSource_basic = `

	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile" {
		appsprofilename   = "my_lsn_profile"
		transportprotocol = "TCP"
		mapping           = "ENDPOINT-INDEPENDENT"
	}

	# Prerequisite: the appsprofile must have a port binding whose range covers the
	# appsattributes port (90) before an appsattributes binding is accepted (NITRO errorcode 257).
	resource "citrixadc_lsnappsprofile_port_binding" "tf_lsnappsprofile_port_binding" {
		appsprofilename = citrixadc_lsnappsprofile.tf_lsnappsprofile.appsprofilename
		lsnport         = "80-100"
	}

	resource "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
		name              = "my_lsn_appattributes"
		transportprotocol = "TCP"
		port              = 90
		sessiontimeout    = 40
	}
resource "citrixadc_lsnappsprofile_lsnappsattributes_binding" "tf_lsnappsprofile_lsnappsattributes_binding" {
	appsprofilename    = citrixadc_lsnappsprofile.tf_lsnappsprofile.appsprofilename
	appsattributesname = citrixadc_lsnappsattributes.tf_lsnappsattributes.name
	depends_on         = [citrixadc_lsnappsprofile_port_binding.tf_lsnappsprofile_port_binding]
}

data "citrixadc_lsnappsprofile_lsnappsattributes_binding" "tf_lsnappsprofile_lsnappsattributes_binding" {
	appsprofilename    = citrixadc_lsnappsprofile_lsnappsattributes_binding.tf_lsnappsprofile_lsnappsattributes_binding.appsprofilename
	appsattributesname = citrixadc_lsnappsprofile_lsnappsattributes_binding.tf_lsnappsprofile_lsnappsattributes_binding.appsattributesname
	depends_on         = [citrixadc_lsnappsprofile_lsnappsattributes_binding.tf_lsnappsprofile_lsnappsattributes_binding]
}
`

func TestAccLsnappsprofile_lsnappsattributes_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsprofile_lsnappsattributes_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsprofile_lsnappsattributes_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsprofile_lsnappsattributes_bindingExist("citrixadc_lsnappsprofile_lsnappsattributes_binding.tf_lsnappsprofile_lsnappsattributes_binding", nil),
				),
			},
			{
				Config: testAccLsnappsprofile_lsnappsattributes_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsprofile_lsnappsattributes_bindingNotExist("citrixadc_lsnappsprofile_lsnappsattributes_binding.tf_lsnappsprofile_lsnappsattributes_binding", "my_lsn_profile,my_lsn_appattributes"),
				),
			},
		},
	})
}

func testAccCheckLsnappsprofile_lsnappsattributes_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnappsprofile_lsnappsattributes_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"appsprofilename", "appsattributesname"}, nil)
		if err != nil {
			return err
		}
		appsprofilename := idMap["appsprofilename"]
		appsattributesname := idMap["appsattributesname"]

		findParams := service.FindParams{
			ResourceType:             "lsnappsprofile_lsnappsattributes_binding",
			ResourceName:             appsprofilename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching appsattributesname
		found := false
		for _, v := range dataArr {
			if v["appsattributesname"].(string) == appsattributesname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsnappsprofile_lsnappsattributes_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnappsprofile_lsnappsattributes_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"appsprofilename", "appsattributesname"}, nil)
		if err != nil {
			return err
		}
		appsprofilename := idMap["appsprofilename"]
		appsattributesname := idMap["appsattributesname"]

		findParams := service.FindParams{
			ResourceType:             "lsnappsprofile_lsnappsattributes_binding",
			ResourceName:             appsprofilename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching appsattributesname
		found := false
		for _, v := range dataArr {
			if v["appsattributesname"].(string) == appsattributesname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsnappsprofile_lsnappsattributes_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsnappsprofile_lsnappsattributes_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnappsprofile_lsnappsattributes_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"appsprofilename", "appsattributesname"}, nil)
		if err != nil {
			return err
		}
		appsprofilename := idMap["appsprofilename"]

		_, err = client.FindResource("lsnappsprofile_lsnappsattributes_binding", appsprofilename)
		if err == nil {
			return fmt.Errorf("lsnappsprofile_lsnappsattributes_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLsnappsprofile_lsnappsattributes_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsprofile_lsnappsattributes_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile_lsnappsattributes_binding.tf_lsnappsprofile_lsnappsattributes_binding", "appsprofilename", "my_lsn_profile"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile_lsnappsattributes_binding.tf_lsnappsprofile_lsnappsattributes_binding", "appsattributesname", "my_lsn_appattributes"),
				),
			},
		},
	})
}
