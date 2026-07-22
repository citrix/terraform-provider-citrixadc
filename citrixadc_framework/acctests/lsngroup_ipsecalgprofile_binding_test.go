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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Note: the `bind lsn group` CLI family is marked deprecated by ADC but is still
// functional, so the binding is created normally with no special handling.
//
// =============================================================================
// Prerequisite for binding ANY ALG profile (IPSEC/SIP/RTSP) to an LSN group
// =============================================================================
// Binding an ALG profile fails with ADC errorcode 3520 / 3882 / 3883
// ("...IP Pool pair is not enabled" / "...the group is not in full cone mode or
//  IP Pool pair is not enabled") unless the LSN group is in FULL CONE mode AND
// has an IP-pool pair. Verified against the live testbed CLI man pages and NITRO,
// the complete recipe is:
//
//   1. lsnclient + lsnclient_network_binding  (subscriber subnet)
//   2. lsnpool   + lsnpool_lsnip_binding       (NAT IP range)
//   3. lsngroup  + lsngroup_lsnpool_binding    (NAT pool paired to group)
//   4. TWO lsnappsprofiles -- one TCP and one UDP -- each with
//      mapping=ENDPOINT-INDEPENDENT + filtering=ENDPOINT-INDEPENDENT (full-cone NAT)
//      AND ippooling=PAIRED (the "IP pool pair"), both bound to the group via
//      lsngroup_lsnappsprofile_binding.
//   5. lsnappsprofile_port_binding binding the full destination-port range 1-65535
//      to EACH of those full-cone profiles -- THIS is the step that actually flips
//      the group into full-cone mode for ALG binding (verified: without it the bind
//      fails with 3520/3882/3883). Both TCP and UDP full-cone coverage is required so
//      the single recipe works for IPSEC (UDP IKE), SIP (UDP) and RTSP (TCP control).
//
// IP ranges are safe for the standalone testbed:
//   - subscriber subnet: 100.64.0.0/10 (RFC 6598 CGN range)
//   - NAT IP range:      10.20.30.40-10.20.30.50 (mirrors lsnpool_lsnip_binding_test.go)
//
// IMPORTANT firmware quirk on this testbed: the standalone GET endpoints for
//   lsnpool_lsnip_binding  and  lsnappsprofile_port_binding
// return an empty body even when the binding exists (only the aggregate
// lsnpool_binding / lsnappsprofile_binding endpoints enumerate them). That breaks
// the Read of the corresponding managed Terraform resources, so those two bindings
// are provisioned OUT-OF-BAND in PreConfig via direct NITRO calls instead of as
// managed HCL resources. They are children of the managed lsnpool / lsnappsprofile
// resources, so Terraform's destroy of those parents removes them automatically.
// All other prerequisite resources are managed HCL (their Reads work fine).

// lsnGroupAlgPrereqConfig contains the managed-HCL portion of the prerequisite.
const lsnGroupAlgPrereqConfig = `
	resource "citrixadc_lsnclient" "tf_lsnclient" {
		clientname = "my_lsnclient"
	}

	resource "citrixadc_lsnclient_network_binding" "tf_lsnclient_network_binding" {
		clientname = citrixadc_lsnclient.tf_lsnclient.clientname
		network    = "100.64.0.0"
		netmask    = "255.192.0.0"
		td         = 0
	}

	resource "citrixadc_lsngroup" "tf_lsngroup" {
		groupname  = "my_lsngroup"
		clientname = citrixadc_lsnclient.tf_lsnclient.clientname
		nattype    = "DYNAMIC"
		depends_on = [citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding]
	}

	resource "citrixadc_lsnpool" "tf_lsnpool" {
		poolname            = "my_lsn_pool"
		nattype             = "DYNAMIC"
		portblockallocation = "DISABLED"
		maxportrealloctmq   = 50
		portrealloctimeout  = 50
	}

	resource "citrixadc_lsngroup_lsnpool_binding" "tf_lsngroup_lsnpool_binding" {
		groupname  = citrixadc_lsngroup.tf_lsngroup.groupname
		poolname   = citrixadc_lsnpool.tf_lsnpool.poolname
		depends_on = [citrixadc_lsngroup.tf_lsngroup, citrixadc_lsnpool.tf_lsnpool]
	}

	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile_udp" {
		appsprofilename   = "my_lsn_appsprofile_udp"
		transportprotocol = "UDP"
		mapping           = "ENDPOINT-INDEPENDENT"
		filtering         = "ENDPOINT-INDEPENDENT"
		ippooling         = "PAIRED"
	}

	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile_tcp" {
		appsprofilename   = "my_lsn_appsprofile_tcp"
		transportprotocol = "TCP"
		mapping           = "ENDPOINT-INDEPENDENT"
		filtering         = "ENDPOINT-INDEPENDENT"
		ippooling         = "PAIRED"
	}

	resource "citrixadc_lsngroup_lsnappsprofile_binding" "tf_lsngroup_lsnappsprofile_binding_udp" {
		groupname       = citrixadc_lsngroup.tf_lsngroup.groupname
		appsprofilename = citrixadc_lsnappsprofile.tf_lsnappsprofile_udp.appsprofilename
		depends_on      = [citrixadc_lsngroup.tf_lsngroup, citrixadc_lsnappsprofile.tf_lsnappsprofile_udp]
	}

	resource "citrixadc_lsngroup_lsnappsprofile_binding" "tf_lsngroup_lsnappsprofile_binding_tcp" {
		groupname       = citrixadc_lsngroup.tf_lsngroup.groupname
		appsprofilename = citrixadc_lsnappsprofile.tf_lsnappsprofile_tcp.appsprofilename
		depends_on      = [citrixadc_lsngroup.tf_lsngroup, citrixadc_lsnappsprofile.tf_lsnappsprofile_tcp]
	}
`

