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
	"net/url"
	"os"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// DANGER: citrixadc_nsconfig_update issues `set ns config`, which changes the
// appliance NSIP/netmask. This test is SAFE ONLY because every settable param is
// pinned to the running box's OWN current value, making the `set ns config` an
// effective no-op that CANNOT disconnect the box. `ipaddress` is derived from the
// NS_URL supplied to the acceptance run and `netmask` is read from the live box,
// so the test is portable to whichever standalone box NS_URL points at (rather
// than being hardcoded to one appliance). nsvlan/ifnum/tagged are intentionally
// omitted so the management data path is never disturbed, and httpport/cipheader/
// ftpportrange/crportrange are omitted because they are unset on the box (setting
// them would not be a no-op). The remaining attributes below exercise the new
// attribute code paths end to end while remaining non-disruptive.
//
// The two %s placeholders are filled with the NS_URL host (ipaddress) and the
// box's current netmask at test time.
const testAccNsconfigUpdate_basic = `
	resource "citrixadc_nsconfig_update" "foo" {
		ipaddress = "%s"
		netmask   = "%s"

		maxconn                 = 0
		maxreq                  = 0
		cip                     = "DISABLED"
		cookieversion           = "0"
		securecookie            = "ENABLED"
		pmtumin                 = 576
		pmtutimeout             = 10
		timezone                = "CoordinatedUniversalTime"
		grantquotamaxclient     = 10
		exclusivequotamaxclient = 80
		grantquotaspillover     = 10
		exclusivequotaspillover = 80
		securemanagementtraffic = "DISABLED"
		securemanagementtd      = 0
	}
`

// nsconfigSelfIPNetmask returns the appliance's own NSIP (parsed from NS_URL) and
// its current netmask (read from the live box). On non-acceptance/unit runs where
// NS_URL is unset, it returns placeholders; resource.Test skips the test anyway.
func nsconfigSelfIPNetmask(t *testing.T) (string, string) {
	t.Helper()
	u, err := url.Parse(os.Getenv("NS_URL"))
	if err != nil || u.Hostname() == "" {
		return "127.0.0.1", "255.255.255.0"
	}
	ipaddress := u.Hostname()
	netmask := "255.255.255.0"
	if client, cerr := testAccGetFrameworkClient(); cerr == nil {
		if data, derr := client.FindResource(service.Nsconfig.Type(), ""); derr == nil {
			if v, ok := data["netmask"]; ok {
				netmask = fmt.Sprintf("%v", v)
			}
		}
	}
	return ipaddress, netmask
}

func TestAccNsconfigUpdate_basic(t *testing.T) {
	ipaddress, netmask := nsconfigSelfIPNetmask(t)
	config := fmt.Sprintf(testAccNsconfigUpdate_basic, ipaddress, netmask)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsconfigUpdateExist("citrixadc_nsconfig_update.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "ipaddress", ipaddress),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "netmask", netmask),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "maxconn", "0"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "maxreq", "0"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "cip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "cookieversion", "0"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "securecookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "pmtumin", "576"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "pmtutimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "timezone", "CoordinatedUniversalTime"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "grantquotamaxclient", "10"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "exclusivequotamaxclient", "80"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "grantquotaspillover", "10"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "exclusivequotaspillover", "80"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "securemanagementtraffic", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.foo", "securemanagementtd", "0"),
				),
			},
		},
	})
}

// Step 1 sets three behavioral params to NON-default values; step 2 removes them
// so the provider issues ?action=unset and the appliance restores the defaults.
// ipaddress/netmask are pinned to the box's own values in both steps (safe no-op).
// The %s placeholders are the NS_URL host and the box's netmask.
const testAccNsconfigUpdate_unset_step1 = `
	resource "citrixadc_nsconfig_update" "unset" {
		ipaddress = "%s"
		netmask   = "%s"

		pmtutimeout         = 20
		grantquotamaxclient = 20
		securecookie        = "DISABLED"
	}
`

