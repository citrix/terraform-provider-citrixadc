/*
Copyright 2024 Citrix Systems, Inc

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
	"log"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccGslbservicegroup_gslbservicegroupmember_binding_basic = `

	resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
		servicegroupname = "test_gslbvservicegroup"
		servicetype      = "HTTP"
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
	resource "citrixadc_server" "tf_server" {
		name = "tf_server"
		ipaddress = "192.168.11.13"
	}
	
	resource "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "tf_binding" {
		servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
		servername       = citrixadc_server.tf_server.name
		port             = 60
	}
	
`

const testAccGslbservicegroup_gslbservicegroupmember_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
		servicegroupname = "test_gslbvservicegroup"
		servicetype      = "HTTP"
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
	resource "citrixadc_server" "tf_server" {
		name = "tf_server"
		ipaddress = "192.168.11.13"
	}
`

func TestAccGslbservicegroup_gslbservicegroupmember_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbservicegroup_gslbservicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_gslbservicegroupmember_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_gslbservicegroupmember_bindingExist("citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccGslbservicegroup_gslbservicegroupmember_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_gslbservicegroupmember_bindingNotExist("citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "test_gslbvservicegroup,10.10.10.10,60"),
				),
			},
		},
	})
}

func TestAccGslbservicegroup_gslbservicegroupmember_binding_import(t *testing.T) {
	const resAddr = "citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: servicegroupname,servername,ip,port) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"servicegroupname", "servername", "ip", "port"}
		parts := make([]string, 0, len(ordr))
		for _, k := range ordr {
			// Skip empty values: servername and ip are mutually exclusive, so the
			// canonical id always carries an empty slot for the unused one. SDK v2
			// wrote a 3-token positional id ("servicegroupname,<servername-or-ip>,port"),
			// never a 4-token id with an empty middle slot, so exclude empties here.
			if v, ok := kv[k]; ok && v != "" {
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
		CheckDestroy:             testAccCheckGslbservicegroup_gslbservicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_gslbservicegroupmember_binding_basic,
			},
			{
				Config:            testAccGslbservicegroup_gslbservicegroupmember_binding_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// The resource Read/SetAttrFromGet preserves identity fields from prior
				// state and does not repopulate them from the ID during import, so these
				// RequiresReplace identity attributes cannot round-trip on a bare import.
				ImportStateVerifyIgnore: []string{"port", "servername", "servicegroupname"},
			},
			{Config: testAccGslbservicegroup_gslbservicegroupmember_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"port", "servername", "servicegroupname"}},
		},
	})
}

// TestAccGslbservicegroup_gslbservicegroupmember_binding_selfHealing verifies drift
// recovery: after the binding is deleted out-of-band on the ADC, the next refresh's Read
// must detect it is gone (data.Id -> null) and drop it from state so the same config
// recreates it.
func TestAccGslbservicegroup_gslbservicegroupmember_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbservicegroup_gslbservicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_gslbservicegroupmember_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbservicegroup_gslbservicegroupmember_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: failed to get client: %v", err)
					}
					// Out-of-band delete of just the bound member (parent servicegroup remains).
					if err := client.DeleteResourceWithArgs(service.Gslbservicegroup_gslbservicegroupmember_binding.Type(), "test_gslbvservicegroup", []string{"servername:tf_server", "port:60"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccGslbservicegroup_gslbservicegroupmember_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbservicegroup_gslbservicegroupmember_bindingExist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckGslbservicegroup_gslbservicegroupmember_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbservicegroup_gslbservicegroupmember_binding id is set")
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
		idMap, _, err := utils.ParseIdString(bindingId, []string{"servicegroupname", "servername", "ip", "port"}, []string{"servername", "ip", "port"})
		if err != nil {
			return err
		}
		servicegroupname := idMap["servicegroupname"]

		servername := idMap["servername"]

		port := 0
		if portStr, ok := idMap["port"]; ok && portStr != "" {
			if port, err = strconv.Atoi(portStr); err != nil {
				return err
			}
		}

		findParams := service.FindParams{
			ResourceType:             "gslbservicegroup_gslbservicegroupmember_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if port != 0 {
				portEqual := int(v["port"].(float64)) == port
				servernameEqual := v["servername"] == servername
				if servernameEqual && portEqual {
					foundIndex = i
					break
				}
			} else {
				log.Printf("[DEBUG] teh val sis  %v, %v", v["servername"].(string), servername)
				if v["servername"].(string) == servername {
					foundIndex = i
					break
				}
			}
			log.Printf("[DEBUG] teh val sis  %v, %v", v["servername"].(string), servername)
		}

		if foundIndex == -1 {
			return fmt.Errorf("gslbservicegroup_gslbservicegroupmember_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbservicegroup_gslbservicegroupmember_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}

		idSlice := strings.SplitN(id, ",", 3)
		servicegroupname := idSlice[0]

		servername := idSlice[1]

		port := 0
		if len(idSlice) == 3 {
			if port, err = strconv.Atoi(idSlice[2]); err != nil {
				return err
			}
		}

		findParams := service.FindParams{
			ResourceType:             "gslbservicegroup_gslbservicegroupmember_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if port != 0 {
				portEqual := int(v["port"].(float64)) == port
				servernameEqual := v["servername"] == servername
				if servernameEqual && portEqual {
					foundIndex = i
					break
				}
			}
		}

		if foundIndex != -1 {
			return fmt.Errorf("servicegroup_servicegroupmember_binding %s found. Should have been deleted", id)
		}

		return nil
	}
}

func testAccCheckGslbservicegroup_gslbservicegroupmember_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("gslbservicegroup_gslbservicegroupmember_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbservicegroup_gslbservicegroupmember_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccGslbservicegroup_gslbservicegroupmember_bindingDataSource_basic = `

	resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
		servicegroupname = "test_gslbvservicegroup"
		servicetype      = "HTTP"
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
	resource "citrixadc_server" "tf_server" {
		name = "tf_server"
		ipaddress = "192.168.11.13"
	}
	
	resource "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "tf_binding" {
		servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
		servername       = citrixadc_server.tf_server.name
		port             = 60
	}

	data "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "tf_binding" {
		servicegroupname = citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding.servicegroupname
		servername       = citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding.servername
		port             = citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding.port
	}
`

func TestAccGslbservicegroup_gslbservicegroupmember_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_gslbservicegroupmember_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "servicegroupname", "test_gslbvservicegroup"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "servername", "tf_server"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "port", "60"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "ip", "192.168.11.13"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// SDK v2 -> Framework state-upgrade tests.
//
// The last SDK v2 release (citrix/citrixadc 2.2.0) wrote a legacy 3-part
// positional id "servicegroupname,<servername-or-ip>,port". These tests verify
// the current Framework provider upgrades that state cleanly: it parses the
// legacy id, locates the member, rewrites the id to the new key:value form, and
// produces an empty follow-up plan (no spurious replace / NITRO 273). Both the
// servername-bound and ip-bound (ADC auto-names server == ip) paths are covered.
// ---------------------------------------------------------------------------

const testAccGslbservicegroup_gslbservicegroupmember_binding_upgrade_servername = `
	resource "citrixadc_gslbsite" "site_local" {
		sitename        = "Site-Local"
		siteipaddress   = "172.31.96.234"
		sessionexchange = "DISABLED"
		sitepassword    = "password123"
	}
	resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
		servicegroupname = "test_gslbvservicegroup"
		servicetype      = "HTTP"
		cip              = "DISABLED"
		healthmonitor    = "NO"
		sitename         = citrixadc_gslbsite.site_local.sitename
	}
	resource "citrixadc_server" "tf_server" {
		name      = "tf_server"
		ipaddress = "192.168.11.13"
	}
	resource "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "tf_binding" {
		servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
		servername       = citrixadc_server.tf_server.name
		port             = 60
	}
`

func TestAccGslbservicegroup_gslbservicegroupmember_binding_sdkv2StateUpgrade(t *testing.T) {
	const resAddr = "citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGslbservicegroup_gslbservicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release -> legacy 3-part id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccGslbservicegroup_gslbservicegroupmember_binding_upgrade_servername,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resAddr, "id", "test_gslbvservicegroup,tf_server,60"),
				),
			},
			{
				// Step 2: manage the SAME config with the current Framework provider.
				// Read parses the legacy id, locates the member, and rewrites the id to
				// the new key:value form. Plan must be empty (no replace / 273).
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccGslbservicegroup_gslbservicegroupmember_binding_upgrade_servername,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_gslbservicegroupmember_bindingExist(resAddr, nil),
					resource.TestCheckResourceAttr(resAddr, "id", "ip:,port:60,servername:tf_server,servicegroupname:test_gslbvservicegroup"),
					resource.TestCheckResourceAttr(resAddr, "servername", "tf_server"),
					resource.TestCheckResourceAttr(resAddr, "port", "60"),
				),
			},
		},
	})
}

const testAccGslbservicegroup_gslbservicegroupmember_binding_upgrade_ip = `
	resource "citrixadc_gslbsite" "site_local" {
		sitename        = "Site-Local"
		siteipaddress   = "172.31.96.234"
		sessionexchange = "DISABLED"
		sitepassword    = "password123"
	}
	resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
		servicegroupname = "test_gslbvservicegroup"
		servicetype      = "HTTP"
		cip              = "DISABLED"
		healthmonitor    = "NO"
		sitename         = citrixadc_gslbsite.site_local.sitename
	}
	resource "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "tf_binding" {
		servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
		ip               = "192.168.11.14"
		port             = 61
	}
`

func TestAccGslbservicegroup_gslbservicegroupmember_binding_sdkv2StateUpgrade_ipBound(t *testing.T) {
	const resAddr = "citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGslbservicegroup_gslbservicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create by IP with the last SDK v2 release. ADC auto-names the
				// server == ip; the legacy id is "servicegroupname,<ip>,port".
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccGslbservicegroup_gslbservicegroupmember_binding_upgrade_ip,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resAddr, "id", "test_gslbvservicegroup,192.168.11.14,61"),
				),
			},
			{
				// Step 2: manage the SAME config with the current Framework provider.
				// The legacy token equals the member's ip, so Read resolves it to ip
				// (servername null) and rewrites the id. Plan must be empty.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccGslbservicegroup_gslbservicegroupmember_binding_upgrade_ip,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbmemberExistByIp(resAddr, "192.168.11.14", 61),
					resource.TestCheckResourceAttr(resAddr, "id", "ip:192.168.11.14,port:61,servername:,servicegroupname:test_gslbvservicegroup"),
					resource.TestCheckResourceAttr(resAddr, "ip", "192.168.11.14"),
					resource.TestCheckResourceAttr(resAddr, "port", "61"),
				),
			},
		},
	})
}

// testAccCheckGslbmemberExistByIp verifies (via the NITRO client) that a member
// with the given ip and port exists under the binding's servicegroupname. Used by
// the ip-bound upgrade test, where the canonical id carries an empty servername
// segment and the generic servername-based Exist helper would not match.
func testAccCheckGslbmemberExistByIp(n, ip string, port int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbservicegroup_gslbservicegroupmember_binding id is set")
		}
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"servicegroupname", "servername", "ip", "port"}, []string{"servername", "ip", "port"})
		if err != nil {
			return err
		}
		findParams := service.FindParams{
			ResourceType:             "gslbservicegroup_gslbservicegroupmember_binding",
			ResourceName:             idMap["servicegroupname"],
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		for _, v := range dataArr {
			portEqual := int(v["port"].(float64)) == port
			ipEqual := v["ip"] == ip || v["servername"] == ip
			if portEqual && ipEqual {
				return nil
			}
		}
		return fmt.Errorf("gslbservicegroup_gslbservicegroupmember_binding ip=%s port=%d not found", ip, port)
	}
}
