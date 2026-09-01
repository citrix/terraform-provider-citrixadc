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
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccCachepolicy_basic = `


resource "citrixadc_cachepolicy" "tf_cachepolicy" {
	policyname  = "my_cachepolicy"
	rule        = "true"
	action      = "CACHE"
	undefaction = "NOCACHE"
	}
  
`
const testAccCachepolicy_update = `

	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "my_cachepolicy"
		rule        = "true"
		action      = "MAY_CACHE"
		undefaction = "RESET"
	}
  
`

func TestAccCachepolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCachepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCachepolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicyExist("citrixadc_cachepolicy.tf_cachepolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "policyname", "my_cachepolicy"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "action", "CACHE"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "undefaction", "NOCACHE"),
				),
			},
			{
				Config: testAccCachepolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicyExist("citrixadc_cachepolicy.tf_cachepolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "policyname", "my_cachepolicy"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "action", "MAY_CACHE"),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_cachepolicy", "undefaction", "RESET"),
				),
			},
		},
	})
}

func TestAccCachepolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_cachepolicy.tf_cachepolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCachepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCachepolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCachepolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Cachepolicy.Type(), "my_cachepolicy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCachepolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCachepolicyExist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckCachepolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cachepolicy name is set")
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
		data, err := client.FindResource(service.Cachepolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cachepolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckCachepolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cachepolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cachepolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cachepolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCachepolicy_import(t *testing.T) {
	const resAddr = "citrixadc_cachepolicy.tf_cachepolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCachepolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCachepolicy_basic},
			{
				Config:                  testAccCachepolicy_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccCachepolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCachepolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccCachepolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCachepolicyExist("citrixadc_cachepolicy.tf_cachepolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccCachepolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCachepolicyExist("citrixadc_cachepolicy.tf_cachepolicy", nil)),
			},
		},
	})
}

// The cachepolicy unset test covers storeingroup, the one unset-eligible
// attribute that reverts to a documented, round-trippable NITRO default
// ("DEFAULT"). Step 1 sets it to a non-default value; step 2 removes it from
// config and the provider must unset it so the appliance reverts to "DEFAULT".
// (undefaction is excluded: its NITRO revert value "Use Global" is not a valid
// input value, so it cannot be cleanly wired for unset.)
const testAccCachepolicy_unset_step1 = `
resource "citrixadc_cachepolicy" "tf_unset" {
	policyname   = "tf_test_cachepolicy_unset"
	rule         = "true"
	action       = "CACHE"
	storeingroup = "BASEFILE"
}
`

const testAccCachepolicy_unset_step2 = `
resource "citrixadc_cachepolicy" "tf_unset" {
	policyname = "tf_test_cachepolicy_unset"
	rule       = "true"
	action     = "CACHE"
	# storeingroup removed from config -> the provider must unset it (revert to
	# the NITRO default "DEFAULT").
}
`

func TestAccCachepolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCachepolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccCachepolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicyExist("citrixadc_cachepolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_unset", "storeingroup", "BASEFILE"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the documented NITRO default, and the
				// implicit post-apply plan must be empty.
				Config: testAccCachepolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicyExist("citrixadc_cachepolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cachepolicy.tf_unset", "storeingroup", "DEFAULT"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckCachepolicyADCValue("tf_test_cachepolicy_unset", "storeingroup", "DEFAULT"),
				),
			},
		},
	})
}

// testAccCheckCachepolicyADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckCachepolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Cachepolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("cachepolicy %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("cachepolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccCachepolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCachepolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_cachepolicy.tf_cachepolicy_ds", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_cachepolicy.tf_cachepolicy_ds", "policyname", "tf_cachepolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_cachepolicy.tf_cachepolicy_ds", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_cachepolicy.tf_cachepolicy_ds", "action", "CACHE"),
					resource.TestCheckResourceAttr("data.citrixadc_cachepolicy.tf_cachepolicy_ds", "undefaction", "NOCACHE"),
				),
			},
		},
	})
}

const testAccCachepolicyDataSource_basic = `

resource "citrixadc_cachepolicy" "tf_cachepolicy_ds" {
    policyname  = "tf_cachepolicy_ds"
    rule        = "true"
    action      = "CACHE"
    undefaction = "NOCACHE"
}

data "citrixadc_cachepolicy" "tf_cachepolicy_ds" {
    policyname = citrixadc_cachepolicy.tf_cachepolicy_ds.policyname
    depends_on = [citrixadc_cachepolicy.tf_cachepolicy_ds]
}

`
