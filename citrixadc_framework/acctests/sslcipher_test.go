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
	"strconv"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccSslcipher_add = `
	resource "citrixadc_sslcipher" "foo" {
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

const testAccSslcipher_transpose = `
	resource "citrixadc_sslcipher" "foo" {
		ciphergroupname = "tfAccsslcipher"

		# ciphersuitebinding is MANDATORY attribute
		# Any change in the ciphersuitebinding will result in re-creation of the whole sslcipher resource.
		ciphersuitebinding {
			ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
			cipherpriority = 3
	}
		ciphersuitebinding {
			ciphername     = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
			cipherpriority = 2
	}
		ciphersuitebinding {
			ciphername     = "TLS1.2-ECDHE-RSA-AES-128-SHA256"
			cipherpriority = 1
	}
	}
`

// Update re-creates the while ciphergroup
const testAccSslcipher_update = `  
	resource "citrixadc_sslcipher" "foo" {
		ciphergroupname = "tfAccsslcipher"

		# ciphersuitebinding is MANDATORY attribute
		# Any change in the ciphersuitebinding will result in re-creation of the whole sslcipher resource.
		ciphersuitebinding {
			ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
			cipherpriority = 1
	}
	}
`

func TestAccSslcipher_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcipherDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcipher_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcipherExist("citrixadc_sslcipher.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcipher.foo", "ciphergroupname", "tfAccsslcipher"),
					testAccCheckSslcipherCiphersuiteBinding("tfAccsslcipher", "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256", 1),
					testAccCheckSslcipherCiphersuiteBinding("tfAccsslcipher", "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384", 2),
					testAccCheckSslcipherCiphersuiteBinding("tfAccsslcipher", "TLS1.2-ECDHE-RSA-AES-128-SHA256", 3),
				),
			},
			{
				Config: testAccSslcipher_transpose,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcipherExist("citrixadc_sslcipher.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcipher.foo", "ciphergroupname", "tfAccsslcipher"),
					testAccCheckSslcipherCiphersuiteBinding("tfAccsslcipher", "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256", 3),
					testAccCheckSslcipherCiphersuiteBinding("tfAccsslcipher", "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384", 2),
					testAccCheckSslcipherCiphersuiteBinding("tfAccsslcipher", "TLS1.2-ECDHE-RSA-AES-128-SHA256", 1),
				),
			},
			{
				Config: testAccSslcipher_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcipherExist("citrixadc_sslcipher.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcipher.foo", "ciphergroupname", "tfAccsslcipher"),
					testAccCheckSslcipherCiphersuiteBinding("tfAccsslcipher", "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256", 1),
				),
			},
		},
	})
}

func testAccCheckSslcipherExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("sslciphergroupname is set")
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
		dataArr, err := client.FindAllResources(service.Sslcipher.Type())

		if err != nil {
			return err
		}
		sslcipherGroupName := rs.Primary.ID

		if len(dataArr) == 0 {
			fmt.Printf("[WARN] citrixadc-provider: Sslcipher does not exist.")
			return fmt.Errorf("SSLCiphergroup %s not found", n)
		}

		foundIndex := -1
		for i, v := range dataArr {
			if v["ciphergroupname"].(string) == sslcipherGroupName {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			fmt.Printf("[WARN] citrixadc-provider: Sslcipher does not exist in the dataArr.")
			return fmt.Errorf("SSLCiphergroup %s not found", n)
		}

		data := dataArr[foundIndex]

		if data == nil {
			return fmt.Errorf("SSLCiphergroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcipherCiphersuiteBinding(ciphergroupname string, ciphername string, expectedpriority int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "citrixadc_sslcipher" {
				continue
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("No name is set")
			}

			bindings, _ := client.FindResourceArray(service.Sslcipher_sslciphersuite_binding.Type(), ciphergroupname)
			for _, binding := range bindings {
				if binding["ciphername"].(string) == ciphername {
					receivedpriority, _ := strconv.Atoi(binding["cipherpriority"].(string))
					if receivedpriority != expectedpriority {
						return fmt.Errorf("Expected cipherpriority %d, got %d for ciphername %s in ciphergroup %s", expectedpriority, receivedpriority, ciphername, ciphergroupname)
					} else {
						return nil
					}
				}
			}
		}

		return fmt.Errorf("ciphername %s not found for ciphergroupname %s", ciphername, ciphergroupname)
	}
}

func testAccCheckSslcipherDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcipher" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslcipher.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslciphergroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
