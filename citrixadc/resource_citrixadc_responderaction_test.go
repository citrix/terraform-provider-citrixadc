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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResponderaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResponderactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResponderaction_target_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderactionExist("citrixadc_responderaction.tfaction", nil),
				),
			},
			resource.TestStep{
				Config: testAccResponderaction_target_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderactionExist("citrixadc_responderaction.tfaction", nil),
				),
			},
			resource.TestStep{
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
		PreCheck:     func() { doResponderactionPreChecks(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResponderactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResponderaction_html_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderactionExist("citrixadc_responderaction.tfaction2", nil),
				),
			},
			resource.TestStep{
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

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Responderaction.Type(), rs.Primary.ID)

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
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_responderaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Responderaction.Type(), rs.Primary.ID)
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
			t.Errorf(err.Error())
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
				t.Errorf(err.Error())
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
