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
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

func TestAccGslbservice_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbserviceDestroy,
		Steps: []resource.TestStep{
			{
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Gslbservice.Type(), rs.Primary.ID)

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbservice" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Gslbservice.Type(), rs.Primary.ID)
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
	sitepassword = "password123"

	}

resource "citrixadc_gslbservice" "foo" {

  ip = "172.16.1.101"
  port = "80"
  servicename = "gslb1vservice"
  servicetype = "HTTP"
  sitename = "${citrixadc_gslbsite.foo.sitename}"

}
`

func TestAccGslbservice_AssertNonUpdateableAttributes(t *testing.T) {
	t.Skip("TODO:")

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
	serverType := service.Server.Type()

	// Defer deletion of requisite resource
	defer testHelperEnsureResourceDeletion(c, t, serverType, serverName, nil)

	serverInstance := basic.Server{
		Name:      serverName,
		Ipaddress: serverAddress,
	}

	// Requisite resource
	siteName := "tf-acc-gslb-site-name"
	siteIpaddress := "10.122.22.22"
	siteType := service.Gslbsite.Type()

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
	serviceType := service.Gslbservice.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, serviceType, serviceName, nil)

	serviceInstance := gslb.Gslbservice{
		Servicename: serviceName,
		Sitename:    siteName,
		Servername:  serverName,
		Servicetype: "HTTP",
		Port:        utils.IntPtr(8080),
	}

	if _, err := c.client.AddResource(serviceType, serviceName, serviceInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// Zero out fields in present service instance
	serviceInstance.Servername = ""
	serviceInstance.Servicetype = ""
	serviceInstance.Port = utils.IntPtr(0)
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
	//serviceInstance.Servername = "other_server"
	//testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "servername")
	//serviceInstance.Servername = ""

	//servicetype
	serviceInstance.Servicetype = "TCP"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "servicetype")
	serviceInstance.Servicetype = ""

	//port
	serviceInstance.Port = utils.IntPtr(9999)
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "port")
	serviceInstance.Port = utils.IntPtr(0)

	//sitename
	serviceInstance.Sitename = "other_site_name"
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "sitename")
	serviceInstance.Sitename = ""

	//cookietimeout
	serviceInstance.Cookietimeout = utils.IntPtr(10)
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "cookietimeout")
	serviceInstance.Cookietimeout = utils.IntPtr(0)

	//clttimeout
	serviceInstance.Clttimeout = utils.IntPtr(10)
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "clttimeout")
	serviceInstance.Clttimeout = utils.IntPtr(0)

	//svrtimeout
	serviceInstance.Svrtimeout = utils.IntPtr(10)
	testHelperVerifyImmutabilityFunc(c, t, serviceType, serviceName, serviceInstance, "svrtimeout")
	serviceInstance.Svrtimeout = utils.IntPtr(0)
}

const testAccGslbserviceEnableDisable_enabled = `
resource "citrixadc_gslbsite" "tf_test_acc_gslbsite" {
  sitename = "tf_test_acc_gslbsite"
  siteipaddress = "192.168.22.33"
  sessionexchange = "DISABLED"
  sitepassword = "password123"
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
  sitepassword = "password123"
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

func TestAccGslbservice_enable_disable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbserviceDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccGslbserviceEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_test_acc_gslbservice", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_test_acc_gslbservice", "state", "ENABLED"),
				),
			},
			// Disable
			{
				Config: testAccGslbserviceEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_test_acc_gslbservice", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_test_acc_gslbservice", "state", "DISABLED"),
				),
			},
			// Re enable
			{
				Config: testAccGslbserviceEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_test_acc_gslbservice", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_test_acc_gslbservice", "state", "ENABLED"),
				),
			},
		},
	})
}

func TestAccGslbservice_lbmonitorbinding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbserviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicelbmonitor_two,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_test_gslbservice", nil),
					verifyLbmonitorbindingExists("tf_test_gslbservice", "tf_test_monitor1", false),
					verifyLbmonitorbindingExists("tf_test_gslbservice", "tf_test_monitor2", false),
				),
			},
			{
				Config: testAccGslbservicelbmonitor_none,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_test_gslbservice", nil),
					verifyLbmonitorbindingExists("tf_test_gslbservice", "tf_test_monitor1", true),
					verifyLbmonitorbindingExists("tf_test_gslbservice", "tf_test_monitor2", true),
				),
			},
			{
				Config: testAccGslbservicelbmonitor_one,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_test_gslbservice", nil),
					verifyLbmonitorbindingExists("tf_test_gslbservice", "tf_test_monitor1", true),
					verifyLbmonitorbindingExists("tf_test_gslbservice", "tf_test_monitor2", false),
				),
			},
		},
	})
}

