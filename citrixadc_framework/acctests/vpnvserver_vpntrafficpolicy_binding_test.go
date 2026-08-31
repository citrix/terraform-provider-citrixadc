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

const testAccVpnvserver_vpntrafficpolicy_binding_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_examplevserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpntrafficaction" "foo" {
		fta  = "ON"
		hdx  = "ON"
		name = "Testingaction"
		qual = "tcp"
		sso  = "ON"
	}
	resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
		name   = "tf_vpntrafficpolicy"
		rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
		action = citrixadc_vpntrafficaction.foo.name
	}
	resource "citrixadc_vpnvserver_vpntrafficpolicy_binding" "tf_bind" {
		name                   = citrixadc_vpnvserver.tf_vpnvserver.name
		policy                 = citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy.name
		priority               = 200
		bindpoint              = "REQUEST"
		gotopriorityexpression = "NEXT"
	}
`

const testAccVpnvserver_vpntrafficpolicy_bindingDataSource_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_examplevserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpntrafficaction" "foo" {
		fta  = "ON"
		hdx  = "ON"
		name = "Testingaction"
		qual = "tcp"
		sso  = "ON"
	}
	resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
		name   = "tf_vpntrafficpolicy"
		rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
		action = citrixadc_vpntrafficaction.foo.name
	}
	resource "citrixadc_vpnvserver_vpntrafficpolicy_binding" "tf_bind" {
		name                   = citrixadc_vpnvserver.tf_vpnvserver.name
		policy                 = citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy.name
		priority               = 200
		bindpoint              = "REQUEST"
		gotopriorityexpression = "NEXT"
	}

	data "citrixadc_vpnvserver_vpntrafficpolicy_binding" "tf_bind" {
		name   = citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind.name
		policy = citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind.policy
	}
`

const testAccVpnvserver_vpntrafficpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_examplevserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpntrafficaction" "foo" {
		fta  = "ON"
		hdx  = "ON"
		name = "Testingaction"
		qual = "tcp"
		sso  = "ON"
	}
	resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
		name   = "tf_vpntrafficpolicy"
		rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
		action = citrixadc_vpntrafficaction.foo.name
	}
`

func TestAccVpnvserver_vpntrafficpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_vpntrafficpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_vpntrafficpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpntrafficpolicy_bindingExist("citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccVpnvserver_vpntrafficpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpntrafficpolicy_bindingNotExist("citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind", "tf_examplevserver,tf_vpntrafficpolicy"),
				),
			},
		},
	})
}

func testAccCheckVpnvserver_vpntrafficpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_vpntrafficpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "policy"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_vpntrafficpolicy_binding",
			ResourceName:             name,
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
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnvserver_vpntrafficpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_vpntrafficpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "policy"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_vpntrafficpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnvserver_vpntrafficpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_vpntrafficpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_vpntrafficpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnvserver_vpntrafficpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver_vpntrafficpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnvserver_vpntrafficpolicy_binding_upgrade_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_examplevserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpntrafficaction" "foo" {
		fta  = "ON"
		hdx  = "ON"
		name = "Testingaction"
		qual = "tcp"
		sso  = "ON"
	}
	resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
		name   = "tf_vpntrafficpolicy"
		rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
		action = citrixadc_vpntrafficaction.foo.name
	}
	resource "citrixadc_vpnvserver_vpntrafficpolicy_binding" "tf_bind" {
		name                   = citrixadc_vpnvserver.tf_vpnvserver.name
		policy                 = citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy.name
		priority               = 200
		bindpoint              = "REQUEST"
		gotopriorityexpression = "NEXT"
	}
`

func TestAccVpnvserver_vpntrafficpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnvserver_vpntrafficpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the resource with the last SDK v2 release (2.2.0),
				// which writes state with the legacy comma-joined id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.0.0",
					},
				},
				Config: testAccVpnvserver_vpntrafficpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpntrafficpolicy_bindingExist("citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind", "id", "tf_examplevserver,tf_vpntrafficpolicy"),
				),
			},
			{
				// Step 2: refresh/plan/apply the legacy-id state through the current
				// framework provider. Read recomputes the id to the new key:value format.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVpnvserver_vpntrafficpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpntrafficpolicy_bindingExist("citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind", "id", "name:tf_examplevserver,policy:tf_vpntrafficpolicy"),
				),
			},
		},
	})
}

func TestAccVpnvserver_vpntrafficpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_vpntrafficpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind", "name", "tf_examplevserver"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind", "policy", "tf_vpntrafficpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind", "priority", "200"),
				),
			},
		},
	})
}

func TestAccVpnvserver_vpntrafficpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,policy) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "policy"}
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
		CheckDestroy:             testAccCheckVpnvserver_vpntrafficpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnvserver_vpntrafficpolicy_binding_basic},
			{Config: testAccVpnvserver_vpntrafficpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"bindpoint", "gotopriorityexpression", "priority"}},
			{Config: testAccVpnvserver_vpntrafficpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"bindpoint", "gotopriorityexpression", "priority"}},
		},
	})
}

func TestAccVpnvserver_vpntrafficpolicy_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver_vpntrafficpolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_vpntrafficpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_vpntrafficpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpntrafficpolicy_bindingExist(resAddr, nil),
				),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Vpnvserver_vpntrafficpolicy_binding.Type(), "tf_examplevserver", []string{"policy:tf_vpntrafficpolicy", "bindpoint:REQUEST"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnvserver_vpntrafficpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpntrafficpolicy_bindingExist(resAddr, nil),
				),
			},
		},
	})
}
