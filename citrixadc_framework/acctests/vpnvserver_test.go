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
	"regexp"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccVpnvserver_add = `
	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}
	resource "citrixadc_vpnvserver" "foo" {
		name                     = "tf.citrix.example.com"
		servicetype              = "SSL"
		ipv46                    = "3.3.3.3"
		port                     = 443
		ipset                    = citrixadc_ipset.tf_ipset.name
		dtls                     = "OFF"
		downstateflush           = "DISABLED"
		listenpolicy             = "NONE"
		tcpprofilename           = "nstcp_default_XA_XD_profile"
		secureprivateaccess		= "ENABLED"
		accessrestrictedpageredirect = "NS"
		deviceposture 		  = "DISABLED"
	}
`

const testAccVpnvserver_update = `
	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}
	resource "citrixadc_vpnvserver" "foo" {
		name                     = "tf.citrix.example.com"
		servicetype              = "SSL"
		ipv46                    = "3.3.3.3"
		port                     = 443
		ipset                    = citrixadc_ipset.tf_ipset.name
		dtls                     = "OFF"
		downstateflush           = "ENABLED"
		listenpolicy             = "NONE"
		tcpprofilename           = "nstcp_default_XA_XD_profile"
		secureprivateaccess		= "DISABLED"
		deviceposture 		  = "ENABLED"
	}
`

func TestAccVpnvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverExist("citrixadc_vpnvserver.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "servicetype", "SSL"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "ipv46", "3.3.3.3"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "downstateflush", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "secureprivateaccess", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "deviceposture", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "accessrestrictedpageredirect", "NS"),
				),
			},
			{
				Config: testAccVpnvserver_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverExist("citrixadc_vpnvserver.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "servicetype", "SSL"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "ipv46", "3.3.3.3"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "downstateflush", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "secureprivateaccess", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "deviceposture", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckVpnvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver name is set")
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
		data, err := client.FindResource(service.Vpnvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnvserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnvserver_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnvserverExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpnvserver.Type(), "tf.citrix.example.com"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnvserverExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVpnvserver_import(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserverDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnvserver_add},
			{
				Config:                  testAccVpnvserver_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccVpnvserver_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnvserverDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccVpnvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnvserverExist("citrixadc_vpnvserver.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVpnvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnvserverExist("citrixadc_vpnvserver.foo", nil)),
			},
		},
	})
}

const testAccVpnvserver_dtls = `
	resource "citrixadc_vpnvserver" "dtls" {
		name           = "tf-acc-dtls-vserver"
		servicetype    = "DTLS"
		ipv46          = "192.0.2.76"
		port           = 443
		downstateflush = "DISABLED"
	}
`

// TestAccVpnvserver_dtlsDefaultUpgrade is a regression test for GH #1476. A DTLS
// vpnvserver created by the SDKv2 provider (2.3.2) must upgrade to the framework
// provider without a spurious, failing update. The framework schema previously put a
// static Default on cginfrahomepageredirect ("ENABLED") and dtls ("ON"), which are
// type-conditional and rejected by the appliance for servicetype DTLS (NITRO
// errorcode 3725). The upgrade therefore planned DISABLED->ENABLED / OFF->ON and the
// apply failed. With the defaults removed (UseStateForUnknown instead), step 2 is a
// clean no-op. On the buggy schema step 2 fails with errorcode 3725.
func TestAccVpnvserver_dtlsDefaultUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnvserverDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.3.2"},
				},
				Config: testAccVpnvserver_dtls,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnvserverExist("citrixadc_vpnvserver.dtls", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVpnvserver_dtls,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverExist("citrixadc_vpnvserver.dtls", nil),
					testAccCheckVpnvserverADCValue("tf-acc-dtls-vserver", "servicetype", "DTLS"),
				),
			},
		},
	})
}

const testAccVpnvserver_dtls_fresh = `
	resource "citrixadc_vpnvserver" "dtls_default" {
		name           = "tf-acc-dtls-fresh"
		servicetype    = "DTLS"
		ipv46          = "192.0.2.77"
		port           = 443
		downstateflush = "DISABLED"
	}
`

// TestAccVpnvserver_dtlsFreshCreate is a regression test for GH #1476 (fresh-create
// path). Creating a DTLS vpnvserver in a single apply must succeed. The gateway-only
// attributes (cginfrahomepageredirect, dtls, loginonce, logoutonsmartcardremoval,
// icmpvsrresponse, rhistate) are forced into the create payload by their schema
// Defaults but are rejected by the appliance for servicetype DTLS on the idempotent
// create (NITRO errorcode 1092 / 3725). The Create path strips the default-forced ones
// for DTLS, so the apply succeeds. Without the filter the apply fails. The implicit
// post-apply plan asserts idempotency.
func TestAccVpnvserver_dtlsFreshCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_dtls_fresh,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverExist("citrixadc_vpnvserver.dtls_default", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.dtls_default", "servicetype", "DTLS"),
					testAccCheckVpnvserverADCValue("tf-acc-dtls-fresh", "servicetype", "DTLS"),
				),
			},
		},
	})
}

