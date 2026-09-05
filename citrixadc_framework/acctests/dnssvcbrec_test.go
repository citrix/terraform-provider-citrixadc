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

// dnssvcbrec is a composite-key DNS record: a record is identified by
// domain + targetname + priority + svcbtype. domain, targetname, priority and
// svcbtype are RequiresReplace; alpn, encryptedclienthello, ipv4hint, ipv6hint,
// mandatory, nodefaultalpn, port and ttl are updatable in place and support the
// NITRO unset operation.

import (
	"fmt"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccDnssvcbrec_add = `
resource "citrixadc_dnssvcbrec" "dnssvcbrec" {
	domain     = "svcb.example.com"
	targetname = "target.example.com"
	priority   = 1
	svcbtype   = "HTTPS"
	alpn       = "h2,h3"
	port       = 443
	ttl        = 3600
}
`

const testAccDnssvcbrec_update = `
resource "citrixadc_dnssvcbrec" "dnssvcbrec" {
	domain     = "svcb.example.com"
	targetname = "target.example.com"
	priority   = 1
	svcbtype   = "HTTPS"
	alpn       = "h2"
	port       = 8443
	ttl        = 7200
}
`

func TestAccDnssvcbrec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssvcbrecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssvcbrec_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssvcbrecExist("citrixadc_dnssvcbrec.dnssvcbrec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.dnssvcbrec", "domain", "svcb.example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.dnssvcbrec", "targetname", "target.example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.dnssvcbrec", "priority", "1"),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.dnssvcbrec", "svcbtype", "HTTPS"),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.dnssvcbrec", "alpn", "h2,h3"),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.dnssvcbrec", "port", "443"),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.dnssvcbrec", "ttl", "3600"),
				),
			},
			{
				Config: testAccDnssvcbrec_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssvcbrecExist("citrixadc_dnssvcbrec.dnssvcbrec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.dnssvcbrec", "alpn", "h2"),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.dnssvcbrec", "port", "8443"),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.dnssvcbrec", "ttl", "7200"),
				),
			},
		},
	})
}

func TestAccDnssvcbrec_import(t *testing.T) {
	const resAddr = "citrixadc_dnssvcbrec.dnssvcbrec"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssvcbrecDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDnssvcbrec_add},
			{
				Config:                  testAccDnssvcbrec_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

// alpn, port and ttl are spec-documented unsettable attributes. step1 sets
// non-default values; step2 removes them so the provider must unset them.
const testAccDnssvcbrec_unset_step1 = `
resource "citrixadc_dnssvcbrec" "tf_unset" {
	domain     = "svcb.example.com"
	targetname = "target.example.com"
	priority   = 1
	svcbtype   = "HTTPS"
	alpn       = "h2,h3"
	port       = 8443
	ttl        = 7200
}
`

const testAccDnssvcbrec_unset_step2 = `
resource "citrixadc_dnssvcbrec" "tf_unset" {
	domain     = "svcb.example.com"
	targetname = "target.example.com"
	priority   = 1
	svcbtype   = "HTTPS"
	# alpn, port and ttl removed -> provider must unset them (revert to NITRO defaults).
}
`

func TestAccDnssvcbrec_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssvcbrecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssvcbrec_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssvcbrecExist("citrixadc_dnssvcbrec.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.tf_unset", "alpn", "h2,h3"),
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.tf_unset", "ttl", "7200"),
				),
			},
			{
				Config: testAccDnssvcbrec_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssvcbrecExist("citrixadc_dnssvcbrec.tf_unset", nil),
					// ttl reverts to the documented NITRO default after unset.
					resource.TestCheckResourceAttr("citrixadc_dnssvcbrec.tf_unset", "ttl", "3600"),
					testAccCheckDnssvcbrecADCValue("svcb.example.com", "target.example.com", "HTTPS", "ttl", "3600"),
				),
			},
		},
	})
}