func verifyLbmonitorbindingExists(servicename, monitor_name string, inverse bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		bindFound := false
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		lbmonitorBindings, _ := client.FindResourceArray("gslbservice_lbmonitor_binding", servicename)
		for _, val := range lbmonitorBindings {
			if val["monitor_name"].(string) == monitor_name {
				bindFound = true
				break
			}
		}

		if !inverse {
			if bindFound {
				return nil
			} else {
				return fmt.Errorf("Verify error cannot find bind for monitor %v for gslb service %v\n", monitor_name, servicename)
			}
		} else {
			if bindFound {
				return fmt.Errorf("Verify error found exessive bind for monitor %v for gslb service %v\n", monitor_name, servicename)
			} else {
				return nil
			}
		}
	}
}

const testAccGslbservicelbmonitor_two = `
resource "citrixadc_lbmonitor" "tf_test_monitor1" {
  monitorname = "tf_test_monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tf_test_monitor2" {
  monitorname = "tf_test_monitor2"
  type        = "PING"
}

resource "citrixadc_gslbsite" "tf_test_site" {
  sitename        = "tf_test_site"
  siteipaddress   = "192.168.22.19"
  sessionexchange = "DISABLED"
  sitepassword    = "password123"
}

resource "citrixadc_gslbservice" "tf_test_gslbservice" {
  ip          = "192.168.18.81"
  port        = "80"
  servicename = "tf_test_gslbservice"
  servicetype = "HTTP"
  sitename = citrixadc_gslbsite.tf_test_site.sitename

  lbmonitorbinding {
      monitor_name = citrixadc_lbmonitor.tf_test_monitor1.monitorname
      weight = 80
	}

  lbmonitorbinding {
      monitor_name = citrixadc_lbmonitor.tf_test_monitor2.monitorname
      weight = 20
	}
}
`
const testAccGslbservicelbmonitor_one = `
resource "citrixadc_lbmonitor" "tf_test_monitor1" {
  monitorname = "tf_test_monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tf_test_monitor2" {
  monitorname = "tf_test_monitor2"
  type        = "PING"
}

resource "citrixadc_gslbsite" "tf_test_site" {
  sitename        = "tf_test_site"
  siteipaddress   = "192.168.22.19"
  sessionexchange = "DISABLED"
  sitepassword    = "password123"
}

resource "citrixadc_gslbservice" "tf_test_gslbservice" {
  ip          = "192.168.18.81"
  port        = "80"
  servicename = "tf_test_gslbservice"
  servicetype = "HTTP"
  sitename = citrixadc_gslbsite.tf_test_site.sitename

  lbmonitorbinding {
      monitor_name = citrixadc_lbmonitor.tf_test_monitor2.monitorname
      weight = 20
	}
}
`

const testAccGslbservicelbmonitor_none = `
resource "citrixadc_lbmonitor" "tf_test_monitor1" {
  monitorname = "tf_test_monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tf_test_monitor2" {
  monitorname = "tf_test_monitor2"
  type        = "PING"
}

resource "citrixadc_gslbsite" "tf_test_site" {
  sitename        = "tf_test_site"
  siteipaddress   = "192.168.22.19"
  sessionexchange = "DISABLED"
  sitepassword    = "password123"
}

resource "citrixadc_gslbservice" "tf_test_gslbservice" {
  ip          = "192.168.18.81"
  port        = "80"
  servicename = "tf_test_gslbservice"
  servicetype = "HTTP"
  sitename = citrixadc_gslbsite.tf_test_site.sitename

}
`

func TestAccGslbservice_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_gslbservice.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbserviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservice_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbserviceExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Gslbservice.Type(), "gslb1vservice"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccGslbservice_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbserviceExist(resAddr, nil)),
			},
		},
	})
}

const testAccGslbservice_unset_step1 = `
resource "citrixadc_gslbsite" "tf_unset_site" {
  sitename        = "tf_unset_site"
  siteipaddress   = "192.168.44.55"
  sessionexchange = "DISABLED"
  sitepassword    = "password123"
}

resource "citrixadc_gslbservice" "tf_unset" {
  ip              = "192.168.44.60"
  port            = "80"
  servicename     = "tf_test_gslbservice_unset"
  servicetype     = "HTTP"
  sitename        = citrixadc_gslbsite.tf_unset_site.sitename
  appflowlog      = "DISABLED"
  cip             = "ENABLED"
  healthmonitor   = "NO"
}
`

