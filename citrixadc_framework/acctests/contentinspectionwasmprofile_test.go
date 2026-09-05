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

// NOTE on the contentinspectionwasmprofile resource:
//   - `wasmmodule` is REQUIRED by NITRO on create (errorcode 1095 "Required
//     argument missing [wasmModule]" if omitted), and the referenced module must
//     be a Content-Inspection-compatible WASM module. The firmware-shipped
//     ns_sigsci module is CI-compatible; the ns_extrahop_* samples are NOT and are
//     rejected here with errorcode 278 "Invalid argument".
//   - Each config therefore first creates a citrixadc_wasmmodule from the CI sample
//     (testdata/tfwasmfile.wasm + tfwasmfile.sig, a copy of ns_sigsci) and
//     references it via wasmmodule. doWasmmodulePreChecks stages those files into
//     /var/netscaler/wasm (where `add wasmmodule` reads and signature-validates
//     them). Terraform's implicit dependency (profile -> module reference) also
//     guarantees the module is destroyed after the profile on teardown.

const testAccCIWasmProfile_module = `
resource "citrixadc_wasmmodule" "tf_ci_mod" {
	name          = "tf_ci_mod"
	modulefile    = "tfwasmfile.wasm"
	signaturefile = "tfwasmfile.sig"
}
`

const testAccContentinspectionwasmprofile_basic = testAccCIWasmProfile_module + `
resource "citrixadc_contentinspectionwasmprofile" "tf_ciwasmprofile" {
	name              = "tf_ciwasmprofile"
	wasmmodule        = citrixadc_wasmmodule.tf_ci_mod.name
	timeout           = 2000
	timeoutaction     = "BYPASS"
	maxbodylen        = 32
	anomalousdatasize = 256
	anomalousttfbtime = 2000
}
`

const testAccContentinspectionwasmprofile_update = testAccCIWasmProfile_module + `
resource "citrixadc_contentinspectionwasmprofile" "tf_ciwasmprofile" {
	name              = "tf_ciwasmprofile"
	wasmmodule        = citrixadc_wasmmodule.tf_ci_mod.name
	timeout           = 3000
	timeoutaction     = "RESET"
	maxbodylen        = 16
	anomalousdatasize = 512
	anomalousttfbtime = 1000
}
`

func TestAccContentinspectionwasmprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmmodulePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionwasmprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionwasmprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionwasmprofileExist("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "name", "tf_ciwasmprofile"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "wasmmodule", "tf_ci_mod"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "timeout", "2000"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "timeoutaction", "BYPASS"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "maxbodylen", "32"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "anomalousdatasize", "256"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "anomalousttfbtime", "2000"),
				),
			},
			{
				Config: testAccContentinspectionwasmprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionwasmprofileExist("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "name", "tf_ciwasmprofile"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "timeout", "3000"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "timeoutaction", "RESET"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "maxbodylen", "16"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "anomalousdatasize", "512"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile", "anomalousttfbtime", "1000"),
				),
			},
		},
	})
}

func TestAccContentinspectionwasmprofile_import(t *testing.T) {
	const resAddr = "citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmmodulePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionwasmprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccContentinspectionwasmprofile_basic},
			{
				Config:            testAccContentinspectionwasmprofile_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckContentinspectionwasmprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No contentinspectionwasmprofile name is set")
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
		data, err := client.FindResource("contentinspectionwasmprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("contentinspectionwasmprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckContentinspectionwasmprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_contentinspectionwasmprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("contentinspectionwasmprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("contentinspectionwasmprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccContentinspectionwasmprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmmodulePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionwasmprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionwasmprofile.test", "name", "tf_ciwasmprofile_ds"),
					resource.TestCheckResourceAttrSet("data.citrixadc_contentinspectionwasmprofile.test", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionwasmprofile.test", "timeoutaction", "RESET"),
					resource.TestCheckResourceAttrSet("data.citrixadc_contentinspectionwasmprofile.test", "timeout"),
				),
			},
		},
	})
}

const testAccContentinspectionwasmprofileDataSource_basic = testAccCIWasmProfile_module + `
resource "citrixadc_contentinspectionwasmprofile" "tf_ciwasmprofile_ds" {
	name          = "tf_ciwasmprofile_ds"
	wasmmodule    = citrixadc_wasmmodule.tf_ci_mod.name
	timeout       = 1500
	timeoutaction = "RESET"
}

data "citrixadc_contentinspectionwasmprofile" "test" {
	name = citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile_ds.name
}
`

// Step 1: every unset-eligible attribute set to a valid non-default value.
const testAccContentinspectionwasmprofile_unset_step1 = testAccCIWasmProfile_module + `
resource "citrixadc_contentinspectionwasmprofile" "tf_unset" {
	name              = "tf_ciwasmprofile_unset"
	wasmmodule        = citrixadc_wasmmodule.tf_ci_mod.name
	timeout           = 2000
	timeoutaction     = "BYPASS"
	maxbodylen        = 32
	anomalousdatasize = 256
	anomalousttfbtime = 2000
}
`

// Step 2: all unset-eligible attributes removed from config -> provider must unset
// them, reverting each to its NITRO default. wasmmodule stays (it is required).
const testAccContentinspectionwasmprofile_unset_step2 = testAccCIWasmProfile_module + `
resource "citrixadc_contentinspectionwasmprofile" "tf_unset" {
	name       = "tf_ciwasmprofile_unset"
	wasmmodule = citrixadc_wasmmodule.tf_ci_mod.name
}
`

func TestAccContentinspectionwasmprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmmodulePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionwasmprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccContentinspectionwasmprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionwasmprofileExist("citrixadc_contentinspectionwasmprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_unset", "timeout", "2000"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_unset", "timeoutaction", "BYPASS"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_unset", "maxbodylen", "32"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_unset", "anomalousdatasize", "256"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_unset", "anomalousttfbtime", "2000"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults.
				Config: testAccContentinspectionwasmprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionwasmprofileExist("citrixadc_contentinspectionwasmprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_unset", "timeout", "1000"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_unset", "timeoutaction", "DROP"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_unset", "maxbodylen", "16"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_unset", "anomalousdatasize", "512"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionwasmprofile.tf_unset", "anomalousttfbtime", "1000"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckContentinspectionwasmprofileADCValue("tf_ciwasmprofile_unset", "timeout", "1000"),
					testAccCheckContentinspectionwasmprofileADCValue("tf_ciwasmprofile_unset", "timeoutaction", "DROP"),
					testAccCheckContentinspectionwasmprofileADCValue("tf_ciwasmprofile_unset", "maxbodylen", "16"),
					testAccCheckContentinspectionwasmprofileADCValue("tf_ciwasmprofile_unset", "anomalousdatasize", "512"),
					testAccCheckContentinspectionwasmprofileADCValue("tf_ciwasmprofile_unset", "anomalousttfbtime", "1000"),
				),
			},
		},
	})
}

// testAccCheckContentinspectionwasmprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckContentinspectionwasmprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Contentinspectionwasmprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("contentinspectionwasmprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("contentinspectionwasmprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

// TestAccContentinspectionwasmprofile_selfHealing verifies the provider re-creates the
// profile when it is deleted out-of-band between apply steps (drift recovery).
func TestAccContentinspectionwasmprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_contentinspectionwasmprofile.tf_ciwasmprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doWasmmodulePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionwasmprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionwasmprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionwasmprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Contentinspectionwasmprofile.Type(), "tf_ciwasmprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccContentinspectionwasmprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionwasmprofileExist(resAddr, nil)),
			},
		},
	})
}
