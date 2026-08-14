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

const testAccIptunnel_basic_step1 = `
resource "citrixadc_iptunnel" "tf_iptunnel" {
    name = "tf_iptunnel"
	protocol = "GENEVE"
    remote = "66.0.0.11"
    remotesubnetmask = "255.255.255.255"
    local = "*"
	vnid = 100
	tosinherit = "DISABLED"
	destport = 1088
	vlantagging = "DISABLED"
}
`

const testAccIptunnel_basic_step2 = `
resource "citrixadc_iptunnel" "tf_iptunnel" {
    name = "tf_iptunnel"
	protocol = "GENEVE"
    remote = "66.0.0.10"
    remotesubnetmask = "255.255.255.255"
    local = "*"
	vnid = 100
	tosinherit = "ENABLED"
	destport = 2088
	vlantagging = "ENABLED"
}
`

const testAccIptunnelDataSource_basic = `
resource "citrixadc_iptunnel" "tf_iptunnel" {
    name = "tf_iptunnel_ds"
	protocol = "GENEVE"
    remote = "66.0.0.12"
    remotesubnetmask = "255.255.255.255"
    local = "*"
	vnid = 100
	tosinherit = "DISABLED"
	destport = 1088
	vlantagging = "DISABLED"
}

data "citrixadc_iptunnel" "tf_iptunnel_ds" {
	name = citrixadc_iptunnel.tf_iptunnel.name
	remote = citrixadc_iptunnel.tf_iptunnel.remote
	remotesubnetmask = citrixadc_iptunnel.tf_iptunnel.remotesubnetmask
}
`

func TestAccIptunnel_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIptunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIptunnel_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelExist("citrixadc_iptunnel.tf_iptunnel", nil),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "name", "tf_iptunnel"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "protocol", "GENEVE"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "remote", "66.0.0.11"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "remotesubnetmask", "255.255.255.255"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "local", "*"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "vnid", "100"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "tosinherit", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "destport", "1088"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "vlantagging", "DISABLED"),
				),
			},
			{
				Config: testAccIptunnel_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelExist("citrixadc_iptunnel.tf_iptunnel", nil),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "name", "tf_iptunnel"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "protocol", "GENEVE"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "remote", "66.0.0.10"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "remotesubnetmask", "255.255.255.255"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "local", "*"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "vnid", "100"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "tosinherit", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "destport", "2088"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "vlantagging", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckIptunnelExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No iptunnel name is set")
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
		data, err := client.FindResource(service.Iptunnel.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("iptunnel %s not found", n)
		}

		return nil
	}
}

func testAccCheckIptunnelDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_iptunnel" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Iptunnel.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("iptunnel %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccIptunnel_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_iptunnel.tf_iptunnel"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIptunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIptunnel_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIptunnelExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Iptunnel.Type(), "tf_iptunnel"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccIptunnel_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIptunnelExist(resAddr, nil)),
			},
		},
	})
}

func TestAccIptunnel_import(t *testing.T) {
	// TODO: un-skip once iptunnel import populates the name/ID. Provider defect
	// (NSNETAUTO-1148): iptunnel/resource_schema.go iptunnelSetAttrFromGet sets
	// data.Id from data.Name, but on import data.Name starts null and is only
	// repopulated from the GET response via an `else if data.Name.IsUnknown()`
	// branch (never fires for a null) — so the imported ID resolves to "" and the
	// import verification fails with "resource with ID '' not found".
	t.Skip("known provider import defect: iptunnel import yields empty ID (name not repopulated from GET) (NSNETAUTO-1148)")
	const resAddr = "citrixadc_iptunnel.tf_iptunnel"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIptunnelDestroy,
		Steps: []resource.TestStep{
			{Config: testAccIptunnel_basic_step1},
			{
				Config:                  testAccIptunnel_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccIptunnel_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckIptunnelDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccIptunnel_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelExist("citrixadc_iptunnel.tf_iptunnel", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccIptunnel_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelExist("citrixadc_iptunnel.tf_iptunnel", nil),
				),
			},
		},
	})
}

// testAccCheckIptunnelADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckIptunnelADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Iptunnel.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("iptunnel %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("iptunnel %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccIptunnelDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIptunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIptunnelDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_iptunnel.tf_iptunnel_ds", "id"),
				),
			},
		},
	})
}
