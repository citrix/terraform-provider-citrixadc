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
	"github.com/chiradeep/go-nitro/config/basic"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

func TestAccService_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccService_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.foo", nil),
					resource.TestCheckResourceAttr(
						"citrixadc_service.foo", "lbvserver", "foo_lb"), resource.TestCheckResourceAttr(
						"citrixadc_service.foo", "name", "foo_svc"), resource.TestCheckResourceAttr(
						"citrixadc_service.foo", "port", "80"), resource.TestCheckResourceAttr(
						"citrixadc_service.foo", "servername", "10.202.22.12"), resource.TestCheckResourceAttr(
						"citrixadc_service.foo", "servicetype", "HTTP"),
				),
			},
		},
	})
}

func testAccCheckServiceExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(netscaler.Service.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckServiceDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_service" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Service.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccService_basic = `

resource "citrixadc_lbvserver" "foo" {

  ipv46 = "10.202.11.11"
  name = "foo_lb"
  port = 80
  servicetype = "HTTP"
}


resource "citrixadc_service" "foo" {

  lbvserver = "foo_lb"
  name = "foo_svc"
  port = 80
  ip = "10.202.22.12"
  servicetype = "HTTP"

  depends_on = ["citrixadc_lbvserver.foo"]

}
`

func TestAccService_AssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	serviceName := "tf-acc-service-test"
	serviceType := netscaler.Service.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, serviceType, serviceName, nil)

	serviceInstance := basic.Service{
		Name:        serviceName,
		Port:        80,
		Ip:          "10.202.22.12",
		Servicetype: "HTTP",
	}

	if _, err := c.client.AddResource(serviceType, serviceName, serviceInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// Zero out immutable members
	serviceInstance.Port = 0
	serviceInstance.Ip = ""
	serviceInstance.Servicetype = ""

	//ip
	serviceInstance.Ip = "1.1.1.1"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "ip")
	serviceInstance.Ip = ""

	//servername
	serviceInstance.Servername = "server1"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "servername")
	serviceInstance.Servername = ""

	//servicetype
	serviceInstance.Servicetype = "HTTP"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "servicetype")
	serviceInstance.Servicetype = ""

	//port
	serviceInstance.Port = 88
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "port")
	serviceInstance.Port = 0

	//cleartextport
	serviceInstance.Cleartextport = 98
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "cleartextport")
	serviceInstance.Cleartextport = 0

	//cachetype
	serviceInstance.Cachetype = "TRANSPARENT"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "cachetype")
	serviceInstance.Cachetype = ""

	//td
	serviceInstance.Td = 2
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "td")
	serviceInstance.Td = 0

	//riseapbrstatsmsgcode
	serviceInstance.Riseapbrstatsmsgcode = 2
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "riseapbrstatsmsgcode")
	serviceInstance.Riseapbrstatsmsgcode = 0

}

const testAccServiceEnableDisable_enabled = `

resource "citrixadc_lbvserver" "tf_acc_lbvsrv" {

  ipv46 = "10.202.11.11"
  name = "tf_acc_lbvsrv"
  port = 80
  servicetype = "HTTP"
}


resource "citrixadc_service" "tf_acc_service" {

  lbvserver = citrixadc_lbvserver.tf_acc_lbvsrv.name
  name = "tf_acc_service"
  port = 80
  ip = "10.202.22.12"
  servicetype = "HTTP"
  comment = "enabled state comment"

  state = "ENABLED"
  graceful = "YES"
  delay = 60
}
`

const testAccServiceEnableDisable_disabled = `

resource "citrixadc_lbvserver" "tf_acc_lbvsrv" {

  ipv46 = "10.202.11.11"
  name = "tf_acc_lbvsrv"
  port = 80
  servicetype = "HTTP"
}


resource "citrixadc_service" "tf_acc_service" {

  lbvserver = citrixadc_lbvserver.tf_acc_lbvsrv.name
  name = "tf_acc_service"
  port = 80
  ip = "10.202.22.12"
  servicetype = "HTTP"
  comment = "disabled state comment"

  state = "DISABLED"
  graceful = "YES"
  delay = 60
}
`

func TestAccService_enable_disable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccServiceEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.tf_acc_service", nil),
					resource.TestCheckResourceAttr("citrixadc_service.tf_acc_service", "state", "ENABLED"),
				),
			},
			resource.TestStep{
				Config: testAccServiceEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.tf_acc_service", nil),
					resource.TestCheckResourceAttr("citrixadc_service.tf_acc_service", "state", "DISABLED"),
				),
			},
			resource.TestStep{
				Config: testAccServiceEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.tf_acc_service", nil),
					resource.TestCheckResourceAttr("citrixadc_service.tf_acc_service", "state", "ENABLED"),
				),
			},
		},
	})
}

func TestAccService_sslservice(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccService_sslservice_config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.test_service", nil),
					resource.TestCheckResourceAttr("citrixadc_service.test_service", "snienable", "ENABLED"),
				),
			},
		},
	})
}

const testAccService_sslservice_config = `

resource "citrixadc_lbvserver" "test_lbvserver" {
    name = "test_lbvserver"
    ipv46 = "10.33.55.33"
    port = 80

}

resource "citrixadc_service" "test_service" {
    servicetype = "SSL"
    name = "test_service"
    ipaddress = "10.77.33.22"
    ip = "10.77.33.22"
    port = "443"
    lbvserver = citrixadc_lbvserver.test_lbvserver.name
    snienable = "ENABLED"
	commonname = "test.com"
}
`
