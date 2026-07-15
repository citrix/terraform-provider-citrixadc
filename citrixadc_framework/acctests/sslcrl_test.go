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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccSslcrl_basic = `
	resource "citrixadc_sslcrl" "tf_sslcrl" {
		crlname = "tf_sslcrl"
		crlpath = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert  = "rootrsa_cert1"
	}
`

const testAccSslcrlDataSource_basic = `
	resource "citrixadc_sslcrl" "tf_sslcrl" {
		crlname = "tf_sslcrl"
		crlpath = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert  = "rootrsa_cert1"
	}

	data "citrixadc_sslcrl" "tf_sslcrl" {
		crlname = citrixadc_sslcrl.tf_sslcrl.crlname
	}
`

func TestAccSslcrl_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcrlPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcrlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrl_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlExist("citrixadc_sslcrl.tf_sslcrl", nil),
				),
			},
		},
	})
}

// doSslcrlPreChecks stages the prerequisites the basic sslcrl config depends on.
// `add ssl crl` requires (a) the CRL file to be present under the partition's
// default CRL directory (/var/netscaler/ssl) and (b) an installed CA certificate
// (certkey) whose subject matches the CRL issuer, or NITRO rejects the CRL with
// errorcode 1583 "Unable to find the CA certificate for the CRL".
//
// Rather than depend on a fixed testdata CRL/CA pair, this precheck generates a
// fresh self-signed CA and a matching (empty) CRL in-process, uploads the CA cert
// to /nsconfig/ssl and creates certkey "rootrsa_cert1" from it, then uploads the
// CRL to /var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem. All three artifacts
// are cleaned up after the test so the appliance is left clean.
func doSslcrlPreChecks(t *testing.T) {
	testAccPreCheck(t)

	const (
		certkeyName  = "rootrsa_cert1"
		caFileName   = "rootrsa_cert1.pem"
		caFileDir    = "/nsconfig/ssl"
		crlFileName  = "crl_config_clnt_rsa1_1cert.pem"
		crlFileDir   = "/var/netscaler/ssl"
		crlFileDeEnc = "%2Fvar%2Fnetscaler%2Fssl"
	)

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v", err)
	}
	client := c.client

	// 1. Generate a self-signed CA (used to sign, and to validate, the CRL).
	caKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate CA key: %v", err)
	}
	caSubject := pkix.Name{Country: []string{"in"}, Organization: []string{"citrix"}, CommonName: "crlca"}
	caTemplate := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               caSubject,
		NotBefore:             time.Now().Add(-1 * time.Hour),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}
	caDER, err := x509.CreateCertificate(rand.Reader, caTemplate, caTemplate, &caKey.PublicKey, caKey)
	if err != nil {
		t.Fatalf("Failed to create CA certificate: %v", err)
	}
	caCert, err := x509.ParseCertificate(caDER)
	if err != nil {
		t.Fatalf("Failed to parse CA certificate: %v", err)
	}
	caCertPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: caDER})

	// 2. Generate an (empty) CRL signed by that CA - its issuer therefore matches
	//    the CA certificate's subject.
	crlTemplate := &x509.RevocationList{
		Number:     big.NewInt(1),
		ThisUpdate: time.Now().Add(-1 * time.Hour),
		NextUpdate: time.Now().AddDate(10, 0, 0),
	}
	crlDER, err := x509.CreateRevocationList(rand.Reader, crlTemplate, caCert, caKey)
	if err != nil {
		t.Fatalf("Failed to create CRL: %v", err)
	}
	crlPEM := pem.EncodeToMemory(&pem.Block{Type: "X509 CRL", Bytes: crlDER})

	// 3. Upload the CA cert and create the certkey referenced by cacert.
	if err := uploadSystemfileContent(client, caFileName, caFileDir, caCertPEM); err != nil {
		t.Fatalf("Failed to upload CA cert file: %v", err)
	}
	// Recreate the certkey so it reads the freshly-generated CA cert.
	_ = client.DeleteResource(service.Sslcertkey.Type(), certkeyName)
	certkey := ssl.Sslcertkey{Certkey: certkeyName, Cert: caFileName}
	if _, err := client.AddResource(service.Sslcertkey.Type(), certkeyName, &certkey); err != nil {
		t.Fatalf("Failed to add certkey %s: %v", certkeyName, err)
	}

	// 4. Upload the CRL file into the default CRL directory.
	if err := uploadSystemfileContent(client, crlFileName, crlFileDir, crlPEM); err != nil {
		t.Fatalf("Failed to upload CRL file: %v", err)
	}

	// 5. Leave the appliance clean once the test completes.
	t.Cleanup(func() {
		_ = client.DeleteResource(service.Sslcertkey.Type(), certkeyName)
		_ = client.DeleteResourceWithArgsMap(service.Systemfile.Type(), caFileName,
			map[string]string{"filelocation": strings.Replace(caFileDir, "/", "%2F", -1)})
		_ = client.DeleteResourceWithArgsMap(service.Systemfile.Type(), crlFileName,
			map[string]string{"filelocation": crlFileDeEnc})
	})
}

