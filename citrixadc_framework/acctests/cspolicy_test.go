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

func TestAccCspolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCspolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCspolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspolicyExist("citrixadc_cspolicy.foo_cspolicy", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_cspolicy.foo_cspolicy", "csvserver", "tst_policy_cs"),
					resource.TestCheckResourceAttr(
						"citrixadc_cspolicy.foo_cspolicy", "targetlbvserver", "tst_policy_lb"),
					resource.TestCheckResourceAttr(
						"citrixadc_cspolicy.foo_cspolicy", "policyname", "test_policy"),
					resource.TestCheckResourceAttr(
						"citrixadc_cspolicy.foo_cspolicy", "rule", "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.84.0)"),
				),
			},
		},
	})
}

func testAccCheckCspolicyExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := client.FindResource(service.Cspolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckCspolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cspolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cspolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCspolicy_basic = `

resource "citrixadc_csvserver" "foo_cspolicy" {

  ipv46 = "10.202.11.11"
  name = "tst_policy_cs"
  port = 8080
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "foo_cspolicy" {

  name = "tst_policy_lb"
  servicetype = "HTTP"
}

resource "citrixadc_cspolicy" "foo_cspolicy" {
  csvserver = "tst_policy_cs"
  targetlbvserver = "tst_policy_lb"
  policyname = "test_policy"
  rule = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.84.0)"
  priority = 10

  depends_on = ["citrixadc_csvserver.foo_cspolicy", "citrixadc_lbvserver.foo_cspolicy"]

}
`

func TestAccCspolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCspolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccCspolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCspolicyExist("citrixadc_cspolicy.foo_cspolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccCspolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCspolicyExist("citrixadc_cspolicy.foo_cspolicy", nil)),
			},
		},
	})
}

// The cspolicy unset test covers the mutable, spec-unsettable attribute
// logaction (the one attribute the NITRO cspolicy unset operation documents).
// action is NOT unsettable (NITRO rejects it with errorcode 278 "Invalid
// argument"), so it is excluded. Step 1 sets logaction to a valid non-default
// value (a real auditmessageaction); step 2 removes it from config, so the
// provider must issue the NITRO ?action=unset and the appliance reverts it to
// its empty default.
const testAccCspolicy_unset_step1 = `
resource "citrixadc_auditmessageaction" "tf_unset_msg" {
  name              = "tf_cspol_unset_msg"
  loglevel          = "NOTICE"
  stringbuilderexpr = "\"hello\""
}

resource "citrixadc_cspolicy" "tf_unset" {
  policyname = "tf_cspol_unset"
  rule       = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.84.0)"
  logaction  = citrixadc_auditmessageaction.tf_unset_msg.name

  depends_on = [citrixadc_auditmessageaction.tf_unset_msg]
}
`

const testAccCspolicy_unset_step2 = `
resource "citrixadc_auditmessageaction" "tf_unset_msg" {
  name              = "tf_cspol_unset_msg"
  loglevel          = "NOTICE"
  stringbuilderexpr = "\"hello\""
}

resource "citrixadc_cspolicy" "tf_unset" {
  policyname = "tf_cspol_unset"
  rule       = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.84.0)"
  # logaction removed from config -> the provider must unset it.
}
`

func TestAccCspolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCspolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccCspolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspolicyExist("citrixadc_cspolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cspolicy.tf_unset", "logaction", "tf_cspol_unset_msg"),
					testAccCheckCspolicyADCValue("tf_cspol_unset", "logaction", "tf_cspol_unset_msg"),
				),
			},
			{
				// Removing the attributes must unset them: the appliance reverts to
				// the empty defaults and the implicit post-apply plan must be empty.
				Config: testAccCspolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspolicyExist("citrixadc_cspolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cspolicy.tf_unset", "rule", "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.84.0)"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckCspolicyADCValue("tf_cspol_unset", "logaction", ""),
				),
			},
		},
	})
}

// testAccCheckCspolicyADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. A missing/empty attribute reads as "".
func testAccCheckCspolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Cspolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("cspolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("cspolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccCspolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCspolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cspolicy.tf_cspolicy_ds", "policyname", "tf_test_policy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_cspolicy.tf_cspolicy_ds", "rule", "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.85.0)"),
				),
			},
		},
	})
}

func TestAccCspolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_cspolicy.foo_cspolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCspolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCspolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCspolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.UnbindResource(service.Csvserver.Type(), "tst_policy_cs", service.Cspolicy.Type(), "test_policy", "policyname"); err != nil {
						t.Fatalf("self-healing: out-of-band unbind failed: %v", err)
					}
					if err := client.DeleteResource(service.Cspolicy.Type(), "test_policy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCspolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCspolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccCspolicy_import(t *testing.T) {
	const resAddr = "citrixadc_cspolicy.foo_cspolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCspolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCspolicy_basic},
			{
				Config:                  testAccCspolicy_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"csvserver", "priority", "targetlbvserver"},
			},
		},
	})
}

const testAccCspolicyDataSource_basic = `

resource "citrixadc_csvserver" "tf_cspolicy_ds" {
  ipv46       = "10.202.11.12"
  name        = "tst_policy_cs_ds"
  port        = 8081
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_cspolicy_ds" {
  name        = "tst_policy_lb_ds"
  servicetype = "HTTP"
}

resource "citrixadc_cspolicy" "tf_cspolicy_ds" {
  csvserver       = citrixadc_csvserver.tf_cspolicy_ds.name
  targetlbvserver = citrixadc_lbvserver.tf_cspolicy_ds.name
  policyname      = "tf_test_policy_ds"
  rule            = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.85.0)"
  priority        = 10

  depends_on = [citrixadc_csvserver.tf_cspolicy_ds, citrixadc_lbvserver.tf_cspolicy_ds]
}

data "citrixadc_cspolicy" "tf_cspolicy_ds" {
  policyname = citrixadc_cspolicy.tf_cspolicy_ds.policyname
  depends_on = [citrixadc_cspolicy.tf_cspolicy_ds]
}

`
