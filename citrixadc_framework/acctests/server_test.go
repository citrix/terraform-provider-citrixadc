/*
Copyright 2016 Citrix Systems, Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.o	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		fv := reflect.ValueOf(client).Elem().FieldByName("headers")

		if fmt.Sprintf("%v", fv) == "map[User-Agent:terraform-ctxadc]" {
			return nil
		} else {/LICENSE-2.0

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
	"reflect"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

func TestAccServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServer_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExist("citrixadc_server.foo", nil),
					testAccCheckUserAgent(),
				),
			},
		},
	})
}

func testAccCheckServerExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No server name is set")
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
		data, err := client.FindResource(service.Server.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("server %s not found", n)
		}

		return nil
	}
}

func testAccCheckServerDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_server" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Server.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("server %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccServer_basic = `


resource "citrixadc_server" "foo" {
	name = "test_server"
	ipaddress = "192.168.11.13"
}
`

// Test for immutability of attributes
// This is to catch any attibute having ForceNew: true while not actually needed
func TestAccServer_AssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	serverName := "tf-acc-server-name"
	serverType := service.Server.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, serverType, serverName, nil)

	serverInstance := basic.Server{
		Domain:      "tfacc.domain.com",
		Ipv6address: "YES",
		Name:        serverName,
		Td:          utils.IntPtr(0),
	}

	if _, err := c.client.AddResource(serverType, serverName, serverInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// Verify immutability of argument td
	serverInstance.Domain = ""
	serverInstance.Ipv6address = ""
	serverInstance.Td = utils.IntPtr(10)
	testHelperVerifyImmutabilityFunc(c, t, serverType, serverName, serverInstance, "td")
	serverInstance.Td = utils.IntPtr(0)

	// Verify immutability of argument domain
	serverInstance.Domain = "newdomain.com"
	serverInstance.Ipv6address = ""
	testHelperVerifyImmutabilityFunc(c, t, serverType, serverName, serverInstance, "domain")
	serverInstance.Domain = ""

	// Verify immutability of argument ipv6address
	serverInstance.Ipv6address = "YES"
	serverInstance.Td = utils.IntPtr(0)
	testHelperVerifyImmutabilityFunc(c, t, serverType, serverName, serverInstance, "ipv6address")
	serverInstance.Ipv6address = ""
}

const testAccServerEnableDisable_enabled = `
resource "citrixadc_server" "tf_enable_disable_test_svr" {
	name = "tf_enable_disable_test_svr"
	ipaddress = "192.168.43.33"
	comment = "enabled state comment"
	state = "ENABLED"
	graceful = "YES"
	delay = 60
}
`

const testAccServerEnableDisable_disabled = `
resource "citrixadc_server" "tf_enable_disable_test_svr" {
	name = "tf_enable_disable_test_svr"
	ipaddress = "192.168.43.33"
	comment = "disabled state comment"
	state = "DISABLED"
	graceful = "YES"
	delay = 60
}
`

func TestAccServer_enable_disable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServerDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccServerEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExist("citrixadc_server.tf_enable_disable_test_svr", nil),
					resource.TestCheckResourceAttr("citrixadc_server.tf_enable_disable_test_svr", "state", "ENABLED"),
				),
			},
			// Disable
			{
				Config: testAccServerEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExist("citrixadc_server.tf_enable_disable_test_svr", nil),
					resource.TestCheckResourceAttr("citrixadc_server.tf_enable_disable_test_svr", "state", "DISABLED"),
				),
			},
			// Re enable
			{
				Config: testAccServerEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExist("citrixadc_server.tf_enable_disable_test_svr", nil),
					resource.TestCheckResourceAttr("citrixadc_server.tf_enable_disable_test_svr", "state", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckUserAgent() resource.TestCheckFunc {
	// TODO check logs of ADC for presence of user agent string
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		fv := reflect.ValueOf(client).Elem().FieldByName("headers")

		if fmt.Sprintf("%v", fv) == "map[User-Agent:terraform-ctxadc]" {
			return nil
		} else {
			return fmt.Errorf("Could not verify headers. fv is %v", fv)
		}
	}
}

const testAccServerDataSource_basic = `

	resource "citrixadc_server" "tf_server" {
		name      = "test_server_ds"
		ipaddress = "192.168.11.14"
	}

	data "citrixadc_server" "tf_server" {
		name       = citrixadc_server.tf_server.name
		depends_on = [citrixadc_server.tf_server]
	}
`

func TestAccServer_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_server.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServer_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckServerExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Server.Type(), "test_server"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccServer_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckServerExist(resAddr, nil)),
			},
		},
	})
}

func TestAccServer_import(t *testing.T) {
	const resAddr = "citrixadc_server.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServerDestroy,
		Steps: []resource.TestStep{
			{Config: testAccServer_basic},
			{
				Config:                  testAccServer_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccServer_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckServerDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccServer_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckServerExist("citrixadc_server.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccServer_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckServerExist("citrixadc_server.foo", nil)),
			},
		},
	})
}

// testAccCheckServerADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. An attribute NITRO omits from GET (its default) is treated as "".
func testAccCheckServerADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Server.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("server %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("server %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccServer_unset_step1 = `
resource "citrixadc_server" "tf_unset" {
	name      = "tf_test_server_unset"
	ipaddress = "192.168.77.61"
	comment   = "unset test comment"
}
`

const testAccServer_unset_step2 = `
resource "citrixadc_server" "tf_unset" {
	name      = "tf_test_server_unset"
	ipaddress = "192.168.77.61"
	# comment removed from config -> the provider must unset it (revert to the
	# NITRO default, empty).
}
`

func TestAccServer_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServerDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccServer_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExist("citrixadc_server.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_server.tf_unset", "comment", "unset test comment"),
					testAccCheckServerADCValue("tf_test_server_unset", "comment", "unset test comment"),
				),
			},
			{
				// Removing the attribute must unset it: the appliance reverts to
				// the documented NITRO default (empty) and the implicit post-apply
				// plan must be empty.
				Config: testAccServer_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExist("citrixadc_server.tf_unset", nil),
					testAccCheckServerADCValue("tf_test_server_unset", "comment", ""),
				),
			},
		},
	})
}

func TestAccServerDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServerDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_server.tf_server", "name", "test_server_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_server.tf_server", "ipaddress", "192.168.11.14"),
				),
			},
		},
	})
}
