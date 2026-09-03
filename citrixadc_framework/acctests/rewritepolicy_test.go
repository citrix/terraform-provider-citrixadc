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

func TestAccRewritepolicy_globalbinding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewritepolicy_globalbinding_not_exists,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					verifyRewriteGlobalBindingExists("REQ_OVERRIDE", "tf_rewrite_policy", true),
				),
			},
			{
				Config: testAccRewritepolicy_globalbinding_exists,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					verifyRewriteGlobalBindingExists("REQ_DEFAULT", "tf_rewrite_policy", false),
				),
			},
			/*
							   // TODO: Find race condition that makes this fail. In manual testing this succeeds
								resource.TestStep{
									Config: testAccRewritepolicy_globalbinding_modified,
									Check: resource.ComposeTestCheckFunc(
										testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
										verifyRewriteGlobalBindingExists("REQ_OVERRIDE", "tf_rewrite_policy", false),
									),
				},
			*/
		},
	})
}

func TestAccRewritepolicy_lbvserverbinding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewritepolicy_lbvserverbindings_none,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", true),
					verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", true),
				),
			},
			{
				Config: testAccRewritepolicy_lbvserverbindings_both,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", false),
					verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", false),
				),
			},
			{
				Config: testAccRewritepolicy_lbvserverbindings_one,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", false),
					verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", true),
				),
			},
		},
	})
}

func TestAccRewritepolicy_csvserverbinding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewritepolicy_csvserverbindings_none,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", true),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", true),
				),
			},
			{
				Config: testAccRewritepolicy_csvserverbindings_both,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", true),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", true),
				),
			},
			{
				Config: testAccRewritepolicy_csvserverbindings_one,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", true),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", true),
				),
			},
		},
	})
}

func testAccCheckRewritepolicyExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := client.FindResource(service.Rewritepolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckRewritepolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rewritepolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Rewritepolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func verifyRewriteGlobalBindingExists(bindtype string, policyname string, inverse bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		bindFound := false
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		globalBindings, _ := client.FindResourceArray("rewritepolicy_rewriteglobal_binding", policyname)
		for _, val := range globalBindings {
			boundtoSlice := strings.Split(val["boundto"].(string), " ")
			if bindtype == boundtoSlice[1] {
				bindFound = true
				break
			}
		}

		if !inverse {
			if bindFound {
				return nil
			} else {
				return fmt.Errorf("Verify error cannot find bind of type %v for policyname %v\n", bindtype, policyname)
			}
		} else {
			if bindFound {
				return fmt.Errorf("Verify error found exessive bind of type %v for policyname %v\n", bindtype, policyname)
			} else {
				return nil
			}
		}
	}
}

func verifyRewriteLbvserverBindingExists(servername string, policyname string, inverse bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		bindFound := false
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		lbVserverBindings, _ := client.FindResourceArray("rewritepolicy_lbvserver_binding", policyname)
		for _, val := range lbVserverBindings {
			boundtoSlice := strings.Split(val["boundto"].(string), " ")
			if servername == boundtoSlice[2] {
				bindFound = true
				break
			}
		}

		if !inverse {
			if bindFound {
				return nil
			} else {
				return fmt.Errorf("Verify error cannot find bind to lbvserver %v for policyname %v\n", servername, policyname)
			}
		} else {
			if bindFound {
				return fmt.Errorf("Verify error found exessive bind to lbvserver %v for policyname %v\n", servername, policyname)
			} else {
				return nil
			}
		}
	}
}

const testAccRewritepolicy_globalbinding_exists = `

resource "citrixadc_lbvserver" "tf_lbvserver" {

  ipv46 = "10.66.22.33"
  name = "tf_lbvserver_name"
  port = 80
  servicetype = "HTTP"

}


resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

	depends_on = [ citrixadc_lbvserver.tf_lbvserver ]

	globalbinding {
		gotopriorityexpression = "END"
		labelname = citrixadc_lbvserver.tf_lbvserver.name
		labeltype = "reqvserver"
		priority = 205
		invoke = true
		type = "REQ_DEFAULT"
	}
}
`

