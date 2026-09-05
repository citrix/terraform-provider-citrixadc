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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// NOTE on the wasmfile resource:
//   - There is NO adc-nitro-go SDK struct for wasmfile. The resource drives the
//     generic NitroClient calls with the literal type string "wasmfile" and an
//     untyped map payload (flagged for review).
//   - Create maps to NITRO POST /wasmfile?action=Import; Delete maps to
//     DELETE /wasmfile?args=name:<name>; Read/get is get(all) filtered by name.
//   - NITRO's `change` (?action=update) accepts only `name`, so every writable
//     attribute is RequiresReplace and there is no true in-place update path.
//   - needs_fixture=true: Import needs a real WASM file reachable by `src`.
//     doWasmfilePreChecks (in helpers_test.go) uploads the testdata/tfwasmfile.wasm
//     fixture to /var/tmp so src="local:tfwasmfile.wasm" resolves without an
//     external host. The fixture is a genuine, validly-signed WASM module (a copy
//     of the firmware-shipped sample /var/netscaler/wasm/ns_sigsci.wasm);
//     wasmfile's Module import does not verify the signature, so the module alone
//     is sufficient here.

// doWasmfilePreChecks (which uploads testdata/tfwasmfile.wasm) lives in
// helpers_test.go alongside the other do*PreChecks; the configs below reference
// it as src="local:tfwasmfile.wasm".

const testAccWasmfile_basic_step1 = `
resource "citrixadc_wasmfile" "tf_wasmfile" {
  name     = "tf_wasmfile"
  src      = "local:tfwasmfile.wasm"
  filetype = "Module"
  comment  = "test_wasmfile_v1"
}
`

// Step 2 changes an attribute; since every attribute is RequiresReplace this
// exercises destroy+recreate rather than a true in-place update.
const testAccWasmfile_basic_step2 = `
resource "citrixadc_wasmfile" "tf_wasmfile" {
  name     = "tf_wasmfile"
  src      = "local:tfwasmfile.wasm"
  filetype = "Module"
  comment  = "test_wasmfile_v2"
}
`

func TestAccWasmfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWasmfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWasmfile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWasmfileExist("citrixadc_wasmfile.tf_wasmfile", nil),
					resource.TestCheckResourceAttr("citrixadc_wasmfile.tf_wasmfile", "name", "tf_wasmfile"),
					resource.TestCheckResourceAttr("citrixadc_wasmfile.tf_wasmfile", "comment", "test_wasmfile_v1"),
				),
			},
			{
				Config: testAccWasmfile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWasmfileExist("citrixadc_wasmfile.tf_wasmfile", nil),
					resource.TestCheckResourceAttr("citrixadc_wasmfile.tf_wasmfile", "comment", "test_wasmfile_v2"),
				),
			},
		},
	})
}

func TestAccWasmfile_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWasmfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWasmfile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckWasmfileExist("citrixadc_wasmfile.tf_wasmfile", nil)),
			},
			{
				ResourceName:      "citrixadc_wasmfile.tf_wasmfile",
				ImportState:       true,
				ImportStateVerify: true,
				// `src` and `overwrite` are Import-only inputs; NITRO get(all)
				// may not echo them exactly, so ignore them on import verify.
				ImportStateVerifyIgnore: []string{"src", "overwrite"},
			},
		},
	})
}

// TestAccWasmfile_selfHealing verifies drift recovery: after the resource is
// deleted out-of-band on the ADC, the next refresh's Read must detect it is gone
// and drop it from state so the same config recreates it.
func TestAccWasmfile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_wasmfile.tf_wasmfile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWasmfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWasmfile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckWasmfileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					args := []string{"name:tf_wasmfile"}
					if err := client.DeleteResourceWithArgs("wasmfile", "", args); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccWasmfile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckWasmfileExist(resAddr, nil)),
			},
		},
	})
}

// Datasource test. The datasource is queryable by `name` and NITRO get(all)
// echoes name/src/filetype/comment.
const testAccWasmfileDataSource_basic = `
resource "citrixadc_wasmfile" "tf_wasmfile" {
  name     = "tf_wasmfile"
  src      = "local:tfwasmfile.wasm"
  filetype = "Module"
  comment  = "test_wasmfile_ds"
}

data "citrixadc_wasmfile" "tf_wasmfile" {
  name       = citrixadc_wasmfile.tf_wasmfile.name
  depends_on = [citrixadc_wasmfile.tf_wasmfile]
}
`

func TestAccWasmfileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccWasmfileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_wasmfile.tf_wasmfile", "name", "tf_wasmfile"),
				),
			},
		},
	})
}

func testAccCheckWasmfileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No wasmfile name is set")
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

		findParams := service.FindParams{
			ResourceType:             "wasmfile",
			FilterMap:                map[string]string{"name": rs.Primary.ID},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		if len(dataArr) == 0 {
			return fmt.Errorf("wasmfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckWasmfileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_wasmfile" {
			continue
		}
		findParams := service.FindParams{
			ResourceType:             "wasmfile",
			FilterMap:                map[string]string{"name": rs.Primary.ID},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err == nil && len(dataArr) > 0 {
			return fmt.Errorf("wasmfile %s still exists after destroy", rs.Primary.ID)
		}
	}
	return nil
}
