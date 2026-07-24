package citrixadc

import (
	"fmt"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLbparameter_basic_step1 = `

	resource "citrixadc_lbparameter" "tf_lbparameter" {
        httponlycookieflag = "DISABLED"
        useencryptedpersistencecookie = "ENABLED"
        consolidatedlconn = "NO"
        useportforhashlb = "NO"
        preferdirectroute = "NO"
        startuprrfactor = 10
        monitorskipmaxclient = "DISABLED"
        monitorconnectionclose = "FIN"
        vserverspecificmac = "DISABLED"
        allowboundsvcremoval = "ENABLED"
        retainservicestate = "OFF"
        dbsttl = 0
        maxpipelinenat = 240
        storemqttclientidandusername = "NO"
        dropmqttjumbomessage = "YES"
        lbhashalgorithm = "JARH"
        lbhashfingers = 512
		
	}

`

const testAccLbparameter_basic_step2 = `

	resource "citrixadc_lbparameter" "tf_lbparameter" {
        httponlycookieflag = "ENABLED"
        useencryptedpersistencecookie = "DISABLED"
        consolidatedlconn = "YES"
        useportforhashlb = "YES"
        preferdirectroute = "YES"
        startuprrfactor = 0
        monitorskipmaxclient = "DISABLED"
        monitorconnectionclose = "FIN"
        vserverspecificmac = "DISABLED"
        allowboundsvcremoval = "ENABLED"
        retainservicestate = "OFF"
        dbsttl = 0
        maxpipelinenat = 255
        storemqttclientidandusername = "NO"
        dropmqttjumbomessage = "YES"
        lbhashalgorithm = "DEFAULT"
        lbhashfingers = 256
		
	}
`

func TestAccLbparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbparameterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbparameter_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbparameterExist("citrixadc_lbparameter.tf_lbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "httponlycookieflag", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "useencryptedpersistencecookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "consolidatedlconn", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "useportforhashlb", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "preferdirectroute", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "startuprrfactor", "10"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "monitorskipmaxclient", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "monitorconnectionclose", "FIN"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "vserverspecificmac", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "allowboundsvcremoval", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "retainservicestate", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "dbsttl", "0"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "maxpipelinenat", "240"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "storemqttclientidandusername", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "dropmqttjumbomessage", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "lbhashalgorithm", "JARH"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "lbhashfingers", "512"),
				),
			},
			{
				Config: testAccLbparameter_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbparameterExist("citrixadc_lbparameter.tf_lbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "httponlycookieflag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "useencryptedpersistencecookie", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "consolidatedlconn", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "useportforhashlb", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "preferdirectroute", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "startuprrfactor", "0"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "monitorskipmaxclient", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "monitorconnectionclose", "FIN"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "vserverspecificmac", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "allowboundsvcremoval", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "retainservicestate", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "dbsttl", "0"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "maxpipelinenat", "255"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "storemqttclientidandusername", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "dropmqttjumbomessage", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "lbhashalgorithm", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "lbhashfingers", "256"),
				),
			},
		},
	})
}

func TestAccLbparameter_import(t *testing.T) {
	const resAddr = "citrixadc_lbparameter.tf_lbparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbparameterDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbparameter_basic_step1},
			{
				Config:            testAccLbparameter_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// cookiepassphrase_wo_version is a write-only version tracker
				// (default 1) that is populated from config, not read back from
				// the ADC, so it is absent after a bare import and cannot round-trip.
				ImportStateVerifyIgnore: []string{"cookiepassphrase_wo_version"},
			},
		},
	})
}

func testAccCheckLbparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbparameter name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Get a configured client from the test helper
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lbparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lbparameter %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbparameterDestroy(s *terraform.State) error {
	// Get a configured client from the test helper
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbparameter" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbparameter.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbparameter %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// Test backward-compatible path: using cookiepassphrase (Sensitive attribute)
const testAccLbparameter_cookiepassphrase_step1 = `

	variable "lbparameter_cookiepassphrase" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_lbparameter" "tf_lbparameter" {
		httponlycookieflag            = "ENABLED"
		useencryptedpersistencecookie = "ENABLED"
		cookiepassphrase              = var.lbparameter_cookiepassphrase
		usesecuredpersistencecookie   = "ENABLED"
	}
`

// Update backward-compatible path: change cookiepassphrase value
const testAccLbparameter_cookiepassphrase_step2 = `

	variable "lbparameter_cookiepassphrase_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_lbparameter" "tf_lbparameter" {
		httponlycookieflag            = "ENABLED"
		useencryptedpersistencecookie = "ENABLED"
		cookiepassphrase              = var.lbparameter_cookiepassphrase_2
		usesecuredpersistencecookie   = "ENABLED"
	}
`

func TestAccLbparameter_cookiepassphrase_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_lbparameter_cookiepassphrase", "oldpassphrase123")
	t.Setenv("TF_VAR_lbparameter_cookiepassphrase_2", "newpassphrase456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbparameterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbparameter_cookiepassphrase_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbparameterExist("citrixadc_lbparameter.tf_lbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "httponlycookieflag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "useencryptedpersistencecookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "usesecuredpersistencecookie", "ENABLED"),
				),
			},
			{
				Config: testAccLbparameter_cookiepassphrase_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbparameterExist("citrixadc_lbparameter.tf_lbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "httponlycookieflag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "useencryptedpersistencecookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "usesecuredpersistencecookie", "ENABLED"),
				),
			},
		},
	})
}

// Test ephemeral path: using cookiepassphrase_wo (WriteOnly attribute) with version tracker
const testAccLbparameter_cookiepassphrase_wo_step1 = `

	variable "lbparameter_cookiepassphrase_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_lbparameter" "tf_lbparameter" {
		httponlycookieflag            = "ENABLED"
		useencryptedpersistencecookie = "ENABLED"
		cookiepassphrase_wo           = var.lbparameter_cookiepassphrase_wo
		cookiepassphrase_wo_version   = 1
		usesecuredpersistencecookie   = "ENABLED"
	}
`

// Update ephemeral path: bump version to trigger update with new passphrase
const testAccLbparameter_cookiepassphrase_wo_step2 = `

	variable "lbparameter_cookiepassphrase_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_lbparameter" "tf_lbparameter" {
		httponlycookieflag            = "ENABLED"
		useencryptedpersistencecookie = "ENABLED"
		cookiepassphrase_wo           = var.lbparameter_cookiepassphrase_wo_2
		cookiepassphrase_wo_version   = 2
		usesecuredpersistencecookie   = "ENABLED"
	}
`

func TestAccLbparameter_cookiepassphrase_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_lbparameter_cookiepassphrase_wo", "ephemeral_pass1")
	t.Setenv("TF_VAR_lbparameter_cookiepassphrase_wo_2", "ephemeral_pass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbparameterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbparameter_cookiepassphrase_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbparameterExist("citrixadc_lbparameter.tf_lbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "httponlycookieflag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "useencryptedpersistencecookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "cookiepassphrase_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "usesecuredpersistencecookie", "ENABLED"),
				),
			},
			{
				Config: testAccLbparameter_cookiepassphrase_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbparameterExist("citrixadc_lbparameter.tf_lbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "httponlycookieflag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "useencryptedpersistencecookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "cookiepassphrase_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_lbparameter", "usesecuredpersistencecookie", "ENABLED"),
				),
			},
		},
	})
}

func TestAccLbparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbparameterDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccLbparameter_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbparameterExist("citrixadc_lbparameter.tf_lbparameter", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbparameter_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbparameterExist("citrixadc_lbparameter.tf_lbparameter", nil),
				),
			},
		},
	})
}

const testAccLbparameterDataSource_basic = `

	resource "citrixadc_lbparameter" "tf_lbparameter" {
        httponlycookieflag = "DISABLED"
        useencryptedpersistencecookie = "ENABLED"
        consolidatedlconn = "NO"
        useportforhashlb = "NO"
        preferdirectroute = "NO"
        startuprrfactor = 10
        monitorskipmaxclient = "DISABLED"
        monitorconnectionclose = "FIN"
        vserverspecificmac = "DISABLED"
        allowboundsvcremoval = "ENABLED"
        retainservicestate = "OFF"
        dbsttl = 0
        maxpipelinenat = 240
        storemqttclientidandusername = "NO"
        dropmqttjumbomessage = "YES"
        lbhashalgorithm = "JARH"
        lbhashfingers = 512
	}

	data "citrixadc_lbparameter" "tf_lbparameter" {
		depends_on = [citrixadc_lbparameter.tf_lbparameter]
	}
`

