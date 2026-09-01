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

const testAccNsacl_basic = `


resource "citrixadc_nsacl" "foo" {

  aclaction = "DENY"
  aclname = "test_acl"
  destipval = "192.168.1.33"
  protocol = "TCP"
  srcportval = "45-1024"
  priority = "100"

}
`

func TestAccNsacl_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaclDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsacl_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("citrixadc_nsacl.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.foo", "aclaction", "DENY"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.foo", "aclname", "test_acl"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.foo", "destipval", "192.168.1.33"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.foo", "protocol", "TCP"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.foo", "srcportval", "45-1024"),
				),
			},
		},
	})
}

func testAccCheckNsaclExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := client.FindResource(service.Nsacl.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsaclDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsacl" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsacl.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNsaclEnableDisable_enabled = `
resource "citrixadc_nsacl" "tf_test_acc_nsacl" {
    aclname = "tf_test_acc_nsacl"
    aclaction = "ALLOW"
    priority = "100"
    srcipval = "192.168.10.22"
    destipval = "172.17.0.20"
    state = "ENABLED"
}
`

const testAccNsaclEnableDisable_disabled = `
resource "citrixadc_nsacl" "tf_test_acc_nsacl" {
    aclname = "tf_test_acc_nsacl"
    aclaction = "ALLOW"
    priority = "99"
    srcipval = "192.168.10.22"
    destipval = "172.17.0.20"
    state = "DISABLED"
}
`

func TestAccNsacl_enable_disable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaclDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccNsaclEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("citrixadc_nsacl.tf_test_acc_nsacl", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacl.tf_test_acc_nsacl", "state", "ENABLED"),
				),
			},
			// Disable
			{
				Config: testAccNsaclEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("citrixadc_nsacl.tf_test_acc_nsacl", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacl.tf_test_acc_nsacl", "state", "DISABLED"),
				),
			},
			// Re enable
			{
				Config: testAccNsaclEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("citrixadc_nsacl.tf_test_acc_nsacl", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacl.tf_test_acc_nsacl", "state", "ENABLED"),
				),
			},
		},
	})
}

const testAccNsacl_dataset = `

resource "citrixadc_policydataset" "tf_srcip_dataset" {
  name    = "tf_srcip_dataset"
  type    = "ipv4"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value1" {
  name = citrixadc_policydataset.tf_srcip_dataset.name
  value    = "192.168.1.1"
}

resource "citrixadc_policydataset" "tf_srcport_dataset" {
  name    = "tf_srcport_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value2" {
  name = citrixadc_policydataset.tf_srcport_dataset.name
  value    = "8080"
}

resource "citrixadc_policydataset" "tf_destip_dataset" {
  name    = "tf_destip_dataset"
  type    = "ipv4"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value3" {
  name = citrixadc_policydataset.tf_destip_dataset.name
  value    = "192.168.1.1"
}

resource "citrixadc_policydataset" "tf_destport_dataset" {
  name    = "tf_destport_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value4" {
  name = citrixadc_policydataset.tf_destport_dataset.name
  value    = "8081"
}

resource "citrixadc_nsacl" "test_acl" {

  aclaction = "DENY"
  aclname = "test_acl"
  srcipdataset = citrixadc_policydataset.tf_srcip_dataset.name
  protocol = "TCP"
  srcportdataset = citrixadc_policydataset.tf_srcport_dataset.name
  priority = "100"
  destipdataset = citrixadc_policydataset.tf_destip_dataset.name
  destportdataset = citrixadc_policydataset.tf_destport_dataset.name

}
`

func TestAccNsacl_dataset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaclDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsacl_dataset,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("citrixadc_nsacl.test_acl", nil),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.test_acl", "srcipdataset", "tf_srcip_dataset"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.test_acl", "srcportdataset", "tf_srcport_dataset"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.test_acl", "destipdataset", "tf_destip_dataset"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.test_acl", "destportdataset", "tf_destport_dataset"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.test_acl", "protocol", "TCP"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.test_acl", "priority", "100"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.test_acl", "aclname", "test_acl"),
				),
			},
		},
	})
}

func TestAccNsacl_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nsacl.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaclDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsacl_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsaclExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nsacl.Type(), "test_acl"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNsacl_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsaclExist(resAddr, nil)),
			},
		},
	})
}

const testAccNsaclDataSource_basic = `

resource "citrixadc_nsacl" "foo" {
  aclaction = "DENY"
  aclname = "test_acl_datasource"
  destipval = "192.168.1.33"
  protocol = "TCP"
  srcportval = "45-1024"
  priority = "100"
}

data "citrixadc_nsacl" "foo" {
  aclname = citrixadc_nsacl.foo.aclname
  type = "CLASSIC"
  depends_on = [citrixadc_nsacl.foo]
}
`

func TestAccNsacl_import(t *testing.T) {
	const resAddr = "citrixadc_nsacl.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaclDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNsacl_basic},
			{
				Config:                  testAccNsacl_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNsacl_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsaclDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccNsacl_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsaclExist("citrixadc_nsacl.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNsacl_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsaclExist("citrixadc_nsacl.foo", nil)),
			},
		},
	})
}

// The nsacl unset test covers the mutable, documented-default attributes wired
// into attributesToUnset: logstate (default DISABLED) and stateful (default NO).
const testAccNsacl_unset_step1 = `
resource "citrixadc_nsacl" "tf_unset" {
  aclname   = "tf_test_acc_nsacl_unset"
  aclaction = "ALLOW"
  priority  = "100"
  srcipval  = "192.168.10.22"
  destipval = "172.17.0.20"
  logstate  = "ENABLED"
  stateful  = "YES"
}
`

const testAccNsacl_unset_step2 = `
resource "citrixadc_nsacl" "tf_unset" {
  aclname   = "tf_test_acc_nsacl_unset"
  aclaction = "ALLOW"
  priority  = "100"
  srcipval  = "192.168.10.22"
  destipval = "172.17.0.20"
  # logstate, stateful removed from config -> provider must unset them (revert
  # to NITRO defaults).
}
`

func TestAccNsacl_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaclDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNsacl_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("citrixadc_nsacl.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacl.tf_unset", "logstate", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsacl.tf_unset", "stateful", "YES"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccNsacl_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("citrixadc_nsacl.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacl.tf_unset", "logstate", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsacl.tf_unset", "stateful", "NO"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsaclADCValue("tf_test_acc_nsacl_unset", "logstate", "DISABLED"),
					testAccCheckNsaclADCValue("tf_test_acc_nsacl_unset", "stateful", "NO"),
				),
			},
		},
	})
}

// testAccCheckNsaclADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckNsaclADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsacl.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsacl %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nsacl %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccNsaclDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsaclDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsacl.foo", "aclname", "test_acl_datasource"),
					resource.TestCheckResourceAttr("data.citrixadc_nsacl.foo", "aclaction", "DENY"),
					resource.TestCheckResourceAttr("data.citrixadc_nsacl.foo", "destipval", "192.168.1.33"),
					resource.TestCheckResourceAttr("data.citrixadc_nsacl.foo", "protocol", "TCP"),
					resource.TestCheckResourceAttr("data.citrixadc_nsacl.foo", "srcportval", "45-1024"),
					resource.TestCheckResourceAttr("data.citrixadc_nsacl.foo", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_nsacl.foo", "type", "CLASSIC"),
					// Universal runtime-binding proof.
					resource.TestCheckResourceAttrSet("data.citrixadc_nsacl.foo", "id"),
					// kernelstate (commit status) is always populated for a created ACL.
					resource.TestCheckResourceAttrSet("data.citrixadc_nsacl.foo", "kernelstate"),
				),
			},
		},
	})
}
