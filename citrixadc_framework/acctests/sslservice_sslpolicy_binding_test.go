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
	"strconv"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// sslservice_sslpolicy_binding binds an SSL policy (policyname) to an SSL service
// (servicename). Participating entities:
//   - citrixadc_service of servicetype SSL (the SSL service is the SSL view of that service)
//   - citrixadc_sslpolicy (+ citrixadc_sslaction; config lifted from sslpolicy_test.go)
//
// priority is set explicitly so the composite ID is deterministic.
//
// BLOCKED (confirmed 2026-07-01 on testbed 10.101.132.152, firmware-level):
// A plain SSL-policy -> SSL-service bind is REJECTED with errorcode 257
// "Operation not permitted". Reproduced both via NITRO and the raw CLI:
//
//	PUT /nitro/v1/config/sslservice_sslpolicy_binding
//	  {"sslservice_sslpolicy_binding":{"servicename":<svc>,"policyname":<pol>,"priority":100}}
//	  -> { "errorcode": 257, "message": "Operation not permitted" }
//	CLI: bind ssl service <svc> -policyName <pol> -priority 100
//	  -> ERROR: Operation not permitted
//
// The SAME sslpolicy binds fine to an sslvserver and globally, so the policy/action
// are valid; only the service bind point is restricted. Per `man bind ssl service`,
// the ONLY permitted policy->service bind is a server-auth OCSP-validation policy
// bound with `-type SERVER_AUTH_VAL` (action must be -ocspCertValidation ...). That
// bind DOES succeed on .152 (verified), but the generated resource schema has NO
// `type` attribute (it is absent from tfdata/sslservice_sslpolicy_binding.json), so
// the resource cannot express the only permitted form. This is a schema/codegen gap
// for FeatureDeveloper, not a fixture/test bug. Until `type` (SERVER_AUTH_VAL) is
// added to the resource+datasource schema/model/payload/SetAttrFromGet/delete-args,
// this test cannot pass and stays generate-only.
//
// NOTE: when `type` is added, the typed GET sslservice_sslpolicy_binding/<svc> DOES
// reflect the binding on this firmware (verified), so the existing Read code needs no
// umbrella fallback (unlike the sslservice_sslcipher_binding sibling).
//
// Composite ID = policyname:<v>,priority:<v>,servicename:<v>. The exist/destroy checks
// read the binding array for the servicename and match on policyname + priority.

// step1: create the binding
const testAccSslserviceSslpolicyBinding_basic_step1 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_sslsvc_policy_lb"
  ipv46       = "10.33.55.36"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
  servicetype = "SSL"
  name        = "tf_sslsvc_policy"
  ipaddress   = "10.77.33.25"
  ip          = "10.77.33.25"
  port        = "443"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslsvc_policy_action"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_sslsvc_policy_pol"
  rule   = "false"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_sslservice_sslpolicy_binding" "tf_binding" {
  servicename = citrixadc_service.tf_service.name
  policyname  = citrixadc_sslpolicy.tf_sslpolicy.name
  priority    = 100

  depends_on = [citrixadc_service.tf_service, citrixadc_sslpolicy.tf_sslpolicy]
}
`

// step2: drop the binding (entities remain)
const testAccSslserviceSslpolicyBinding_basic_step2 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_sslsvc_policy_lb"
  ipv46       = "10.33.55.36"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
  servicetype = "SSL"
  name        = "tf_sslsvc_policy"
  ipaddress   = "10.77.33.25"
  ip          = "10.77.33.25"
  port        = "443"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslsvc_policy_action"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_sslsvc_policy_pol"
  rule   = "false"
  action = citrixadc_sslaction.tf_sslaction.name
}
`

func TestAccSslserviceSslpolicyBinding_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslserviceSslpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslserviceSslpolicyBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslserviceSslpolicyBindingExist("citrixadc_sslservice_sslpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslpolicy_binding.tf_binding", "servicename", "tf_sslsvc_policy"),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslpolicy_binding.tf_binding", "policyname", "tf_sslsvc_policy_pol"),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslpolicy_binding.tf_binding", "priority", "100"),
				),
			},
			{
				// Binding dropped; verify it no longer exists on the ADC.
				Config: testAccSslserviceSslpolicyBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslserviceSslpolicyBindingNotExist("tf_sslsvc_policy", "tf_sslsvc_policy_pol", 100),
				),
			},
		},
	})
}

