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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func doSslcacertgroup_sslcertkey_bindingPreChecks(t *testing.T) {
	testAccPreCheck(t)

	uploads := []string{
		"rootcert1.cert",
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	for _, filename := range uploads {
		err := uploadTestdataFile(c, t, filename, "/var/tmp")
		if err != nil {
			t.Errorf("%v", err)
		}
	}
}

const testAccSslcacertgroup_sslcertkey_binding_basic_step1 = `
	resource "citrixadc_sslcacertgroup_sslcertkey_binding" "sslcacertgroup_sslcertkey_binding_demo" {	
        cacertgroupname = citrixadc_sslcacertgroup.ns_callout_certs1.cacertgroupname
		certkeyname = citrixadc_sslcertkey.tf_cacertkey.certkey
        ocspcheck = "Mandatory"
	}

	resource "citrixadc_sslcertkey" "tf_cacertkey" {
		certkey = "tf_cacertkey"
		cert = "/var/tmp/rootcert1.cert"
	}
		
	resource "citrixadc_sslcacertgroup" "ns_callout_certs1" {
		cacertgroupname = "ns_callout_certs1"
	}
`

const testAccSslcacertgroup_sslcertkey_binding_basic_step2 = `

resource "citrixadc_sslcertkey" "tf_cacertkey" {
	certkey = "tf_cacertkey"
	cert = "/var/tmp/rootcert1.cert"
}
	
resource "citrixadc_sslcacertgroup" "ns_callout_certs1" {
	cacertgroupname = "ns_callout_certs1"
}
`

func TestAccSslcacertgroup_sslcertkey_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcacertgroup_sslcertkey_bindingPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcacertgroup_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcacertgroup_sslcertkey_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcacertgroup_sslcertkey_bindingExist("citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo", nil),
				),
			},
			{
				Config: testAccSslcacertgroup_sslcertkey_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcacertgroup_sslcertkey_bindingNotExist("citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo", "ns_callout_certs1,tf_cacertkey"),
				),
			},
		},
	})
}

func testAccCheckSslcacertgroup_sslcertkey_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcacertgroup_sslcertkey_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"cacertgroupname", "certkeyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}

		cacertgroupname := idMap["cacertgroupname"]
		certkeyname := idMap["certkeyname"]

		findParams := service.FindParams{
			ResourceType:             "sslcacertgroup_sslcertkey_binding",
			ResourceName:             cacertgroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor cacertgroupname
		found := false
		for _, v := range dataArr {
			if v["certkeyname"].(string) == certkeyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslcacertgroup_sslcertkey_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcacertgroup_sslcertkey_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}

		idSlice := strings.SplitN(id, ",", 2)

		certkeyname := idSlice[0]
		cacertgroupname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslcacertgroup_sslcertkey_binding",
			ResourceName:             certkeyname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["cacertgroupname"].(string) == cacertgroupname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslcacertgroup_sslcertkey_binding %s not deleted", n)
		}

		return nil
	}
}

func testAccCheckSslcacertgroup_sslcertkey_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcacertgroup_sslcertkey_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cacertgroupname is set")
		}

		_, err := client.FindResource("sslcacertgroup_sslcertkey_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslcacertgroup_sslcertkey_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSslcacertgroup_sslcertkey_bindingDataSource_basic = `
	resource "citrixadc_sslcacertgroup_sslcertkey_binding" "sslcacertgroup_sslcertkey_binding_demo" {	
        cacertgroupname = citrixadc_sslcacertgroup.ns_callout_certs1.cacertgroupname
		certkeyname = citrixadc_sslcertkey.tf_cacertkey.certkey
        ocspcheck = "Mandatory"
	}

	resource "citrixadc_sslcertkey" "tf_cacertkey" {
		certkey = "tf_cacertkey"
		cert = "/var/tmp/rootcert1.cert"
	}
		
	resource "citrixadc_sslcacertgroup" "ns_callout_certs1" {
		cacertgroupname = "ns_callout_certs1"
	}

	data "citrixadc_sslcacertgroup_sslcertkey_binding" "sslcacertgroup_sslcertkey_binding_demo" {
		cacertgroupname = citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo.cacertgroupname
		certkeyname = citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo.certkeyname
	}
`

func TestAccSslcacertgroup_sslcertkey_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcacertgroup_sslcertkey_bindingPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcacertgroup_sslcertkey_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo", "cacertgroupname", "ns_callout_certs1"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo", "certkeyname", "tf_cacertkey"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo", "ocspcheck", "Mandatory"),
				),
			},
		},
	})
}

const testAccsslcacertgroup_sslcertkey_binding_upgrade_basic = `
	resource "citrixadc_sslcacertgroup_sslcertkey_binding" "sslcacertgroup_sslcertkey_binding_demo" {
        cacertgroupname = citrixadc_sslcacertgroup.ns_callout_certs1.cacertgroupname
		certkeyname = citrixadc_sslcertkey.tf_cacertkey.certkey
        ocspcheck = "Mandatory"
	}

	resource "citrixadc_sslcertkey" "tf_cacertkey" {
		certkey = "tf_cacertkey"
		cert = "/var/tmp/rootcert1.cert"
	}

	resource "citrixadc_sslcacertgroup" "ns_callout_certs1" {
		cacertgroupname = "ns_callout_certs1"
	}
`

func TestAccSslcacertgroup_sslcertkey_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { doSslcacertgroup_sslcertkey_bindingPreChecks(t) },
		CheckDestroy: testAccCheckSslcacertgroup_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy id
				// (d.SetId("<cacertgroupname>,<certkeyname>")).
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccsslcacertgroup_sslcertkey_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcacertgroup_sslcertkey_bindingExist("citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo", "id", "ns_callout_certs1,tf_cacertkey"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id to the canonical new format
				// (comma-joined key:UrlEncode(value) pairs in idParts order).
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccsslcacertgroup_sslcertkey_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcacertgroup_sslcertkey_bindingExist("citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo", "id", "cacertgroupname:ns_callout_certs1,certkeyname:tf_cacertkey"),
				),
			},
		},
	})
}

func TestAccSslcacertgroup_sslcertkey_binding_import(t *testing.T) {
	const resAddr = "citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcacertgroup_sslcertkey_bindingPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcacertgroup_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslcacertgroup_sslcertkey_binding_basic_step1},
			{Config: testAccSslcacertgroup_sslcertkey_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
