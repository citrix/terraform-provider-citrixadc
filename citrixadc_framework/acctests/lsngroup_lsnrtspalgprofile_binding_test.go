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
// Prerequisite for binding an RTSP ALG profile to an LSN group: the group must be in
// FULL CONE mode AND have an IP-pool pair (else ADC errorcode 3883
// "...the group is not in full cone mode or IP Pool pair is not enabled"). See the
// detailed recipe in lsngroup_ipsecalgprofile_binding_test.go. The shared helper
// provisionLsnAlgOutOfBandPrereq and lsnAlgFullPortRange (also in that file) create
// the two broken-Read child bindings (lsnpool_lsnip_binding,
// lsnappsprofile_port_binding) out-of-band. This file uses rtspbind-specific entity
// names + a distinct subnet/NAT range to avoid collisions with sibling LSN tests.
//   - subscriber subnet uses CGN private range 100.65.0.0/16 (RFC 6598)
//   - NAT IP range uses 10.20.31.40-10.20.31.50

const lsnRtspGroupAlgPrereqConfig = `
	resource "citrixadc_lsnclient" "tf_rtspbind_lsnclient" {
		clientname = "my_rtspbind_lsnclient"
	}

	resource "citrixadc_lsnclient_network_binding" "tf_rtspbind_lsnclient_network_binding" {
		clientname = citrixadc_lsnclient.tf_rtspbind_lsnclient.clientname
		network    = "100.65.0.0"
		netmask    = "255.255.0.0"
		td         = 0
	}

	resource "citrixadc_lsngroup" "tf_rtspbind_lsngroup" {
		groupname  = "my_rtspbind_lsngroup"
		clientname = citrixadc_lsnclient.tf_rtspbind_lsnclient.clientname
		nattype    = "DYNAMIC"
		depends_on = [citrixadc_lsnclient_network_binding.tf_rtspbind_lsnclient_network_binding]
	}

	resource "citrixadc_lsnpool" "tf_rtspbind_lsnpool" {
		poolname            = "my_rtspbind_lsn_pool"
		nattype             = "DYNAMIC"
		portblockallocation = "DISABLED"
		maxportrealloctmq   = 50
		portrealloctimeout  = 50
	}

	resource "citrixadc_lsngroup_lsnpool_binding" "tf_rtspbind_lsngroup_lsnpool_binding" {
		groupname  = citrixadc_lsngroup.tf_rtspbind_lsngroup.groupname
		poolname   = citrixadc_lsnpool.tf_rtspbind_lsnpool.poolname
		depends_on = [citrixadc_lsngroup.tf_rtspbind_lsngroup, citrixadc_lsnpool.tf_rtspbind_lsnpool]
	}

	resource "citrixadc_lsnappsprofile" "tf_rtspbind_lsnappsprofile_udp" {
		appsprofilename   = "my_rtspbind_appsprofile_udp"
		transportprotocol = "UDP"
		mapping           = "ENDPOINT-INDEPENDENT"
		filtering         = "ENDPOINT-INDEPENDENT"
		ippooling         = "PAIRED"
	}

	resource "citrixadc_lsnappsprofile" "tf_rtspbind_lsnappsprofile_tcp" {
		appsprofilename   = "my_rtspbind_appsprofile_tcp"
		transportprotocol = "TCP"
		mapping           = "ENDPOINT-INDEPENDENT"
		filtering         = "ENDPOINT-INDEPENDENT"
		ippooling         = "PAIRED"
	}

	resource "citrixadc_lsngroup_lsnappsprofile_binding" "tf_rtspbind_lsngroup_lsnappsprofile_binding_udp" {
		groupname       = citrixadc_lsngroup.tf_rtspbind_lsngroup.groupname
		appsprofilename = citrixadc_lsnappsprofile.tf_rtspbind_lsnappsprofile_udp.appsprofilename
		depends_on      = [citrixadc_lsngroup.tf_rtspbind_lsngroup, citrixadc_lsnappsprofile.tf_rtspbind_lsnappsprofile_udp]
	}

	resource "citrixadc_lsngroup_lsnappsprofile_binding" "tf_rtspbind_lsngroup_lsnappsprofile_binding_tcp" {
		groupname       = citrixadc_lsngroup.tf_rtspbind_lsngroup.groupname
		appsprofilename = citrixadc_lsnappsprofile.tf_rtspbind_lsnappsprofile_tcp.appsprofilename
		depends_on      = [citrixadc_lsngroup.tf_rtspbind_lsngroup, citrixadc_lsnappsprofile.tf_rtspbind_lsnappsprofile_tcp]
	}
`

