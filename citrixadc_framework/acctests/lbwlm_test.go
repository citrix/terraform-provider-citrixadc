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

// lbwlm/WLM is a deprecated NetScaler feature but still functional.
// wlmname/lbuid/ipaddress/port are create-only (RequiresReplace); only katimeout
// is updateable in place, so step2 changes only katimeout (2 -> 10) to exercise
// the set/update path without forcing a replacement.

const testAccLbwlm_basic_step1 = `
resource "citrixadc_lbwlm" "tf_lbwlm" {
  wlmname   = "tf_lbwlm"
  lbuid     = "TF_LBWLM_UID"
  ipaddress = "10.222.74.128"
  port      = 3060
  katimeout = 2
}

`

const testAccLbwlm_basic_step2 = `
resource "citrixadc_lbwlm" "tf_lbwlm" {
  wlmname   = "tf_lbwlm"
  lbuid     = "TF_LBWLM_UID"
  ipaddress = "10.222.74.128"
  port      = 3060
  katimeout = 10
}

`

func TestAccLbwlm_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbwlmDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbwlm_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbwlmExist("citrixadc_lbwlm.tf_lbwlm", nil),
					resource.TestCheckResourceAttr("citrixadc_lbwlm.tf_lbwlm", "wlmname", "tf_lbwlm"),
					resource.TestCheckResourceAttr("citrixadc_lbwlm.tf_lbwlm", "lbuid", "TF_LBWLM_UID"),
					resource.TestCheckResourceAttr("citrixadc_lbwlm.tf_lbwlm", "ipaddress", "10.222.74.128"),
					resource.TestCheckResourceAttr("citrixadc_lbwlm.tf_lbwlm", "port", "3060"),
					resource.TestCheckResourceAttr("citrixadc_lbwlm.tf_lbwlm", "katimeout", "2"),
				),
			},
			{
				Config: testAccLbwlm_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbwlmExist("citrixadc_lbwlm.tf_lbwlm", nil),
					resource.TestCheckResourceAttr("citrixadc_lbwlm.tf_lbwlm", "wlmname", "tf_lbwlm"),
					resource.TestCheckResourceAttr("citrixadc_lbwlm.tf_lbwlm", "lbuid", "TF_LBWLM_UID"),
					resource.TestCheckResourceAttr("citrixadc_lbwlm.tf_lbwlm", "ipaddress", "10.222.74.128"),
					resource.TestCheckResourceAttr("citrixadc_lbwlm.tf_lbwlm", "port", "3060"),
					resource.TestCheckResourceAttr("citrixadc_lbwlm.tf_lbwlm", "katimeout", "10"),
				),
			},
		},
	})
}

func testAccCheckLbwlmExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbwlm name is set")
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
		data, err := client.FindResource(service.Lbwlm.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lbwlm %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbwlmDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbwlm" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbwlm.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbwlm %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbwlmDataSource_basic = `

resource "citrixadc_lbwlm" "tf_lbwlm" {
  wlmname   = "tf_lbwlm"
  lbuid     = "TF_LBWLM_UID"
  ipaddress = "10.222.74.128"
  port      = 3060
  katimeout = 2
}

data "citrixadc_lbwlm" "tf_lbwlm_data" {
  wlmname    = citrixadc_lbwlm.tf_lbwlm.wlmname
  depends_on = [citrixadc_lbwlm.tf_lbwlm]
}
`

func TestAccLbwlmDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbwlmDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbwlmDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbwlm.tf_lbwlm_data", "wlmname", "tf_lbwlm"),
					resource.TestCheckResourceAttr("data.citrixadc_lbwlm.tf_lbwlm_data", "lbuid", "TF_LBWLM_UID"),
					resource.TestCheckResourceAttr("data.citrixadc_lbwlm.tf_lbwlm_data", "katimeout", "2"),
				),
			},
		},
	})
}
