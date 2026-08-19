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
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccSystemuser_nspartition_binding_basic = `

resource "citrixadc_systemuser_nspartition_binding" "tf_systemuser_nspartition_binding" {
	username      = citrixadc_systemuser.user.username
	partitionname = citrixadc_nspartition.tf_nspartition.partitionname
	}
  
  resource "citrixadc_nspartition" "tf_nspartition" {
	partitionname = "tf_nspartition"
	maxbandwidth  = 10240
	minbandwidth  = 512
	maxconn       = 512
	maxmemlimit   = 11
	}
  
  
  resource "citrixadc_systemuser" "user" {
	username = "george"
	password = "12345"
	timeout  = 900
	}
`

const testAccSystemuser_nspartition_binding_basic_step2 = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
	
	
	resource "citrixadc_systemuser" "user" {
		username = "george"
		password = "12345"
		timeout  = 900
	}
`

func TestAccSystemuser_nspartition_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemuser_nspartition_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemuser_nspartition_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuser_nspartition_bindingExist("citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding", nil),
				),
			},
			{
				Config: testAccSystemuser_nspartition_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuser_nspartition_bindingNotExist("citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding", "george,tf_nspartition"),
				),
			},
		},
	})
}

func testAccCheckSystemuser_nspartition_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemuser_nspartition_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"username", "partitionname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		username := idMap["username"]
		partitionname := idMap["partitionname"]

		findParams := service.FindParams{
			ResourceType:             "systemuser_nspartition_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching partitionname
		found := false
		for _, v := range dataArr {
			if v["partitionname"].(string) == partitionname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("systemuser_nspartition_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemuser_nspartition_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"username", "partitionname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		username := idMap["username"]
		partitionname := idMap["partitionname"]

		findParams := service.FindParams{
			ResourceType:             "systemuser_nspartition_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching partitionname
		found := false
		for _, v := range dataArr {
			if v["partitionname"].(string) == partitionname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("systemuser_nspartition_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSystemuser_nspartition_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemuser_nspartition_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Systemuser_nspartition_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("systemuser_nspartition_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystemuser_nspartition_bindingDataSource_basic = `

resource "citrixadc_nspartition" "tf_nspartition" {
	partitionname = "tf_nspartition"
	maxbandwidth  = 10240
	minbandwidth  = 512
	maxconn       = 512
	maxmemlimit   = 11
}


resource "citrixadc_systemuser" "user" {
	username = "george"
	password = "12345"
	timeout  = 900
}

resource "citrixadc_systemuser_nspartition_binding" "tf_systemuser_nspartition_binding" {
	username      = citrixadc_systemuser.user.username
	partitionname = citrixadc_nspartition.tf_nspartition.partitionname
}

data "citrixadc_systemuser_nspartition_binding" "tf_systemuser_nspartition_binding" {
	username      = citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding.username
	partitionname = citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding.partitionname
	depends_on    = [citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding]
}
`

const testAccSystemuser_nspartition_binding_upgrade_basic = `

resource "citrixadc_systemuser_nspartition_binding" "tf_systemuser_nspartition_binding" {
	username      = citrixadc_systemuser.user.username
	partitionname = citrixadc_nspartition.tf_nspartition.partitionname
	}

  resource "citrixadc_nspartition" "tf_nspartition" {
	partitionname = "tf_nspartition"
	maxbandwidth  = 10240
	minbandwidth  = 512
	maxconn       = 512
	maxmemlimit   = 11
	}


  resource "citrixadc_systemuser" "user" {
	username = "george"
	password = "12345"
	timeout  = 900
	}
`

func TestAccSystemuser_nspartition_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSystemuser_nspartition_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with the last SDK v2 release (writes legacy id)
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccSystemuser_nspartition_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuser_nspartition_bindingExist("citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding", "id", "george,tf_nspartition"),
				),
			},
			{
				// Step 2: Refresh/plan/apply the legacy-id state through the current framework provider
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSystemuser_nspartition_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuser_nspartition_bindingExist("citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding", "id", "partitionname:tf_nspartition,username:george"),
				),
			},
		},
	})
}

func TestAccSystemuser_nspartition_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemuser_nspartition_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemuser_nspartition_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding", "username", "george"),
					resource.TestCheckResourceAttr("data.citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding", "partitionname", "tf_nspartition"),
				),
			},
		},
	})
}

func TestAccSystemuser_nspartition_binding_import(t *testing.T) {
	const resAddr = "citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: username,partitionname) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"username", "partitionname"}
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
		CheckDestroy:             testAccCheckSystemuser_nspartition_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSystemuser_nspartition_binding_basic},
			{Config: testAccSystemuser_nspartition_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccSystemuser_nspartition_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccSystemuser_nspartition_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemuser_nspartition_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemuser_nspartition_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSystemuser_nspartition_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Systemuser_nspartition_binding.Type(), "george", map[string]string{"partitionname": "tf_nspartition"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSystemuser_nspartition_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSystemuser_nspartition_bindingExist(resAddr, nil)),
			},
		},
	})
}
