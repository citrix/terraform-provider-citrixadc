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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// NOTE on the wasmmodule resource:
//   - Creating a wasmmodule requires a module file (modulefile) AND its matching
//     cryptographic signature (signaturefile) to be present in /var/netscaler/wasm
//     on the appliance. `add wasmmodule` VALIDATES the signature against the module
//     (NITRO errorcode 1074 "Module and signature files not matched" on mismatch),
//     so a fabricated hash cannot be used — the fixture must be a genuine signed
//     pair. doWasmmodulePreChecks (helpers_test.go) stages testdata/tfwasmfile.wasm
//     + testdata/tfwasmfile.sig (a copy of the firmware-shipped sample
//     /var/netscaler/wasm/ns_sigsci.wasm + ns_sigsci_wasm.sig) into
//     /var/netscaler/wasm.
//   - modulefile and signaturefile are create-only (RequiresReplace); only
//     settingfile and comment are updatable/unsettable. settingfile is omitted here
//     (the shipped samples include no setting file), so the update/unset paths are
//     exercised via comment.

const testAccWasmmodule_basic = `
resource "citrixadc_wasmmodule" "tf_wasmmodule" {
	name          = "tf_wasmmodule"
	modulefile    = "tfwasmfile.wasm"
	signaturefile = "tfwasmfile.sig"
	comment       = "created by acc test"
}
`

const testAccWasmmodule_update = `
resource "citrixadc_wasmmodule" "tf_wasmmodule" {
	name          = "tf_wasmmodule"
	modulefile    = "tfwasmfile.wasm"
	signaturefile = "tfwasmfile.sig"
	comment       = "updated by acc test"
}
`

func TestAccWasmmodule_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmmodulePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWasmmoduleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWasmmodule_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWasmmoduleExist("citrixadc_wasmmodule.tf_wasmmodule", nil),
					resource.TestCheckResourceAttr("citrixadc_wasmmodule.tf_wasmmodule", "name", "tf_wasmmodule"),
					resource.TestCheckResourceAttr("citrixadc_wasmmodule.tf_wasmmodule", "modulefile", "tfwasmfile.wasm"),
					resource.TestCheckResourceAttr("citrixadc_wasmmodule.tf_wasmmodule", "signaturefile", "tfwasmfile.sig"),
					resource.TestCheckResourceAttr("citrixadc_wasmmodule.tf_wasmmodule", "comment", "created by acc test"),
				),
			},
			{
				Config: testAccWasmmodule_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWasmmoduleExist("citrixadc_wasmmodule.tf_wasmmodule", nil),
					resource.TestCheckResourceAttr("citrixadc_wasmmodule.tf_wasmmodule", "name", "tf_wasmmodule"),
					resource.TestCheckResourceAttr("citrixadc_wasmmodule.tf_wasmmodule", "comment", "updated by acc test"),
				),
			},
		},
	})
}

func TestAccWasmmodule_import(t *testing.T) {
	const resAddr = "citrixadc_wasmmodule.tf_wasmmodule"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmmodulePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWasmmoduleDestroy,
		Steps: []resource.TestStep{
			{Config: testAccWasmmodule_basic},
			{
				Config:            testAccWasmmodule_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckWasmmoduleExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No wasmmodule name is set")
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
		data, err := client.FindResource(service.Wasmmodule.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("wasmmodule %s not found", n)
		}

		return nil
	}
}

func testAccCheckWasmmoduleDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_wasmmodule" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Wasmmodule.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("wasmmodule %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccWasmmoduleDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmmodulePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccWasmmoduleDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_wasmmodule.test", "name", "tf_wasmmodule_ds"),
					resource.TestCheckResourceAttrSet("data.citrixadc_wasmmodule.test", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_wasmmodule.test", "modulefile", "tfwasmfile.wasm"),
					resource.TestCheckResourceAttr("data.citrixadc_wasmmodule.test", "comment", "ds fixture"),
				),
			},
		},
	})
}

const testAccWasmmoduleDataSource_basic = `
resource "citrixadc_wasmmodule" "tf_wasmmodule_ds" {
	name          = "tf_wasmmodule_ds"
	modulefile    = "tfwasmfile.wasm"
	signaturefile = "tfwasmfile.sig"
	comment       = "ds fixture"
}

data "citrixadc_wasmmodule" "test" {
	name = citrixadc_wasmmodule.tf_wasmmodule_ds.name
}
`

// Step 1: comment (the unset-eligible attribute) set to a valid non-default value.
const testAccWasmmodule_unset_step1 = `
resource "citrixadc_wasmmodule" "tf_unset" {
	name          = "tf_wasmmodule_unset"
	modulefile    = "tfwasmfile.wasm"
	signaturefile = "tfwasmfile.sig"
	comment       = "unset fixture"
}
`

// Step 2: comment removed from config -> provider must unset it on the appliance.
const testAccWasmmodule_unset_step2 = `
resource "citrixadc_wasmmodule" "tf_unset" {
	name          = "tf_wasmmodule_unset"
	modulefile    = "tfwasmfile.wasm"
	signaturefile = "tfwasmfile.sig"
}
`

func TestAccWasmmodule_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmmodulePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWasmmoduleDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value applies and persists.
				Config: testAccWasmmodule_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWasmmoduleExist("citrixadc_wasmmodule.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_wasmmodule.tf_unset", "comment", "unset fixture"),
				),
			},
			{
				// Removing it must unset -> appliance drops the value.
				Config: testAccWasmmodule_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWasmmoduleExist("citrixadc_wasmmodule.tf_unset", nil),
					testAccCheckWasmmoduleADCUnset("tf_wasmmodule_unset", "comment"),
				),
			},
		},
	})
}

// testAccCheckWasmmoduleADCUnset asserts an attribute is absent/empty directly on
// the appliance (not just in Terraform state), proving the unset took effect.
func testAccCheckWasmmoduleADCUnset(name, attr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Wasmmodule.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("wasmmodule %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != "" && got != "<nil>" {
			return fmt.Errorf("wasmmodule %s: appliance attr %q = %q, want empty (unset did not revert it)", name, attr, got)
		}
		return nil
	}
}

// TestAccWasmmodule_selfHealing verifies the provider re-creates the module when
// it is deleted out-of-band between apply steps (drift recovery).
func TestAccWasmmodule_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_wasmmodule.tf_wasmmodule"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmmodulePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWasmmoduleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWasmmodule_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckWasmmoduleExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Wasmmodule.Type(), "tf_wasmmodule"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccWasmmodule_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckWasmmoduleExist(resAddr, nil)),
			},
		},
	})
}