// lsnAlgPrereqParams holds the entity names used by the out-of-band prerequisite
// provisioning so the helper can be reused across the three ALG binding tests.
type lsnAlgPrereqParams struct {
	poolname           string
	lsnip              string
	appsprofilenameUDP string
	appsprofilenameTCP string
}

// lsnAlgFullPortRange covers all destination ports on the full-cone appsprofiles.
// Binding the whole range (rather than a single ALG-specific port) makes the group
// genuinely "full cone" so IPSEC/SIP/RTSP ALG profiles can all be bound.
const lsnAlgFullPortRange = "1-65535"

// provisionLsnAlgOutOfBandPrereq creates, via direct NITRO calls, the child
// bindings whose standalone GET endpoints are broken on this firmware and therefore
// cannot be managed as Terraform resources: lsnpool_lsnip_binding and
// lsnappsprofile_port_binding. They are children of managed resources (lsnpool /
// lsnappsprofile), so Terraform's destroy of those parents cleans them up. The
// parents are created by the managed HCL in the preceding test step; the AddResource
// calls below are idempotent safeguards.
func provisionLsnAlgOutOfBandPrereq(t *testing.T, p lsnAlgPrereqParams) {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		t.Fatalf("Failed to get test client for LSN ALG prereq: %v", err)
	}

	// Ensure parents exist (idempotent; managed HCL also creates these).
	_, _ = client.AddResource(service.Lsnpool.Type(), p.poolname, map[string]interface{}{
		"poolname": p.poolname,
		"nattype":  "DYNAMIC",
	})
	for _, ap := range []struct{ name, proto string }{
		{p.appsprofilenameUDP, "UDP"},
		{p.appsprofilenameTCP, "TCP"},
	} {
		_, _ = client.AddResource(service.Lsnappsprofile.Type(), ap.name, map[string]interface{}{
			"appsprofilename":   ap.name,
			"transportprotocol": ap.proto,
			"mapping":           "ENDPOINT-INDEPENDENT",
			"filtering":         "ENDPOINT-INDEPENDENT",
			"ippooling":         "PAIRED",
		})
	}

	// NAT IP range bound to the pool (broken-Read child binding).
	if err := client.UpdateUnnamedResource(service.Lsnpool_lsnip_binding.Type(), map[string]interface{}{
		"poolname": p.poolname,
		"lsnip":    p.lsnip,
	}); err != nil {
		t.Fatalf("Failed to bind lsnip %s to pool %s: %v", p.lsnip, p.poolname, err)
	}

	// Full destination-port range bound to BOTH full-cone appsprofiles (broken-Read
	// child bindings). This is what flips the group into full-cone mode so the ALG
	// profile bind succeeds.
	for _, name := range []string{p.appsprofilenameUDP, p.appsprofilenameTCP} {
		if err := client.UpdateUnnamedResource(service.Lsnappsprofile_port_binding.Type(), map[string]interface{}{
			"appsprofilename": name,
			"lsnport":         lsnAlgFullPortRange,
		}); err != nil {
			t.Fatalf("Failed to bind lsnport %s to appsprofile %s: %v", lsnAlgFullPortRange, name, err)
		}
	}
}

