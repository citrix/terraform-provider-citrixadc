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

const testAccCsvserver_vpnvserver_binding_basic = `

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
	name           = "tf_vpnvserver"
	servicetype    = "SSL"
}

resource "citrixadc_csvserver_vpnvserver_binding" "tf_csvserver_vpnvserver_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	vserver = citrixadc_vpnvserver.tf_vpnvserver.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "SSL"
	sslprofile = citrixadc_sslprofile.tf_sslprofile.name
}

resource "citrixadc_sslprofile" "tf_sslprofile" {
	name = "tf_sslprofile"
	ecccurvebindings = []
}
`

const testAccCsvserver_vpnvserver_binding_basic_step2 = `

	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name           = "tf_vpnvserver"
		servicetype    = "SSL"
	}
	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "SSL"
		sslprofile = citrixadc_sslprofile.tf_sslprofile.name
	}

	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name = "tf_sslprofile"
		ecccurvebindings = []
	}
`

func TestAccCsvserver_vpnvserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_vpnvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_vpnvserver_bindingExist("citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_vpnvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_vpnvserver_bindingNotExist("citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding", "tf_csvserver,tf_vpnvserver"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_vpnvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_vpnvserver_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "vserver"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		vserver := idMap["vserver"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_vpnvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["vserver"].(string) == vserver {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("csvserver_vpnvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_vpnvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "vserver"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		vserver := idMap["vserver"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_vpnvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right vpn vserver name
		found := false
		for _, v := range dataArr {
			if v["vserver"].(string) == vserver {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("csvserver_vpnvserver_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_vpnvserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_vpnvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Csvserver_vpnvserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_vpnvserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCsvserver_vpnvserver_bindingDataSource_basic = `

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
	name           = "tf_vpnvserver"
	servicetype    = "SSL"
}

resource "citrixadc_csvserver_vpnvserver_binding" "tf_csvserver_vpnvserver_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	vserver = citrixadc_vpnvserver.tf_vpnvserver.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "SSL"
	sslprofile = citrixadc_sslprofile.tf_sslprofile.name
}

resource "citrixadc_sslprofile" "tf_sslprofile" {
	name = "tf_sslprofile"
	ecccurvebindings = []
}

data "citrixadc_csvserver_vpnvserver_binding" "tf_csvserver_vpnvserver_binding" {
	name = citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding.name
	vserver = citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding.vserver
}
`

func TestAccCsvserver_vpnvserver_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_vpnvserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding", "name", "tf_csvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding", "vserver", "tf_vpnvserver"),
					resource.TestCheckResourceAttrSet("data.citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding", "id"),
				),
			},
		},
	})
}

// testAccCsvserver_vpnvserver_binding_upgrade_basic is the config used by the
// sdkv2 -> framework state-upgrade test. It reuses the same values and resource
// labels as testAccCsvserver_vpnvserver_binding_basic so it is valid under BOTH
// the SDK v2 2.2.0 schema and the current framework schema.
const testAccCsvserver_vpnvserver_binding_upgrade_basic = `

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
	name           = "tf_vpnvserver"
	servicetype    = "SSL"
}

resource "citrixadc_csvserver_vpnvserver_binding" "tf_csvserver_vpnvserver_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	vserver = citrixadc_vpnvserver.tf_vpnvserver.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "SSL"
}
`

// TestAccCsvserver_vpnvserver_binding_sdkv2StateUpgrade verifies that a binding
// created with the last SDK v2 release (2.2.0, legacy comma-separated ID) is
// correctly refreshed/planned/applied by the current framework provider.
func TestAccCsvserver_vpnvserver_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserver_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create the binding with the last SDK v2 release.
			// State is written with the LEGACY comma-separated id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.0.0",
					},
				},
				Config: testAccCsvserver_vpnvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_vpnvserver_bindingExist("citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding", "id", "tf_csvserver,tf_vpnvserver"),
				),
			},
			// Step 2: same config, current (framework) provider. Terraform
			// refreshes the legacy-id state through the framework Read
			// (exercising ParseIdString on the legacy id) then plans/applies.
			// The framework recomputes the id on read to the new key:value form.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccCsvserver_vpnvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_vpnvserver_bindingExist("citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding", "id", "name:tf_csvserver,vserver:tf_vpnvserver"),
				),
			},
		},
	})
}

func TestAccCsvserver_vpnvserver_binding_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,vserver) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "vserver"}
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
		CheckDestroy:             testAccCheckCsvserver_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_vpnvserver_binding_basic},
			{Config: testAccCsvserver_vpnvserver_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccCsvserver_vpnvserver_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccCsvserver_vpnvserver_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_csvserver_vpnvserver_binding.tf_csvserver_vpnvserver_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_vpnvserver_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserver_vpnvserver_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Csvserver_vpnvserver_binding.Type(), "tf_csvserver", map[string]string{"vserver": "tf_vpnvserver"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCsvserver_vpnvserver_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserver_vpnvserver_bindingExist(resAddr, nil)),
			},
		},
	})
}