const testAccGslbservice_unset_step2 = `
resource "citrixadc_gslbsite" "tf_unset_site" {
  sitename        = "tf_unset_site"
  siteipaddress   = "192.168.44.55"
  sessionexchange = "DISABLED"
  sitepassword    = "password123"
}

resource "citrixadc_gslbservice" "tf_unset" {
  ip          = "192.168.44.60"
  port        = "80"
  servicename = "tf_test_gslbservice_unset"
  servicetype = "HTTP"
  sitename    = citrixadc_gslbsite.tf_unset_site.sitename
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccGslbservice_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbserviceDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccGslbservice_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_unset", "appflowlog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_unset", "cip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_unset", "healthmonitor", "NO"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccGslbservice_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_unset", "appflowlog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_unset", "cip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservice.tf_unset", "healthmonitor", "YES"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckGslbserviceADCValue("tf_test_gslbservice_unset", "appflowlog", "ENABLED"),
					testAccCheckGslbserviceADCValue("tf_test_gslbservice_unset", "healthmonitor", "YES"),
				),
			},
		},
	})
}

// testAccCheckGslbserviceADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckGslbserviceADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Gslbservice.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("gslbservice %s not found on appliance", name)
		}
		got := fmt.Sprintf("%v", data[attr])
		if got != want {
			return fmt.Errorf("gslbservice %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccGslbservice_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGslbserviceDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccGslbservice_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbserviceExist("citrixadc_gslbservice.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccGslbservice_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbserviceExist("citrixadc_gslbservice.foo", nil)),
			},
		},
	})
}

func TestAccGslbserviceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbserviceDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_gslbservice.tf_gslbservice_ds", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservice.tf_gslbservice_ds", "servicename", "tf_gslbservice_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservice.tf_gslbservice_ds", "ipaddress", "172.16.1.102"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservice.tf_gslbservice_ds", "port", "80"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservice.tf_gslbservice_ds", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservice.tf_gslbservice_ds", "sitename", "Site-DS-East-Coast"),
				),
			},
		},
	})
}

func TestAccGslbservice_import(t *testing.T) {
	const resAddr = "citrixadc_gslbservice.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbserviceDestroy,
		Steps: []resource.TestStep{
			{Config: testAccGslbservice_basic},
			{
				Config:                  testAccGslbservice_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccGslbserviceDataSource_basic = `

resource "citrixadc_gslbsite" "tf_gslbsite_ds" {
  siteipaddress = "172.31.11.21"
  sitename      = "Site-DS-East-Coast"
  sitepassword  = "password123"
}

resource "citrixadc_gslbservice" "tf_gslbservice_ds" {
  ip          = "172.16.1.102"
  port        = "80"
  servicename = "tf_gslbservice_ds"
  servicetype = "HTTP"
  sitename    = citrixadc_gslbsite.tf_gslbsite_ds.sitename
}

data "citrixadc_gslbservice" "tf_gslbservice_ds" {
  servicename = citrixadc_gslbservice.tf_gslbservice_ds.servicename
  depends_on  = [citrixadc_gslbservice.tf_gslbservice_ds]
}

`

// TestAccGslbservice_lbmonitorbinding_editweight edits only a NON-KEY sub-attribute
// (weight) of an lbmonitor binding while its diff key (monitor_name) is unchanged.
// A key-only reconciliation would silently drop the edit; the second step asserts
// the new weight actually reached the appliance (both on the box and in state).
func TestAccGslbservice_lbmonitorbinding_editweight(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbserviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicelbmonitor_editweight(80),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_test_gslbservice", nil),
					verifyLbmonitorbindingWeight("tf_test_gslbservice", "tf_test_monitor1", 80),
				),
			},
			{
				Config: testAccGslbservicelbmonitor_editweight(50),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("citrixadc_gslbservice.tf_test_gslbservice", nil),
					verifyLbmonitorbindingWeight("tf_test_gslbservice", "tf_test_monitor1", 50),
					resource.TestCheckTypeSetElemNestedAttrs(
						"citrixadc_gslbservice.tf_test_gslbservice", "lbmonitorbinding.*",
						map[string]string{"monitor_name": "tf_test_monitor1", "weight": "50"}),
				),
			},
		},
	})
}

// verifyLbmonitorbindingWeight asserts the weight bound for monitorName on the
// appliance equals expected.
func verifyLbmonitorbindingWeight(servicename, monitorName string, expected int64) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		bindings, _ := client.FindResourceArray("gslbservice_lbmonitor_binding", servicename)
		for _, val := range bindings {
			if mn, ok := val["monitor_name"].(string); ok && mn == monitorName {
				got, cErr := utils.ConvertToInt64(val["weight"])
				if cErr != nil {
					return fmt.Errorf("could not parse weight for monitor %s: %v", monitorName, cErr)
				}
				if got != expected {
					return fmt.Errorf("weight for monitor %s = %d on the appliance, expected %d", monitorName, got, expected)
				}
				return nil
			}
		}
		return fmt.Errorf("no lbmonitor binding for monitor %s on gslbservice %s", monitorName, servicename)
	}
}

func testAccGslbservicelbmonitor_editweight(weight int) string {
	return fmt.Sprintf(`
resource "citrixadc_lbmonitor" "tf_test_monitor1" {
  monitorname = "tf_test_monitor1"
  type        = "HTTP"
}

resource "citrixadc_gslbsite" "tf_test_site" {
  sitename        = "tf_test_site"
  siteipaddress   = "192.168.22.19"
  sessionexchange = "DISABLED"
  sitepassword    = "password123"
}

resource "citrixadc_gslbservice" "tf_test_gslbservice" {
  ip          = "192.168.18.81"
  port        = "80"
  servicename = "tf_test_gslbservice"
  servicetype = "HTTP"
  sitename    = citrixadc_gslbsite.tf_test_site.sitename

  lbmonitorbinding {
    monitor_name = citrixadc_lbmonitor.tf_test_monitor1.monitorname
    weight       = %d
  }
}
`, weight)
}
