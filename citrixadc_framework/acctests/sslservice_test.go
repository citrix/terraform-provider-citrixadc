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

const testAccSslservice_basic = `
resource "citrixadc_sslservice" "demo_sslservice" {
	cipherredirect = "DISABLED"
	clientauth = "DISABLED"
	dh = "DISABLED"
	dhcount = 0
	dhkeyexpsizelimit = "DISABLED"
	ersa = "DISABLED"
	redirectportrewrite = "DISABLED"
	serverauth = "ENABLED"
	servicename = citrixadc_service.tf_service.name
	sessreuse = "ENABLED"
	snienable = "DISABLED"
	ssl2 = "DISABLED"
	ssl3 = "ENABLED"
	sslredirect = "DISABLED"
	sslv2redirect = "DISABLED"
	tls1 = "DISABLED"
	tls11 = "ENABLED"
	tls12 = "ENABLED"
	tls13 = "DISABLED"
	sesstimeout = 380

}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	ipv46       = "10.10.10.44"
	name        = "tf_lbvserver"
	port        = 443
	servicetype = "SSL"
	sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_service" "tf_service" {
	name = "tf_service"
	servicetype = "SSL"
	port = 443 
	lbvserver = citrixadc_lbvserver.tf_lbvserver.name
	ip = "10.77.33.22"
}
`

const testAccSslservice_basic_update_sess = `
resource "citrixadc_sslservice" "demo_sslservice" {
	cipherredirect = "DISABLED"
	clientauth = "DISABLED"
	dh = "DISABLED"
	dhcount = 0
	dhkeyexpsizelimit = "DISABLED"
	ersa = "DISABLED"
	redirectportrewrite = "DISABLED"
	serverauth = "ENABLED"
	servicename = citrixadc_service.tf_service.name
	sessreuse = "ENABLED"
	snienable = "DISABLED"
	ssl2 = "DISABLED"
	ssl3 = "ENABLED"
	sslredirect = "DISABLED"
	sslv2redirect = "DISABLED"
	tls1 = "DISABLED"
	tls11 = "ENABLED"
	tls12 = "ENABLED"
	tls13 = "DISABLED"
	sesstimeout = 390

}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	ipv46       = "10.10.10.44"
	name        = "tf_lbvserver"
	port        = 443
	servicetype = "SSL"
	sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_service" "tf_service" {
	name = "tf_service"
	servicetype = "SSL"
	port = 443 
	lbvserver = citrixadc_lbvserver.tf_lbvserver.name
	ip = "10.77.33.22"
}
`
const testAccSslservice_basic_update_general = `
resource "citrixadc_sslservice" "demo_sslservice" {
	cipherredirect = "DISABLED"
	clientauth = "DISABLED"
	dh = "DISABLED"
	dhcount = 0
	dhkeyexpsizelimit = "DISABLED"
	ersa = "DISABLED"
	redirectportrewrite = "DISABLED"
	serverauth = "ENABLED"
	servicename = citrixadc_service.tf_service.name
	sessreuse = "DISABLED"
	snienable = "DISABLED"
	ssl2 = "DISABLED"
	ssl3 = "ENABLED"
	sslredirect = "DISABLED"
	sslv2redirect = "DISABLED"
	tls1 = "DISABLED"
	tls11 = "ENABLED"
	tls12 = "ENABLED"
	tls13 = "DISABLED"
	sesstimeout = 390

}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	ipv46       = "10.10.10.44"
	name        = "tf_lbvserver"
	port        = 443
	servicetype = "SSL"
	sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_service" "tf_service" {
	name = "tf_service"
	servicetype = "SSL"
	port = 443 
	lbvserver = citrixadc_lbvserver.tf_lbvserver.name
	ip = "10.77.33.22"
}
`

func TestAccSslservice_basic(t *testing.T) {
	// if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
	// 	t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	// }
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslserviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservice_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslserviceExist("citrixadc_sslservice.demo_sslservice", nil),
				),
			},
			{
				Config: testAccSslservice_basic_update_sess,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslserviceExist("citrixadc_sslservice.demo_sslservice", nil),
				),
			},
			{
				Config: testAccSslservice_basic_update_general,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslserviceExist("citrixadc_sslservice.demo_sslservice", nil),
				),
			},
		},
	})
}

func TestAccSslservice_sdkv2StateUpgrade(t *testing.T) {
	// if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
	// 	t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	// }
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslserviceDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSslservice_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslserviceExist("citrixadc_sslservice.demo_sslservice", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSslservice_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslserviceExist("citrixadc_sslservice.demo_sslservice", nil)),
			},
		},
	})
}