// lbparameter is a singleton, so the unset test has no mandatory identity
// fields. Step 1 sets every unset-eligible attribute to a valid non-default
// value; step 2 removes them all so the provider issues ?action=unset and each
// reverts to its NITRO default in state and on the appliance.
const testAccLbparameter_unset_step1 = `
	resource "citrixadc_lbparameter" "tf_unset" {
		allowboundsvcremoval          = "DISABLED"
		consolidatedlconn             = "NO"
		dbsttl                        = 3600
		dropmqttjumbomessage          = "NO"
		httponlycookieflag            = "DISABLED"
		lbhashalgorithm               = "JARH"
		lbhashfingers                 = 512
		monitorconnectionclose        = "RESET"
		monitorskipmaxclient          = "ENABLED"
		preferdirectroute             = "NO"
		proximityfromself             = "YES"
		retainservicestate            = "ON"
		storemqttclientidandusername  = "YES"
		undefaction                   = "RESET"
		useencryptedpersistencecookie = "ENABLED"
		useportforhashlb              = "NO"
		vserverspecificmac            = "ENABLED"
	}
`

const testAccLbparameter_unset_step2 = `
	resource "citrixadc_lbparameter" "tf_unset" {
		# All unset-eligible attributes removed from config -> provider must unset
		# them, reverting each to its NITRO default.
	}
`

func TestAccLbparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbparameterDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccLbparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbparameterExist("citrixadc_lbparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "allowboundsvcremoval", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "consolidatedlconn", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "dbsttl", "3600"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "dropmqttjumbomessage", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "httponlycookieflag", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "lbhashalgorithm", "JARH"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "lbhashfingers", "512"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "monitorconnectionclose", "RESET"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "monitorskipmaxclient", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "preferdirectroute", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "proximityfromself", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "retainservicestate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "storemqttclientidandusername", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "undefaction", "RESET"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "useencryptedpersistencecookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "useportforhashlb", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "vserverspecificmac", "ENABLED"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccLbparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbparameterExist("citrixadc_lbparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "allowboundsvcremoval", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "consolidatedlconn", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "dbsttl", "0"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "dropmqttjumbomessage", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "httponlycookieflag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "lbhashalgorithm", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "lbhashfingers", "256"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "monitorconnectionclose", "FIN"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "monitorskipmaxclient", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "preferdirectroute", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "proximityfromself", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "retainservicestate", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "storemqttclientidandusername", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "undefaction", "NOLBACTION"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "useencryptedpersistencecookie", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "useportforhashlb", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbparameter.tf_unset", "vserverspecificmac", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLbparameterADCValue("allowboundsvcremoval", "ENABLED"),
					testAccCheckLbparameterADCValue("consolidatedlconn", "YES"),
					testAccCheckLbparameterADCValue("dbsttl", "0"),
					testAccCheckLbparameterADCValue("dropmqttjumbomessage", "YES"),
					testAccCheckLbparameterADCValue("httponlycookieflag", "ENABLED"),
					testAccCheckLbparameterADCValue("lbhashalgorithm", "DEFAULT"),
					testAccCheckLbparameterADCValue("lbhashfingers", "256"),
					testAccCheckLbparameterADCValue("monitorconnectionclose", "FIN"),
					testAccCheckLbparameterADCValue("monitorskipmaxclient", "DISABLED"),
					testAccCheckLbparameterADCValue("preferdirectroute", "YES"),
					testAccCheckLbparameterADCValue("proximityfromself", "NO"),
					testAccCheckLbparameterADCValue("retainservicestate", "OFF"),
					testAccCheckLbparameterADCValue("storemqttclientidandusername", "NO"),
					testAccCheckLbparameterADCValue("undefaction", "NOLBACTION"),
					testAccCheckLbparameterADCValue("useencryptedpersistencecookie", "DISABLED"),
					testAccCheckLbparameterADCValue("useportforhashlb", "YES"),
					testAccCheckLbparameterADCValue("vserverspecificmac", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckLbparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. lbparameter is a singleton, so it is fetched with an empty resource name.
func testAccCheckLbparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lbparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lbparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lbparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccLbparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "httponlycookieflag", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "useencryptedpersistencecookie", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "consolidatedlconn", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "useportforhashlb", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "preferdirectroute", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "startuprrfactor", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "monitorskipmaxclient", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "monitorconnectionclose", "FIN"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "vserverspecificmac", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "allowboundsvcremoval", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "retainservicestate", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "dbsttl", "0"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "maxpipelinenat", "240"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "storemqttclientidandusername", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "dropmqttjumbomessage", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "lbhashalgorithm", "JARH"),
					resource.TestCheckResourceAttr("data.citrixadc_lbparameter.tf_lbparameter", "lbhashfingers", "512"),
				),
			},
		},
	})
}
