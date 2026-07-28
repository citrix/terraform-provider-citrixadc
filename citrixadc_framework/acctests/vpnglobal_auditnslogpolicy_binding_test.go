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

// vpnglobal is a singleton on the ADC, so there is no parent resource to create.
// The binding's participating entity is the auditnslogpolicy, reused from
// auditnslogpolicy_test.go (name + rule + action, using the built-in
// SETASLEARNNSLOG_ACT nslog action so no auditnslogaction prerequisite is needed).

const testAccVpnglobalAuditnslogpolicyBinding_basic_step1 = `
	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "10.222.74.180"
		loglevel = ["ALERT", "CRITICAL"]
		tcp      = "ALL"
		acl      = "ENABLED"
		protocolviolations = "NONE"
		}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "tf_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_vpnglobal_auditnslogpolicy_binding" "tf_binding" {
		policyname = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		priority   = 90
		secondary  = false

		depends_on = [citrixadc_auditnslogpolicy.tf_auditnslogpolicy]
	}
`

const testAccVpnglobalAuditnslogpolicyBinding_basic_step2 = `
	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "10.222.74.180"
		loglevel = ["ALERT", "CRITICAL"]
		tcp      = "ALL"
		acl      = "ENABLED"
		protocolviolations = "NONE"
		}

	# Keep the participating entity without the actual binding to confirm proper deletion.
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "tf_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}
`

func TestAccVpnglobalAuditnslogpolicyBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobalAuditnslogpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalAuditnslogpolicyBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobalAuditnslogpolicyBindingExist("citrixadc_vpnglobal_auditnslogpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_auditnslogpolicy_binding.tf_binding", "policyname", "tf_auditnslogpolicy"),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_auditnslogpolicy_binding.tf_binding", "priority", "90"),
				),
			},
			{
				Config: testAccVpnglobalAuditnslogpolicyBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobalAuditnslogpolicyBindingNotExist("citrixadc_vpnglobal_auditnslogpolicy_binding.tf_binding", "tf_auditnslogpolicy"),
				),
			},
		},
	})
}

func TestAccVpnglobalAuditnslogpolicyBinding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnglobal_auditnslogpolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobalAuditnslogpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalAuditnslogpolicyBinding_basic_step1,
			},
			{
				Config:            testAccVpnglobalAuditnslogpolicyBinding_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// policyname (identity) and priority (config) round-trip from the GET.
				// gotopriorityexpression/groupextraction are non-recoverable but omitted
				// from config (null==null) so they need no ignore. secondary is the only
				// remaining ignore: verified live that the GET row returns only
				// policyname/priority/policysubtype/stateflag and that stateflag does NOT
				// encode secondary (identical for true/false); it is a write-only bind flag
				// the appliance never echoes and it is not part of the composite ID
				// (ID = policyname only), so it is genuinely non-recoverable on import.
				ImportStateVerifyIgnore: []string{"secondary"},
			},
		},
	})
}

func testAccCheckVpnglobalAuditnslogpolicyBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_auditnslogpolicy_binding id is set")
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

		// ID is a plain value (single unique attr: policyname)
		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_auditnslogpolicy_binding.Type(),
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
			return fmt.Errorf("vpnglobal_auditnslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobalAuditnslogpolicyBindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		policyname := id

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_auditnslogpolicy_binding.Type(),
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
			return fmt.Errorf("vpnglobal_auditnslogpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobalAuditnslogpolicyBindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_auditnslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_auditnslogpolicy_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// A missing-resource error means the binding is gone, which is what we want.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("vpnglobal_auditnslogpolicy_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccVpnglobalAuditnslogpolicyBindingDataSource_basic = `
	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "10.222.74.180"
		loglevel = ["ALERT", "CRITICAL"]
		tcp      = "ALL"
		acl      = "ENABLED"
		protocolviolations = "NONE"
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "tf_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_vpnglobal_auditnslogpolicy_binding" "tf_binding" {
		policyname = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		priority   = 90
		secondary  = false

		depends_on = [citrixadc_auditnslogpolicy.tf_auditnslogpolicy]
	}

	data "citrixadc_vpnglobal_auditnslogpolicy_binding" "tf_binding" {
		policyname = citrixadc_vpnglobal_auditnslogpolicy_binding.tf_binding.policyname
		depends_on = [citrixadc_vpnglobal_auditnslogpolicy_binding.tf_binding]
	}
`

func TestAccVpnglobalAuditnslogpolicyBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalAuditnslogpolicyBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_auditnslogpolicy_binding.tf_binding", "policyname", "tf_auditnslogpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_auditnslogpolicy_binding.tf_binding", "priority", "90"),
				),
			},
		},
	})
}
