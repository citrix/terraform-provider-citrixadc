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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccFeoaction_basic = `

	resource "citrixadc_feoaction" "tf_feoaction" {
		name              = "my_feoaction"
		cachemaxage       = 50
		imgshrinktoattrib = "false"
		imggiftopng       = "false"
	}
`
const testAccFeoaction_update = `

	resource "citrixadc_feoaction" "tf_feoaction" {
		name              = "my_feoaction"
		cachemaxage       = 40
		imgshrinktoattrib = "true"
		imggiftopng       = "true"
	}
`

func TestAccFeoaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFeoactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFeoaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoactionExist("citrixadc_feoaction.tf_feoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "name", "my_feoaction"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "cachemaxage", "50"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "imgshrinktoattrib", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "imggiftopng", "false"),
				),
			},
			{
				Config: testAccFeoaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoactionExist("citrixadc_feoaction.tf_feoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "name", "my_feoaction"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "cachemaxage", "40"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "imgshrinktoattrib", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "imggiftopng", "true"),
				),
			},
		},
	})
}

func TestAccFeoaction_import(t *testing.T) {
	const resAddr = "citrixadc_feoaction.tf_feoaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFeoactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccFeoaction_basic},
			{
				Config:                  testAccFeoaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckFeoactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No feoaction name is set")
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
		data, err := client.FindResource("feoaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("feoaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckFeoactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_feoaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("feoaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("feoaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccFeoactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccFeoactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "name", "tf_feoaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "cachemaxage", "60"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "imgshrinktoattrib", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "imggiftopng", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "cssminify", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "jsminify", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "htmlminify", "true"),
				),
			},
		},
	})
}

func TestAccFeoaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_feoaction.tf_feoaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFeoactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFeoaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckFeoactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Feoaction.Type(), "my_feoaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccFeoaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckFeoactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccFeoaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckFeoactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccFeoaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckFeoactionExist("citrixadc_feoaction.tf_feoaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccFeoaction_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckFeoactionExist("citrixadc_feoaction.tf_feoaction", nil)),
			},
		},
	})
}

const testAccFeoaction_unset_step1 = `
resource "citrixadc_feoaction" "tf_unset" {
  name                = "tf_test_feoaction_unset"
  pageextendcache     = "true"
  imgshrinktoattrib   = "true"
  imggiftopng         = "true"
  imgtowebp           = "true"
  imgtojpegxr         = "true"
  imginline           = "true"
  cssimginline        = "true"
  jpgoptimize         = "true"
  imglazyload         = "true"
  cssminify           = "true"
  cssinline           = "true"
  csscombine          = "true"
  convertimporttolink = "true"
  jsminify            = "true"
  jsinline            = "true"
  htmlminify          = "true"
  cssmovetohead       = "true"
  jsmovetoend         = "true"
}
`

const testAccFeoaction_unset_step2 = `
resource "citrixadc_feoaction" "tf_unset" {
  name = "tf_test_feoaction_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults, "false").
}
`

func TestAccFeoaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFeoactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccFeoaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoactionExist("citrixadc_feoaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "pageextendcache", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imgshrinktoattrib", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imggiftopng", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imgtowebp", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imgtojpegxr", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imginline", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "cssimginline", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "jpgoptimize", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imglazyload", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "cssminify", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "cssinline", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "csscombine", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "convertimporttolink", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "jsminify", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "jsinline", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "htmlminify", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "cssmovetohead", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "jsmovetoend", "true"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccFeoaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoactionExist("citrixadc_feoaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "pageextendcache", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imgshrinktoattrib", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imggiftopng", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imgtowebp", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imgtojpegxr", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imginline", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "cssimginline", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "jpgoptimize", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "imglazyload", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "cssminify", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "cssinline", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "csscombine", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "convertimporttolink", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "jsminify", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "jsinline", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "htmlminify", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "cssmovetohead", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_unset", "jsmovetoend", "false"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckFeoactionADCValue("tf_test_feoaction_unset", "imgshrinktoattrib", "false"),
					testAccCheckFeoactionADCValue("tf_test_feoaction_unset", "cssminify", "false"),
					testAccCheckFeoactionADCValue("tf_test_feoaction_unset", "jsminify", "false"),
					testAccCheckFeoactionADCValue("tf_test_feoaction_unset", "htmlminify", "false"),
				),
			},
		},
	})
}

// testAccCheckFeoactionADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckFeoactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Feoaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("feoaction %s not found on appliance", name)
		}
		got := fmt.Sprintf("%v", data[attr])
		if got != want {
			return fmt.Errorf("feoaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccFeoactionDataSource_basic = `

resource "citrixadc_feoaction" "tf_feoaction_ds" {
    name              = "tf_feoaction_ds"
    cachemaxage       = 60
    imgshrinktoattrib = "true"
    imggiftopng       = "true"
    cssminify         = "true"
    jsminify          = "true"
    htmlminify        = "true"
}

data "citrixadc_feoaction" "tf_feoaction_ds" {
    name = citrixadc_feoaction.tf_feoaction_ds.name
    depends_on = [citrixadc_feoaction.tf_feoaction_ds]
}

`
