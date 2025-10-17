/*
Copyright 2016 Citrix Systems, Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.o	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		fv := reflect.ValueOf(client).Elem().FieldByName("headers")

		if fmt.Sprintf("%v", fv) == "map[User-Agent:terraform-ctxadc]" {
			return nil
		} else {/LICENSE-2.0

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
	"reflect"
	"testing"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServer_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExist("citrixadc_server.foo", nil),
					testAccCheckUserAgent(),
				),
			},
		},
	})
}

func testAccCheckServerExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No server name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Server.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("server %s not found", n)
		}

		return nil
	}
}

func testAccCheckServerDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_server" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Server.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("server %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccServer_basic = `


resource "citrixadc_server" "foo" {
	name = "test_server"
	ipaddress = "192.168.11.13"
}
`

// Test for immutability of attributes
// This is to catch any attibute having ForceNew: true while not actually needed
func TestAccServer_AssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	serverName := "tf-acc-server-name"
	serverType := service.Server.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, serverType, serverName, nil)

	serverInstance := basic.Server{
		Domain:      "tfacc.domain.com",
		Ipv6address: "YES",
		Name:        serverName,
		Td:          intPtr(0),
	}

	if _, err := c.client.AddResource(serverType, serverName, serverInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// Verify immutability of argument td
	serverInstance.Domain = ""
	serverInstance.Ipv6address = ""
	serverInstance.Td = intPtr(10)
	testHelperVerifyImmutabilityFunc(c, t, serverType, serverName, serverInstance, "td")
	serverInstance.Td = intPtr(0)

	// Verify immutability of argument domain
	serverInstance.Domain = "newdomain.com"
	serverInstance.Ipv6address = ""
	testHelperVerifyImmutabilityFunc(c, t, serverType, serverName, serverInstance, "domain")
	serverInstance.Domain = ""

	// Verify immutability of argument ipv6address
	serverInstance.Ipv6address = "YES"
	serverInstance.Td = intPtr(0)
	testHelperVerifyImmutabilityFunc(c, t, serverType, serverName, serverInstance, "ipv6address")
	serverInstance.Ipv6address = ""
}

const testAccServerEnableDisable_enabled = `
resource "citrixadc_server" "tf_enable_disable_test_svr" {
	name = "tf_enable_disable_test_svr"
	ipaddress = "192.168.43.33"
	comment = "enabled state comment"
	state = "ENABLED"
	graceful = "YES"
	delay = 60
}
`

const testAccServerEnableDisable_disabled = `
resource "citrixadc_server" "tf_enable_disable_test_svr" {
	name = "tf_enable_disable_test_svr"
	ipaddress = "192.168.43.33"
	comment = "disabled state comment"
	state = "DISABLED"
	graceful = "YES"
	delay = 60
}
`

func TestAccServer_enable_disable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckServerDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccServerEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExist("citrixadc_server.tf_enable_disable_test_svr", nil),
					resource.TestCheckResourceAttr("citrixadc_server.tf_enable_disable_test_svr", "state", "ENABLED"),
				),
			},
			// Disable
			{
				Config: testAccServerEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExist("citrixadc_server.tf_enable_disable_test_svr", nil),
					resource.TestCheckResourceAttr("citrixadc_server.tf_enable_disable_test_svr", "state", "DISABLED"),
				),
			},
			// Re enable
			{
				Config: testAccServerEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExist("citrixadc_server.tf_enable_disable_test_svr", nil),
					resource.TestCheckResourceAttr("citrixadc_server.tf_enable_disable_test_svr", "state", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckUserAgent() resource.TestCheckFunc {
	// TODO check logs of ADC for presence of user agent string
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		fv := reflect.ValueOf(client).Elem().FieldByName("headers")

		if fmt.Sprintf("%v", fv) == "map[User-Agent:terraform-ctxadc]" {
			return nil
		} else {
			return fmt.Errorf("Could not verify headers. fv is %v", fv)
		}
	}
}
