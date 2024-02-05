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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strings"
	"testing"
)

const testAccSslservicegroup_sslciphersuite_binding_basic = `

resource "citrixadc_sslservicegroup_sslciphersuite_binding" "tf_sslservicegroup_sslciphersuite_binding" {
	ciphername       = citrixadc_sslcipher.tf_sslcipher.ciphergroupname
	servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
  }
  resource "citrixadc_sslcipher" "tf_sslcipher" {
	  ciphergroupname = "my_ciphersuite"
	 
	  ciphersuitebinding {
		  ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		  cipherpriority = 1
	  }    
  }
  
  resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "my_gslbvservicegroup"
	servicetype      = "SSL_TCP"
	cip              = "DISABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
  }
  resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
  }
  
`

const testAccSslservicegroup_sslciphersuite_binding_basic_step2 = `
resource "citrixadc_sslcipher" "tf_sslcipher" {
	ciphergroupname = "my_ciphersuite"
   
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		cipherpriority = 1
	}    
}

resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
  servicegroupname = "my_gslbvservicegroup"
  servicetype      = "SSL_TCP"
  cip              = "DISABLED"
  healthmonitor    = "NO"
  sitename         = citrixadc_gslbsite.site_local.sitename
}
resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}
`

func TestAccSslservicegroup_sslciphersuite_binding_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslservicegroup_sslciphersuite_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroup_sslciphersuite_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_sslciphersuite_bindingExist("citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding", nil),
				),
			},
			{
				Config: testAccSslservicegroup_sslciphersuite_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_sslciphersuite_bindingNotExist("citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding", "my_gslbvservicegroup,my_ciphersuite"),
				),
			},
		},
	})
}

func testAccCheckSslservicegroup_sslciphersuite_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservicegroup_sslciphersuite_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		servicegroupname := idSlice[0]
		ciphername := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslservicegroup_sslciphersuite_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ciphername
		found := false
		for _, v := range dataArr {
			if v["ciphername"].(string) == ciphername {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslservicegroup_sslciphersuite_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslservicegroup_sslciphersuite_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		servicegroupname := idSlice[0]
		ciphername := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslservicegroup_sslciphersuite_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ciphername
		found := false
		for _, v := range dataArr {
			if v["ciphername"].(string) == ciphername {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslservicegroup_sslciphersuite_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSslservicegroup_sslciphersuite_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservicegroup_sslciphersuite_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Sslservicegroup_sslciphersuite_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslservicegroup_sslciphersuite_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
