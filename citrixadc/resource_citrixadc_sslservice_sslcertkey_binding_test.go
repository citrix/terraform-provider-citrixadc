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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func doSslservice_sslcertkey_bindingPreChecks(t *testing.T) {
	testAccPreCheck(t)

	uploads := []string{
		"ca.crt",
		"certificate1.crt",
		"key1.pem",
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	for _, filename := range uploads {
		err := uploadTestdataFile(c, t, filename, "/var/tmp")
		if err != nil {
			t.Errorf("%v", err)
		}
	}
}

const testAccSslservice_sslcertkey_binding_basic_step1 = `
resource "citrixadc_sslcertkey" "tf_certkey" {
	certkey = "tf_certkey"
	cert = "/nsconfig/ssl/ns-root.cert"
	key = "/nsconfig/ssl/ns-root.key"
	notificationperiod = 40
	expirymonitor = "ENABLED"
}

resource "citrixadc_sslservice" "demo_sslservice" {
	cipherredirect = "DISABLED"
	clientauth = "DISABLED"
	dh = "DISABLED"
	dhcount = 0
	dhkeyexpsizelimit = "DISABLED"
	dtls12 = "DISABLED"
	ersa = "DISABLED"
	redirectportrewrite = "DISABLED"
	serverauth = "ENABLED"
	servicename = citrixadc_service.tf_service.name
	sessreuse = "ENABLED"
	sesstimeout = 300
	snienable = "DISABLED"
	ssl2 = "DISABLED"
	ssl3 = "ENABLED"
	sslredirect = "DISABLED"
	sslv2redirect = "DISABLED"
	tls1 = "ENABLED"
	tls11 = "ENABLED"
	tls12 = "ENABLED"
	tls13 = "DISABLED"
	
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
	depends_on = [citrixadc_sslcertkey.tf_certkey]

}
  
resource "citrixadc_sslservice_sslcertkey_binding" "tf_sslservice_sslcertkey_binding" {
	certkeyname = citrixadc_sslcertkey.tf_certkey.certkey
	servicename = citrixadc_service.tf_service.name
	ca = true
	ocspcheck = "Optional"
}
`

const testAccSslservice_sslcertkey_binding_basic_step2 = `
resource "citrixadc_sslcertkey" "tf_certkey" {
	certkey = "tf_certkey"
	cert = "/nsconfig/ssl/ns-root.cert"
	key = "/nsconfig/ssl/ns-root.key"
	notificationperiod = 40
	expirymonitor = "ENABLED"
}

resource "citrixadc_sslservice" "demo_sslservice" {
	cipherredirect = "DISABLED"
	clientauth = "DISABLED"
	dh = "DISABLED"
	dhcount = 0
	dhkeyexpsizelimit = "DISABLED"
	dtls12 = "DISABLED"
	ersa = "DISABLED"
	redirectportrewrite = "DISABLED"
	serverauth = "ENABLED"
	servicename = citrixadc_service.tf_service.name
	sessreuse = "ENABLED"
	sesstimeout = 300
	snienable = "DISABLED"
	ssl2 = "DISABLED"
	ssl3 = "ENABLED"
	sslredirect = "DISABLED"
	sslv2redirect = "DISABLED"
	tls1 = "ENABLED"
	tls11 = "ENABLED"
	tls12 = "ENABLED"
	tls13 = "DISABLED"
	
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
	depends_on = [citrixadc_sslcertkey.tf_certkey]

}
`

const testAccSslservice_sslcertkey_binding_basic_no_ca_step1 = `
resource "citrixadc_sslcertkey" "tf_certkey" {
	certkey = "tf_certkey"
	cert = "/nsconfig/ssl/ns-root.cert"
	key = "/nsconfig/ssl/ns-root.key"
	notificationperiod = 40
	expirymonitor = "ENABLED"
}

resource "citrixadc_sslservice" "demo_sslservice" {
	cipherredirect = "DISABLED"
	clientauth = "DISABLED"
	dh = "DISABLED"
	dhcount = 0
	dhkeyexpsizelimit = "DISABLED"
	dtls12 = "DISABLED"
	ersa = "DISABLED"
	redirectportrewrite = "DISABLED"
	serverauth = "ENABLED"
	servicename = citrixadc_service.tf_service.name
	sessreuse = "ENABLED"
	sesstimeout = 300
	snienable = "DISABLED"
	ssl2 = "DISABLED"
	ssl3 = "ENABLED"
	sslredirect = "DISABLED"
	sslv2redirect = "DISABLED"
	tls1 = "ENABLED"
	tls11 = "ENABLED"
	tls12 = "ENABLED"
	tls13 = "DISABLED"
	
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
	depends_on = [citrixadc_sslcertkey.tf_certkey]

}
  
resource "citrixadc_sslservice_sslcertkey_binding" "tf_sslservice_sslcertkey_binding" {
	certkeyname = citrixadc_sslcertkey.tf_certkey.certkey
	servicename = citrixadc_service.tf_service.name
}
`

const testAccSslservice_sslcertkey_binding_basic_no_ca_step2 = `
resource "citrixadc_sslcertkey" "tf_certkey" {
	certkey = "tf_certkey"
	cert = "/nsconfig/ssl/ns-root.cert"
	key = "/nsconfig/ssl/ns-root.key"
	notificationperiod = 40
	expirymonitor = "ENABLED"
}

resource "citrixadc_sslservice" "demo_sslservice" {
	cipherredirect = "DISABLED"
	clientauth = "DISABLED"
	dh = "DISABLED"
	dhcount = 0
	dhkeyexpsizelimit = "DISABLED"
	dtls12 = "DISABLED"
	ersa = "DISABLED"
	redirectportrewrite = "DISABLED"
	serverauth = "ENABLED"
	servicename = citrixadc_service.tf_service.name
	sessreuse = "ENABLED"
	sesstimeout = 300
	snienable = "DISABLED"
	ssl2 = "DISABLED"
	ssl3 = "ENABLED"
	sslredirect = "DISABLED"
	sslv2redirect = "DISABLED"
	tls1 = "ENABLED"
	tls11 = "ENABLED"
	tls12 = "ENABLED"
	tls13 = "DISABLED"
	
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
	depends_on = [citrixadc_sslcertkey.tf_certkey]

}
`

func TestAccSslservice_sslcertkey_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { doSslservice_sslcertkey_bindingPreChecks(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckSslservice_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservice_sslcertkey_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservice_sslcertkey_bindingExist("citrixadc_sslservice_sslcertkey_binding.tf_sslservice_sslcertkey_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslcertkey_binding.tf_sslservice_sslcertkey_binding", "ca", "true"),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslcertkey_binding.tf_sslservice_sslcertkey_binding", "ocspcheck", "Optional"),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslcertkey_binding.tf_sslservice_sslcertkey_binding", "snicert", "false"),
				),
			},
			{
				Config: testAccSslservice_sslcertkey_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservice_sslcertkey_bindingNotExist("citrixadc_sslservice_sslcertkey_binding.tf_sslservice_sslcertkey_binding", "tf_service,tf_certkey,false,true"),
				),
			},
		},
	})
}

