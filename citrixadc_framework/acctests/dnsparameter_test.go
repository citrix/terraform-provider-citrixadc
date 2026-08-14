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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccDnsparameter_basic_step1 = `

resource "citrixadc_dnsparameter" "tf_dnsparameter" {
  cacheecszeroprefix         = "DISABLED"
  cachehitbypass             = "ENABLED"
  cachenoexpire              = "ENABLED"
  dns64timeout               = 1200
  dnsrootreferral            = "ENABLED"
  dnssec                     = "ENABLED"
  ecsmaxsubnets              = 5
  maxcachesize               = 10
  maxnegativecachesize       = 10
  maxnegcachettl             = 404800
  maxpipeline                = 245
  maxttl                     = 404800
  maxudppacketsize           = 1180
  minttl                     = 2
  namelookuppriority         = "DNS"
  nxdomainratelimitthreshold = 10
  recursion                  = "ENABLED"
  resolutionorder            = "OnlyAAAAQuery"
  retries                    = 2
  splitpktqueryprocessing    = "DROP"
  zonetransfer 			  = "DISABLED"
  resolvermaxtcptimeout	   = 10
  resolvermaxtcpconnections  = 100
  resolvermaxactiveresolutions = 500
  autosavekeyops 			  = "DISABLED"
}


`

const testAccDnsparameter_basic_step2 = `

resource "citrixadc_dnsparameter" "tf_dnsparameter" {
  cacheecszeroprefix         = "ENABLED"
  cachehitbypass             = "DISABLED"
  cachenoexpire              = "DISABLED"
  dns64timeout               = 1000
  dnsrootreferral            = "DISABLED"
  dnssec                     = "ENABLED"
  ecsmaxsubnets              = 0
  maxcachesize               = 0
  maxnegativecachesize       = 0
  maxnegcachettl             = 604800
  maxpipeline                = 255
  maxttl                     = 604800
  maxudppacketsize           = 1280
  minttl                     = 0
  namelookuppriority         = "WINS"
  nxdomainratelimitthreshold = 0
  recursion                  = "DISABLED"
  resolutionorder            = "OnlyAQuery"
  retries                    = 5
  splitpktqueryprocessing    = "ALLOW"
  zonetransfer 			  = "ENABLED"
  resolvermaxtcptimeout	   = 20
  resolvermaxtcpconnections  = 110
  resolvermaxactiveresolutions = 510
  autosavekeyops 			  = "ENABLED"
}


`

func TestAccDnsparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsparameter_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsparameterExist("citrixadc_dnsparameter.tf_dnsparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "cacheecszeroprefix", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "cachehitbypass", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "cachenoexpire", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "dns64timeout", "1200"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "dnsrootreferral", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "dnssec", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "ecsmaxsubnets", "5"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxcachesize", "10"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxnegativecachesize", "10"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxnegcachettl", "404800"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxpipeline", "245"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxttl", "404800"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxudppacketsize", "1180"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "minttl", "2"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "namelookuppriority", "DNS"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "nxdomainratelimitthreshold", "10"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "recursion", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "resolutionorder", "OnlyAAAAQuery"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "retries", "2"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "splitpktqueryprocessing", "DROP"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "zonetransfer", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "resolvermaxtcptimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "resolvermaxtcpconnections", "100"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "resolvermaxactiveresolutions", "500"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "autosavekeyops", "DISABLED"),
				),
			},
			{
				Config: testAccDnsparameter_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsparameterExist("citrixadc_dnsparameter.tf_dnsparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "cacheecszeroprefix", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "cachehitbypass", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "cachenoexpire", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "dns64timeout", "1000"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "dnsrootreferral", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "dnssec", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "ecsmaxsubnets", "0"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxcachesize", "0"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxnegativecachesize", "0"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxnegcachettl", "604800"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxpipeline", "255"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxttl", "604800"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "maxudppacketsize", "1280"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "minttl", "0"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "namelookuppriority", "WINS"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "nxdomainratelimitthreshold", "0"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "recursion", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "resolutionorder", "OnlyAQuery"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "retries", "5"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "splitpktqueryprocessing", "ALLOW"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "zonetransfer", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "resolvermaxtcptimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "resolvermaxtcpconnections", "110"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "resolvermaxactiveresolutions", "510"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_dnsparameter", "autosavekeyops", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckDnsparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsparameter name is set")
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
		data, err := client.FindResource(service.Dnsparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsparameter %s not found", n)
		}

		return nil
	}
}

func TestAccDnsparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccDnsparameter_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsparameterExist("citrixadc_dnsparameter.tf_dnsparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccDnsparameter_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsparameterExist("citrixadc_dnsparameter.tf_dnsparameter", nil)),
			},
		},
	})
}

// TestAccDnsparameter_unset covers the spec-unsettable attributes that have a
// documented NITRO server default. step1 sets them to non-default values; step2
// removes them from config, so the provider must issue a NITRO unset that reverts
// each to its documented default (also asserted directly on the appliance).
const testAccDnsparameter_unset_step1 = `
resource "citrixadc_dnsparameter" "tf_unset" {
  autosavekeyops               = "ENABLED"
  cacheecszeroprefix           = "DISABLED"
  cachehitbypass               = "ENABLED"
  cachenoexpire                = "ENABLED"
  maxcachesize                 = 10
  cacherecords                 = "NO"
  dnsrootreferral              = "ENABLED"
  dnssec                       = "DISABLED"
  ecsmaxsubnets                = 5
  maxnegcachettl               = 404800
  maxttl                       = 404800
  maxudppacketsize             = 1180
  namelookuppriority           = "DNS"
  recursion                    = "ENABLED"
  resolutionorder              = "OnlyAAAAQuery"
  resolvermaxactiveresolutions = 500
  resolvermaxtcpconnections    = 110
  resolvermaxtcptimeout        = 20
  retries                      = 2
  splitpktqueryprocessing      = "DROP"
  zonetransfer                 = "ENABLED"
}
`

