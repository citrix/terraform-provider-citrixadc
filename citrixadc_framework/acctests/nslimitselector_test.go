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

// NOTE: the nslimitselector CLI command is deprecated but still functional.
// step1 creates the selector with a single rule expression; step2 appends a
// second rule expression to exercise the update path (rule is updateable).
const testAccNslimitselector_basic_step1 = `
resource "citrixadc_nslimitselector" "tf_nslimitselector" {
  selectorname = "tf_nslimitselector"
  rule         = ["CLIENT.IP.SRC"]
}

`

const testAccNslimitselector_basic_step2 = `
resource "citrixadc_nslimitselector" "tf_nslimitselector" {
  selectorname = "tf_nslimitselector"
  rule         = ["CLIENT.IP.SRC", "HTTP.REQ.URL"]
}

`

func TestAccNslimitselector_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNslimitselectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNslimitselector_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslimitselectorExist("citrixadc_nslimitselector.tf_nslimitselector", nil),
					resource.TestCheckResourceAttr("citrixadc_nslimitselector.tf_nslimitselector", "selectorname", "tf_nslimitselector"),
					resource.TestCheckResourceAttr("citrixadc_nslimitselector.tf_nslimitselector", "rule.#", "1"),
					resource.TestCheckResourceAttr("citrixadc_nslimitselector.tf_nslimitselector", "rule.0", "CLIENT.IP.SRC"),
				),
			},
			{
				Config: testAccNslimitselector_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslimitselectorExist("citrixadc_nslimitselector.tf_nslimitselector", nil),
					resource.TestCheckResourceAttr("citrixadc_nslimitselector.tf_nslimitselector", "selectorname", "tf_nslimitselector"),
					resource.TestCheckResourceAttr("citrixadc_nslimitselector.tf_nslimitselector", "rule.#", "2"),
					resource.TestCheckResourceAttr("citrixadc_nslimitselector.tf_nslimitselector", "rule.0", "CLIENT.IP.SRC"),
					resource.TestCheckResourceAttr("citrixadc_nslimitselector.tf_nslimitselector", "rule.1", "HTTP.REQ.URL"),
				),
			},
		},
	})
}

func testAccCheckNslimitselectorExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nslimitselector name is set")
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
		// NITRO returns this object under the "streamselector" key (nslimitselector
		// is an alias of streamselector); a typed GET against "nslimitselector"
		// returns a body the client cannot map back. Read via "streamselector".
		data, err := client.FindResource(service.Streamselector.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nslimitselector %s not found", n)
		}

		return nil
	}
}

func testAccCheckNslimitselectorDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nslimitselector" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		// Read via "streamselector" (canonical NITRO key) - see note in Exist check.
		_, err := client.FindResource(service.Streamselector.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nslimitselector %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNslimitselectorDataSource_basic = `

resource "citrixadc_nslimitselector" "tf_nslimitselector" {
  selectorname = "tf_nslimitselector_ds"
  rule         = ["CLIENT.IP.SRC"]
}

data "citrixadc_nslimitselector" "tf_nslimitselector" {
  selectorname = citrixadc_nslimitselector.tf_nslimitselector.selectorname
  depends_on   = [citrixadc_nslimitselector.tf_nslimitselector]
}
`

func TestAccNslimitselectorDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNslimitselectorDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nslimitselector.tf_nslimitselector", "selectorname", "tf_nslimitselector_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nslimitselector.tf_nslimitselector", "rule.#", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_nslimitselector.tf_nslimitselector", "rule.0", "CLIENT.IP.SRC"),
				),
			},
		},
	})
}
