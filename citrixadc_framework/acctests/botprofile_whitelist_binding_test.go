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
	"net/url"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccBotprofile_whitelist_binding_basic = `

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
	resource "citrixadc_botprofile_whitelist_binding" "tf_binding" {
		name                  = citrixadc_botprofile.tf_botprofile.name
		bot_whitelist         = "true"
		bot_whitelist_type    = "IPv4"
		bot_whitelist_value   = "1.2.1.2"
		bot_bind_comment      = "TestingWhiteList"
		bot_whitelist_enabled = "ON"
		log                   = "ON"
		logmessage            = "BotWhiteListAdded"
	}
`

const testAccBotprofile_whitelist_binding_basic_step2 = `
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

const testAccBotprofile_whitelist_bindingDataSource_basic = `
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
	resource "citrixadc_botprofile_whitelist_binding" "tf_binding" {
		name                  = citrixadc_botprofile.tf_botprofile.name
		bot_whitelist         = "true"
		bot_whitelist_type    = "IPv4"
		bot_whitelist_value   = "1.2.1.2"
		bot_bind_comment      = "TestingWhiteList"
		bot_whitelist_enabled = "ON"
		log                   = "ON"
		logmessage            = "BotWhiteListAdded"
	}

	data "citrixadc_botprofile_whitelist_binding" "tf_binding" {
		name                = citrixadc_botprofile_whitelist_binding.tf_binding.name
		bot_whitelist_value = citrixadc_botprofile_whitelist_binding.tf_binding.bot_whitelist_value
	}
`

const testAccBotprofile_whitelist_binding_upgrade_basic = `

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
	resource "citrixadc_botprofile_whitelist_binding" "tf_binding" {
		name                  = citrixadc_botprofile.tf_botprofile.name
		bot_whitelist         = "true"
		bot_whitelist_type    = "IPv4"
		bot_whitelist_value   = "1.2.1.2"
		bot_bind_comment      = "TestingWhiteList"
		bot_whitelist_enabled = "ON"
		log                   = "ON"
		logmessage            = "BotWhiteListAdded"
	}
`

// TestAccBotprofile_whitelist_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release (2.2.0), which uses the legacy comma-joined ID
// "name,bot_whitelist_value", is correctly refreshed/planned/applied by the current
// Framework provider. On Read the Framework recomputes the ID to the new key:value
// format, so the ID upgrades from "tf_botprofile,1.2.1.2" to
// "name:tf_botprofile,bot_whitelist_value:1.2.1.2".
func TestAccBotprofile_whitelist_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckBotprofile_whitelist_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release (legacy comma ID).
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.0.0",
					},
				},
				Config: testAccBotprofile_whitelist_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofile_whitelist_bindingExist("citrixadc_botprofile_whitelist_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_botprofile_whitelist_binding.tf_binding", "id", "tf_botprofile,1.2.1.2"),
				),
			},
			{
				// Step 2: refresh/plan/apply the legacy-ID state through the current
				// Framework provider. Read recomputes the ID to the new key:value format.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccBotprofile_whitelist_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofile_whitelist_bindingExist("citrixadc_botprofile_whitelist_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_botprofile_whitelist_binding.tf_binding", "id", "name:tf_botprofile,bot_whitelist_value:1.2.1.2"),
				),
			},
		},
	})
}

func TestAccBotprofile_whitelist_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofile_whitelist_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotprofile_whitelist_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofile_whitelist_bindingExist("citrixadc_botprofile_whitelist_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccBotprofile_whitelist_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofile_whitelist_bindingNotExist("citrixadc_botprofile_whitelist_binding.tf_binding", "tf_botprofile,1.2.1.2"),
				),
			},
		},
	})
}

func TestAccbotprofile_whitelist_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBotprofile_whitelist_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_whitelist_binding.tf_binding", "name", "tf_botprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_whitelist_binding.tf_binding", "bot_whitelist", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_whitelist_binding.tf_binding", "bot_whitelist_value", "1.2.1.2"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_whitelist_binding.tf_binding", "bot_whitelist_type", "IPv4"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_whitelist_binding.tf_binding", "bot_bind_comment", "TestingWhiteList"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_whitelist_binding.tf_binding", "bot_whitelist_enabled", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_whitelist_binding.tf_binding", "log", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile_whitelist_binding.tf_binding", "logmessage", "BotWhiteListAdded"),
				),
			},
		},
	})
}

func testAccCheckBotprofile_whitelist_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No botprofile_whitelist_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "bot_whitelist_value"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		bot_whitelist_value := idMap["bot_whitelist_value"]

		findParams := service.FindParams{
			ResourceType:             "botprofile_whitelist_binding",
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
			if v["bot_whitelist_value"].(string) == bot_whitelist_value {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("botprofile_whitelist_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckBotprofile_whitelist_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"name", "bot_whitelist_value"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		bot_whitelist_value := idMap["bot_whitelist_value"]

		findParams := service.FindParams{
			ResourceType:             "botprofile_whitelist_binding",
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
			if v["bot_whitelist_value"].(string) == bot_whitelist_value {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("botprofile_whitelist_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckBotprofile_whitelist_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_botprofile_whitelist_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("botprofile_whitelist_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("botprofile_whitelist_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccBotprofile_whitelist_binding_import(t *testing.T) {
	const resAddr = "citrixadc_botprofile_whitelist_binding.tf_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,bot_whitelist_value) so it matches exactly what SDK v2 wrote.
	legacyID := func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resAddr]
		if !ok {
			return "", fmt.Errorf("resource not found in state: %s", resAddr)
		}
		kv := map[string]string{}
		for _, p := range strings.Split(rs.Primary.ID, ",") {
			if i := strings.Index(p, ":"); i >= 0 {
				v, _ := url.QueryUnescape(p[i+1:])
				kv[p[:i]] = v
			}
		}
		ordr := []string{"name", "bot_whitelist_value"}
		parts := make([]string, 0, len(ordr))
		for _, k := range ordr {
			if v, ok := kv[k]; ok {
				parts = append(parts, v)
			}
		}
		// Fallback: a positional (non key:value) id has no key:value parts to reorder; import it as-is.
		if len(parts) == 0 {
			return rs.Primary.ID, nil
		}
		return strings.Join(parts, ","), nil
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofile_whitelist_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccBotprofile_whitelist_binding_basic},
			{Config: testAccBotprofile_whitelist_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccBotprofile_whitelist_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

// TestAccBotprofile_whitelist_binding_selfHealing verifies drift recovery:
// after the binding is deleted out-of-band, the next apply of the same config recreates it.
func TestAccBotprofile_whitelist_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_botprofile_whitelist_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofile_whitelist_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotprofile_whitelist_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotprofile_whitelist_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Botprofile_whitelist_binding.Type(), "tf_botprofile", []string{"bot_whitelist:true", "bot_whitelist_value:1.2.1.2"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccBotprofile_whitelist_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotprofile_whitelist_bindingExist(resAddr, nil)),
			},
		},
	})
}
