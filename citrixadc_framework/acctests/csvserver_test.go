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

	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

const testAccCsvserver_basic = `

resource "citrixadc_csvserver" "foo" {

  ipv46 = "10.202.11.11"
  name = "terraform-cs"
  port = 8080
  servicetype = "HTTP"
  v6persistmasklen = 128
  # timeout is the persistence-session timeout; the ADC only honors a
  # non-default value when a persistence type is configured. Without
  # persistencetype the ADC silently resets timeout to its default (2),
  # producing a perpetual diff. Configure persistence so timeout sticks.
  persistencetype = "SOURCEIP"
  timeout = 180
  probesuccessresponsecode = "200-399"
  probeprotocol = "HTTP"
  probeport = 8085
  persistencebackup = "NONE"
  dtls = "OFF"
  backuppersistencetimeout = 30

}
`

func TestAccCsvserver_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_csvserver.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserverExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Csvserver.Type(), "terraform-cs"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCsvserver_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserverExist(resAddr, nil)),
			},
		},
	})
}

func TestAccCsvserver_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserverDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccCsvserver_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserverExist("citrixadc_csvserver.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccCsvserver_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserverExist("citrixadc_csvserver.foo", nil)),
			},
		},
	})
}

const testAccCsvserverDataSource_basic = `

resource "citrixadc_csvserver" "foo" {

  ipv46 = "10.202.11.11"
  name = "terraform-cs"
  port = 8080
  servicetype = "HTTP"

}

data "citrixadc_csvserver" "foo" {
  name = citrixadc_csvserver.foo.name
}
`

func TestAccCsvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "ipv46", "10.202.11.11"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "name", "terraform-cs"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "port", "8080"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "v6persistmasklen", "128"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "persistencetype", "SOURCEIP"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "timeout", "180"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "probesuccessresponsecode", "200-399"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "probeprotocol", "HTTP"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "probeport", "8085"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "persistencebackup", "NONE"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "dtls", "OFF"),
					resource.TestCheckResourceAttr(
						"citrixadc_csvserver.foo", "backuppersistencetimeout", "30"),
				),
			},
		},
	})
}

func TestAccCsvserver_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserverDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_basic},
			{
				Config:                  testAccCsvserver_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccCsvserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.citrixadc_csvserver.foo", "name", "terraform-cs"),
					resource.TestCheckResourceAttr(
						"data.citrixadc_csvserver.foo", "ipv46", "10.202.11.11"),
					resource.TestCheckResourceAttr(
						"data.citrixadc_csvserver.foo", "port", "8080"),
					resource.TestCheckResourceAttr(
						"data.citrixadc_csvserver.foo", "servicetype", "HTTP"),
				),
			},
		},
	})
}

// skipIfDefaultSSLProfileEnabled skips the test unless the ADC's default SSL
// profile is DISABLED. PREREQUISITE: these tests bind ciphers/SNI certs directly
// on the vserver, which NITRO only permits when the default SSL profile is
// DISABLED (a STANDALONE_NON_DEFAULT_SSL_PROFILE box). With it ENABLED, cipher/SNI
// binding must go through an SSL profile and NITRO rejects the direct binding with
// errorcode 3740 "Operation not permitted. Use profile command to do this
// operation". State is read from sslparameter.defaultprofile.
func skipIfDefaultSSLProfileEnabled(t *testing.T) {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		t.Fatalf("Failed to get test client: %v", err)
	}
	data, err := client.FindResource(service.Sslparameter.Type(), "")
	if err != nil {
		t.Fatalf("Failed to read sslparameter to check the default-SSL-profile prerequisite: %v", err)
	}
	if got := strings.TrimSpace(fmt.Sprintf("%v", data["defaultprofile"])); got == "ENABLED" {
		t.Skipf("Prerequisite not met: this test binds ciphers/SNI directly on the vserver, which requires the "+
			"default SSL profile to be DISABLED (STANDALONE_NON_DEFAULT_SSL_PROFILE); appliance reports "+
			"sslparameter.defaultprofile=%q, so NITRO rejects direct binding with errorcode 3740. Skipping.", got)
	}
}

func TestAccCsvserver_standalone_ciphersuites_mixed(t *testing.T) {
	// if isCluster {
	// 	t.Skip("cluster ADC deployment")
	// }
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t); skipIfDefaultSSLProfileEnabled(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Initial
			{
				Config: testCiphersuitesConfig(templateCsvserverCiphersConfig, []string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_csvserver.ciphers", []string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				),
			},
			// Transpose
			{
				Config: testCiphersuitesConfig(templateCsvserverCiphersConfig, []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_csvserver.ciphers", []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}),
				),
			},
			// Empty list
			{
				Config: testCiphersuitesConfig(templateCsvserverCiphersConfig, nil),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_csvserver.ciphers", []string{}),
				),
			},
		},
	})
}

