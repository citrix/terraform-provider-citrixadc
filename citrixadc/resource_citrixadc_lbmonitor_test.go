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
	"os"
	"testing"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccLbmonitor_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_lbmonitor.foo", "monitorname", "sample_lb_monitor"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbmonitor.foo", "type", "HTTP"),
				),
			},
		},
	})
}

func testAccCheckLbmonitorExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(service.Lbmonitor.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbmonitorDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbmonitor" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Lbmonitor.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbmonitor_basic = `


resource "citrixadc_lbmonitor" "foo" {
  monitorname = "sample_lb_monitor"
  type = "HTTP"
}
`

func TestAccLbmonitor_AssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	monitorName := "tf-acc-lbmonitor-test"
	monitorType := service.Lbmonitor.Type()

	// Defer deletion of actual resource
	deleteArgsMap := make(map[string]string)
	deleteArgsMap["type"] = "HTTP"
	defer testHelperEnsureResourceDeletion(c, t, monitorType, monitorName, deleteArgsMap)

	monitorInstance := lb.Lbmonitor{
		Monitorname: monitorName,
		Type:        "HTTP",
	}

	if _, err := c.client.AddResource(monitorType, monitorName, monitorInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	//servicename
	monitorInstance.Servicename = "foo"
	testHelperVerifyImmutabilityFunc(c, t, monitorType, monitorName, monitorInstance, "servicename")
	monitorInstance.Servicename = ""

	//servicegroupname
	monitorInstance.Servicegroupname = "foo"
	testHelperVerifyImmutabilityFunc(c, t, monitorType, monitorName, monitorInstance, "servicegroupname")
	monitorInstance.Servicegroupname = ""
}
