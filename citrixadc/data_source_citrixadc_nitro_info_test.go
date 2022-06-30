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

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/stretchr/testify/assert"
)

const testAccNitroInfoBindingList_step1 = `
data "citrixadc_nitro_info" "sample" {
    workflow = {
		lifecycle = "binding_list"
		endpoint = "sslcertkey_sslvserver_binding"
		bound_resource_missing_errorcode = 1540
	}
    primary_id = "tf_sslcertkey"
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/var/tmp/certificate1.crt"
  key = "/var/tmp/key1.pem"
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
`

func TestAccNitroInfo_binding_list(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNitroInfoBindingList_step1,
				Check:  resource.ComposeTestCheckFunc(
				//resource.TestCheckResourceAttr("data.citrixadc_nitro_info.sample", "nitro_list", "tolist"),
				),
			},
			resource.TestStep{
				Config: testAccNitroInfoBindingList_step1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nitro_info.sample", "nitro_list", "tolist"),
				),
			},
		},
	})
}

func TestTerratestNitroInfoBindingList(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/nitro_info/binding_list",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
	terraform.RunTerraformCommand(t, terraformOptions, "refresh")

	output := terraform.OutputList(t, terraformOptions, "object_output")
	emptyList := []string{"tf_csvserver", "tf_lbvserver"}
	assert.Equal(t, emptyList, output)
}

func TestTerratestNitroInfoObjectByName(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/nitro_info/object_by_name",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
	terraform.RunTerraformCommand(t, terraformOptions, "refresh")

	output := terraform.Output(t, terraformOptions, "nitro_object_length")
	assert.Equal(t, "33", output)
}
