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

// Participating-entity config (citrixadc_lsnpool) is lifted from
// citrixadc_framework/acctests/lsnpool_test.go (testAccLsnpool_basic).
// ownernode is intentionally NOT set: tests target the standalone testbed (not cluster).

const testAccLsnpool_lsnip_binding_basic_step1 = `
	resource "citrixadc_lsnpool" "tf_lsnpool" {
		poolname            = "my_lsn_pool"
		nattype             = "DYNAMIC"
		portblockallocation = "DISABLED"
		maxportrealloctmq   = 50
		portrealloctimeout  = 50
	}

	resource "citrixadc_lsnpool_lsnip_binding" "tf_lsnpool_lsnip_binding" {
		poolname   = citrixadc_lsnpool.tf_lsnpool.poolname
		lsnip      = "10.20.30.40-10.20.30.50"
		depends_on = [citrixadc_lsnpool.tf_lsnpool]
	}

`

const testAccLsnpool_lsnip_binding_basic_step2 = `
	# Keep the parent lsnpool without the actual binding to verify proper deletion
	resource "citrixadc_lsnpool" "tf_lsnpool" {
		poolname            = "my_lsn_pool"
		nattype             = "DYNAMIC"
		portblockallocation = "DISABLED"
		maxportrealloctmq   = 50
		portrealloctimeout  = 50
	}

`

func TestAccLsnpool_lsnip_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnpool_lsnip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnpool_lsnip_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnpool_lsnip_bindingExist("citrixadc_lsnpool_lsnip_binding.tf_lsnpool_lsnip_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnpool_lsnip_binding.tf_lsnpool_lsnip_binding", "poolname", "my_lsn_pool"),
					resource.TestCheckResourceAttr("citrixadc_lsnpool_lsnip_binding.tf_lsnpool_lsnip_binding", "lsnip", "10.20.30.40-10.20.30.50"),
				),
			},
			{
				Config: testAccLsnpool_lsnip_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnpool_lsnip_bindingNotExist("citrixadc_lsnpool_lsnip_binding.tf_lsnpool_lsnip_binding", "poolname:my_lsn_pool,lsnip:10.20.30.40-10.20.30.50"),
				),
			},
		},
	})
}

// lsnpoolLsnipBindingAggregateReadForTest reads the bound IPs via the AGGREGATE
// parent endpoint (lsnpool_binding/<poolname>) and flattens the nested
// "lsnpool_lsnip_binding" arrays. The direct lsnpool_lsnip_binding endpoint
// returns a keyless empty body on NS14.1, so the check helpers must use the
// aggregate parent endpoint to locate the binding.
func lsnpoolLsnipBindingAggregateReadForTest(client *service.NitroClient, poolname string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType:             "lsnpool_binding",
		ResourceName:             poolname,
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		nested, ok := parent["lsnpool_lsnip_binding"]
		if !ok || nested == nil {
			continue
		}
		nestedArr, ok := nested.([]interface{})
		if !ok {
			continue
		}
		for _, item := range nestedArr {
			if m, ok := item.(map[string]interface{}); ok {
				rows = append(rows, m)
			}
		}
	}
	return rows, nil
}

func testAccCheckLsnpool_lsnip_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnpool_lsnip_binding id is set")
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

		// Composite ID: poolname:UrlEncode(value),lsnip:UrlEncode(value)
		// Legacy attr order is [poolname, lsnip].
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"poolname", "lsnip"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		poolname := idMap["poolname"]
		lsnip := idMap["lsnip"]

		dataArr, err := lsnpoolLsnipBindingAggregateReadForTest(client, poolname)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching lsnip
		found := false
		for _, v := range dataArr {
			if val, ok := v["lsnip"].(string); ok && val == lsnip {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsnpool_lsnip_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnpool_lsnip_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"poolname", "lsnip"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}

		poolname := idMap["poolname"]
		lsnip := idMap["lsnip"]

		dataArr, err := lsnpoolLsnipBindingAggregateReadForTest(client, poolname)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching lsnip
		found := false
		for _, v := range dataArr {
			if val, ok := v["lsnip"].(string); ok && val == lsnip {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsnpool_lsnip_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsnpool_lsnip_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnpool_lsnip_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"poolname", "lsnip"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		poolname := idMap["poolname"]
		lsnip := idMap["lsnip"]

		dataArr, err := lsnpoolLsnipBindingAggregateReadForTest(client, poolname)

		// If the parent pool itself is gone, the binding is certainly gone.
		if err != nil {
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["lsnip"].(string); ok && val == lsnip {
				return fmt.Errorf("lsnpool_lsnip_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

func TestAccLsnpool_lsnip_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lsnpool_lsnip_binding.tf_lsnpool_lsnip_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnpool_lsnip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnpool_lsnip_binding_basic_step1,
			},
			{
				Config:                  testAccLsnpool_lsnip_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccLsnpool_lsnip_binding_DataSource_basic = `
	resource "citrixadc_lsnpool" "tf_lsnpool" {
		poolname            = "my_lsn_pool"
		nattype             = "DYNAMIC"
		portblockallocation = "DISABLED"
		maxportrealloctmq   = 50
		portrealloctimeout  = 50
	}

	resource "citrixadc_lsnpool_lsnip_binding" "tf_lsnpool_lsnip_binding" {
		poolname   = citrixadc_lsnpool.tf_lsnpool.poolname
		lsnip      = "10.20.30.40-10.20.30.50"
		depends_on = [citrixadc_lsnpool.tf_lsnpool]
	}

	data "citrixadc_lsnpool_lsnip_binding" "tf_lsnpool_lsnip_binding" {
		poolname   = citrixadc_lsnpool.tf_lsnpool.poolname
		lsnip      = "10.20.30.40-10.20.30.50"
		depends_on = [citrixadc_lsnpool_lsnip_binding.tf_lsnpool_lsnip_binding]
	}

`

func TestAccLsnpool_lsnip_binding_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnpool_lsnip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnpool_lsnip_binding_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnpool_lsnip_binding.tf_lsnpool_lsnip_binding", "poolname", "my_lsn_pool"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnpool_lsnip_binding.tf_lsnpool_lsnip_binding", "lsnip", "10.20.30.40-10.20.30.50"),
				),
			},
		},
	})
}
