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

const testAccGslbvserver_domain_binding_basic = `

resource "citrixadc_gslbvserver_domain_binding" "tf_gslbvserver_domain_binding"{
	name = citrixadc_gslbvserver.tf_gslbvserver.name
	domainname = "www.exampledomain.com"
	backupipflag = false
	}
  resource "citrixadc_gslbvserver" "tf_gslbvserver" {
	dnsrecordtype = "A"
	name          = "GSLB-East-Coast-Vserver"
	servicetype   = "HTTP"
	}
`

const testAccGslbvserver_domain_binding_basic_step2 = `
  resource "citrixadc_gslbvserver" "tf_gslbvserver" {
	dnsrecordtype = "A"
	name          = "GSLB-East-Coast-Vserver"
	servicetype   = "HTTP"
	}
`

const testAccGslbvserver_domain_bindingDataSource_basic = `

resource "citrixadc_gslbvserver_domain_binding" "tf_gslbvserver_domain_binding"{
	name = citrixadc_gslbvserver.tf_gslbvserver.name
	domainname = "www.exampledomain.com"
	backupipflag = false
}
resource "citrixadc_gslbvserver" "tf_gslbvserver" {
	dnsrecordtype = "A"
	name          = "GSLB-East-Coast-Vserver"
	servicetype   = "HTTP"
}

data "citrixadc_gslbvserver_domain_binding" "tf_gslbvserver_domain_binding" {
	name = citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding.name
	domainname = citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding.domainname
	depends_on = [citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding]
}
`

func TestAccGslbvserver_domain_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbvserver_domain_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbvserver_domain_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserver_domain_bindingExist("citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding", nil),
				),
			},
			{
				Config: testAccGslbvserver_domain_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserver_domain_bindingNotExist("citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding", "GSLB-East-Coast-Vserver,www.exampledomain.com"),
				),
			},
		},
	})
}

func testAccCheckGslbvserver_domain_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbvserver_domain_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "domainname"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		domainname := idMap["domainname"]

		findParams := service.FindParams{
			ResourceType:             "gslbvserver_domain_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching domainname
		found := false
		for _, v := range dataArr {
			if v["domainname"].(string) == domainname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("gslbvserver_domain_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbvserver_domain_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "domainname"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		domainname := idMap["domainname"]

		findParams := service.FindParams{
			ResourceType:             "gslbvserver_domain_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching domainname
		found := false
		for _, v := range dataArr {
			if v["domainname"].(string) == domainname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("gslbvserver_domain_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckGslbvserver_domain_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbvserver_domain_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Gslbvserver_domain_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbvserver_domain_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccGslbvserver_domain_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbvserver_domain_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding", "name", "GSLB-East-Coast-Vserver"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding", "domainname", "www.exampledomain.com"),
				),
			},
		},
	})
}

// testAccgslbvserver_domain_binding_upgrade_basic mirrors testAccGslbvserver_domain_binding_basic
// using only SDK v2 attribute names, so it is valid under BOTH the last SDK v2 release (2.2.0)
// schema and the current framework schema. The resource label is kept identical so the Exist /
// Destroy helpers and the state address match.
const testAccgslbvserver_domain_binding_upgrade_basic = `

resource "citrixadc_gslbvserver_domain_binding" "tf_gslbvserver_domain_binding"{
	name = citrixadc_gslbvserver.tf_gslbvserver.name
	domainname = "www.exampledomain.com"
	backupipflag = false
	}
  resource "citrixadc_gslbvserver" "tf_gslbvserver" {
	dnsrecordtype = "A"
	name          = "GSLB-East-Coast-Vserver"
	servicetype   = "HTTP"
	}
`

// TestAccGslbvserver_domain_binding_sdkv2StateUpgrade verifies that state written by the last
// SDK v2 release (with the legacy comma-joined ID) is refreshed and re-applied cleanly by the
// current framework provider.
//
//	Step 1: create the binding with citrix/citrixadc 2.2.0 (SDK v2). State carries the legacy
//	        ID "name,domainname".
//	Step 2: the SAME config served by the current framework provider. Terraform refreshes the
//	        legacy-id state through the framework Read (exercising utils.ParseIdString on the
//	        legacy id) then plans/applies. The framework Read does NOT recompute the ID
//	        (gslbvserver_domain_bindingSetAttrFromGet leaves data.Id untouched), so the legacy
//	        ID is retained after the upgrade.
func TestAccGslbvserver_domain_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGslbvserver_domain_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Create with the last SDK v2 release from the registry.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccgslbvserver_domain_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserver_domain_bindingExist("citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding", "id", "GSLB-East-Coast-Vserver,www.exampledomain.com"),
				),
			},
			{
				// Refresh/plan/apply the legacy-id state through the current framework provider.
				// The framework Read re-derives the canonical new-format id from name+domainname,
				// so after the upgrade the id is the key:value form (legacy->new-format only).
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccgslbvserver_domain_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserver_domain_bindingExist("citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding", "id", "name:GSLB-East-Coast-Vserver,domainname:www.exampledomain.com"),
				),
			},
		},
	})
}

func TestAccGslbvserver_domain_binding_import(t *testing.T) {
	const resAddr = "citrixadc_gslbvserver_domain_binding.tf_gslbvserver_domain_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbvserver_domain_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbvserver_domain_binding_basic,
			},
			{
				Config:                  testAccGslbvserver_domain_binding_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backupipflag"},
			},
		},
	})
}
