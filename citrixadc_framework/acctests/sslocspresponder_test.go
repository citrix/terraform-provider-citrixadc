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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccSslocspresponder_basic = `
	resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = "tf_sslocspresponder"
		url = "http://www.citrix.com"
		batchingdelay = 5
		batchingdepth = 2
		cache = "ENABLED"
		cachetimeout = 1
		httpmethod = "GET"
		insertclientcert = "YES"
		ocspurlresolvetimeout = 100
		producedattimeskew = 300
		resptimeout = 100
		trustresponder = false
		usenonce = "NO"
	}
`

const testAccSslocspresponder_basic_update1 = `
	resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = "tf_sslocspresponder"
		url = "http://www.google.com"
		batchingdelay = 6
		batchingdepth = 3
		cache = "DISABLED"
		httpmethod = "POST"
		insertclientcert = "NO"
		ocspurlresolvetimeout = 101
		producedattimeskew = 301
		resptimeout = 101
		trustresponder = true
		usenonce = "YES"
		cachetimeout = 0
	}
`

const testAccSslocspresponder_basic_update2 = `
	resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = "tf_sslocspresponder"
		url = "http://www.google.com"
		batchingdelay = 6
		batchingdepth = 3
		cache = "DISABLED"
		httpmethod = "POST"
		insertclientcert = "NO"
		ocspurlresolvetimeout = 101
		producedattimeskew = 301
		respondercert = "ns-server-certificate"
		resptimeout = 101
		signingcert = "ns-server-certificate"
		trustresponder = false
		usenonce = "YES"
	}
`

const testAccSslocspresponderDataSource_basic = `
	resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = "tf_sslocspresponder"
		url = "http://www.citrix.com"
		batchingdelay = 5
		batchingdepth = 2
		cache = "ENABLED"
		cachetimeout = 1
		httpmethod = "GET"
		insertclientcert = "YES"
		ocspurlresolvetimeout = 100
		producedattimeskew = 300
		resptimeout = 100
		trustresponder = false
		usenonce = "NO"
	}

	data "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = citrixadc_sslocspresponder.tf_sslocspresponder.name
	}
`

func TestAccSslocspresponder_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslocspresponderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslocspresponder_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslocspresponderExist("citrixadc_sslocspresponder.tf_sslocspresponder", nil),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "url", "http://www.citrix.com"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "batchingdelay", "5"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "batchingdepth", "2"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "cache", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "cachetimeout", "1"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "httpmethod", "GET"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "insertclientcert", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "ocspurlresolvetimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "producedattimeskew", "300"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "resptimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "trustresponder", "false"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "usenonce", "NO"),
				),
			},
			{
				Config: testAccSslocspresponder_basic_update1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslocspresponderExist("citrixadc_sslocspresponder.tf_sslocspresponder", nil),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "url", "http://www.google.com"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "batchingdelay", "6"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "batchingdepth", "3"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "cache", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "cachetimeout", "0"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "httpmethod", "POST"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "insertclientcert", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "ocspurlresolvetimeout", "101"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "producedattimeskew", "301"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "resptimeout", "101"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "trustresponder", "true"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "usenonce", "YES"),
				),
			},
			{
				Config: testAccSslocspresponder_basic_update2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslocspresponderExist("citrixadc_sslocspresponder.tf_sslocspresponder", nil),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "respondercert", "ns-server-certificate"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "signingcert", "ns-server-certificate"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_sslocspresponder", "trustresponder", "false"),
				),
			},
		},
	})
}

func testAccCheckSslocspresponderExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslocspresponder name is set")
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
		data, err := client.FindResource(service.Sslocspresponder.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslocspresponder %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslocspresponderDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslocspresponder" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslocspresponder.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslocspresponder %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslocspresponder_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_sslocspresponder.tf_sslocspresponder"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslocspresponderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslocspresponder_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslocspresponderExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Sslocspresponder.Type(), "tf_sslocspresponder"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSslocspresponder_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslocspresponderExist(resAddr, nil)),
			},
		},
	})
}

func TestAccSslocspresponder_import(t *testing.T) {
	const resAddr = "citrixadc_sslocspresponder.tf_sslocspresponder"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslocspresponderDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslocspresponder_basic},
			{
				Config:                  testAccSslocspresponder_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSslocspresponder_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslocspresponderDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSslocspresponder_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslocspresponderExist("citrixadc_sslocspresponder.tf_sslocspresponder", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSslocspresponder_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckSslocspresponderExist("citrixadc_sslocspresponder.tf_sslocspresponder", nil)),
			},
		},
	})
}

// testAccSslocspresponder_unset_step1 sets the unset-eligible attributes to
// valid NON-default values. testAccSslocspresponder_unset_step2 removes them
// (keeping only the required name + url), so the provider must unset them and
// the appliance must revert them to their NITRO defaults.
const testAccSslocspresponder_unset_step1 = `
	resource "citrixadc_sslocspresponder" "tf_unset" {
		name               = "tf_sslocspresponder_unset"
		url                = "http://www.citrix.com"
		cache              = "ENABLED"
		cachetimeout       = 100
		httpmethod         = "GET"
		producedattimeskew = 600
		trustresponder     = true
	}
`

const testAccSslocspresponder_unset_step2 = `
	resource "citrixadc_sslocspresponder" "tf_unset" {
		name = "tf_sslocspresponder_unset"
		url  = "http://www.citrix.com"
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccSslocspresponder_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslocspresponderDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSslocspresponder_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslocspresponderExist("citrixadc_sslocspresponder.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_unset", "cache", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_unset", "cachetimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_unset", "httpmethod", "GET"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_unset", "producedattimeskew", "600"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_unset", "trustresponder", "true"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSslocspresponder_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslocspresponderExist("citrixadc_sslocspresponder.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_unset", "cache", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_unset", "cachetimeout", "1"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_unset", "httpmethod", "POST"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_unset", "producedattimeskew", "300"),
					resource.TestCheckResourceAttr("citrixadc_sslocspresponder.tf_unset", "trustresponder", "false"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSslocspresponderADCValue("tf_sslocspresponder_unset", "cache", "DISABLED"),
					testAccCheckSslocspresponderADCValue("tf_sslocspresponder_unset", "httpmethod", "POST"),
					testAccCheckSslocspresponderADCValue("tf_sslocspresponder_unset", "producedattimeskew", "300"),
				),
			},
		},
	})
}

// testAccCheckSslocspresponderADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckSslocspresponderADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Sslocspresponder.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("sslocspresponder %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("sslocspresponder %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccSslocspresponderDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslocspresponderDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslocspresponder.tf_sslocspresponder", "name", "tf_sslocspresponder"),
					resource.TestCheckResourceAttr("data.citrixadc_sslocspresponder.tf_sslocspresponder", "url", "http://www.citrix.com"),
					resource.TestCheckResourceAttr("data.citrixadc_sslocspresponder.tf_sslocspresponder", "batchingdelay", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_sslocspresponder.tf_sslocspresponder", "batchingdepth", "2"),
					resource.TestCheckResourceAttr("data.citrixadc_sslocspresponder.tf_sslocspresponder", "cache", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslocspresponder.tf_sslocspresponder", "cachetimeout", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_sslocspresponder.tf_sslocspresponder", "httpmethod", "GET"),
					resource.TestCheckResourceAttr("data.citrixadc_sslocspresponder.tf_sslocspresponder", "insertclientcert", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_sslocspresponder.tf_sslocspresponder", "usenonce", "NO"),
				),
			},
		},
	})
}