// uploadSystemfileContent writes the given bytes to a systemfile at targetDir,
// overwriting any existing file of the same name.
func uploadSystemfileContent(client *service.NitroClient, filename, targetDir string, content []byte) error {
	sf := system.Systemfile{
		Filename:     filename,
		Filecontent:  base64.StdEncoding.EncodeToString(content),
		Filelocation: targetDir,
	}
	_, err := client.AddResource(service.Systemfile.Type(), filename, &sf)
	if err != nil && strings.Contains(err.Error(), "File already exists") {
		urlArgs := map[string]string{"filelocation": strings.Replace(targetDir, "/", "%2F", -1)}
		if derr := client.DeleteResourceWithArgsMap(service.Systemfile.Type(), filename, urlArgs); derr != nil {
			return derr
		}
		_, err = client.AddResource(service.Systemfile.Type(), filename, &sf)
	}
	return err
}

func TestAccSslcrl_import(t *testing.T) {
	const resAddr = "citrixadc_sslcrl.tf_sslcrl"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcrlPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcrlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrl_basic,
			},
			{
				Config:            testAccSslcrl_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// password_wo_version is a Computed write-only version tracker that
				// NITRO does not return; on import it cannot be repopulated from the
				// API, so it legitimately cannot round-trip.
				ImportStateVerifyIgnore: []string{"password_wo_version"},
			},
		},
	})
}

func testAccCheckSslcrlExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcrl name is set")
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
		data, err := client.FindResource(service.Sslcrl.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslcrl %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcrlDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcrl" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslcrl.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslcrl %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslcrl_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcrlPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrlDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslcrl.tf_sslcrl", "crlname", "tf_sslcrl"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcrl.tf_sslcrl", "crlpath", "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcrl.tf_sslcrl", "cacert", "rootrsa_cert1"),
				),
			},
		},
	})
}

// Test backward-compatible path: using password (Sensitive attribute)
// password is the LDAP password used when refreshing CRL from an LDAP server
const testAccSslcrl_password_step1 = `
	variable "sslcrl_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcrl" "tf_sslcrl_ephem" {
		crlname  = "tf_sslcrl_ephem"
		crlpath  = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert   = "rootrsa_cert1"
		password = var.sslcrl_password
	}
`

func TestAccSslcrl_password_backward_compat(t *testing.T) {
	t.Skipf("Need a valid CRL file on the ADC instance before running this test")
	t.Setenv("TF_VAR_sslcrl_password", "crlldappass1")
	t.Setenv("TF_VAR_sslcrl_password_2", "crlldappass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcrlPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcrlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrl_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlExist("citrixadc_sslcrl.tf_sslcrl_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcrl.tf_sslcrl_ephem", "crlname", "tf_sslcrl_ephem"),
				),
			},
		},
	})
}

func TestAccSslcrl_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { doSslcrlPreChecks(t) },
		CheckDestroy: testAccCheckSslcrlDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSslcrl_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlExist("citrixadc_sslcrl.tf_sslcrl", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSslcrl_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlExist("citrixadc_sslcrl.tf_sslcrl", nil),
				),
			},
		},
	})
}

// Test ephemeral path: using password_wo (WriteOnly attribute) with version tracker
const testAccSslcrl_password_wo_step1 = `
	variable "sslcrl_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcrl" "tf_sslcrl_ephem" {
		crlname             = "tf_sslcrl_ephem"
		crlpath             = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert              = "rootrsa_cert1"
		password_wo         = var.sslcrl_password_wo
		password_wo_version = 1
	}
`

func TestAccSslcrl_password_wo_ephemeral(t *testing.T) {
	t.Skipf("Need a valid CRL file on the ADC instance before running this test")
	t.Setenv("TF_VAR_sslcrl_password_wo", "ephem_crlpass1")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcrlPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcrlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrl_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlExist("citrixadc_sslcrl.tf_sslcrl_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcrl.tf_sslcrl_ephem", "crlname", "tf_sslcrl_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslcrl.tf_sslcrl_ephem", "password_wo_version", "1"),
				),
			},
		},
	})
}
