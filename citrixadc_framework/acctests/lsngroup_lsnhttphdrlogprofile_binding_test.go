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

const testAccLsngroup_lsnhttphdrlogprofile_binding_basic = `
resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsn_client"
}

resource "citrixadc_lsngroup" "tf_lsngroup" {
	groupname  = "my_lsn_group"
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
}

resource "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile" {
	httphdrlogprofilename = "my_httplogprofile"
}

resource "citrixadc_lsngroup_lsnhttphdrlogprofile_binding" "tf_lsngroup_lsnhttphdrlogprofile_binding" {
	groupname             = citrixadc_lsngroup.tf_lsngroup.groupname
	httphdrlogprofilename = citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile.httphdrlogprofilename
}
  
`

const testAccLsngroup_lsnhttphdrlogprofile_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccLsngroup_lsnhttphdrlogprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_lsnhttphdrlogprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnhttphdrlogprofile_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnhttphdrlogprofile_bindingExist("citrixadc_lsngroup_lsnhttphdrlogprofile_binding.tf_lsngroup_lsnhttphdrlogprofile_binding", nil),
				),
			},
			{
				Config: testAccLsngroup_lsnhttphdrlogprofile_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnhttphdrlogprofile_bindingNotExist("citrixadc_lsngroup_lsnhttphdrlogprofile_binding.tf_lsngroup_lsnhttphdrlogprofile_binding", "my_lsn_group,my_httplogprofile"),
				),
			},
		},
	})
}

