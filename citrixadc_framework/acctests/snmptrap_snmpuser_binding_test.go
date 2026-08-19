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

const testAccSnmptrap_snmpuser_binding_basic = `
	resource "citrixadc_snmpuser" "tf_snmpuser" {
	name       = "tf_snmpuser"
	group      = "all_group"
	authtype   = "SHA"
	authpasswd = "secretpassword"
	privtype   = "AES"
	privpasswd = "secretpassword"
	}
	resource "citrixadc_snmptrap" "tf_snmptrap" {
	trapclass       = "generic"
	trapdestination = "10.50.50.10"
	version         = "V3"
	}
	resource "citrixadc_snmptrap_snmpuser_binding" "tf_binding" {
	trapclass       = citrixadc_snmptrap.tf_snmptrap.trapclass
	trapdestination = citrixadc_snmptrap.tf_snmptrap.trapdestination
	username        = citrixadc_snmpuser.tf_snmpuser.name
	securitylevel   = "authPriv"
	}
`

// const testAccSnmptrap_snmpuser_binding_basic_step2 = `
// 	# Keep the above bound resources without the actual binding to check proper deletion
// 	resource "citrixadc_snmpuser" "tf_snmpuser" {
// 		name       = "tf_snmpuser"
// 		group      = "all_group"
// 		authtype   = "SHA"
// 		authpasswd = "secretpassword"
// 		privtype   = "AES"
// 		privpasswd = "secretpassword"
// 	}
// 	resource "citrixadc_snmptrap" "tf_snmptrap" {
// 		trapclass       = "generic"
// 		trapdestination = "10.50.50.10"
// 		version         = "V3"
// 		td = 0
// 	}
// `

func TestAccSnmptrap_snmpuser_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmptrap_snmpuser_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmptrap_snmpuser_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmptrap_snmpuser_bindingExist("citrixadc_snmptrap_snmpuser_binding.tf_binding", nil),
				),
			},
			// resource.TestStep{
			// 	Config: testAccSnmptrap_snmpuser_binding_basic_step2,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckSnmptrap_snmpuser_bindingNotExist("citrixadc_snmptrap_snmpuser_binding.tf_binding", "generic,10.50.50.10,tf_snmpuser"),
			// 	),
			// },
		},
	})
}

func testAccCheckSnmptrap_snmpuser_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmptrap_snmpuser_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"trapclass", "trapdestination", "username"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		trapclass := idMap["trapclass"]
		trapdestination := idMap["trapdestination"]
		username := idMap["username"]

		args := make(map[string]string, 0)
		args["trapclass"] = trapclass
		args["trapdestination"] = trapdestination
		args["version"] = rs.Primary.Attributes["version"]
		args["td"] = rs.Primary.Attributes["td"]

		findParams := service.FindParams{
			ResourceType:             "snmptrap_snmpuser_binding",
			ArgsMap:                  args,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["username"].(string) == username {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("snmptrap_snmpuser_binding %s not found", n)
		}

		return nil
	}
}

// FIXME: Know how to access the other attributes of the resource other than attributes that are included in the ID of the resource
// func testAccCheckSnmptrap_snmpuser_bindingNotExist(n string, id string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		Use the shared utility function to get a configured client
// 		client, err := testAccGetFrameworkClient()
// 		if err != nil {
// 			return fmt.Errorf("Failed to get test client: %v", err)
// 		}
// 		rs, _ := s.RootModule().Resources[n]
// 		if !strings.Contains(id, ",") {
// 			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
// 		}

// 		bindingId := rs.Primary.ID

// 		idSlice := strings.SplitN(bindingId, ",", 3)

// 		trapclass := idSlice[0]
// 		trapdestination := idSlice[1]
// 		username := idSlice[2]

// 		args := make(map[string]string, 0)
// 		args["trapclass"] = trapclass
// 		args["trapdestination"] = trapdestination
// 		args["version"] = rs.Primary.Attributes["version"]
// 		args["td"] = rs.Primary.Attributes["td"]

// 		findParams := service.FindParams{
// 			ResourceType:             "snmptrap_snmpuser_binding",
// 			ArgsMap:                  args,
// 			ResourceMissingErrorCode: 258,
// 		}
// 		dataArr, err := client.FindResourceArrayWithParams(findParams)

// 		// Unexpected error
// 		if err != nil {
// 			return err
// 		}

// 		// Iterate through results to hopefully not find the one with the matching secondIdComponent
// 		found := false
// 		for _, v := range dataArr {
// 			if v["username"].(string) == username {
// 				found = true
// 				break
// 			}
// 		}

// 		if found {
// 			return fmt.Errorf("snmptrap_snmpuser_binding %s was found, but it should have been destroyed", n)
// 		}

// 		return nil
// 	}
// }

