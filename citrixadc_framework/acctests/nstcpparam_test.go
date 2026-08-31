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

const testAccNstcpparam_basic = `
	resource "citrixadc_nstcpparam" "tf_tcpparam" {
		
		
	}
`

var testAccNstcpparam_zero_values_map = map[string]string{
	"maxpktpermss":              "0",
	"mptcpmaxpendingsf":         "0",
	"mptcppendingjointhreshold": "0",
	"mptcpsfreplacetimeout":     "0",
	"mptcpsftimeout":            "0",
	"oooqsize":                  "0",
	"tcpfastopencookietimeout":  "0",
	"wsval":                     "0",
	"rfc5961chlgacklimit":       "0",
}

var testAccNstcpparam_default_values_map = map[string]string{
	"ws":                                  "ENABLED",
	"wsval":                               "8",
	"sack":                                "ENABLED",
	"learnvsvrmss":                        "DISABLED",
	"maxburst":                            "6",
	"initialcwnd":                         "10",
	"delayedack":                          "100",
	"downstaterst":                        "DISABLED",
	"nagle":                               "DISABLED",
	"limitedpersist":                      "ENABLED",
	"oooqsize":                            "300",
	"ackonpush":                           "ENABLED",
	"maxpktpermss":                        "0",
	"pktperretx":                          "1",
	"minrto":                              "1000",
	"slowstartincr":                       "2",
	"maxdynserverprobes":                  "7",
	"synholdfastgiveup":                   "1024",
	"maxsynholdperprobe":                  "128",
	"maxsynhold":                          "16384",
	"msslearninterval":                    "180",
	"msslearndelay":                       "3600",
	"maxtimewaitconn":                     "7000",
	"maxsynackretx":                       "100",
	"synattackdetection":                  "ENABLED",
	"connflushifnomem":                    "NONE ",
	"connflushthres":                      "4294967295",
	"mptcpconcloseonpassivesf":            "ENABLED",
	"mptcpchecksum":                       "ENABLED",
	"mptcpsftimeout":                      "0",
	"mptcpsfreplacetimeout":               "10",
	"mptcpmaxsf":                          "4",
	"mptcpmaxpendingsf":                   "4",
	"mptcppendingjointhreshold":           "0",
	"mptcprtostoswitchsf":                 "2",
	"mptcpusebackupondss":                 "ENABLED",
	"tcpmaxretries":                       "7",
	"mptcpimmediatesfcloseonfin":          "DISABLED",
	"mptcpclosemptcpsessiononlastsfclose": "DISABLED",
	"tcpfastopencookietimeout":            "0",
	"autosyncookietimeout":                "30",
	"tcpfintimeout":                       "40",
	"rfc5961chlgacklimit":                 "0",
	"mptcpsendsfresetoption":              "DISABLED",
	"mptcpreliableaddaddr":                "DISABLED",
	"mptcpfastcloseoption":                "ACK",
	"enhancedisngeneration":               "DISABLED",
	"delinkclientserveronrst":             "DISABLED",
	"compacttcpoptionnoop":                "DISABLED",
}

var testAccNstcpparam_non_default_values_map = map[string]string{
	"ws":                                  "DISABLED",
	"wsval":                               "9",
	"sack":                                "DISABLED",
	"learnvsvrmss":                        "ENABLED",
	"maxburst":                            "7",
	"initialcwnd":                         "11",
	"delayedack":                          "110",
	"downstaterst":                        "ENABLED",
	"nagle":                               "ENABLED",
	"limitedpersist":                      "DISABLED",
	"oooqsize":                            "500",
	"ackonpush":                           "DISABLED",
	"maxpktpermss":                        "10",
	"pktperretx":                          "2",
	"minrto":                              "2000",
	"slowstartincr":                       "4",
	"maxdynserverprobes":                  "8",
	"synholdfastgiveup":                   "2046",
	"maxsynholdperprobe":                  "130",
	"maxsynhold":                          "20000",
	"msslearninterval":                    "200",
	"msslearndelay":                       "4000",
	"maxtimewaitconn":                     "6000",
	"maxsynackretx":                       "200",
	"synattackdetection":                  "DISABLED",
	"connflushifnomem":                    "FIFO",
	"connflushthres":                      "4294967290",
	"mptcpconcloseonpassivesf":            "DISABLED",
	"mptcpchecksum":                       "DISABLED",
	"mptcpsftimeout":                      "10",
	"mptcpsfreplacetimeout":               "20",
	"mptcpmaxsf":                          "6",
	"mptcpmaxpendingsf":                   "2",
	"mptcppendingjointhreshold":           "10",
	"mptcprtostoswitchsf":                 "4",
	"mptcpusebackupondss":                 "DISABLED",
	"tcpmaxretries":                       "6",
	"mptcpimmediatesfcloseonfin":          "ENABLED",
	"mptcpclosemptcpsessiononlastsfclose": "ENABLED",
	"tcpfastopencookietimeout":            "10",
	"autosyncookietimeout":                "40",
	"tcpfintimeout":                       "20",
	"rfc5961chlgacklimit":                 "2200",
	"mptcpsendsfresetoption":              "ENABLED",
	"mptcpreliableaddaddr":                "ENABLED",
	"mptcpfastcloseoption":                "RESET",
	"enhancedisngeneration":               "ENABLED",
	"delinkclientserveronrst":             "ENABLED",
	"compacttcpoptionnoop":                "ENABLED",
}

