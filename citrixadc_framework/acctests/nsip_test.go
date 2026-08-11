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
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccNsip_basic_step1 = `

resource "citrixadc_nsip" "tf_test_nsip" {
    ipaddress = "192.168.2.55"
    type = "VIP"
    netmask = "255.255.255.0"
    icmp = "ENABLED"
}
`

const testAccNsip_basic_step2 = `

resource "citrixadc_nsip" "tf_test_nsip" {
    ipaddress = "192.168.2.55"
    type = "VIP"
    netmask = "255.255.254.0"
    icmp = "ENABLED"
}
`

const testAccNsip_basic_step3 = `

resource "citrixadc_nsip" "tf_test_nsip" {
    ipaddress = "192.168.2.55"
    type = "VIP"
    netmask = "255.255.254.0"
    icmp = "DISABLED"
}
`

const testAccNsip_basic_step4 = `

resource "citrixadc_nsip" "tf_test_nsip" {
    ipaddress = "192.168.2.55"
    type = "VIP"
    netmask = "255.255.254.0"
    icmp = "DISABLED"
	state = "DISABLED"
}
`

const testAccNsipDataSource_basic = `
	resource "citrixadc_nsip" "tf_nsip_ds" {
		ipaddress = "10.222.74.149"
		netmask   = "255.255.255.0"
		type      = "SNIP"
		arp       = "ENABLED"
		icmp      = "ENABLED"
		snmp      = "ENABLED"
	}

	data "citrixadc_nsip" "tf_nsip_ds" {
		ipaddress = citrixadc_nsip.tf_nsip_ds.ipaddress
		td        = 0
	}
`

func TestAccNsip_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsip_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip", nil),
				),
			},
			{
				Config: testAccNsip_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip", nil),
				),
			},
			{
				Config: testAccNsip_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip", nil),
				),
			},
			{
				Config: testAccNsip_basic_step4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip", nil),
				),
			},
		},
	})
}

const testAccNsip_mptcpadvertise = `
	resource "citrixadc_nsip" "tf_test_nsip_mptcpadvertise" {
		ipaddress = "192.168.1.55"
		type = "VIP"
		netmask = "255.255.255.0"
		icmp = "ENABLED"
		mptcpadvertise = "YES"
	}
`

func TestAccNsip_mptcpadvertise(t *testing.T) {
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsip_mptcpadvertise,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip_mptcpadvertise", nil),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_test_nsip_mptcpadvertise", "mptcpadvertise", "YES"),
				),
			},
		},
	})
}

const testAccNsip_trafficdomain_create = `

resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
  td        = 2
  aliasname = "tf_trafficdomain"
  vmac      = "ENABLED"
}

resource "citrixadc_nsip" "tf_test_nsip" {
    ipaddress = "192.168.2.155"
    type = "VIP"
    netmask = "255.255.255.0"
    td = citrixadc_nstrafficdomain.tf_trafficdomain.td
}
`

func TestAccNsip_trafficdomain(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsip_trafficdomain_create,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip", nil),
				),
			},
		},
	})
}

func testAccCheckNsipExist(n string, id *string) resource.TestCheckFunc {
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

		argsMap := make(map[string]string)
		nsipName := rs.Primary.ID
		netmask := rs.Primary.Attributes["netmask"]
		trafficDomain := 0
		if val, ok := rs.Primary.Attributes["td"]; ok {
			trafficDomain, _ = strconv.Atoi(val)
		}
		argsMap["td"] = fmt.Sprintf("%d", trafficDomain)
		findParams := service.FindParams{
			ResourceType:             service.Nsip.Type(),
			ResourceName:             nsipName,
			ResourceMissingErrorCode: 258,
			ArgsMap:                  argsMap,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		// Unexpected error
		if err != nil {
			log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
			return fmt.Errorf("Error while finding resource array!")
		}

		// Resource is missing
		if len(dataArr) == 0 {
			log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
			return fmt.Errorf("Error: Resource not found!")
		}

		// Iterate through results to find the one with the right id
		foundIndex := -1
		for i, v := range dataArr {
			if v["ipaddress"].(string) == nsipName && v["netmask"].(string) == netmask {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
			return fmt.Errorf("Error: Resource not found!")
		}

		return nil
	}
}

func testAccCheckNsipDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsip" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		argsMap := make(map[string]string)
		nsipName := rs.Primary.ID
		netmask := rs.Primary.Attributes["netmask"]
		trafficDomain := 0
		if val, ok := rs.Primary.Attributes["td"]; ok {
			trafficDomain, _ = strconv.Atoi(val)
		}
		argsMap["td"] = fmt.Sprintf("%d", trafficDomain)
		findParams := service.FindParams{
			ResourceType:             service.Nsip.Type(),
			ResourceName:             nsipName,
			ResourceMissingErrorCode: 258,
			ArgsMap:                  argsMap,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		// Unexpected error
		if err != nil {
			// If the traffic domain itself is not configured (error 946),
			// the NSIP is implicitly deleted along with the traffic domain
			if strings.Contains(err.Error(), "errorcode\": 946") {
				log.Printf("[DEBUG] citrixadc-provider: Traffic domain not configured, NSIP considered destroyed")
				continue
			}
			log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
			return fmt.Errorf("Error while finding resource array!")
		}

		// Resource is missing
		if len(dataArr) == 0 {
			log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
			return nil
		}

		// Iterate through results to find the one with the right id
		foundIndex := -1
		for i, v := range dataArr {
			if v["ipaddress"].(string) == nsipName && v["netmask"].(string) == netmask {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex != -1 {
			log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams resource still found in array")
			return fmt.Errorf("Error: Resource still found!")
		}

	}

	return nil
}

func TestAccNsip_import(t *testing.T) {
	const resAddr = "citrixadc_nsip.tf_test_nsip"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsipDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNsip_basic_step1},
			{
				Config:                  testAccNsip_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNsip_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsipDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsip_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNsip_basic_step1,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip", nil)),
			},
		},
	})
}