func TestAccSslservice_sslcertkey_binding_basic_no_ca(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { doSslservice_sslcertkey_bindingPreChecks(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckSslservice_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservice_sslcertkey_binding_basic_no_ca_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservice_sslcertkey_bindingExist("citrixadc_sslservice_sslcertkey_binding.tf_sslservice_sslcertkey_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslcertkey_binding.tf_sslservice_sslcertkey_binding", "ca", "false"),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslcertkey_binding.tf_sslservice_sslcertkey_binding", "snicert", "false"),
				),
			},
			{
				Config: testAccSslservice_sslcertkey_binding_basic_no_ca_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservice_sslcertkey_bindingNotExist("citrixadc_sslservice_sslcertkey_binding.tf_sslservice_sslcertkey_binding", "tf_service,tf_certkey,false,false"),
				),
			},
		},
	})
}

func testAccCheckSslservice_sslcertkey_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservice_sslcertkey_binding id is set")
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

		bindingId := rs.Primary.ID

		idSlice := strings.Split(bindingId, ",")

		servicename := idSlice[0]
		certkeyname := idSlice[1]
		snicert := false
		ca := false
		if val, ok := rs.Primary.Attributes["snicert"]; ok {
			snicert = val == "true"
		}
		if val, ok := rs.Primary.Attributes["ca"]; ok {
			ca = val == "true"
		}

		findParams := service.FindParams{
			ResourceType:             "sslservice_sslcertkey_binding",
			ResourceName:             servicename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["certkeyname"].(string) == certkeyname && v["snicert"].(bool) == snicert && v["ca"].(bool) == ca {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslservice_sslcertkey_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslservice_sslcertkey_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}

		idSlice := strings.Split(id, ",")
		servicename := idSlice[0]
		certkeyname := idSlice[1]
		snicert := idSlice[2] == "true"
		ca := idSlice[3] == "true"

		findParams := service.FindParams{
			ResourceType:             "sslservice_sslcertkey_binding",
			ResourceName:             servicename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["certkeyname"].(string) == certkeyname && v["snicert"].(bool) == snicert && v["ca"].(bool) == ca {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslservice_sslcertkey_binding %s not deleted", n)
		}

		return nil
	}
}

func testAccCheckSslservice_sslcertkey_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservice_sslcertkey_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslservice_sslcertkey_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslservice_sslcertkey_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
