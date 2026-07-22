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

// PREREQUISITE / TODO_PLACEHOLDER: The Cloud Tunnel feature (cloudtunnelvserver)
// may be license/feature gated on the testbed. On an unlicensed ADC the NITRO
// add/get calls can fail with "Feature not supported in this release" or
// "not licensed". Ensure the cloud-tunnel feature/license is enabled on the
// target appliance before running these tests; otherwise Create/Read will fail
// through no fault of the test configuration.
//
// TODO_PLACEHOLDER: `listenpolicy` is a NetScaler policy EXPRESSION. Step 1 uses
// the safe default "none". Step 2 uses a simple client-IP expression
// ("CLIENT.IP.SRC.EQ(1.1.1.1)"). If the appliance rejects this expression for a
// cloud tunnel vserver, replace it with a valid named/default expression.

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccCloudtunnelvserver_basic_step1 = `
resource "citrixadc_cloudtunnelvserver" "tf_cloudtunnelvserver" {
  name           = "tf_cloudtunnelvserver"
  servicetype    = "TCP"
  listenpolicy   = "none"
  listenpriority = 50
}

`

const testAccCloudtunnelvserver_basic_step2 = `
resource "citrixadc_cloudtunnelvserver" "tf_cloudtunnelvserver" {
  name           = "tf_cloudtunnelvserver"
  servicetype    = "TCP"
  listenpolicy   = "CLIENT.IP.SRC.EQ(1.1.1.1)"
  listenpriority = 80
}

`

func TestAccCloudtunnelvserver_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudtunnelvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudtunnelvserver_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudtunnelvserverExist("citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "name", "tf_cloudtunnelvserver"),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "servicetype", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "listenpolicy", "none"),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "listenpriority", "50"),
				),
			},
			{
				Config: testAccCloudtunnelvserver_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudtunnelvserverExist("citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "name", "tf_cloudtunnelvserver"),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "servicetype", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "listenpolicy", "CLIENT.IP.SRC.EQ(1.1.1.1)"),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "listenpriority", "80"),
				),
			},
		},
	})
}

func TestAccCloudtunnelvserver_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	const resAddr = "citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudtunnelvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudtunnelvserver_basic_step1,
			},
			{
				Config:                  testAccCloudtunnelvserver_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckCloudtunnelvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cloudtunnelvserver name is set")
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
		data, err := client.FindResource(service.Cloudtunnelvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cloudtunnelvserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckCloudtunnelvserverDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cloudtunnelvserver" {
			continue
		}

		_, err := client.FindResource(service.Cloudtunnelvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cloudtunnelvserver %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccCloudtunnelvserverDataSource_basic = `

resource "citrixadc_cloudtunnelvserver" "tf_cloudtunnelvserver" {
  name           = "tf_cloudtunnelvserver"
  servicetype    = "TCP"
  listenpolicy   = "none"
  listenpriority = 50
}

data "citrixadc_cloudtunnelvserver" "tf_cloudtunnelvserver" {
  name       = citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver.name
  depends_on = [citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver]
}
`

func TestAccCloudtunnelvserverDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudtunnelvserverDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "name", "tf_cloudtunnelvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "servicetype", "TCP"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "listenpolicy", "none"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver", "listenpriority", "50"),
				),
			},
		},
	})
}
