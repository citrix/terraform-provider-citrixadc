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

const testAccAaauser_intranetip6_binding_basic = `
	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_aaauser_intranetip6_binding" "tf_aaauser_intranetip6_binding" {
		username    = citrixadc_aaauser.tf_aaauser.username
		intranetip6 = "2003:db8:100::fb/128"
	}
`

const testAccAaauser_intranetip6_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
`

const testAccAaauser_intranetip6_bindingDataSource_basic = `
	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_aaauser_intranetip6_binding" "tf_aaauser_intranetip6_binding" {
		username    = citrixadc_aaauser.tf_aaauser.username
		intranetip6 = "2003:db8:100::fb/128"
		numaddr     = 1
	}

	data "citrixadc_aaauser_intranetip6_binding" "tf_aaauser_intranetip6_binding" {
		username    = citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding.username
		intranetip6 = citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding.intranetip6
		depends_on = [citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding]
	}
`

func TestAccAaauser_intranetip6_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaauser_intranetip6_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaauser_intranetip6_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_intranetip6_bindingExist("citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding", nil),
				),
			},
			{
				Config: testAccAaauser_intranetip6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_intranetip6_bindingNotExist("citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding", "user1,2003:db8:100::fb/128"),
				),
			},
		},
	})
}

func testAccCheckAaauser_intranetip6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaauser_intranetip6_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"username", "intranetip6"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		username := idMap["username"]
		intranetip6 := idMap["intranetip6"]

		findParams := service.FindParams{
			ResourceType:             "aaauser_intranetip6_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching intranetip6
		found := false
		for _, v := range dataArr {
			if v["intranetip6"].(string) == intranetip6 {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("aaauser_intranetip6_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaauser_intranetip6_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		username := idSlice[0]
		intranetip6 := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaauser_intranetip6_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching intranetip6
		found := false
		for _, v := range dataArr {
			if v["intranetip6"].(string) == intranetip6 {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("aaauser_intranetip6_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaauser_intranetip6_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaauser_intranetip6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Aaauser_intranetip6_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaauser_intranetip6_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAaauser_intranetip6_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaauser_intranetip6_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding", "username", "user1"),
					resource.TestCheckResourceAttr("data.citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding", "intranetip6", "2003:db8:100::fb/128"),
					resource.TestCheckResourceAttr("data.citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding", "numaddr", "1"),
				),
			},
		},
	})
}

const testAccAaauser_intranetip6_binding_upgrade_basic = `
	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_aaauser_intranetip6_binding" "tf_aaauser_intranetip6_binding" {
		username    = citrixadc_aaauser.tf_aaauser.username
		intranetip6 = "2003:db8:100::fb/128"
	}
`

func TestAccAaauser_intranetip6_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAaauser_intranetip6_bindingDestroy,
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
				Config: testAccAaauser_intranetip6_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_intranetip6_bindingExist("citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding", "id", "user1,2003:db8:100::fb/128"),
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
				Config: testAccAaauser_intranetip6_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_intranetip6_bindingExist("citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding", "id", "intranetip6:2003%3Adb8%3A100%3A%3Afb%2F128,username:user1"),
				),
			},
		},
	})
}

func TestAccAaauser_intranetip6_binding_import(t *testing.T) {
	const resAddr = "citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: username,intranetip6) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"username", "intranetip6"}
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
		CheckDestroy:             testAccCheckAaauser_intranetip6_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAaauser_intranetip6_binding_basic},
			{Config: testAccAaauser_intranetip6_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccAaauser_intranetip6_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccAaauser_intranetip6_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaauser_intranetip6_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaauser_intranetip6_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaauser_intranetip6_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Aaauser_intranetip6_binding.Type(), "user1", map[string]string{"intranetip6": url.PathEscape("2003:db8:100::fb/128")}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAaauser_intranetip6_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaauser_intranetip6_bindingExist(resAddr, nil)),
			},
		},
	})
}
