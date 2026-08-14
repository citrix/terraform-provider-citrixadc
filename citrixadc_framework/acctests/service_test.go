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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

func TestAccService_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccService_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_service.foo", "lbvserver", "foo_lb"),
					resource.TestCheckResourceAttr("citrixadc_service.foo", "name", "foo_svc"),
					resource.TestCheckResourceAttr("citrixadc_service.foo", "port", "80"),
					resource.TestCheckResourceAttr("citrixadc_service.foo", "servername", "10.202.22.12"),
					resource.TestCheckResourceAttr("citrixadc_service.foo", "servicetype", "HTTP"),
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Service.Type(), rs.Primary.ID)

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_service" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Service.Type(), rs.Primary.ID)
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
	serviceType := service.Service.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, serviceType, serviceName, nil)

	serviceInstance := basic.Service{
		Name:        serviceName,
		Port:        utils.IntPtr(80),
		Ip:          "10.202.22.12",
		Servicetype: "HTTP",
	}

	if _, err := c.client.AddResource(serviceType, serviceName, serviceInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// Zero out immutable members
	serviceInstance.Port = utils.IntPtr(0)
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

	serviceInstance = basic.Service{
		Name: serviceName,
	}

	//port
	serviceInstance.Port = utils.IntPtr(88)
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "port")

	serviceInstance = basic.Service{
		Name: serviceName,
	}

	//cleartextport
	serviceInstance.Cleartextport = utils.IntPtr(98)
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "cleartextport")
	serviceInstance.Cleartextport = utils.IntPtr(0)

	serviceInstance = basic.Service{
		Name: serviceName,
	}

	//cachetype
	serviceInstance.Cachetype = "TRANSPARENT"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "cachetype")
	serviceInstance.Cachetype = ""

	//td
	serviceInstance.Td = utils.IntPtr(2)
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "td")
	serviceInstance.Td = utils.IntPtr(0)

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
  wait_until_disabled = true
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
  wait_until_disabled = true
}
`

func TestAccService_enable_disable(t *testing.T) {
	t.Skip("TODO: Disable")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.tf_acc_service", nil),
					resource.TestCheckResourceAttr("citrixadc_service.tf_acc_service", "state", "ENABLED"),
				),
			},
			{
				Config: testAccServiceEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.tf_acc_service", nil),
					resource.TestCheckResourceAttr("citrixadc_service.tf_acc_service", "state", "DISABLED"),
				),
			},
			{
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
	// if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
	// 	t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	// }
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
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

func TestAccService_rebind_default_monitor(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccService_rebind_default_monitor_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.test_service", nil),
					resource.TestCheckResourceAttr("citrixadc_service.test_service", "lbmonitor", "tcp-default"),
				),
			},
			{
				Config: testAccService_rebind_default_monitor_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.test_service", nil),
					resource.TestCheckResourceAttr("citrixadc_service.test_service", "lbmonitor", "tf_monitor"),
				),
			},
			{
				Config: testAccService_rebind_default_monitor_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.test_service", nil),
					resource.TestCheckResourceAttr("citrixadc_service.test_service", "lbmonitor", "tcp-default"),
				),
			},
		},
	})
}

const testAccService_rebind_default_monitor_step1 = `
resource "citrixadc_lbmonitor" "tf_monitor" {
	monitorname = "tf_monitor"
	type = "HTTP"
}

resource "citrixadc_service" "test_service" {
    servicetype = "HTTP"
    name = "test_service"
    ipaddress = "10.77.33.22"
    ip = "10.77.33.22"
    port = "80"
}
`

const testAccService_rebind_default_monitor_step2 = `
resource "citrixadc_lbmonitor" "tf_monitor" {
	monitorname = "tf_monitor"
	type = "HTTP"
}

resource "citrixadc_service" "test_service" {
    servicetype = "HTTP"
    name = "test_service"
    ipaddress = "10.77.33.22"
    ip = "10.77.33.22"
    port = "80"
    lbmonitor = citrixadc_lbmonitor.tf_monitor.monitorname
}
`

const testAccService_rebind_default_monitor_step3 = `
resource "citrixadc_lbmonitor" "tf_monitor" {
	monitorname = "tf_monitor"
	type = "HTTP"
}

resource "citrixadc_service" "test_service" {
    servicetype = "HTTP"
    name = "test_service"
    ipaddress = "10.77.33.22"
    ip = "10.77.33.22"
    port = "80"
    lbmonitor = "tcp-default"
}
`

const testAccServiceDataSource_basic = `

	resource "citrixadc_server" "tf_server" {
		name      = "test_service_ds_server"
		ipaddress = "192.168.11.15"
	}

	resource "citrixadc_service" "tf_service" {
		name        = "test_service_ds"
		servicetype = "HTTP"
		servername  = citrixadc_server.tf_server.name
		port        = 80
	}

	data "citrixadc_service" "tf_service" {
		name       = citrixadc_service.tf_service.name
		depends_on = [citrixadc_service.tf_service]
	}
`

func TestAccServiceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_service.tf_service", "name", "test_service_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_service.tf_service", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("data.citrixadc_service.tf_service", "port", "80"),
				),
			},
		},
	})
}

func TestAccService_import(t *testing.T) {
	const resAddr = "citrixadc_service.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{Config: testAccService_basic},
			{
				Config:                  testAccService_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ip", "lbmonitor", "lbvserver"},
			},
		},
	})
}

func TestAccService_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_service.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccService_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckServiceExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Service.Type(), "foo_svc"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccService_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckServiceExist(resAddr, nil)),
			},
		},
	})
}

// The service unset test covers the mutable, spec-unsettable attributes that
// have a documented NITRO default which is also a valid SET input value:
// accessdown, appflowlog, cacheable, downstateflush, healthmonitor,
// processlocal. Step 1 sets them to non-default values;
// step 2 removes them from config so the provider must unset them (revert to
// the documented NITRO defaults) and the implicit post-apply plan must be empty.
const testAccService_unset_step1 = `
resource "citrixadc_service" "tf_unset" {
  name               = "tf_test_service_unset"
  servicetype        = "HTTP"
  ip                 = "10.222.111.56"
  port               = 80
  accessdown         = "YES"
  appflowlog         = "DISABLED"
  cacheable          = "YES"
  downstateflush     = "DISABLED"
  healthmonitor      = "NO"
  processlocal       = "ENABLED"
}
`

const testAccService_unset_step2 = `
resource "citrixadc_service" "tf_unset" {
  name        = "tf_test_service_unset"
  servicetype = "HTTP"
  ip          = "10.222.111.56"
  port        = 80
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to the documented NITRO defaults).
}
`

func TestAccService_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccService_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "accessdown", "YES"),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "appflowlog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "cacheable", "YES"),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "downstateflush", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "healthmonitor", "NO"),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "processlocal", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccService_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("citrixadc_service.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "accessdown", "NO"),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "appflowlog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "cacheable", "NO"),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "downstateflush", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "healthmonitor", "YES"),
					resource.TestCheckResourceAttr("citrixadc_service.tf_unset", "processlocal", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckServiceADCValue("tf_test_service_unset", "accessdown", "NO"),
					testAccCheckServiceADCValue("tf_test_service_unset", "healthmonitor", "YES"),
					testAccCheckServiceADCValue("tf_test_service_unset", "appflowlog", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckServiceADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckServiceADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Service.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("service %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("service %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccService_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccService_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckServiceExist("citrixadc_service.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccService_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckServiceExist("citrixadc_service.foo", nil)),
			},
		},
	})
}
