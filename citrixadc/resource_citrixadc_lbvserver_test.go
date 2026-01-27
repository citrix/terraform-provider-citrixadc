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
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLbvserver_basic = `

	resource "citrixadc_lbvserver" "foo" {
	ipv46 = "10.202.11.11"
	lbmethod = "ROUNDROBIN"
	name = "terraform-lb"
	persistencetype = "COOKIEINSERT"
	port = 80
	servicetype = "HTTP"
	}
`

const testAccLbvserver_basic_step2 = `

	resource "citrixadc_lbvserver" "foo" {
	ipv46 = "10.202.11.11"
	lbmethod = "ROUNDROBIN"
	name = "terraform-lb"
	persistencetype = "COOKIEINSERT"
	timeout = 0
	port = 80
	servicetype = "HTTP"
	toggleorder = "ASCENDING"
	probeprotocol = "HTTP"
	probeport = 8085
	probesuccessresponsecode = "200 OK"
	orderthreshold = "10"
	}
`

func TestAccLbvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "ipv46", "10.202.11.11"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "lbmethod", "ROUNDROBIN"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "name", "terraform-lb"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "persistencetype", "COOKIEINSERT"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "port", "80"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "timeout", "2"),
				),
			},
			{
				Config: testAccLbvserver_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "ipv46", "10.202.11.11"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "lbmethod", "ROUNDROBIN"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "name", "terraform-lb"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "persistencetype", "COOKIEINSERT"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "port", "80"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "timeout", "0"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "toggleorder", "ASCENDING"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "probeprotocol", "HTTP"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "probeport", "8085"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "probesuccessresponsecode", "200 OK"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbvserver.foo", "orderthreshold", "10"),
				),
			},
		},
	})
}

const testAccLbvserver_quicbridgeprofile = `
	resource citrixadc_quicbridgeprofile demo_quicbridge {
		name             = "demo_quicbridge"
		routingalgorithm = "PLAINTEXT"
		serveridlength   = 4
	}
	resource "citrixadc_lbvserver" "tfAcc_lbvserver" {
		name                  = "demo_quicbridge_vserver"
		ipv46                 = "10.202.11.11"
		lbmethod              = "TOKEN"
		persistencetype       = "CUSTOMSERVERID"
		rule = "QUIC.CONNECTIONID"
		port                  = 8080
		servicetype           = "QUIC_BRIDGE"
		quicbridgeprofilename = citrixadc_quicbridgeprofile.demo_quicbridge.name
	}
`

func TestAccLbvserver_quicbridgeprofile(t *testing.T) {
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_quicbridgeprofile,
				Check: resource.ComposeTestCheckFunc(testAccCheckLbvserverExist("citrixadc_lbvserver.tfAcc_lbvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.tfAcc_lbvserver", "servicetype", "QUIC_BRIDGE"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.tfAcc_lbvserver", "persistencetype", "CUSTOMSERVERID"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.tfAcc_lbvserver", "quicbridgeprofilename", "demo_quicbridge"),
				),
			},
		},
	})
}

const sniCertsTemplateConfig = `
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

	resource "citrixadc_lbvserver" "lbsni" {
	  ipv46 = "10.202.11.11"
	  lbmethod = "ROUNDROBIN"
	  name = "terraform-lb"
	  persistencetype = "COOKIEINSERT"
	  port = 443
	  servicetype = "SSL"
	  ciphers = ["DEFAULT"]
	  %v
	  %v
	}
`

func TestAccLbvserver_snicerts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { doPreChecks(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSslcertificateBindingsConfig(sniCertsTemplateConfig, "", "cert2-cert3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.lbsni", nil),
				),
			},
			{
				Config: testSslcertificateBindingsConfig(sniCertsTemplateConfig, "", "cert2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.lbsni", nil),
				),
			},
			{
				Config: testSslcertificateBindingsConfig(sniCertsTemplateConfig, "cert3", "cert2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.lbsni", nil),
				),
			},
			{
				Config: testSslcertificateBindingsConfig(sniCertsTemplateConfig, "cert3", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.lbsni", nil),
				),
			},
			{
				Config: testSslcertificateBindingsConfig(sniCertsTemplateConfig, "cert2", "cert3-cert2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.lbsni", nil),
				),
			},
		},
	})
}

