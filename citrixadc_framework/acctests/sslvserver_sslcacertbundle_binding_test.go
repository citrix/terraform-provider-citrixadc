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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// sslvserver_sslcacertbundle_binding joins an SSL vserver (the SSL view of an
// SSL-type lbvserver) with an existing CA-cert-bundle (citrixadc_sslcacertbundle).
//
// The SSL vserver is created here as a citrixadc_lbvserver of servicetype "SSL".
// The CA-cert-bundle requires a bundle file that ALREADY EXISTS on the appliance
// under /nsconfig/ssl/ - the bundlefile value below is a TODO_PLACEHOLDER and must
// be replaced with a real CA-cert-bundle file present on your testbed.
//
// Composite ID = "cacertbundlename:<v>,vservername:<v>" (key:value pairs). By-name
// GET works: FindResourceArrayWithParams keyed on vservername, then match
// cacertbundlename in the returned array. Parse the ID with utils.ParseIdString.

const testAccSslvserver_sslcacertbundle_binding_basic_step1 = `
	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name        = "tf_sslvserver"
		servicetype = "SSL"
		ipv46       = "5.5.5.5"
		port        = 443
	}

	resource "citrixadc_sslcacertbundle" "tf_sslcacertbundle" {
		cacertbundlename = "tf_sslcacertbundle"
		bundlefile       = "ca-bundle.pem" // CA-cert-bundle file is uploaded to /nsconfig/ssl/ by doSslcacertbundlePreChecks
	}

	resource "citrixadc_sslvserver_sslcacertbundle_binding" "tf_sslvserver_sslcacertbundle_binding" {
		vservername      = citrixadc_lbvserver.tf_sslvserver.name
		cacertbundlename = citrixadc_sslcacertbundle.tf_sslcacertbundle.cacertbundlename
		depends_on       = [citrixadc_sslcacertbundle.tf_sslcacertbundle]
	}
`

// step2 drops the binding (and the bundle) but keeps the SSL vserver, so the
// binding's removal can be verified.
const testAccSslvserver_sslcacertbundle_binding_basic_step2 = `
	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name        = "tf_sslvserver"
		servicetype = "SSL"
		ipv46       = "5.5.5.5"
		port        = 443
	}
`

func TestAccSslvserver_sslcacertbundle_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcacertbundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserver_sslcacertbundle_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_sslcacertbundle_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcacertbundle_bindingExist("citrixadc_sslvserver_sslcacertbundle_binding.tf_sslvserver_sslcacertbundle_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcacertbundle_binding.tf_sslvserver_sslcacertbundle_binding", "vservername", "tf_sslvserver"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcacertbundle_binding.tf_sslvserver_sslcacertbundle_binding", "cacertbundlename", "tf_sslcacertbundle"),
				),
			},
			{
				Config: testAccSslvserver_sslcacertbundle_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcacertbundle_bindingNotExist("tf_sslvserver", "tf_sslcacertbundle"),
				),
			},
		},
	})
}

func testAccCheckSslvserver_sslcacertbundle_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslvserver_sslcacertbundle_binding id is set")
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

		// Composite ID = "cacertbundlename:<v>,vservername:<v>"
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		vservername := idMap["vservername"]
		cacertbundlename := idMap["cacertbundlename"]

		findParams := service.FindParams{
			ResourceType:             service.Sslvserver_sslcacertbundle_binding.Type(),
			ResourceName:             vservername,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		// Match the member (cacertbundlename) in the vservername array.
		found := false
		for _, v := range dataArr {
			if val, ok := v["cacertbundlename"].(string); ok && val == cacertbundlename {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslvserver_sslcacertbundle_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckSslvserver_sslcacertbundle_bindingNotExist verifies the binding is gone
// while the parent vserver still exists (step2 keeps the vserver).
func testAccCheckSslvserver_sslcacertbundle_bindingNotExist(vservername, cacertbundlename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Sslvserver_sslcacertbundle_binding.Type(),
			ResourceName:             vservername,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Vserver may have no bindings at all - acceptable.
			return nil
		}

		for _, v := range dataArr {
			if val, ok := v["cacertbundlename"].(string); ok && val == cacertbundlename {
				return fmt.Errorf("sslvserver_sslcacertbundle_binding for %s/%s still exists", vservername, cacertbundlename)
			}
		}

		return nil
	}
}

func testAccCheckSslvserver_sslcacertbundle_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslvserver_sslcacertbundle_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		vservername := idMap["vservername"]
		cacertbundlename := idMap["cacertbundlename"]

		findParams := service.FindParams{
			ResourceType:             service.Sslvserver_sslcacertbundle_binding.Type(),
			ResourceName:             vservername,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone or no bindings - binding is destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["cacertbundlename"].(string); ok && val == cacertbundlename {
				return fmt.Errorf("sslvserver_sslcacertbundle_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccSslvserver_sslcacertbundle_bindingDataSource_basic = `
	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name        = "tf_sslvserver"
		servicetype = "SSL"
		ipv46       = "5.5.5.5"
		port        = 443
	}

	resource "citrixadc_sslcacertbundle" "tf_sslcacertbundle" {
		cacertbundlename = "tf_sslcacertbundle"
		bundlefile       = "ca-bundle.pem" // CA-cert-bundle file is uploaded to /nsconfig/ssl/ by doSslcacertbundlePreChecks
	}

	resource "citrixadc_sslvserver_sslcacertbundle_binding" "tf_sslvserver_sslcacertbundle_binding" {
		vservername      = citrixadc_lbvserver.tf_sslvserver.name
		cacertbundlename = citrixadc_sslcacertbundle.tf_sslcacertbundle.cacertbundlename
		depends_on       = [citrixadc_sslcacertbundle.tf_sslcacertbundle]
	}

	data "citrixadc_sslvserver_sslcacertbundle_binding" "tf_sslvserver_sslcacertbundle_binding" {
		vservername      = citrixadc_lbvserver.tf_sslvserver.name
		cacertbundlename = citrixadc_sslcacertbundle.tf_sslcacertbundle.cacertbundlename
		depends_on       = [citrixadc_sslvserver_sslcacertbundle_binding.tf_sslvserver_sslcacertbundle_binding]
	}
`

func TestAccSslvserver_sslcacertbundle_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcacertbundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserver_sslcacertbundle_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_sslcacertbundle_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslvserver_sslcacertbundle_binding.tf_sslvserver_sslcacertbundle_binding", "vservername", "tf_sslvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_sslvserver_sslcacertbundle_binding.tf_sslvserver_sslcacertbundle_binding", "cacertbundlename", "tf_sslcacertbundle"),
				),
			},
		},
	})
}
