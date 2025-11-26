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
package acctests

import (
	"fmt"
	"os"
	"testing"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccSslcertkey_basic = `

resource "citrixadc_sslcertkey" "foo" {
  certkey = "sample_ssl_cert"
  cert = "/nsconfig/ssl/servercert1.cert"
  key = "/nsconfig/ssl/servercert1.key"
  notificationperiod = 40
  expirymonitor = "ENABLED"
}
`

const testAccSslcertkey_basic_update = `

variable "sslcertkey_passplain_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslcertkey" "foo" {
  certkey = "sample_ssl_cert"
  cert = "/nsconfig/ssl/servercert2.cert"
  key = "/nsconfig/ssl/servercert2.key"
  nodomaincheck = true
  password = true
  passplain_wo  = var.sslcertkey_passplain_wo
  passplain_wo_version = 2
}
`

const testAccSslcertkey_basic_update_2 = `

variable "sslcertkey_passplain_wo_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslcertkey" "foo" {
  certkey = "sample_ssl_cert"
  cert = "/nsconfig/ssl/servercert3.cert"
  key = "/nsconfig/ssl/servercert3.key"
  nodomaincheck = true
  password = true
  passplain_wo  = var.sslcertkey_passplain_wo_2
  passplain_wo_version = 3
}
`

func TestAccSslcertkey_basic(t *testing.T) {
	// t.Skip("TODO: Need to find a way to test this resource!")
	t.Setenv("TF_VAR_sslcertkey_passplain_wo", "123456")
	t.Setenv("TF_VAR_sslcertkey_passplain_wo_2", "1234567")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertkeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertkey_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.foo", "cert", "/nsconfig/ssl/servercert1.cert"),
					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.foo", "certkey", "sample_ssl_cert"),
					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.foo", "key", "/nsconfig/ssl/servercert1.key"),
				),
			},
			{
				Config: testAccSslcertkey_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.foo", "cert", "/nsconfig/ssl/servercert2.cert"),
					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.foo", "certkey", "sample_ssl_cert"),
					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.foo", "key", "/nsconfig/ssl/servercert2.key"),
				),
			},
			{
				Config: testAccSslcertkey_basic_update_2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.foo", "cert", "/nsconfig/ssl/servercert3.cert"),
					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.foo", "certkey", "sample_ssl_cert"),
					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.foo", "key", "/nsconfig/ssl/servercert3.key"),
				),
			},
		},
	})
}

const testAccSslcertkey_linkcert_nolink = `

resource "citrixadc_sslcertkey" "client" {
    cert = "/nsconfig/ssl/servercert1.cert"
    key = "/nsconfig/ssl/servercert1.key"
    certkey = "client"
}

resource "citrixadc_sslcertkey" "intermediate" {
    cert = "/nsconfig/ssl/intermediate1.cert"
    certkey = "intermediate"
}

`

// TODO Add use case with cross signed certificate to do a link-unlink operation in one pass
const testAccSslcertkey_linkcert_linked = `

resource "citrixadc_sslcertkey" "client" {
    cert = "/nsconfig/ssl/servercert1.cert"
    key = "/nsconfig/ssl/servercert1.key"
    certkey = "client"
    linkcertkeyname = citrixadc_sslcertkey.intermediate.certkey
}

resource "citrixadc_sslcertkey" "intermediate" {
    cert = "/nsconfig/ssl/intermediate1.cert"
    certkey = "intermediate"
}

`

const testAccSslcertkey_linkcert_client_key_removed = `

resource "citrixadc_sslcertkey" "intermediate" {
    cert = "/nsconfig/ssl/intermediate1.cert"
    certkey = "intermediate"
}

`

func TestAccSslcertkey_linkcert(t *testing.T) {
	// t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertkeyDestroy,
		Steps: []resource.TestStep{

			// Check initial link
			{
				Config: testAccSslcertkey_linkcert_linked,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.client", nil),
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.intermediate", nil),
					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.client", "linkcertkeyname", "intermediate"),
				),
			},

			// Check unlink
			{
				Config: testAccSslcertkey_linkcert_nolink,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.client", nil),
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.intermediate", nil),
					// resource.TestCheckResourceAttr("citrixadc_sslcertkey.client", "linkcertkeyname", ""),
				),
			},

			// Check relink
			{
				Config: testAccSslcertkey_linkcert_linked,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.client", nil),
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.intermediate", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.client", "linkcertkeyname", "intermediate"),
				),
			},

			// Check removal of linked key
			{
				Config: testAccSslcertkey_linkcert_client_key_removed,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.intermediate", nil),
				),
			},

			// Recreate unlinked to check subsequent removal
			{
				Config: testAccSslcertkey_linkcert_nolink,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.client", nil),
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.intermediate", nil),

					// resource.TestCheckResourceAttr("citrixadc_sslcertkey.client", "linkcertkeyname", ""),
				),
			},

			// Check removal of unlinked key
			{
				Config: testAccSslcertkey_linkcert_client_key_removed,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.intermediate", nil),
				),
			},

			// Relink to test removal of both entries by end of test
			{
				Config: testAccSslcertkey_linkcert_linked,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.client", nil),
					testAccCheckSslcertkeyExist("citrixadc_sslcertkey.intermediate", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_sslcertkey.client", "linkcertkeyname", "intermediate"),
				),
			},
		},
	})
}

func TestAccSslcertkey_AssertNonUpdateableAttributes(t *testing.T) {
	// t.Skip("TODO: Need to find a way to test this resource!")

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	certkeyName := "tf-acc-certkey-test"
	certkeyType := service.Sslcertkey.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, certkeyType, certkeyName, nil)

	certkeyInstance := ssl.Sslcertkey{
		Certkey: certkeyName,
		Cert:    "/nsconfig/ssl/servercert1.cert",
		Key:     "/nsconfig/ssl/servercert1.key",
	}

	if _, err := c.client.AddResource(certkeyType, certkeyName, certkeyInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// Zero out immutable members
	certkeyInstance.Cert = ""
	certkeyInstance.Key = ""

	//cert
	certkeyInstance.Cert = "/nsconfig/ssl/new/cert"
	testHelperVerifyImmutabilityFunc(c, t, certkeyType, certkeyName, certkeyInstance, "cert")
	certkeyInstance.Cert = ""

	//key
	certkeyInstance.Key = "/nsconfig/ssl/new/key"
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Sslcertkey.Type(), rs.Primary.ID)

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcertkey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslcertkey.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("SSL certkey %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
