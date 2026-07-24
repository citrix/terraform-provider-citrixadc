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

// CAUTION: nsmgmtparam controls the appliance management interface. Changing
// mgmthttpport / mgmthttpsport can sever the NITRO connection mid-test, so these
// tests deliberately exercise only httpdmaxclients and leave mgmthttpport /
// mgmthttpsport unset (defaults 80 / 443 retained on the ADC).

import (
	"fmt"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccNsmgmtparam_basic_step1 = `
	resource "citrixadc_nsmgmtparam" "tf_nsmgmtparam" {
		httpdmaxclients = 40
	}
`

const testAccNsmgmtparam_basic_step2 = `
	resource "citrixadc_nsmgmtparam" "tf_nsmgmtparam" {
		httpdmaxclients = 50
	}
`

func TestAccNsmgmtparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: never truly deleted from the ADC (state-only removal).
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsmgmtparam_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsmgmtparamExist("citrixadc_nsmgmtparam.tf_nsmgmtparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nsmgmtparam.tf_nsmgmtparam", "httpdmaxclients", "40"),
				),
			},
			{
				Config: testAccNsmgmtparam_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsmgmtparamExist("citrixadc_nsmgmtparam.tf_nsmgmtparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nsmgmtparam.tf_nsmgmtparam", "httpdmaxclients", "50"),
				),
			},
		},
	})
}

func TestAccNsmgmtparam_import(t *testing.T) {
	const resAddr = "citrixadc_nsmgmtparam.tf_nsmgmtparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: never truly deleted from the ADC (state-only removal).
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{Config: testAccNsmgmtparam_basic_step1},
			{
				Config:                  testAccNsmgmtparam_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckNsmgmtparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsmgmtparam name is set")
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
		// Singleton: get-all, no name argument.
		data, err := client.FindResource(service.Nsmgmtparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsmgmtparam %s not found", n)
		}

		return nil
	}
}

const testAccNsmgmtparamDataSource_basic = `

	resource "citrixadc_nsmgmtparam" "tf_nsmgmtparam" {
		httpdmaxclients = 60
	}

	data "citrixadc_nsmgmtparam" "tf_nsmgmtparam" {
		depends_on = [citrixadc_nsmgmtparam.tf_nsmgmtparam]
	}
`

func TestAccNsmgmtparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsmgmtparamDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsmgmtparam.tf_nsmgmtparam", "httpdmaxclients", "60"),
				),
			},
		},
	})
}

// Unset support (attributesToUnset in resource_nsmgmtparam.go) covers TWO
// attributes: mgmthttpport (default 80) and mgmthttpsport (default 443).
//
// Only mgmthttpsport is exercised here. mgmthttpport controls the appliance's
// management HTTP port, and the provider/NITRO connection in these tests is HTTP
// on port 80 (NS_URL=http://.../). Setting mgmthttpport to a non-default value
// takes effect IMMEDIATELY and moves the management HTTP interface off port 80,
// which severs the provider's own connection mid-apply (verified live: the set
// returns 200, then every subsequent GET on port 80 fails and the appliance is
// only reachable on the new port). That is a testbed-transport hazard, not a
// property of the unset logic (the unset itself works: reconnecting on the new
// port and issuing action=unset restores port 80). It therefore cannot be
// acceptance-tested over the port-80 NITRO connection without breaking the
// testbed. mgmthttpsport (the management HTTPS port) is orthogonal to the HTTP
// connection used here, so it can be set to a non-default and unset safely.
// httpdmaxclients is intentionally NOT set: nsmgmtparam rejects a "set" that
// carries httpdmaxclients equal to the value already configured (NITRO
// errorcode 2431, "The same HTTPD max clients/max request workers value cannot
// be configured"). It is not the attribute under test, so it is left out.
const testAccNsmgmtparam_unset_step1 = `
	resource "citrixadc_nsmgmtparam" "tf_unset" {
		mgmthttpsport = 8443
	}
`

const testAccNsmgmtparam_unset_step2 = `
	resource "citrixadc_nsmgmtparam" "tf_unset" {
		# mgmthttpsport removed from config -> provider must unset it (revert to 443)
	}
`

func TestAccNsmgmtparam_unset(t *testing.T) {
	// The resource's other tests (TestAccNsmgmtparam_basic / _import / DataSource)
	// have no skip guards -> they run on the default standalone testbed, so this
	// test adds none either.
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: never truly deleted from the ADC (state-only removal).
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				// Non-default value applies and persists.
				Config: testAccNsmgmtparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsmgmtparamExist("citrixadc_nsmgmtparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsmgmtparam.tf_unset", "mgmthttpsport", "8443"),
					// Sanity: mgmthttpport left at its default throughout (not touched).
					resource.TestCheckResourceAttr("citrixadc_nsmgmtparam.tf_unset", "mgmthttpport", "80"),
				),
			},
			{
				// Removing mgmthttpsport must unset it -> state reverts to the NITRO
				// default (443), and the implicit post-apply plan must be empty.
				Config: testAccNsmgmtparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsmgmtparamExist("citrixadc_nsmgmtparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsmgmtparam.tf_unset", "mgmthttpsport", "443"),
					resource.TestCheckResourceAttr("citrixadc_nsmgmtparam.tf_unset", "mgmthttpport", "80"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsmgmtparamADCValue("mgmthttpsport", "443"),
				),
			},
		},
	})
}

// testAccCheckNsmgmtparamADCValue asserts a nsmgmtparam attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually
// reverted it. nsmgmtparam is a singleton, so the GET takes no name argument.
func testAccCheckNsmgmtparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsmgmtparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsmgmtparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nsmgmtparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}