func TestAccCsvserver_cluster_ciphersuites(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t); skipIfDefaultSSLProfileEnabled(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			// Initial
			{
				Config: testCiphersuitesConfig(templateCsvserverCiphersConfig, []string{"SSL3-EXP-ADH-RC4-MD5", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{"SSL3-EXP-ADH-RC4-MD5", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_csvserver.ciphers", []string{"SSL3-EXP-ADH-RC4-MD5", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				),
			},
			// Transpose
			{
				Config: testCiphersuitesConfig(templateCsvserverCiphersConfig, []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "SSL3-EXP-ADH-RC4-MD5"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "SSL3-EXP-ADH-RC4-MD5"}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_csvserver.ciphers", []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "SSL3-EXP-ADH-RC4-MD5"}),
				),
			},
			// Empty list
			{
				Config: testCiphersuitesConfig(templateCsvserverCiphersConfig, nil),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_csvserver.ciphers", []string{}),
				),
			},
		},
	})
}

func TestAccCsvserver_cluster_ciphers(t *testing.T) {
	// if !isCluster {
	// 	t.Skip("standalone ADC deployment")
	// }
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			// Initial
			{
				Config: testCiphersConfig(templateCsvserverCiphersConfig, []string{"HIGH", "MEDIUM"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{"HIGH", "MEDIUM"}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("citrixadc_csvserver.ciphers", []string{"HIGH", "MEDIUM"}),
				),
			},
			// Transpose
			{
				Config: testCiphersConfig(templateCsvserverCiphersConfig, []string{"MEDIUM", "HIGH"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{"MEDIUM", "HIGH"}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("citrixadc_csvserver.ciphers", []string{"MEDIUM", "HIGH"}),
				),
			},
			// Empty list
			{
				Config: testCiphersConfig(templateCsvserverCiphersConfig, nil),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("citrixadc_csvserver.ciphers", []string{}),
				),
			},
		},
	})
}

const templateCsvserverCiphersConfig = `

resource "citrixadc_csvserver" "ciphers" {

  ipv46 = "10.202.11.11"
  name = "tf-acc-ciphers-test"
  port = 443
  servicetype = "SSL"
  %v
}

`

func testAccCheckCsvserverExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := client.FindResource(service.Csvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Csvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCsvserver_AssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	vserverName := "tf-acc-cs-vserver-name"
	vserverType := service.Csvserver.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, vserverType, vserverName, nil)

	vserverInstance := cs.Csvserver{
		Ipv46:       "192.23.23.23",
		Name:        vserverName,
		Servicetype: "HTTP",
		Port:        utils.IntPtr(80),
	}

	if _, err := c.client.AddResource(vserverType, vserverName, vserverInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// Set to zero values all immutables already defined
	vserverInstance.Port = utils.IntPtr(0)
	vserverInstance.Servicetype = ""

	//port
	vserverInstance.Port = utils.IntPtr(88)
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "port")
	vserverInstance.Port = utils.IntPtr(0)

	//servicetype
	vserverInstance.Servicetype = "TCP"
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "servicetype")
	vserverInstance.Servicetype = ""

	//targettype
	vserverInstance.Targettype = "GSLB"
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "targettype")
	vserverInstance.Targettype = ""

	//range
	vserverInstance.Range = utils.IntPtr(1)
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "range")
	vserverInstance.Range = utils.IntPtr(0)

	//td
	vserverInstance.Td = utils.IntPtr(1)
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "td")
	vserverInstance.Td = utils.IntPtr(0)
}

const testAccCsvserverEnableDisable_enabled = `
resource "citrixadc_csvserver" "tf_test_acc_csvserver" {
  name        = "tf_test_acc_csvserver"
  ipv46       = "192.168.33.22"
  port        = 80
  servicetype = "HTTP"
  comment = "enabled state comment"
  state       = "ENABLED"
}
`
const testAccCsvserverEnableDisable_disabled = `
resource "citrixadc_csvserver" "tf_test_acc_csvserver" {
  name        = "tf_test_acc_csvserver"
  ipv46       = "192.168.33.22"
  port        = 80
  servicetype = "HTTP"
  comment = "disabled state comment"
  state       = "DISABLED"
}
`

func TestAccCsvserver_enable_disable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserverDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccCsvserverEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_test_acc_csvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_test_acc_csvserver", "state", "ENABLED"),
				),
			},
			// Disable
			{
				Config: testAccCsvserverEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_test_acc_csvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_test_acc_csvserver", "state", "DISABLED"),
				),
			},
			// Re enable
			{
				Config: testAccCsvserverEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_test_acc_csvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_test_acc_csvserver", "state", "ENABLED"),
				),
			},
		},
	})
}