const testAccRewritepolicy_globalbinding_modified = `

resource "citrixadc_lbvserver" "tf_lbvserver" {

  ipv46 = "10.66.22.33"
  name = "tf_lbvserver"
  port = 80
  servicetype = "HTTP"

}


resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
	depends_on = [ citrixadc_lbvserver.tf_lbvserver ]

	globalbinding {
            gotopriorityexpression = "END"
            labelname = citrixadc_lbvserver.tf_lbvserver.name
            labeltype = "reqvserver"
            priority = 208
            invoke = true
            type = "REQ_OVERRIDE"
	}
}
`

const testAccRewritepolicy_globalbinding_not_exists = `

resource "citrixadc_lbvserver" "tf_lbvserver" {

  ipv46 = "10.66.22.33"
  name = "tf_lbvserver"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
	depends_on = [ citrixadc_lbvserver.tf_lbvserver ]
}
`

const testAccRewritepolicy_lbvserverbindings_none = `
resource "citrixadc_lbvserver" "tf_lbvserver1" {

  ipv46 = "10.22.22.22"
  name = "tf_lbvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_lbvserver" "tf_lbvserver2" {

  ipv46 = "10.33.22.66"
  name = "tf_lbvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}
`

const testAccRewritepolicy_lbvserverbindings_one = `
resource "citrixadc_lbvserver" "tf_lbvserver1" {

  ipv46 = "10.22.22.22"
  name = "tf_lbvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_lbvserver" "tf_lbvserver2" {

  ipv46 = "10.33.22.66"
  name = "tf_lbvserver2"
  port = 80
  servicetype = "HTTP"

}
resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

	lbvserverbinding {
        name = citrixadc_lbvserver.tf_lbvserver1.name
        bindpoint = "REQUEST"
        priority = 114
        gotopriorityexpression = "END"
	}
}
`

const testAccRewritepolicy_lbvserverbindings_both = `
resource "citrixadc_lbvserver" "tf_lbvserver1" {

  ipv46 = "10.22.22.22"
  name = "tf_lbvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_lbvserver" "tf_lbvserver2" {

  ipv46 = "10.33.22.66"
  name = "tf_lbvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

	lbvserverbinding {
        name = citrixadc_lbvserver.tf_lbvserver1.name
        bindpoint = "RESPONSE"
        priority = 114
        gotopriorityexpression = "END"
	}

	lbvserverbinding {
        name = citrixadc_lbvserver.tf_lbvserver2.name
        bindpoint = "REQUEST"
        priority = 114
        gotopriorityexpression = "END"
	}
}
`

const testAccRewritepolicy_csvserverbindings_both = `
resource "citrixadc_csvserver" "tf_csvserver1" {

  ipv46 = "192.168.45.66"
  name = "tf_csvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_csvserver" "tf_csvserver2" {

  ipv46 = "192.168.45.67"
  name = "tf_csvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

	csvserverbinding {
        name = citrixadc_csvserver.tf_csvserver1.name
        bindpoint = "RESPONSE"
        priority = 114
        gotopriorityexpression = "END"
	}

	csvserverbinding {
        name = citrixadc_csvserver.tf_csvserver2.name
        bindpoint = "REQUEST"
        priority = 114
        gotopriorityexpression = "END"
	}
}
`

const testAccRewritepolicy_csvserverbindings_one = `
resource "citrixadc_csvserver" "tf_csvserver1" {

  ipv46 = "192.168.45.66"
  name = "tf_csvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_csvserver" "tf_csvserver2" {

  ipv46 = "192.168.45.67"
  name = "tf_csvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

	csvserverbinding {
        name = citrixadc_csvserver.tf_csvserver2.name
        bindpoint = "REQUEST"
        priority = 114
        gotopriorityexpression = "END"
	}
}
`

