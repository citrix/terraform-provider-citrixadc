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

// NOTE (TODO_PLACEHOLDER - prerequisite): The "bot" feature may need to be
// licensed and enabled on the ADC before these tests can run, e.g.:
//   enable ns feature bot
// If bot is not licensed/enabled the parent botprofile creation will fail.
//
// NOTE (TODO_PLACEHOLDER - value): bot_km_expression_value expects a JavaScript
// snippet/file that is inserted when the keyboard-mouse expression evaluates
// true. The value used below is a best-effort placeholder; replace it with a
// value that your ADC accepts if the bind call is rejected.

import (
	"fmt"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccBotprofile_kmdetectionexpr_binding_basic = `

	resource "citrixadc_botprofile" "tf_botprofile" {
		name                     = "tf_botprofile"
		errorurl                 = "http://www.citrix.com"
		trapurl                  = "/http://www.citrix.com"
		comment                  = "tf_botprofile comment"
		bot_enable_white_list    = "ON"
		bot_enable_black_list    = "ON"
		bot_enable_rate_limit    = "ON"
		devicefingerprint        = "ON"
		devicefingerprintaction  = ["LOG", "RESET"]
		bot_enable_ip_reputation = "ON"
		trap                     = "ON"
		trapaction               = ["LOG", "RESET"]
		bot_enable_tps           = "ON"
	}
	resource "citrixadc_botprofile_kmdetectionexpr_binding" "tf_binding" {
		name                     = citrixadc_botprofile.tf_botprofile.name
		kmdetectionexpr          = "true"
		bot_km_expression_name   = "tf_kmname"
		bot_km_expression_value  = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"login\")"
		bot_km_detection_enabled = "ON"
		bot_bind_comment         = "KmTesting"
	}
`

const testAccBotprofile_kmdetectionexpr_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_botprofile" "tf_botprofile" {
		name                     = "tf_botprofile"
		errorurl                 = "http://www.citrix.com"
		trapurl                  = "/http://www.citrix.com"
		comment                  = "tf_botprofile comment"
		bot_enable_white_list    = "ON"
		bot_enable_black_list    = "ON"
		bot_enable_rate_limit    = "ON"
		devicefingerprint        = "ON"
		devicefingerprintaction  = ["LOG", "RESET"]
		bot_enable_ip_reputation = "ON"
		trap                     = "ON"
		trapaction               = ["LOG", "RESET"]
		bot_enable_tps           = "ON"
	}
`

func TestAccBotprofile_kmdetectionexpr_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofile_kmdetectionexpr_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotprofile_kmdetectionexpr_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofile_kmdetectionexpr_bindingExist("citrixadc_botprofile_kmdetectionexpr_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccBotprofile_kmdetectionexpr_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofile_kmdetectionexpr_bindingNotExist("citrixadc_botprofile_kmdetectionexpr_binding.tf_binding", "tf_botprofile,tf_kmname"),
				),
			},
		},
	})
}

func TestAccBotprofile_kmdetectionexpr_binding_import(t *testing.T) {
	const resAddr = "citrixadc_botprofile_kmdetectionexpr_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofile_kmdetectionexpr_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotprofile_kmdetectionexpr_binding_basic,
			},
			{
				Config:                  testAccBotprofile_kmdetectionexpr_binding_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckBotprofile_kmdetectionexpr_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No botprofile_kmdetectionexpr_binding id is set")
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

		// The framework ID is key:UrlEncode(value) comma-separated; the parent
		// (name) and the matched attribute (bot_km_expression_name) are extracted
		// from the resource attributes rather than parsing the composite ID.
		name := rs.Primary.Attributes["name"]
		bot_km_expression_name := rs.Primary.Attributes["bot_km_expression_name"]

		findParams := service.FindParams{
			ResourceType:             "botprofile_kmdetectionexpr_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching bot_km_expression_name
		found := false
		for _, v := range dataArr {
			if v["bot_km_expression_name"].(string) == bot_km_expression_name {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("botprofile_kmdetectionexpr_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckBotprofile_kmdetectionexpr_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		bot_km_expression_name := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "botprofile_kmdetectionexpr_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching bot_km_expression_name
		found := false
		for _, v := range dataArr {
			if v["bot_km_expression_name"].(string) == bot_km_expression_name {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("botprofile_kmdetectionexpr_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckBotprofile_kmdetectionexpr_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_botprofile_kmdetectionexpr_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("botprofile_kmdetectionexpr_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("botprofile_kmdetectionexpr_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccBotprofileKmdetectionexprBindingDataSource_basic = `
	resource "citrixadc_botprofile" "tf_botprofile" {
		name                     = "tf_botprofile"
		errorurl                 = "http://www.citrix.com"
		trapurl                  = "/http://www.citrix.com"
		comment                  = "tf_botprofile comment"
		bot_enable_white_list    = "ON"
		bot_enable_black_list    = "ON"
		bot_enable_rate_limit    = "ON"
		devicefingerprint        = "ON"
		devicefingerprintaction  = ["LOG", "RESET"]
		bot_enable_ip_reputation = "ON"
		trap                     = "ON"
		trapaction               = ["LOG", "RESET"]
		bot_enable_tps           = "ON"
	}
	resource "citrixadc_botprofile_kmdetectionexpr_binding" "tf_binding" {
		name                     = citrixadc_botprofile.tf_botprofile.name
		kmdetectionexpr          = "true"
		bot_km_expression_name   = "tf_kmname"
		bot_km_expression_value  = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"login\")"
		bot_km_detection_enabled = "ON"
		bot_bind_comment         = "KmTesting"
	}

	data "citrixadc_botprofile_kmdetectionexpr_binding" "tf_binding" {
		name                   = citrixadc_botprofile_kmdetectionexpr_binding.tf_binding.name
		bot_km_expression_name = citrixadc_botprofile_kmdetectionexpr_binding.tf_binding.bot_km_expression_name
		kmdetectionexpr        = citrixadc_botprofile_kmdetectionexpr_binding.tf_binding.kmdetectionexpr
	}
`

func TestAccBotprofileKmdetectionexprBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBotprofileKmdetectionexprBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_kmdetectionexpr_binding.tf_binding", "name", "tf_botprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_kmdetectionexpr_binding.tf_binding", "kmdetectionexpr", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_kmdetectionexpr_binding.tf_binding", "bot_km_expression_name", "tf_kmname"),
				),
			},
		},
	})
}
