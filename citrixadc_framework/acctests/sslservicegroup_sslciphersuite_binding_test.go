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

const testAccSslservicegroup_sslciphersuite_binding_basic = `

resource "citrixadc_sslservicegroup_sslciphersuite_binding" "tf_sslservicegroup_sslciphersuite_binding" {
	ciphername       = citrixadc_sslcipher.tf_sslcipher.ciphergroupname
	servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
	}
  resource "citrixadc_sslcipher" "tf_sslcipher" {
	  ciphergroupname = "my_ciphersuite"
	 
	  ciphersuitebinding {
		  ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		  cipherpriority = 1
	}
	}
  
  resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "my_gslbvservicegroup"
	servicetype      = "SSL_TCP"
	cip              = "DISABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
	}
  resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword = "password123"
	}
  
`

const testAccSslservicegroup_sslciphersuite_binding_basic_step2 = `
resource "citrixadc_sslcipher" "tf_sslcipher" {
	ciphergroupname = "my_ciphersuite"
   
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		cipherpriority = 1
	}
}

resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
  servicegroupname = "my_gslbvservicegroup"
  servicetype      = "SSL_TCP"
  cip              = "DISABLED"
  healthmonitor    = "NO"
  sitename         = citrixadc_gslbsite.site_local.sitename
}
resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
  sitepassword    = "password123"
}
`

func TestAccSslservicegroup_sslciphersuite_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslservicegroup_sslciphersuite_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroup_sslciphersuite_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_sslciphersuite_bindingExist("citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding", nil),
				),
			},
			{
				Config: testAccSslservicegroup_sslciphersuite_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_sslciphersuite_bindingNotExist("citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding", "my_gslbvservicegroup,my_ciphersuite"),
				),
			},
		},
	})
}

func testAccCheckSslservicegroup_sslciphersuite_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservicegroup_sslciphersuite_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"servicegroupname", "ciphername"}, nil)
		if err != nil {
			return err
		}
		servicegroupname := idMap["servicegroupname"]
		ciphername := idMap["ciphername"]

		findParams := service.FindParams{
			ResourceType:             "sslservicegroup_sslciphersuite_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ciphername
		found := false
		for _, v := range dataArr {
			if v["ciphername"].(string) == ciphername {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslservicegroup_sslciphersuite_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslservicegroup_sslciphersuite_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"servicegroupname", "ciphername"}, nil)
		if err != nil {
			return err
		}
		servicegroupname := idMap["servicegroupname"]
		ciphername := idMap["ciphername"]

		findParams := service.FindParams{
			ResourceType:             "sslservicegroup_sslciphersuite_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ciphername
		found := false
		for _, v := range dataArr {
			if v["ciphername"].(string) == ciphername {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslservicegroup_sslciphersuite_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSslservicegroup_sslciphersuite_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservicegroup_sslciphersuite_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"servicegroupname", "ciphername"}, nil)
		if err != nil {
			return err
		}
		servicegroupname := idMap["servicegroupname"]

		_, err = client.FindResource(service.Sslservicegroup_sslciphersuite_binding.Type(), servicegroupname)
		if err == nil {
			return fmt.Errorf("sslservicegroup_sslciphersuite_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSslservicegroup_sslciphersuite_bindingDataSource_basic = `

resource "citrixadc_sslservicegroup_sslciphersuite_binding" "tf_sslservicegroup_sslciphersuite_binding" {
	ciphername       = citrixadc_sslcipher.tf_sslcipher.ciphergroupname
	servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
}

resource "citrixadc_sslcipher" "tf_sslcipher" {
	ciphergroupname = "my_ciphersuite"
   
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		cipherpriority = 1
	}
}

resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "my_gslbvservicegroup"
	servicetype      = "SSL_TCP"
	cip              = "DISABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
}

resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
}

data "citrixadc_sslservicegroup_sslciphersuite_binding" "tf_sslservicegroup_sslciphersuite_binding" {
	servicegroupname = citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding.servicegroupname
	ciphername       = citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding.ciphername
}
`

const testAccSslservicegroup_sslciphersuite_binding_upgrade_basic = `
resource "citrixadc_sslservicegroup_sslciphersuite_binding" "tf_sslservicegroup_sslciphersuite_binding" {
	ciphername       = citrixadc_sslcipher.tf_sslcipher.ciphergroupname
	servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
}

resource "citrixadc_sslcipher" "tf_sslcipher" {
	ciphergroupname = "my_ciphersuite"

	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		cipherpriority = 1
	}
}

resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "my_gslbvservicegroup"
	servicetype      = "SSL_TCP"
	cip              = "DISABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
}

resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
}
`

// TestAccSslservicegroup_sslciphersuite_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release (legacy comma-joined id) is transparently upgraded
// by the current Framework provider. Step 1 creates the binding with citrix/citrixadc
// 2.2.0 (legacy id "servicegroupname,ciphername"); step 2 refreshes/plans the same
// config through the current Framework provider, whose Read parses the legacy id and
// recomputes it to the new "ciphername:<v>,servicegroupname:<v>" format (SetAttrFromGet).
func TestAccSslservicegroup_sslciphersuite_binding_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslservicegroup_sslciphersuite_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release, writing the legacy id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.0.0",
					},
				},
				Config: testAccSslservicegroup_sslciphersuite_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_sslciphersuite_bindingExist("citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding", "id", "my_gslbvservicegroup,my_ciphersuite"),
				),
			},
			{
				// Step 2: refresh/apply the same config through the current Framework
				// provider. Read exercises ParseIdString on the legacy id, then
				// recomputes the id to the new key:value canonical format.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSslservicegroup_sslciphersuite_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_sslciphersuite_bindingExist("citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding", "id", "ciphername:my_ciphersuite,servicegroupname:my_gslbvservicegroup"),
				),
			},
		},
	})
}

func TestAccSslservicegroup_sslciphersuite_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroup_sslciphersuite_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding", "servicegroupname", "my_gslbvservicegroup"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding", "ciphername", "my_ciphersuite"),
				),
			},
		},
	})
}

func TestAccSslservicegroup_sslciphersuite_binding_import(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	const resAddr = "citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: servicegroupname,ciphername) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"servicegroupname", "ciphername"}
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
		CheckDestroy:             testAccCheckSslservicegroup_sslciphersuite_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslservicegroup_sslciphersuite_binding_basic},
			{Config: testAccSslservicegroup_sslciphersuite_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccSslservicegroup_sslciphersuite_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccSslservicegroup_sslciphersuite_binding_selfHealing(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	const resAddr = "citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslservicegroup_sslciphersuite_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroup_sslciphersuite_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslservicegroup_sslciphersuite_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Sslservicegroup_sslciphersuite_binding.Type(), "my_gslbvservicegroup", map[string]string{"ciphername": "my_ciphersuite"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSslservicegroup_sslciphersuite_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslservicegroup_sslciphersuite_bindingExist(resAddr, nil)),
			},
		},
	})
}
