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
	"os"
	"testing"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGslbsite_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGslbsiteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbsite_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbsiteExist("citrixadc_gslbsite.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_gslbsite.foo", "siteipaddress", "172.31.11.20"),
					resource.TestCheckResourceAttr(
						"citrixadc_gslbsite.foo", "sitename", "Site-GSLB-East-Coast"),
				),
			},
		},
	})
}

func testAccCheckGslbsiteExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Gslbsite.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbsiteDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbsite" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Gslbsite.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccGslbsite_basic = `


resource "citrixadc_gslbsite" "foo" {

  siteipaddress = "172.31.11.20"
  sitename = "Site-GSLB-East-Coast"
  sitepassword = "password123"

}
`

func TestAccGslbsite_AssertNonUpdateableAttributes(t *testing.T) {
	t.Skip("TODO: The GSLB site does not exist")

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	siteName := "tf-acc-glsb-site-test"
	siteType := service.Gslbsite.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, siteType, siteName, nil)

	siteInstance := gslb.Gslbsite{
		Sitename:      siteName,
		Siteipaddress: "12.12.12.12",
	}

	if _, err := c.client.AddResource(siteType, siteName, siteInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}
	// Zero out fields of created instance
	siteInstance.Siteipaddress = ""

	// sitetype
	siteInstance.Sitetype = "HTTP"
	testHelperVerifyImmutabilityFunc(c, t, siteType, siteName, siteInstance, "sitetype")
	siteInstance.Sitetype = ""

	// siteipaddress
	siteInstance.Siteipaddress = "12.22.22.22"
	testHelperVerifyImmutabilityFunc(c, t, siteType, siteName, siteInstance, "siteipaddress")
	siteInstance.Siteipaddress = ""

	// publicip
	siteInstance.Publicip = "1.2.3.4"
	testHelperVerifyImmutabilityFunc(c, t, siteType, siteName, siteInstance, "publicip")
	siteInstance.Publicip = ""

	// parentsite
	siteInstance.Parentsite = "some_site"
	testHelperVerifyImmutabilityFunc(c, t, siteType, siteName, siteInstance, "parentsite")
	siteInstance.Parentsite = ""

	// clip
	siteInstance.Clip = "some_clip"
	testHelperVerifyImmutabilityFunc(c, t, siteType, siteName, siteInstance, "clip")
	siteInstance.Clip = ""

	// publicclip
	siteInstance.Publicclip = "some_clip"
	testHelperVerifyImmutabilityFunc(c, t, siteType, siteName, siteInstance, "publicclip")
	siteInstance.Publicclip = ""
}
