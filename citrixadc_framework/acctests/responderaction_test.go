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

	"github.com/citrix/adc-nitro-go/resource/config/responder"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResponderaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckResponderactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResponderaction_target_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderactionExist("citrixadc_responderaction.tfaction", nil),
				),
			},
			{
				Config: testAccResponderaction_target_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderactionExist("citrixadc_responderaction.tfaction", nil),
				),
			},
			{
				Config: testAccResponderaction_target_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderactionExist("citrixadc_responderaction.tfaction", nil),
				),
			},
		},
	})
}

func TestAccResponderaction_html(t *testing.T) {

	if isCpxRun {
		t.Skip("Skipping responder action html test because CPX cannot import responder html page")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doResponderactionPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckResponderactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResponderaction_html_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderactionExist("citrixadc_responderaction.tfaction2", nil),
				),
			},
			{
				Config: testAccResponderaction_html_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderactionExist("citrixadc_responderaction.tfaction2", nil),
				),
			},
		},
	})
}

func testAccCheckResponderactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No responder action exists")
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
		data, err := client.FindResource(service.Responderaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckResponderactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_responderaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Responderaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func doResponderactionPreChecks(t *testing.T) {
	testAccPreCheck(t)

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	uploads := []string{
		"error.html",
		"other_error.html",
	}
	//c := testAccProvider.Meta().(*NetScalerNitroClient)
	for _, filename := range uploads {
		err := uploadTestdataFile(c, t, filename, "/var/tmp")
		if err != nil {
			t.Errorf("%v", err)
		}
	}

	errorPage := responder.Responderhtmlpage{
		Name: "tf-error-page",
		Src:  "local://error.html",
	}
	otherErrorPage := responder.Responderhtmlpage{
		Name: "tf-other-error-page",
		Src:  "local://other_error.html",
	}
	pages := make([]responder.Responderhtmlpage, 0, 2)
	pages = append(pages, errorPage)
	pages = append(pages, otherErrorPage)

	for _, page := range pages {
		if err := c.client.ActOnResource(service.Responderhtmlpage.Type(), page, "Import"); err != nil {
			if !strings.Contains(err.Error(), "Object already exists") {
				t.Errorf("%v", err)
			}
		}
	}
}

// Create
const testAccResponderaction_target_step1 = `
resource "citrixadc_responderaction" "tfaction" {
  name    = "tfaction"
  type    = "respondwith"
  bypasssafetycheck = "YES"
  target  = "HTTP.REQ.URL.SUFFIX.EQ(\"goodbye\")"
  comment = "some comment"
}
`

// Update target to include bypasssafetycheck
const testAccResponderaction_target_step2 = `
resource "citrixadc_responderaction" "tfaction" {
  name    = "tfaction"
  type    = "respondwith"
  bypasssafetycheck = "NO"
  target  = "HTTP.REQ.URL.SUFFIX.EQ(\"hello\")"
  comment = "some comment"
}
`

// Update irrelevant field comment to check non inclusion of bypasssafetycheck
const testAccResponderaction_target_step3 = `
resource "citrixadc_responderaction" "tfaction" {
  name    = "tfaction"
  type    = "respondwith"
  bypasssafetycheck = "YES"
  target  = "HTTP.REQ.URL.SUFFIX.EQ(\"hello\")"
  comment = "other comment"
}
`

// Initial html response action
const testAccResponderaction_html_step1 = `
resource "citrixadc_responderaction" "tfaction2" {
  name    = "tfaction2"
  type    = "respondwithhtmlpage"
  htmlpage = "tf-error-page"
  comment = "some comment"
  reasonphrase = "HTTP.REQ.URL"
  responsestatuscode = 202
}
`

// Update html response action

const testAccResponderaction_html_step2 = `
resource "citrixadc_responderaction" "tfaction2" {
  name    = "tfaction2"
  type    = "respondwithhtmlpage"
  htmlpage = "tf-other-error-page"
  comment = "some other comment"
  reasonphrase = "HTTP.REQ.URL.SUFFIX.EQ(\"goodbye1\")"
  responsestatuscode = 201
}
`

func TestAccResponderaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_responderaction.tfaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckResponderactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResponderaction_target_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckResponderactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Responderaction.Type(), "tfaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccResponderaction_target_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckResponderactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccResponderaction_import(t *testing.T) {
	const resAddr = "citrixadc_responderaction.tfaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckResponderactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccResponderaction_target_step1},
			{
				Config:                  testAccResponderaction_target_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bypasssafetycheck"},
			},
		},
	})
}

