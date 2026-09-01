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

const testAccDnsprofile_add = `


	resource "citrixadc_dnsprofile" "tf_add" {
		
		dnsprofilename      = "tf_profile1"
  		dnsquerylogging     = "DISABLED"
  		dnsanswerseclogging = "DISABLED"
  		dnsextendedlogging  = "DISABLED"
  		dnserrorlogging     = "DISABLED"
  		cacherecords        = "ENABLED"
  		cachenegativeresponses="ENABLED"
  		dropmultiqueryrequest="DISABLED"
  		cacheecsresponses ="DISABLED"
		recursiveresolution = "ENABLED"
		insertecs = "ENABLED"
		replaceecs = "ENABLED"
		maxcacheableecsprefixlength = 16
		maxcacheableecsprefixlength6 = 16
		
	}
`
const testAccDnsprofile_update = `


	resource "citrixadc_dnsprofile" "tf_add" {
		
		dnsprofilename      = "tf_profile1"
  		dnsquerylogging     = "DISABLED"
  		dnsanswerseclogging = "DISABLED"
  		dnsextendedlogging  = "DISABLED"
  		dnserrorlogging     = "DISABLED"
  		cacherecords        = "ENABLED"
  		cachenegativeresponses="ENABLED"
  		dropmultiqueryrequest="ENABLED"
  		cacheecsresponses ="DISABLED"
		recursiveresolution = "DISABLED"
		insertecs = "DISABLED"
		replaceecs = "DISABLED"
		maxcacheableecsprefixlength = 18
		maxcacheableecsprefixlength6 = 18
		
	}
`

const testAccDnsprofileDataSource_basic = `

resource "citrixadc_dnsprofile" "tf_dnsprofile_ds" {
	dnsprofilename      = "tf_profile_ds"
	dnsquerylogging     = "DISABLED"
	dnsanswerseclogging = "DISABLED"
	dnsextendedlogging  = "DISABLED"
	dnserrorlogging     = "DISABLED"
	cacherecords        = "ENABLED"
	cachenegativeresponses="ENABLED"
	dropmultiqueryrequest="DISABLED"
	cacheecsresponses ="DISABLED"
	recursiveresolution = "ENABLED"
	insertecs = "ENABLED"
	replaceecs = "ENABLED"
	maxcacheableecsprefixlength = 16
	maxcacheableecsprefixlength6 = 16
}

data "citrixadc_dnsprofile" "tf_dnsprofile_ds" {
	dnsprofilename = citrixadc_dnsprofile.tf_dnsprofile_ds.dnsprofilename
}
`

func TestAccDnsprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsprofileExist("citrixadc_dnsprofile.tf_add", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsprofilename", "tf_profile1"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsquerylogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsanswerseclogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsextendedlogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnserrorlogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cacherecords", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cachenegativeresponses", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dropmultiqueryrequest", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cacheecsresponses", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "recursiveresolution", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "insertecs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "replaceecs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "maxcacheableecsprefixlength", "16"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "maxcacheableecsprefixlength6", "16"),
				),
			},
			{
				Config: testAccDnsprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsprofileExist("citrixadc_dnsprofile.tf_add", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsprofilename", "tf_profile1"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsquerylogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsanswerseclogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsextendedlogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnserrorlogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cacherecords", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cachenegativeresponses", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dropmultiqueryrequest", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cacheecsresponses", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "recursiveresolution", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "insertecs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "replaceecs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "maxcacheableecsprefixlength", "18"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "maxcacheableecsprefixlength6", "18"),
				),
			},
		},
	})
}

func testAccCheckDnsprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsprofile name is set")
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
		data, err := client.FindResource(service.Dnsprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnsprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccDnsprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dnsprofile.tf_add"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Dnsprofile.Type(), "tf_profile1"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDnsprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccDnsprofile_import(t *testing.T) {
	const resAddr = "citrixadc_dnsprofile.tf_add"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDnsprofile_add},
			{
				Config:                  testAccDnsprofile_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccDnsprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDnsprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccDnsprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsprofileExist("citrixadc_dnsprofile.tf_add", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccDnsprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsprofileExist("citrixadc_dnsprofile.tf_add", nil)),
			},
		},
	})
}

// The dnsprofile unset test exercises every unset-eligible (Optional, mutable,
// non-key) attribute. step1 sets all of them to non-default values; step2
// removes them from config so the provider must unset them, reverting each to
// its documented NITRO default.
const testAccDnsprofile_unset_step1 = `
resource "citrixadc_dnsprofile" "tf_unset" {
  dnsprofilename               = "tf_dnsprofile_unset"
  recursiveresolution          = "ENABLED"
  dnsquerylogging              = "ENABLED"
  dnsanswerseclogging          = "ENABLED"
  dnsextendedlogging           = "ENABLED"
  dnserrorlogging              = "ENABLED"
  cacherecords                 = "DISABLED"
  cachenegativeresponses       = "DISABLED"
  dropmultiqueryrequest        = "ENABLED"
  cacheecsresponses            = "ENABLED"
  insertecs                    = "ENABLED"
  replaceecs                   = "ENABLED"
  maxcacheableecsprefixlength  = 16
  maxcacheableecsprefixlength6 = 64
}
`

const testAccDnsprofile_unset_step2 = `
resource "citrixadc_dnsprofile" "tf_unset" {
  dnsprofilename = "tf_dnsprofile_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccDnsprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccDnsprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsprofileExist("citrixadc_dnsprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "recursiveresolution", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "dnsquerylogging", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "dnsanswerseclogging", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "dnsextendedlogging", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "dnserrorlogging", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "cacherecords", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "cachenegativeresponses", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "dropmultiqueryrequest", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "cacheecsresponses", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "insertecs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "replaceecs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "maxcacheableecsprefixlength", "16"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "maxcacheableecsprefixlength6", "64"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccDnsprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsprofileExist("citrixadc_dnsprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "recursiveresolution", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "dnsquerylogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "dnsanswerseclogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "dnsextendedlogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "dnserrorlogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "cacherecords", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "cachenegativeresponses", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "dropmultiqueryrequest", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "cacheecsresponses", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "insertecs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "replaceecs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "maxcacheableecsprefixlength", "32"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_unset", "maxcacheableecsprefixlength6", "128"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckDnsprofileADCValue("tf_dnsprofile_unset", "recursiveresolution", "DISABLED"),
					testAccCheckDnsprofileADCValue("tf_dnsprofile_unset", "cacherecords", "ENABLED"),
					testAccCheckDnsprofileADCValue("tf_dnsprofile_unset", "maxcacheableecsprefixlength", "32"),
				),
			},
		},
	})
}

// testAccCheckDnsprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckDnsprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Dnsprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("dnsprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("dnsprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccDnsprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					// id is the universal runtime-binding proof.
					resource.TestCheckResourceAttrSet("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "dnsprofilename", "tf_profile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "dnsquerylogging", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "dnsanswerseclogging", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "dnsextendedlogging", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "dnserrorlogging", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "cacherecords", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "cachenegativeresponses", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "dropmultiqueryrequest", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "cacheecsresponses", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "recursiveresolution", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "insertecs", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "replaceecs", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "maxcacheableecsprefixlength", "16"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsprofile.tf_dnsprofile_ds", "maxcacheableecsprefixlength6", "16"),
				),
			},
		},
	})
}
