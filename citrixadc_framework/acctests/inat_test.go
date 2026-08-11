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
	"regexp"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

func TestAccInat_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInat_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInatExist("citrixadc_inat.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_inat.foo", "name", "ip4ip"),
					resource.TestCheckResourceAttr(
						"citrixadc_inat.foo", "privateip", "192.168.1.1"),
					resource.TestCheckResourceAttr(
						"citrixadc_inat.foo", "publicip", "172.16.1.2"),
					resource.TestCheckResourceAttr(
						"citrixadc_inat.foo", "tcpproxy", "ENABLED"),
					resource.TestCheckResourceAttr(
						"citrixadc_inat.foo", "usnip", "ON"),
					resource.TestCheckResourceAttr(
						"citrixadc_inat.foo", "connfailover", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckInatExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := client.FindResource(service.Inat.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Inat rule %s not found", n)
		}

		return nil
	}
}

func testAccCheckInatDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_inat" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Inat.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Inat rule %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccInat_basic = `


resource "citrixadc_inat" "foo" {

  name = "ip4ip"
  privateip = "192.168.1.1"
  publicip = "172.16.1.2"
  tcpproxy = "ENABLED"
  usnip = "ON"
  connfailover = "DISABLED"

}
`

func TestAccInat_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_inat.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInat_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckInatExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Inat.Type(), "ip4ip"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccInat_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckInatExist(resAddr, nil)),
			},
		},
	})
}

func TestAccInat_AssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	inatName := "tf-acc-inat-test"
	inatType := service.Inat.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, inatType, inatName, nil)

	inatInstance := network.Inat{
		Name:      inatName,
		Privateip: "192.168.1.1",
		Publicip:  "172.16.1.2",
	}

	if _, err := c.client.AddResource(inatType, inatName, inatInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// publicip
	inatInstance.Publicip = "172.16.1.3"
	testHelperVerifyImmutabilityFunc(c, t, inatType, inatName, inatInstance, "publicip")
	inatInstance.Publicip = ""

	// name
	newName := "inat-new-name"
	inatInstance.Name = newName
	if _, err := c.client.UpdateResource(inatType, inatName, inatInstance); err != nil {
		r := regexp.MustCompile(fmt.Sprintf("errorcode.*258.*No such resource \\[name, %s\\]", newName))
		if r.Match([]byte(err.Error())) {
			t.Logf("Succesfully verified immutability of attribute name")
		} else {
			t.Errorf("Error while assesing immutability of attribute name")
			t.Fatal(err)
		}
	}

	// td
	inatInstance.Td = utils.IntPtr(1)
	testHelperVerifyImmutabilityFunc(c, t, inatType, inatName, inatInstance, "td")
	inatInstance.Td = utils.IntPtr(0)
}

func TestAccInat_import(t *testing.T) {
	const resAddr = "citrixadc_inat.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInatDestroy,
		Steps: []resource.TestStep{
			{Config: testAccInat_basic},
			{
				Config:                  testAccInat_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccInatDataSource_basic = `

resource "citrixadc_inat" "foo" {
  name = "ip4ip"
  privateip = "192.168.1.1"
  publicip = "172.16.1.2"
  tcpproxy = "ENABLED"
  usnip = "ON"
  connfailover = "DISABLED"
}

data "citrixadc_inat" "foo" {
  name = citrixadc_inat.foo.name
}
`

func TestAccInat_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckInatDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccInat_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckInatExist("citrixadc_inat.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccInat_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckInatExist("citrixadc_inat.foo", nil)),
			},
		},
	})
}

// The inat unset test covers the mutable, unset-eligible attributes that have a
// documented NITRO default: tcpproxy, ftp, tftp (DISABLED), useproxyport
// (ENABLED) and connfailover (DISABLED). step1 sets each to a non-default value;
// step2 removes them from config so the provider unsets them and the appliance
// reverts them to their defaults.
const testAccInat_unset_step1 = `
resource "citrixadc_inat" "tf_unset" {
  name         = "tf_test_inat_unset"
  publicip     = "172.16.9.9"
  privateip    = "192.168.9.9"
  tcpproxy     = "ENABLED"
  ftp          = "ENABLED"
  tftp         = "ENABLED"
  useproxyport = "DISABLED"
  connfailover = "ENABLED"
}
`

const testAccInat_unset_step2 = `
resource "citrixadc_inat" "tf_unset" {
  name      = "tf_test_inat_unset"
  publicip  = "172.16.9.9"
  privateip = "192.168.9.9"
  # All unset-eligible attributes removed from config -> the provider must unset
  # them (revert to NITRO defaults).
}
`

func TestAccInat_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInatDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccInat_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInatExist("citrixadc_inat.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_inat.tf_unset", "tcpproxy", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_inat.tf_unset", "ftp", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_inat.tf_unset", "tftp", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_inat.tf_unset", "useproxyport", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_inat.tf_unset", "connfailover", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccInat_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInatExist("citrixadc_inat.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_inat.tf_unset", "tcpproxy", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_inat.tf_unset", "ftp", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_inat.tf_unset", "tftp", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_inat.tf_unset", "useproxyport", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_inat.tf_unset", "connfailover", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckInatADCValue("tf_test_inat_unset", "tcpproxy", "DISABLED"),
					testAccCheckInatADCValue("tf_test_inat_unset", "useproxyport", "ENABLED"),
					testAccCheckInatADCValue("tf_test_inat_unset", "connfailover", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckInatADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckInatADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Inat.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("inat %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("inat %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccInatDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInatDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_inat.foo", "name", "ip4ip"),
					resource.TestCheckResourceAttr("data.citrixadc_inat.foo", "privateip", "192.168.1.1"),
					resource.TestCheckResourceAttr("data.citrixadc_inat.foo", "publicip", "172.16.1.2"),
					resource.TestCheckResourceAttr("data.citrixadc_inat.foo", "tcpproxy", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_inat.foo", "usnip", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_inat.foo", "connfailover", "DISABLED"),
				),
			},
		},
	})
}
