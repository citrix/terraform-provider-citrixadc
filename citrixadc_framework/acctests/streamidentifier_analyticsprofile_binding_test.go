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

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// streamidentifier_analyticsprofile_binding is an immutable bind/unbind resource
// (name + analyticsprofile are both RequiresReplace), so there is no update step.
// Participating-entity config (streamselector -> streamidentifier, analyticsprofile)
// is reused from streamidentifier_test.go and analyticsprofile_test.go.
const testAccStreamidentifier_analyticsprofile_binding_basic_step1 = `
resource "citrixadc_streamselector" "tf_streamselector" {
  name = "my_streamselector"
  rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC"]
}

resource "citrixadc_streamidentifier" "tf_streamidentifier" {
  name         = "my_streamidentifier"
  selectorname = citrixadc_streamselector.tf_streamselector.name
  samplecount  = 10
  sort         = "CONNECTIONS"
  snmptrap     = "ENABLED"
  loglimit     = 500
  loginterval  = 60
  log          = "NONE"
}

resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
  name             = "my_analyticsprofile"
  type             = "timeseries"
}

resource "citrixadc_streamidentifier_analyticsprofile_binding" "tf_binding" {
  name             = citrixadc_streamidentifier.tf_streamidentifier.name
  analyticsprofile = citrixadc_analyticsprofile.tf_analyticsprofile.name

  depends_on = [
    citrixadc_streamidentifier.tf_streamidentifier,
    citrixadc_analyticsprofile.tf_analyticsprofile,
  ]
}
`

func TestAccStreamidentifier_analyticsprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckStreamidentifier_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStreamidentifier_analyticsprofile_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStreamidentifier_analyticsprofile_bindingExist("citrixadc_streamidentifier_analyticsprofile_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier_analyticsprofile_binding.tf_binding", "name", "my_streamidentifier"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier_analyticsprofile_binding.tf_binding", "analyticsprofile", "my_analyticsprofile"),
				),
			},
		},
	})
}

func testAccCheckStreamidentifier_analyticsprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No streamidentifier_analyticsprofile_binding ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %q: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		analyticsprofile := idMap["analyticsprofile"]

		findParams := service.FindParams{
			ResourceType:             service.Streamidentifier_analyticsprofile_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["analyticsprofile"].(string); ok && val == analyticsprofile {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("streamidentifier_analyticsprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckStreamidentifier_analyticsprofile_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_streamidentifier_analyticsprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %q: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		analyticsprofile := idMap["analyticsprofile"]

		findParams := service.FindParams{
			ResourceType:             service.Streamidentifier_analyticsprofile_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// If the parent is gone the binding is gone too.
			continue
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["analyticsprofile"].(string); ok && val == analyticsprofile {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("streamidentifier_analyticsprofile_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccStreamidentifier_analyticsprofile_bindingDataSource_basic = `
resource "citrixadc_streamselector" "tf_streamselector" {
  name = "my_streamselector"
  rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC"]
}

resource "citrixadc_streamidentifier" "tf_streamidentifier" {
  name         = "my_streamidentifier"
  selectorname = citrixadc_streamselector.tf_streamselector.name
  samplecount  = 10
  sort         = "CONNECTIONS"
  snmptrap     = "ENABLED"
  loglimit     = 500
  loginterval  = 60
  log          = "NONE"
}

resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
  name             = "my_analyticsprofile"
  type             = "webinsight"
  httppagetracking = "DISABLED"
  httpurl          = "DISABLED"
}

resource "citrixadc_streamidentifier_analyticsprofile_binding" "tf_binding" {
  name             = citrixadc_streamidentifier.tf_streamidentifier.name
  analyticsprofile = citrixadc_analyticsprofile.tf_analyticsprofile.name

  depends_on = [
    citrixadc_streamidentifier.tf_streamidentifier,
    citrixadc_analyticsprofile.tf_analyticsprofile,
  ]
}

data "citrixadc_streamidentifier_analyticsprofile_binding" "tf_binding" {
  name             = citrixadc_streamidentifier_analyticsprofile_binding.tf_binding.name
  analyticsprofile = citrixadc_streamidentifier_analyticsprofile_binding.tf_binding.analyticsprofile

  depends_on = [citrixadc_streamidentifier_analyticsprofile_binding.tf_binding]
}
`

func TestAccStreamidentifier_analyticsprofile_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccStreamidentifier_analyticsprofile_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_streamidentifier_analyticsprofile_binding.tf_binding", "name", "my_streamidentifier"),
					resource.TestCheckResourceAttr("data.citrixadc_streamidentifier_analyticsprofile_binding.tf_binding", "analyticsprofile", "my_analyticsprofile"),
					resource.TestCheckResourceAttrSet("data.citrixadc_streamidentifier_analyticsprofile_binding.tf_binding", "id"),
				),
			},
		},
	})
}