const testAccRewritepolicy_csvserverbindings_none = `
resource "citrixadc_csvserver" "tf_csvserver1" {

  ipv46 = "192.168.45.66"
  name = "tf_csvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_csvserver" "tf_csvserver2" {

  ipv46 = "192.168.45.67"
  name = "tf_csvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

}
`

func TestAccRewritepolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_rewritepolicy.tf_rewrite_policy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewritepolicy_globalbinding_not_exists,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRewritepolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Rewritepolicy.Type(), "tf_rewrite_policy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccRewritepolicy_globalbinding_not_exists,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRewritepolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccRewritepolicy_import(t *testing.T) {
	const resAddr = "citrixadc_rewritepolicy.tf_rewrite_policy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccRewritepolicy_globalbinding_not_exists},
			{
				Config:                  testAccRewritepolicy_globalbinding_not_exists,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccRewritepolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccRewritepolicy_globalbinding_not_exists,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccRewritepolicy_globalbinding_not_exists,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
				),
			},
		},
	})
}

// Unset test: undefaction and logaction are the type-independent unset-eligible
// scalar attributes of rewritepolicy (per the NITRO spec unset operation).
// comment is excluded: NITRO reverts it to absent (no server default returned by
// GET), so it cannot carry a stable schema Default and is not safely unsettable.
const testAccRewritepolicy_unset_step1 = `
resource "citrixadc_auditmessageaction" "tf_msg" {
	name              = "tf_rw_unset_msg"
	loglevel          = "INFORMATIONAL"
	stringbuilderexpr = "\"probe\""
}

resource "citrixadc_rewritepolicy" "tf_unset" {
	name        = "tf_rewritepolicy_unset"
	action      = "DROP"
	rule        = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
	undefaction = "RESET"
	logaction   = citrixadc_auditmessageaction.tf_msg.name
}
`

const testAccRewritepolicy_unset_step2 = `
resource "citrixadc_auditmessageaction" "tf_msg" {
	name              = "tf_rw_unset_msg"
	loglevel          = "INFORMATIONAL"
	stringbuilderexpr = "\"probe\""
}

resource "citrixadc_rewritepolicy" "tf_unset" {
	name   = "tf_rewritepolicy_unset"
	action = "DROP"
	rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
	# undefaction and logaction removed from config -> provider must unset them
	# (revert to NITRO defaults "Use Global" / "None").
}
`

func TestAccRewritepolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccRewritepolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rewritepolicy.tf_unset", "undefaction", "RESET"),
					resource.TestCheckResourceAttr("citrixadc_rewritepolicy.tf_unset", "logaction", "tf_rw_unset_msg"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccRewritepolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rewritepolicy.tf_unset", "undefaction", "Use Global"),
					resource.TestCheckResourceAttr("citrixadc_rewritepolicy.tf_unset", "logaction", "None"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckRewritepolicyADCValue("tf_rewritepolicy_unset", "undefaction", "Use Global"),
					testAccCheckRewritepolicyADCValue("tf_rewritepolicy_unset", "logaction", "None"),
				),
			},
		},
	})
}

// testAccCheckRewritepolicyADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckRewritepolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Rewritepolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("rewritepolicy %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("rewritepolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccRewritepolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRewritepolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_rewritepolicy.test", "name", "tf_rewritepolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_rewritepolicy.test", "action", "DROP"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rewritepolicy.test", "rule"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rewritepolicy.test", "id"),
				),
			},
		},
	})
}

const testAccRewritepolicyDataSource_basic = `
resource "citrixadc_rewritepolicy" "test" {
	name   = "tf_rewritepolicy_ds"
	action = "DROP"
	rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"test\")"
}

data "citrixadc_rewritepolicy" "test" {
	name = citrixadc_rewritepolicy.test.name
}
`

// The three tests below each edit only a NON-KEY sub-attribute
// (gotopriorityexpression) of a convenience-block binding while its diff key is
// unchanged (globalbinding: type+priority; lb/csvserverbinding:
// name+bindpoint+priority). A key-only reconciliation would silently drop the
// edit, and the post-apply refresh would then disagree with the plan
// ("inconsistent result after apply"). The second step asserts the edit lands.

