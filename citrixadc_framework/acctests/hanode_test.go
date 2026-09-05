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

const testAccHanodeLocal_basic = `
  
resource "citrixadc_hanode" "local_node" {
	hanode_id     = 0
	hellointerval = 200
	deadinterval = 5
	}
   
`
const testAccHanodeLocal_update = `
	resource "citrixadc_hanode" "local_node" {
		hanode_id     = 0
		hellointerval = 400
		deadinterval = 30
	}
	
`

func TestAccHanodeLocal_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccHanodeLocal_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.local_node", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "hanode_id", "0"),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "hellointerval", "200"),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "deadinterval", "5"),
				),
			},
			{
				Config: testAccHanodeLocal_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.local_node", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "hanode_id", "0"),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "hellointerval", "400"),
					resource.TestCheckResourceAttr("citrixadc_hanode.local_node", "deadinterval", "30"),
				),
			},
		},
	})
}

const testAccHanodeRemote_basic = `
  
resource "citrixadc_hanode" "remote_node" {
	hanode_id = 2
	ipaddress = "10.222.74.145"
	}
  
   
`
const testAccHanodeRemote_update = `
	resource "citrixadc_hanode" "remote_node" {
		hanode_id = 3
		ipaddress = "10.222.74.145"
	}
	
`

func TestAccHanodeRemote_basic(t *testing.T) {
	if adcTestbed != "HA" {
		t.Skipf("ADC testbed is %s. Expected HA.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHanodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHanodeRemote_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.remote_node", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.remote_node", "hanode_id", "2"),
					resource.TestCheckResourceAttr("citrixadc_hanode.remote_node", "ipaddress", "10.222.74.145"),
				),
			},
			{
				Config: testAccHanodeRemote_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.remote_node", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.remote_node", "hanode_id", "3"),
					resource.TestCheckResourceAttr("citrixadc_hanode.remote_node", "ipaddress", "10.222.74.145"),
				),
			},
		},
	})
}

func testAccCheckHanodeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No hanode name is set")
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
		data, err := client.FindResource(service.Hanode.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("hanode %s not found", n)
		}

		return nil
	}
}

const testAccHanodeLocalDataSource_basic = `
  
resource "citrixadc_hanode" "local_node" {
	hanode_id     = 0
	hellointerval = 200
	deadinterval = 5
}

data "citrixadc_hanode" "local_node" {
	hanode_id = citrixadc_hanode.local_node.hanode_id
	depends_on = [citrixadc_hanode.local_node]
}
   
`

func TestAccHanodeLocalDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccHanodeLocalDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					// Universal runtime-binding proof that the data source resolved.
					resource.TestCheckResourceAttrSet("data.citrixadc_hanode.local_node", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_hanode.local_node", "hanode_id", "0"),
					resource.TestCheckResourceAttr("data.citrixadc_hanode.local_node", "hellointerval", "200"),
					resource.TestCheckResourceAttr("data.citrixadc_hanode.local_node", "deadinterval", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_hanode.local_node", "hasync", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_hanode.local_node", "inc", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_hanode.local_node", "state", "Primary"),
					resource.TestCheckResourceAttr("data.citrixadc_hanode.local_node", "haprop", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_hanode.local_node", "failsafe", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_hanode.local_node", "syncstatusstrictmode", "DISABLED"),
				),
			},
		},
	})
}

func TestAccHanode_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_hanode.local_node"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccHanodeLocal_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckHanodeExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Hanode.Type(), "0"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccHanodeLocal_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckHanodeExist(resAddr, nil)),
			},
		},
	})
}

func TestAccHanode_import(t *testing.T) {
	const resAddr = "citrixadc_hanode.local_node"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccHanodeLocal_basic},
			{
				Config:            testAccHanodeLocal_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// masterstatetime is volatile; rpcnodepassword_wo_version defaults to 1
				// on create but is a config-only tracker NITRO never returns on import.
				ImportStateVerifyIgnore: []string{"masterstatetime", "rpcnodepassword_wo_version"},
			},
		},
	})
}

func TestAccHanode_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		// CheckDestroy is nil: hanode 0 is the permanent local self-node and can
		// never be deleted, so the standard Destroy check always reports it as
		// "still exists". The upgrade itself (2.2.0 -> current) is what we verify.
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccHanodeLocal_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckHanodeExist("citrixadc_hanode.local_node", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccHanodeLocal_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckHanodeExist("citrixadc_hanode.local_node", nil)),
			},
		},
	})
}

// The hanode unset test runs against the permanent local self-node (id 0) so it
// works on a standalone appliance. Step 1 sets the unset-eligible mutable
// attributes to non-default values; step 2 removes them from config, which must
// unset them back to their documented NITRO defaults.
const testAccHanode_unset_step1 = `
resource "citrixadc_hanode" "tf_unset" {
	hanode_id            = 0
	deadinterval         = 30
	failsafe             = "ON"
	haprop               = "DISABLED"
	hasync               = "DISABLED"
	hellointerval        = 400
	maxflips             = 5
	maxfliptime          = 60
	syncstatusstrictmode = "ENABLED"
}
`