// nsip unset test: step1 sets a set of type-independent, VIP-applicable
// attributes to non-default values; step2 removes them from config so the
// provider must unset them (revert to the documented NITRO defaults).
const testAccNsip_unset_step1 = `
resource "citrixadc_nsip" "tf_unset" {
    ipaddress       = "192.168.3.77"
    netmask         = "255.255.255.0"
    type            = "VIP"
    arp             = "DISABLED"
    mptcpadvertise  = "YES"
    decrementttl    = "ENABLED"
    arpresponse     = "ALL_VSERVERS"
    icmpresponse    = "ALL_VSERVERS"
}
`

const testAccNsip_unset_step2 = `
resource "citrixadc_nsip" "tf_unset" {
    ipaddress = "192.168.3.77"
    netmask   = "255.255.255.0"
    type      = "VIP"
    # All unset-eligible attributes removed from config -> provider unsets them.
}
`

func TestAccNsip_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsipDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNsip_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_unset", "arp", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_unset", "mptcpadvertise", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_unset", "decrementttl", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_unset", "arpresponse", "ALL_VSERVERS"),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_unset", "icmpresponse", "ALL_VSERVERS"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNsip_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_unset", "arp", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_unset", "mptcpadvertise", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_unset", "decrementttl", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_unset", "arpresponse", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_unset", "icmpresponse", "NONE"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsipADCValue("192.168.3.77", 0, "arp", "ENABLED"),
					testAccCheckNsipADCValue("192.168.3.77", 0, "arpresponse", "NONE"),
					testAccCheckNsipADCValue("192.168.3.77", 0, "tag", "0"),
				),
			},
		},
	})
}

// testAccCheckNsipADCValue asserts an attribute's value directly on the
// appliance (not just Terraform state), proving the unset actually reverted it.
func testAccCheckNsipADCValue(ipaddress string, td int, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		argsMap := map[string]string{"td": fmt.Sprintf("%d", td)}
		findParams := service.FindParams{
			ResourceType:             service.Nsip.Type(),
			ResourceName:             ipaddress,
			ResourceMissingErrorCode: 258,
			ArgsMap:                  argsMap,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		foundIndex := -1
		for i, v := range dataArr {
			if addr, ok := v["ipaddress"].(string); ok && addr == ipaddress {
				foundIndex = i
				break
			}
		}
		if foundIndex == -1 {
			return fmt.Errorf("nsip %s not found on appliance", ipaddress)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", dataArr[foundIndex][attr]))
		if got != want {
			return fmt.Errorf("nsip %s: appliance attr %q = %q, want %q (unset did not revert it)", ipaddress, attr, got, want)
		}
		return nil
	}
}

func TestAccNsipDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsipDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsip.tf_nsip_ds", "ipaddress", "10.222.74.149"),
					resource.TestCheckResourceAttr("data.citrixadc_nsip.tf_nsip_ds", "netmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("data.citrixadc_nsip.tf_nsip_ds", "type", "SNIP"),
					resource.TestCheckResourceAttr("data.citrixadc_nsip.tf_nsip_ds", "arp", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nsip.tf_nsip_ds", "icmp", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nsip.tf_nsip_ds", "snmp", "ENABLED"),
				),
			},
		},
	})
}