func TestAccRewritepolicy_globalbinding_editgotopri(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewritepolicy_globalbinding_editgotopri("END"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					resource.TestCheckTypeSetElemNestedAttrs(
						"citrixadc_rewritepolicy.tf_rewrite_policy", "globalbinding.*",
						map[string]string{"type": "REQ_OVERRIDE", "priority": "666", "gotopriorityexpression": "END"}),
				),
			},
			{
				Config: testAccRewritepolicy_globalbinding_editgotopri("NEXT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					resource.TestCheckTypeSetElemNestedAttrs(
						"citrixadc_rewritepolicy.tf_rewrite_policy", "globalbinding.*",
						map[string]string{"type": "REQ_OVERRIDE", "priority": "666", "gotopriorityexpression": "NEXT"}),
				),
			},
		},
	})
}

func TestAccRewritepolicy_lbvserverbinding_editgotopri(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewritepolicy_lbvserverbinding_editgotopri("END"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					resource.TestCheckTypeSetElemNestedAttrs(
						"citrixadc_rewritepolicy.tf_rewrite_policy", "lbvserverbinding.*",
						map[string]string{"name": "tf_lbvserver1", "bindpoint": "REQUEST", "priority": "114", "gotopriorityexpression": "END"}),
				),
			},
			{
				Config: testAccRewritepolicy_lbvserverbinding_editgotopri("NEXT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					resource.TestCheckTypeSetElemNestedAttrs(
						"citrixadc_rewritepolicy.tf_rewrite_policy", "lbvserverbinding.*",
						map[string]string{"name": "tf_lbvserver1", "bindpoint": "REQUEST", "priority": "114", "gotopriorityexpression": "NEXT"}),
				),
			},
		},
	})
}

func TestAccRewritepolicy_csvserverbinding_editgotopri(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewritepolicy_csvserverbinding_editgotopri("END"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					resource.TestCheckTypeSetElemNestedAttrs(
						"citrixadc_rewritepolicy.tf_rewrite_policy", "csvserverbinding.*",
						map[string]string{"name": "tf_csvserver1", "bindpoint": "REQUEST", "priority": "114", "gotopriorityexpression": "END"}),
				),
			},
			{
				Config: testAccRewritepolicy_csvserverbinding_editgotopri("NEXT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					resource.TestCheckTypeSetElemNestedAttrs(
						"citrixadc_rewritepolicy.tf_rewrite_policy", "csvserverbinding.*",
						map[string]string{"name": "tf_csvserver1", "bindpoint": "REQUEST", "priority": "114", "gotopriorityexpression": "NEXT"}),
				),
			},
		},
	})
}

func testAccRewritepolicy_globalbinding_editgotopri(gotopri string) string {
	return fmt.Sprintf(`
resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

  globalbinding {
    type                   = "REQ_OVERRIDE"
    priority               = 666
    gotopriorityexpression = "%s"
  }
}
`, gotopri)
}

func testAccRewritepolicy_lbvserverbinding_editgotopri(gotopri string) string {
	return fmt.Sprintf(`
resource "citrixadc_lbvserver" "tf_lbvserver1" {
  ipv46       = "10.22.22.22"
  name        = "tf_lbvserver1"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

  lbvserverbinding {
    name                   = citrixadc_lbvserver.tf_lbvserver1.name
    bindpoint              = "REQUEST"
    priority               = 114
    gotopriorityexpression = "%s"
  }
}
`, gotopri)
}

func testAccRewritepolicy_csvserverbinding_editgotopri(gotopri string) string {
	return fmt.Sprintf(`
resource "citrixadc_csvserver" "tf_csvserver1" {
  ipv46       = "192.168.45.66"
  name        = "tf_csvserver1"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

  csvserverbinding {
    name                   = citrixadc_csvserver.tf_csvserver1.name
    bindpoint              = "REQUEST"
    priority               = 114
    gotopriorityexpression = "%s"
  }
}
`, gotopri)
}
