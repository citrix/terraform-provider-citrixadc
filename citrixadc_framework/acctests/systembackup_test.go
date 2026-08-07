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

const testAccSystembackup_basic = `

resource "citrixadc_systembackup" "tf_systembackup" {
	filename         = "new.tgz"
	}
  
`

func TestAccSystembackup_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystembackupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystembackup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystembackupExist("citrixadc_systembackup.tf_systembackup", nil),
				),
			},
		},
	})
}

func TestAccSystembackup_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSystembackupDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSystembackup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSystembackupExist("citrixadc_systembackup.tf_systembackup", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSystembackup_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckSystembackupExist("citrixadc_systembackup.tf_systembackup", nil)),
			},
		},
	})
}

func TestAccSystembackup_selfHealing(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	const resAddr = "citrixadc_systembackup.tf_systembackup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystembackupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystembackup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSystembackupExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Systembackup.Type(), "new.tgz"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSystembackup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSystembackupExist(resAddr, nil)),
			},
		},
	})
}

func TestAccSystembackup_import(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	const resAddr = "citrixadc_systembackup.tf_systembackup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystembackupDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSystembackup_basic},
			{
				Config:                  testAccSystembackup_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSystembackupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systembackup name is set")
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
		data, err := client.FindResource(service.Systembackup.Type(), rs.Primary.Attributes["filename"])

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("systembackup %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystembackupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systembackup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Systembackup.Type(), rs.Primary.Attributes["filename"])
		if err == nil {
			return fmt.Errorf("systembackup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystembackupDataSource_basic = `

resource "citrixadc_systembackup_create" "tf_systembackup_create" {
	filename         = "my_backup_file"
	level            = "basic"
	uselocaltimezone = "true"
}

data "citrixadc_systembackup" "tf_systembackup" {
	filename = "my_backup_file.tgz"
	depends_on = [citrixadc_systembackup_create.tf_systembackup_create]
}
`

func TestAccSystembackupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSystembackupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systembackup.tf_systembackup", "filename", "my_backup_file.tgz"),
					resource.TestCheckResourceAttr("data.citrixadc_systembackup.tf_systembackup", "level", "basic"),
				),
			},
		},
	})
}
