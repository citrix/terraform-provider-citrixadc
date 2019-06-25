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
	"github.com/chiradeep/go-nitro/config/ssl"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

func TestAccSslcertkey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { doSslcertkeyPreChecks(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslcertkeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslcertkey_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyExist("netscaler_sslcertkey.foo", nil),

					resource.TestCheckResourceAttr(
						"netscaler_sslcertkey.foo", "cert", "/var/tmp/server.crt"),
					resource.TestCheckResourceAttr(
						"netscaler_sslcertkey.foo", "certkey", "sample_ssl_cert"),
					resource.TestCheckResourceAttr(
						"netscaler_sslcertkey.foo", "key", "/var/tmp/server.key"),
				),
			},
		},
	})
}

func testAccCheckSslcertkeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ssl cert name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(netscaler.Sslcertkey.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("SSL cert %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcertkeyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscaler_sslcertkey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Sslcertkey.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("SSL certkey %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func doSslcertkeyPreChecks(t *testing.T) {
	testAccPreCheck(t)

	uploads := []string{"server.crt", "server.key"}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	//c := testAccProvider.Meta().(*NetScalerNitroClient)
	for _, filename := range uploads {
		err := uploadTestdataFile(c, t, filename, "/var/tmp")
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}

const testAccSslcertkey_basic = `


resource "netscaler_sslcertkey" "foo" {
  certkey = "sample_ssl_cert"
  cert = "/var/tmp/server.crt"
  key = "/var/tmp/server.key"
  notificationperiod = 40
  expirymonitor = "ENABLED"
}
`

func TestAccSslcertkeyAssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	uploads := []string{"server.crt", "server.key"}

	for _, filename := range uploads {
		err := uploadTestdataFile(c, t, filename, "/var/tmp")
		if err != nil {
			t.Fatalf(err.Error())
		}
	}

	// Create resource
	certkeyName := "tf-acc-certkey-test"
	certkeyType := netscaler.Sslcertkey.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, certkeyType, certkeyName, nil)

	certkeyInstance := ssl.Sslcertkey{
		Certkey: certkeyName,
		Cert:    "/var/tmp/server.crt",
		Key:     "/var/tmp/server.key",
	}

	if _, err := c.client.AddResource(certkeyType, certkeyName, certkeyInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// Zero out immutable members
	certkeyInstance.Cert = ""
	certkeyInstance.Key = ""

	//cert
	certkeyInstance.Cert = "/var/tmp/new/crt"
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "cert")
	certkeyInstance.Cert = ""

	//key
	certkeyInstance.Key = "/var/tmp/new/key"
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "key")
	certkeyInstance.Key = ""

	//password
	certkeyInstance.Password = true
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "password")
	certkeyInstance.Password = false

	//fipskey
	certkeyInstance.Fipskey = "newfips"
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "fipskey")
	certkeyInstance.Fipskey = ""

	//hsmkey
	certkeyInstance.Hsmkey = "newhsm"
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "hsmkey")
	certkeyInstance.Hsmkey = ""

	//inform
	certkeyInstance.Inform = "PEM"
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "inform")
	certkeyInstance.Inform = ""

	//passplain
	certkeyInstance.Passplain = "passwordnew"
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "passplain")
	certkeyInstance.Passplain = ""

	//bundle
	certkeyInstance.Bundle = "YES"
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "bundle")
	certkeyInstance.Bundle = ""

	//linkcertkeyname
	certkeyInstance.Linkcertkeyname = "certkeyname"
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "linkcertkeyname")
	certkeyInstance.Linkcertkeyname = ""

	//nodomaincheck
	certkeyInstance.Nodomaincheck = true
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "nodomaincheck")
	certkeyInstance.Nodomaincheck = false

	//ocspstaplingcache
	certkeyInstance.Ocspstaplingcache = true
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "ocspstaplingcache")
	certkeyInstance.Ocspstaplingcache = false
}