func TestAccSslserviceSslpolicyBinding_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	const resAddr = "citrixadc_sslservice_sslpolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslserviceSslpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslserviceSslpolicyBinding_basic_step1,
			},
			{
				Config:                  testAccSslserviceSslpolicyBinding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSslserviceSslpolicyBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservice_sslpolicy_binding ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}
			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		servicename := idMap["servicename"]
		policyname := idMap["policyname"]
		priority := idMap["priority"]

		findParams := service.FindParams{
			ResourceType: service.Sslservice_sslpolicy_binding.Type(),
			ResourceName: servicename,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if !sslservicePolicyMatches(v, policyname, priority) {
				continue
			}
			found = true
			break
		}
		if !found {
			return fmt.Errorf("sslservice_sslpolicy_binding %s not found on ADC", rs.Primary.ID)
		}

		return nil
	}
}

// sslservicePolicyMatches reports whether a binding array entry matches the given
// policyname and priority (priority compared numerically).
func sslservicePolicyMatches(v map[string]interface{}, policyname, priority string) bool {
	val, ok := v["policyname"].(string)
	if !ok || val != policyname {
		return false
	}
	if priority == "" {
		return true
	}
	pv, ok := v["priority"]
	if !ok {
		return false
	}
	gotPrio, _ := utils.ConvertToInt64(pv)
	wantPrio, _ := strconv.ParseInt(priority, 10, 64)
	return gotPrio == wantPrio
}

// testAccCheckSslserviceSslpolicyBindingNotExist verifies the binding is gone
// while the parent service still exists (step2 keeps the service).
func testAccCheckSslserviceSslpolicyBindingNotExist(servicename, policyname string, priority int64) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Sslservice_sslpolicy_binding.Type(),
			ResourceName:             servicename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return nil
		}
		for _, v := range dataArr {
			if sslservicePolicyMatches(v, policyname, strconv.FormatInt(priority, 10)) {
				return fmt.Errorf("sslservice_sslpolicy_binding for %s/%s still exists", servicename, policyname)
			}
		}
		return nil
	}
}

func testAccCheckSslserviceSslpolicyBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservice_sslpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		servicename := idMap["servicename"]
		policyname := idMap["policyname"]
		priority := idMap["priority"]

		findParams := service.FindParams{
			ResourceType:             service.Sslservice_sslpolicy_binding.Type(),
			ResourceName:             servicename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			continue
		}
		for _, v := range dataArr {
			if sslservicePolicyMatches(v, policyname, priority) {
				return fmt.Errorf("sslservice_sslpolicy_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccSslserviceSslpolicyBindingDataSource_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_sslsvc_policy_lb"
  ipv46       = "10.33.55.36"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
  servicetype = "SSL"
  name        = "tf_sslsvc_policy"
  ipaddress   = "10.77.33.25"
  ip          = "10.77.33.25"
  port        = "443"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslsvc_policy_action"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_sslsvc_policy_pol"
  rule   = "false"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_sslservice_sslpolicy_binding" "tf_binding" {
  servicename = citrixadc_service.tf_service.name
  policyname  = citrixadc_sslpolicy.tf_sslpolicy.name
  priority    = 100

  depends_on = [citrixadc_service.tf_service, citrixadc_sslpolicy.tf_sslpolicy]
}

data "citrixadc_sslservice_sslpolicy_binding" "tf_binding" {
  servicename = citrixadc_sslservice_sslpolicy_binding.tf_binding.servicename
  policyname  = citrixadc_sslservice_sslpolicy_binding.tf_binding.policyname
  priority    = citrixadc_sslservice_sslpolicy_binding.tf_binding.priority
  depends_on  = [citrixadc_sslservice_sslpolicy_binding.tf_binding]
}
`

func TestAccSslserviceSslpolicyBindingDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslserviceSslpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslserviceSslpolicyBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslservice_sslpolicy_binding.tf_binding", "servicename", "tf_sslsvc_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservice_sslpolicy_binding.tf_binding", "policyname", "tf_sslsvc_policy_pol"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservice_sslpolicy_binding.tf_binding", "priority", "100"),
				),
			},
		},
	})
}