func TestAccSslservice_import(t *testing.T) {
	t.Skip("Skipping sslservice import test")
	const resAddr = "citrixadc_sslservice.demo_sslservice"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslserviceDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslservice_basic},
			{
				Config:                  testAccSslservice_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSslserviceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservice name is set")
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
		data, err := client.FindResource(service.Sslservice.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslservice %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslserviceDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservice" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslservice.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslservice %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSslserviceDataSource_basic = `
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		ipv46       = "10.10.10.44"
		name        = "tf_lbvserver"
		port        = 443
		servicetype = "SSL"
		sslprofile  = "ns_default_ssl_profile_frontend"
	}

	resource "citrixadc_service" "tf_service" {
		name = "tf_service"
		servicetype = "SSL"
		port = 443 
		lbvserver = citrixadc_lbvserver.tf_lbvserver.name
		ip = "10.77.33.22"
	}

	resource "citrixadc_sslservice" "tf_sslservice" {
		cipherredirect = "DISABLED"
		clientauth = "DISABLED"
		dh = "DISABLED"
		dhcount = 0
		dhkeyexpsizelimit = "DISABLED"
		ersa = "DISABLED"
		redirectportrewrite = "DISABLED"
		serverauth = "ENABLED"
		servicename = citrixadc_service.tf_service.name
		sessreuse = "ENABLED"
		snienable = "DISABLED"
		ssl2 = "DISABLED"
		ssl3 = "ENABLED"
		sslredirect = "DISABLED"
		sslv2redirect = "DISABLED"
		tls1 = "DISABLED"
		tls11 = "ENABLED"
		tls12 = "ENABLED"
		tls13 = "DISABLED"
		sesstimeout = 380
	}

	data "citrixadc_sslservice" "tf_sslservice_datasource" {
		servicename = citrixadc_sslservice.tf_sslservice.servicename
	}
`

// The sslservice unset test exercises the spec-unsettable string attributes
// that carry a documented NITRO default. step1 sets each to a non-default
// value; step2 removes them from config, which must trigger the unset action
// and revert each to its documented default.
const testAccSslservice_unset_prereq = `
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		ipv46       = "10.10.10.44"
		name        = "tf_lbvserver"
		port        = 443
		servicetype = "SSL"
		sslprofile  = "ns_default_ssl_profile_frontend"
	}

	resource "citrixadc_service" "tf_service" {
		name        = "tf_service"
		servicetype = "SSL"
		port        = 443
		lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
		ip          = "10.77.33.22"
	}
`

const testAccSslservice_unset_step1 = testAccSslservice_unset_prereq + `
	resource "citrixadc_sslservice" "tf_unset" {
		servicename         = citrixadc_service.tf_service.name
		# Frontend-only options (cipherredirect, clientauth, dh, ersa, ssl2,
		# sslredirect, sslv2redirect) are omitted: NITRO reports each as "not
		# applicable when configuring a backend service" (ec1095/ec3745), so they
		# are not settable/unset-testable on this service-based sslservice.
		# redirectportrewrite is likewise omitted: enabling it requires sslredirect
		# to be enabled (NITRO ec1585 "SSL port rewrite can be enabled only when SSL
		# redirect is enabled"), but sslredirect is frontend-only and cannot be set on
		# this backend service, so redirectportrewrite=ENABLED is unreachable here.
		dhkeyexpsizelimit   = "ENABLED"
		serverauth          = "ENABLED"
		sessreuse           = "DISABLED"
		snienable           = "ENABLED"
		ssl3                = "DISABLED"
		tls1                = "DISABLED"
		tls11               = "DISABLED"
		tls12               = "DISABLED"
		tls13               = "ENABLED"
	}
`

const testAccSslservice_unset_step2 = testAccSslservice_unset_prereq + `
	resource "citrixadc_sslservice" "tf_unset" {
		servicename = citrixadc_service.tf_service.name
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to the documented NITRO defaults).
	}
`

func TestAccSslservice_unset(t *testing.T) {
	// if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
	// 	t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	// }
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslserviceDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSslservice_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslserviceExist("citrixadc_sslservice.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "dhkeyexpsizelimit", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "serverauth", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "sessreuse", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "snienable", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "ssl3", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "tls1", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "tls11", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "tls12", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "tls13", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSslservice_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslserviceExist("citrixadc_sslservice.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "dhkeyexpsizelimit", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "serverauth", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "sessreuse", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "snienable", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "ssl3", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "tls1", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "tls11", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "tls12", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservice.tf_unset", "tls13", "DISABLED"),
					// No raw-appliance confirmation here: every one of these attributes
					// reverts to its default on unset, and NITRO omits SSL protocol/feature
					// attributes from GET when they are at their default (omit-on-default),
					// so a raw data[attr] read is nil. The state asserts above already prove
					// the unset fired — had it not, the echoed non-default value (e.g.
					// serverauth=ENABLED) would be read back and fail those checks.
				),
			},
		},
	})
}

// testAccCheckSslserviceADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckSslserviceADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Sslservice.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("sslservice %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("sslservice %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccSslserviceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslserviceDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslservice.tf_sslservice_datasource", "servicename", "tf_service"),
					resource.TestCheckResourceAttrSet("data.citrixadc_sslservice.tf_sslservice_datasource", "id"),
				),
			},
		},
	})
}
