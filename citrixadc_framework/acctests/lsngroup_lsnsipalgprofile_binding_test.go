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
// Prerequisite for binding a SIP ALG profile to an LSN group: the group must be in
// FULL CONE mode AND have an IP-pool pair (else ADC errorcode 3882
// "...the group is not in full cone mode or IP Pool pair is not enabled"). See the
// detailed recipe in lsngroup_ipsecalgprofile_binding_test.go. The shared helper
// provisionLsnAlgOutOfBandPrereq and lsnAlgFullPortRange (also in that file) create
// the two broken-Read child bindings (lsnpool_lsnip_binding,
// lsnappsprofile_port_binding) out-of-band. This file uses sipbind-specific entity
// names + a distinct subnet/NAT range to avoid collisions with sibling LSN tests.
//   - subscriber subnet uses CGN private range 100.66.0.0/16 (RFC 6598)
//   - NAT IP range uses 10.20.32.40-10.20.32.50
//
// NOTE: the SIP ALG profile uses siptransportprotocol = UDP so that the UDP
// full-cone appsprofile governs its data path; with TCP transport the bind is
// rejected for not being in full cone mode on this firmware.

const lsnSipGroupAlgPrereqConfig = `
	resource "citrixadc_lsnclient" "tf_sipbind_lsnclient" {
		clientname = "my_sipbind_lsnclient"
	}

	resource "citrixadc_lsnclient_network_binding" "tf_sipbind_lsnclient_network_binding" {
		clientname = citrixadc_lsnclient.tf_sipbind_lsnclient.clientname
		network    = "100.66.0.0"
		netmask    = "255.255.0.0"
		td         = 0
	}

	resource "citrixadc_lsngroup" "tf_sipbind_lsngroup" {
		groupname  = "my_sipbind_lsngroup"
		clientname = citrixadc_lsnclient.tf_sipbind_lsnclient.clientname
		nattype    = "DYNAMIC"
		depends_on = [citrixadc_lsnclient_network_binding.tf_sipbind_lsnclient_network_binding]
	}

	resource "citrixadc_lsnpool" "tf_sipbind_lsnpool" {
		poolname            = "my_sipbind_lsn_pool"
		nattype             = "DYNAMIC"
		portblockallocation = "DISABLED"
		maxportrealloctmq   = 50
		portrealloctimeout  = 50
	}

	resource "citrixadc_lsngroup_lsnpool_binding" "tf_sipbind_lsngroup_lsnpool_binding" {
		groupname  = citrixadc_lsngroup.tf_sipbind_lsngroup.groupname
		poolname   = citrixadc_lsnpool.tf_sipbind_lsnpool.poolname
		depends_on = [citrixadc_lsngroup.tf_sipbind_lsngroup, citrixadc_lsnpool.tf_sipbind_lsnpool]
	}

	resource "citrixadc_lsnappsprofile" "tf_sipbind_lsnappsprofile_udp" {
		appsprofilename   = "my_sipbind_appsprofile_udp"
		transportprotocol = "UDP"
		mapping           = "ENDPOINT-INDEPENDENT"
		filtering         = "ENDPOINT-INDEPENDENT"
		ippooling         = "PAIRED"
	}

	resource "citrixadc_lsnappsprofile" "tf_sipbind_lsnappsprofile_tcp" {
		appsprofilename   = "my_sipbind_appsprofile_tcp"
		transportprotocol = "TCP"
		mapping           = "ENDPOINT-INDEPENDENT"
		filtering         = "ENDPOINT-INDEPENDENT"
		ippooling         = "PAIRED"
	}

	resource "citrixadc_lsngroup_lsnappsprofile_binding" "tf_sipbind_lsngroup_lsnappsprofile_binding_udp" {
		groupname       = citrixadc_lsngroup.tf_sipbind_lsngroup.groupname
		appsprofilename = citrixadc_lsnappsprofile.tf_sipbind_lsnappsprofile_udp.appsprofilename
		depends_on      = [citrixadc_lsngroup.tf_sipbind_lsngroup, citrixadc_lsnappsprofile.tf_sipbind_lsnappsprofile_udp]
	}

	resource "citrixadc_lsngroup_lsnappsprofile_binding" "tf_sipbind_lsngroup_lsnappsprofile_binding_tcp" {
		groupname       = citrixadc_lsngroup.tf_sipbind_lsngroup.groupname
		appsprofilename = citrixadc_lsnappsprofile.tf_sipbind_lsnappsprofile_tcp.appsprofilename
		depends_on      = [citrixadc_lsngroup.tf_sipbind_lsngroup, citrixadc_lsnappsprofile.tf_sipbind_lsnappsprofile_tcp]
	}
`

