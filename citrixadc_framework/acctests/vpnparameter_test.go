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

const testAccVpnparameter_add = `

	resource "citrixadc_vpnparameter" "tf_vpnparameter" {
		splitdns              = "LOCAL"
		sesstimeout           = 30
		clientsecuritylog     = "OFF"
		smartgroup            = 10
		splittunnel           = "ON"
		locallanaccess        = "ON"
		winsip                = "4.45.5.4"
		samesite              = "None"
		backendcertvalidation = "DISABLED"
		backendserversni      = "DISABLED"
		icasessiontimeout     = "OFF"
		iconwithreceiver      = "OFF"
		linuxpluginupgrade    = "Always"
		uitheme               = "DEFAULT"
		httpport              = [80]
		secureprivateaccess	= "ENABLED"
		maxiipperuser         = 5
		httptrackconnproxy	= "OFF"
		deviceposture = "DISABLED"
		backenddtls12 = "DISABLED"
		accessrestrictedpageredirect = "NS"
	}
`
const testAccVpnparameter_update = `

	resource "citrixadc_vpnparameter" "tf_vpnparameter" {
		splitdns              = "LOCAL"
		sesstimeout           = 30
		clientsecuritylog     = "OFF"
		smartgroup            = 10
		splittunnel           = "OFF"
		locallanaccess        = "OFF"
		winsip                = "4.45.5.4"
		samesite              = "None"
		backendcertvalidation = "DISABLED"
		backendserversni      = "DISABLED"
		icasessiontimeout     = "OFF"
		iconwithreceiver      = "OFF"
		linuxpluginupgrade    = "Always"
		uitheme               = "DEFAULT"
		httpport              = [80]
		secureprivateaccess	= "DISABLED"
		maxiipperuser         = 10
		httptrackconnproxy	= "ON"
		deviceposture = "ENABLED"
		backenddtls12 = "ENABLED"
	}
`

func TestAccVpnparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// vpnparameter resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnparameter_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnparameterExist("citrixadc_vpnparameter.tf_vpnparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "splittunnel", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "locallanaccess", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "secureprivateaccess", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "httptrackconnproxy", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "deviceposture", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "backenddtls12", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "accessrestrictedpageredirect", "NS"),
				),
			},
			{
				Config: testAccVpnparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnparameterExist("citrixadc_vpnparameter.tf_vpnparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "splittunnel", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "locallanaccess", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "secureprivateaccess", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "httptrackconnproxy", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "deviceposture", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "backenddtls12", "ENABLED"),
				),
			},
		},
	})
}

func TestAccVpnparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccVpnparameter_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnparameterExist("citrixadc_vpnparameter.tf_vpnparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVpnparameter_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnparameterExist("citrixadc_vpnparameter.tf_vpnparameter", nil)),
			},
		},
	})
}

func testAccCheckVpnparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnparameter name is set")
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
		data, err := client.FindResource(service.Vpnparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnparameter %s not found", n)
		}

		return nil
	}
}

func TestAccVpnparameter_import(t *testing.T) {
	const resAddr = "citrixadc_vpnparameter.tf_vpnparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccVpnparameter_add},
			{
				Config:            testAccVpnparameter_add,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// spoofiip and windowsclienttype have server defaults but this
				// appliance's GET omits them from the response, so they cannot
				// round-trip through import.
				ImportStateVerifyIgnore: []string{"spoofiip", "windowsclienttype"},
			},
		},
	})
}

const testAccVpnparameterDataSource_basic = `

	resource "citrixadc_vpnparameter" "tf_vpnparameter" {
		splitdns              = "LOCAL"
		sesstimeout           = 30
		clientsecuritylog     = "OFF"
		smartgroup            = 10
		splittunnel           = "ON"
		locallanaccess        = "ON"
		winsip                = "4.45.5.4"
		samesite              = "None"
		backendcertvalidation = "DISABLED"
		backendserversni      = "DISABLED"
		icasessiontimeout     = "OFF"
		iconwithreceiver      = "OFF"
		linuxpluginupgrade    = "Always"
		uitheme               = "DEFAULT"
		httpport              = [80]
		secureprivateaccess	= "ENABLED"
		maxiipperuser         = 5
		httptrackconnproxy	= "OFF"
		deviceposture = "DISABLED"
		backenddtls12 = "DISABLED"
		accessrestrictedpageredirect = "NS"
	}

data "citrixadc_vpnparameter" "tf_vpnparameter" {
	depends_on = [citrixadc_vpnparameter.tf_vpnparameter]
}
`

func TestAccVpnparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnparameter.tf_vpnparameter", "splittunnel", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnparameter.tf_vpnparameter", "locallanaccess", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnparameter.tf_vpnparameter", "secureprivateaccess", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnparameter.tf_vpnparameter", "httptrackconnproxy", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnparameter.tf_vpnparameter", "deviceposture", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnparameter.tf_vpnparameter", "backenddtls12", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnparameter.tf_vpnparameter", "accessrestrictedpageredirect", "NS"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnparameter.tf_vpnparameter", "splitdns", "LOCAL"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnparameter.tf_vpnparameter", "sesstimeout", "30"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnparameter.tf_vpnparameter", "clientsecuritylog", "OFF"),
				),
			},
		},
	})
}

// vpnparameter is a singleton. The unset test sets a broad set of mutable,
// spec-unsettable attributes to non-default values, then removes them all from
// config so the provider must issue a NITRO unset that reverts each to its
// documented default. Only attributes whose unset the appliance accepts as part
// of the atomic unset batch are included (others are omitted; see the resource).
const testAccVpnparameter_unset_step1 = `
	resource "citrixadc_vpnparameter" "tf_unset" {
		deviceposture           = "ENABLED"
		backenddtls12           = "ENABLED"
		backendserversni        = "ENABLED"
		backendcertvalidation   = "ENABLED"
		httptrackconnproxy      = "ON"
		splittunnel             = "ON"
		locallanaccess          = "ON"
		rfc1918                 = "ON"
		killconnections         = "ON"
		transparentinterception = "ON"
		windowsautologon        = "ON"
		sso                     = "ON"
		encryptcsecexp          = "DISABLED"
		securebrowse            = "DISABLED"
		clientcleanupprompt     = "OFF"
		ssocredential           = "SECONDARY"
		usemip                  = "OFF"
		useiip                  = "SPILLOVER"
		windowsclienttype       = "PLUGIN"
		spoofiip                = "OFF"
		proxylocalbypass        = "ENABLED"
		maxiipperuser           = 5
	}
`

const testAccVpnparameter_unset_step2 = `
	resource "citrixadc_vpnparameter" "tf_unset" {
		# All previously-set attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccVpnparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// vpnparameter resource does not have a DELETE operation.
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVpnparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnparameterExist("citrixadc_vpnparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "deviceposture", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "backenddtls12", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "backendserversni", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "backendcertvalidation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "httptrackconnproxy", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "splittunnel", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "locallanaccess", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "rfc1918", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "killconnections", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "transparentinterception", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "windowsautologon", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "sso", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "encryptcsecexp", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "securebrowse", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "clientcleanupprompt", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "ssocredential", "SECONDARY"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "usemip", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "useiip", "SPILLOVER"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "windowsclienttype", "PLUGIN"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "spoofiip", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "proxylocalbypass", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "maxiipperuser", "5"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccVpnparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnparameterExist("citrixadc_vpnparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "deviceposture", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "backenddtls12", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "backendserversni", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "backendcertvalidation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "httptrackconnproxy", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "splittunnel", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "locallanaccess", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "rfc1918", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "killconnections", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "transparentinterception", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "windowsautologon", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "sso", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "encryptcsecexp", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "securebrowse", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "clientcleanupprompt", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "ssocredential", "PRIMARY"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "usemip", "NS"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "useiip", "NOSPILLOVER"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "windowsclienttype", "AGENT"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "spoofiip", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "proxylocalbypass", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_unset", "maxiipperuser", "1"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVpnparameterADCValue("deviceposture", "DISABLED"),
					testAccCheckVpnparameterADCValue("splittunnel", "OFF"),
					testAccCheckVpnparameterADCValue("encryptcsecexp", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckVpnparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. vpnparameter is a singleton, so it is fetched with an empty name.
func testAccCheckVpnparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vpnparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}
