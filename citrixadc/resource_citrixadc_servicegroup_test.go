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

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccServicegroup_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServicegroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccServicegroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExist("citrixadc_servicegroup.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_servicegroup.foo", "servicegroupname", "test_servicegroup"),
					resource.TestCheckResourceAttr(
						"citrixadc_servicegroup.foo", "servicetype", "HTTP"),
				),
			},
		},
	})
}

func testAccCheckServicegroupExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(service.Servicegroup.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckServicegroupDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_servicegroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Servicegroup.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// TODO add testcase when we have servicegroupmembers_by_servername defined
const testAccServicegroup_basic = `

resource "citrixadc_lbvserver" "foo1" {

  name = "foo_lb_1"
  ipv46 = "10.202.11.11"
  port = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "foo2" {

  name = "foo_lb_2"
  ipv46 = "10.202.11.12"
  port = 80
  servicetype = "HTTP"
}


resource "citrixadc_servicegroup" "foo" {

  servicegroupname = "test_servicegroup"
  servicetype = "HTTP"
  servicegroupmembers = ["172.20.0.9:80:10", "172.20.0.10:80:10", "172.20.0.11:8080:20"]
  lbvservers = ["foo_lb_1", "foo_lb_2"]
  depends_on = ["citrixadc_lbvserver.foo1", "citrixadc_lbvserver.foo2"]
}
`

func TestAccServicegroup_AssertNonUpdateableAttributes(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	servicegroupName := "tf-acc-servicegroup-test"
	servicegroupType := service.Servicegroup.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, servicegroupType, servicegroupName, nil)

	servicegroupInstance := basic.Servicegroup{
		Servicegroupname: servicegroupName,
		Servicetype:      "HTTP",
	}

	if _, err := c.client.AddResource(servicegroupType, servicegroupName, servicegroupInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	//servicetype
	servicegroupInstance.Servicetype = "HTTP"
	testHelperVerifyImmutabilityFunc(c, t, servicegroupType, servicegroupName, servicegroupInstance, "servicetype")
	servicegroupInstance.Servicetype = ""

	//cachetype
	servicegroupInstance.Cachetype = "TRANSPARENT"
	testHelperVerifyImmutabilityFunc(c, t, servicegroupType, servicegroupName, servicegroupInstance, "cachetype")
	servicegroupInstance.Cachetype = ""

	//td
	servicegroupInstance.Td = 2
	testHelperVerifyImmutabilityFunc(c, t, servicegroupType, servicegroupName, servicegroupInstance, "td")
	servicegroupInstance.Td = 0

	//memberport
	servicegroupInstance.Memberport = 80
	testHelperVerifyImmutabilityFunc(c, t, servicegroupType, servicegroupName, servicegroupInstance, "memberport")
	servicegroupInstance.Memberport = 0

	//includemembers
	servicegroupInstance.Includemembers = true
	testHelperVerifyImmutabilityFunc(c, t, servicegroupType, servicegroupName, servicegroupInstance, "includemembers")
	servicegroupInstance.Includemembers = false

}

const testAccServicegroupEnableDisable_enabled = `
resource "citrixadc_servicegroup" "tf_enable_disable_test_svcgroup" {
	servicegroupname = "tf_enable_disable_test_svcgroup"
    servicetype = "HTTP"
	servicegroupmembers = []
	comment = "enabled state comment"
	state = "ENABLED"
	graceful = "YES"
	delay = 60
}
`

const testAccServicegroupEnableDisable_disabled = `
resource "citrixadc_servicegroup" "tf_enable_disable_test_svcgroup" {
	servicegroupname = "tf_enable_disable_test_svcgroup"
    servicetype = "HTTP"
	servicegroupmembers = []
	comment = "disabled state comment"
	state = "DISABLED"
	graceful = "YES"
	delay = 60
}
`

func TestAccServicegroup_enable_disable(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServicegroupDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			resource.TestStep{
				Config: testAccServicegroupEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExist("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", "state", "ENABLED"),
				),
			},
			// Disable
			resource.TestStep{
				Config: testAccServicegroupEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExist("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", "state", "DISABLED"),
				),
			},
			// Re enable
			resource.TestStep{
				Config: testAccServicegroupEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExist("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", "state", "ENABLED"),
				),
			},
		},
	})
}