const testAccCsvserver_binding_add = `
	resource "citrixadc_csvserver" "testbindingfoo" {
		ipv46       = "10.10.10.22"
		name        = "testAccCsVserver"
		port        = 80
		servicetype = "HTTP"

		lbvserverbinding = citrixadc_lbvserver.test_lbvserver_old.name
	}
	resource "citrixadc_lbvserver" "test_lbvserver_old" {
		ipv46       = "10.10.10.33"
		name        = "testAccLbVserver_old"
		port        = 80
		servicetype = "HTTP"
	}
	resource "citrixadc_lbvserver" "test_lbvserver_new" {
		ipv46       = "10.10.10.44"
		name        = "testAccLbVserver_new"
		port        = 80
		servicetype = "HTTP"
	}
`
const testAccCsvserver_binding_update = `
	resource "citrixadc_csvserver" "testbindingfoo" {
		ipv46       = "10.10.10.22"
		name        = "testAccCsVserver"
		port        = 80
		servicetype = "HTTP"

		lbvserverbinding = citrixadc_lbvserver.test_lbvserver_new.name
	}
	resource "citrixadc_lbvserver" "test_lbvserver_old" {
		ipv46       = "10.10.10.33"
		name        = "testAccLbVserver_old"
		port        = 80
		servicetype = "HTTP"
	}
	resource "citrixadc_lbvserver" "test_lbvserver_new" {
		ipv46       = "10.10.10.44"
		name        = "testAccLbVserver_new"
		port        = 80
		servicetype = "HTTP"
	}
`

func TestAccCsvserver_lbvserverbinding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_binding_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserverExist("citrixadc_csvserver.testbindingfoo", nil)),
			},
			{
				Config: testAccCsvserver_binding_update,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserverExist("citrixadc_csvserver.testbindingfoo", nil)),
			},
		},
	})
}

func TestAccCsvserver_snicerts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doPreChecks(t); skipIfDefaultSSLProfileEnabled(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSslcertificateBindingsConfig(sniCertsCsvserverTemplateConfig, "", "cert2-cert3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.cssni", nil),
				),
			},
			{
				Config: testSslcertificateBindingsConfig(sniCertsCsvserverTemplateConfig, "", "cert2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.cssni", nil),
				),
			},
			{
				Config: testSslcertificateBindingsConfig(sniCertsCsvserverTemplateConfig, "cert3", "cert2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.cssni", nil),
				),
			},
			{
				Config: testSslcertificateBindingsConfig(sniCertsCsvserverTemplateConfig, "cert3", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.cssni", nil),
				),
			},
			{
				Config: testSslcertificateBindingsConfig(sniCertsCsvserverTemplateConfig, "cert2", "cert3-cert2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.cssni", nil),
				),
			},
		},
	})
}

const sniCertsCsvserverTemplateConfig = `
	resource "citrixadc_sslcertkey" "cert2" {
	  certkey = "cert2"
	  cert = "/var/tmp/certificate2.crt"
	  key = "/var/tmp/key2.pem"
	  expirymonitor = "DISABLED"
	}

	resource "citrixadc_sslcertkey" "cert3" {
	  certkey = "cert3"
	  cert = "/var/tmp/certificate3.crt"
	  key = "/var/tmp/key3.pem"
	  expirymonitor = "DISABLED"
	}

	resource "citrixadc_csvserver" "cssni" {
	  ipv46 = "10.202.11.11"
	  name = "terraform-cs"
	  port = 443
	  servicetype = "SSL"
	  ciphers = ["DEFAULT"]
	  %v
	  %v
	}
`

func TestAccCsvserver_sslpolicy(t *testing.T) {
	// if isCluster {
	// 	t.Skip("cluster ADC deployment")
	// }
	if isCpxRun {
		t.Skip("TODO fix sslaction for CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: sslpolicy_config_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_csvserver", nil),
				),
			},
			{
				Config: sslpolicy_config_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_csvserver", nil),
				),
			},
			{
				Config: sslpolicy_config_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_csvserver", nil),
				),
			},
			{
				Config: sslpolicy_config_step4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_csvserver", nil),
				),
			},
		},
	})
}

const sslpolicy_config_step1 = `
resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslaction"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_policy"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_sslpolicy" "tf_sslpolicy2" {
  name   = "tf_policy2"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"

  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy.name
     priority = 100
	 gotopriorityexpression = "END"
	}
  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy2.name
     priority = 200
	 gotopriorityexpression = "END"
	}

}
`

const sslpolicy_config_step2 = `
resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslaction"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_policy"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_sslpolicy" "tf_sslpolicy2" {
  name   = "tf_policy2"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"

  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy.name
     priority = 300
	 gotopriorityexpression = "END"
	}
  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy2.name
     priority = 100
	 gotopriorityexpression = "END"
	}

}
`

