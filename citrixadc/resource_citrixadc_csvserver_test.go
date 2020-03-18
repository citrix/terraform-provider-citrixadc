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

	"github.com/chiradeep/go-nitro/config/cs"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCsvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCsvserverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCsvserver_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "ipv46", "10.202.11.11"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "name", "terraform-cs"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "port", "8080"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "servicetype", "HTTP"),
				),
			},
		},
	})
}

func TestAccCsvserver_ciphers(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Initial
			resource.TestStep{
				Config: testCiphersConfig(templateCsvserverCiphersConfig, []string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("citrixadc_csvserver.ciphers", []string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				),
			},
			// Transpose
			resource.TestStep{
				Config: testCiphersConfig(templateCsvserverCiphersConfig, []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("citrixadc_csvserver.ciphers", []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}),
				),
			},
			// Empty list
			resource.TestStep{
				Config: testCiphersConfig(templateCsvserverCiphersConfig, nil),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("citrixadc_csvserver.ciphers", []string{}),
				),
			},
		},
	})
}

const templateCsvserverCiphersConfig = `

resource "citrixadc_csvserver" "ciphers" {

  ipv46 = "10.202.11.11"
  name = "tf-acc-ciphers-test"
  port = 443
  servicetype = "SSL"
  %v
}

`

func testAccCheckCsvserverExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(netscaler.Csvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserverDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Csvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCsvserver_basic = `


resource "citrixadc_csvserver" "foo" {

  ipv46 = "10.202.11.11"
  name = "terraform-cs"
  port = 8080
  servicetype = "HTTP"

}
`

func TestAccCsvserverAssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	vserverName := "tf-acc-cs-vserver-name"
	vserverType := netscaler.Csvserver.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, vserverType, vserverName, nil)

	vserverInstance := cs.Csvserver{
		Ipv46:       "192.23.23.23",
		Name:        vserverName,
		Servicetype: "HTTP",
		Port:        80,
	}

	if _, err := c.client.AddResource(vserverType, vserverName, vserverInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// Set to zero values all immutables already defined
	vserverInstance.Port = 0
	vserverInstance.Servicetype = ""

	//port
	vserverInstance.Port = 88
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "port")
	vserverInstance.Port = 0

	//td
	vserverInstance.Td = 1
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "td")
	vserverInstance.Td = 0

	//servicetype
	vserverInstance.Servicetype = "TCP"
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "servicetype")
	vserverInstance.Servicetype = ""

	//targettype
	vserverInstance.Targettype = "GSLB"
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "targettype")
	vserverInstance.Targettype = ""

	//range
	vserverInstance.Range = 1
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "range")
	vserverInstance.Range = 0
}

const testAccCsvserverEnableDisable_enabled = `
resource "citrixadc_csvserver" "tf_test_acc_csvserver" {
  name        = "tf_test_acc_csvserver"
  ipv46       = "192.168.33.22"
  port        = 80
  servicetype = "HTTP"
  comment = "enabled state comment"
  state       = "ENABLED"
}
`
const testAccCsvserverEnableDisable_disabled = `
resource "citrixadc_csvserver" "tf_test_acc_csvserver" {
  name        = "tf_test_acc_csvserver"
  ipv46       = "192.168.33.22"
  port        = 80
  servicetype = "HTTP"
  comment = "disabled state comment"
  state       = "DISABLED"
}
`

func TestAccCsvserver_enable_disable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCsvserverDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			resource.TestStep{
				Config: testAccCsvserverEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_test_acc_csvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_test_acc_csvserver", "state", "ENABLED"),
				),
			},
			// Disable
			resource.TestStep{
				Config: testAccCsvserverEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_test_acc_csvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_test_acc_csvserver", "state", "DISABLED"),
				),
			},
			// Re enable
			resource.TestStep{
				Config: testAccCsvserverEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_test_acc_csvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_test_acc_csvserver", "state", "ENABLED"),
				),
			},
		},
	})
}

const testAccCsvserver_binding_add = `
	resource "citrixadc_csvserver" "testbindingfoo" {
		ipv46       = "10.10.10.22"
		name        = "testAccCsVserver"
		port        = 80
		servicetype = "HTTP"

		lbvserverbinding = citrixadc_lbvserver.test_lbvserver_old.name
	}
	resource "citrixadc_lbvserver" "test_lbvserver_old" {
		ipv46       = "10.10.10.33"
		name        = "testAccLbVserver_old"
		port        = 80
		servicetype = "HTTP"
	}
	resource "citrixadc_lbvserver" "test_lbvserver_new" {
		ipv46       = "10.10.10.44"
		name        = "testAccLbVserver_new"
		port        = 80
		servicetype = "HTTP"
	  }
`
const testAccCsvserver_binding_update = `
	resource "citrixadc_csvserver" "testbindingfoo" {
		ipv46       = "10.10.10.22"
		name        = "testAccCsVserver"
		port        = 80
		servicetype = "HTTP"

		lbvserverbinding = citrixadc_lbvserver.test_lbvserver_new.name
	}
	resource "citrixadc_lbvserver" "test_lbvserver_old" {
		ipv46       = "10.10.10.33"
		name        = "testAccLbVserver_old"
		port        = 80
		servicetype = "HTTP"
	}
	resource "citrixadc_lbvserver" "test_lbvserver_new" {
		ipv46       = "10.10.10.44"
		name        = "testAccLbVserver_new"
		port        = 80
		servicetype = "HTTP"
	  }
`

func TestAccCsvserver_lbvserverbinding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCsvserverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCsvserver_binding_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserverExist("citrixadc_csvserver.testbindingfoo", nil)),
			},
			resource.TestStep{
				Config: testAccCsvserver_binding_update,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserverExist("citrixadc_csvserver.testbindingfoo", nil)),
			},
		},
	})
}