const templateCiphersConfig = `

	resource "citrixadc_lbvserver" "ciphers" {

	ipv46 = "10.202.11.11"
	lbmethod = "ROUNDROBIN"
	name = "tf-acc-ciphers-test"
	persistencetype = "COOKIEINSERT"
	port = 443
	servicetype = "SSL"
	%v
	}

`

func TestAccLbvserver_standalone_ciphersuites_mixed(t *testing.T) {
	// if isCluster {
	// 	t.Skip("cluster ADC deployment")
	// }
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			// Initial
			{
				Config: testCiphersuitesConfig(templateCiphersConfig, []string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_lbvserver.ciphers", []string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				),
			},
			// Transpose
			{
				Config: testCiphersuitesConfig(templateCiphersConfig, []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_lbvserver.ciphers", []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}),
				),
			},
			// Empty list
			{
				Config: testCiphersuitesConfig(templateCiphersConfig, nil),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_lbvserver.ciphers", []string{}),
				),
			},
		},
	})
}

func TestAccLbvserver_cluster_ciphersuites(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			// Initial
			{
				Config: testCiphersuitesConfig(templateCiphersConfig, []string{"SSL3-EXP-ADH-RC4-MD5", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{"SSL3-EXP-ADH-RC4-MD5", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_lbvserver.ciphers", []string{"SSL3-EXP-ADH-RC4-MD5", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				),
			},
			// Transpose
			{
				Config: testCiphersuitesConfig(templateCiphersConfig, []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "SSL3-EXP-ADH-RC4-MD5"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "SSL3-EXP-ADH-RC4-MD5"}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_lbvserver.ciphers", []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "SSL3-EXP-ADH-RC4-MD5"}),
				),
			},
			// Empty list
			{
				Config: testCiphersuitesConfig(templateCiphersConfig, nil),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersuitesEqualToActual([]string{}, "tf-acc-ciphers-test"),
					testCheckCiphersuitesConfiguredExpected("citrixadc_lbvserver.ciphers", []string{}),
				),
			},
		},
	})
}

func TestAccLbvserver_cluster_ciphers(t *testing.T) {
	// if !isCluster {
	// 	t.Skip("standalone ADC deployment")
	// }
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			// Initial
			{
				Config: testCiphersConfig(templateCiphersConfig, []string{"HIGH", "MEDIUM"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{"HIGH", "MEDIUM"}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("citrixadc_lbvserver.ciphers", []string{"HIGH", "MEDIUM"}),
				),
			},
			// Transpose
			{
				Config: testCiphersConfig(templateCiphersConfig, []string{"MEDIUM", "HIGH"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{"MEDIUM", "HIGH"}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("citrixadc_lbvserver.ciphers", []string{"MEDIUM", "HIGH"}),
				),
			},
			// Empty list
			{
				Config: testCiphersConfig(templateCiphersConfig, nil),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("citrixadc_lbvserver.ciphers", []string{}),
				),
			},
		},
	})
}

func testAccCheckLbvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		//lintignore:R018
		time.Sleep(5000 * time.Millisecond)
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
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lbvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbvserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func doPreChecks(t *testing.T) {
	testAccPreCheck(t)

	uploads := []string{"certificate2.crt", "key2.pem", "certificate3.crt", "key3.pem"}

	//c := testAccProvider.Meta().(*NetScalerNitroClient)
	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Error instantiating helper NITRO client: %s", err.Error())
	}
	for _, filename := range uploads {
		err := uploadTestdataFile(c, t, filename, "/var/tmp")
		if err != nil {
			t.Errorf("%v", err)
		}
	}

}

