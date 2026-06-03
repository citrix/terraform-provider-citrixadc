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

// DEPRECATION NOTE:
// NitroValidator confirmed that current NetScaler firmware REJECTS the
// "bind vpn vserver -epaprofile" operation with the error:
//   "There has been a design change in the support of OPSWAT specific EPA
//    scans. EPA Profile Configuration is no longer needed."
// This binding is therefore non-functional on current firmware and a live
// apply will fail. All tests below are guarded with t.Skip(...) and should be
// re-enabled only on firmware that still supports EPA profile bindings.
// The full, correct test bodies are retained so they are ready to run on
// supporting firmware.

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Step 1: create the vpn vserver + vpnepaprofile, then bind them.
// (vpnvserver block lifted from vpnvserver_test.go; vpnepaprofile block lifted
//
//	from vpnepaprofile_test.go.)
const testAccVpnvserverVpnepaprofileBinding_basic_step1 = `
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf.citrix.example.com"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}

resource "citrixadc_vpnepaprofile" "tf_vpnepaprofile" {
  name     = "tf_vpnepaprofile"
  filename = "tf_vpnepaprofile.xml"
}

resource "citrixadc_vpnvserver_vpnepaprofile_binding" "tf_binding" {
  name               = citrixadc_vpnvserver.tf_vpnvserver.name
  epaprofile         = citrixadc_vpnepaprofile.tf_vpnepaprofile.name
  epaprofileoptional = true

  depends_on = [
    citrixadc_vpnvserver.tf_vpnvserver,
    citrixadc_vpnepaprofile.tf_vpnepaprofile,
  ]
}
`

// Step 2: drop the binding, keep the participating entities. This confirms the
// binding is removed (CheckDestroy / FindResourceArrayWithParams filter).
const testAccVpnvserverVpnepaprofileBinding_basic_step2 = `
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf.citrix.example.com"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}

resource "citrixadc_vpnepaprofile" "tf_vpnepaprofile" {
  name     = "tf_vpnepaprofile"
  filename = "tf_vpnepaprofile.xml"
}
`

func TestAccVpnvserverVpnepaprofileBinding_basic(t *testing.T) {
	t.Skip("vpnvserver_vpnepaprofile_binding is deprecated: current NetScaler firmware rejects 'bind vpn vserver -epaprofile' (OPSWAT EPA design change - EPA Profile Configuration is no longer needed). Re-enable only on firmware that still supports EPA profile bindings.")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserverVpnepaprofileBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserverVpnepaprofileBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverVpnepaprofileBindingExist("citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding", "epaprofile", "tf_vpnepaprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding", "epaprofileoptional", "true"),
				),
			},
			{
				// Binding removed - the participating entities remain. CheckDestroy
				// verifies the binding no longer exists on the ADC.
				Config: testAccVpnvserverVpnepaprofileBinding_basic_step2,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func testAccCheckVpnvserverVpnepaprofileBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_vpnepaprofile_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}
			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// Composite ID is "epaprofile:<v>,name:<v>" (key:value pairs).
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		name := idMap["name"]
		epaprofile := idMap["epaprofile"]

		findParams := service.FindParams{
			ResourceType:             service.Vpnvserver_vpnepaprofile_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["epaprofile"].(string); ok && val == epaprofile {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("vpnvserver_vpnepaprofile_binding %s not found", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckVpnvserverVpnepaprofileBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_vpnepaprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_vpnepaprofile_binding id is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		name := idMap["name"]
		epaprofile := idMap["epaprofile"]

		findParams := service.FindParams{
			ResourceType:             service.Vpnvserver_vpnepaprofile_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone / binding gone - treat as destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["epaprofile"].(string); ok && val == epaprofile {
				return fmt.Errorf("vpnvserver_vpnepaprofile_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

// Datasource references the binding by its unique attributes (name + epaprofile).
const testAccVpnvserverVpnepaprofileBindingDataSource_basic = `
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf.citrix.example.com"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}

resource "citrixadc_vpnepaprofile" "tf_vpnepaprofile" {
  name     = "tf_vpnepaprofile"
  filename = "tf_vpnepaprofile.xml"
}

resource "citrixadc_vpnvserver_vpnepaprofile_binding" "tf_binding" {
  name               = citrixadc_vpnvserver.tf_vpnvserver.name
  epaprofile         = citrixadc_vpnepaprofile.tf_vpnepaprofile.name
  epaprofileoptional = true

  depends_on = [
    citrixadc_vpnvserver.tf_vpnvserver,
    citrixadc_vpnepaprofile.tf_vpnepaprofile,
  ]
}

data "citrixadc_vpnvserver_vpnepaprofile_binding" "tf_binding" {
  name       = citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding.name
  epaprofile = citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding.epaprofile

  depends_on = [citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding]
}
`

func TestAccVpnvserverVpnepaprofileBindingDataSource_basic(t *testing.T) {
	t.Skip("vpnvserver_vpnepaprofile_binding is deprecated: current NetScaler firmware rejects 'bind vpn vserver -epaprofile' (OPSWAT EPA design change - EPA Profile Configuration is no longer needed). Re-enable only on firmware that still supports EPA profile bindings.")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserverVpnepaprofileBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding", "epaprofile", "tf_vpnepaprofile"),
				),
			},
		},
	})
}
