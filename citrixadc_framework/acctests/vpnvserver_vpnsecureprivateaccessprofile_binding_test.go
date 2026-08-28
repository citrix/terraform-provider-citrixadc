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
	"net/url"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// NOTE: vpnvserver_vpnsecureprivateaccessprofile_binding binds a vpnvserver to a
// vpnsecureprivateaccessprofile (Secure Private Access profile). The configs below
// create both parents (a citrixadc_vpnsecureprivateaccessprofile fixture and a
// citrixadc_vpnvserver) inline so the tests are self-contained.
const testAccVpnvserver_vpnsecureprivateaccessprofile_binding_basic = `
resource "citrixadc_vpnsecureprivateaccessprofile" "tf_spa_profile" {
	name = "tf_spa_profile"
	url  = "https://www.citrix.com"
}

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
	name        = "tf_vserver"
	servicetype = "SSL"
	ipv46       = "3.3.3.3"
	port        = 443
}

resource "citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding" "tf_bind" {
	name                       = citrixadc_vpnvserver.tf_vpnvserver.name
	secureprivateaccessprofile = citrixadc_vpnsecureprivateaccessprofile.tf_spa_profile.name
}
`

const testAccVpnvserver_vpnsecureprivateaccessprofile_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

resource "citrixadc_vpnsecureprivateaccessprofile" "tf_spa_profile" {
	name = "tf_spa_profile"
	url  = "https://www.citrix.com"
}

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
	name        = "tf_vserver"
	servicetype = "SSL"
	ipv46       = "3.3.3.3"
	port        = 443
}
`

func TestAccVpnvserver_vpnsecureprivateaccessprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_vpnsecureprivateaccessprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_vpnsecureprivateaccessprofile_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpnsecureprivateaccessprofile_bindingExist("citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccVpnvserver_vpnsecureprivateaccessprofile_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpnsecureprivateaccessprofile_bindingNotExist("citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind", "tf_vserver,tf_spa_profile"),
				),
			},
		},
	})
}

func testAccCheckVpnvserver_vpnsecureprivateaccessprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_vpnsecureprivateaccessprofile_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "secureprivateaccessprofile"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		secureprivateaccessprofile := idMap["secureprivateaccessprofile"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_vpnsecureprivateaccessprofile_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["secureprivateaccessprofile"].(string) == secureprivateaccessprofile {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnvserver_vpnsecureprivateaccessprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_vpnsecureprivateaccessprofile_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"name", "secureprivateaccessprofile"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		name := idMap["name"]
		secureprivateaccessprofile := idMap["secureprivateaccessprofile"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_vpnsecureprivateaccessprofile_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["secureprivateaccessprofile"].(string) == secureprivateaccessprofile {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnvserver_vpnsecureprivateaccessprofile_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_vpnsecureprivateaccessprofile_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnvserver_vpnsecureprivateaccessprofile_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver_vpnsecureprivateaccessprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnvserver_vpnsecureprivateaccessprofile_bindingDataSource_basic = `
resource "citrixadc_vpnsecureprivateaccessprofile" "tf_spa_profile" {
	name = "tf_spa_profile"
	url  = "https://www.citrix.com"
}

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
	name        = "tf_vserver"
	servicetype = "SSL"
	ipv46       = "3.3.3.3"
	port        = 443
}

resource "citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding" "tf_bind" {
	name                       = citrixadc_vpnvserver.tf_vpnvserver.name
	secureprivateaccessprofile = citrixadc_vpnsecureprivateaccessprofile.tf_spa_profile.name
}

data "citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding" "tf_bind" {
	name                       = citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind.name
	secureprivateaccessprofile = citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind.secureprivateaccessprofile
	depends_on                 = [citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind]
}
`

func TestAccVpnvserver_vpnsecureprivateaccessprofile_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_vpnsecureprivateaccessprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_vpnsecureprivateaccessprofile_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind", "name", "tf_vserver"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind", "secureprivateaccessprofile", "tf_spa_profile"),
				),
			},
		},
	})
}

func TestAccVpnvserver_vpnsecureprivateaccessprofile_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,secureprivateaccessprofile) so it matches exactly what SDK v2 wrote.
	legacyID := func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resAddr]
		if !ok {
			return "", fmt.Errorf("resource not found in state: %s", resAddr)
		}
		kv := map[string]string{}
		for _, p := range strings.Split(rs.Primary.ID, ",") {
			if i := strings.Index(p, ":"); i >= 0 {
				v, _ := url.QueryUnescape(p[i+1:])
				kv[p[:i]] = v
			}
		}
		ordr := []string{"name", "secureprivateaccessprofile"}
		parts := make([]string, 0, len(ordr))
		for _, k := range ordr {
			if v, ok := kv[k]; ok {
				parts = append(parts, v)
			}
		}
		// Fallback: a positional (non key:value) id has no key:value parts to reorder; import it as-is.
		if len(parts) == 0 {
			return rs.Primary.ID, nil
		}
		return strings.Join(parts, ","), nil
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_vpnsecureprivateaccessprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnvserver_vpnsecureprivateaccessprofile_binding_basic},
			{Config: testAccVpnvserver_vpnsecureprivateaccessprofile_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccVpnvserver_vpnsecureprivateaccessprofile_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccVpnvserver_vpnsecureprivateaccessprofile_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_vpnsecureprivateaccessprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_vpnsecureprivateaccessprofile_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpnsecureprivateaccessprofile_bindingExist(resAddr, nil),
				),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Vpnvserver_vpnsecureprivateaccessprofile_binding.Type(), "tf_vserver", map[string]string{"secureprivateaccessprofile": "tf_spa_profile"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnvserver_vpnsecureprivateaccessprofile_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_vpnsecureprivateaccessprofile_bindingExist(resAddr, nil),
				),
			},
		},
	})
}