// rtspAlgPrereqParams is the out-of-band prerequisite for the rtsp ALG test.
var rtspAlgPrereqParams = lsnAlgPrereqParams{
	poolname:           "my_rtspbind_lsn_pool",
	lsnip:              "10.20.31.40-10.20.31.50",
	appsprofilenameUDP: "my_rtspbind_appsprofile_udp",
	appsprofilenameTCP: "my_rtspbind_appsprofile_tcp",
}

// step0: managed prereq parents only (no ALG binding yet).
const testAccLsngroup_lsnrtspalgprofile_binding_basic_step0 = lsnRtspGroupAlgPrereqConfig + `
	resource "citrixadc_lsnrtspalgprofile" "tf_rtspbind_lsnrtspalgprofile" {
		rtspalgprofilename = "my_rtspbind_lsnrtspalgprofile"
		rtspportrange      = 4200
		rtspidletimeout    = 150
	}
`

const testAccLsngroup_lsnrtspalgprofile_binding_basic_step1 = lsnRtspGroupAlgPrereqConfig + `
	resource "citrixadc_lsnrtspalgprofile" "tf_rtspbind_lsnrtspalgprofile" {
		rtspalgprofilename = "my_rtspbind_lsnrtspalgprofile"
		rtspportrange      = 4200
		rtspidletimeout    = 150
	}

	resource "citrixadc_lsngroup_lsnrtspalgprofile_binding" "tf_rtspbind_lsngroup_lsnrtspalgprofile_binding" {
		groupname          = citrixadc_lsngroup.tf_rtspbind_lsngroup.groupname
		rtspalgprofilename = citrixadc_lsnrtspalgprofile.tf_rtspbind_lsnrtspalgprofile.rtspalgprofilename
		depends_on = [
			citrixadc_lsngroup.tf_rtspbind_lsngroup,
			citrixadc_lsnrtspalgprofile.tf_rtspbind_lsnrtspalgprofile,
			citrixadc_lsngroup_lsnpool_binding.tf_rtspbind_lsngroup_lsnpool_binding,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_rtspbind_lsngroup_lsnappsprofile_binding_udp,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_rtspbind_lsngroup_lsnappsprofile_binding_tcp,
		]
	}
`

const testAccLsngroup_lsnrtspalgprofile_binding_basic_step2 = lsnRtspGroupAlgPrereqConfig + `
	# Keep the participating resources but drop the binding to verify proper deletion.
	resource "citrixadc_lsnrtspalgprofile" "tf_rtspbind_lsnrtspalgprofile" {
		rtspalgprofilename = "my_rtspbind_lsnrtspalgprofile"
		rtspportrange      = 4200
		rtspidletimeout    = 150
	}
`

func TestAccLsngroup_lsnrtspalgprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_lsnrtspalgprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnrtspalgprofile_binding_basic_step0,
			},
			{
				PreConfig: func() { provisionLsnAlgOutOfBandPrereq(t, rtspAlgPrereqParams) },
				Config:    testAccLsngroup_lsnrtspalgprofile_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnrtspalgprofile_bindingExist("citrixadc_lsngroup_lsnrtspalgprofile_binding.tf_rtspbind_lsngroup_lsnrtspalgprofile_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lsngroup_lsnrtspalgprofile_binding.tf_rtspbind_lsngroup_lsnrtspalgprofile_binding", "groupname", "my_rtspbind_lsngroup"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup_lsnrtspalgprofile_binding.tf_rtspbind_lsngroup_lsnrtspalgprofile_binding", "rtspalgprofilename", "my_rtspbind_lsnrtspalgprofile"),
				),
			},
			{
				Config: testAccLsngroup_lsnrtspalgprofile_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnrtspalgprofile_bindingNotExist("citrixadc_lsngroup_lsnrtspalgprofile_binding.tf_rtspbind_lsngroup_lsnrtspalgprofile_binding", "my_rtspbind_lsngroup,my_rtspbind_lsnrtspalgprofile"),
				),
			},
		},
	})
}

