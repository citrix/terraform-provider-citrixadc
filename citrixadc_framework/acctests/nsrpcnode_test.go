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
	"net/url"
	"os"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// nsrpcnode is set-only (Create uses set, there is no add/delete): the ADC
// auto-owns an rpcNode for the appliance's own NSIP, and NITRO rejects adding an
// rpcNode for a foreign IP (ec275 "Operation not supported" / ec1088 "No such
// command"). So these tests target the box's OWN NSIP.
//
// Testbed: run on STANDALONE only.
//   - CLUSTER: setting rpcNode 'srcip' is not permitted (ec257), and config must
//     go via the CLIP (node IPs are read-only, ec1203).
//   - HA: excluded on purpose — the rpcNode password is used for HA config sync,
//     so changing it on an HA node would disrupt the pair.

// nsrpcnodeSelfIP returns the appliance's own NSIP, parsed from NS_URL. Falls back
// to a placeholder for non-acceptance/unit runs (where NS_URL is unset and the
// test is skipped by resource.Test anyway).
func nsrpcnodeSelfIP() string {
	if u, err := url.Parse(os.Getenv("NS_URL")); err == nil && u.Hostname() != "" {
		return u.Hostname()
	}
	return "127.0.0.1"
}

// nsrpcnodeSkipUnlessStandalone skips on CLUSTER (srcip not settable / read-only
// nodes) and HA (rpcNode password drives HA sync) and CPX.
func nsrpcnodeSkipUnlessStandalone(t *testing.T) {
	switch adcTestbed {
	case "CLUSTER":
		t.Skipf("ADC testbed is %s: rpcNode 'srcip' is not settable on a cluster (ec257) and config must use the CLIP; nsrpcnode is exercised on a standalone box.", adcTestbed)
	case "HA", "HA_PAIR":
		t.Skipf("ADC testbed is %s: skipping to avoid changing the HA node's rpcNode password (used for HA config sync); nsrpcnode is exercised on a standalone box.", adcTestbed)
	}
	if isCpxRun {
		t.Skip("Operation not permitted under CPX")
	}
}

func nsrpcnodeConfig(ip, secure string) string {
	return fmt.Sprintf(`
resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "%s"
    password  = "CADS123$%%^"
    secure    = "%s"
    srcip     = "%s"
}
`, ip, secure, ip)
}

func nsrpcnodePasswordVarConfig(ip, varName, secure string) string {
	return fmt.Sprintf(`
variable %q {
  type      = string
  sensitive = true
}

resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "%s"
    password  = var.%s
    secure    = "%s"
    srcip     = "%s"
}
`, varName, ip, varName, secure, ip)
}

func nsrpcnodePasswordWoConfig(ip, varName string, version int, secure string) string {
	return fmt.Sprintf(`
variable %q {
  type      = string
  sensitive = true
}

resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress           = "%s"
    password_wo         = var.%s
    password_wo_version = %d
    secure              = "%s"
    srcip               = "%s"
}
`, varName, ip, varName, version, secure, ip)
}

func nsrpcnodeDataSourceConfig(ip string) string {
	return fmt.Sprintf(`
resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "%s"
    password  = "CADS123$%%^"
    secure    = "ON"
    srcip     = "%s"
}

data "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = citrixadc_nsrpcnode.tf_nsrpcnode.ipaddress
}
`, ip, ip)
}

func TestAccNsrpcnode_basic(t *testing.T) {
	nsrpcnodeSkipUnlessStandalone(t)
	ip := nsrpcnodeSelfIP()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: nsrpcnodeConfig(ip, "ON"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "ON"),
				),
			},
			{
				Config: nsrpcnodeConfig(ip, "OFF"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "OFF"),
				),
			},
		},
	})
}

func TestAccNsrpcnode_import(t *testing.T) {
	t.Skip("TODO: Requires review")
}

func testAccCheckNsrpcnodeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
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
		data, err := client.FindResource(service.Nsrpcnode.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("RPC node %s not found", n)
		}

		return nil
	}
}

// Test backward-compatible path: using password (Sensitive attribute)
func TestAccNsrpcnode_password_backward_compat(t *testing.T) {
	nsrpcnodeSkipUnlessStandalone(t)
	ip := nsrpcnodeSelfIP()
	t.Setenv("TF_VAR_nsrpcnode_password", "oldpassword123")
	t.Setenv("TF_VAR_nsrpcnode_password_2", "newpassword456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: nsrpcnodePasswordVarConfig(ip, "nsrpcnode_password", "ON"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "ON"),
				),
			},
			{
				Config: nsrpcnodePasswordVarConfig(ip, "nsrpcnode_password_2", "ON"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "ON"),
				),
			},
		},
	})
}

// Test ephemeral path: using password_wo (WriteOnly attribute) with version tracker
func TestAccNsrpcnode_password_wo_ephemeral(t *testing.T) {
	nsrpcnodeSkipUnlessStandalone(t)
	ip := nsrpcnodeSelfIP()
	t.Setenv("TF_VAR_nsrpcnode_password_wo", "ephemeral_pass1")
	t.Setenv("TF_VAR_nsrpcnode_password_wo_2", "ephemeral_pass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: nsrpcnodePasswordWoConfig(ip, "nsrpcnode_password_wo", 1, "ON"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "password_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "ON"),
				),
			},
			{
				Config: nsrpcnodePasswordWoConfig(ip, "nsrpcnode_password_wo_2", 2, "ON"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "password_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "ON"),
				),
			},
		},
	})
}

func TestAccNsrpcnode_sdkv2StateUpgrade(t *testing.T) {
	nsrpcnodeSkipUnlessStandalone(t)
	ip := nsrpcnodeSelfIP()
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: nsrpcnodeConfig(ip, "ON"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   nsrpcnodeConfig(ip, "ON"),
				// GH #1441 write-only phantom: apply the upgrade and assert no destroy+recreate
				// (expectNoReplace) instead of asserting the strict non-refresh PlanOnly plan.
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
			},
		},
	})
}

func TestAccNsrpcnodeDataSource_basic(t *testing.T) {
	nsrpcnodeSkipUnlessStandalone(t)
	ip := nsrpcnodeSelfIP()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: nsrpcnodeDataSourceConfig(ip),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsrpcnode.tf_nsrpcnode", "ipaddress", ip),
					resource.TestCheckResourceAttr("data.citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "ON"),
				),
			},
		},
	})
}
