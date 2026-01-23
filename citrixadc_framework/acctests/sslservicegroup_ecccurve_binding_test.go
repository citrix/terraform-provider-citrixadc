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
	"strings"
	"testing"
)

const testAccSslservicegroup_ecccurve_binding_basic = `
	resource "citrixadc_sslservicegroup_ecccurve_binding" "tf_sslservicegroup_ecccurve_binding" {
		ecccurvename = "P_256"
        servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
	}

	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype = "SSL"
	}
`

const testAccSslservicegroup_ecccurve_binding_basic_step2 = `
	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype = "SSL"
	}
`

func TestAccSslservicegroup_ecccurve_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslservicegroup_ecccurve_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroup_ecccurve_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_ecccurve_bindingExist("citrixadc_sslservicegroup_ecccurve_binding.tf_sslservicegroup_ecccurve_binding", nil),
				),
			},
			{
				Config: testAccSslservicegroup_ecccurve_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_ecccurve_bindingNotExist("citrixadc_sslservicegroup_ecccurve_binding.tf_sslservicegroup_ecccurve_binding", "tf_sslservicegroup,P_256"),
				),
			},
		},
	})
}

func testAccCheckSslservicegroup_ecccurve_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservicegroup_ecccurve_binding id is set")
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

		idSlice := strings.SplitN(bindingId, ",", 2)

		servicegroupname := idSlice[0]
		ecccurvename := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslservicegroup_ecccurve_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["ecccurvename"].(string) == ecccurvename {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslservicegroup_ecccurve_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslservicegroup_ecccurve_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		sslservicegroupName := idSlice[0]
		ecccurvename := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslservicegroup_ecccurve_binding",
			ResourceName:             sslservicegroupName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right ecccurve name
		found := false
		for _, v := range dataArr {
			if v["ecccurvename"].(string) == ecccurvename {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslservicegroup_ecccurve_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSslservicegroup_ecccurve_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservicegroup_ecccurve_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslservicegroup_ecccurve_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslservicegroup_ecccurve_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
