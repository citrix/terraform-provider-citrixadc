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

const testAccSslprofile_sslcertkey_binding_basic_step1 = `
resource "citrixadc_sslprofile" "tfUnit_sslprofile-hello" {
	name = "tfUnit_sslprofile-hello"

	// ecccurvebindings is REQUIRED attribute.
	// The default ecccurvebindings will be DELETED and only the explicitly given ecccurvebindings will be retained
	// To unbind all the ecccurvebindings, an empty list [] is to be assinged to ecccurvebindings attribute

	ecccurvebindings = ["P_256"]
	sslinterception = "ENABLED"
  
}
  
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	  certkey = "tf_sslcertkey"
	  cert = "/nsconfig/ssl/ns-root.cert"
	  key = "/nsconfig/ssl/ns-root.key"
}
	
resource "citrixadc_sslprofile_sslcertkey_binding" "demo_sslprofile_sslcertkey_binding" {
	  name = citrixadc_sslprofile.tfUnit_sslprofile-hello.name
	  sslicacertkey = citrixadc_sslcertkey.tf_sslcertkey.certkey 
}
  `

const testAccSslprofile_sslcertkey_binding_basic_step2 = `
resource "citrixadc_sslprofile" "tfUnit_sslprofile-hello" {
	name = "tfUnit_sslprofile-hello"

	// ecccurvebindings is REQUIRED attribute.
	// The default ecccurvebindings will be DELETED and only the explicitly given ecccurvebindings will be retained
	// To unbind all the ecccurvebindings, an empty list [] is to be assinged to ecccurvebindings attribute
	
	ecccurvebindings = ["P_256"]
	sslinterception = "ENABLED"
  
}
  
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	  certkey = "tf_sslcertkey"
	  cert = "/nsconfig/ssl/ns-root.cert"
	  key = "/nsconfig/ssl/ns-root.key"
}
`

func TestAccSslprofile_sslcertkey_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofile_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_sslcertkey_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_sslcertkey_bindingExist("citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding", nil),
				),
			},
			{
				Config: testAccSslprofile_sslcertkey_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_sslcertkey_bindingNotExist("citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding", "tfUnit_sslprofile-hello,tf_sslcertkey"),
				),
			},
		},
	})
}

func testAccCheckSslprofile_sslcertkey_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslprofile_sslcertkey_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "sslicacertkey"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		sslicacertkey := idMap["sslicacertkey"]

		findParams := service.FindParams{
			ResourceType:             "sslprofile_sslcertkey_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 3248,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["sslicacertkey"].(string) == sslicacertkey {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslprofile_sslcertkey_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslprofile_sslcertkey_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		name := idSlice[0]
		sslicacertkey := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslprofile_sslcertkey_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 3248,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["sslicacertkey"].(string) == sslicacertkey {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslprofile_sslcertkey_binding %s not deleted", n)
		}

		return nil
	}
}

func testAccCheckSslprofile_sslcertkey_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslprofile_sslcertkey_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("sslprofile_sslcertkey_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslprofile_sslcertkey_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSslprofile_sslcertkey_bindingDataSource_basic = `
resource "citrixadc_sslprofile" "tfUnit_sslprofile-hello" {
	name = "tfUnit_sslprofile-hello"
	ecccurvebindings = ["P_256"]
	sslinterception = "ENABLED"
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	certkey = "tf_sslcertkey"
	cert = "/nsconfig/ssl/ns-root.cert"
	key = "/nsconfig/ssl/ns-root.key"
}

resource "citrixadc_sslprofile_sslcertkey_binding" "demo_sslprofile_sslcertkey_binding" {
	name = citrixadc_sslprofile.tfUnit_sslprofile-hello.name
	sslicacertkey = citrixadc_sslcertkey.tf_sslcertkey.certkey
}

data "citrixadc_sslprofile_sslcertkey_binding" "demo_sslprofile_sslcertkey_binding" {
	name = citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding.name
	sslicacertkey = citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding.sslicacertkey
}
`

func TestAccSslprofile_sslcertkey_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_sslcertkey_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding", "name", "tfUnit_sslprofile-hello"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding", "sslicacertkey", "tf_sslcertkey"),
				),
			},
		},
	})
}

const testAccsslprofile_sslcertkey_binding_upgrade_basic = `
resource "citrixadc_sslprofile" "tfUnit_sslprofile-hello" {
	name = "tfUnit_sslprofile-hello"

	// ecccurvebindings is REQUIRED attribute.
	// The default ecccurvebindings will be DELETED and only the explicitly given ecccurvebindings will be retained
	// To unbind all the ecccurvebindings, an empty list [] is to be assinged to ecccurvebindings attribute

	ecccurvebindings = ["P_256"]
	sslinterception = "ENABLED"

}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	  certkey = "tf_sslcertkey"
	  cert = "/nsconfig/ssl/ns-root.cert"
	  key = "/nsconfig/ssl/ns-root.key"
}

resource "citrixadc_sslprofile_sslcertkey_binding" "demo_sslprofile_sslcertkey_binding" {
	  name = citrixadc_sslprofile.tfUnit_sslprofile-hello.name
	  sslicacertkey = citrixadc_sslcertkey.tf_sslcertkey.certkey
}
  `

func TestAccSslprofile_sslcertkey_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			// sslprofile_* bindings require the default SSL profile enabled; skip unless the
			// run is labelled for that testbed (matches the sibling _basic gate on this resource family).
			if adcTestbed != "STANDALONE_DEFAULT_SSL_PROFILE" {
				t.Skipf("ADC testbed is %s. Expected STANDALONE_DEFAULT_SSL_PROFILE.", adcTestbed)
			}
		},
		CheckDestroy: testAccCheckSslprofile_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create the resource with the last SDK v2 release (writes legacy-format id).
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.0.0",
					},
				},
				Config: testAccsslprofile_sslcertkey_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_sslcertkey_bindingExist("citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding", "id", "tfUnit_sslprofile-hello,tf_sslcertkey"),
				),
			},
			// Step 2: Refresh/apply the legacy-id state through the current (framework) provider.
			// Read recomputes the id into the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccsslprofile_sslcertkey_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_sslcertkey_bindingExist("citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding", "id", "name:tfUnit_sslprofile-hello,sslicacertkey:tf_sslcertkey"),
				),
			},
		},
	})
}

func TestAccSslprofile_sslcertkey_binding_import(t *testing.T) {
	const resAddr = "citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,sslicacertkey) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "sslicacertkey"}
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
		PreCheck: func() {
			testAccPreCheck(t)
			// sslprofile_* bindings require the default SSL profile enabled; skip unless the
			// run is labelled for that testbed (matches the sibling gate on this resource family).
			if adcTestbed != "STANDALONE_DEFAULT_SSL_PROFILE" {
				t.Skipf("ADC testbed is %s. Expected STANDALONE_DEFAULT_SSL_PROFILE.", adcTestbed)
			}
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofile_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslprofile_sslcertkey_binding_basic_step1},
			{Config: testAccSslprofile_sslcertkey_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccSslprofile_sslcertkey_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccSslprofile_sslcertkey_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofile_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_sslcertkey_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslprofile_sslcertkey_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Sslprofile_sslcertkey_binding.Type(), "tfUnit_sslprofile-hello", map[string]string{"sslicacertkey": "tf_sslcertkey"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSslprofile_sslcertkey_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslprofile_sslcertkey_bindingExist(resAddr, nil)),
			},
		},
	})
}
