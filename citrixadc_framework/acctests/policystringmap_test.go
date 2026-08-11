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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccPolicystringmap_basic_step1 = `

resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap"
    comment = "Some comment"
}
`

const testAccPolicystringmap_basic_step2 = `

resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap"
    comment = "Some other comment"
}
`

func TestAccPolicystringmap_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicystringmapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicystringmap_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmapExist("citrixadc_policystringmap.tf_policystringmap", nil),
				),
			},
			{
				Config: testAccPolicystringmap_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmapExist("citrixadc_policystringmap.tf_policystringmap", nil),
				),
			},
		},
	})
}

func testAccCheckPolicystringmapExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policystringmap name is set")
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
		data, err := client.FindResource(service.Policystringmap.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("policystringmap %s not found", n)
		}

		return nil
	}
}

func testAccCheckPolicystringmapDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policystringmap" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Policystringmap.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("policystringmap %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccPolicystringmap_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_policystringmap.tf_policystringmap"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicystringmapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicystringmap_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicystringmapExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Policystringmap.Type(), "tf_policystringmap"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccPolicystringmap_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicystringmapExist(resAddr, nil)),
			},
		},
	})
}

func TestAccPolicystringmap_import(t *testing.T) {
	const resAddr = "citrixadc_policystringmap.tf_policystringmap"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicystringmapDestroy,
		Steps: []resource.TestStep{
			{Config: testAccPolicystringmap_basic_step1},
			{
				Config:                  testAccPolicystringmap_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccPolicystringmapDataSource_basic = `
	resource "citrixadc_policystringmap" "tf_policystringmap_ds" {
		name = "tf_policystringmap_ds"
		comment = "datasource test stringmap"
	}

	data "citrixadc_policystringmap" "tf_policystringmap_ds" {
		name = citrixadc_policystringmap.tf_policystringmap_ds.name
	}
`

func TestAccPolicystringmap_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckPolicystringmapDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccPolicystringmap_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicystringmapExist("citrixadc_policystringmap.tf_policystringmap", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccPolicystringmap_basic_step1,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckPolicystringmapExist("citrixadc_policystringmap.tf_policystringmap", nil)),
			},
		},
	})
}

const testAccPolicystringmap_unset_step1 = `
resource "citrixadc_policystringmap" "tf_unset" {
    name    = "tf_test_policystringmap_unset"
    comment = "non default comment"
}
`

const testAccPolicystringmap_unset_step2 = `
resource "citrixadc_policystringmap" "tf_unset" {
    name = "tf_test_policystringmap_unset"
    # comment removed from config -> the provider must unset it (revert to the
    # NITRO default, an empty string).
}
`

func TestAccPolicystringmap_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicystringmapDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccPolicystringmap_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmapExist("citrixadc_policystringmap.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_policystringmap.tf_unset", "comment", "non default comment"),
				),
			},
			{
				// Removing comment must unset it: state (read back from the
				// appliance) reverts to the NITRO default (empty string), and the
				// implicit post-apply plan must be empty.
				Config: testAccPolicystringmap_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmapExist("citrixadc_policystringmap.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_policystringmap.tf_unset", "comment", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckPolicystringmapADCValue("tf_test_policystringmap_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckPolicystringmapADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckPolicystringmapADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Policystringmap.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("policystringmap %s not found on appliance", name)
		}
		// NITRO omits an empty comment from GET, so a missing key means the
		// default (empty string) is in effect.
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("policystringmap %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccPolicystringmapDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicystringmapDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_policystringmap.tf_policystringmap_ds", "name", "tf_policystringmap_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_policystringmap.tf_policystringmap_ds", "comment", "datasource test stringmap"),
				),
			},
		},
	})
}