func testAccCheckLsngroup_lsnhttphdrlogprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsngroup_lsnhttphdrlogprofile_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"groupname", "httphdrlogprofilename"}, nil)
		if err != nil {
			return err
		}
		groupname := idMap["groupname"]
		httphdrlogprofilename := idMap["httphdrlogprofilename"]

		findParams := service.FindParams{
			ResourceType:             "lsngroup_lsnhttphdrlogprofile_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching httphdrlogprofilename
		found := false
		for _, v := range dataArr {
			if v["httphdrlogprofilename"].(string) == httphdrlogprofilename {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsngroup_lsnhttphdrlogprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_lsnhttphdrlogprofile_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"groupname", "httphdrlogprofilename"}, nil)
		if err != nil {
			return err
		}
		groupname := idMap["groupname"]
		httphdrlogprofilename := idMap["httphdrlogprofilename"]

		findParams := service.FindParams{
			ResourceType:             "lsngroup_lsnhttphdrlogprofile_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching httphdrlogprofilename
		found := false
		for _, v := range dataArr {
			if v["httphdrlogprofilename"].(string) == httphdrlogprofilename {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsngroup_lsnhttphdrlogprofile_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_lsnhttphdrlogprofile_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsngroup_lsnhttphdrlogprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsngroup_lsnhttphdrlogprofile_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsngroup_lsnhttphdrlogprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsngroup_lsnhttphdrlogprofile_bindingDataSource_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsn_client"
}

resource "citrixadc_lsngroup" "tf_lsngroup" {
	groupname  = "my_lsn_group"
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
}

resource "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile" {
	httphdrlogprofilename = "my_httplogprofile"
}

resource "citrixadc_lsngroup_lsnhttphdrlogprofile_binding" "tf_lsngroup_lsnhttphdrlogprofile_binding" {
	groupname             = citrixadc_lsngroup.tf_lsngroup.groupname
	httphdrlogprofilename = citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile.httphdrlogprofilename
}

data "citrixadc_lsngroup_lsnhttphdrlogprofile_binding" "tf_lsngroup_lsnhttphdrlogprofile_binding" {
	groupname             = citrixadc_lsngroup_lsnhttphdrlogprofile_binding.tf_lsngroup_lsnhttphdrlogprofile_binding.groupname
	httphdrlogprofilename = citrixadc_lsngroup_lsnhttphdrlogprofile_binding.tf_lsngroup_lsnhttphdrlogprofile_binding.httphdrlogprofilename
	depends_on            = [citrixadc_lsngroup_lsnhttphdrlogprofile_binding.tf_lsngroup_lsnhttphdrlogprofile_binding]
}
`

func TestAccLsngroup_lsnhttphdrlogprofile_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_lsnhttphdrlogprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnhttphdrlogprofile_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_lsnhttphdrlogprofile_binding.tf_lsngroup_lsnhttphdrlogprofile_binding", "groupname", "my_lsn_group"),
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_lsnhttphdrlogprofile_binding.tf_lsngroup_lsnhttphdrlogprofile_binding", "httphdrlogprofilename", "my_httplogprofile"),
				),
			},
		},
	})
}

// testAccLsngroup_lsnhttphdrlogprofile_binding_upgrade_basic reuses the _basic config (binding +
// all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0 schema and the current
// Framework schema because the migration restored the SDK v2 attribute names. The terraform
// resource label (tf_lsngroup_lsnhttphdrlogprofile_binding) is kept identical to _basic so the
// Exist/Destroy helpers and resource addresses match.
const testAccLsngroup_lsnhttphdrlogprofile_binding_upgrade_basic = `
resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsn_client"
}

resource "citrixadc_lsngroup" "tf_lsngroup" {
	groupname  = "my_lsn_group"
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
}

resource "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile" {
	httphdrlogprofilename = "my_httplogprofile"
}

resource "citrixadc_lsngroup_lsnhttphdrlogprofile_binding" "tf_lsngroup_lsnhttphdrlogprofile_binding" {
	groupname             = citrixadc_lsngroup.tf_lsngroup.groupname
	httphdrlogprofilename = citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile.httphdrlogprofilename
}
`

// TestAccLsngroup_lsnhttphdrlogprofile_binding_sdkv2StateUpgrade verifies that state written by the
// last SDK v2 release (legacy comma-separated ID) is correctly upgraded when the same config is
// subsequently managed by the current Framework provider. Step 1 creates the binding with
// citrix/citrixadc 2.2.0 (writes the legacy id "my_lsn_group,my_httplogprofile"). Step 2
// refreshes/plans/applies the same config through the Framework provider, exercising ParseIdString
// on the legacy id; because the Framework recomputes the id on Read (SetAttrFromGet), the id
// upgrades to the new "key:value" form
// "groupname:my_lsn_group,httphdrlogprofilename:my_httplogprofile".
func TestAccLsngroup_lsnhttphdrlogprofile_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_lsngroup_lsnhttphdrlogprofile_binding.tf_lsngroup_lsnhttphdrlogprofile_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsngroup_lsnhttphdrlogprofile_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLsngroup_lsnhttphdrlogprofile_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnhttphdrlogprofile_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "my_lsn_group,my_httplogprofile"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework provider.
			// The legacy-id state is read via ParseIdString and the id is recomputed to the new
			// key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLsngroup_lsnhttphdrlogprofile_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnhttphdrlogprofile_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "groupname:my_lsn_group,httphdrlogprofilename:my_httplogprofile"),
				),
			},
		},
	})
}

func TestAccLsngroup_lsnhttphdrlogprofile_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lsngroup_lsnhttphdrlogprofile_binding.tf_lsngroup_lsnhttphdrlogprofile_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: groupname,httphdrlogprofilename) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"groupname", "httphdrlogprofilename"}
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
		CheckDestroy:             testAccCheckLsngroup_lsnhttphdrlogprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsngroup_lsnhttphdrlogprofile_binding_basic},
			{Config: testAccLsngroup_lsnhttphdrlogprofile_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccLsngroup_lsnhttphdrlogprofile_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccLsngroup_lsnhttphdrlogprofile_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lsngroup_lsnhttphdrlogprofile_binding.tf_lsngroup_lsnhttphdrlogprofile_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_lsnhttphdrlogprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnhttphdrlogprofile_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsngroup_lsnhttphdrlogprofile_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Lsngroup_lsnhttphdrlogprofile_binding.Type(), "my_lsn_group", map[string]string{"httphdrlogprofilename": "my_httplogprofile"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsngroup_lsnhttphdrlogprofile_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsngroup_lsnhttphdrlogprofile_bindingExist(resAddr, nil)),
			},
		},
	})
}
