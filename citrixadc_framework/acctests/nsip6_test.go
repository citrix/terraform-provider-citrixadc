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
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccNsip6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsip6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsip6_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_nsip6", nil, "2002:db8:100::ff/64"),
				),
			},
			{
				Config: testAccNsip6_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_nsip6", nil, "2002:db8:100::ff/64"),
				),
			},
			{
				Config: testAccNsip6_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_nsip6", nil, "2002:db8:100::ff/64"),
				),
			},
		},
	})
}

const testAccNsip6_mptcpadvertise = `
	resource "citrixadc_nsip6" "tf_test_nsip6_mptcpadvertise" {
		ipv6address = "2002:db8:100::ff/64"
		type = "VIP"
		icmp = "ENABLED"
		mptcpadvertise = "YES"
	}
`

func TestAccNsip6_mptcpadvertise(t *testing.T) {
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsip6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsip6_mptcpadvertise,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_test_nsip6_mptcpadvertise", nil, "2002:db8:100::ff/64"),
					resource.TestCheckResourceAttr("citrixadc_nsip6.tf_test_nsip6_mptcpadvertise", "mptcpadvertise", "YES"),
				),
			},
		},
	})
}

func testAccCheckNsip6Exist(n string, id *string, ipv6address string) resource.TestCheckFunc {
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

		array, _ := client.FindAllResources(service.Nsip6.Type())

		foundAddress := false
		for _, item := range array {
			if item["ipv6address"] == ipv6address {
				foundAddress = true
				break
			}
		}
		if !foundAddress {
			return errors.New("Cannot find resource nsip6 with ipv6address %v")
		}

		return nil
	}
}

func testAccCheckNsip6Destroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsip6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsip6.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNsip6_basic_step1 = `

resource "citrixadc_nsip6" "tf_nsip6" {
    ipv6address = "2002:db8:100::ff/64"
    type = "VIP"
    icmp = "DISABLED"
}
`

const testAccNsip6_basic_step2 = `

resource "citrixadc_nsip6" "tf_nsip6" {
    ipv6address = "2002:db8:100::ff/64"
    type = "VIP"
    icmp = "ENABLED"
}
`

const testAccNsip6_basic_step3 = `

resource "citrixadc_nsip6" "tf_nsip6" {
    ipv6address = "2002:db8:100::ff/64"
    type = "SNIP"
    icmp = "ENABLED"
}
`

const testAccNsip6DataSource_basic = `

resource "citrixadc_nsip6" "tf_nsip6_ds" {
    ipv6address = "2002:db8:100::aa/64"
    type = "VIP"
    icmp = "ENABLED"
    nd = "ENABLED"
    state = "ENABLED"
}

data "citrixadc_nsip6" "tf_nsip6_datasource" {
    ipv6address = citrixadc_nsip6.tf_nsip6_ds.ipv6address
    td = 0
}
`

func TestAccNsip6_import(t *testing.T) {
	const resAddr = "citrixadc_nsip6.tf_nsip6"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsip6Destroy,
		Steps: []resource.TestStep{
			{Config: testAccNsip6_basic_step1},
			{
				Config:                  testAccNsip6_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNsip6_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsip6Destroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsip6_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_nsip6", nil, "2002:db8:100::ff/64"),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNsip6_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_nsip6", nil, "2002:db8:100::ff/64"),
				),
			},
		},
	})
}

// The nsip6 unset test covers the type-independent, unset-eligible attributes
// whose static NITRO default (DISABLED) holds regardless of IP type. The
// management-access "allow" attributes (ftp/gui/ssh/snmp/telnet) are excluded:
// they default to ENABLED but the appliance forces them to DISABLED on a VIP,
// so a static ENABLED default would be type-dependent and cannot be wired
// safely. mgmtaccess and restrictaccess default to DISABLED (matching the
// appliance on every IP type) and can be set on a SNIP6 and unset cleanly.
const testAccNsip6_unset_step1 = `
resource "citrixadc_nsip6" "tf_unset" {
    ipv6address    = "2002:db8:200::ff/64"
    type           = "SNIP"
    mgmtaccess     = "ENABLED"
    restrictaccess = "ENABLED"
}
`

const testAccNsip6_unset_step2 = `
resource "citrixadc_nsip6" "tf_unset" {
    ipv6address = "2002:db8:200::ff/64"
    type        = "SNIP"
    # All unset-eligible attributes removed from config -> the provider must
    # unset them (revert to NITRO defaults).
}
`

func TestAccNsip6_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsip6Destroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNsip6_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_unset", nil, "2002:db8:200::ff/64"),
					resource.TestCheckResourceAttr("citrixadc_nsip6.tf_unset", "mgmtaccess", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsip6.tf_unset", "restrictaccess", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNsip6_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_unset", nil, "2002:db8:200::ff/64"),
					resource.TestCheckResourceAttr("citrixadc_nsip6.tf_unset", "mgmtaccess", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsip6.tf_unset", "restrictaccess", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsip6ADCValue("2002:db8:200::ff/64", "mgmtaccess", "DISABLED"),
					testAccCheckNsip6ADCValue("2002:db8:200::ff/64", "restrictaccess", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckNsip6ADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. The nsip6 address cannot go in the URL path, so match by listing.
func testAccCheckNsip6ADCValue(ipv6address, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		array, err := client.FindAllResources(service.Nsip6.Type())
		if err != nil {
			return err
		}
		for _, item := range array {
			if item["ipv6address"] == ipv6address {
				got := strings.TrimSpace(fmt.Sprintf("%v", item[attr]))
				if got != want {
					return fmt.Errorf("nsip6 %s: appliance attr %q = %q, want %q (unset did not revert it)", ipv6address, attr, got, want)
				}
				return nil
			}
		}
		return fmt.Errorf("nsip6 %s not found on appliance", ipv6address)
	}
}

func TestAccNsip6DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsip6DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsip6.tf_nsip6_datasource", "ipv6address", "2002:db8:100::aa/64"),
					resource.TestCheckResourceAttr("data.citrixadc_nsip6.tf_nsip6_datasource", "icmp", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nsip6.tf_nsip6_datasource", "nd", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nsip6.tf_nsip6_datasource", "state", "ENABLED"),
				),
			},
		},
	})
}
