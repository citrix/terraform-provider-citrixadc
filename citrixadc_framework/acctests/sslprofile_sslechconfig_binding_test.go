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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Participating entities (deep dependency chain):
//   - citrixadc_sslhpkekey   (HPKE key; file staged under /nsconfig/ssl by the PreCheck)
//   - citrixadc_sslechconfig (ECH config referencing the HPKE key)
//   - citrixadc_sslprofile   (the SSL profile the ech config is bound to)
// Composite ID = name,echconfigname.
//
// The sslhpkekey "file" is testdata/hpke_key.der (X25519 PKCS#8 DER - the only
// format `add ssl hpkekey` accepts), staged under /nsconfig/ssl by the PreCheck
// (doSslhpkekeyPreChecks). dhkem = X_25519 and echcipher = AES128-GCM-HKDFSHA256
// are the NITRO-documented values. NOTE: creating citrixadc_sslprofile requires
// the default SSL profile feature enabled on the appliance.

const testAccSslprofileSslechconfigBinding_basic_step1 = `
	resource "citrixadc_sslhpkekey" "tf_sslhpkekey" {
		hpkekeyname = "tf_hpkekey_ech"
		dhkem       = "X_25519"
		file        = "hpke_key.der"
	}

	resource "citrixadc_sslechconfig" "tf_sslechconfig" {
		echconfigname = "tf_echconfig"
		echconfigid   = 1
		echpublicname = "example.com"
		echcipher     = "AES128-GCM-HKDFSHA256"
		hpkekeyname   = citrixadc_sslhpkekey.tf_sslhpkekey.hpkekeyname

		depends_on = [citrixadc_sslhpkekey.tf_sslhpkekey]
	}

	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name             = "tf_sslprofile_ech"
		ecccurvebindings = []
	}

	resource "citrixadc_sslprofile_sslechconfig_binding" "tf_sslprofile_sslechconfig_binding" {
		name          = citrixadc_sslprofile.tf_sslprofile.name
		echconfigname = citrixadc_sslechconfig.tf_sslechconfig.echconfigname

		depends_on = [
			citrixadc_sslprofile.tf_sslprofile,
			citrixadc_sslechconfig.tf_sslechconfig,
		]
	}
`

// step2 drops the binding (keeps the participating entities), verifying clean unbind.
const testAccSslprofileSslechconfigBinding_basic_step2 = `
	resource "citrixadc_sslhpkekey" "tf_sslhpkekey" {
		hpkekeyname = "tf_hpkekey_ech"
		dhkem       = "X_25519"
		file        = "hpke_key.der"
	}

	resource "citrixadc_sslechconfig" "tf_sslechconfig" {
		echconfigname = "tf_echconfig"
		echconfigid   = 1
		echpublicname = "example.com"
		echcipher     = "AES128-GCM-HKDFSHA256"
		hpkekeyname   = citrixadc_sslhpkekey.tf_sslhpkekey.hpkekeyname

		depends_on = [citrixadc_sslhpkekey.tf_sslhpkekey]
	}

	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name             = "tf_sslprofile_ech"
		ecccurvebindings = []
	}
`

func TestAccSslprofileSslechconfigBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslhpkekeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileSslechconfigBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofileSslechconfigBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileSslechconfigBindingExist("citrixadc_sslprofile_sslechconfig_binding.tf_sslprofile_sslechconfig_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile_sslechconfig_binding.tf_sslprofile_sslechconfig_binding", "name", "tf_sslprofile_ech"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile_sslechconfig_binding.tf_sslprofile_sslechconfig_binding", "echconfigname", "tf_echconfig"),
				),
			},
			{
				Config: testAccSslprofileSslechconfigBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_sslprofile", nil),
				),
			},
		},
	})
}

func testAccCheckSslprofileSslechconfigBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslprofile_sslechconfig_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}
			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// Composite ID = name,echconfigname (key:UrlEncode(value) form)
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"name", "echconfigname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Sslprofile_sslechconfig_binding.Type(),
			ResourceName:             idMap["name"],
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if en, ok := v["echconfigname"].(string); ok && en == idMap["echconfigname"] {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslprofile_sslechconfig_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslprofileSslechconfigBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslprofile_sslechconfig_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"name", "echconfigname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Sslprofile_sslechconfig_binding.Type(),
			ResourceName:             idMap["name"],
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Error (e.g. parent gone) means the binding is gone.
			continue
		}

		for _, v := range dataArr {
			if en, ok := v["echconfigname"].(string); ok && en == idMap["echconfigname"] {
				return fmt.Errorf("sslprofile_sslechconfig_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccSslprofileSslechconfigBindingDataSource_basic = `
	resource "citrixadc_sslhpkekey" "tf_sslhpkekey" {
		hpkekeyname = "tf_hpkekey_ech"
		dhkem       = "X_25519"
		file        = "hpke_key.der"
	}

	resource "citrixadc_sslechconfig" "tf_sslechconfig" {
		echconfigname = "tf_echconfig"
		echconfigid   = 1
		echpublicname = "example.com"
		echcipher     = "AES128-GCM-HKDFSHA256"
		hpkekeyname   = citrixadc_sslhpkekey.tf_sslhpkekey.hpkekeyname

		depends_on = [citrixadc_sslhpkekey.tf_sslhpkekey]
	}

	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name             = "tf_sslprofile_ech"
		ecccurvebindings = []
	}

	resource "citrixadc_sslprofile_sslechconfig_binding" "tf_sslprofile_sslechconfig_binding" {
		name          = citrixadc_sslprofile.tf_sslprofile.name
		echconfigname = citrixadc_sslechconfig.tf_sslechconfig.echconfigname

		depends_on = [
			citrixadc_sslprofile.tf_sslprofile,
			citrixadc_sslechconfig.tf_sslechconfig,
		]
	}

	data "citrixadc_sslprofile_sslechconfig_binding" "tf_sslprofile_sslechconfig_binding" {
		name          = citrixadc_sslprofile_sslechconfig_binding.tf_sslprofile_sslechconfig_binding.name
		echconfigname = citrixadc_sslprofile_sslechconfig_binding.tf_sslprofile_sslechconfig_binding.echconfigname

		depends_on = [citrixadc_sslprofile_sslechconfig_binding.tf_sslprofile_sslechconfig_binding]
	}
`

func TestAccSslprofileSslechconfigBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslhpkekeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileSslechconfigBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofileSslechconfigBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile_sslechconfig_binding.tf_sslprofile_sslechconfig_binding", "name", "tf_sslprofile_ech"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile_sslechconfig_binding.tf_sslprofile_sslechconfig_binding", "echconfigname", "tf_echconfig"),
				),
			},
		},
	})
}
