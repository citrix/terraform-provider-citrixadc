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

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccLbmonitor_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_lbmonitor.foo", "monitorname", "sample_lb_monitor"),
					resource.TestCheckResourceAttr(
						"citrixadc_lbmonitor.foo", "type", "HTTP"),
				),
			},
		},
	})
}

func testAccCheckLbmonitorExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := client.FindResource(service.Lbmonitor.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbmonitorDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbmonitor" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbmonitor.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbmonitor_basic = `


resource "citrixadc_lbmonitor" "foo" {
  monitorname = "sample_lb_monitor"
  type = "HTTP"
}
`

func TestAccLbmonitor_AssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	monitorName := "tf-acc-lbmonitor-test"
	monitorType := service.Lbmonitor.Type()

	// Defer deletion of actual resource
	deleteArgsMap := make(map[string]string)
	deleteArgsMap["type"] = "HTTP"
	defer testHelperEnsureResourceDeletion(c, t, monitorType, monitorName, deleteArgsMap)

	monitorInstance := lb.Lbmonitor{
		Monitorname: monitorName,
		Type:        "HTTP",
	}

	if _, err := c.client.AddResource(monitorType, monitorName, monitorInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	//servicename
	monitorInstance.Servicename = "foo"
	testHelperVerifyImmutabilityFunc(c, t, monitorType, monitorName, monitorInstance, "servicename")
	monitorInstance.Servicename = ""

	//servicegroupname
	monitorInstance.Servicegroupname = "foo"
	testHelperVerifyImmutabilityFunc(c, t, monitorType, monitorName, monitorInstance, "servicegroupname")
	monitorInstance.Servicegroupname = ""
}

const testAccLbmonitorDataSource_basic = `

	resource "citrixadc_lbmonitor" "tf_lbmonitor" {
		monitorname = "tf_test_lbmonitor_datasource"
		type        = "HTTP"
		interval    = 350
		resptimeout = 2
	}
	
	data "citrixadc_lbmonitor" "tf_lbmonitor" {
		monitorname = citrixadc_lbmonitor.tf_lbmonitor.monitorname
	}
`

func TestAccLbmonitorDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitorDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbmonitor.tf_lbmonitor", "monitorname", "tf_test_lbmonitor_datasource"),
					resource.TestCheckResourceAttr("data.citrixadc_lbmonitor.tf_lbmonitor", "type", "HTTP"),
					resource.TestCheckResourceAttr("data.citrixadc_lbmonitor.tf_lbmonitor", "interval", "350"),
					resource.TestCheckResourceAttr("data.citrixadc_lbmonitor.tf_lbmonitor", "resptimeout", "2"),
				),
			},
		},
	})
}

const testAccLbmonitor_respcode = `

resource "citrixadc_lbmonitor" "tf_lbmonitor_respcode" {
	monitorname = "tf_test_lbmonitor_respcode"
	type        = "HTTP"
	interval    = 5
	resptimeout = 2
	respcode    = ["200", "301"]
}
`

func TestAccLbmonitor_respcode(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_respcode,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_respcode", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode", "monitorname", "tf_test_lbmonitor_respcode"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode", "type", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode", "respcode.#", "2"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode", "respcode.0", "200"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode", "respcode.1", "301"),
				),
			},
		},
	})
}

const testAccLbmonitor_respcode_empty = `

resource "citrixadc_lbmonitor" "tf_lbmonitor_respcode_empty" {
	monitorname = "tf_test_lbmonitor_respcode_empty"
	type        = "HTTP"
	interval    = 5
	resptimeout = 2
	# respcode not specified - NetScaler will use default ["200"]
}
`

func TestAccLbmonitor_respcode_empty(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_respcode_empty,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_respcode_empty", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode_empty", "monitorname", "tf_test_lbmonitor_respcode_empty"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode_empty", "type", "HTTP"),
					// When respcode is not specified, NetScaler returns default ["200"]
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode_empty", "respcode.#", "1"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode_empty", "respcode.0", "200"),
				),
			},
		},
	})
}

const testAccLbmonitor_respcode_single = `

resource "citrixadc_lbmonitor" "tf_lbmonitor_respcode_single" {
	monitorname = "tf_test_lbmonitor_respcode_single"
	type        = "HTTP"
	interval    = 5
	resptimeout = 2
	respcode    = ["200"]
}
`

func TestAccLbmonitor_respcode_single(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_respcode_single,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_respcode_single", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode_single", "monitorname", "tf_test_lbmonitor_respcode_single"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode_single", "type", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode_single", "respcode.#", "1"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_respcode_single", "respcode.0", "200"),
				),
			},
		},
	})
}

// DNS Monitor Tests - respcode does not apply to DNS type monitors

const testAccLbmonitor_dns_basic = `

resource "citrixadc_lbmonitor" "tf_lbmonitor_dns_basic" {
	monitorname = "tf_test_lbmonitor_dns_basic"
	type        = "DNS"
	query       = "example.com"
	querytype   = "Address"
	interval    = 5
	resptimeout = 2
}
`

func TestAccLbmonitor_dns_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_dns_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_dns_basic", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_dns_basic", "monitorname", "tf_test_lbmonitor_dns_basic"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_dns_basic", "type", "DNS"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_dns_basic", "query", "example.com"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_dns_basic", "querytype", "Address"),
					// NetScaler returns nil for respcode on DNS monitors
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_dns_basic", "respcode.#", "0"),
				),
			},
		},
	})
}

const testAccLbmonitor_dns_empty = `

resource "citrixadc_lbmonitor" "tf_lbmonitor_dns_empty" {
	monitorname = "tf_test_lbmonitor_dns_empty"
	type        = "DNS"
	interval    = 5
	resptimeout = 2
	# No query, querytype, or respcode specified
}
`

func TestAccLbmonitor_dns_empty(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_dns_empty,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_dns_empty", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_dns_empty", "monitorname", "tf_test_lbmonitor_dns_empty"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_dns_empty", "type", "DNS"),
					// NetScaler returns nil for respcode on DNS monitors
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_dns_empty", "respcode.#", "0"),
				),
			},
		},
	})
}


