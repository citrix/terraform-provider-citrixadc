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

const testAccLsngroup_lsnlogprofile_binding_basic = `
	resource "citrixadc_lsnclient" "tf_lsnclient" {
		clientname = "my_lsnclient"
	}

	resource "citrixadc_lsngroup" "tf_lsngroup" {
		groupname     = "my_lsngroup"
		clientname    = resource.citrixadc_lsnclient.tf_lsnclient.clientname
		logging       = "DISABLED"
		nattype       = "DYNAMIC"
		snmptraplimit = 50
	}

	resource "citrixadc_lsnlogprofile" "tf_lsnlogprofile" {
		logprofilename = "my_lsn_logprofile"
		logsubscrinfo   = "ENABLED"
		logcompact      = "ENABLED"
		logipfix        = "ENABLED"
	}

resource "citrixadc_lsngroup_lsnlogprofile_binding" "tf_lsngroup_lsnlogprofile_binding" {
	groupname      = citrixadc_lsngroup.tf_lsngroup.groupname
	logprofilename = citrixadc_lsnlogprofile.tf_lsnlogprofile.logprofilename
}
  
`

const testAccLsngroup_lsnlogprofile_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

const testAccLsngroup_lsnlogprofile_bindingDataSource_basic = `

	resource "citrixadc_lsnclient" "tf_lsnclient" {
		clientname = "my_lsnclient"
	}

	resource "citrixadc_lsngroup" "tf_lsngroup" {
		groupname     = "my_lsngroup"
		clientname    = resource.citrixadc_lsnclient.tf_lsnclient.clientname
		logging       = "DISABLED"
		nattype       = "DYNAMIC"
		snmptraplimit = 50
	}

	resource "citrixadc_lsnlogprofile" "tf_lsnlogprofile" {
		logprofilename = "my_lsn_logprofile"
		logsubscrinfo   = "ENABLED"
		logcompact      = "ENABLED"
		logipfix        = "ENABLED"
	}

resource "citrixadc_lsngroup_lsnlogprofile_binding" "tf_lsngroup_lsnlogprofile_binding" {
	groupname      = citrixadc_lsngroup.tf_lsngroup.groupname
	logprofilename = citrixadc_lsnlogprofile.tf_lsnlogprofile.logprofilename
}

data "citrixadc_lsngroup_lsnlogprofile_binding" "tf_lsngroup_lsnlogprofile_binding" {
	groupname      = citrixadc_lsngroup_lsnlogprofile_binding.tf_lsngroup_lsnlogprofile_binding.groupname
	logprofilename = citrixadc_lsngroup_lsnlogprofile_binding.tf_lsngroup_lsnlogprofile_binding.logprofilename
}
`

func TestAccLsngroup_lsnlogprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_lsnlogprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnlogprofile_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnlogprofile_bindingExist("citrixadc_lsngroup_lsnlogprofile_binding.tf_lsngroup_lsnlogprofile_binding", nil),
				),
			},
			{
				Config: testAccLsngroup_lsnlogprofile_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnlogprofile_bindingNotExist("citrixadc_lsngroup_lsnlogprofile_binding.tf_lsngroup_lsnlogprofile_binding", "my_lsn_group,my_lsn_logprofile"),
				),
			},
		},
	})
}

func testAccCheckLsngroup_lsnlogprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsngroup_lsnlogprofile_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"groupname", "logprofilename"}, nil)
		if err != nil {
			return err
		}
		groupname := idMap["groupname"]
		logprofilename := idMap["logprofilename"]

		findParams := service.FindParams{
			ResourceType:             "lsngroup_lsnlogprofile_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching logprofilename
		found := false
		for _, v := range dataArr {
			if v["logprofilename"].(string) == logprofilename {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsngroup_lsnlogprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_lsnlogprofile_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"groupname", "logprofilename"}, nil)
		if err != nil {
			return err
		}
		groupname := idMap["groupname"]
		logprofilename := idMap["logprofilename"]

		findParams := service.FindParams{
			ResourceType:             "lsngroup_lsnlogprofile_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching logprofilename
		found := false
		for _, v := range dataArr {
			if v["logprofilename"].(string) == logprofilename {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsngroup_lsnlogprofile_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_lsnlogprofile_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsngroup_lsnlogprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsngroup_lsnlogprofile_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsngroup_lsnlogprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLsngroup_lsnlogprofile_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnlogprofile_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_lsnlogprofile_binding.tf_lsngroup_lsnlogprofile_binding", "groupname", "my_lsngroup"),
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_lsnlogprofile_binding.tf_lsngroup_lsnlogprofile_binding", "logprofilename", "my_lsn_logprofile"),
				),
			},
		},
	})
}