// The responderaction unset test uses type "redirect", the only type that
// accepts comment, reasonphrase and responsestatuscode together (respondwith
// rejects responsestatuscode/headers; reasonphrase has a min length of 1). Each
// of these Optional+Computed attributes is wired for NITRO ?action=unset via an
// unsetOnRemove plan modifier: removing it from config reverts it on the
// appliance (GET no longer returns it).
const testAccResponderaction_unset_step1 = `
resource "citrixadc_responderaction" "tf_unset" {
  name               = "tf_responderaction_unset"
  type               = "redirect"
  target             = "\"http://backupsite.com\" + HTTP.REQ.URL"
  comment            = "unset test comment"
  reasonphrase       = "\"Moved\""
  responsestatuscode = 307
}
`

const testAccResponderaction_unset_step2 = `
resource "citrixadc_responderaction" "tf_unset" {
  name   = "tf_responderaction_unset"
  type   = "redirect"
  target = "\"http://backupsite.com\" + HTTP.REQ.URL"
  # comment, reasonphrase and responsestatuscode removed from config -> the
  # provider must unset them (revert to NITRO defaults / absent).
}
`

func TestAccResponderaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckResponderactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccResponderaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderactionExist("citrixadc_responderaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_responderaction.tf_unset", "comment", "unset test comment"),
					resource.TestCheckResourceAttr("citrixadc_responderaction.tf_unset", "reasonphrase", "\"Moved\""),
					resource.TestCheckResourceAttr("citrixadc_responderaction.tf_unset", "responsestatuscode", "307"),
				),
			},
			{
				// Removing the attributes must unset them: the appliance no longer
				// returns them, and the implicit post-apply plan must be empty.
				Config: testAccResponderaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderactionExist("citrixadc_responderaction.tf_unset", nil),
					// Independent appliance-level confirmation the unset took effect
					// (unset reverts these to absent on the appliance).
					testAccCheckResponderactionADCValue("tf_responderaction_unset", "comment", ""),
					testAccCheckResponderactionADCValue("tf_responderaction_unset", "reasonphrase", ""),
					testAccCheckResponderactionADCValue("tf_responderaction_unset", "responsestatuscode", ""),
				),
			},
		},
	})
}

// testAccCheckResponderactionADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it. want "" means the attribute is expected to be absent/empty.
func testAccCheckResponderactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Responderaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("responderaction %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("responderaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccResponderaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckResponderactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccResponderaction_target_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckResponderactionExist("citrixadc_responderaction.tfaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccResponderaction_target_step1,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckResponderactionExist("citrixadc_responderaction.tfaction", nil)),
			},
		},
	})
}

const testAccResponderactionDataSource_basic = `
resource "citrixadc_responderaction" "tfaction_ds" {
  name    = "tfaction_ds"
  type    = "respondwith"
  target  = "\"test_response\""
  comment = "datasource test comment"
}

data "citrixadc_responderaction" "tfaction_ds" {
  name = citrixadc_responderaction.tfaction_ds.name
}
`

func TestAccResponderactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResponderactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_responderaction.tfaction_ds", "name", "tfaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_responderaction.tfaction_ds", "type", "respondwith"),
					resource.TestCheckResourceAttr("data.citrixadc_responderaction.tfaction_ds", "target", "\"test_response\""),
					resource.TestCheckResourceAttr("data.citrixadc_responderaction.tfaction_ds", "comment", "datasource test comment"),
				),
			},
		},
	})
}