// ipsecAlgPrereqParams is the out-of-band prerequisite for the ipsec ALG test.
var ipsecAlgPrereqParams = lsnAlgPrereqParams{
	poolname:           "my_lsn_pool",
	lsnip:              "10.20.30.40-10.20.30.50",
	appsprofilenameUDP: "my_lsn_appsprofile_udp",
	appsprofilenameTCP: "my_lsn_appsprofile_tcp",
}

// step0: managed prereq parents only (no ALG binding yet). After this step applies,
// the out-of-band child bindings (lsnip + appsprofile port) are provisioned in the
// next step's PreConfig, flipping the group into full-cone mode.
const testAccLsngroup_ipsecalgprofile_binding_basic_step0 = lsnGroupAlgPrereqConfig + `
	resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
		name              = "my_ipsecalgprofile"
		ikesessiontimeout = 50
		espsessiontimeout = 20
		connfailover      = "DISABLED"
	}
`

const testAccLsngroup_ipsecalgprofile_binding_basic_step1 = lsnGroupAlgPrereqConfig + `
	resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
		name              = "my_ipsecalgprofile"
		ikesessiontimeout = 50
		espsessiontimeout = 20
		connfailover      = "DISABLED"
	}

	resource "citrixadc_lsngroup_ipsecalgprofile_binding" "tf_lsngroup_ipsecalgprofile_binding" {
		groupname       = citrixadc_lsngroup.tf_lsngroup.groupname
		ipsecalgprofile = citrixadc_ipsecalgprofile.tf_ipsecalgprofile.name
		depends_on = [
			citrixadc_lsngroup.tf_lsngroup,
			citrixadc_ipsecalgprofile.tf_ipsecalgprofile,
			citrixadc_lsngroup_lsnpool_binding.tf_lsngroup_lsnpool_binding,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_lsngroup_lsnappsprofile_binding_udp,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_lsngroup_lsnappsprofile_binding_tcp,
		]
	}
`

const testAccLsngroup_ipsecalgprofile_binding_basic_step2 = lsnGroupAlgPrereqConfig + `
	# Keep the participating resources but drop the binding to verify proper deletion.
	resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
		name              = "my_ipsecalgprofile"
		ikesessiontimeout = 50
		espsessiontimeout = 20
		connfailover      = "DISABLED"
	}
`

func TestAccLsngroup_ipsecalgprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_ipsecalgprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Create the managed prerequisite parents (client/subnet, pool,
				// group+pool, full-cone appsprofile+group binding, ipsecalgprofile).
				Config: testAccLsngroup_ipsecalgprofile_binding_basic_step0,
			},
			{
				// PreConfig provisions the two broken-Read child bindings out-of-band
				// (lsnip on the pool, UDP 500 on the full-cone appsprofile), enabling
				// full-cone mode / IP pool pair so the ALG bind can succeed.
				PreConfig: func() { provisionLsnAlgOutOfBandPrereq(t, ipsecAlgPrereqParams) },
				Config:    testAccLsngroup_ipsecalgprofile_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_ipsecalgprofile_bindingExist("citrixadc_lsngroup_ipsecalgprofile_binding.tf_lsngroup_ipsecalgprofile_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lsngroup_ipsecalgprofile_binding.tf_lsngroup_ipsecalgprofile_binding", "groupname", "my_lsngroup"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup_ipsecalgprofile_binding.tf_lsngroup_ipsecalgprofile_binding", "ipsecalgprofile", "my_ipsecalgprofile"),
				),
			},
			{
				Config: testAccLsngroup_ipsecalgprofile_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_ipsecalgprofile_bindingNotExist("citrixadc_lsngroup_ipsecalgprofile_binding.tf_lsngroup_ipsecalgprofile_binding", "my_lsngroup,my_ipsecalgprofile"),
				),
			},
		},
	})
}