func TestAccDnssvcbrec_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dnssvcbrec.dnssvcbrec"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssvcbrecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssvcbrec_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnssvcbrecExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					// Delete the record out-of-band exactly as the resource Delete does.
					argsMap := map[string]string{
						"targetname": "target.example.com",
						"priority":   "1",
						"svcbtype":   "HTTPS",
					}
					if err := client.DeleteResourceWithArgsMap(service.Dnssvcbrec.Type(), "svcb.example.com", argsMap); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDnssvcbrec_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnssvcbrecExist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckDnssvcbrecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnssvcbrec id is set")
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

		found, err := dnssvcbrecRecordExists(client,
			rs.Primary.Attributes["domain"],
			rs.Primary.Attributes["targetname"],
			rs.Primary.Attributes["priority"],
			rs.Primary.Attributes["svcbtype"])
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("dnssvcbrec %s does not exist on the appliance", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckDnssvcbrecDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnssvcbrec" {
			continue
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}
		found, err := dnssvcbrecRecordExists(client,
			rs.Primary.Attributes["domain"],
			rs.Primary.Attributes["targetname"],
			rs.Primary.Attributes["priority"],
			rs.Primary.Attributes["svcbtype"])
		if err != nil {
			return err
		}
		if found {
			return fmt.Errorf("dnssvcbrec %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

// dnssvcbrecRecordExists reports whether a record with the given identity is
// present in the get-all response.
func dnssvcbrecRecordExists(client *service.NitroClient, domain, targetname, priority, svcbtype string) (bool, error) {
	// dnssvcbrec's get requires the svcbtype key arg; a plain get-all returns empty.
	findParams := service.FindParams{
		ResourceType:             service.Dnssvcbrec.Type(),
		ArgsMap:                  map[string]string{"svcbtype": svcbtype},
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return false, nil
	}
	for _, v := range dataArr {
		if d, _ := v["domain"].(string); d != domain {
			continue
		}
		if tn, _ := v["targetname"].(string); tn != targetname {
			continue
		}
		if st, _ := v["svcbtype"].(string); st != svcbtype {
			continue
		}
		if pv, ok := v["priority"]; ok && pv != nil {
			if strings.TrimSpace(fmt.Sprintf("%v", pv)) != priority {
				continue
			}
		}
		return true, nil
	}
	return false, nil
}

// testAccCheckDnssvcbrecADCValue asserts an attribute's value directly on the
// appliance, proving the unset actually reverted it.
func testAccCheckDnssvcbrecADCValue(domain, targetname, svcbtype, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		findParams := service.FindParams{
			ResourceType: service.Dnssvcbrec.Type(),
			// svcbtype is a required get key; without it the get-all returns empty.
			ArgsMap: map[string]string{"svcbtype": svcbtype},
		}
		dataArray, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		for _, d := range dataArray {
			if dm, _ := d["domain"].(string); dm != domain {
				continue
			}
			if tn, _ := d["targetname"].(string); tn != targetname {
				continue
			}
			if st, _ := d["svcbtype"].(string); st != svcbtype {
				continue
			}
			got := strings.TrimSpace(fmt.Sprintf("%v", d[attr]))
			if got != want {
				return fmt.Errorf("dnssvcbrec %s,%s,%s: appliance attr %q = %q, want %q (unset did not revert it)", domain, targetname, svcbtype, attr, got, want)
			}
			return nil
		}
		return fmt.Errorf("dnssvcbrec %s,%s,%s not found on appliance", domain, targetname, svcbtype)
	}
}

const testAccDnssvcbrecDataSource_basic = `
resource "citrixadc_dnssvcbrec" "tf_dnssvcbrec_ds" {
	domain     = "svcb.example.com"
	targetname = "target.example.com"
	priority   = 1
	svcbtype   = "HTTPS"
	alpn       = "h2,h3"
	port       = 443
	ttl        = 3600
}

data "citrixadc_dnssvcbrec" "tf_dnssvcbrec_ds" {
	domain     = citrixadc_dnssvcbrec.tf_dnssvcbrec_ds.domain
	targetname = citrixadc_dnssvcbrec.tf_dnssvcbrec_ds.targetname
	depends_on = [citrixadc_dnssvcbrec.tf_dnssvcbrec_ds]
}
`

func TestAccDnssvcbrecDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssvcbrecDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnssvcbrec.tf_dnssvcbrec_ds", "domain", "svcb.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_dnssvcbrec.tf_dnssvcbrec_ds", "targetname", "target.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_dnssvcbrec.tf_dnssvcbrec_ds", "priority", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_dnssvcbrec.tf_dnssvcbrec_ds", "svcbtype", "HTTPS"),
					resource.TestCheckResourceAttr("data.citrixadc_dnssvcbrec.tf_dnssvcbrec_ds", "port", "443"),
					resource.TestCheckResourceAttr("data.citrixadc_dnssvcbrec.tf_dnssvcbrec_ds", "ttl", "3600"),
				),
			},
		},
	})
}