const sslpolicy_config_step3 = `
resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslaction"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_policy"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_sslpolicy" "tf_sslpolicy2" {
  name   = "tf_policy2"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"

  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy2.name
     priority = 100
	 gotopriorityexpression = "END"
	}

}
`

const sslpolicy_config_step4 = `
resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslaction"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_policy"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_sslpolicy" "tf_sslpolicy2" {
  name   = "tf_policy2"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"

  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy.name
     priority = 300
	 gotopriorityexpression = "END"
	}
  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy2.name
     priority = 100
	 gotopriorityexpression = "END"
	}

}
`

func TestAccCsvserver_sslpolicy_cluster(t *testing.T) {
	// if !isCluster {
	// 	t.Skip("standalone ADC deployment")
	// }
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: sslpolicy_config_cluster_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_csvserver", nil),
				),
			},
			{
				Config: sslpolicy_config_cluster_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_csvserver", nil),
				),
			},
			{
				Config: sslpolicy_config_cluster_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_csvserver", nil),
				),
			},
			{
				Config: sslpolicy_config_cluster_step4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_csvserver", nil),
				),
			},
		},
	})
}

const sslpolicy_config_cluster_step1 = `
resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslaction"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_policy"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_sslpolicy" "tf_sslpolicy2" {
  name   = "tf_policy2"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"

  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy.name
     priority = 100
	 gotopriorityexpression = "END"
	}
  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy2.name
     priority = 200
	 gotopriorityexpression = "END"
	}

}
`

// type = "REQUEST"

const sslpolicy_config_cluster_step2 = `
resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslaction"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_policy"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_sslpolicy" "tf_sslpolicy2" {
  name   = "tf_policy2"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"

  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy.name
     priority = 300
	 gotopriorityexpression = "END"
	}
  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy2.name
     priority = 100
	 gotopriorityexpression = "END"
	}

}
`

const sslpolicy_config_cluster_step3 = `
resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslaction"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_policy"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_sslpolicy" "tf_sslpolicy2" {
  name   = "tf_policy2"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"

  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy2.name
     priority = 100
	 gotopriorityexpression = "END"
	}

}
`

const sslpolicy_config_cluster_step4 = `
resource "citrixadc_sslaction" "tf_sslaction" {
  name                   = "tf_sslaction"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_policy"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_sslpolicy" "tf_sslpolicy2" {
  name   = "tf_policy2"
  rule   = "true"
  action = citrixadc_sslaction.tf_sslaction.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"

  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy.name
     priority = 300
	 gotopriorityexpression = "END"
	}
  sslpolicybinding {
     policyname = citrixadc_sslpolicy.tf_sslpolicy2.name
     priority = 100
	 gotopriorityexpression = "END"
	}

}
`

// The csvserver unset test covers the mutable, spec-unsettable attributes whose
// NITRO ?action=unset cleanly reverts them to their documented server default
// for an HTTP content switching virtual server: appflowlog (ENABLED),
// clttimeout (180), icmpvsrresponse (PASSIVE) and l2conn (OFF). Other listed
// unset attributes (e.g. cacheable, downstateflush, redirectportrewrite,
// disableprimaryondown, rhistate, stateupdate) do NOT revert on unset for a
// csvserver (the appliance keeps the configured value), so they are excluded.
const testAccCsvserver_unset_step1 = `
resource "citrixadc_csvserver" "tf_unset" {
  name            = "tf_test_csvserver_unset"
  ipv46           = "10.202.11.44"
  port            = 8080
  servicetype     = "HTTP"
  appflowlog      = "DISABLED"
  clttimeout      = 200
  icmpvsrresponse = "ACTIVE"
  l2conn          = "ON"
}
`

const testAccCsvserver_unset_step2 = `
resource "citrixadc_csvserver" "tf_unset" {
  name        = "tf_test_csvserver_unset"
  ipv46       = "10.202.11.44"
  port        = 8080
  servicetype = "HTTP"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccCsvserver_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserverDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccCsvserver_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_unset", "appflowlog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_unset", "clttimeout", "200"),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_unset", "icmpvsrresponse", "ACTIVE"),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_unset", "l2conn", "ON"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccCsvserver_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserverExist("citrixadc_csvserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_unset", "appflowlog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_unset", "clttimeout", "180"),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_unset", "icmpvsrresponse", "PASSIVE"),
					resource.TestCheckResourceAttr("citrixadc_csvserver.tf_unset", "l2conn", "OFF"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckCsvserverADCValue("tf_test_csvserver_unset", "appflowlog", "ENABLED"),
					testAccCheckCsvserverADCValue("tf_test_csvserver_unset", "l2conn", "OFF"),
				),
			},
		},
	})
}

// testAccCheckCsvserverADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckCsvserverADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Csvserver.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("csvserver %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("csvserver %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}
