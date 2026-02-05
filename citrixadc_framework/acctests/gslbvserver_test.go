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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGslbvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbvserver_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserverExist("citrixadc_gslbvserver.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_gslbvserver.foo", "dnsrecordtype", "A"),
					resource.TestCheckResourceAttr(
						"citrixadc_gslbvserver.foo", "name", "GSLB-East-Coast-Vserver"),
					resource.TestCheckResourceAttr(
						"citrixadc_gslbvserver.foo", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr(
						"citrixadc_gslbvserver.foo", "toggleorder", "ASCENDING"),
					resource.TestCheckResourceAttr(
						"citrixadc_gslbvserver.foo", "orderthreshold", "50"),
				),
			},
		},
	})
}

func testAccCheckGslbvserverExist(n string, id *string) resource.TestCheckFunc {
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Gslbvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("GSLB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbvserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Gslbvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("GSLB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccGslbvserver_basic = `


resource "citrixadc_gslbvserver" "foo" {

  dnsrecordtype = "A"
  name = "GSLB-East-Coast-Vserver"
  servicetype = "HTTP"
  toggleorder = "ASCENDING"
  orderthreshold = "50"
  domain {
	  domainname =  "www.fooco.co"
	  ttl = "60"
	}
  domain {
	  domainname = "www.barco.com"
	  ttl = "55"
	}
}
`

func TestAccGslbvserver_AssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	serverName := "tf-acc-glsb-vserver-test"
	serverType := service.Gslbvserver.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, serverType, serverName, nil)

	serverInstance := gslb.Gslbvserver{
		Name:        serverName,
		Servicetype: "HTTP",
	}

	if _, err := c.client.AddResource(serverType, serverName, serverInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// servicetype
	serverInstance.Servicetype = "SSL"
	testHelperVerifyImmutabilityFunc(c, t, serverType, serverName, serverInstance, "servicetype")
}

const testAccGslbvserverEnableDisable_enabled = `
resource "citrixadc_gslbvserver" "tf_test_acc_gslbvsever" {
	dnsrecordtype = "A"
	name = "tf_test_acc_gslbvsever"
	servicetype = "HTTP"
	state = "ENABLED"
	comment = "enabled state comment"
}
`

const testAccGslbvserverEnableDisable_disabled = `
resource "citrixadc_gslbvserver" "tf_test_acc_gslbvsever" {
	dnsrecordtype = "A"
	name = "tf_test_acc_gslbvsever"
	servicetype = "HTTP"
	state = "DISABLED"
	comment = "disabled state comment"
}
`

func TestAccGslbvserver_enable_disable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbvserverDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccGslbvserverEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserverExist("citrixadc_gslbvserver.tf_test_acc_gslbvsever", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbvserver.tf_test_acc_gslbvsever", "state", "ENABLED"),
				),
			},
			// Disable
			{
				Config: testAccGslbvserverEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserverExist("citrixadc_gslbvserver.tf_test_acc_gslbvsever", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbvserver.tf_test_acc_gslbvsever", "state", "DISABLED"),
				),
			},
			// Re enable
			{
				Config: testAccGslbvserverEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserverExist("citrixadc_gslbvserver.tf_test_acc_gslbvsever", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbvserver.tf_test_acc_gslbvsever", "state", "ENABLED"),
				),
			},
		},
	})
}

const testAccGslbvserverDataSource_basic = `

resource "citrixadc_gslbvserver" "tf_gslbvserver" {
  dnsrecordtype = "A"
  name = "tf_test_gslbvserver_ds"
  servicetype = "HTTP"
  toggleorder = "ASCENDING"
  orderthreshold = "50"
}

data "citrixadc_gslbvserver" "tf_gslbvserver" {
  name = citrixadc_gslbvserver.tf_gslbvserver.name
}
`

func TestAccGslbvserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbvserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_gslbvserver.tf_gslbvserver", "name", "tf_test_gslbvserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbvserver.tf_gslbvserver", "dnsrecordtype", "A"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbvserver.tf_gslbvserver", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbvserver.tf_gslbvserver", "toggleorder", "ASCENDING"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbvserver.tf_gslbvserver", "orderthreshold", "50"),
				),
			},
		},
	})
}
