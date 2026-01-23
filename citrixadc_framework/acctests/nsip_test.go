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
