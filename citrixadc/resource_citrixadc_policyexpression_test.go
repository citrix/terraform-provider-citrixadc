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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccPolicyexpression_advanced(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPolicyexpressionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccPolicyexpression_advanced_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyexpressionExist("citrixadc_policyexpression.tf_advanced_policyexpression", nil),
				),
			},
			resource.TestStep{
				Config: testAccPolicyexpression_advanced_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyexpressionExist("citrixadc_policyexpression.tf_advanced_policyexpression", nil),
				),
			},
			resource.TestStep{
				Config: testAccPolicyexpression_advanced_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyexpressionExist("citrixadc_policyexpression.tf_advanced_policyexpression", nil),
				),
			},
		},
	})
}

func TestAccPolicyexpression_classic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPolicyexpressionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccPolicyexpression_classic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyexpressionExist("citrixadc_policyexpression.tf_classic_policyexpression", nil),
				),
			},
			resource.TestStep{
				Config: testAccPolicyexpression_classic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyexpressionExist("citrixadc_policyexpression.tf_classic_policyexpression", nil),
				),
			},
			resource.TestStep{
				Config: testAccPolicyexpression_classic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyexpressionExist("citrixadc_policyexpression.tf_classic_policyexpression", nil),
				),
			},
		},
	})
}

func testAccCheckPolicyexpressionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Policyexpression.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckPolicyexpressionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policyexpression" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Policyexpression.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccPolicyexpression_advanced_step1 = `

resource "citrixadc_policyexpression" "tf_advanced_policyexpression" {
    name = "tf_advanced_policyexrpession"
    value = "HTTP.REQ.URL.SUFFIX.EQ(\"cgi\")"
    comment = "comment"
}
`

const testAccPolicyexpression_advanced_step2 = `

resource "citrixadc_policyexpression" "tf_advanced_policyexpression" {
    name = "tf_advanced_policyexrpession"
    value = "HTTP.REQ.URL.SUFFIX.EQ(\"cginew\")"
    comment = "comment"
}
`

const testAccPolicyexpression_advanced_step3 = `

resource "citrixadc_policyexpression" "tf_advanced_policyexpression" {
    name = "tf_advanced_policyexrpession_new"
    value = "HTTP.REQ.URL.SUFFIX.EQ(\"cginew\")"
    comment = "comment"
}
`

const testAccPolicyexpression_classic_step1 = `

resource "citrixadc_policyexpression" "tf_classic_policyexpression" {
    name = "tf_classic_policyexrpession"
    value = "HEADER Cookie EXISTS"
    clientsecuritymessage = "security message"
}
`

const testAccPolicyexpression_classic_step2 = `

resource "citrixadc_policyexpression" "tf_classic_policyexpression" {
    name = "tf_classic_policyexrpession"
    value = "METHOD != GET"
    clientsecuritymessage = "security message"
}
`

const testAccPolicyexpression_classic_step3 = `

resource "citrixadc_policyexpression" "tf_classic_policyexpression" {
    name = "tf_classic_policyexrpession_new"
    value = "METHOD != GET"
    clientsecuritymessage = "new security message"
}
`
