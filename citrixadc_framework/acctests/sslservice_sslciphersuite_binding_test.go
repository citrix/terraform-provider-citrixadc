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

const testAccSslservice_sslciphersuite_binding_basic_step1 = `
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

}

resource "citrixadc_sslcipher" "tfAccsslcipher" {
	ciphergroupname = "tfAccsslcipher"

	# ciphersuitebinding is MANDATORY attribute
	# Any change in the ciphersuitebinding will result in re-creation of the whole sslcipher resource.
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		cipherpriority = 1
	}
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
		cipherpriority = 2
	}
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES-128-SHA256"
		cipherpriority = 3
	}
}

resource "citrixadc_sslservice_sslciphersuite_binding" "tf_sslservice_sslciphersuite_binding" {
	ciphername = citrixadc_sslcipher.tfAccsslcipher.ciphergroupname
	servicename = citrixadc_service.tf_service.name   
}
`

const testAccSslservice_sslciphersuite_binding_basic_step2 = `
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

}

resource "citrixadc_sslcipher" "tfAccsslcipher" {
	ciphergroupname = "tfAccsslcipher"

	# ciphersuitebinding is MANDATORY attribute
	# Any change in the ciphersuitebinding will result in re-creation of the whole sslcipher resource.
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		cipherpriority = 1
	}
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
		cipherpriority = 2
	}
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES-128-SHA256"
		cipherpriority = 3
	}
}
`

func TestAccSslservice_sslciphersuite_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslservice_sslciphersuite_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservice_sslciphersuite_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservice_sslciphersuite_bindingExist("citrixadc_sslservice_sslciphersuite_binding.tf_sslservice_sslciphersuite_binding", nil),
				),
			},
			{
				Config: testAccSslservice_sslciphersuite_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservice_sslciphersuite_bindingNotExist("citrixadc_sslservice_sslciphersuite_binding.tf_sslservice_sslciphersuite_binding", "tf_service,tfAccsslcipher"),
				),
			},
		},
	})
}

func testAccCheckSslservice_sslciphersuite_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservice_sslciphersuite_binding id is set")
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

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		servicename := idSlice[0]
		ciphername := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslservice_sslciphersuite_binding",
			ResourceName:             servicename,
			ResourceMissingErrorCode: 463,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["ciphername"].(string) == ciphername {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslservice_sslciphersuite_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslservice_sslciphersuite_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		servicename := idSlice[0]
		ciphername := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslservice_sslcipher_binding",
			ResourceName:             servicename,
			ResourceMissingErrorCode: 463,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		found := false
		for _, v := range dataArr {
			if v["ciphername"].(string) == ciphername {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslservice_sslcipher_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSslservice_sslciphersuite_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservice_sslciphersuite_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslservice_sslciphersuite_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslservice_sslciphersuite_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