const testAccHanode_unset_step2 = `
resource "citrixadc_hanode" "tf_unset" {
	hanode_id = 0
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccHanode_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy is nil: hanode 0 is the permanent local self-node and can
		// never be deleted.
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccHanode_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "deadinterval", "30"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "failsafe", "ON"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "haprop", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "hasync", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "hellointerval", "400"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "maxflips", "5"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "maxfliptime", "60"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "syncstatusstrictmode", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccHanode_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "deadinterval", "3"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "failsafe", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "haprop", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "hasync", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "hellointerval", "200"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "maxflips", "0"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "maxfliptime", "0"),
					resource.TestCheckResourceAttr("citrixadc_hanode.tf_unset", "syncstatusstrictmode", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckHanodeADCValue("0", "deadinterval", "3"),
					testAccCheckHanodeADCValue("0", "hellointerval", "200"),
					testAccCheckHanodeADCValue("0", "failsafe", "OFF"),
				),
			},
		},
	})
}

// testAccCheckHanodeADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckHanodeADCValue(id, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Hanode.Type(), id)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("hanode %s not found on appliance", id)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("hanode %s: appliance attr %q = %q, want %q (unset did not revert it)", id, attr, got, want)
		}
		return nil
	}
}

func testAccCheckHanodeDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_hanode" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Hanode.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("hanode %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// --- rpcnodepassword write-only (ephemeral) support (GH #1441) ---
//
// rpcnodepassword is create-only (ForceNew) on hanode, so its write-only twin
// rpcnodepassword_wo pairs with rpcnodepassword_wo_version, and a version bump is
// RequiresReplace (rotation re-adds the peer node). These exercise a remote peer
// node and therefore require an HA testbed, matching TestAccHanodeRemote_basic.

// Backward-compatible path: the plain (state-persisted) rpcnodepassword still works.
const testAccHanodeRemote_rpcnodepassword_backward_compat = `
	variable "hanode_rpcnodepassword" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_hanode" "remote_node" {
		hanode_id       = 2
		ipaddress       = "10.222.74.145"
		rpcnodepassword = var.hanode_rpcnodepassword
	}
`

func TestAccHanodeRemote_rpcnodepassword_backward_compat(t *testing.T) {
	if adcTestbed != "HA" {
		t.Skipf("ADC testbed is %s. Expected HA.", adcTestbed)
	}
	t.Setenv("TF_VAR_hanode_rpcnodepassword", "rpcpass1")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHanodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHanodeRemote_rpcnodepassword_backward_compat,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.remote_node", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.remote_node", "hanode_id", "2"),
				),
			},
		},
	})
}

// Ephemeral path: rpcnodepassword_wo (WriteOnly, never persisted) + version tracker.
// Bumping rpcnodepassword_wo_version forces a replace (rpcnodepassword is create-only).
const testAccHanodeRemote_rpcnodepassword_wo_step1 = `
	variable "hanode_rpcnodepassword_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_hanode" "remote_node" {
		hanode_id                  = 2
		ipaddress                  = "10.222.74.145"
		rpcnodepassword_wo         = var.hanode_rpcnodepassword_wo
		rpcnodepassword_wo_version = 1
	}
`

const testAccHanodeRemote_rpcnodepassword_wo_step2 = `
	variable "hanode_rpcnodepassword_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_hanode" "remote_node" {
		hanode_id                  = 2
		ipaddress                  = "10.222.74.145"
		rpcnodepassword_wo         = var.hanode_rpcnodepassword_wo_2
		rpcnodepassword_wo_version = 2
	}
`

func TestAccHanodeRemote_rpcnodepassword_wo_ephemeral(t *testing.T) {
	if adcTestbed != "HA" {
		t.Skipf("ADC testbed is %s. Expected HA.", adcTestbed)
	}
	t.Setenv("TF_VAR_hanode_rpcnodepassword_wo", "ephem_rpc1")
	t.Setenv("TF_VAR_hanode_rpcnodepassword_wo_2", "ephem_rpc2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHanodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHanodeRemote_rpcnodepassword_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.remote_node", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.remote_node", "rpcnodepassword_wo_version", "1"),
					// The write-only value must never be persisted in state.
					resource.TestCheckNoResourceAttr("citrixadc_hanode.remote_node", "rpcnodepassword_wo"),
				),
			},
			{
				// Bump the version (and value) -> RequiresReplace re-adds the node.
				Config: testAccHanodeRemote_rpcnodepassword_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanodeExist("citrixadc_hanode.remote_node", nil),
					resource.TestCheckResourceAttr("citrixadc_hanode.remote_node", "rpcnodepassword_wo_version", "2"),
				),
			},
		},
	})
}
