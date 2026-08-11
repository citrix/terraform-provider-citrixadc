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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccNstcpprofile_mpcapablecbit = `

resource "citrixadc_nstcpprofile" "tf_test_profile_mpcapablecbit" {
    name = "test_profile_mpcapablecbit"
    ws = "ENABLED"
	ackaggregation = "ENABLED"
	mpcapablecbit = "ENABLED"
}
`

func TestAccNstcpprofile_mpcapablecbit(t *testing.T) {
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNstcpprofile_mpcapablecbit,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpprofileExist("citrixadc_nstcpprofile.tf_test_profile_mpcapablecbit", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_test_profile_mpcapablecbit", "mpcapablecbit", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckNstcpprofileExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := client.FindResource(service.Nstcpprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckNstcpprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nstcpprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nstcpprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNstcpprofile_basic_step1 = `

resource "citrixadc_nstcpprofile" "tf_test_profile" {
    name = "test_tf_profile"
    ws = "ENABLED"
    ackaggregation = "DISABLED"
	rfc5961compliance = "DISABLED"
}
`

const testAccNstcpprofile_basic_step2 = `

resource "citrixadc_nstcpprofile" "tf_test_profile" {
    name = "test_tf_profile"
    ws = "ENABLED"
    ackaggregation = "ENABLED"
	rfc5961compliance = "ENABLED"
}
`

func TestAccNstcpprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNstcpprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpprofileExist("citrixadc_nstcpprofile.tf_test_profile", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_test_profile", "rfc5961compliance", "DISABLED"),
				),
			},
			{
				Config: testAccNstcpprofile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpprofileExist("citrixadc_nstcpprofile.tf_test_profile", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_test_profile", "rfc5961compliance", "ENABLED"),
				),
			},
		},
	})
}

const testAccNstcpprofileDataSource_basic = `

	resource "citrixadc_nstcpprofile" "tf_nstcpprofile" {
		name = "test_profile_datasource"
		ws = "ENABLED"
		ackaggregation = "ENABLED"
	}

	data "citrixadc_nstcpprofile" "tf_nstcpprofile_data" {
		name = citrixadc_nstcpprofile.tf_nstcpprofile.name
	}
`

func TestAccNstcpprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nstcpprofile.tf_test_profile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNstcpprofile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNstcpprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nstcpprofile.Type(), "test_tf_profile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNstcpprofile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNstcpprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNstcpprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNstcpprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nstcpprofile.tf_nstcpprofile_data", "name", "test_profile_datasource"),
					resource.TestCheckResourceAttr("data.citrixadc_nstcpprofile.tf_nstcpprofile_data", "ws", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nstcpprofile.tf_nstcpprofile_data", "ackaggregation", "ENABLED"),
				),
			},
		},
	})
}

func TestAccNstcpprofile_import(t *testing.T) {
	const resAddr = "citrixadc_nstcpprofile.tf_test_profile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstcpprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNstcpprofile_basic_step1},
			{
				Config:                  testAccNstcpprofile_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNstcpprofile_unset_step1 = `
resource "citrixadc_nstcpprofile" "tf_unset" {
  name              = "tf_test_nstcpprofile_unset"
  ws                = "DISABLED"
  sack              = "DISABLED"
  ackaggregation    = "ENABLED"
  nagle             = "ENABLED"
  ecn               = "ENABLED"
  rfc5961compliance = "ENABLED"
  timestamp         = "ENABLED"
  mptcp             = "ENABLED"
  taillossprobe     = "ENABLED"
  maxburst          = 10
  dupackthresh      = 5
  minrto            = 2000
}
`

const testAccNstcpprofile_unset_step2 = `
resource "citrixadc_nstcpprofile" "tf_unset" {
  name = "tf_test_nstcpprofile_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccNstcpprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNstcpprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpprofileExist("citrixadc_nstcpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "ws", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "sack", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "ackaggregation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "nagle", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "ecn", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "rfc5961compliance", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "timestamp", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "mptcp", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "taillossprobe", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "maxburst", "10"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "dupackthresh", "5"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "minrto", "2000"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNstcpprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpprofileExist("citrixadc_nstcpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "ws", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "sack", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "ackaggregation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "nagle", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "ecn", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "rfc5961compliance", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "timestamp", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "mptcp", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "taillossprobe", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "maxburst", "6"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "dupackthresh", "3"),
					resource.TestCheckResourceAttr("citrixadc_nstcpprofile.tf_unset", "minrto", "1000"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNstcpprofileADCValue("tf_test_nstcpprofile_unset", "ws", "ENABLED"),
					testAccCheckNstcpprofileADCValue("tf_test_nstcpprofile_unset", "ackaggregation", "DISABLED"),
					testAccCheckNstcpprofileADCValue("tf_test_nstcpprofile_unset", "maxburst", "6"),
				),
			},
		},
	})
}

// testAccCheckNstcpprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNstcpprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nstcpprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nstcpprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nstcpprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccNstcpprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNstcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNstcpprofile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNstcpprofileExist("citrixadc_nstcpprofile.tf_test_profile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNstcpprofile_basic_step1,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNstcpprofileExist("citrixadc_nstcpprofile.tf_test_profile", nil)),
			},
		},
	})
}