// sipAlgPrereqParams is the out-of-band prerequisite for the sip ALG test.
var sipAlgPrereqParams = lsnAlgPrereqParams{
	poolname:           "my_sipbind_lsn_pool",
	lsnip:              "10.20.32.40-10.20.32.50",
	appsprofilenameUDP: "my_sipbind_appsprofile_udp",
	appsprofilenameTCP: "my_sipbind_appsprofile_tcp",
}

// step0: managed prereq parents only (no ALG binding yet).
const testAccLsngroup_lsnsipalgprofile_binding_basic_step0 = lsnSipGroupAlgPrereqConfig + `
	resource "citrixadc_lsnsipalgprofile" "tf_sipbind_lsnsipalgprofile" {
		sipalgprofilename      = "my_sipbind_lsnsipalgprofile"
		datasessionidletimeout = 150
		sipsessiontimeout      = 150
		registrationtimeout    = 150
		sipsrcportrange        = "4400"
		siptransportprotocol   = "UDP"
	}
`

const testAccLsngroup_lsnsipalgprofile_binding_basic_step1 = lsnSipGroupAlgPrereqConfig + `
	resource "citrixadc_lsnsipalgprofile" "tf_sipbind_lsnsipalgprofile" {
		sipalgprofilename      = "my_sipbind_lsnsipalgprofile"
		datasessionidletimeout = 150
		sipsessiontimeout      = 150
		registrationtimeout    = 150
		sipsrcportrange        = "4400"
		siptransportprotocol   = "UDP"
	}

	resource "citrixadc_lsngroup_lsnsipalgprofile_binding" "tf_sipbind_binding" {
		groupname         = citrixadc_lsngroup.tf_sipbind_lsngroup.groupname
		sipalgprofilename = citrixadc_lsnsipalgprofile.tf_sipbind_lsnsipalgprofile.sipalgprofilename
		depends_on = [
			citrixadc_lsngroup.tf_sipbind_lsngroup,
			citrixadc_lsnsipalgprofile.tf_sipbind_lsnsipalgprofile,
			citrixadc_lsngroup_lsnpool_binding.tf_sipbind_lsngroup_lsnpool_binding,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_sipbind_lsngroup_lsnappsprofile_binding_udp,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_sipbind_lsngroup_lsnappsprofile_binding_tcp,
		]
	}
`

const testAccLsngroup_lsnsipalgprofile_binding_basic_step2 = lsnSipGroupAlgPrereqConfig + `
	# Keep the participating resources but drop the binding to verify proper deletion.
	resource "citrixadc_lsnsipalgprofile" "tf_sipbind_lsnsipalgprofile" {
		sipalgprofilename      = "my_sipbind_lsnsipalgprofile"
		datasessionidletimeout = 150
		sipsessiontimeout      = 150
		registrationtimeout    = 150
		sipsrcportrange        = "4400"
		siptransportprotocol   = "UDP"
	}
`

func TestAccLsngroup_lsnsipalgprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_lsnsipalgprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnsipalgprofile_binding_basic_step0,
			},
			{
				PreConfig: func() { provisionLsnAlgOutOfBandPrereq(t, sipAlgPrereqParams) },
				Config:    testAccLsngroup_lsnsipalgprofile_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnsipalgprofile_bindingExist("citrixadc_lsngroup_lsnsipalgprofile_binding.tf_sipbind_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lsngroup_lsnsipalgprofile_binding.tf_sipbind_binding", "groupname", "my_sipbind_lsngroup"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup_lsnsipalgprofile_binding.tf_sipbind_binding", "sipalgprofilename", "my_sipbind_lsnsipalgprofile"),
				),
			},
			{
				Config: testAccLsngroup_lsnsipalgprofile_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnsipalgprofile_bindingNotExist("citrixadc_lsngroup_lsnsipalgprofile_binding.tf_sipbind_binding", "my_sipbind_lsngroup,my_sipbind_lsnsipalgprofile"),
				),
			},
		},
	})
}