const testAccNsconfigUpdate_unset_step2 = `
	resource "citrixadc_nsconfig_update" "unset" {
		ipaddress = "%s"
		netmask   = "%s"
	}
`

func TestAccNsconfigUpdate_unsetOnRemove(t *testing.T) {
	ipaddress, netmask := nsconfigSelfIPNetmask(t)
	step1 := fmt.Sprintf(testAccNsconfigUpdate_unset_step1, ipaddress, netmask)
	step2 := fmt.Sprintf(testAccNsconfigUpdate_unset_step2, ipaddress, netmask)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Set the three behavioral params to non-default values.
				Config: step1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.unset", "pmtutimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.unset", "grantquotamaxclient", "20"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_update.unset", "securecookie", "DISABLED"),
					testAccCheckNsconfigAdcValues(map[string]string{
						"pmtutimeout":         "20",
						"grantquotamaxclient": "20",
						"securecookie":        "DISABLED",
					}),
				),
			},
			{
				// Remove them from config -> provider unsets them -> appliance
				// defaults restored.
				Config: step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsconfigAdcValues(map[string]string{
						"pmtutimeout":         "10",
						"grantquotamaxclient": "10",
						"securecookie":        "ENABLED",
					}),
				),
			},
		},
	})
}

// testAccCheckNsconfigAdcValues reads the live nsconfig and asserts the given
// attribute -> expected-value pairs.
func testAccCheckNsconfigAdcValues(want map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Error creating client for nsconfig read-back: %s", err.Error())
		}
		data, err := client.FindResource(service.Nsconfig.Type(), "")
		if err != nil {
			return fmt.Errorf("Error reading nsconfig from ADC: %s", err.Error())
		}
		for attr, exp := range want {
			got, ok := data[attr]
			if !ok {
				return fmt.Errorf("nsconfig %s missing from ADC response", attr)
			}
			if fmt.Sprintf("%v", got) != exp {
				return fmt.Errorf("nsconfig %s: want %s, adc reports %v", attr, exp, got)
			}
		}
		return nil
	}
}

func testAccCheckNsconfigUpdateExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No NsConfigUpdate is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Read the live nsconfig back from the ADC and confirm the settable
		// params we applied match what the appliance now reports.
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Error creating client for nsconfig read-back: %s", err.Error())
		}
		data, err := client.FindResource(service.Nsconfig.Type(), "")
		if err != nil {
			return fmt.Errorf("Error reading nsconfig from ADC: %s", err.Error())
		}

		wantIP := rs.Primary.Attributes["ipaddress"]
		if got, ok := data["ipaddress"]; ok && wantIP != "" {
			if fmt.Sprintf("%v", got) != wantIP {
				return fmt.Errorf("nsconfig ipaddress mismatch: state=%s adc=%v", wantIP, got)
			}
		}
		wantNetmask := rs.Primary.Attributes["netmask"]
		if got, ok := data["netmask"]; ok && wantNetmask != "" {
			if fmt.Sprintf("%v", got) != wantNetmask {
				return fmt.Errorf("nsconfig netmask mismatch: state=%s adc=%v", wantNetmask, got)
			}
		}

		// Cross-check a representative sample of the newly added settable params
		// against what the appliance reports, confirming they round-trip.
		for _, attr := range []string{
			"maxconn", "maxreq", "cip", "cookieversion", "securecookie",
			"pmtumin", "pmtutimeout", "timezone", "grantquotamaxclient",
			"exclusivequotamaxclient", "grantquotaspillover",
			"exclusivequotaspillover", "securemanagementtraffic", "securemanagementtd",
		} {
			want := rs.Primary.Attributes[attr]
			if got, ok := data[attr]; ok && want != "" {
				if fmt.Sprintf("%v", got) != want {
					return fmt.Errorf("nsconfig %s mismatch: state=%s adc=%v", attr, want, got)
				}
			}
		}
		return nil
	}
}
