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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccVpnvserver_authenticationsamlpolicy_binding_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name           = "tf_vserver_examples"
		servicetype    = "SSL"
		ipv46          = "3.3.3.3"
		port           = 443
	}
	resource "citrixadc_authenticationsamlaction" "tf_samlaction" {
		name                    = "tf_samlaction"
		metadataurl             = "http://www.example.com"
		samltwofactor           = "OFF"
		requestedauthncontext   = "minimum"
		digestmethod            = "SHA1"
		signaturealg            = "RSA-SHA256"
		metadatarefreshinterval = 1
	}
	resource "citrixadc_authenticationsamlpolicy" "tf_samlpolicy" {
		name      = "tf_samlpolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationsamlaction.tf_samlaction.name
	}
	resource "citrixadc_vpnvserver_authenticationsamlpolicy_binding" "tf_bind" {
		name 	  = citrixadc_vpnvserver.tf_vpnvserver.name
		policy    = citrixadc_authenticationsamlpolicy.tf_samlpolicy.name
		priority  = 80
		bindpoint = "RESPONSE"
	}
`

const testAccVpnvserver_authenticationsamlpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name           = "tf_vserver_examples"
		servicetype    = "SSL"
		ipv46          = "3.3.3.3"
		port           = 443
	}
	resource "citrixadc_authenticationsamlaction" "tf_samlaction" {
		name                    = "tf_samlaction"
		metadataurl             = "http://www.example.com"
		samltwofactor           = "OFF"
		requestedauthncontext   = "minimum"
		digestmethod            = "SHA1"
		signaturealg            = "RSA-SHA256"
		metadatarefreshinterval = 1
	}
	resource "citrixadc_authenticationsamlpolicy" "tf_samlpolicy" {
		name      = "tf_samlpolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationsamlaction.tf_samlaction.name
	}
`

func TestAccVpnvserver_authenticationsamlpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_authenticationsamlpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_authenticationsamlpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_authenticationsamlpolicy_bindingExist("citrixadc_vpnvserver_authenticationsamlpolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccVpnvserver_authenticationsamlpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_authenticationsamlpolicy_bindingNotExist("citrixadc_vpnvserver_authenticationsamlpolicy_binding.tf_bind", "tf_vserver_examples,tf_samlpolicy"),
				),
			},
		},
	})
}

func testAccCheckVpnvserver_authenticationsamlpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_authenticationsamlpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "policy"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_authenticationsamlpolicy_binding",
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
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnvserver_authenticationsamlpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_authenticationsamlpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "policy"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		name := idMap["name"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_authenticationsamlpolicy_binding",
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
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnvserver_authenticationsamlpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_authenticationsamlpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_authenticationsamlpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnvserver_authenticationsamlpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver_authenticationsamlpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnvserver_authenticationsamlpolicy_bindingDataSource_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name           = "tf_vserver_examples"
		servicetype    = "SSL"
		ipv46          = "3.3.3.3"
		port           = 443
	}
	resource "citrixadc_authenticationsamlaction" "tf_samlaction" {
		name                    = "tf_samlaction"
		metadataurl             = "http://www.example.com"
		samltwofactor           = "OFF"
		requestedauthncontext   = "minimum"
		digestmethod            = "SHA1"
		signaturealg            = "RSA-SHA256"
		metadatarefreshinterval = 1
	}
	resource "citrixadc_authenticationsamlpolicy" "tf_samlpolicy" {
		name      = "tf_samlpolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationsamlaction.tf_samlaction.name
	}
	resource "citrixadc_vpnvserver_authenticationsamlpolicy_binding" "tf_bind" {
		name 	  = citrixadc_vpnvserver.tf_vpnvserver.name
		policy    = citrixadc_authenticationsamlpolicy.tf_samlpolicy.name
		priority  = 80
		bindpoint = "RESPONSE"
	}

	data "citrixadc_vpnvserver_authenticationsamlpolicy_binding" "tf_bind" {
		name   = citrixadc_vpnvserver_authenticationsamlpolicy_binding.tf_bind.name
		policy = citrixadc_vpnvserver_authenticationsamlpolicy_binding.tf_bind.policy
	}
`

func TestAccVpnvserver_authenticationsamlpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_authenticationsamlpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_authenticationsamlpolicy_binding.tf_bind", "name", "tf_vserver_examples"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_authenticationsamlpolicy_binding.tf_bind", "policy", "tf_samlpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_authenticationsamlpolicy_binding.tf_bind", "priority", "80"),
				),
			},
		},
	})
}

func TestAccVpnvserver_authenticationsamlpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver_authenticationsamlpolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_authenticationsamlpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnvserver_authenticationsamlpolicy_binding_basic},
			{Config: testAccVpnvserver_authenticationsamlpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"bindpoint", "priority"}},
		},
	})
}

// testAccVpnvserver_authenticationsamlpolicy_binding_upgrade_basic reuses the _basic config
// (binding + all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0 schema and
// the current Framework schema because the migration restored the SDK v2 attribute names. The
// terraform resource label (tf_bind) is kept identical to _basic so the Exist/Destroy helpers
// and resource addresses match.
const testAccVpnvserver_authenticationsamlpolicy_binding_upgrade_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name           = "tf_vserver_examples"
		servicetype    = "SSL"
		ipv46          = "3.3.3.3"
		port           = 443
	}
	resource "citrixadc_authenticationsamlaction" "tf_samlaction" {
		name                    = "tf_samlaction"
		metadataurl             = "http://www.example.com"
		samltwofactor           = "OFF"
		requestedauthncontext   = "minimum"
		digestmethod            = "SHA1"
		signaturealg            = "RSA-SHA256"
		metadatarefreshinterval = 1
	}
	resource "citrixadc_authenticationsamlpolicy" "tf_samlpolicy" {
		name      = "tf_samlpolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationsamlaction.tf_samlaction.name
	}
	resource "citrixadc_vpnvserver_authenticationsamlpolicy_binding" "tf_bind" {
		name 	  = citrixadc_vpnvserver.tf_vpnvserver.name
		policy    = citrixadc_authenticationsamlpolicy.tf_samlpolicy.name
		priority  = 80
		bindpoint = "RESPONSE"
	}
`

// TestAccVpnvserver_authenticationsamlpolicy_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release (legacy comma-separated ID) is correctly upgraded when the
// same config is subsequently managed by the current Framework provider. Step 1 creates the
// binding with citrix/citrixadc 2.2.0 (writes the legacy id "tf_vserver_examples,tf_samlpolicy").
// Step 2 refreshes/plans/applies the same config through the Framework provider, exercising
// ParseIdString on the legacy id; because the Framework recomputes the id on Read
// (SetAttrFromGet), the id upgrades to the new "key:value" form
// "name:tf_vserver_examples,policy:tf_samlpolicy".
func TestAccVpnvserver_authenticationsamlpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_vpnvserver_authenticationsamlpolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnvserver_authenticationsamlpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccVpnvserver_authenticationsamlpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_authenticationsamlpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_vserver_examples,tf_samlpolicy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVpnvserver_authenticationsamlpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_authenticationsamlpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "name:tf_vserver_examples,policy:tf_samlpolicy"),
				),
			},
		},
	})
}