func TestAccLsngroup_lsnsipalgprofile_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lsngroup_lsnsipalgprofile_binding.tf_sipbind_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_lsnsipalgprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnsipalgprofile_binding_basic_step0,
			},
			{
				PreConfig: func() { provisionLsnAlgOutOfBandPrereq(t, sipAlgPrereqParams) },
				Config:    testAccLsngroup_lsnsipalgprofile_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnsipalgprofile_bindingExist(resAddr, nil),
				),
			},
			{
				Config:                  testAccLsngroup_lsnsipalgprofile_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckLsngroup_lsnsipalgprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsngroup_lsnsipalgprofile_binding id is set")
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

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"groupname", "sipalgprofilename"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}

		groupname := idMap["groupname"]
		sipalgprofilename := idMap["sipalgprofilename"]

		findParams := service.FindParams{
			ResourceType:             service.Lsngroup_lsnsipalgprofile_binding.Type(),
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching sipalgprofilename
		found := false
		for _, v := range dataArr {
			if val, ok := v["sipalgprofilename"].(string); ok && val == sipalgprofilename {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsngroup_lsnsipalgprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_lsnsipalgprofile_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// id is supplied in legacy comma-separated order: groupname,sipalgprofilename
		idMap, _, err := utils.ParseIdString(id, []string{"groupname", "sipalgprofilename"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}

		groupname := idMap["groupname"]
		sipalgprofilename := idMap["sipalgprofilename"]

		findParams := service.FindParams{
			ResourceType:             service.Lsngroup_lsnsipalgprofile_binding.Type(),
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching sipalgprofilename
		found := false
		for _, v := range dataArr {
			if val, ok := v["sipalgprofilename"].(string); ok && val == sipalgprofilename {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsngroup_lsnsipalgprofile_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_lsnsipalgprofile_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsngroup_lsnsipalgprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lsngroup_lsnsipalgprofile_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsngroup_lsnsipalgprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsngroup_lsnsipalgprofile_binding_DataSource_basic = lsnSipGroupAlgPrereqConfig + `
	resource "citrixadc_lsnsipalgprofile" "tf_sipbind_lsnsipalgprofile" {
		sipalgprofilename      = "my_sipbind_lsnsipalgprofile"
		datasessionidletimeout = 150
		sipsessiontimeout      = 150
		registrationtimeout    = 150
		sipsrcportrange        = "4400"
		siptransportprotocol   = "UDP"
	}

	resource "citrixadc_lsngroup_lsnsipalgprofile_binding" "tf_sipbind_binding" {
		groupname         = citrixadc_lsngroup.tf_sipbind_lsngroup.groupname
		sipalgprofilename = citrixadc_lsnsipalgprofile.tf_sipbind_lsnsipalgprofile.sipalgprofilename
		depends_on = [
			citrixadc_lsngroup.tf_sipbind_lsngroup,
			citrixadc_lsnsipalgprofile.tf_sipbind_lsnsipalgprofile,
			citrixadc_lsngroup_lsnpool_binding.tf_sipbind_lsngroup_lsnpool_binding,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_sipbind_lsngroup_lsnappsprofile_binding_udp,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_sipbind_lsngroup_lsnappsprofile_binding_tcp,
		]
	}

	data "citrixadc_lsngroup_lsnsipalgprofile_binding" "tf_sipbind_binding" {
		groupname         = citrixadc_lsngroup.tf_sipbind_lsngroup.groupname
		sipalgprofilename = citrixadc_lsnsipalgprofile.tf_sipbind_lsnsipalgprofile.sipalgprofilename
		depends_on        = [citrixadc_lsngroup_lsnsipalgprofile_binding.tf_sipbind_binding]
	}
`

func TestAccLsngroup_lsnsipalgprofile_binding_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_lsnsipalgprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnsipalgprofile_binding_basic_step0,
			},
			{
				PreConfig: func() { provisionLsnAlgOutOfBandPrereq(t, sipAlgPrereqParams) },
				Config:    testAccLsngroup_lsnsipalgprofile_binding_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_lsnsipalgprofile_binding.tf_sipbind_binding", "groupname", "my_sipbind_lsngroup"),
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_lsnsipalgprofile_binding.tf_sipbind_binding", "sipalgprofilename", "my_sipbind_lsnsipalgprofile"),
				),
			},
		},
	})
}
