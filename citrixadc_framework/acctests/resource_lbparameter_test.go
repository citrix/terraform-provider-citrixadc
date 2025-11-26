package acctests

import (
	"fmt"
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
