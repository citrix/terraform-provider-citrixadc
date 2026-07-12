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

const testAccSslvserver_ecccurve_binding_basic = `

	resource "citrixadc_sslvserver_ecccurve_binding" "tf_sslvserver_ecccurve_binding" {
		ecccurvename = "P_256"
        vservername = citrixadc_lbvserver.tf_sslvserver.name
        
	}

	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name        = "tf_sslvserver"
		servicetype = "SSL"
	}
`

const testAccSslvserver_ecccurve_binding_basic_step2 = `

	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name        = "tf_sslvserver"
		servicetype = "SSL"
	}
`

func TestAccSslvserver_ecccurve_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserver_ecccurve_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_ecccurve_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_ecccurve_bindingExist("citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding", nil),
				),
			},
			{
				Config: testAccSslvserver_ecccurve_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_ecccurve_bindingNotExist("citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding", "tf_sslvserver,P_256"),
				),
			},
		},
	})
}

func testAccCheckSslvserver_ecccurve_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslvserver_ecccurve_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"vservername", "ecccurvename"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["vservername"]
		ecccurvename := idMap["ecccurvename"]

		findParams := service.FindParams{
			ResourceType:             "sslvserver_ecccurve_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ecccurvename
		found := false
		for _, v := range dataArr {
			if v["ecccurvename"].(string) == ecccurvename {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslvserver_ecccurve_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslvserver_ecccurve_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		ecccurvename := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslvserver_ecccurve_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ecccurvename
		found := false
		for _, v := range dataArr {
			if v["ecccurvename"].(string) == ecccurvename {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslvserver_ecccurve_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSslvserver_ecccurve_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslvserver_ecccurve_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslvserver_ecccurve_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslvserver_ecccurve_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSslvserver_ecccurve_bindingDataSource_basic = `

	resource "citrixadc_sslvserver_ecccurve_binding" "tf_sslvserver_ecccurve_binding" {
		ecccurvename = "P_256"
        vservername = citrixadc_lbvserver.tf_sslvserver.name
        
	}

	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name        = "tf_sslvserver"
		servicetype = "SSL"
	}

	data "citrixadc_sslvserver_ecccurve_binding" "tf_sslvserver_ecccurve_binding" {
		ecccurvename = citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding.ecccurvename
		vservername  = citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding.vservername
		depends_on   = [citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding]
	}
`

func TestAccSslvserver_ecccurve_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_ecccurve_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding", "ecccurvename", "P_256"),
					resource.TestCheckResourceAttr("data.citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding", "vservername", "tf_sslvserver"),
				),
			},
		},
	})
}

const testAccSslvserver_ecccurve_binding_upgrade_basic = `

	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name        = "tf_sslvserver"
		servicetype = "SSL"
	}

	resource "citrixadc_sslvserver_ecccurve_binding" "tf_sslvserver_ecccurve_binding" {
		ecccurvename = "P_256"
		vservername  = citrixadc_lbvserver.tf_sslvserver.name
	}
`

// TestAccSslvserver_ecccurve_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release (legacy comma-joined id) is transparently upgraded by
// the current Framework provider. Step 1 creates the binding with citrix/citrixadc
// 2.2.0 (legacy id "vservername,ecccurvename"); step 2 refreshes/plans the same
// config through the current Framework provider, whose Read parses the legacy id and
// recomputes it to the new "ecccurvename:<v>,vservername:<v>" format (SetAttrFromGet).
func TestAccSslvserver_ecccurve_binding_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslvserver_ecccurve_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release, writing the legacy id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccSslvserver_ecccurve_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_ecccurve_bindingExist("citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding", "id", "tf_sslvserver,P_256"),
				),
			},
			{
				// Step 2: refresh/apply the same config through the current Framework
				// provider. Read exercises ParseIdString on the legacy id, then
				// recomputes the id to the new key:value canonical format.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSslvserver_ecccurve_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_ecccurve_bindingExist("citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding", "id", "ecccurvename:P_256,vservername:tf_sslvserver"),
				),
			},
		},
	})
}

func TestAccSslvserver_ecccurve_binding_import(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	const resAddr = "citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserver_ecccurve_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslvserver_ecccurve_binding_basic},
			{Config: testAccSslvserver_ecccurve_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
