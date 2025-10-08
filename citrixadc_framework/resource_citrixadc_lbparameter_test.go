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
package citrixadc_framework

import (
	"fmt"
	"os"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"citrixadc": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add any setup code here
	if v := os.Getenv("NS_URL"); v == "" {
		t.Fatal("NS_URL must be set for acceptance tests")
	}
}

func testAccGetFrameworkClient() (*service.NitroClient, error) {
	username := os.Getenv("NS_LOGIN")
	if username == "" {
		username = "nsroot"
	}

	password := os.Getenv("NS_PASSWORD")
	if password == "" {
		password = "nsroot"
	}

	endpoint := os.Getenv("NS_URL")
	if endpoint == "" {
		return nil, fmt.Errorf("NS_URL environment variable must be set")
	}

	params := service.NitroParams{
		Url:      endpoint,
		Username: username,
		Password: password,
	}

	client, err := service.NewNitroClientFromParams(params)
	if err != nil {
		return nil, err
	}

	return client, nil
}

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
				),
			},
			{
				Config: testAccLbparameter_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbparameterExist("citrixadc_lbparameter.tf_lbparameter", nil),
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
