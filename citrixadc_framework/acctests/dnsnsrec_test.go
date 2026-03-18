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
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"fmt"
	"strings"
	"testing"
)

const testAccDnsnsrec_basic_step1 = `

resource "citrixadc_dnsnsrec" "tf_dnsnsrec1" {
    domain = "www.test.com"
    nameserver = "192.168.1.100"
	ttl = 4000
}

resource "citrixadc_dnsnsrec" "tf_dnsnsrec2" {
    domain = "www.test.com"
    nameserver = "192.168.1.99"
	ttl = 4000
}
`

const testAccDnsnsrec_basic_step2 = `

resource "citrixadc_dnsnsrec" "tf_dnsnsrec1" {
    domain = "www.test.com"
    nameserver = "192.168.1.100"
	ttl = 4000
}

resource "citrixadc_dnsnsrec" "tf_dnsnsrec2" {
    domain = "www.test.com"
    nameserver = "192.168.1.98"
	ttl = 4000
}
`

const testAccDnsnsrecDataSource_basic = `

resource "citrixadc_dnsnsrec" "tf_dnsnsrec" {
    domain = "tf-datasource-test-001.local"
    nameserver = "192.168.99.200"
	ttl = 4000
}

data "citrixadc_dnsnsrec" "tf_dnsnsrec_ds" {
	depends_on = [citrixadc_dnsnsrec.tf_dnsnsrec]
	domain = citrixadc_dnsnsrec.tf_dnsnsrec.domain
}
`

func TestAccDnsnsrec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsnsrecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsnsrec_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsnsrecExist("citrixadc_dnsnsrec.tf_dnsnsrec1", nil),
					testAccCheckDnsnsrecExist("citrixadc_dnsnsrec.tf_dnsnsrec2", nil),
				),
			},
			{
				Config: testAccDnsnsrec_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsnsrecExist("citrixadc_dnsnsrec.tf_dnsnsrec1", nil),
					testAccCheckDnsnsrecExist("citrixadc_dnsnsrec.tf_dnsnsrec2", nil),
				),
			},
		},
	})
}

func testAccCheckDnsnsrecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsnsrec name is set")
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

		dnsnsrecId := rs.Primary.ID

		idSlice := strings.SplitN(dnsnsrecId, ",", 2)
		domain := idSlice[0]
		nameserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType: "dnsnsrec",
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		foundIndex := -1
		for i, v := range dataArr {
			if v["domain"] == domain && v["nameserver"] == nameserver {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			return fmt.Errorf("Cannot find dnsnsrec with id %v", dnsnsrecId)
		}

		return nil
	}
}

func testAccCheckDnsnsrecDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsnsrec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnsnsrec.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsnsrec %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccDnsnsrecDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsnsrecDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnsnsrec.tf_dnsnsrec_ds", "domain", "tf-datasource-test-001.local"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsnsrec.tf_dnsnsrec_ds", "nameserver", "192.168.99.200"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsnsrec.tf_dnsnsrec_ds", "ttl", "4000"),
					resource.TestCheckResourceAttrSet("data.citrixadc_dnsnsrec.tf_dnsnsrec_ds", "id"),
				),
			},
		},
	})
}
