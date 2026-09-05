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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccSslvserver_basic = `
	resource "citrixadc_sslvserver" "tf_sslvserver" {
		cipherredirect = "ENABLED"
		cipherurl = "http://www.citrix.com"
		cleartextport = "80"
		clientauth = "ENABLED"
		clientcert = "Optional"
		hsts = "ENABLED"
		includesubdomains = "YES"
		maxage = "100"
		ocspstapling = "ENABLED"
		preload = "YES"
		sendclosenotify = "YES"
		sessreuse = "ENABLED"
		sesstimeout = "180"
		snienable = "ENABLED"
		sslredirect = "ENABLED"
		strictsigdigestcheck = "ENABLED"
		tls1 = "ENABLED"
		tls11 = "ENABLED"
		tls12 = "ENABLED"
		tls13 = "ENABLED"
		tls13sessionticketsperauthcontext = "7"
		zerorttearlydata = "ENABLED"
		vservername = citrixadc_lbvserver.tf_lbvserver.name
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_vserver"
		servicetype = "SSL"
	}
`

func TestAccSslvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserverExist("citrixadc_sslvserver.tf_sslvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "cipherredirect", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "cipherurl", "http://www.citrix.com"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "cleartextport", "80"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "clientauth", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "clientcert", "Optional"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "hsts", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "includesubdomains", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "maxage", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "ocspstapling", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "preload", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "sendclosenotify", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "sessreuse", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "sesstimeout", "180"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "snienable", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "sslredirect", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "strictsigdigestcheck", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "tls1", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "tls11", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "tls12", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "tls13", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "tls13sessionticketsperauthcontext", "7"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "zerorttearlydata", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_sslvserver", "vservername", "tf_vserver"),
				),
			},
		},
	})
}

func testAccCheckSslvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslvserver name is set")
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
		data, err := client.FindResource(service.Sslvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslvserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslvserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslvserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslvserver_import(t *testing.T) {
	const resAddr = "citrixadc_sslvserver.tf_sslvserver"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserverDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslvserver_basic},
			{
				Config:                  testAccSslvserver_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccSslvserverDataSource_basic = `
	resource "citrixadc_sslvserver" "tf_sslvserver" {
		cipherredirect = "ENABLED"
		cipherurl = "http://www.citrix.com"
		cleartextport = "80"
		clientauth = "ENABLED"
		clientcert = "Optional"
		hsts = "ENABLED"
		includesubdomains = "YES"
		maxage = "100"
		ocspstapling = "ENABLED"
		preload = "YES"
		sendclosenotify = "YES"
		sessreuse = "ENABLED"
		sesstimeout = "180"
		snienable = "ENABLED"
		sslredirect = "ENABLED"
		strictsigdigestcheck = "ENABLED"
		tls1 = "ENABLED"
		tls11 = "ENABLED"
		tls12 = "ENABLED"
		tls13 = "ENABLED"
		tls13sessionticketsperauthcontext = "7"
		zerorttearlydata = "ENABLED"
		vservername = citrixadc_lbvserver.tf_lbvserver.name
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_vserver"
		servicetype = "SSL"
	}

	data "citrixadc_sslvserver" "tf_sslvserver_datasource" {
		vservername = citrixadc_sslvserver.tf_sslvserver.vservername
	}
`

func TestAccSslvserver_sdkv2StateUpgrade(t *testing.T) {
	// No upgrade baseline possible: the released citrix/citrixadc 2.2.0 provider
	// CRASHES ("plugin crashed") applying this sslvserver config, so step 1 never
	// establishes SDK-v2 state. The current provider is unaffected.
	t.Skip("no 2.2.0 baseline: released 2.2.0 provider crashes on sslvserver apply")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslvserverDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSslvserver_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslvserverExist("citrixadc_sslvserver.tf_sslvserver", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSslvserver_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslvserverExist("citrixadc_sslvserver.tf_sslvserver", nil)),
			},
		},
	})
}

const testAccSslvserver_unset_step1 = `
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_vserver_unset"
		servicetype = "SSL"
	}

	resource "citrixadc_sslvserver" "tf_unset" {
		vservername                       = citrixadc_lbvserver.tf_lbvserver.name
		cipherredirect                    = "ENABLED"
		cleartextport                     = 80
		clientauth                        = "ENABLED"
		ersa                              = "DISABLED"
		hsts                              = "ENABLED"
		ocspstapling                      = "ENABLED"
		redirectportrewrite               = "ENABLED"
		sendclosenotify                   = "NO"
		sessreuse                         = "ENABLED"
		sesstimeout                       = 180
		snienable                         = "ENABLED"
		ssl3                              = "DISABLED"
		sslclientlogs                     = "ENABLED"
		sslredirect                       = "ENABLED"
		strictsigdigestcheck              = "ENABLED"
		tls1                              = "DISABLED"
		tls11                             = "DISABLED"
		tls12                             = "DISABLED"
		tls13                             = "ENABLED"
		tls13sessionticketsperauthcontext = 7
	}
`

const testAccSslvserver_unset_step2 = `
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_vserver_unset"
		servicetype = "SSL"
	}

	resource "citrixadc_sslvserver" "tf_unset" {
		vservername = citrixadc_lbvserver.tf_lbvserver.name
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccSslvserver_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserverDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSslvserver_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserverExist("citrixadc_sslvserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "cipherredirect", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "cleartextport", "80"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "clientauth", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "ersa", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "hsts", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "ocspstapling", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "redirectportrewrite", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "sendclosenotify", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "sesstimeout", "180"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "snienable", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "ssl3", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "sslclientlogs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "sslredirect", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "strictsigdigestcheck", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "tls1", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "tls11", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "tls12", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "tls13", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "tls13sessionticketsperauthcontext", "7"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSslvserver_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserverExist("citrixadc_sslvserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "cipherredirect", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "cleartextport", "0"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "clientauth", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "ersa", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "hsts", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "ocspstapling", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "redirectportrewrite", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "sendclosenotify", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "sesstimeout", "120"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "snienable", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "ssl3", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "sslclientlogs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "sslredirect", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "strictsigdigestcheck", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "tls1", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "tls11", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "tls12", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "tls13", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver.tf_unset", "tls13sessionticketsperauthcontext", "1"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSslvserverADCValue("tf_vserver_unset", "cipherredirect", "DISABLED"),
					testAccCheckSslvserverADCValue("tf_vserver_unset", "hsts", "DISABLED"),
					testAccCheckSslvserverADCValue("tf_vserver_unset", "tls13", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckSslvserverADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckSslvserverADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Sslvserver.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("sslvserver %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("sslvserver %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccSslvserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslvserver.tf_sslvserver_datasource", "vservername", "tf_vserver"),
					resource.TestCheckResourceAttrSet("data.citrixadc_sslvserver.tf_sslvserver_datasource", "id"),
				),
			},
		},
	})
}