func TestAccLsngroup_ipsecalgprofile_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lsngroup_ipsecalgprofile_binding.tf_lsngroup_ipsecalgprofile_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_ipsecalgprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Create the managed prerequisite parents first.
				Config: testAccLsngroup_ipsecalgprofile_binding_basic_step0,
			},
			{
				// Provision broken-Read child bindings out-of-band, then apply the
				// ALG binding so it exists in state ready to be imported.
				PreConfig: func() { provisionLsnAlgOutOfBandPrereq(t, ipsecAlgPrereqParams) },
				Config:    testAccLsngroup_ipsecalgprofile_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_ipsecalgprofile_bindingExist(resAddr, nil),
				),
			},
			{
				Config:                  testAccLsngroup_ipsecalgprofile_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckLsngroup_ipsecalgprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsngroup_ipsecalgprofile_binding id is set")
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

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"groupname", "ipsecalgprofile"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}

		groupname := idMap["groupname"]
		ipsecalgprofile := idMap["ipsecalgprofile"]

		findParams := service.FindParams{
			ResourceType:             service.Lsngroup_ipsecalgprofile_binding.Type(),
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ipsecalgprofile
		found := false
		for _, v := range dataArr {
			if val, ok := v["ipsecalgprofile"].(string); ok && val == ipsecalgprofile {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsngroup_ipsecalgprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_ipsecalgprofile_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// id is supplied in legacy comma-separated order: groupname,ipsecalgprofile
		idMap, _, err := utils.ParseIdString(id, []string{"groupname", "ipsecalgprofile"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}

		groupname := idMap["groupname"]
		ipsecalgprofile := idMap["ipsecalgprofile"]

		findParams := service.FindParams{
			ResourceType:             service.Lsngroup_ipsecalgprofile_binding.Type(),
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ipsecalgprofile
		found := false
		for _, v := range dataArr {
			if val, ok := v["ipsecalgprofile"].(string); ok && val == ipsecalgprofile {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsngroup_ipsecalgprofile_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_ipsecalgprofile_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsngroup_ipsecalgprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lsngroup_ipsecalgprofile_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsngroup_ipsecalgprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsngroup_ipsecalgprofile_binding_DataSource_basic = lsnGroupAlgPrereqConfig + `
	resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
		name              = "my_ipsecalgprofile"
		ikesessiontimeout = 50
		espsessiontimeout = 20
		connfailover      = "DISABLED"
	}

	resource "citrixadc_lsngroup_ipsecalgprofile_binding" "tf_lsngroup_ipsecalgprofile_binding" {
		groupname       = citrixadc_lsngroup.tf_lsngroup.groupname
		ipsecalgprofile = citrixadc_ipsecalgprofile.tf_ipsecalgprofile.name
		depends_on = [
			citrixadc_lsngroup.tf_lsngroup,
			citrixadc_ipsecalgprofile.tf_ipsecalgprofile,
			citrixadc_lsngroup_lsnpool_binding.tf_lsngroup_lsnpool_binding,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_lsngroup_lsnappsprofile_binding_udp,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_lsngroup_lsnappsprofile_binding_tcp,
		]
	}

	data "citrixadc_lsngroup_ipsecalgprofile_binding" "tf_lsngroup_ipsecalgprofile_binding" {
		groupname       = citrixadc_lsngroup.tf_lsngroup.groupname
		ipsecalgprofile = citrixadc_ipsecalgprofile.tf_ipsecalgprofile.name
		depends_on      = [citrixadc_lsngroup_ipsecalgprofile_binding.tf_lsngroup_ipsecalgprofile_binding]
	}
`

func TestAccLsngroup_ipsecalgprofile_binding_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_ipsecalgprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Create managed prereq parents first.
				Config: testAccLsngroup_ipsecalgprofile_binding_basic_step0,
			},
			{
				// Provision broken-Read child bindings out-of-band, then apply the
				// ALG binding + datasource.
				PreConfig: func() { provisionLsnAlgOutOfBandPrereq(t, ipsecAlgPrereqParams) },
				Config:    testAccLsngroup_ipsecalgprofile_binding_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_ipsecalgprofile_binding.tf_lsngroup_ipsecalgprofile_binding", "groupname", "my_lsngroup"),
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_ipsecalgprofile_binding.tf_lsngroup_ipsecalgprofile_binding", "ipsecalgprofile", "my_ipsecalgprofile"),
				),
			},
		},
	})
}
