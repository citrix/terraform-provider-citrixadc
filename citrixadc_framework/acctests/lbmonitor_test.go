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
	"strings"
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

		// Parse compound ID (e.g. "monitorname:foo,type:HTTP") to extract the monitorname
		monitorName := rs.Primary.ID
		for _, part := range strings.Split(rs.Primary.ID, ",") {
			if kv := strings.SplitN(part, ":", 2); len(kv) == 2 && kv[0] == "monitorname" {
				monitorName = kv[1]
				break
			}
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lbmonitor.Type(), monitorName)

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

		// Parse compound ID (e.g. "monitorname:foo,type:HTTP") to extract the monitorname
		monitorName := rs.Primary.ID
		for _, part := range strings.Split(rs.Primary.ID, ",") {
			if kv := strings.SplitN(part, ":", 2); len(kv) == 2 && kv[0] == "monitorname" {
				monitorName = kv[1]
				break
			}
		}

		_, err := client.FindResource(service.Lbmonitor.Type(), monitorName)
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
		type        = citrixadc_lbmonitor.tf_lbmonitor.type
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

// ============================================================
// Ephemeral / Write-Only tests for secret attribute: password
// ============================================================

const testAccLbmonitor_password_step1 = `

variable "lbmonitor_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_password" {
  monitorname = "tf_test_lbmonitor_password"
  type        = "FTP"
  username    = "testuser"
  password    = var.lbmonitor_password
}
`

const testAccLbmonitor_password_step2 = `

variable "lbmonitor_password_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_password" {
  monitorname = "tf_test_lbmonitor_password"
  type        = "FTP"
  username    = "testuser"
  password    = var.lbmonitor_password_2
}
`

func TestAccLbmonitor_password_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_lbmonitor_password", "Password1!")
	t.Setenv("TF_VAR_lbmonitor_password_2", "Password2!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_password", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_password", "monitorname", "tf_test_lbmonitor_password"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_password", "type", "FTP"),
				),
			},
			{
				Config: testAccLbmonitor_password_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_password", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_password", "monitorname", "tf_test_lbmonitor_password"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_password", "type", "FTP"),
				),
			},
		},
	})
}

const testAccLbmonitor_password_wo_step1 = `

variable "lbmonitor_password_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_password_wo" {
  monitorname        = "tf_test_lbmonitor_password_wo"
  type               = "FTP"
  username           = "testuser"
  password_wo        = var.lbmonitor_password_wo
  password_wo_version = 1
}
`

const testAccLbmonitor_password_wo_step2 = `

variable "lbmonitor_password_wo_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_password_wo" {
  monitorname        = "tf_test_lbmonitor_password_wo"
  type               = "FTP"
  username           = "testuser"
  password_wo        = var.lbmonitor_password_wo_2
  password_wo_version = 2
}
`

func TestAccLbmonitor_password_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_lbmonitor_password_wo", "Password1!")
	t.Setenv("TF_VAR_lbmonitor_password_wo_2", "Password2!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_password_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_password_wo", "monitorname", "tf_test_lbmonitor_password_wo"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_password_wo", "type", "FTP"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_password_wo", "password_wo_version", "1"),
				),
			},
			{
				Config: testAccLbmonitor_password_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_password_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_password_wo", "monitorname", "tf_test_lbmonitor_password_wo"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_password_wo", "type", "FTP"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_password_wo", "password_wo_version", "2"),
				),
			},
		},
	})
}

// ============================================================
// Ephemeral / Write-Only tests for secret attribute: radkey
// ============================================================

const testAccLbmonitor_radkey_step1 = `

variable "lbmonitor_radkey" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_radkey" {
  monitorname = "tf_test_lbmonitor_radkey"
  type        = "RADIUS"
  username    = "raduser"
  password    = "RadPass1!"
  radkey      = var.lbmonitor_radkey
}
`

const testAccLbmonitor_radkey_step2 = `

variable "lbmonitor_radkey_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_radkey" {
  monitorname = "tf_test_lbmonitor_radkey"
  type        = "RADIUS"
  username    = "raduser"
  password    = "RadPass1!"
  radkey      = var.lbmonitor_radkey_2
}
`

