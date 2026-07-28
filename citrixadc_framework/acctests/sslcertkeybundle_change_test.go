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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NOTE on the sslcertkeybundle_change resource:
//   - Models the NITRO sslcertkeybundle `change` action. The NITRO doc labels the
//     operation `change`, but the on-wire/CLI verb is `update`, so Create calls
//     ActOnResource(service.Sslcertkeybundle.Type(), &payload, "update").
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     change action, Read is a no-op (preserves state), Update is a no-op (every
//     attribute is RequiresReplace), and Delete is a state-only removal. There is
//     NO get/update/delete endpoint for the change op, so the resource CANNOT be
//     verified by reading it back from the ADC, and it has NO datasource.
//   - PREREQUISITE: `change` operates on an EXISTING cert-key bundle. The test
//     therefore first creates a citrixadc_sslcertkeybundle (the parent bundle)
//     and points the change resource at it via certkeybundlename. The parent's
//     PreCheck (doSslcertkeybundlePreChecks) uploads the required PEM bundle file
//     to /nsconfig/ssl so `bundlefile = "servercert1_certkeybundle.pem"` resolves.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("sslcertkeybundle_change-<certkeybundlename>");
//     it does NOT (and cannot) verify the change side-effect via NITRO.
//   - There is no CheckDestroy for the action resource itself (the change action
//     has no inverse on NITRO); CheckDestroy is reused from the parent test only
//     to confirm the real parent bundle is cleaned up (it filters on
//     rs.Type == "citrixadc_sslcertkeybundle", which excludes the change type).
//
// This mirrors the action-only test precedent in protocolhttpband_clear_test.go
// (single apply step, state-only Exist check, TestCheckResourceAttrSet on "id").
// The parent's tests are NOT skip-gated, so no ADC_TESTBED gate is applied here.

// Single apply step: every attribute of the change resource is RequiresReplace,
// so there is no in-place update to exercise. The parent bundle is created first
// and the change action re-applies the bundle file to it.
const testAccSslcertkeybundleChange_basic = `

resource "citrixadc_sslcertkeybundle" "tf_sslcertkeybundle" {
  certkeybundlename = "tf_sslcertkeybundle_change"
  bundlefile        = "servercert1_certkeybundle.pem"
}

resource "citrixadc_sslcertkeybundle_change" "tf_sslcertkeybundle_change" {
  certkeybundlename = citrixadc_sslcertkeybundle.tf_sslcertkeybundle.certkeybundlename
  bundlefile        = "servercert1_certkeybundle.pem"
  depends_on        = [citrixadc_sslcertkeybundle.tf_sslcertkeybundle]
}

`

func TestAccSslcertkeybundleChange_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeybundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy only verifies the real parent bundle is removed; the
		// change action resource has no inverse on NITRO and no GET-by-id.
		CheckDestroy: testAccCheckSslcertkeybundleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertkeybundleChange_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeybundleChangeExist("citrixadc_sslcertkeybundle_change.tf_sslcertkeybundle_change", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcertkeybundle_change.tf_sslcertkeybundle_change", "certkeybundlename", "tf_sslcertkeybundle_change"),
					resource.TestCheckResourceAttr("citrixadc_sslcertkeybundle_change.tf_sslcertkeybundle_change", "bundlefile", "servercert1_certkeybundle.pem"),
					// "id" is the synthetic state handle "sslcertkeybundle_change-tf_sslcertkeybundle_change".
					resource.TestCheckResourceAttrSet("citrixadc_sslcertkeybundle_change.tf_sslcertkeybundle_change", "id"),
				),
			},
		},
	})
}

// testAccCheckSslcertkeybundleChangeExist is a state-only existence check.
//
// sslcertkeybundle_change is an action-only resource: Read is a no-op and there is
// no GET-by-id endpoint for the change op, so we CANNOT verify the change via
// NITRO. We only assert that Terraform recorded the resource in state with a
// non-empty ID (which equals the synthetic
// "sslcertkeybundle_change-<certkeybundlename>" after a successful ?action=update).
// This mirrors testAccCheckProtocolhttpbandClearExist.
func testAccCheckSslcertkeybundleChangeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcertkeybundle_change ID is set")
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
