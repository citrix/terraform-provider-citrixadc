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
package netscaler

import (
	"fmt"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"strings"
	"testing"
	"time"
)

func TestAccLbvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLbvserver_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("netscaler_lbvserver.foo", nil),

					resource.TestCheckResourceAttr(
						"netscaler_lbvserver.foo", "ipv46", "10.202.11.11"),
					resource.TestCheckResourceAttr(
						"netscaler_lbvserver.foo", "lbmethod", "ROUNDROBIN"),
					resource.TestCheckResourceAttr(
						"netscaler_lbvserver.foo", "name", "terraform-lb"),
					resource.TestCheckResourceAttr(
						"netscaler_lbvserver.foo", "persistencetype", "COOKIEINSERT"),
					resource.TestCheckResourceAttr(
						"netscaler_lbvserver.foo", "port", "80"),
					resource.TestCheckResourceAttr(
						"netscaler_lbvserver.foo", "servicetype", "HTTP"),
				),
			},
		},
	})
}

func TestAccLbvserver_snicerts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { doPreChecks(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testSslcertificateBindingsConfig("", "cert2-cert3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("netscaler_lbvserver.lbsni", nil),
				),
			},
			resource.TestStep{
				Config: testSslcertificateBindingsConfig("", "cert2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("netscaler_lbvserver.lbsni", nil),
				),
			},
			resource.TestStep{
				Config: testSslcertificateBindingsConfig("cert3", "cert2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("netscaler_lbvserver.lbsni", nil),
				),
			},
			resource.TestStep{
				Config: testSslcertificateBindingsConfig("cert3", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("netscaler_lbvserver.lbsni", nil),
				),
			},
			resource.TestStep{
				Config: testSslcertificateBindingsConfig("cert2", "cert3-cert2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserverExist("netscaler_lbvserver.lbsni", nil),
				),
			},
		},
	})
}

func TestAccLbvserver_ciphers(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLbvserverDestroy,
		Steps: []resource.TestStep{
			// Initial
			resource.TestStep{
				Config: testCiphersConfig(templateCiphersConfig, []string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("netscaler_lbvserver.ciphers", []string{"HIGH", "TLS1.2-DHE-RSA-CHACHA20-POLY1305"}),
				),
			},
			// Transpose
			resource.TestStep{
				Config: testCiphersConfig(templateCiphersConfig, []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("netscaler_lbvserver.ciphers", []string{"TLS1.2-DHE-RSA-CHACHA20-POLY1305", "HIGH"}),
				),
			},
			// Empty list
			resource.TestStep{
				Config: testCiphersConfig(templateCiphersConfig, nil),
				Check: resource.ComposeTestCheckFunc(
					testCheckCiphersEqualToActual([]string{}, "tf-acc-ciphers-test"),
					testCheckCiphersConfiguredExpected("netscaler_lbvserver.ciphers", []string{}),
				),
			},
		},
	})
}

func testAccCheckLbvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
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

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(netscaler.Lbvserver.Type(), rs.Primary.ID)

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
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscaler_lbvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Lbvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func doPreChecks(t *testing.T) {
	testAccPreCheck(t)

	uploads := []string{"certificate2.pem", "key2.pem", "certificate3.pem", "key3.pem"}

	for _, filename := range uploads {
		err := uploadTestdataFile(t, filename, "/var/tmp")
		if err != nil {
			t.Errorf(err.Error())
		}
	}

}

const templateCiphersConfig = `

resource "netscaler_lbvserver" "ciphers" {

  ipv46 = "10.202.11.11"
  lbmethod = "ROUNDROBIN"
  name = "tf-acc-ciphers-test"
  persistencetype = "COOKIEINSERT"
  port = 443
  servicetype = "SSL"
  %v
}

`

const testAccLbvserver_basic = `

resource "netscaler_lbvserver" "foo" {
  
  ipv46 = "10.202.11.11"
  lbmethod = "ROUNDROBIN"
  name = "terraform-lb"
  persistencetype = "COOKIEINSERT"
  port = 80
  servicetype = "HTTP"
}
`

func testSslcertificateBindingsConfig(sslcertkey string, snicertskeys string) string {
	sslcertkeyReplacement := ""
	snisslcertkeysReplacement := "snisslcertkeys = []"
	if sslcertkey != "" {
		sslcertkeyReplacement = fmt.Sprintf("sslcertkey = \"${netscaler_sslcertkey.%v.certkey}\"\n", sslcertkey)
	}
	snicerts := strings.Split(snicertskeys, "-")
	log.Printf("len of snicerts %v", len(snicerts))
	if snicertskeys != "" && len(snicerts) > 0 {
		snisslcertkeysReplacement = "\nsnisslcertkeys = [\n"
		for _, certkey := range snicerts {

			line := fmt.Sprintf("\"${netscaler_sslcertkey.%v.certkey}\",\n", certkey)
			snisslcertkeysReplacement += line
		}
		snisslcertkeysReplacement += "]\n"
	}
	log.Printf("sslcertkeyReplacement \"%v\"", sslcertkeyReplacement)
	log.Printf("snisslcertkeysReplacement \"%v\"", snisslcertkeysReplacement)
	return fmt.Sprintf(`
	resource "netscaler_sslcertkey" "cert2" {
	  certkey = "cert2"
	  cert = "/var/tmp/certificate2.pem"
	  key = "/var/tmp/key2.pem"
	  expirymonitor = "DISABLED"
	}

	resource "netscaler_sslcertkey" "cert3" {
	  certkey = "cert3"
	  cert = "/var/tmp/certificate3.pem"
	  key = "/var/tmp/key3.pem"
	  expirymonitor = "DISABLED"
	}

	resource "netscaler_lbvserver" "lbsni" {
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
	`, sslcertkeyReplacement, snisslcertkeysReplacement)
}
