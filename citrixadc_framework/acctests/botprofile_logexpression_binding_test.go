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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccBotprofile_logexpression_binding_basic = `

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
	resource "citrixadc_botprofile_logexpression_binding" "tf_binding" {
		name                     = citrixadc_botprofile.tf_botprofile.name
		logexpression            = "true"
		bot_log_expression_name  = "tf_logname"
		bot_log_expression_value = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"ANDROID\")"
		bot_bind_comment         = "LogTesting"
	}
`

const testAccBotprofile_logexpression_binding_basic_step2 = `
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

func TestAccBotprofile_logexpression_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofile_logexpression_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotprofile_logexpression_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofile_logexpression_bindingExist("citrixadc_botprofile_logexpression_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccBotprofile_logexpression_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofile_logexpression_bindingNotExist("citrixadc_botprofile_logexpression_binding.tf_binding", "tf_botprofile,tf_logname"),
				),
			},
		},
	})
}

func testAccCheckBotprofile_logexpression_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No botprofile_logexpression_binding id is set")
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

		bindingId := rs.Primary.ID

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "bot_log_expression_name"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		bot_log_expression_name := idMap["bot_log_expression_name"]

		findParams := service.FindParams{
			ResourceType:             "botprofile_logexpression_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["bot_log_expression_name"].(string) == bot_log_expression_name {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("botprofile_logexpression_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckBotprofile_logexpression_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		bot_log_expression_name := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "botprofile_logexpression_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["bot_log_expression_name"].(string) == bot_log_expression_name {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("botprofile_logexpression_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckBotprofile_logexpression_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_botprofile_logexpression_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("botprofile_logexpression_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("botprofile_logexpression_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccBotprofileLogexpressionBindingDataSource_basic = `
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
	resource "citrixadc_botprofile_logexpression_binding" "tf_binding" {
		name                     = citrixadc_botprofile.tf_botprofile.name
		logexpression            = "true"
		bot_log_expression_name  = "tf_logname"
		bot_log_expression_value = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"ANDROID\")"
		bot_bind_comment         = "LogTesting"
	}

	data "citrixadc_botprofile_logexpression_binding" "tf_binding" {
		name                    = citrixadc_botprofile_logexpression_binding.tf_binding.name
		bot_log_expression_name = citrixadc_botprofile_logexpression_binding.tf_binding.bot_log_expression_name
	}
`

func TestAccBotprofileLogexpressionBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBotprofileLogexpressionBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_logexpression_binding.tf_binding", "name", "tf_botprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_logexpression_binding.tf_binding", "logexpression", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_logexpression_binding.tf_binding", "bot_log_expression_name", "tf_logname"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_logexpression_binding.tf_binding", "bot_log_expression_value", "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"ANDROID\")"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_logexpression_binding.tf_binding", "bot_bind_comment", "LogTesting"),
				),
			},
		},
	})
}

const testAccBotprofile_logexpression_binding_upgrade_basic = `
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
	resource "citrixadc_botprofile_logexpression_binding" "tf_binding" {
		name                     = citrixadc_botprofile.tf_botprofile.name
		logexpression            = "true"
		bot_log_expression_name  = "tf_logname"
		bot_log_expression_value = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"ANDROID\")"
		bot_bind_comment         = "LogTesting"
	}
`

func TestAccBotprofile_logexpression_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckBotprofile_logexpression_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy comma-joined id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccBotprofile_logexpression_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofile_logexpression_bindingExist("citrixadc_botprofile_logexpression_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_botprofile_logexpression_binding.tf_binding", "id", "tf_botprofile,tf_logname"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccBotprofile_logexpression_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofile_logexpression_bindingExist("citrixadc_botprofile_logexpression_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_botprofile_logexpression_binding.tf_binding", "id", "bot_log_expression_name:tf_logname,logexpression:true,name:tf_botprofile"),
				),
			},
		},
	})
}

func TestAccBotprofile_logexpression_binding_import(t *testing.T) {
	const resAddr = "citrixadc_botprofile_logexpression_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofile_logexpression_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccBotprofile_logexpression_binding_basic},
			{Config: testAccBotprofile_logexpression_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
