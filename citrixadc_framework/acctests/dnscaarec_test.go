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

// dnscaarec is add/rm only: CAA records cannot be modified, so every settable
// attribute uses RequiresReplace. A changed attribute therefore forces
// replacement rather than an in-place update.
//
// The composite ID is "domain:<domain>,recordid:<recordid>", where recordid is
// server-assigned (Computed). Delete removes the record via the domain (URL)
// plus recordid as a query arg. If a test run aborts mid-way and leaves a
// dangling record, manually clean it up with:
//   curl -X DELETE 'http://<NS_URL>/nitro/v1/config/dnscaarec/<domain>?args=recordid:<id>'

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccDnscaarec_basic_step1 = `
resource "citrixadc_dnscaarec" "tf_dnscaarec" {
  domain      = "tf-caa.example.com"
  valuestring = "letsencrypt.org"
  tag         = "issue"
  flag        = "NONE"
  ttl         = 3600
}
`

// Step 2 changes an attribute value. Because all attributes are RequiresReplace,
// this forces replacement (destroy + create), not an in-place update.
const testAccDnscaarec_basic_step2 = `
resource "citrixadc_dnscaarec" "tf_dnscaarec" {
  domain      = "tf-caa.example.com"
  valuestring = "sectigo.com"
  tag         = "issue"
  flag        = "CRITICAL"
  ttl         = 7200
}
`

func TestAccDnscaarec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnscaarecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnscaarec_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnscaarecExist("citrixadc_dnscaarec.tf_dnscaarec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnscaarec.tf_dnscaarec", "domain", "tf-caa.example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnscaarec.tf_dnscaarec", "valuestring", "letsencrypt.org"),
					resource.TestCheckResourceAttr("citrixadc_dnscaarec.tf_dnscaarec", "tag", "issue"),
					resource.TestCheckResourceAttr("citrixadc_dnscaarec.tf_dnscaarec", "flag", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_dnscaarec.tf_dnscaarec", "ttl", "3600"),
				),
			},
			{
				// Attribute changes force replacement (RequiresReplace on all attrs).
				Config: testAccDnscaarec_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnscaarecExist("citrixadc_dnscaarec.tf_dnscaarec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnscaarec.tf_dnscaarec", "domain", "tf-caa.example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnscaarec.tf_dnscaarec", "valuestring", "sectigo.com"),
					resource.TestCheckResourceAttr("citrixadc_dnscaarec.tf_dnscaarec", "tag", "issue"),
					resource.TestCheckResourceAttr("citrixadc_dnscaarec.tf_dnscaarec", "flag", "CRITICAL"),
					resource.TestCheckResourceAttr("citrixadc_dnscaarec.tf_dnscaarec", "ttl", "7200"),
				),
			},
		},
	})
}

func TestAccDnscaarec_import(t *testing.T) {
	const resAddr = "citrixadc_dnscaarec.tf_dnscaarec"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnscaarecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnscaarec_basic_step1,
			},
			{
				Config:                  testAccDnscaarec_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckDnscaarecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnscaarec ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// dnscaarec is keyed by domain (URL) plus a server-assigned recordid.
		// Look up the record set for the domain and confirm the specific recordid
		// is present.
		domain := rs.Primary.Attributes["domain"]
		recordid := rs.Primary.Attributes["recordid"]

		findParams := service.FindParams{
			ResourceType:             service.Dnscaarec.Type(),
			ResourceName:             domain,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		if len(dataArr) == 0 {
			return fmt.Errorf("dnscaarec %s not found", n)
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["recordid"]; ok && val != nil {
				if fmt.Sprintf("%v", val) == recordid {
					found = true
					break
				}
			}
		}
		if !found {
			return fmt.Errorf("dnscaarec %s (recordid %s) not found", n, recordid)
		}

		return nil
	}
}

func testAccCheckDnscaarecDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnscaarec" {
			continue
		}

		domain := rs.Primary.Attributes["domain"]
		recordid := rs.Primary.Attributes["recordid"]

		findParams := service.FindParams{
			ResourceType:             service.Dnscaarec.Type(),
			ResourceName:             domain,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Domain record set no longer present - considered destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["recordid"]; ok && val != nil {
				if fmt.Sprintf("%v", val) == recordid {
					return fmt.Errorf("dnscaarec %s (recordid %s) still exists", rs.Primary.ID, recordid)
				}
			}
		}
	}

	return nil
}

const testAccDnscaarecDataSource_basic = `
resource "citrixadc_dnscaarec" "tf_dnscaarec" {
  domain      = "tf-caa-ds.example.com"
  valuestring = "letsencrypt.org"
  tag         = "issue"
  flag        = "NONE"
  ttl         = 3600
}

data "citrixadc_dnscaarec" "tf_dnscaarec" {
  domain     = citrixadc_dnscaarec.tf_dnscaarec.domain
  recordid   = citrixadc_dnscaarec.tf_dnscaarec.recordid
  depends_on = [citrixadc_dnscaarec.tf_dnscaarec]
}
`

func TestAccDnscaarecDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnscaarecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnscaarecDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnscaarec.tf_dnscaarec", "domain", "tf-caa-ds.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_dnscaarec.tf_dnscaarec", "valuestring", "letsencrypt.org"),
					resource.TestCheckResourceAttr("data.citrixadc_dnscaarec.tf_dnscaarec", "tag", "issue"),
					resource.TestCheckResourceAttr("data.citrixadc_dnscaarec.tf_dnscaarec", "ttl", "3600"),
				),
			},
		},
	})
}
