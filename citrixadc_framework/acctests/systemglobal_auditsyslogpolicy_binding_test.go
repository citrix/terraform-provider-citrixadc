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

	"github.com/citrix/adc-nitro-go/resource/config/audit"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NOTE on participating entities:
// The standalone citrixadc_systemglobal_auditsyslogpolicy_binding resource manages
// the system-global binding of an auditsyslog policy. Its participating policy must
// exist first. We deliberately do NOT create the policy via the legacy SDK v2
// `citrixadc_auditsyslogpolicy` resource here, because that resource has an embedded
// `globalbinding` sub-block whose Read pulls in the live
// systemglobal_auditsyslogpolicy_binding and whose Delete unbinds it. That makes the
// SDK v2 policy resource fight the standalone binding resource over the same ADC
// object (perpetual plan diff + double-delete "Policy not bound" errorcode 2097).
// Instead we create the auditsyslogaction + auditsyslogpolicy out-of-band via NITRO
// in PreConfig and reference the policy by its literal name, so Terraform only manages
// the standalone binding. The participating entities are torn down in CheckDestroy.

const tfAuditsyslogpolicyBindingPolicyName = "tf_auditsyslogpolicy_bind"
const tfAuditsyslogpolicyBindingActionName = "tf_syslogaction_bind"

// setupAuditsyslogpolicyBindingParticipants creates the auditsyslogaction +
// auditsyslogpolicy on the appliance (idempotently) so the standalone binding has a
// real policy to bind to.
func setupAuditsyslogpolicyBindingParticipants(t *testing.T) {
	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client: %v", err)
	}

	// Clean up any leftovers from a previous aborted run first.
	teardownAuditsyslogpolicyBindingParticipants(t)

	port := 514
	action := audit.Auditsyslogaction{
		Name:       tfAuditsyslogpolicyBindingActionName,
		Serverip:   "10.78.60.33",
		Serverport: &port,
		Loglevel:   []string{"ERROR", "NOTICE"},
	}
	if _, err := c.client.AddResource(service.Auditsyslogaction.Type(), tfAuditsyslogpolicyBindingActionName, &action); err != nil {
		t.Fatalf("Failed to create participating auditsyslogaction: %v", err)
	}

	policy := audit.Auditsyslogpolicy{
		Name:   tfAuditsyslogpolicyBindingPolicyName,
		Rule:   "ns_true",
		Action: tfAuditsyslogpolicyBindingActionName,
	}
	if _, err := c.client.AddResource(service.Auditsyslogpolicy.Type(), tfAuditsyslogpolicyBindingPolicyName, &policy); err != nil {
		t.Fatalf("Failed to create participating auditsyslogpolicy: %v", err)
	}
}

// teardownAuditsyslogpolicyBindingParticipants removes the participating policy and
// action from the appliance. Safe to call when they do not exist.
func teardownAuditsyslogpolicyBindingParticipants(t *testing.T) {
	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client: %v", err)
	}
	// Best-effort unbind in case a binding was left dangling.
	_ = c.client.DeleteResourceWithArgs(service.Systemglobal_auditsyslogpolicy_binding.Type(), "", []string{"policyname:" + tfAuditsyslogpolicyBindingPolicyName})
	_ = c.client.DeleteResource(service.Auditsyslogpolicy.Type(), tfAuditsyslogpolicyBindingPolicyName)
	_ = c.client.DeleteResource(service.Auditsyslogaction.Type(), tfAuditsyslogpolicyBindingActionName)
}

const testAccSystemglobal_auditsyslogpolicy_binding_basic = `
	resource "citrixadc_systemglobal_auditsyslogpolicy_binding" "tf_systemglobal_auditsyslogpolicy_binding" {
		policyname = "` + tfAuditsyslogpolicyBindingPolicyName + `"
		priority   = 50
	}
`

const testAccSystemglobal_auditsyslogpolicy_binding_basic_step2 = `
	# Drop the binding (config is empty) to verify the standalone resource removes it.
`

