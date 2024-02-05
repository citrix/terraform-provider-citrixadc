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
	"testing"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccNsconfigSave_basic = `
	resource "citrixadc_nsconfig_save" "foo" {
		# all        = true # "all" attribute not present in 12.0
		timestamp  = "2020-03-24T12:37:06Z"
	}
`

func TestAccNsconfigSave_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNsconfigSave_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsconfigSaveExist("citrixadc_nsconfig_save.foo", nil),
				),
			},
		},
	})
}

const testAccNsconfigSave_save_race_no_retry = `
	resource "citrixadc_nsconfig_save" "foo" {
		# all        = true # "all" attribute not present in 12.0
		timestamp  = "2020-03-24T12:37:06Z"

		concurrent_save_ok = true
		concurrent_save_retries = 0
	}
`

const testAccNsconfigSave_save_race_retry = `
	resource "citrixadc_nsconfig_save" "foo" {
		# all        = true # "all" attribute not present in 12.0
		timestamp  = "2020-03-24T12:37:06Z"

		concurrent_save_ok = true
		concurrent_save_retries = 1
	}
`

func testAccCheckNsconfigSaveExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No NsConfigSave is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		return nil
	}
}

func TestAccNsconfigSave_save_race_no_retry(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccSaveRaceSetup(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNsconfigSave_save_race_no_retry,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsconfigSaveExist("citrixadc_nsconfig_save.foo", nil),
				),
			},
		},
	})
}

func TestAccNsconfigSave_save_race_retry(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccSaveRaceSetup(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNsconfigSave_save_race_retry,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsconfigSaveExist("citrixadc_nsconfig_save.foo", nil),
				),
			},
		},
	})
}

func testAccSaveRaceSetup(t *testing.T) {
	// Run Save config concurrently
	url := os.Getenv("NS_URL")
	password := os.Getenv("NS_PASSWORD")
	params := service.NitroParams{
		Url:       url,
		Username:  "nsroot",
		Password:  password,
		SslVerify: false,
	}
	client, err := service.NewNitroClientFromParams(params)
	if err != nil {
		t.Fatalf("extra client instantiation error: %s", err.Error())
	}

	lbvserverName := "tf_acc_saveconfig_race_lb"
	create_lb_vserver := func() {
		lbvserver := lb.Lbvserver{
			Name:        lbvserverName,
			Ipv46:       "10.3.4.25",
			Port:        80,
			Servicetype: "HTTP",
		}
		_, err := client.AddResource("lbvserver", lbvserverName, &lbvserver)
		if err != nil {
			t.Fatalf("Error creating race lb vserver: %s", err.Error())
		}
	}

	remove_lb_vserver := func() {
		err := client.DeleteResource("lbvserver", lbvserverName)
		if err != nil {
			t.Fatalf("Error deletin race lb vserver: %s", err.Error())
		}
	}

	do_save := func() {
		t.Log("start of save")
		nsconfig := ns.Nsconfig{}
		err := client.ActOnResource("nsconfig", &nsconfig, "save")
		if err != nil {
			t.Fatalf("do_save error: %s", err.Error())
		}
		t.Log("end of save")
	}
	// Do some changes to ensure save config is not a NOOP
	create_lb_vserver()
	remove_lb_vserver()

	// Concurrently run the first save config to ensure
	// save config from the configuration will raise the NITRO errorcode 293
	go do_save()

	// Give it 500ms to head start operation

	//lintignore:R018
	time.Sleep(500 * time.Millisecond)
	t.Log("end of setup")
}