// testAccLsngroup_lsnlogprofile_binding_upgrade_basic reuses the _basic config
// values. It is valid under BOTH the last SDK v2 schema (2.2.0) and the current
// Framework schema (the migration restored the SDK v2 attribute names).
const testAccLsngroup_lsnlogprofile_binding_upgrade_basic = `
	resource "citrixadc_lsnclient" "tf_lsnclient" {
		clientname = "my_lsnclient"
	}

	resource "citrixadc_lsngroup" "tf_lsngroup" {
		groupname     = "my_lsngroup"
		clientname    = resource.citrixadc_lsnclient.tf_lsnclient.clientname
		logging       = "DISABLED"
		nattype       = "DYNAMIC"
		snmptraplimit = 50
	}

	resource "citrixadc_lsnlogprofile" "tf_lsnlogprofile" {
		logprofilename = "my_lsn_logprofile"
		logsubscrinfo   = "ENABLED"
		logcompact      = "ENABLED"
		logipfix        = "ENABLED"
	}

	resource "citrixadc_lsngroup_lsnlogprofile_binding" "tf_lsngroup_lsnlogprofile_binding" {
		groupname      = citrixadc_lsngroup.tf_lsngroup.groupname
		logprofilename = citrixadc_lsnlogprofile.tf_lsnlogprofile.logprofilename
	}
`

// TestAccLsngroup_lsnlogprofile_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release is correctly upgraded when the same config is subsequently
// managed by the current Framework provider. Step 1 creates the binding with
// citrix/citrixadc 2.2.0 (writes the legacy comma id "my_lsngroup,my_lsn_logprofile" —
// the SDK v2 d.SetId(fmt.Sprintf("%s,%s", groupname, logprofilename))). Step 2 refreshes/
// plans/applies the same config through the Framework provider, exercising ParseIdString
// on the legacy id; the Framework recomputes the id on Read (SetAttrFromGet), so the
// canonical new-format id becomes "groupname:my_lsngroup,logprofilename:my_lsn_logprofile".
func TestAccLsngroup_lsnlogprofile_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_lsngroup_lsnlogprofile_binding.tf_lsngroup_lsnlogprofile_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsngroup_lsnlogprofile_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLsngroup_lsnlogprofile_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnlogprofile_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "my_lsngroup,my_lsn_logprofile"),
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
				Config: testAccLsngroup_lsnlogprofile_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnlogprofile_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "groupname:my_lsngroup,logprofilename:my_lsn_logprofile"),
				),
			},
		},
	})
}

func TestAccLsngroup_lsnlogprofile_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lsngroup_lsnlogprofile_binding.tf_lsngroup_lsnlogprofile_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: groupname,logprofilename) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"groupname", "logprofilename"}
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
		CheckDestroy:             testAccCheckLsngroup_lsnlogprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsngroup_lsnlogprofile_binding_basic},
			{Config: testAccLsngroup_lsnlogprofile_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccLsngroup_lsnlogprofile_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccLsngroup_lsnlogprofile_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lsngroup_lsnlogprofile_binding.tf_lsngroup_lsnlogprofile_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_lsnlogprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnlogprofile_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsngroup_lsnlogprofile_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Lsngroup_lsnlogprofile_binding.Type(), "my_lsngroup", map[string]string{"logprofilename": "my_lsn_logprofile"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsngroup_lsnlogprofile_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsngroup_lsnlogprofile_bindingExist(resAddr, nil)),
			},
		},
	})
}