func TestAccSystemglobal_auditsyslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t); setupAuditsyslogpolicyBindingParticipants(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			err := testAccCheckSystemglobal_auditsyslogpolicy_bindingDestroy(s)
			teardownAuditsyslogpolicyBindingParticipants(t)
			return err
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSystemglobal_auditsyslogpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_auditsyslogpolicy_bindingExist("citrixadc_systemglobal_auditsyslogpolicy_binding.tf_systemglobal_auditsyslogpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_systemglobal_auditsyslogpolicy_binding.tf_systemglobal_auditsyslogpolicy_binding", "policyname", tfAuditsyslogpolicyBindingPolicyName),
					resource.TestCheckResourceAttr("citrixadc_systemglobal_auditsyslogpolicy_binding.tf_systemglobal_auditsyslogpolicy_binding", "priority", "50"),
				),
			},
			{
				Config: testAccSystemglobal_auditsyslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_auditsyslogpolicy_bindingNotExist("citrixadc_systemglobal_auditsyslogpolicy_binding.tf_systemglobal_auditsyslogpolicy_binding", tfAuditsyslogpolicyBindingPolicyName),
				),
			},
		},
	})
}

func TestAccSystemglobal_auditsyslogpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_systemglobal_auditsyslogpolicy_binding.tf_systemglobal_auditsyslogpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t); setupAuditsyslogpolicyBindingParticipants(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			err := testAccCheckSystemglobal_auditsyslogpolicy_bindingDestroy(s)
			teardownAuditsyslogpolicyBindingParticipants(t)
			return err
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSystemglobal_auditsyslogpolicy_binding_basic,
			},
			{
				Config:                  testAccSystemglobal_auditsyslogpolicy_binding_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSystemglobal_auditsyslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemglobal_auditsyslogpolicy_binding id is set")
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

		// Mirror the resource's keyless aggregate read: fetch the full array and
		// filter client-side by policyname (the single-token ID).
		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Systemglobal_auditsyslogpolicy_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("systemglobal_auditsyslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemglobal_auditsyslogpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		policyname := id
		findParams := service.FindParams{
			ResourceType:             service.Systemglobal_auditsyslogpolicy_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("systemglobal_auditsyslogpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSystemglobal_auditsyslogpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemglobal_auditsyslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		// Keyless aggregate read filtered by policyname (the single-token ID).
		policyname := rs.Primary.ID
		findParams := service.FindParams{
			ResourceType:             service.Systemglobal_auditsyslogpolicy_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("systemglobal_auditsyslogpolicy_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccSystemglobal_auditsyslogpolicy_bindingDataSource_basic = `
	resource "citrixadc_systemglobal_auditsyslogpolicy_binding" "tf_systemglobal_auditsyslogpolicy_binding" {
		policyname = "` + tfAuditsyslogpolicyBindingPolicyName + `"
		priority   = 50
	}

	data "citrixadc_systemglobal_auditsyslogpolicy_binding" "tf_systemglobal_auditsyslogpolicy_binding" {
		policyname = citrixadc_systemglobal_auditsyslogpolicy_binding.tf_systemglobal_auditsyslogpolicy_binding.policyname
	}
`

func TestAccSystemglobal_auditsyslogpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t); setupAuditsyslogpolicyBindingParticipants(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			err := testAccCheckSystemglobal_auditsyslogpolicy_bindingDestroy(s)
			teardownAuditsyslogpolicyBindingParticipants(t)
			return err
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSystemglobal_auditsyslogpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemglobal_auditsyslogpolicy_binding.tf_systemglobal_auditsyslogpolicy_binding", "policyname", tfAuditsyslogpolicyBindingPolicyName),
					resource.TestCheckResourceAttr("data.citrixadc_systemglobal_auditsyslogpolicy_binding.tf_systemglobal_auditsyslogpolicy_binding", "priority", "50"),
				),
			},
		},
	})
}