func testSslcertificateBindingsConfig(template string, sslcertkey string, snicertskeys string) string {
	sslcertkeyReplacement := ""
	snisslcertkeysReplacement := "snisslcertkeys = []"
	if sslcertkey != "" {
		sslcertkeyReplacement = fmt.Sprintf("sslcertkey = \"${citrixadc_sslcertkey.%v.certkey}\"\n", sslcertkey)
	}
	snicerts := strings.Split(snicertskeys, "-")
	log.Printf("len of snicerts %v", len(snicerts))
	if snicertskeys != "" && len(snicerts) > 0 {
		snisslcertkeysReplacement = "\nsnisslcertkeys = [\n"
		for _, certkey := range snicerts {

			line := fmt.Sprintf("\"${citrixadc_sslcertkey.%v.certkey}\",\n", certkey)
			snisslcertkeysReplacement += line
		}
		snisslcertkeysReplacement += "]\n"
	}
	log.Printf("sslcertkeyReplacement \"%v\"", sslcertkeyReplacement)
	log.Printf("snisslcertkeysReplacement \"%v\"", snisslcertkeysReplacement)
	retval := fmt.Sprintf(template, sslcertkeyReplacement, snisslcertkeysReplacement)
	log.Printf("Full config:\n`\n%s\n`", retval)
	return retval
}

func TestAccLbvserver_AssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	vserverName := "tf-acc-lb-vserver-name"
	vserverType := service.Lbvserver.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, vserverType, vserverName, nil)

	vserverInstance := lb.Lbvserver{
		Ipv46:       "192.23.23.23",
		Name:        vserverName,
		Servicetype: "HTTP",
		Port:        intPtr(80),
	}

	if _, err := c.client.AddResource(vserverType, vserverName, vserverInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// Remove servicetype since it is only valid on create
	vserverInstance.Servicetype = ""

	//port
	vserverInstance.Port = intPtr(80)
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "port")

	vserverInstance = lb.Lbvserver{
		Ipv46: "192.23.23.23",
		Name:  vserverName,
	}

	//servicetype
	vserverInstance.Servicetype = "TCP"
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "servicetype")
	vserverInstance.Servicetype = ""

	vserverInstance = lb.Lbvserver{
		Ipv46: "192.23.23.23",
		Name:  vserverName,
	}

	//range
	vserverInstance.Range = intPtr(10)
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "range")
	vserverInstance.Range = intPtr(0)

	vserverInstance = lb.Lbvserver{
		Ipv46: "192.23.23.23",
		Name:  vserverName,
	}

	//td
	vserverInstance.Td = intPtr(1)
	vserverInstance.Servicetype = ""
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "td")
	vserverInstance.Td = intPtr(0)

	vserverInstance = lb.Lbvserver{
		Ipv46: "192.23.23.23",
		Name:  vserverName,
	}

	//redirurlflags
	vserverInstance.Redirurlflags = true
	testHelperVerifyImmutabilityFunc(c, t, vserverType, vserverName, vserverInstance, "redirurlflags")
	vserverInstance.Redirurlflags = false
}

const testAccLbvserverEnableDisable_enabled = `
resource "citrixadc_lbvserver" "tf_acc_lb_vserver" {
  name = "tf_acc_lb_vserver"
  ipv46 = "192.168.12.67"
  port = "80"
  servicetype = "HTTP"
  comment = "enabled state comment"
  state = "ENABLED"
}
`

const testAccLbvserverEnableDisable_disabled = `
resource "citrixadc_lbvserver" "tf_acc_lb_vserver" {
  name = "tf_acc_lb_vserver"
  ipv46 = "192.168.12.67"
  port = "80"
  servicetype = "HTTP"
  comment = "disabled state comment"
  state = "DISABLED"
}
`

