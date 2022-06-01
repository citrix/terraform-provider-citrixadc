/*
Copyright 2022 Citrix Systems, Inc

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
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testAccSslvserverBindings_basic_step1 = `
resource "citrixadc_lbvserver" "tf_sslvserver" {
	name = "tf_sslvserver"
	servicetype = "SSL"
	ipv46 = "5.5.5.5"
	port = 443
}

resource "citrixadc_sslcipher" "tfsslcipher" {
  ciphergroupname = "tfsslcipher"

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

resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding" {
	ciphername = citrixadc_sslcipher.tfsslcipher.ciphergroupname
	vservername = citrixadc_lbvserver.tf_sslvserver.name
}

data "citrixadc_sslcipher_sslvserver_bindings" "sslbindings" {
    ciphername = citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding.ciphername
}
`

func TestAccSslcipherSslvserverBindings_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslvserverBindings_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslcipher_sslvserver_bindings.sslbindings", "bound_sslvservers", "tf_sslvserver"),
				),
			},
		},
	})
}
