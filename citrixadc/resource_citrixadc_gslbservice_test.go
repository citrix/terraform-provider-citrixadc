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

	"github.com/chiradeep/go-nitro/config/basic"
	"github.com/chiradeep/go-nitro/config/gslb"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGslbservice_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGslbserviceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGslbservice_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_gslbservice.foo", "ipaddress", "172.16.1.101"),
					resource.TestCheckResourceAttr(
						"citrixadc_gslbservice.foo", "port", "80"),
					resource.TestCheckResourceAttr(
						"citrixadc_gslbservice.foo", "servicename", "gslb1vservice"),
					resource.TestCheckResourceAttr(
						"citrixadc_gslbservice.foo", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr(
						"citrixadc_gslbservice.foo", "sitename", "Site-GSLB-East-Coast"),
				),
			},
		},
	})
}

func testAccCheckGslbserviceExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(netscaler.Gslbservice.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbserviceDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbservice" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Gslbservice.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccGslbservice_basic = `
resource "citrixadc_gslbsite" "foo" {

	siteipaddress = "172.31.11.20"
	sitename = "Site-GSLB-East-Coast"

  }

resource "citrixadc_gslbservice" "foo" {

  ip = "172.16.1.101"
  port = "80"
  servicename = "gslb1vservice"
  servicetype = "HTTP"
  sitename = "${citrixadc_gslbsite.foo.sitename}"

}
`

func TestAccGslbserviceAssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Requisite resource
	serverName := "tf-acc-server-helper"
	serverAddress := "10.12.32.33"
	serverType := netscaler.Server.Type()

	// Defer deletion of requisite resource
	defer testHelperEnsureResourceDeletion(c, t, serverType, serverName, nil)

	serverInstance := basic.Server{
		Name:      serverName,
		Ipaddress: serverAddress,
	}

	// Requisite resource
	siteName := "tf-acc-gslb-site-name"
	siteIpaddress := "10.122.22.22"
	siteType := netscaler.Gslbsite.Type()

	if _, err := c.client.AddResource(serverType, serverName, serverInstance); err != nil {
		t.Logf("Error while creating requisite resource")
		t.Fatal(err)
	}

	// Defer deletion of requisite resource
	defer testHelperEnsureResourceDeletion(c, t, siteType, siteName, nil)

	siteInstance := gslb.Gslbsite{
		Sitename:      siteName,
		Siteipaddress: siteIpaddress,
	}

	if _, err := c.client.AddResource(siteType, siteName, siteInstance); err != nil {
		t.Logf("Error while creating requisite resource")
		t.Fatal(err)
	}

	// Create resource
	serviceName := "tf-acc-gslb-service-test"
	serviceType := netscaler.Gslbservice.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, serviceType, serviceName, nil)

	serviceInstance := gslb.Gslbservice{
		Servicename: serviceName,
		Sitename:    siteName,
		Servername:  serverName,
		Servicetype: "HTTP",
		Port:        8080,
	}

	if _, err := c.client.AddResource(serviceType, serviceName, serviceInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// Zero out fields in present service instance
	serviceInstance.Servername = ""
	serviceInstance.Servicetype = ""
	serviceInstance.Port = 0
	serviceInstance.Sitename = ""

	//cnameentry
	serviceInstance.Cnameentry = "cname"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "cnameentry")
	serviceInstance.Cnameentry = ""

	//ip
	serviceInstance.Ip = "29.2.2.2"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "ip")
	serviceInstance.Ip = ""

	//servername
	serviceInstance.Servername = "other_server"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "servername")
	serviceInstance.Servername = ""

	//servicetype
	serviceInstance.Servicetype = "TCP"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "servicetype")
	serviceInstance.Servicetype = ""

	//port
	serviceInstance.Port = 9999
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "port")
	serviceInstance.Port = 0

	//sitename
	serviceInstance.Sitename = "other_site_name"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "sitename")
	serviceInstance.Sitename = ""

	//cookietimeout
	serviceInstance.Cookietimeout = 10
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "cookietimeout")
	serviceInstance.Cookietimeout = 0

	//clttimeout
	serviceInstance.Clttimeout = 10
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "clttimeout")
	serviceInstance.Clttimeout = 0

	//svrtimeout
	serviceInstance.Svrtimeout = 10
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "svrtimeout")
	serviceInstance.Svrtimeout = 0
}

const testAccGslbserviceEnableDisable_enabled = `
resource "citrixadc_gslbsite" "tf_test_acc_gslbsite" {
  sitename = "tf_test_acc_gslbsite"
  siteipaddress = "192.168.22.33"
  sessionexchange = "DISABLED"
}

resource "citrixadc_gslbservice" "tf_test_acc_gslbservice" {
  ip = "192.168.11.66"
  port = "80"
  servicename = "tf_test_acc_gslbservice"
  servicetype = "HTTP"
  sitename = "${citrixadc_gslbsite.tf_test_acc_gslbsite.sitename}"
  comment = "enabled state comment"
  state = "ENABLED"
  delay = 60
}
`

const testAccGslbserviceEnableDisable_disabled = `
resource "citrixadc_gslbsite" "tf_test_acc_gslbsite" {
  sitename = "tf_test_acc_gslbsite"
  siteipaddress = "192.168.22.33"
  sessionexchange = "DISABLED"
}

resource "citrixadc_gslbservice" "tf_test_acc_gslbservice" {
  ip = "192.168.11.66"
  port = "80"
  servicename = "tf_test_acc_gslbservice"
  servicetype = "HTTP"
  sitename = "${citrixadc_gslbsite.tf_test_acc_gslbsite.sitename}"
  comment = "disabled state comment"
  state = "DISABLED"
  delay = 60
}
`

func TestAccGsblservice_enable_disable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGslbserviceDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			resource.TestStep{
				Config: testAccGslbserviceEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_test_acc_gslbservice", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_test_acc_gslbservice", "state", "ENABLED"),
				),
			},
			// Disable
			resource.TestStep{
				Config: testAccGslbserviceEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_test_acc_gslbservice", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_test_acc_gslbservice", "state", "DISABLED"),
				),
			},
			// Re enable
			resource.TestStep{
				Config: testAccGslbserviceEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_test_acc_gslbservice", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_test_acc_gslbservice", "state", "ENABLED"),
				),
			},
		},
	})
}
