/*
Copyright 2020 Citrix Systems, Inc

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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// The nitro_info data source is a generic NITRO query escape-hatch. This exercises
// the "binding_list" workflow: after binding an sslcertkey to an SSL vserver, query
// sslcertkey_sslvserver_binding for the certkey and expect exactly one entry.
// The data source depends_on the binding so it is read AFTER the binding is created.
const testAccNitroInfoBindingList_step1 = `
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert    = "/nsconfig/ssl/certificate1.crt"
  key     = "/nsconfig/ssl/key1.pem"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.44"
  name        = "tf_lbvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
  vservername = citrixadc_lbvserver.tf_lbvserver.name
  certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
}

data "citrixadc_nitro_info" "sample" {
  workflow = {
    lifecycle                        = "binding_list"
    endpoint                         = "sslcertkey_sslvserver_binding"
    bound_resource_missing_errorcode = "1540"
  }
  primary_id = "tf_sslcertkey"
  depends_on = [citrixadc_sslvserver_sslcertkey_binding.tf_binding]
}
`

func TestAccNitroInfo_binding_list(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNitroInfoBindingList_step1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nitro_info.sample", "nitro_list.#", "1"),
				),
			},
		},
	})
}
