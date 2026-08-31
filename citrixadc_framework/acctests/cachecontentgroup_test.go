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

const testAccCachecontentgroup_basic = `

	resource "citrixadc_cachecontentgroup" "tf_cachecontentgroup" {
		name                 = "my_cachecontentgroup"
		heurexpiryparam      = 30
		prefetch             = "YES"
		quickabortsize       = 40
		ignorereqcachinghdrs = "YES"
	}
`
const testAccCachecontentgroup_update = `

	resource "citrixadc_cachecontentgroup" "tf_cachecontentgroup" {
		name                 = "my_cachecontentgroup"
		heurexpiryparam      = 50
		prefetch             = "NO"
		quickabortsize       = 50
		ignorereqcachinghdrs = "NO"
	}
`

func TestAccCachecontentgroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCachecontentgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCachecontentgroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachecontentgroupExist("citrixadc_cachecontentgroup.tf_cachecontentgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "name", "my_cachecontentgroup"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "heurexpiryparam", "30"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "prefetch", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "quickabortsize", "40"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "ignorereqcachinghdrs", "YES"),
				),
			},
			{
				Config: testAccCachecontentgroup_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachecontentgroupExist("citrixadc_cachecontentgroup.tf_cachecontentgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "name", "my_cachecontentgroup"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "heurexpiryparam", "50"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "prefetch", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "quickabortsize", "50"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_cachecontentgroup", "ignorereqcachinghdrs", "NO"),
				),
			},
		},
	})
}

func TestAccCachecontentgroup_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCachecontentgroupDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccCachecontentgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCachecontentgroupExist("citrixadc_cachecontentgroup.tf_cachecontentgroup", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccCachecontentgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCachecontentgroupExist("citrixadc_cachecontentgroup.tf_cachecontentgroup", nil)),
			},
		},
	})
}

func TestAccCachecontentgroup_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_cachecontentgroup.tf_cachecontentgroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCachecontentgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCachecontentgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCachecontentgroupExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Cachecontentgroup.Type(), "my_cachecontentgroup"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCachecontentgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCachecontentgroupExist(resAddr, nil)),
			},
		},
	})
}

func TestAccCachecontentgroup_import(t *testing.T) {
	const resAddr = "citrixadc_cachecontentgroup.tf_cachecontentgroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCachecontentgroupDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCachecontentgroup_basic},
			{
				Config:                  testAccCachecontentgroup_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckCachecontentgroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cachecontentgroup name is set")
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
		data, err := client.FindResource(service.Cachecontentgroup.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cachecontentgroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckCachecontentgroupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cachecontentgroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cachecontentgroup.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cachecontentgroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// The unset test exercises the type-independent, mutable attributes wired into
// attributesToUnset. Step 1 sets each to a valid non-default value; step 2
// removes them so the provider must unset them (revert to the documented NITRO
// defaults). Attribute interdependencies are respected: prefetch is set to NO so
// alwaysevalpolicies can be YES, and flashcache is left at its default (NO) to
// avoid the flashcache/polleverytime (PET) mutual exclusion.
const testAccCachecontentgroup_unset_step1 = `
resource "citrixadc_cachecontentgroup" "tf_unset" {
  name                 = "tf_test_cachecontentgroup_unset"
  alwaysevalpolicies   = "YES"
  expireatlastbyte     = "YES"
  ignorereloadreq      = "NO"
  ignorereqcachinghdrs = "NO"
  insertage            = "NO"
  insertetag           = "NO"
  insertvia            = "NO"
  lazydnsresolve       = "NO"
  maxressize           = 100
  memlimit             = 32768
  minhits              = 5
  minressize           = 10
  persistha            = "YES"
  pinned               = "YES"
  polleverytime        = "YES"
  prefetch             = "NO"
  quickabortsize       = 40
  removecookies        = "NO"
}
`

const testAccCachecontentgroup_unset_step2 = `
resource "citrixadc_cachecontentgroup" "tf_unset" {
  name = "tf_test_cachecontentgroup_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccCachecontentgroup_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCachecontentgroupDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccCachecontentgroup_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachecontentgroupExist("citrixadc_cachecontentgroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "alwaysevalpolicies", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "expireatlastbyte", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "ignorereloadreq", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "ignorereqcachinghdrs", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "insertage", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "insertetag", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "insertvia", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "lazydnsresolve", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "maxressize", "100"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "memlimit", "32768"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "minhits", "5"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "minressize", "10"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "persistha", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "pinned", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "polleverytime", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "prefetch", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "quickabortsize", "40"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "removecookies", "NO"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccCachecontentgroup_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachecontentgroupExist("citrixadc_cachecontentgroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "alwaysevalpolicies", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "expireatlastbyte", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "ignorereloadreq", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "ignorereqcachinghdrs", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "insertage", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "insertetag", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "insertvia", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "lazydnsresolve", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "maxressize", "80"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "memlimit", "65536"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "minhits", "0"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "minressize", "0"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "persistha", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "pinned", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "polleverytime", "NO"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "prefetch", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "quickabortsize", "4194303"),
					resource.TestCheckResourceAttr("citrixadc_cachecontentgroup.tf_unset", "removecookies", "YES"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckCachecontentgroupADCValue("tf_test_cachecontentgroup_unset", "prefetch", "YES"),
					testAccCheckCachecontentgroupADCValue("tf_test_cachecontentgroup_unset", "insertvia", "YES"),
					testAccCheckCachecontentgroupADCValue("tf_test_cachecontentgroup_unset", "pinned", "NO"),
					testAccCheckCachecontentgroupADCValue("tf_test_cachecontentgroup_unset", "removecookies", "YES"),
				),
			},
		},
	})
}

// testAccCheckCachecontentgroupADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckCachecontentgroupADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Cachecontentgroup.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("cachecontentgroup %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("cachecontentgroup %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccCachecontentgroupDataSource_basic = `

	resource "citrixadc_cachecontentgroup" "tf_cachecontentgroup" {
		name                 = "my_cachecontentgroup_ds"
		heurexpiryparam      = 30
		prefetch             = "YES"
		quickabortsize       = 40
		ignorereqcachinghdrs = "YES"
	}

	data "citrixadc_cachecontentgroup" "tf_cachecontentgroup_ds" {
		name = citrixadc_cachecontentgroup.tf_cachecontentgroup.name
	}
`

func TestAccCachecontentgroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCachecontentgroupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cachecontentgroup.tf_cachecontentgroup_ds", "name", "my_cachecontentgroup_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_cachecontentgroup.tf_cachecontentgroup_ds", "heurexpiryparam", "30"),
					resource.TestCheckResourceAttr("data.citrixadc_cachecontentgroup.tf_cachecontentgroup_ds", "prefetch", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_cachecontentgroup.tf_cachecontentgroup_ds", "quickabortsize", "40"),
					resource.TestCheckResourceAttr("data.citrixadc_cachecontentgroup.tf_cachecontentgroup_ds", "ignorereqcachinghdrs", "YES"),
				),
			},
		},
	})
}
