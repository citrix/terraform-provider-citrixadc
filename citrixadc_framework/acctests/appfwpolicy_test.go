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

const testAccAppfwpolicy_add = `
	resource citrixadc_appfwprofile tfAcc_appfwprofile {
		name = "tfAcc_appfwprofile"
		type = ["HTML"]
	}

	resource citrixadc_appfwpolicy tfAcc_appfwpolicy1 {
		name = "tfAcc_appfwpolicy1"
		profilename = citrixadc_appfwprofile.tfAcc_appfwprofile.name
		rule = "true"
	}
`
const testAccAppfwpolicy_update = `
	resource citrixadc_appfwprofile tfAcc_appfwprofile {
		name = "tfAcc_appfwprofile"
		type = ["HTML"]
	}

	resource citrixadc_appfwpolicy tfAcc_appfwpolicy1 {
		name = "tfAcc_appfwpolicy1"
		profilename = citrixadc_appfwprofile.tfAcc_appfwprofile.name
		rule = "true"
        comment = "test comment"
	}
`

const testAccAppfwpolicyDataSource_basic = `
	resource citrixadc_appfwprofile tfAcc_appfwprofile {
		name = "tfAcc_appfwprofile"
		type = ["HTML"]
	}

	resource citrixadc_appfwpolicy tfAcc_appfwpolicy1 {
		name = "tfAcc_appfwpolicy1"
		profilename = citrixadc_appfwprofile.tfAcc_appfwprofile.name
		rule = "true"
	}

	data "citrixadc_appfwpolicy" "tfAcc_appfwpolicy1" {
		name = citrixadc_appfwpolicy.tfAcc_appfwpolicy1.name
		depends_on = [citrixadc_appfwpolicy.tfAcc_appfwpolicy1]
	}
`

func TestAccAppfwpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwpolicyExist("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "name", "tfAcc_appfwpolicy1"),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "profilename", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "rule", "true"),
				),
			},
			{
				Config: testAccAppfwpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwpolicyExist("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "name", "tfAcc_appfwpolicy1"),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "profilename", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "comment", "test comment"),
				),
			},
		},
	})
}

func testAccCheckAppfwpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
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
		data, err := client.FindResource(service.Appfwpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAppfwpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_appfwpolicy.tfAcc_appfwpolicy1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Appfwpolicy.Type(), "tfAcc_appfwpolicy1"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAppfwpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAppfwpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_appfwpolicy.tfAcc_appfwpolicy1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwpolicy_add},
			{
				Config:                  testAccAppfwpolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAppfwpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAppfwpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwpolicyExist("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAppfwpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwpolicyExist("citrixadc_appfwpolicy.tfAcc_appfwpolicy1", nil)),
			},
		},
	})
}

// Unset test: step1 sets the unset-eligible attributes (comment, logaction) to
// valid non-default values; step2 removes them from config so the provider must
// issue the NITRO ?action=unset, reverting them to their defaults (absent -> null).
const testAccAppfwpolicy_unset_step1 = `
	resource citrixadc_appfwprofile tfAcc_appfwprofile_unset {
		name = "tfAcc_appfwprofile_unset"
		type = ["HTML"]
	}

	resource citrixadc_auditmessageaction tfAcc_msgact_unset {
		name              = "tfAcc_msgact_unset"
		loglevel          = "INFORMATIONAL"
		stringbuilderexpr = "\"unset test\""
	}

	resource citrixadc_appfwpolicy tfAcc_appfwpolicy_unset {
		name        = "tfAcc_appfwpolicy_unset"
		profilename = citrixadc_appfwprofile.tfAcc_appfwprofile_unset.name
		rule        = "true"
		comment     = "unset acc comment"
		logaction   = citrixadc_auditmessageaction.tfAcc_msgact_unset.name
	}
`

const testAccAppfwpolicy_unset_step2 = `
	resource citrixadc_appfwprofile tfAcc_appfwprofile_unset {
		name = "tfAcc_appfwprofile_unset"
		type = ["HTML"]
	}

	resource citrixadc_auditmessageaction tfAcc_msgact_unset {
		name              = "tfAcc_msgact_unset"
		loglevel          = "INFORMATIONAL"
		stringbuilderexpr = "\"unset test\""
	}

	resource citrixadc_appfwpolicy tfAcc_appfwpolicy_unset {
		name        = "tfAcc_appfwpolicy_unset"
		profilename = citrixadc_appfwprofile.tfAcc_appfwprofile_unset.name
		rule        = "true"
		# comment and logaction removed from config -> provider must unset them.
	}
`

func TestAccAppfwpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAppfwpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwpolicyExist("citrixadc_appfwpolicy.tfAcc_appfwpolicy_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy_unset", "comment", "unset acc comment"),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicy.tfAcc_appfwpolicy_unset", "logaction", "tfAcc_msgact_unset"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the NITRO default (absent -> empty), and the
				// implicit post-apply plan must be empty.
				Config: testAccAppfwpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwpolicyExist("citrixadc_appfwpolicy.tfAcc_appfwpolicy_unset", nil),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAppfwpolicyADCValue("tfAcc_appfwpolicy_unset", "comment", ""),
					testAccCheckAppfwpolicyADCValue("tfAcc_appfwpolicy_unset", "logaction", ""),
				),
			},
		},
	})
}

// testAccCheckAppfwpolicyADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. After an unset the attribute is omitted from the GET response, so the
// expected value for a reverted attribute is the empty string.
func testAccCheckAppfwpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Appfwpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("appfwpolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("appfwpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAppfwpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "name", "tfAcc_appfwpolicy1"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "profilename", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "rule", "true"),
					// id is the universal runtime-binding proof for the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "id"),
					// Read-only metadata exposed only by the data source. hits/undefhits
					// are counter-style and policytype is a state field, all always
					// populated for a freshly-created policy.
					resource.TestCheckResourceAttrSet("data.citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "hits"),
					resource.TestCheckResourceAttrSet("data.citrixadc_appfwpolicy.tfAcc_appfwpolicy1", "undefhits"),
				),
			},
		},
	})
}
