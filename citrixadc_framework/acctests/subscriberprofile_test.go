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

const testAccSubscriberprofile_basic = `


resource "citrixadc_subscriberprofile" "tf_subscriberprofile" {
	ip                  = "10.222.74.185"
	vlan                = 1
	}
  
`

const testAccSubscriberprofileDataSource_basic = `

resource "citrixadc_subscriberprofile" "tf_subscriberprofile" {
	ip                  = "10.222.74.185"
	vlan                = 1
}

data "citrixadc_subscriberprofile" "tf_subscriberprofile" {
	ip   = citrixadc_subscriberprofile.tf_subscriberprofile.ip
	vlan = citrixadc_subscriberprofile.tf_subscriberprofile.vlan
	depends_on = [citrixadc_subscriberprofile.tf_subscriberprofile]
}
`

const testAccSubscriberprofile_update = `


resource "citrixadc_subscriberprofile" "tf_subscriberprofile" {
	ip                  = "10.222.74.185"
	vlan                = 1
	}
  
`

func TestAccSubscriberprofile_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubscriberprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriberprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberprofileExist("citrixadc_subscriberprofile.tf_subscriberprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "ip", "10.222.74.185"),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "subscriptionidtype", "E164"),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "subscriptionidvalue", "5"),
				),
			},
			{
				Config: testAccSubscriberprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberprofileExist("citrixadc_subscriberprofile.tf_subscriberprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "ip", "10.222.74.185"),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "subscriptionidtype", "IMSI"),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_subscriberprofile", "subscriptionidvalue", "10"),
				),
			},
		},
	})
}

func testAccCheckSubscriberprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No subscriberprofile name is set")
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
		data, err := client.FindResource("subscriberprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("subscriberprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSubscriberprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_subscriberprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("subscriberprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("subscriberprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSubscriberprofile_selfHealing(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	const resAddr = "citrixadc_subscriberprofile.tf_subscriberprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubscriberprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriberprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSubscriberprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Subscriberprofile.Type(), "10.222.74.185", []string{"vlan:1"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSubscriberprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSubscriberprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccSubscriberprofile_import(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	const resAddr = "citrixadc_subscriberprofile.tf_subscriberprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubscriberprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSubscriberprofile_basic},
			{
				Config:                  testAccSubscriberprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSubscriberprofile_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSubscriberprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSubscriberprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSubscriberprofileExist("citrixadc_subscriberprofile.tf_subscriberprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSubscriberprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSubscriberprofileExist("citrixadc_subscriberprofile.tf_subscriberprofile", nil)),
			},
		},
	})
}

// servicepath is the only attribute the NITRO subscriberprofile spec documents
// as unsettable (it appears in the ?action=unset payload example). step1 sets it
// to a non-default value; step2 removes it, so the provider must issue the unset
// and the appliance reverts it to its default (empty).
const testAccSubscriberprofile_unset_step1 = `
resource "citrixadc_subscriberprofile" "tf_unset" {
	ip          = "10.222.74.186"
	vlan        = 1
	servicepath = "tf_test_servicepath"
}
`

const testAccSubscriberprofile_unset_step2 = `
resource "citrixadc_subscriberprofile" "tf_unset" {
	ip   = "10.222.74.186"
	vlan = 1
	# servicepath removed from config -> the provider must unset it
	# (revert to the NITRO default, empty).
}
`

func TestAccSubscriberprofile_unset(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubscriberprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccSubscriberprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberprofileExist("citrixadc_subscriberprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_unset", "servicepath", "tf_test_servicepath"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the NITRO default, and the implicit
				// post-apply plan must be empty.
				Config: testAccSubscriberprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberprofileExist("citrixadc_subscriberprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberprofile.tf_unset", "servicepath", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSubscriberprofileADCValue("10.222.74.186", "servicepath", ""),
				),
			},
		},
	})
}

// testAccCheckSubscriberprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckSubscriberprofileADCValue(ip, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Subscriberprofile.Type(), ip)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("subscriberprofile %s not found on appliance", ip)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("subscriberprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", ip, attr, got, want)
		}
		return nil
	}
}

func TestAccSubscriberprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriberprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_subscriberprofile.tf_subscriberprofile", "ip", "10.222.74.185"),
					resource.TestCheckResourceAttr("data.citrixadc_subscriberprofile.tf_subscriberprofile", "vlan", "1"),
				),
			},
		},
	})
}
