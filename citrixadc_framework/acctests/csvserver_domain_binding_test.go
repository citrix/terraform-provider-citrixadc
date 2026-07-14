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

// Participating entity reused from existing acceptance test:
//   - csvserver : citrixadc_framework/acctests/csvserver_test.go (testAccCsvserver_basic block)
//
// Composite ID is "domainname:<v>,name:<v>" (see resource_schema.go SetAttrFromGet / Create).
// The optional attributes (ttl, backupip, cookiedomain, cookietimeout, sitedomainttl) are
// updateable in place, while name + domainname are RequiresReplace identity keys.
// The basic test therefore creates the binding in step1, then changes updateable optionals
// (ttl, cookietimeout) in step2 while keeping name + domainname constant, exercising the
// in-place PUT-bind update path.
//
// IMPORTANT: A domain can only be bound to a CS vserver whose targetType is GSLB
// (verified via CLI: "bind cs vserver <name> -domainName ..." returns errorcode 257
// "Operation not permitted" on an ordinary IP/port-based CS vserver, but succeeds on a
// GSLB-targetType CS vserver). A GSLB-targetType CS vserver is created without an IP/port.

const testAccCsvserver_domain_binding_basic_step1 = `
resource "citrixadc_csvserver" "tf_csvserver" {
  name        = "tf_csvserver_domain"
  servicetype = "HTTP"
  targettype  = "GSLB"
}

resource "citrixadc_csvserver_domain_binding" "tf_csvserver_domain_binding" {
  name          = citrixadc_csvserver.tf_csvserver.name
  domainname    = "example.com"
  ttl           = 3600
  cookietimeout = 10
  depends_on    = [citrixadc_csvserver.tf_csvserver]
}

`

const testAccCsvserver_domain_binding_basic_step2 = `
resource "citrixadc_csvserver" "tf_csvserver" {
  name        = "tf_csvserver_domain"
  servicetype = "HTTP"
  targettype  = "GSLB"
}

resource "citrixadc_csvserver_domain_binding" "tf_csvserver_domain_binding" {
  name          = citrixadc_csvserver.tf_csvserver.name
  domainname    = "example.com"
  ttl           = 7200
  cookietimeout = 20
  depends_on    = [citrixadc_csvserver.tf_csvserver]
}

`

func TestAccCsvserver_domain_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_domain_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_domain_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_domain_bindingExist("citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", "name", "tf_csvserver_domain"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", "domainname", "example.com"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", "ttl", "3600"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", "cookietimeout", "10"),
				),
			},
			{
				Config: testAccCsvserver_domain_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_domain_bindingExist("citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", "name", "tf_csvserver_domain"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", "domainname", "example.com"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", "ttl", "7200"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", "cookietimeout", "20"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_domain_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_domain_binding id is set")
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

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		domainname := idMap["domainname"]

		findParams := service.FindParams{
			ResourceType:             service.Csvserver_domain_binding.Type(),
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
			if dn, ok := v["domainname"].(string); ok && dn == domainname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("csvserver_domain_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_domain_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_domain_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		domainname := idMap["domainname"]

		findParams := service.FindParams{
			ResourceType:             service.Csvserver_domain_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent csvserver or binding no longer exists - destroyed as expected
			continue
		}

		found := false
		for _, v := range dataArr {
			if dn, ok := v["domainname"].(string); ok && dn == domainname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("csvserver_domain_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccCsvserver_domain_binding_DataSource_basic = `
resource "citrixadc_csvserver" "tf_csvserver" {
  name        = "tf_csvserver_domain"
  servicetype = "HTTP"
  targettype  = "GSLB"
}

resource "citrixadc_csvserver_domain_binding" "tf_csvserver_domain_binding" {
  name          = citrixadc_csvserver.tf_csvserver.name
  domainname    = "example.com"
  ttl           = 3600
  cookietimeout = 10
  depends_on    = [citrixadc_csvserver.tf_csvserver]
}

data "citrixadc_csvserver_domain_binding" "tf_csvserver_domain_binding" {
  name       = citrixadc_csvserver.tf_csvserver.name
  domainname = "example.com"
  depends_on = [citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding]
}
`

func TestAccCsvserver_domain_binding_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_domain_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_domain_binding_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", "name", "tf_csvserver_domain"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_domain_binding.tf_csvserver_domain_binding", "domainname", "example.com"),
				),
			},
		},
	})
}
