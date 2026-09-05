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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccNtpserver_basic_ip = `

resource "citrixadc_ntpserver" "tf_ntpserver" {
	serverip          = "10.222.74.200"
	minpoll            = 5
	maxpoll            = 9
	preferredntpserver = "NO"
  
	}
  
`
const testAccNtpserver_update_ip = `

resource "citrixadc_ntpserver" "tf_ntpserver" {
	serverip         = "10.222.74.200"
	minpoll            = 6
	maxpoll            = 10
	preferredntpserver = "YES"
  
	}
`

func TestAccNtpserver_ip(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNtpserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNtpserver_basic_ip,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpserverExist("citrixadc_ntpserver.tf_ntpserver", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "serverip", "10.222.74.200"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "minpoll", "5"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "maxpoll", "9"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "preferredntpserver", "NO"),
				),
			},
			{
				Config: testAccNtpserver_update_ip,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpserverExist("citrixadc_ntpserver.tf_ntpserver", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "serverip", "10.222.74.200"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "minpoll", "6"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "maxpoll", "10"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "preferredntpserver", "YES"),
				),
			},
		},
	})
}

const testAccNtpserver_basic_servername = `

resource "citrixadc_ntpserver" "tf_ntpserver" {
	servername         = "www.example.com"
	minpoll            = 5
	maxpoll            = 9
	preferredntpserver = "NO"
  
	}
  
`
const testAccNtpserver_update_servername = `

resource "citrixadc_ntpserver" "tf_ntpserver" {
	servername        = "www.example.com"
	minpoll            = 6
	maxpoll            = 10
	preferredntpserver = "YES"
  
	}
  
`

func TestAccNtpserver_servername(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNtpserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNtpserver_basic_servername,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpserverExist("citrixadc_ntpserver.tf_ntpserver", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "servername", "www.example.com"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "minpoll", "5"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "maxpoll", "9"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "preferredntpserver", "NO"),
				),
			},
			{
				Config: testAccNtpserver_update_servername,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpserverExist("citrixadc_ntpserver.tf_ntpserver", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "servername", "www.example.com"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "minpoll", "6"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "maxpoll", "10"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "preferredntpserver", "YES"),
				),
			},
		},
	})
}
func testAccCheckNtpserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ntpserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		ntpserverName := rs.Primary.ID
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Ntpserver.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["serverip"] == ntpserverName || v["servername"] == ntpserverName {
				found = true
				break
			}

		}
		if !found {
			return fmt.Errorf("ntpserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckNtpserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ntpserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		ntpserverName := rs.Primary.ID
		dataArr, err := client.FindAllResources(service.Ntpserver.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["serverip"] == ntpserverName || v["servername"] == ntpserverName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("ntpserver %s still exists", ntpserverName)
		}
	}
	return nil
}

func TestAccNtpserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNtpserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNtpserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_ntpserver.tf_ntpserver_ds", "serverip", "10.222.74.150"),
					resource.TestCheckResourceAttr("data.citrixadc_ntpserver.tf_ntpserver_ds", "minpoll", "6"),
					resource.TestCheckResourceAttr("data.citrixadc_ntpserver.tf_ntpserver_ds", "maxpoll", "10"),
				),
			},
		},
	})
}

func TestAccNtpserver_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_ntpserver.tf_ntpserver"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNtpserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNtpserver_basic_ip,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNtpserverExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Ntpserver.Type(), "10.222.74.200"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNtpserver_basic_ip,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNtpserverExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNtpserver_import(t *testing.T) {
	const resAddr = "citrixadc_ntpserver.tf_ntpserver"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNtpserverDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNtpserver_basic_ip},
			{
				Config:                  testAccNtpserver_basic_ip,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"serverip", "servername"},
			},
		},
	})
}

func TestAccNtpserver_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNtpserverDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccNtpserver_basic_ip,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNtpserverExist("citrixadc_ntpserver.tf_ntpserver", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNtpserver_basic_ip,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNtpserverExist("citrixadc_ntpserver.tf_ntpserver", nil)),
			},
		},
	})
}

// testAccNtpserver_unset_step1 sets the unset-eligible attributes to valid
// NON-default values; step2 removes them so the provider must unset them,
// reverting to the NITRO defaults (minpoll=6, maxpoll=10,
// preferredntpserver=NO, autokey=false).
const testAccNtpserver_unset_step1 = `
resource "citrixadc_ntpserver" "tf_unset" {
	serverip           = "10.222.74.201"
	minpoll            = 5
	maxpoll            = 11
	preferredntpserver = "YES"
	autokey            = true
}
`

const testAccNtpserver_unset_step2 = `
resource "citrixadc_ntpserver" "tf_unset" {
	serverip = "10.222.74.201"
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccNtpserver_unset(t *testing.T) {
	const resAddr = "citrixadc_ntpserver.tf_unset"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNtpserverDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNtpserver_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpserverExist(resAddr, nil),
					resource.TestCheckResourceAttr(resAddr, "minpoll", "5"),
					resource.TestCheckResourceAttr(resAddr, "maxpoll", "11"),
					resource.TestCheckResourceAttr(resAddr, "preferredntpserver", "YES"),
					resource.TestCheckResourceAttr(resAddr, "autokey", "true"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the NITRO defaults, and the implicit
				// post-apply plan must be empty.
				Config: testAccNtpserver_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpserverExist(resAddr, nil),
					resource.TestCheckResourceAttr(resAddr, "minpoll", "6"),
					resource.TestCheckResourceAttr(resAddr, "maxpoll", "10"),
					resource.TestCheckResourceAttr(resAddr, "preferredntpserver", "NO"),
					resource.TestCheckResourceAttr(resAddr, "autokey", "false"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNtpserverADCValue("10.222.74.201", "minpoll", "6"),
					testAccCheckNtpserverADCValue("10.222.74.201", "maxpoll", "10"),
					testAccCheckNtpserverADCValue("10.222.74.201", "preferredntpserver", "NO"),
					testAccCheckNtpserverADCValue("10.222.74.201", "autokey", "false"),
				),
			},
		},
	})
}

// testAccCheckNtpserverADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. The identifier matches either serverip or servername.
func testAccCheckNtpserverADCValue(identifier, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Ntpserver.Type())
		if err != nil {
			return err
		}
		for _, v := range dataArr {
			if v["serverip"] == identifier || v["servername"] == identifier {
				got := strings.TrimSpace(fmt.Sprintf("%v", v[attr]))
				if got != want {
					return fmt.Errorf("ntpserver %s: appliance attr %q = %q, want %q (unset did not revert it)", identifier, attr, got, want)
				}
				return nil
			}
		}
		return fmt.Errorf("ntpserver %s not found on appliance", identifier)
	}
}

const testAccNtpserverDataSource_basic = `

resource "citrixadc_ntpserver" "tf_ntpserver_ds" {
	serverip          = "10.222.74.150"
	minpoll            = 6
	maxpoll            = 10
	preferredntpserver = "NO"
}

data "citrixadc_ntpserver" "tf_ntpserver_ds" {
	serverip = citrixadc_ntpserver.tf_ntpserver_ds.serverip
	depends_on = [citrixadc_ntpserver.tf_ntpserver_ds]
}
`