func TestAccLsngroup_lsnrtspalgprofile_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lsngroup_lsnrtspalgprofile_binding.tf_rtspbind_lsngroup_lsnrtspalgprofile_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_lsnrtspalgprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnrtspalgprofile_binding_basic_step0,
			},
			{
				PreConfig: func() { provisionLsnAlgOutOfBandPrereq(t, rtspAlgPrereqParams) },
				Config:    testAccLsngroup_lsnrtspalgprofile_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsnrtspalgprofile_bindingExist(resAddr, nil),
				),
			},
			{
				Config:                  testAccLsngroup_lsnrtspalgprofile_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckLsngroup_lsnrtspalgprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsngroup_lsnrtspalgprofile_binding id is set")
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

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"groupname", "rtspalgprofilename"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}

		groupname := idMap["groupname"]
		rtspalgprofilename := idMap["rtspalgprofilename"]

		findParams := service.FindParams{
			ResourceType:             service.Lsngroup_lsnrtspalgprofile_binding.Type(),
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching rtspalgprofilename
		found := false
		for _, v := range dataArr {
			if val, ok := v["rtspalgprofilename"].(string); ok && val == rtspalgprofilename {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsngroup_lsnrtspalgprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_lsnrtspalgprofile_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// id is supplied in legacy comma-separated order: groupname,rtspalgprofilename
		idMap, _, err := utils.ParseIdString(id, []string{"groupname", "rtspalgprofilename"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}

		groupname := idMap["groupname"]
		rtspalgprofilename := idMap["rtspalgprofilename"]

		findParams := service.FindParams{
			ResourceType:             service.Lsngroup_lsnrtspalgprofile_binding.Type(),
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching rtspalgprofilename
		found := false
		for _, v := range dataArr {
			if val, ok := v["rtspalgprofilename"].(string); ok && val == rtspalgprofilename {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsngroup_lsnrtspalgprofile_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_lsnrtspalgprofile_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsngroup_lsnrtspalgprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lsngroup_lsnrtspalgprofile_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsngroup_lsnrtspalgprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsngroup_lsnrtspalgprofile_binding_DataSource_basic = lsnRtspGroupAlgPrereqConfig + `
	resource "citrixadc_lsnrtspalgprofile" "tf_rtspbind_lsnrtspalgprofile" {
		rtspalgprofilename = "my_rtspbind_lsnrtspalgprofile"
		rtspportrange      = 4200
		rtspidletimeout    = 150
	}

	resource "citrixadc_lsngroup_lsnrtspalgprofile_binding" "tf_rtspbind_lsngroup_lsnrtspalgprofile_binding" {
		groupname          = citrixadc_lsngroup.tf_rtspbind_lsngroup.groupname
		rtspalgprofilename = citrixadc_lsnrtspalgprofile.tf_rtspbind_lsnrtspalgprofile.rtspalgprofilename
		depends_on = [
			citrixadc_lsngroup.tf_rtspbind_lsngroup,
			citrixadc_lsnrtspalgprofile.tf_rtspbind_lsnrtspalgprofile,
			citrixadc_lsngroup_lsnpool_binding.tf_rtspbind_lsngroup_lsnpool_binding,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_rtspbind_lsngroup_lsnappsprofile_binding_udp,
			citrixadc_lsngroup_lsnappsprofile_binding.tf_rtspbind_lsngroup_lsnappsprofile_binding_tcp,
		]
	}

	data "citrixadc_lsngroup_lsnrtspalgprofile_binding" "tf_rtspbind_lsngroup_lsnrtspalgprofile_binding" {
		groupname          = citrixadc_lsngroup.tf_rtspbind_lsngroup.groupname
		rtspalgprofilename = citrixadc_lsnrtspalgprofile.tf_rtspbind_lsnrtspalgprofile.rtspalgprofilename
		depends_on         = [citrixadc_lsngroup_lsnrtspalgprofile_binding.tf_rtspbind_lsngroup_lsnrtspalgprofile_binding]
	}
`

func TestAccLsngroup_lsnrtspalgprofile_binding_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroup_lsnrtspalgprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsnrtspalgprofile_binding_basic_step0,
			},
			{
				PreConfig: func() { provisionLsnAlgOutOfBandPrereq(t, rtspAlgPrereqParams) },
				Config:    testAccLsngroup_lsnrtspalgprofile_binding_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_lsnrtspalgprofile_binding.tf_rtspbind_lsngroup_lsnrtspalgprofile_binding", "groupname", "my_rtspbind_lsngroup"),
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup_lsnrtspalgprofile_binding.tf_rtspbind_lsngroup_lsnrtspalgprofile_binding", "rtspalgprofilename", "my_rtspbind_lsnrtspalgprofile"),
				),
			},
		},
	})
}
