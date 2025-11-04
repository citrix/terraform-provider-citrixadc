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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
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
		client, err := testAccGetClient()
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
