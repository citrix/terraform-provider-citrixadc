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

const testAccSslvserver_sslciphersuite_binding_basic = `
	resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding" {
		ciphername = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		vservername = citrixadc_lbvserver.tf_sslvserver.name
	}

	resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding2" {
		ciphername = "TLS1.3-CHACHA20-POLY1305-SHA256"
		vservername = citrixadc_lbvserver.tf_sslvserver.name
	}

	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name = "tf_sslvserver"
		servicetype = "SSL"
		ipv46 = "5.5.5.5"
		port = 80
	}
`

const testAccSslvserver_sslciphersuite_binding_basic_step2 = `
	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name = "tf_sslvserver"
		servicetype = "SSL"
		ipv46 = "5.5.5.5"
		port = 80
	}
`

const testAccSslvserver_sslciphersuite_bindingDataSource_basic = `
	resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding" {
		ciphername = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		vservername = citrixadc_lbvserver.tf_sslvserver.name
	}

	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name = "tf_sslvserver_ds"
		servicetype = "SSL"
		ipv46 = "5.5.5.6"
		port = 443
	}

	data "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding" {
		vservername = citrixadc_lbvserver.tf_sslvserver.name
		ciphername = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		depends_on = [citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding]
	}
`

func TestAccSslvserver_sslciphersuite_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserver_sslciphersuite_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_sslciphersuite_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslciphersuite_bindingExist("citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding", nil),
				),
			},
			{
				Config: testAccSslvserver_sslciphersuite_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslciphersuite_bindingExist("citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding2", nil),
				),
			},
			{
				Config: testAccSslvserver_sslciphersuite_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslciphersuite_bindingNotExist("citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding", "tf_sslvserver,TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"),
				),
			},
			{
				Config: testAccSslvserver_sslciphersuite_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslciphersuite_bindingNotExist("citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding2", "tf_sslvserver,TLS1.3-CHACHA20-POLY1305-SHA256"),
				),
			},
		},
	})
}

func TestAccSslvserver_sslciphersuite_binding_import(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	const resAddr = "citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserver_sslciphersuite_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslvserver_sslciphersuite_binding_basic},
			{Config: testAccSslvserver_sslciphersuite_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckSslvserver_sslciphersuite_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslvserver_sslciphersuite_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"vservername", "ciphername"}, nil)
		if err != nil {
			return err
		}
		vservername := idMap["vservername"]
		ciphername := idMap["ciphername"]

		findParams := service.FindParams{
			ResourceType:             "sslvserver_sslciphersuite_binding",
			ResourceName:             vservername,
			ResourceMissingErrorCode: 461,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["ciphername"].(string) == ciphername {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslvserver_sslciphersuite_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslvserver_sslciphersuite_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"vservername", "ciphername"}, nil)
		if err != nil {
			return err
		}
		name := idMap["vservername"]
		ciphername := idMap["ciphername"]

		findParams := service.FindParams{
			ResourceType:             "sslvserver_sslciphersuite_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 461,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right certkey name
		found := false
		for _, v := range dataArr {
			if v["ciphername"].(string) == ciphername {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslvserver_sslciphersuite_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSslvserver_sslciphersuite_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslvserver_sslciphersuite_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"vservername", "ciphername"}, nil)
		if err != nil {
			return err
		}
		vservername := idMap["vservername"]

		_, err = client.FindResource(service.Sslvserver_sslciphersuite_binding.Type(), vservername)
		if err == nil {
			return fmt.Errorf("sslvserver_sslciphersuite_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslvserver_sslciphersuite_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_sslciphersuite_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding", "vservername", "tf_sslvserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding", "ciphername", "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"),
				),
			},
		},
	})
}

// testAccSslvserver_sslciphersuite_binding_upgrade_basic reuses the _basic config values
// (a single binding + its prerequisite lbvserver). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names (ciphername, vservername).
const testAccSslvserver_sslciphersuite_binding_upgrade_basic = `
	resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding" {
		ciphername = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		vservername = citrixadc_lbvserver.tf_sslvserver.name
	}

	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name = "tf_sslvserver"
		servicetype = "SSL"
		ipv46 = "5.5.5.5"
		port = 80
	}
`

// TestAccSslvserver_sslciphersuite_binding_sdkv2StateUpgrade verifies that state written by
// the last SDK v2 release is correctly upgraded when the same config is subsequently managed
// by the current Framework provider. Step 1 creates the binding with citrix/citrixadc 2.2.0
// (writes the legacy comma id "tf_sslvserver,TLS1.2-ECDHE-RSA-AES128-GCM-SHA256" — the SDK v2
// d.SetId(fmt.Sprintf("%s,%s", vservername, ciphername))). Step 2 refreshes/plans/applies the
// same config through the Framework provider, exercising ParseIdString on the legacy id; the
// Framework recomputes the id on Read (SetAttrFromGet), so the canonical new-format id becomes
// "ciphername:TLS1.2-ECDHE-RSA-AES128-GCM-SHA256,vservername:tf_sslvserver".
func TestAccSslvserver_sslciphersuite_binding_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resourceAddr := "citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslvserver_sslciphersuite_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccSslvserver_sslciphersuite_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslciphersuite_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_sslvserver,TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed on Read into the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSslvserver_sslciphersuite_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslciphersuite_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "ciphername:TLS1.2-ECDHE-RSA-AES128-GCM-SHA256,vservername:tf_sslvserver"),
				),
			},
		},
	})
}
