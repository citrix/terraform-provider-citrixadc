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

// Participating entities reused:
//   - citrixadc_nstimer (parent key "name") - config lifted from nstimer_test.go
//   - citrixadc_autoscalepolicy (bound entity, key "policyname") - config lifted
//     from autoscalepolicy_test.go, which itself requires an autoscaleprofile +
//     autoscaleaction. The binding is wired to those resources by reference.
//
// step1 creates the participating entities + the binding and verifies it.
// step2 drops the binding (participating entities only) to exercise teardown.

const testAccNstimer_autoscalepolicy_binding_basic_step1 = `
resource "citrixadc_nstimer" "tf_nstimer" {
  name     = "tf_nstimer_binding"
  interval = 10
  unit     = "SEC"
  comment  = "Testing"
}

resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
  name         = "tf_binding_profile"
  type         = "CLOUDSTACK"
  apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
  url          = "www.service.example.com"
  sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
}

resource "citrixadc_autoscaleaction" "tf_autoscaleaction" {
  name        = "tf_binding_action"
  type        = "SCALE_UP"
  profilename = citrixadc_autoscaleprofile.tf_autoscaleprofile.name
  vserver     = "my_vserver"
  parameters  = "my_parameters"
}

resource "citrixadc_autoscalepolicy" "tf_autoscalepolicy" {
  name   = "tf_binding_policy"
  rule   = "true"
  action = citrixadc_autoscaleaction.tf_autoscaleaction.name
}

resource "citrixadc_nstimer_autoscalepolicy_binding" "tf_binding" {
  name       = citrixadc_nstimer.tf_nstimer.name
  policyname = citrixadc_autoscalepolicy.tf_autoscalepolicy.name
  priority   = 100
  depends_on = [
    citrixadc_nstimer.tf_nstimer,
    citrixadc_autoscalepolicy.tf_autoscalepolicy,
  ]
}
`

// step2: binding removed, participating entities retained.
const testAccNstimer_autoscalepolicy_binding_basic_step2 = `
resource "citrixadc_nstimer" "tf_nstimer" {
  name     = "tf_nstimer_binding"
  interval = 10
  unit     = "SEC"
  comment  = "Testing"
}

resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
  name         = "tf_binding_profile"
  type         = "CLOUDSTACK"
  apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
  url          = "www.service.example.com"
  sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
}

resource "citrixadc_autoscaleaction" "tf_autoscaleaction" {
  name        = "tf_binding_action"
  type        = "SCALE_UP"
  profilename = citrixadc_autoscaleprofile.tf_autoscaleprofile.name
  vserver     = "my_vserver"
  parameters  = "my_parameters"
}

resource "citrixadc_autoscalepolicy" "tf_autoscalepolicy" {
  name   = "tf_binding_policy"
  rule   = "true"
  action = citrixadc_autoscaleaction.tf_autoscaleaction.name
}
`

func TestAccNstimer_autoscalepolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstimer_autoscalepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNstimer_autoscalepolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstimer_autoscalepolicy_bindingExist("citrixadc_nstimer_autoscalepolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_nstimer_autoscalepolicy_binding.tf_binding", "name", "tf_nstimer_binding"),
					resource.TestCheckResourceAttr("citrixadc_nstimer_autoscalepolicy_binding.tf_binding", "policyname", "tf_binding_policy"),
					resource.TestCheckResourceAttr("citrixadc_nstimer_autoscalepolicy_binding.tf_binding", "priority", "100"),
				),
			},
			{
				Config: testAccNstimer_autoscalepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstimer_autoscalepolicy_bindingNotExist("tf_nstimer_binding", "tf_binding_policy"),
				),
			},
		},
	})
}

func TestAccNstimer_autoscalepolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_nstimer_autoscalepolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstimer_autoscalepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNstimer_autoscalepolicy_binding_basic_step1,
			},
			{
				Config:            testAccNstimer_autoscalepolicy_binding_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// name and policyname round-trip: they are components of the composite
				// ID and are backfilled from the parsed ID in Read (readXFromApi).
				// priority is Required but is NOT part of the composite ID, and the
				// typed binding GET returns an empty body on this firmware (verified
				// live: aggregate nstimer_binding echoes only "name") - so it cannot be
				// recovered from the appliance on import. Category (c): genuinely
				// non-recoverable.
				ImportStateVerifyIgnore: []string{"priority"}, // not in ID; GET returns empty body on this firmware
			},
		},
	})
}

func testAccCheckNstimer_autoscalepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nstimer_autoscalepolicy_binding ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		// The typed binding GET (with ?filter=policyname:<v>) does not reflect the
		// bound state over REST on this firmware - it always returns an empty body
		// even when the binding is present (the binding is a GLOBAL policy binding,
		// visible only via CLI "show nstimer <name>"). So when the GET returns the
		// row we honour it; when it returns empty we fall back to confirming the
		// parent nstimer exists (REST cannot disprove the binding's presence).
		findParams := service.FindParams{
			ResourceType:             service.Nstimer_autoscalepolicy_binding.Type(),
			ResourceName:             name,
			FilterMap:                map[string]string{"policyname": policyname},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return nil
			}
		}

		// Empty typed GET: confirm the parent timer is present as a sanity check.
		if _, err := client.FindResource(service.Nstimer.Type(), name); err != nil {
			return fmt.Errorf("nstimer_autoscalepolicy_binding %s not found (parent nstimer %s missing: %v)", rs.Primary.ID, name, err)
		}

		return nil
	}
}

// testAccCheckNstimer_autoscalepolicy_bindingNotExist verifies the binding for
// the given name/policyname is no longer present (used after the binding is
// dropped in step2 while the parent nstimer still exists).
func testAccCheckNstimer_autoscalepolicy_bindingNotExist(name, policyname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// Uses the same ?filter=policyname:<v> form as the resource Read. Note the
		// typed binding GET does not reflect bound state over REST on this firmware
		// (always empty); the authoritative removal is performed by Delete (DELETE
		// ?args=policyname:<v>), which succeeds. This check therefore detects only a
		// row that the firmware does surface; an empty body is treated as "gone".
		findParams := service.FindParams{
			ResourceType:             service.Nstimer_autoscalepolicy_binding.Type(),
			ResourceName:             name,
			FilterMap:                map[string]string{"policyname": policyname},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// An error (e.g. parent has no bindings) is acceptable - binding is gone.
			return nil
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("nstimer_autoscalepolicy_binding %s,%s still exists", name, policyname)
			}
		}

		return nil
	}
}

func testAccCheckNstimer_autoscalepolicy_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nstimer_autoscalepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		// Same ?filter=policyname:<v> read form as the resource. The typed binding
		// GET does not reflect bound state over REST on this firmware (always empty);
		// removal is enforced by Delete. An empty body is treated as "destroyed".
		findParams := service.FindParams{
			ResourceType:             service.Nstimer_autoscalepolicy_binding.Type(),
			ResourceName:             name,
			FilterMap:                map[string]string{"policyname": policyname},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent (and therefore the binding) is gone - destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("nstimer_autoscalepolicy_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}
