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

// NOTE: These tests manage actual cluster membership and are gated to the CLUSTER testbed.
// They add a *self-contained, non-destructive* spare node definition (an UNUSED nodeid with a
// phantom NSIP that is not a live appliance). nodeid=2 is used deliberately: on the CLUSTER
// testbed nodeids 0 and 1 are the existing UP members (node 1 = local config-coordinator that
// carries the management/CLIP connection), so touching them would break quorum and sever the
// test's own connection. Adding nodeid=2 with a phantom IP creates a config-only node that never
// joins (stays DOWN) and is cleanly removable, leaving the live cluster untouched.
const testAccClusternode_basic_nogroup_config = `


resource "citrixadc_clusternode" "tf_clusternode" {
	nodeid             = 2
	ipaddress          = "10.101.132.153"
	state              = "PASSIVE"
	}
  
`
const testAccClusternode_update_nogroup_config = `


resource "citrixadc_clusternode" "tf_clusternode" {
	nodeid             = 2
	ipaddress          = "10.101.132.153"
	state              = "ACTIVE"  
	}
  
`

func TestAccClusternode_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternode_basic_nogroup_config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.101.132.153"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "PASSIVE"),
				),
			},
			{
				Config: testAccClusternode_update_nogroup_config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.101.132.153"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "ACTIVE"),
				),
			},
		},
	})
}

const testAccClusternode_basic_group_config_yes = `


	resource "citrixadc_clusternode" "tf_clusternode" {
		nodeid               = 2
		ipaddress            = "10.101.132.153"
		state                = "PASSIVE"
		clearnodegroupconfig = "YES"
	}
`
const testAccClusternode_update_group_config_yes = `


	resource "citrixadc_clusternode" "tf_clusternode" {
		nodeid               = 2
		ipaddress            = "10.101.132.153"
		state                = "ACTIVE"
		clearnodegroupconfig = "YES"
	}
  
`

func TestAccClusternode_group_config_yes(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternode_basic_group_config_yes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.101.132.153"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "PASSIVE"),
				),
			},
			{
				Config: testAccClusternode_update_group_config_yes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.101.132.153"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "ACTIVE"),
				),
			},
		},
	})
}

const testAccClusternode_basic_group_config_no = `


	resource "citrixadc_clusternode" "tf_clusternode" {
		nodeid               = 2
		ipaddress            = "10.101.132.153"
		state                = "PASSIVE"
		clearnodegroupconfig = "NO"
	}
`
const testAccClusternode_update_group_config_no = `


	resource "citrixadc_clusternode" "tf_clusternode" {
		nodeid               = 2
		ipaddress            = "10.101.132.153"
		state                = "ACTIVE"
		clearnodegroupconfig = "NO"
	}
  
`

func TestAccClusternode_group_config_no(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternode_basic_group_config_no,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.101.132.153"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "PASSIVE"),
				),
			},
			{
				Config: testAccClusternode_update_group_config_no,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "ipaddress", "10.101.132.153"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_clusternode", "state", "ACTIVE"),
				),
			},
		},
	})
}

func testAccCheckClusternodeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusternode name is set")
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
		data, err := client.FindResource(service.Clusternode.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("clusternode %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusternodeDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusternode" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Clusternode.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusternode %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccClusternodeDataSource_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodeDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_clusternode.tf_clusternode_ds", "nodeid", "0"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternode.tf_clusternode_ds", "ipaddress", "10.101.132.123"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternode.tf_clusternode_ds", "state", "ACTIVE"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternode.tf_clusternode_ds", "nodegroup", "DEFAULT_NG"),
				),
			},
		},
	})
}

func TestAccClusternode_selfHealing(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	const resAddr = "citrixadc_clusternode.tf_clusternode"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternode_basic_nogroup_config,
				Check:  resource.ComposeTestCheckFunc(testAccCheckClusternodeExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Clusternode.Type(), "2", []string{"clearnodegroupconfig:YES"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccClusternode_basic_nogroup_config,
				Check:  resource.ComposeTestCheckFunc(testAccCheckClusternodeExist(resAddr, nil)),
			},
		},
	})
}

func TestAccClusternode_import(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	const resAddr = "citrixadc_clusternode.tf_clusternode"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodeDestroy,
		Steps: []resource.TestStep{
			{Config: testAccClusternode_basic_nogroup_config},
			{
				Config:                  testAccClusternode_basic_nogroup_config,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccClusternode_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckClusternodeDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccClusternode_basic_nogroup_config,
				Check:  resource.ComposeTestCheckFunc(testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccClusternode_basic_nogroup_config,
				Check:  resource.ComposeTestCheckFunc(testAccCheckClusternodeExist("citrixadc_clusternode.tf_clusternode", nil)),
			},
		},
	})
}

// testAccClusternode_unset exercises the NITRO unset op for the mutable,
// spec-unsettable attributes with a documented server default: state
// (PASSIVE), priority (31) and tunnelmode (NONE). backplane is in
// the unset payload but has no documented server default, so it is excluded.
// delay is intentionally NOT exercised here: NITRO only honors delay for a
// PASSIVE node, so on this ACTIVE node it is silently kept at 0 (plan 5 vs
// applied 0) -- an appliance semantic, not a provider defect.
// nodeid=2 with a phantom NSIP is used for the same isolation reasons as the
// basic test (a config-only node that never joins the live cluster).
const testAccClusternode_unset_step1 = `


resource "citrixadc_clusternode" "tf_unset" {
	nodeid     = 2
	ipaddress  = "10.101.132.153"
	state      = "ACTIVE"
	priority   = 15
	tunnelmode = "GRE"
}
`

const testAccClusternode_unset_step2 = `


resource "citrixadc_clusternode" "tf_unset" {
	nodeid    = 2
	ipaddress = "10.101.132.153"
	# state, priority and tunnelmode removed from config -> the provider
	# must unset them (revert to NITRO defaults).
}
`

func TestAccClusternode_unset(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodeDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccClusternode_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_unset", "state", "ACTIVE"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_unset", "priority", "15"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_unset", "tunnelmode", "GRE"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccClusternode_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodeExist("citrixadc_clusternode.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_unset", "state", "PASSIVE"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_unset", "priority", "31"),
					resource.TestCheckResourceAttr("citrixadc_clusternode.tf_unset", "tunnelmode", "NONE"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckClusternodeADCValue("2", "priority", "31"),
					testAccCheckClusternodeADCValue("2", "tunnelmode", "NONE"),
				),
			},
		},
	})
}

// testAccCheckClusternodeADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckClusternodeADCValue(nodeid, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Clusternode.Type(), nodeid)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("clusternode %s not found on appliance", nodeid)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("clusternode %s: appliance attr %q = %q, want %q (unset did not revert it)", nodeid, attr, got, want)
		}
		return nil
	}
}

const testAccClusternodeDataSource_basic = `


data "citrixadc_clusternode" "tf_clusternode_ds" {
	nodeid     = "0"
}

`
