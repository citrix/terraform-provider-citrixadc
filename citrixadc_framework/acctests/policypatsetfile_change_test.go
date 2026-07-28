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

	"github.com/citrix/adc-nitro-go/resource/config/policy"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NOTE on the policypatsetfile_change resource:
//   - Models the NITRO POST /policypatsetfile?action=update endpoint (CLI verb is
//     `update policy patsetfile`; there is NO `?action=change` endpoint). It
//     re-loads an updated pattern set from an already-imported patset file.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     update action via ActOnResource(service.Policypatsetfile.Type(), &payload,
//     "update"), Read is a no-op (preserves state), Update is a no-op (the single
//     attribute `name` is RequiresReplace), and Delete is a state-only removal.
//     There is NO get/update/delete-back endpoint for the action, so the resource
//     CANNOT be verified by reading it back from the ADC, and it has NO datasource.
//   - PREREQUISITE (verified on the testbed): `?action=update` refreshes an
//     already-LOADED patset file. Loading a patset file is a two-step NITRO
//     sequence that has NO dedicated Terraform resource in this provider:
//        1. Import the raw file   POST ?action=Import   -> imported list only
//        2. Load the patterns     POST /policypatsetfile -> `show policy patsetfile`
//     `update` resolves the name only after step (2); with just the imported file
//     (what the `citrixadc_policypatsetfile` import resource creates) it returns
//     errorcode 258 "No such resource [name, <name>]". Because there is no
//     Terraform resource for the load step, this test stages both steps in its
//     PreCheck (doPolicypatsetfileChangePreChecks) via the NITRO client and cleans
//     them up afterwards with a single DELETE (which removes both the loaded entry
//     and the imported file). The parent test's PreCheck upload helper is reused to
//     place testdata/tftest.patset on the appliance so the import resolves.
//   - The single attribute `name` is Required and RequiresReplace.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("policypatsetfile_change-<name>"); it does NOT
//     (and cannot) verify the update side-effect via NITRO.
//   - There is no CheckDestroy: the update action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal. The
//     staged patset file is cleaned up by the PreCheck's t.Cleanup.
//
// The parent policypatsetfile test is NOT skip-gated (runs on standalone), so no
// ADC_TESTBED gate is applied here either.

const testAccPolicypatsetfileChange_name = "tf_policypatsetfile_change"

// Single apply step: `name` is RequiresReplace, so there is no in-place update to
// exercise. The patset file is imported+loaded in the PreCheck; the change action
// then re-loads (updates) it.
const testAccPolicypatsetfileChange_basic = `
resource "citrixadc_policypatsetfile_change" "tf_policypatsetfile_change" {
  name = "tf_policypatsetfile_change"
}

`

// doPolicypatsetfileChangePreChecks stages the prerequisite LOADED patset file the
// `update` action operates on (see the resource note above), and registers cleanup.
func doPolicypatsetfileChangePreChecks(t *testing.T) {
	// Uploads testdata/tftest.patset to /var/tmp and runs testAccPreCheck.
	doPolicyPatSetFilePreChecks(t)

	client, err := testAccGetFrameworkClient()
	if err != nil {
		t.Fatalf("Failed to get test client: %v", err)
	}

	name := testAccPolicypatsetfileChange_name

	// Step 1: import the raw file (POST ?action=Import).
	importPayload := policy.Policypatsetfile{
		Name:      name,
		Src:       "local:tftest.patset",
		Charset:   "ASCII",
		Overwrite: true,
	}
	if err := client.ActOnResource(service.Policypatsetfile.Type(), &importPayload, "Import"); err != nil {
		t.Fatalf("PreCheck: failed to import patset file %s: %v", name, err)
	}

	// Step 2: load the patterns (POST /policypatsetfile, no action) so the name is
	// resolvable by `update`.
	loadPayload := policy.Policypatsetfile{Name: name}
	if _, err := client.AddResource(service.Policypatsetfile.Type(), name, &loadPayload); err != nil {
		t.Fatalf("PreCheck: failed to load (add) patset file %s: %v", name, err)
	}

	// A single DELETE /policypatsetfile/<name> removes both the loaded entry and
	// the imported file, leaving the appliance clean.
	t.Cleanup(func() {
		_ = client.DeleteResource(service.Policypatsetfile.Type(), name)
	})
}

func TestAccPolicypatsetfileChange_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doPolicypatsetfileChangePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the update action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccPolicypatsetfileChange_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicypatsetfileChangeExist("citrixadc_policypatsetfile_change.tf_policypatsetfile_change", nil),
					resource.TestCheckResourceAttr("citrixadc_policypatsetfile_change.tf_policypatsetfile_change", "name", "tf_policypatsetfile_change"),
					// "id" is the synthetic state handle "policypatsetfile_change-tf_policypatsetfile_change".
					resource.TestCheckResourceAttrSet("citrixadc_policypatsetfile_change.tf_policypatsetfile_change", "id"),
				),
			},
		},
	})
}

// testAccCheckPolicypatsetfileChangeExist is a state-only existence check.
//
// policypatsetfile_change is an action-only resource: Read is a no-op and there is
// no GET-by-id endpoint for the update action, so we CANNOT verify the update via
// NITRO. We only assert that Terraform recorded the resource in state with a
// non-empty ID (which equals the synthetic "policypatsetfile_change-<name>" after a
// successful POST ?action=update). This mirrors testAccCheckProtocolhttpbandClearExist.
func testAccCheckPolicypatsetfileChangeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policypatsetfile_change ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against; presence of the synthetic state
		// ID is the only confirmation we can make for an action-only resource.
		return nil
	}
}