func TestAccLbmonitor_radkey_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_lbmonitor_radkey", "secret1key")
	t.Setenv("TF_VAR_lbmonitor_radkey_2", "secret2key")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_radkey_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_radkey", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_radkey", "monitorname", "tf_test_lbmonitor_radkey"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_radkey", "type", "RADIUS"),
				),
			},
			{
				Config: testAccLbmonitor_radkey_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_radkey", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_radkey", "monitorname", "tf_test_lbmonitor_radkey"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_radkey", "type", "RADIUS"),
				),
			},
		},
	})
}

const testAccLbmonitor_radkey_wo_step1 = `

variable "lbmonitor_radkey_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_radkey_wo" {
  monitorname      = "tf_test_lbmonitor_radkey_wo"
  type             = "RADIUS"
  username         = "raduser"
  password         = "RadPass1!"
  radkey_wo        = var.lbmonitor_radkey_wo
  radkey_wo_version = 1
}
`

const testAccLbmonitor_radkey_wo_step2 = `

variable "lbmonitor_radkey_wo_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_radkey_wo" {
  monitorname      = "tf_test_lbmonitor_radkey_wo"
  type             = "RADIUS"
  username         = "raduser"
  password         = "RadPass1!"
  radkey_wo        = var.lbmonitor_radkey_wo_2
  radkey_wo_version = 2
}
`

func TestAccLbmonitor_radkey_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_lbmonitor_radkey_wo", "secret1key")
	t.Setenv("TF_VAR_lbmonitor_radkey_wo_2", "secret2key")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_radkey_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_radkey_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_radkey_wo", "monitorname", "tf_test_lbmonitor_radkey_wo"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_radkey_wo", "type", "RADIUS"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_radkey_wo", "radkey_wo_version", "1"),
				),
			},
			{
				Config: testAccLbmonitor_radkey_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_radkey_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_radkey_wo", "monitorname", "tf_test_lbmonitor_radkey_wo"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_radkey_wo", "type", "RADIUS"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_radkey_wo", "radkey_wo_version", "2"),
				),
			},
		},
	})
}

// ==================================================================
// Ephemeral / Write-Only tests for secret attribute: secondarypassword
// ==================================================================

const testAccLbmonitor_secondarypassword_step1 = `

variable "lbmonitor_secondarypassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_secondarypassword" {
  monitorname       = "tf_test_lbmonitor_secpwd"
  type              = "CITRIX-AG"
  secondarypassword = var.lbmonitor_secondarypassword
}
`

const testAccLbmonitor_secondarypassword_step2 = `

variable "lbmonitor_secondarypassword_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_secondarypassword" {
  monitorname       = "tf_test_lbmonitor_secpwd"
  type              = "CITRIX-AG"
  secondarypassword = var.lbmonitor_secondarypassword_2
}
`

func TestAccLbmonitor_secondarypassword_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_lbmonitor_secondarypassword", "SecPwd1!")
	t.Setenv("TF_VAR_lbmonitor_secondarypassword_2", "SecPwd2!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_secondarypassword_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword", "monitorname", "tf_test_lbmonitor_secpwd"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword", "type", "CITRIX-AG"),
				),
			},
			{
				Config: testAccLbmonitor_secondarypassword_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword", "monitorname", "tf_test_lbmonitor_secpwd"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword", "type", "CITRIX-AG"),
				),
			},
		},
	})
}

const testAccLbmonitor_secondarypassword_wo_step1 = `

variable "lbmonitor_secondarypassword_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_secondarypassword_wo" {
  monitorname                  = "tf_test_lbmonitor_secpwd_wo"
  type                         = "CITRIX-AG"
  secondarypassword_wo         = var.lbmonitor_secondarypassword_wo
  secondarypassword_wo_version = 1
}
`

