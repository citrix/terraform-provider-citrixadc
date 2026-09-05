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

const testAccVpnvserver_aaapreauthenticationpolicy_binding_basic = `
	# Since the aaapreauthenticationpolicy resource is not yet available on Terraform,
	# the tf_aaapolicy policy must be created by hand in order for the script to run correctly.
	# You can do that by using the following Citrix ADC cli commands:
	# add aaa preauthenticationaction tf_aaaaction DENY
	# add aaa preauthenticationpolicy tf_aaapolicy NS_TRUE tf_aaaaction
	

	resource "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
		name                    = "tf_aaaaction"
		preauthenticationaction = "DENY"
		deletefiles             = "/var/tmp/new/hello.txt"
	}
	resource "citrixadc_aaapreauthenticationpolicy" "tf_aaapreauthenticationpolicy" {
		name 	  = "tf_aaapolicy"
		rule 	  = "NS_TRUE"
		reqaction = citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction.name
	}
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vpnvserverexample"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpnvserver_aaapreauthenticationpolicy_binding" "tf_binding" {
		name      = citrixadc_vpnvserver.tf_vpnvserver.name
		policy    = citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy.name
		priority  = 40
		secondary = "false"
		bindpoint = "OTHERTCP_REQUEST"
	}
`

const testAccVpnvserver_aaapreauthenticationpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	
	resource "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
		name                    = "tf_aaaaction"
		preauthenticationaction = "DENY"
		deletefiles             = "/var/tmp/new/hello.txt"
	}
	resource "citrixadc_aaapreauthenticationpolicy" "tf_aaapreauthenticationpolicy" {
		name 	  = "tf_aaapolicy"
		rule 	  = "NS_TRUE"
		reqaction = citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction.name
	}
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vpnvserverexample"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
`

func TestAccVpnvserver_aaapreauthenticationpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_aaapreauthenticationpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingExist("citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccVpnvserver_aaapreauthenticationpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingNotExist("citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding", "tf_vpnvserverexample,tf_aaapolicy"),
				),
			},
		},
	})
}

func testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_aaapreauthenticationpolicy_binding id is set")
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
			ResourceType:             "vpnvserver_aaapreauthenticationpolicy_binding",
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
			return fmt.Errorf("vpnvserver_aaapreauthenticationpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"name", "policy"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_aaapreauthenticationpolicy_binding",
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
			return fmt.Errorf("vpnvserver_aaapreauthenticationpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_aaapreauthenticationpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnvserver_aaapreauthenticationpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver_aaapreauthenticationpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnvserver_aaapreauthenticationpolicy_binding_upgrade_basic = `
	resource "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
		name                    = "tf_aaaaction"
		preauthenticationaction = "DENY"
		deletefiles             = "/var/tmp/new/hello.txt"
	}
	resource "citrixadc_aaapreauthenticationpolicy" "tf_aaapreauthenticationpolicy" {
		name 	  = "tf_aaapolicy"
		rule 	  = "NS_TRUE"
		reqaction = citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction.name
	}
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vpnvserverexample"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpnvserver_aaapreauthenticationpolicy_binding" "tf_binding" {
		name      = citrixadc_vpnvserver.tf_vpnvserver.name
		policy    = citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy.name
		priority  = 40
		secondary = "false"
		bindpoint = "OTHERTCP_REQUEST"
	}
`

// TestAccVpnvserver_aaapreauthenticationpolicy_binding_sdkv2StateUpgrade verifies that a
// resource created with the last SDK v2 release (2.2.0), which writes the legacy
// comma-joined id "name,policy", can be refreshed/planned/applied by the current
// Plugin Framework provider. On the framework Read the id is recomputed to the new
// key:value format via SetAttrFromGet -> BuildId.
func TestAccVpnvserver_aaapreauthenticationpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.0.0",
					},
				},
				Config: testAccVpnvserver_aaapreauthenticationpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingExist("citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding", "id", "tf_vpnvserverexample,tf_aaapolicy"),
				),
			},
			// Step 2: refresh the legacy-id state through the current framework provider.
			// The framework recomputes the id to the new key:value format on Read.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVpnvserver_aaapreauthenticationpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingExist("citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding", "id", "name:tf_vpnvserverexample,policy:tf_aaapolicy"),
				),
			},
		},
	})
}

const testAccVpnvserver_aaapreauthenticationpolicy_bindingDataSource_basic = `

	resource "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
		name                    = "tf_aaaaction"
		preauthenticationaction = "DENY"
		deletefiles             = "/var/tmp/new/hello.txt"
	}
	resource "citrixadc_aaapreauthenticationpolicy" "tf_aaapreauthenticationpolicy" {
		name 	  = "tf_aaapolicy"
		rule 	  = "NS_TRUE"
		reqaction = citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction.name
	}
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vpnvserverexample"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpnvserver_aaapreauthenticationpolicy_binding" "tf_binding" {
		name      = citrixadc_vpnvserver.tf_vpnvserver.name
		policy    = citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy.name
		priority  = 40
		secondary = "false"
		bindpoint = "OTHERTCP_REQUEST"
	}

	data "citrixadc_vpnvserver_aaapreauthenticationpolicy_binding" "tf_binding" {
		name   = citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding.name
		policy = citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding.policy
	}
`

func TestAccVpnvserver_aaapreauthenticationpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_aaapreauthenticationpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding", "name", "tf_vpnvserverexample"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding", "policy", "tf_aaapolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding", "priority", "40"),
					// Universal runtime-binding proof for the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding", "id"),
				),
			},
		},
	})
}

func TestAccVpnvserver_aaapreauthenticationpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding"

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
		CheckDestroy:             testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnvserver_aaapreauthenticationpolicy_binding_basic},
			{Config: testAccVpnvserver_aaapreauthenticationpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"bindpoint", "priority", "secondary"}},
			{Config: testAccVpnvserver_aaapreauthenticationpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"bindpoint", "priority", "secondary"}},
		},
	})
}

func TestAccVpnvserver_aaapreauthenticationpolicy_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_aaapreauthenticationpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingExist(resAddr, nil),
				),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Vpnvserver_aaapreauthenticationpolicy_binding.Type(), "tf_vpnvserverexample", []string{"policy:tf_aaapolicy", "secondary:false", "bindpoint:OTHERTCP_REQUEST"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnvserver_aaapreauthenticationpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_aaapreauthenticationpolicy_bindingExist(resAddr, nil),
				),
			},
		},
	})
}
