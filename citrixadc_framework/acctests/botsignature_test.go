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

const testAccBotsignature_basic = `
	resource "citrixadc_systemfile" "tf_signature" {
		filename     = "bot_signature.json"
		filelocation = "/var/tmp"
		filecontent  = file("testdata/bot_signatures.json")
	}
	resource "citrixadc_botsignature" "tf_botsignature" {
		name       = "tf_botsignature"
		src        = "local://bot_signature.json"
		depends_on = [citrixadc_systemfile.tf_signature]
		comment    = "TestingExample"
	}
`

func TestAccBotsignature_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotsignatureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotsignature_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsignatureExist("citrixadc_botsignature.tf_botsignature", nil),
				),
			},
		},
	})
}

func testAccCheckBotsignatureExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No botsignature name is set")
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
		data, err := client.FindResource("botsignature", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("botsignature %s not found", n)
		}

		return nil
	}
}

func testAccCheckBotsignatureDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_botsignature" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("botsignature", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("botsignature %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccBotsignature_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_botsignature.tf_botsignature"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotsignatureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotsignature_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotsignatureExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Botsignature.Type(), "tf_botsignature"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccBotsignature_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotsignatureExist(resAddr, nil)),
			},
		},
	})
}

func TestAccBotsignature_import(t *testing.T) {
	const resAddr = "citrixadc_botsignature.tf_botsignature"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotsignatureDestroy,
		Steps: []resource.TestStep{
			{Config: testAccBotsignature_basic},
			{
				Config:                  testAccBotsignature_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"comment", "src"},
			},
		},
	})
}

func TestAccBotsignature_sdkv2StateUpgrade(t *testing.T) {
	// No upgrade baseline possible: the released citrix/citrixadc 2.2.0 provider
	// rejects this botsignature config at step 1 ("Invalid function argument" on
	// the signature-file import), so SDK-v2 state is never established. The current
	// provider is unaffected.
	t.Skip("no 2.2.0 baseline: released 2.2.0 provider errors on botsignature step 1 (invalid function argument)")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckBotsignatureDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccBotsignature_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotsignatureExist("citrixadc_botsignature.tf_botsignature", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccBotsignature_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckBotsignatureExist("citrixadc_botsignature.tf_botsignature", nil)),
			},
		},
	})
}

const testAccBotsignatureDataSource_basic = `
	resource "citrixadc_systemfile" "tf_signature_ds" {
		filename     = "bot_signature_ds.json"
		filelocation = "/var/tmp"
		filecontent  = file("testdata/bot_signatures.json")
	}
	resource "citrixadc_botsignature" "tf_botsignature_ds" {
		name       = "tf_botsignature_ds"
		src        = "local://bot_signature_ds.json"
		depends_on = [citrixadc_systemfile.tf_signature_ds]
		comment    = "TestingDataSource"
	}
	data "citrixadc_botsignature" "tf_botsignature_ds_data" {
		name       = citrixadc_botsignature.tf_botsignature_ds.name
		depends_on = [citrixadc_botsignature.tf_botsignature_ds]
	}
`

func TestAccBotsignatureDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBotsignatureDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_botsignature.tf_botsignature_ds_data", "name", "tf_botsignature_ds"),
					resource.TestCheckResourceAttrSet("data.citrixadc_botsignature.tf_botsignature_ds_data", "id"),
				),
			},
		},
	})
}