func testAccCheckSnmptrap_snmpuser_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_snmptrap_snmpuser_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Snmptrap_snmpuser_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("snmptrap_snmpuser_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSnmptrap_snmpuser_bindingDataSource_basic = `
	resource "citrixadc_snmpuser" "tf_snmpuser" {
		name       = "tf_snmpuser_ds"
		group      = "all_group"
		authtype   = "SHA"
		authpasswd = "secretpassword"
		privtype   = "AES"
		privpasswd = "secretpassword"
	}
	resource "citrixadc_snmptrap" "tf_snmptrap" {
		trapclass       = "generic"
		trapdestination = "10.50.50.11"
		version         = "V3"
	}
	resource "citrixadc_snmptrap_snmpuser_binding" "tf_binding" {
		trapclass       = citrixadc_snmptrap.tf_snmptrap.trapclass
		trapdestination = citrixadc_snmptrap.tf_snmptrap.trapdestination
		username        = citrixadc_snmpuser.tf_snmpuser.name
		securitylevel   = "authPriv"
	}

	data "citrixadc_snmptrap_snmpuser_binding" "tf_binding_ds" {
		trapclass       = citrixadc_snmptrap_snmpuser_binding.tf_binding.trapclass
		trapdestination = citrixadc_snmptrap_snmpuser_binding.tf_binding.trapdestination
		username        = citrixadc_snmptrap_snmpuser_binding.tf_binding.username
		td              = citrixadc_snmptrap_snmpuser_binding.tf_binding.td
		version         = citrixadc_snmptrap_snmpuser_binding.tf_binding.version
		depends_on      = [citrixadc_snmptrap_snmpuser_binding.tf_binding]
	}
`

func TestAccSnmptrap_snmpuser_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmptrap_snmpuser_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_snmptrap_snmpuser_binding.tf_binding_ds", "trapclass", "generic"),
					resource.TestCheckResourceAttr("data.citrixadc_snmptrap_snmpuser_binding.tf_binding_ds", "trapdestination", "10.50.50.11"),
					resource.TestCheckResourceAttr("data.citrixadc_snmptrap_snmpuser_binding.tf_binding_ds", "username", "tf_snmpuser_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_snmptrap_snmpuser_binding.tf_binding_ds", "securitylevel", "authPriv"),
				),
			},
		},
	})
}

// testAccSnmptrap_snmpuser_binding_upgrade_basic reuses the _basic config (binding +
// all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0 schema and the
// current Framework schema because the migration restored the SDK v2 attribute names.
const testAccSnmptrap_snmpuser_binding_upgrade_basic = `
	resource "citrixadc_snmpuser" "tf_snmpuser" {
	name       = "tf_snmpuser"
	group      = "all_group"
	authtype   = "SHA"
	authpasswd = "secretpassword"
	privtype   = "AES"
	privpasswd = "secretpassword"
	}
	resource "citrixadc_snmptrap" "tf_snmptrap" {
	trapclass       = "generic"
	trapdestination = "10.50.50.10"
	version         = "V3"
	}
	resource "citrixadc_snmptrap_snmpuser_binding" "tf_binding" {
	trapclass       = citrixadc_snmptrap.tf_snmptrap.trapclass
	trapdestination = citrixadc_snmptrap.tf_snmptrap.trapdestination
	username        = citrixadc_snmpuser.tf_snmpuser.name
	securitylevel   = "authPriv"
	}
`

// TestAccSnmptrap_snmpuser_binding_sdkv2StateUpgrade verifies that state written by the
// last SDK v2 release is correctly upgraded when the same config is subsequently managed
// by the current Framework provider. Step 1 creates the binding with citrix/citrixadc
// 2.2.0 (writes the legacy comma id "generic,10.50.50.10,tf_snmpuser" — the SDK v2
// d.SetId(fmt.Sprintf("%s,%s", trapclass+","+trapdestination, username))). Step 2
// refreshes/plans/applies the same config through the Framework provider, exercising
// ParseIdString on the legacy id; the Framework recomputes the id on Read
// (SetAttrFromGet), so the canonical new-format id becomes
// "td:0,trapclass:generic,trapdestination:10.50.50.10,username:tf_snmpuser,version:V3".
func TestAccSnmptrap_snmpuser_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_snmptrap_snmpuser_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSnmptrap_snmpuser_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccSnmptrap_snmpuser_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmptrap_snmpuser_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "generic,10.50.50.10,tf_snmpuser"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed on Read into the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSnmptrap_snmpuser_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmptrap_snmpuser_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "td:0,trapclass:generic,trapdestination:10.50.50.10,username:tf_snmpuser,version:V3"),
				),
			},
		},
	})
}

func TestAccSnmptrap_snmpuser_binding_import(t *testing.T) {
	const resAddr = "citrixadc_snmptrap_snmpuser_binding.tf_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: trapclass,trapdestination,username) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"trapclass", "trapdestination", "username"}
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
		CheckDestroy:             testAccCheckSnmptrap_snmpuser_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSnmptrap_snmpuser_binding_basic},
			{Config: testAccSnmptrap_snmpuser_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccSnmptrap_snmpuser_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccSnmptrap_snmpuser_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_snmptrap_snmpuser_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmptrap_snmpuser_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmptrap_snmpuser_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmptrap_snmpuser_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Snmptrap_snmpuser_binding.Type(), "generic", []string{"trapdestination:10.50.50.10", "username:tf_snmpuser", "version:V3", "td:0"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSnmptrap_snmpuser_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmptrap_snmpuser_bindingExist(resAddr, nil)),
			},
		},
	})
}