const testAccVpnvserver_dtls_explicit_gw = `
	resource "citrixadc_vpnvserver" "dtls_gw" {
		name        = "tf-acc-dtls-gw"
		servicetype = "DTLS"
		ipv46       = "192.0.2.80"
		port        = 443
		loginonce   = "ON"
	}
`

// TestAccVpnvserver_dtlsExplicitGatewayAttr is the counterpart to the fresh-create fix:
// a gateway-only attribute the user EXPLICITLY sets on a DTLS vserver must NOT be
// silently dropped. Only the default-forced values are stripped (see
// TestAccVpnvserver_dtlsFreshCreate); an explicit value is left in the create payload so
// the appliance returns its own clear error (NITRO errorcode 1092 / 3725, "... type
// DTLS") instead of the provider producing an inconsistent result.
func TestAccVpnvserver_dtlsExplicitGatewayAttr(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccVpnvserver_dtls_explicit_gw,
				ExpectError: regexp.MustCompile(`DTLS`),
			},
		},
	})
}

const testAccVpnvserverDataSource_basic = `
	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}
	resource "citrixadc_vpnvserver" "foo" {
		name                     = "tf.citrix.example.com"
		servicetype              = "SSL"
		ipv46                    = "3.3.3.3"
		port                     = 443
		ipset                    = citrixadc_ipset.tf_ipset.name
		dtls                     = "OFF"
		downstateflush           = "DISABLED"
		listenpolicy             = "NONE"
		tcpprofilename           = "nstcp_default_XA_XD_profile"
		secureprivateaccess		= "ENABLED"
		accessrestrictedpageredirect = "NS"
		deviceposture 		  = "DISABLED"
	}

data "citrixadc_vpnvserver" "foo" {
	name = citrixadc_vpnvserver.foo.name
}
`

// testAccVpnvserver_unset_step1 sets a set of type-independent, mutable
// unset-eligible attributes to valid NON-default values on an SSL gateway
// vserver. step2 removes them so the provider must unset them (revert to the
// documented NITRO defaults).
const testAccVpnvserver_unset_step1 = `
	resource "citrixadc_vpnvserver" "tf_unset" {
		name                         = "tf_vpnvserver_unset"
		servicetype                  = "SSL"
		ipv46                        = "5.5.5.5"
		port                         = 443
		# cginfrahomepageredirect and dtls are intentionally excluded: they are
		# type-conditional (rejected on servicetype DTLS, GH #1476) and therefore NOT
		# unset-eligible, so removing them must not revert to a static default.
		appflowlog                   = "DISABLED"
		deviceposture                = "ENABLED"
		downstateflush               = "DISABLED"
		icmpvsrresponse              = "ACTIVE"
		loginonce                    = "ON"
		logoutonsmartcardremoval     = "ON"
		rhistate                     = "ACTIVE"
		secureprivateaccess          = "ENABLED"
	}
`

const testAccVpnvserver_unset_step2 = `
	resource "citrixadc_vpnvserver" "tf_unset" {
		name        = "tf_vpnvserver_unset"
		servicetype = "SSL"
		ipv46       = "5.5.5.5"
		port        = 443
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to their documented NITRO defaults).
	}
`

func TestAccVpnvserver_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserverDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVpnvserver_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverExist("citrixadc_vpnvserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "appflowlog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "deviceposture", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "downstateflush", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "icmpvsrresponse", "ACTIVE"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "loginonce", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "logoutonsmartcardremoval", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "rhistate", "ACTIVE"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "secureprivateaccess", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccVpnvserver_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverExist("citrixadc_vpnvserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "appflowlog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "deviceposture", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "downstateflush", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "icmpvsrresponse", "PASSIVE"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "loginonce", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "logoutonsmartcardremoval", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "rhistate", "PASSIVE"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.tf_unset", "secureprivateaccess", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVpnvserverADCValue("tf_vpnvserver_unset", "downstateflush", "ENABLED"),
					testAccCheckVpnvserverADCValue("tf_vpnvserver_unset", "secureprivateaccess", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckVpnvserverADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckVpnvserverADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnvserver.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnvserver %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vpnvserver %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVpnvserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "servicetype", "SSL"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "ipv46", "3.3.3.3"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "downstateflush", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "secureprivateaccess", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "deviceposture", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "accessrestrictedpageredirect", "NS"),
					// Runtime-binding proof + read-only runtime state exposed only by the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_vpnvserver.foo", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_vpnvserver.foo", "curstate"),
				),
			},
		},
	})
}
