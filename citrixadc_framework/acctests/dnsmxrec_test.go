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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"net/url"
	"testing"
)

const testAccDnsmxrec_add = `


resource "citrixadc_dnsmxrec" "dnsmxrec" {
	domain = "example.com"
	mx     = "mail.example.com"
	pref   = 1
	ttl    = 3600
	}
`

const testAccDnsmxrec_update = `


resource "citrixadc_dnsmxrec" "dnsmxrec" {
	domain = "example.com"
	mx     = "mail.example.com"
	pref   = 2
	ttl    = 3601
	}
`

func TestAccDnsmxrec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsmxrecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsmxrec_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsmxrecExist("citrixadc_dnsmxrec.dnsmxrec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsmxrec.dnsmxrec", "domain", "example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnsmxrec.dnsmxrec", "mx", "mail.example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnsmxrec.dnsmxrec", "pref", "1"),
					resource.TestCheckResourceAttr("citrixadc_dnsmxrec.dnsmxrec", "ttl", "3600"),
				),
			},
			{
				Config: testAccDnsmxrec_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsmxrecExist("citrixadc_dnsmxrec.dnsmxrec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsmxrec.dnsmxrec", "domain", "example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnsmxrec.dnsmxrec", "mx", "mail.example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnsmxrec.dnsmxrec", "pref", "2"),
					resource.TestCheckResourceAttr("citrixadc_dnsmxrec.dnsmxrec", "ttl", "3601"),
				),
			},
		},
	})
}

func testAccCheckDnsmxrecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsmxrec name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		dnsmxrecName := rs.Primary.ID
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		argsMap := make(map[string]string)
		argsMap["ecssubnet"] = url.QueryEscape(rs.Primary.Attributes["ecssubnet"])
		argsMap["mx"] = url.QueryEscape(rs.Primary.Attributes["mx"])
		findParams := service.FindParams{
			ResourceType: service.Dnsmxrec.Type(),
			ArgsMap:      argsMap,
		}
		dataArray, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			log.Printf("[WARN] citrix-provider: acceptance test: Clearing lb route state %s", dnsmxrecName)
			return nil
		}
		if len(dataArray) == 0 {
			log.Printf("[WARN] citrix-provider: acceptance test: Dnsmxrec does not exist. Clearing state.")
			return nil
		}

		if len(dataArray) > 1 {
			return fmt.Errorf("[ERROR] citrix-provider: acceptance test: multiple entries found for Dnsmxrec")
		}

		return nil
	}
}

func testAccCheckDnsmxrecDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsmxrec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		argsMap := make(map[string]string)
		argsMap["ecssubnet"] = url.QueryEscape(rs.Primary.Attributes["ecssubnet"])
		argsMap["mx"] = url.QueryEscape(rs.Primary.Attributes["mx"])
		findParams := service.FindParams{
			ResourceType: service.Dnsmxrec.Type(),
			ArgsMap:      argsMap,
		}
		_, err := client.FindResourceArrayWithParams(findParams)

		if err == nil {
			return fmt.Errorf("dnsmxrec %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