const testAccNstcpparam_zero_values = `
resource "citrixadc_nstcpparam" "tf_tcpparam" {
	maxpktpermss = 0
	mptcpmaxpendingsf = 0
	mptcppendingjointhreshold = 0
	mptcpsfreplacetimeout = 0
	mptcpsftimeout = 0
	oooqsize = 0
	tcpfastopencookietimeout = 0
	wsval = 0
	rfc5961chlgacklimit = 0
}
`

const testAccNstcpparam_default_values = `
resource "citrixadc_nstcpparam" "tf_tcpparam" {
        ws = "ENABLED"
        wsval = 8
        sack = "ENABLED"
        learnvsvrmss = "DISABLED"
        maxburst = 6
        initialcwnd = 10
        delayedack = 100
        downstaterst = "DISABLED"
        nagle = "DISABLED"
        limitedpersist = "ENABLED"
        oooqsize = 300
        ackonpush = "ENABLED"
        maxpktpermss = 0
        pktperretx = 1
        minrto = 1000
        slowstartincr = 2
        maxdynserverprobes = 7
        synholdfastgiveup = 1024
        maxsynholdperprobe = 128
        maxsynhold = 16384 
        msslearninterval = 180
        msslearndelay = 3600
        maxtimewaitconn = 7000
        maxsynackretx = 100
        synattackdetection = "ENABLED"
        connflushifnomem = "NONE "
        connflushthres = 4294967295
        mptcpconcloseonpassivesf = "ENABLED"
        mptcpchecksum = "ENABLED"
        mptcpsftimeout = 0
        mptcpsfreplacetimeout = 10
        mptcpmaxsf = 4
        mptcpmaxpendingsf = 4
        mptcppendingjointhreshold = 0
        mptcprtostoswitchsf = 2
        mptcpusebackupondss = "ENABLED"
        tcpmaxretries = 7
        mptcpimmediatesfcloseonfin = "DISABLED"
        mptcpclosemptcpsessiononlastsfclose = "DISABLED"
        tcpfastopencookietimeout = 0
        autosyncookietimeout = 30
        tcpfintimeout = 40
		rfc5961chlgacklimit = 0
		mptcpsendsfresetoption = "DISABLED"
		mptcpreliableaddaddr = "DISABLED"
		mptcpfastcloseoption = "ACK"
		enhancedisngeneration = "DISABLED"
		delinkclientserveronrst = "DISABLED"
		compacttcpoptionnoop = "DISABLED"
}
`

const testAccNstcpparam_non_default_values = `
resource "citrixadc_nstcpparam" "tf_tcpparam" {
        ws = "DISABLED"
        wsval = 9
        sack = "DISABLED"
        learnvsvrmss = "ENABLED"
        maxburst = 7
        initialcwnd = 11
        delayedack = 110
        downstaterst = "ENABLED"
        nagle = "ENABLED"
        limitedpersist = "DISABLED"
        oooqsize = 500
        ackonpush = "DISABLED"
        maxpktpermss = 10
        pktperretx = 2
        minrto = 2000
        slowstartincr = 4
        maxdynserverprobes = 8
        synholdfastgiveup = 2046
        maxsynholdperprobe = 130
        maxsynhold = 20000
        msslearninterval = 200
        msslearndelay = 4000
        maxtimewaitconn = 6000
        maxsynackretx = 200
        synattackdetection = "DISABLED"
        connflushifnomem = "FIFO"
        connflushthres = 4294967290
        mptcpconcloseonpassivesf = "DISABLED"
        mptcpchecksum = "DISABLED"
        mptcpsftimeout = 10
        mptcpsfreplacetimeout = 20
        mptcpmaxsf = 6
        mptcpmaxpendingsf = 2
        mptcppendingjointhreshold = 10
        mptcprtostoswitchsf = 4
        mptcpusebackupondss = "DISABLED"
        tcpmaxretries = 6
        mptcpimmediatesfcloseonfin = "ENABLED"
        mptcpclosemptcpsessiononlastsfclose = "ENABLED"
        tcpfastopencookietimeout = 10
        autosyncookietimeout = 40
        tcpfintimeout = 20
		rfc5961chlgacklimit = 2200
		mptcpsendsfresetoption = "ENABLED"
		mptcpreliableaddaddr = "ENABLED"
		mptcpfastcloseoption = "RESET"
		enhancedisngeneration = "ENABLED"
		delinkclientserveronrst = "ENABLED"
		compacttcpoptionnoop = "ENABLED"
}
`

func TestAccNstcpparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstcpparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNstcpparam_zero_values,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpparamExist("citrixadc_nstcpparam.tf_tcpparam", nil),
					testAccCheckTcpparamMapvalues(testAccNstcpparam_zero_values_map),
				),
			},
			{
				Config: testAccNstcpparam_non_default_values,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpparamExist("citrixadc_nstcpparam.tf_tcpparam", nil),
					testAccCheckTcpparamMapvalues(testAccNstcpparam_non_default_values_map),
				),
			},
			{
				Config: testAccNstcpparam_default_values,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpparamExist("citrixadc_nstcpparam.tf_tcpparam", nil),
					testAccCheckTcpparamMapvalues(testAccNstcpparam_default_values_map),
				),
			},
		},
	})
}

func TestAccNstcpparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNstcpparamDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccNstcpparam_zero_values,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpparamExist("citrixadc_nstcpparam.tf_tcpparam", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNstcpparam_zero_values,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpparamExist("citrixadc_nstcpparam.tf_tcpparam", nil),
				),
			},
		},
	})
}

func testAccCheckNstcpparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nstcpparam name is set")
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
		data, err := client.FindResource(service.Nstcpparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nstcpparam %s not found", n)
		}

		return nil
	}
}

func testAccCheckTcpparamMapvalues(mapData map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType: "nstcpparam",
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		data := dataArr[0]

		// Compare fetched values to map data
		for k, v := range mapData {
			if _, ok := data[k]; !ok {
				return fmt.Errorf("key %s does not exist in fetched data", k)
			}

			// Get the str repesentation to avoid type mismatches
			str_val := fmt.Sprintf("%v", data[k])
			if str_val != v {
				return fmt.Errorf("key %s value differs. Fetched:%s, Map:%s", k, str_val, v)
			}
		}
		return nil
	}
}

func testAccCheckNstcpparamDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nstcpparam" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nstcpparam.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nstcpparam %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNstcpparam_import(t *testing.T) {
	const resAddr = "citrixadc_nstcpparam.tf_tcpparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstcpparamDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNstcpparam_zero_values},
			{
				Config:                  testAccNstcpparam_zero_values,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNstcpparamDataSource_basic = `

	resource "citrixadc_nstcpparam" "tf_nstcpparam" {
		delayedack = 100
		maxburst   = 6
	}

	data "citrixadc_nstcpparam" "tf_nstcpparam_data" {
		depends_on = [citrixadc_nstcpparam.tf_nstcpparam]
	}
`

// The nstcpparam unset test covers the unset-eligible attributes that carry a
// schema Default so that removing them from config routes through Update and
// triggers a NITRO unset (revert to the documented default). All are global TCP
// parameters with no cross-attribute prerequisite.
const testAccNstcpparam_unset_step1 = `
resource "citrixadc_nstcpparam" "tf_unset" {
	mptcpmaxpendingsf         = 2
	mptcppendingjointhreshold = 10
	mptcpsfreplacetimeout     = 20
	mptcpsftimeout            = 10
	oooqsize                  = 500
	rfc5961chlgacklimit       = 2200
	tcpfastopencookietimeout  = 130
	wsval                     = 9
}
`

const testAccNstcpparam_unset_step2 = `
resource "citrixadc_nstcpparam" "tf_unset" {
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccNstcpparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstcpparamDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNstcpparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpparamExist("citrixadc_nstcpparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "mptcpmaxpendingsf", "2"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "mptcppendingjointhreshold", "10"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "mptcpsfreplacetimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "mptcpsftimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "oooqsize", "500"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "rfc5961chlgacklimit", "2200"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "tcpfastopencookietimeout", "130"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "wsval", "9"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNstcpparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpparamExist("citrixadc_nstcpparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "mptcpmaxpendingsf", "4"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "mptcppendingjointhreshold", "0"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "mptcpsfreplacetimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "mptcpsftimeout", "0"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "oooqsize", "300"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "rfc5961chlgacklimit", "0"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "tcpfastopencookietimeout", "0"),
					resource.TestCheckResourceAttr("citrixadc_nstcpparam.tf_unset", "wsval", "8"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNstcpparamADCValue("oooqsize", "300"),
					testAccCheckNstcpparamADCValue("wsval", "8"),
					testAccCheckNstcpparamADCValue("mptcpmaxpendingsf", "4"),
				),
			},
		},
	})
}

// testAccCheckNstcpparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. nstcpparam is a singleton, so it is fetched with an empty name.
func testAccCheckNstcpparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nstcpparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nstcpparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nstcpparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccNstcpparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNstcpparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nstcpparam.tf_nstcpparam_data", "delayedack", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_nstcpparam.tf_nstcpparam_data", "maxburst", "6"),
				),
			},
		},
	})
}
