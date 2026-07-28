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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// endpointinfo is a STANDARD CRUD resource whose Read is a get-all-and-filter
// against endpointkind + endpointname. The composite ID is
// "endpointkind:<kind>,endpointname:<name>". For an IP endpoint, endpointname
// is an IP address.
//
// No feature-enable is required: verified against a live ADC that
// POST/GET/PUT/DELETE on /config/endpointinfo all succeed without enabling any
// ns feature.
//
// IMPORTANT (server normalization): the ADC always pads endpointmetadata to six
// dotted qualifiers with wildcards, e.g. "cluster.default.frontend" is stored
// (and read back) as "cluster.default.frontend.*.*.*". Because Read reflects the
// authoritative server value into state, the fixture MUST supply the fully
// padded 6-qualifier form so that plan == apply == state and no
// "inconsistent result after apply" occurs.

const testAccEndpointinfo_basic_step1 = `
resource "citrixadc_endpointinfo" "tf_endpointinfo" {
  endpointname       = "10.222.74.100"
  endpointkind       = "IP"
  endpointmetadata   = "cluster.default.frontend.*.*.*"
  endpointlabelsjson = "{\"env\":\"test\"}"
}

`

const testAccEndpointinfo_basic_step2 = `
resource "citrixadc_endpointinfo" "tf_endpointinfo" {
  endpointname       = "10.222.74.100"
  endpointkind       = "IP"
  endpointmetadata   = "cluster.default.backend.*.*.*"
  endpointlabelsjson = "{\"env\":\"prod\"}"
}

`

func TestAccEndpointinfo_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEndpointinfoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointinfo_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointinfoExist("citrixadc_endpointinfo.tf_endpointinfo", nil),
					resource.TestCheckResourceAttr("citrixadc_endpointinfo.tf_endpointinfo", "endpointname", "10.222.74.100"),
					resource.TestCheckResourceAttr("citrixadc_endpointinfo.tf_endpointinfo", "endpointkind", "IP"),
					resource.TestCheckResourceAttr("citrixadc_endpointinfo.tf_endpointinfo", "endpointmetadata", "cluster.default.frontend.*.*.*"),
					resource.TestCheckResourceAttr("citrixadc_endpointinfo.tf_endpointinfo", "endpointlabelsjson", "{\"env\":\"test\"}"),
				),
			},
			{
				// Update the updateable attributes only. endpointkind is
				// RequiresReplace, so it is intentionally kept unchanged here.
				Config: testAccEndpointinfo_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointinfoExist("citrixadc_endpointinfo.tf_endpointinfo", nil),
					resource.TestCheckResourceAttr("citrixadc_endpointinfo.tf_endpointinfo", "endpointname", "10.222.74.100"),
					resource.TestCheckResourceAttr("citrixadc_endpointinfo.tf_endpointinfo", "endpointkind", "IP"),
					resource.TestCheckResourceAttr("citrixadc_endpointinfo.tf_endpointinfo", "endpointmetadata", "cluster.default.backend.*.*.*"),
					resource.TestCheckResourceAttr("citrixadc_endpointinfo.tf_endpointinfo", "endpointlabelsjson", "{\"env\":\"prod\"}"),
				),
			},
		},
	})
}

func testAccCheckEndpointinfoExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No endpointinfo id is set")
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

		// endpointinfo Read is a get-all-and-filter by endpointkind + endpointname.
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		endpointkind := idMap["endpointkind"]
		endpointname := idMap["endpointname"]

		findParams := service.FindParams{
			ResourceType:             service.Endpointinfo.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if kind, ok := v["endpointkind"].(string); ok && kind == endpointkind {
				if name, ok := v["endpointname"].(string); ok && name == endpointname {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("endpointinfo %s not found", n)
		}

		return nil
	}
}

func testAccCheckEndpointinfoDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_endpointinfo" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No endpointinfo id is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		endpointkind := idMap["endpointkind"]
		endpointname := idMap["endpointname"]

		findParams := service.FindParams{
			ResourceType:             service.Endpointinfo.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// A missing resource (empty get-all) means it was destroyed.
			continue
		}

		found := false
		for _, v := range dataArr {
			if kind, ok := v["endpointkind"].(string); ok && kind == endpointkind {
				if name, ok := v["endpointname"].(string); ok && name == endpointname {
					found = true
					break
				}
			}
		}

		if found {
			return fmt.Errorf("endpointinfo %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

func TestAccEndpointinfo_import(t *testing.T) {
	const resAddr = "citrixadc_endpointinfo.tf_endpointinfo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEndpointinfoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointinfo_basic_step1,
			},
			{
				Config:                  testAccEndpointinfo_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccEndpointinfoDataSource_basic = `

resource "citrixadc_endpointinfo" "tf_endpointinfo" {
  endpointname       = "10.222.74.100"
  endpointkind       = "IP"
  endpointmetadata   = "cluster.default.frontend.*.*.*"
  endpointlabelsjson = "{\"env\":\"test\"}"
}

data "citrixadc_endpointinfo" "tf_endpointinfo" {
  endpointkind = citrixadc_endpointinfo.tf_endpointinfo.endpointkind
  endpointname = citrixadc_endpointinfo.tf_endpointinfo.endpointname
  depends_on   = [citrixadc_endpointinfo.tf_endpointinfo]
}
`

func TestAccEndpointinfoDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointinfoDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_endpointinfo.tf_endpointinfo", "endpointkind", "IP"),
					resource.TestCheckResourceAttr("data.citrixadc_endpointinfo.tf_endpointinfo", "endpointname", "10.222.74.100"),
					resource.TestCheckResourceAttr("data.citrixadc_endpointinfo.tf_endpointinfo", "endpointmetadata", "cluster.default.frontend.*.*.*"),
					resource.TestCheckResourceAttr("data.citrixadc_endpointinfo.tf_endpointinfo", "endpointlabelsjson", "{\"env\":\"test\"}"),
				),
			},
		},
	})
}
