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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"net/url"
	"strings"
	"testing"
)

const testAccService_dospolicy_binding_basic = `
	# Since the dospolicy resource is not yet available on Terraform,
	# the tf_dospolicy policy must be created by hand in order for the script to run correctly.
	# You can do that by using the following Citrix ADC cli commands:
	# add dospolicy tf_dospolicy -qDepth 25

	resource "citrixadc_service" "tf_service" {
		servicetype         = "HTTP"
		name                = "tf_service"
		ipaddress           = "10.77.33.22"
		ip                  = "10.77.33.22"
		port                = "80"
		state               = "ENABLED"
		wait_until_disabled = true
	}
	resource "citrixadc_service_dospolicy_binding" "tf_binding" {
		name       = citrixadc_service.tf_service.name
		policyname = "tf_dospolicy"
	}
`

const testAccService_dospolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	# Since the dospolicy resource is not yet available on Terraform,
	# the tf_dospolicy policy must be created by hand in order for the script to run correctly.
	# You can do that by using the following Citrix ADC cli commands:
	# add dospolicy tf_dospolicy -qDepth 25

	resource "citrixadc_service" "tf_service" {
		servicetype         = "HTTP"
		name                = "tf_service"
		ipaddress           = "10.77.33.22"
		ip                  = "10.77.33.22"
		port                = "80"
		state               = "ENABLED"
		wait_until_disabled = true
	}
`

func TestAccService_dospolicy_binding_basic(t *testing.T) {
	t.Skipf("citrixadc_service_dospolicy_binding is not supported in 13.1")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckService_dospolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccService_dospolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckService_dospolicy_bindingExist("citrixadc_service_dospolicy_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccService_dospolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckService_dospolicy_bindingNotExist("citrixadc_service_dospolicy_binding.tf_binding", "tf_service,tf_dospolicy"),
				),
			},
		},
	})
}

func testAccCheckService_dospolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No service_dospolicy_binding id is set")
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

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "service_dospolicy_binding",
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("service_dospolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckService_dospolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "service_dospolicy_binding",
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("service_dospolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

// TestAccService_dospolicy_binding_sdkv2StateUpgrade verifies that a resource
// created by the LAST SDK v2 release (2.2.0) — which writes the legacy
// comma-joined id "name,policyname" — is refreshed and re-applied correctly by
// the CURRENT framework provider. Step 2 exercises ParseIdString on the legacy id
// during the framework Read.
//
// The resource-side SetAttrFromGet RECOMPUTES data.Id to the new canonical
// key:value format on Read (see resource_schema.go), so after the step-2 refresh
// the id becomes "name:tf_service,policyname:tf_dospolicy" — asserted below.
func TestAccService_dospolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	t.Skipf("citrixadc_service_dospolicy_binding is not supported in 13.1")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckService_dospolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release from the registry. This
			// writes state carrying the LEGACY comma-joined id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccService_dospolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckService_dospolicy_bindingExist("citrixadc_service_dospolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_service_dospolicy_binding.tf_binding", "id", "tf_service,tf_dospolicy"),
				),
			},
			// Step 2: same config through the CURRENT framework provider. Terraform
			// refreshes the legacy-id state through the framework Read (exercising
			// ParseIdString on the legacy id) then plans/applies. The framework Read
			// recomputes the id to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccService_dospolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckService_dospolicy_bindingExist("citrixadc_service_dospolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_service_dospolicy_binding.tf_binding", "id", "name:tf_service,policyname:tf_dospolicy"),
				),
			},
		},
	})
}

func testAccCheckService_dospolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_service_dospolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Service_dospolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("service_dospolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccService_dospolicy_binding_import(t *testing.T) {
	t.Skipf("citrixadc_service_dospolicy_binding is not supported in 13.1")
	const resAddr = "citrixadc_service_dospolicy_binding.tf_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,policyname) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "policyname"}
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
		CheckDestroy:             testAccCheckService_dospolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccService_dospolicy_binding_basic},
			{Config: testAccService_dospolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccService_dospolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccService_dospolicy_binding_selfHealing(t *testing.T) {
	t.Skipf("citrixadc_service_dospolicy_binding is not supported in 13.1")
	const resAddr = "citrixadc_service_dospolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckService_dospolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccService_dospolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckService_dospolicy_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Service_dospolicy_binding.Type(), "tf_service", map[string]string{"policyname": "tf_dospolicy"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccService_dospolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckService_dospolicy_bindingExist(resAddr, nil)),
			},
		},
	})
}
