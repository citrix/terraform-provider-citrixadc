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
	"net/url"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccIpset_nsip_binding_basic = `

	resource "citrixadc_ipset_nsip_binding" "tf_ipset_nsip_binding" {
		name    = citrixadc_ipset.tf_ipset.name
		ipaddress = citrixadc_nsip.tf_nsip.ipaddress
	}
	
	
	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}
	
	resource "citrixadc_nsip" "tf_nsip" {
		ipaddress = "10.1.1.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
`

const testAccIpset_nsip_binding_basic_step2 = `
 
	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}
	
	resource "citrixadc_nsip" "tf_nsip" {
		ipaddress = "10.1.1.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
`

const testAccIpset_nsip_bindingDataSource_basic = `

	resource "citrixadc_ipset_nsip_binding" "tf_ipset_nsip_binding" {
		name    = citrixadc_ipset.tf_ipset.name
		ipaddress = citrixadc_nsip.tf_nsip.ipaddress
	}
	
	
	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}
	
	resource "citrixadc_nsip" "tf_nsip" {
		ipaddress = "10.1.1.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}

	data "citrixadc_ipset_nsip_binding" "tf_ipset_nsip_binding" {
		name       = citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding.name
		ipaddress  = citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding.ipaddress
		depends_on = [citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding]
	}
`

func TestAccIpset_nsip_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpset_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpset_nsip_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpset_nsip_bindingExist("citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding", nil),
				),
			},
			{
				Config: testAccIpset_nsip_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpset_nsip_bindingNotExist("citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding", "tf_test_ipset,10.1.1.1"),
				),
			},
		},
	})
}

func testAccCheckIpset_nsip_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ipset_nsip_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "ipaddress"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %v: %v", bindingId, err)
		}
		name := idMap["name"]
		ipaddress := idMap["ipaddress"]

		findParams := service.FindParams{
			ResourceType:             "ipset_nsip_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ipaddress
		found := false
		for _, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("ipset_nsip_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckIpset_nsip_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		name := idSlice[0]
		ipaddress := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "ipset_nsip_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ipaddress
		found := false
		for _, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("ipset_nsip_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckIpset_nsip_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ipset_nsip_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Ipset_nsip_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("ipset_nsip_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccIpset_nsip_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccIpset_nsip_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding", "name", "tf_test_ipset"),
					resource.TestCheckResourceAttr("data.citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding", "ipaddress", "10.1.1.1"),
				),
			},
		},
	})
}

const testAccIpset_nsip_binding_upgrade_basic = `

	resource "citrixadc_ipset_nsip_binding" "tf_ipset_nsip_binding" {
		name    = citrixadc_ipset.tf_ipset.name
		ipaddress = citrixadc_nsip.tf_nsip.ipaddress
	}


	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}

	resource "citrixadc_nsip" "tf_nsip" {
		ipaddress = "10.1.1.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
`

func TestAccIpset_nsip_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckIpset_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy comma-joined id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccIpset_nsip_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpset_nsip_bindingExist("citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding", "id", "tf_test_ipset,10.1.1.1"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccIpset_nsip_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpset_nsip_bindingExist("citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding", "id", "name:tf_test_ipset,ipaddress:10.1.1.1"),
				),
			},
		},
	})
}

func TestAccIpset_nsip_binding_import(t *testing.T) {
	const resAddr = "citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,ipaddress) so it matches exactly what SDK v2 wrote.
	legacyID := func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resAddr]
		if !ok {
			return "", fmt.Errorf("resource not found in state: %s", resAddr)
		}
		kv := map[string]string{}
		for _, p := range strings.Split(rs.Primary.ID, ",") {
			if i := strings.Index(p, ":"); i >= 0 {
				v, _ := url.QueryUnescape(p[i+1:])
				kv[p[:i]] = v
			}
		}
		ordr := []string{"name", "ipaddress"}
		parts := make([]string, 0, len(ordr))
		for _, k := range ordr {
			if v, ok := kv[k]; ok {
				parts = append(parts, v)
			}
		}
		// Fallback: a positional (non key:value) id has no key:value parts to reorder; import it as-is.
		if len(parts) == 0 {
			return rs.Primary.ID, nil
		}
		return strings.Join(parts, ","), nil
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpset_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpset_nsip_binding_basic,
			},
			{
				Config:                  testAccIpset_nsip_binding_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
			{Config: testAccIpset_nsip_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

// TestAccIpset_nsip_binding_selfHealing verifies drift recovery: after the binding is
// created, it is deleted out-of-band on the ADC; the next apply of the same config must
// detect the missing binding and recreate it.
func TestAccIpset_nsip_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_ipset_nsip_binding.tf_ipset_nsip_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpset_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpset_nsip_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpset_nsip_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Ipset_nsip_binding.Type(), "tf_test_ipset", []string{"ipaddress:" + utils.UrlEncode("10.1.1.1")}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccIpset_nsip_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpset_nsip_bindingExist(resAddr, nil)),
			},
		},
	})
}