const testAccLbmonitor_secondarypassword_wo_step2 = `

variable "lbmonitor_secondarypassword_wo_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_secondarypassword_wo" {
  monitorname                  = "tf_test_lbmonitor_secpwd_wo"
  type                         = "CITRIX-AG"
  secondarypassword_wo         = var.lbmonitor_secondarypassword_wo_2
  secondarypassword_wo_version = 2
}
`

func TestAccLbmonitor_secondarypassword_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_lbmonitor_secondarypassword_wo", "SecPwd1!")
	t.Setenv("TF_VAR_lbmonitor_secondarypassword_wo_2", "SecPwd2!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_secondarypassword_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword_wo", "monitorname", "tf_test_lbmonitor_secpwd_wo"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword_wo", "type", "CITRIX-AG"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword_wo", "secondarypassword_wo_version", "1"),
				),
			},
			{
				Config: testAccLbmonitor_secondarypassword_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword_wo", "monitorname", "tf_test_lbmonitor_secpwd_wo"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword_wo", "type", "CITRIX-AG"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secondarypassword_wo", "secondarypassword_wo_version", "2"),
				),
			},
		},
	})
}

// ============================================================
// Ephemeral / Write-Only tests for secret attribute: secureargs
// ============================================================

const testAccLbmonitor_secureargs_step1 = `

variable "lbmonitor_secureargs" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_secureargs" {
  monitorname = "tf_test_lbmonitor_secureargs"
  type        = "USER"
  scriptname  = "nssimpleauth.pl"
  secureargs  = var.lbmonitor_secureargs
}
`

const testAccLbmonitor_secureargs_step2 = `

variable "lbmonitor_secureargs_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_secureargs" {
  monitorname = "tf_test_lbmonitor_secureargs"
  type        = "USER"
  scriptname  = "nssimpleauth.pl"
  secureargs  = var.lbmonitor_secureargs_2
}
`

func TestAccLbmonitor_secureargs_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_lbmonitor_secureargs", "secret1=val1")
	t.Setenv("TF_VAR_lbmonitor_secureargs_2", "secret2=val2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_secureargs_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_secureargs", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secureargs", "monitorname", "tf_test_lbmonitor_secureargs"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secureargs", "type", "USER"),
				),
			},
			{
				Config: testAccLbmonitor_secureargs_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_secureargs", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secureargs", "monitorname", "tf_test_lbmonitor_secureargs"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secureargs", "type", "USER"),
				),
			},
		},
	})
}

const testAccLbmonitor_secureargs_wo_step1 = `

variable "lbmonitor_secureargs_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_secureargs_wo" {
  monitorname           = "tf_test_lbmonitor_secureargs_wo"
  type                  = "USER"
  scriptname            = "nssimpleauth.pl"
  secureargs_wo         = var.lbmonitor_secureargs_wo
  secureargs_wo_version = 1
}
`

const testAccLbmonitor_secureargs_wo_step2 = `

variable "lbmonitor_secureargs_wo_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_lbmonitor" "tf_lbmonitor_secureargs_wo" {
  monitorname           = "tf_test_lbmonitor_secureargs_wo"
  type                  = "USER"
  scriptname            = "nssimpleauth.pl"
  secureargs_wo         = var.lbmonitor_secureargs_wo_2
  secureargs_wo_version = 2
}
`

func TestAccLbmonitor_secureargs_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_lbmonitor_secureargs_wo", "secret1=val1")
	t.Setenv("TF_VAR_lbmonitor_secureargs_wo_2", "secret2=val2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitor_secureargs_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_secureargs_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secureargs_wo", "monitorname", "tf_test_lbmonitor_secureargs_wo"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secureargs_wo", "type", "USER"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secureargs_wo", "secureargs_wo_version", "1"),
				),
			},
			{
				Config: testAccLbmonitor_secureargs_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorExist("citrixadc_lbmonitor.tf_lbmonitor_secureargs_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secureargs_wo", "monitorname", "tf_test_lbmonitor_secureargs_wo"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secureargs_wo", "type", "USER"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor.tf_lbmonitor_secureargs_wo", "secureargs_wo_version", "2"),
				),
			},
		},
	})
}