func TestAccLbvserver_enable_disable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccLbvserverEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.tf_acc_lb_vserver", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.tf_acc_lb_vserver", "state", "ENABLED"),
				),
			},
			// Disable
			{
				Config: testAccLbvserverEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.tf_acc_lb_vserver", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.tf_acc_lb_vserver", "state", "DISABLED"),
				),
			},
			// Re enable
			{
				Config: testAccLbvserverEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.tf_acc_lb_vserver", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.tf_acc_lb_vserver", "state", "ENABLED"),
				),
			},
		},
	})
}

const testAccLbvserverNonAddressable = `
resource "citrixadc_lbvserver" "test_non_addressable_lb_vserver" {
	name = "test-na-lb"
	ipv46 = "0.0.0.0"
	port = 0
	servicetype = "HTTP"
}
`

func TestAccLbvserver_nonAddressable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccLbvserverNonAddressable,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.test_non_addressable_lb_vserver", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.test_non_addressable_lb_vserver", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.test_non_addressable_lb_vserver", "ipv46", "0.0.0.0"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.test_non_addressable_lb_vserver", "port", "0"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.test_non_addressable_lb_vserver", "servicetype", "HTTP"),
				),
			},
		},
	})
}

const testAccLbvserver_backupvserver_step1 = `
resource "citrixadc_lbvserver" "tf_test_backup_lb" {
	name = "bkp1"
	ipv46 = "10.202.11.20"
	port = 80
	servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_test_main_lb" {
	name = "lbvs1"
	ipv46 = "10.202.11.21"
	port = 80
	servicetype = "HTTP"
	backupvserver = citrixadc_lbvserver.tf_test_backup_lb.name
}
`

const testAccLbvserver_backupvserver_step2 = `
resource "citrixadc_lbvserver" "tf_test_backup_lb" {
	name = "bkp1"
	ipv46 = "10.202.11.20"
	port = 80
	servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_test_main_lb" {
	name = "lbvs1"
	ipv46 = "10.202.11.21"
	port = 80
	servicetype = "HTTP"
	backupvserver = ""
}
`

const testAccLbvserver_backupvserver_step3 = `

resource "citrixadc_lbvserver" "tf_test_backup_lb" {
	name = "bkp1"
	ipv46 = "10.202.11.20"
	port = 80
	servicetype = "HTTP"
}
	

resource "citrixadc_lbvserver" "tf_test_backup_lb2" {
	name = "bkpnew"
	ipv46 = "10.202.11.22"
	port = 80
	servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_test_main_lb" {
	name = "lbvs1"
	ipv46 = "10.202.11.21"
	port = 80
	servicetype = "HTTP"
	backupvserver = citrixadc_lbvserver.tf_test_backup_lb2.name
}
`

const testAccLbvserver_backupvserver_step4 = `
resource "citrixadc_lbvserver" "tf_test_backup_lb" {
	name = "bkp1"
	ipv46 = "10.202.11.20"
	port = 80
	servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_test_backup_lb2" {
	name = "bkpnew"
	ipv46 = "10.202.11.22"
	port = 80
	servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_test_main_lb" {
	name = "lbvs1"
	ipv46 = "10.202.11.21"
	port = 80
	servicetype = "HTTP"
}
`

func TestAccLbvserver_backupvserver(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			// Create with backupvserver
			{
				Config: testAccLbvserver_backupvserver_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.tf_test_main_lb", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.tf_test_main_lb", "backupvserver", "bkp1"),
				),
			},
			// Unset backupvserver (change from value to empty)
			{
				Config: testAccLbvserver_backupvserver_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.tf_test_main_lb", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.tf_test_main_lb", "backupvserver", ""),
				),
			},
			// Change to different backupvserver
			{
				Config: testAccLbvserver_backupvserver_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.tf_test_main_lb", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.tf_test_main_lb", "backupvserver", "bkpnew"),
				),
			},
			// Remove backupvserver attribute entirely
			{
				Config: testAccLbvserver_backupvserver_step4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("citrixadc_lbvserver.tf_test_main_lb", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver.tf_test_main_lb", "backupvserver", ""),
				),
			},
		},
	})
}