const testAccDnsparameter_unset_step2 = `
resource "citrixadc_dnsparameter" "tf_unset" {
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccDnsparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccDnsparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsparameterExist("citrixadc_dnsparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "autosavekeyops", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "cacheecszeroprefix", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "cachehitbypass", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "cachenoexpire", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "cacherecords", "NO"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "dnsrootreferral", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "dnssec", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "ecsmaxsubnets", "5"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "maxnegcachettl", "404800"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "maxttl", "404800"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "maxudppacketsize", "1180"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "namelookuppriority", "DNS"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "recursion", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "resolutionorder", "OnlyAAAAQuery"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "resolvermaxactiveresolutions", "500"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "resolvermaxtcpconnections", "110"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "resolvermaxtcptimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "retries", "2"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "splitpktqueryprocessing", "DROP"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "zonetransfer", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccDnsparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsparameterExist("citrixadc_dnsparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "autosavekeyops", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "cacheecszeroprefix", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "cachehitbypass", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "cachenoexpire", "DISABLED"),
					resource.TestCheckNoResourceAttr("citrixadc_dnsparameter.tf_unset", "cacherecords"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "dnsrootreferral", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "dnssec", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "ecsmaxsubnets", "0"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "maxnegcachettl", "604800"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "maxttl", "604800"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "maxudppacketsize", "1280"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "namelookuppriority", "WINS"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "recursion", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "resolutionorder", "OnlyAQuery"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "resolvermaxactiveresolutions", "0"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "resolvermaxtcpconnections", "1000"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "resolvermaxtcptimeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "retries", "5"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "splitpktqueryprocessing", "ALLOW"),
					resource.TestCheckResourceAttr("citrixadc_dnsparameter.tf_unset", "zonetransfer", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckDnsparameterADCValue("recursion", "DISABLED"),
					testAccCheckDnsparameterADCValue("maxttl", "604800"),
					testAccCheckDnsparameterADCValue("splitpktqueryprocessing", "ALLOW"),
				),
			},
		},
	})
}

// testAccCheckDnsparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckDnsparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Dnsparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("dnsparameter not found on appliance")
		}
		got := fmt.Sprintf("%v", data[attr])
		if got != want {
			return fmt.Errorf("dnsparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccDnsparameter_import(t *testing.T) {
	const resAddr = "citrixadc_dnsparameter.tf_dnsparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccDnsparameter_basic_step1},
			{
				Config:                  testAccDnsparameter_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccDnsparameterDataSource_basic = `
resource "citrixadc_dnsparameter" "tf_dnsparameter" {
  cacheecszeroprefix         = "DISABLED"
  cachehitbypass             = "ENABLED"
  cachenoexpire              = "ENABLED"
  dns64timeout               = 1200
  dnsrootreferral            = "ENABLED"
  dnssec                     = "ENABLED"
  ecsmaxsubnets              = 5
  maxcachesize               = 10
  maxnegativecachesize       = 10
  maxnegcachettl             = 404800
  maxpipeline                = 245
  maxttl                     = 404800
  maxudppacketsize           = 1180
  minttl                     = 2
  namelookuppriority         = "DNS"
  nxdomainratelimitthreshold = 10
  recursion                  = "ENABLED"
  resolutionorder            = "OnlyAAAAQuery"
  retries                    = 2
  splitpktqueryprocessing    = "DROP"
  zonetransfer               = "DISABLED"
  resolvermaxtcptimeout      = 10
  resolvermaxtcpconnections  = 100
  resolvermaxactiveresolutions = 500
  autosavekeyops             = "DISABLED"
}

data "citrixadc_dnsparameter" "tf_dnsparameter_ds" {
  depends_on = [citrixadc_dnsparameter.tf_dnsparameter]
}
`

func TestAccDnsparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "cacheecszeroprefix", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "cachehitbypass", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "cachenoexpire", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "dns64timeout", "1200"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "dnsrootreferral", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "dnssec", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "ecsmaxsubnets", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "maxcachesize", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "maxnegativecachesize", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "maxnegcachettl", "404800"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "maxpipeline", "245"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "maxttl", "404800"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "maxudppacketsize", "1180"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "minttl", "2"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "namelookuppriority", "DNS"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "nxdomainratelimitthreshold", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "recursion", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "resolutionorder", "OnlyAAAAQuery"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "retries", "2"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "splitpktqueryprocessing", "DROP"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "zonetransfer", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "resolvermaxtcptimeout", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "resolvermaxtcpconnections", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "resolvermaxactiveresolutions", "500"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "autosavekeyops", "DISABLED"),
					resource.TestCheckResourceAttrSet("data.citrixadc_dnsparameter.tf_dnsparameter_ds", "id"),
				),
			},
		},
	})
}
