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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAppfwProfile_uber_payload = `

resource citrixadc_csvserver test_csvserver {
  ipv46 = "10.202.11.11"
  name = "test_csvserver"
  port = 8080
  servicetype = "HTTP"
}

resource "citrixadc_appfwsignatures" "test_appfwsignatures" {
  name       = "test_appfwsignatures"
  src        = "local://appfw_signatures.xml"
}

resource "citrixadc_appfwsignatures" "test_appfwsignatures_custom" {
  name       = "test_appfwsignatures"
  src        = "local://appfw_signatures.xml"
  ruleid = [ 999848, 999849, 999850, 999851, 999854, 999855 ]
  enabled = "ON"
  overwrite = true
  action = ["block"]
}

resource citrixadc_appfwprofile test_appfw_profile {
  name = "test_appfw_profile"
  signatures = citrixadc_appfwsignatures.test_appfwsignatures.name
}

resource citrixadc_appfwpolicy test_appfw_policy {
  name = "test_appfw_policy"
  profilename = citrixadc_appfwprofile.test_appfw_profile.name
  rule = "true"
  comment = "test comment"
}

resource citrixadc_csvserver_appfwpolicy_binding test_csvserver_appfwpolicy_binding {
  name = citrixadc_csvserver.test_csvserver.name
  priority = 100
  policyname  = citrixadc_appfwpolicy.test_appfw_policy.name
  gotopriorityexpression = "END"
}

resource "citrixadc_appfwpolicylabel" "test_appfw_policylabel" {
  labelname = "test_appfw_policylabel"
  policylabeltype = "http_req"
}


resource "citrixadc_appfwpolicylabel_appfwpolicy_binding" "test_appfw_policy_label_policy_binding" {
  labelname  = citrixadc_appfwpolicylabel.test_appfw_policylabel.labelname
  policyname = citrixadc_appfwpolicy.test_appfw_policy.name
  priority   = 20
}

resource citrixadc_appfwprofile_starturl_binding appfwprofile_starturl1 {
  name = citrixadc_appfwprofile.test_appfw_profile.name
  starturl = "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"
  alertonly      = "OFF"
  isautodeployed = "NOTAUTODEPLOYED"
  state          = "ENABLED"
}

resource citrixadc_appfwprofile_denyurl_binding appfwprofile_denyurl1 {
  name = citrixadc_appfwprofile.test_appfw_profile.name
  denyurl = "debug[.][^/?]*(|[?].*)$"
  alertonly      = "OFF"
  isautodeployed = "NOTAUTODEPLOYED"
  state          = "ENABLED"
}

resource "citrixadc_appfwprofile_contenttype_binding" "tf_binding" {
  name           = citrixadc_appfwprofile.test_appfw_profile.name
  contenttype    = "hello"
  state          = "ENABLED"
  alertonly      = "ON"
  isautodeployed = "NOTAUTODEPLOYED"
  comment        = "Testing"
}

resource citrixadc_appfwprofile_cookieconsistency_binding demo_binding {
  name              = citrixadc_appfwprofile.test_appfw_profile.name
  cookieconsistency = "^logon_[0-9A-Za-z]{2,15}$"
}

resource citrixadc_appfwprofile_crosssitescripting_binding demo_binding {
	name                 = citrixadc_appfwprofile.test_appfw_profile.name
	crosssitescripting   = "demoxss"
	formactionurl_xss    = "http://www.example.com"
	as_scan_location_xss = "HEADER"
	isregex_xss          = "NOTREGEX"
	comment              = "democomment"
	state                = "ENABLED"
	as_value_type_xss    = "Attribute"
	as_value_expr_xss    = "value"
	isvalueregex_xss     = "NOTREGEX"
}

resource "citrixadc_appfwprofile_sqlinjection_binding" "demo_binding" {
  name                 = citrixadc_appfwprofile.test_appfw_profile.name
  sqlinjection         = "demo_binding"
  as_scan_location_sql = "HEADER"
  formactionurl_sql    = "www.example.com"
  state                = "ENABLED"
  isregex_sql          = "NOTREGEX"
  as_value_type_sql    = "Keyword"
  as_value_expr_sql    = "example1"
  isvalueregex_sql     = "NOTREGEX"
}

resource "citrixadc_appfwprofile_fieldformat_binding" "tf_binding" {
  name                 = citrixadc_appfwprofile.test_appfw_profile.name
  fieldformat          = "tf_field"
  formactionurl_ff     = "http://www.example.com"
  comment              = "Testing"
  state                = "ENABLED"
  fieldformatmaxlength = 20
  isregexff            = "NOTREGEX"
  fieldtype            = "alpha"
  alertonly            = "OFF"
}

resource "citrixadc_appfwprofile_fieldconsistency_binding" "tf_binding" {
  name              = citrixadc_appfwprofile.test_appfw_profile.name
  fieldconsistency  = "tf_field"
  formactionurl_ffc = "www.example.com"
  isautodeployed    = "NOTAUTODEPLOYED"
  state             = "DISABLED"
  alertonly         = "OFF"
  isregex_ffc       = "REGEX"
  comment           = "Testing"
}

resource "citrixadc_appfwprofile_cmdinjection_binding" "tf_binding" {
  name                 = citrixadc_appfwprofile.test_appfw_profile.name
  cmdinjection         = "tf_cmdinjection"
  formactionurl_cmd    = "http://10.10.10.10/"
  as_scan_location_cmd = "HEADER"
  as_value_type_cmd    = "Keyword"
  as_value_expr_cmd    = "[a-z]+grep"
  alertonly            = "OFF"
  isvalueregex_cmd     = "REGEX"
  isautodeployed       = "NOTAUTODEPLOYED"
  comment              = "Testing"
}

resource "citrixadc_appfwprofile_csrftag_binding" "tf_binding" {
  name              = citrixadc_appfwprofile.test_appfw_profile.name
  csrftag           = "www.source.com"
  csrfformactionurl = "www.action.com"
  isautodeployed    = "NOTAUTODEPLOYED"
  comment           = "Testing"
  state             = "ENABLED"
  alertonly         = "ON"
}

resource "citrixadc_appfwprofile_fileuploadtype_binding" "tf_binding" {
  name                   = citrixadc_appfwprofile.test_appfw_profile.name
  fileuploadtype         = "tf_uploadtype"
  as_fileuploadtypes_url = "www.example.com"
  filetype               = ["pdf", "text"]
}

resource "citrixadc_appfwprofile_creditcardnumber_binding" "tf_binding" {
  name                = citrixadc_appfwprofile.test_appfw_profile.name
  creditcardnumber    = "123456789"
  creditcardnumberurl = "www.example.com"
  isautodeployed      = "AUTODEPLOYED"
  alertonly           = "ON"
  state               = "ENABLED"
  comment             = "Testing"
}

resource "citrixadc_appfwprofile_safeobject_binding" "tf_binding" {
  name           = citrixadc_appfwprofile.test_appfw_profile.name
  safeobject     = "tf_safeobject"
  as_expression  = "regularexpression"
  maxmatchlength = 10
  state          = "DISABLED"
  alertonly      = "OFF"
  isautodeployed = "AUTODEPLOYED"
  comment        = "Example"
  action         = ["block", "log"]
}

resource "citrixadc_appfwprofile_logexpression_binding" "tf_binding" {
  name             = citrixadc_appfwprofile.test_appfw_profile.name
  logexpression    = "tf_logexp"
  as_logexpression = "HTTP.REQ.IS_VALID"
  alertonly        = "ON"
  isautodeployed   = "AUTODEPLOYED"
  comment          = "Testing"
  state            = "ENABLED"
}

resource "citrixadc_appfwprofile_trustedlearningclients_binding" "tf_binding" {
  name                   = citrixadc_appfwprofile.test_appfw_profile.name
  trustedlearningclients = "1.2.31.1/32"
  state                  = "ENABLED"
  alertonly              = "ON"
  isautodeployed         = "AUTODEPLOYED"
  comment                = "Testing"
}

resource "citrixadc_appfwprofile_jsonxssurl_binding" "tf_binding" {
  name           = citrixadc_appfwprofile.test_appfw_profile.name
  jsonxssurl     = "www.example.com"
  alertonly      = "OFF"
  state          = "ENABLED"
  isautodeployed = "NOTAUTODEPLOYED"
  comment        = "Testing"
}

resource "citrixadc_appfwprofile_jsonsqlurl_binding" "tf_binding" {
  name           = citrixadc_appfwprofile.test_appfw_profile.name
  jsonsqlurl     = "[abc][a-z]a*"
  isautodeployed = "AUTODEPLOYED"
  state          = "ENABLED"
  alertonly      = "ON"
  comment        = "Testing"
}

resource "citrixadc_appfwprofile_jsoncmdurl_binding" "tf_binding" {
  name           = citrixadc_appfwprofile.test_appfw_profile.name
  jsoncmdurl     = "www.example.com"
  alertonly      = "ON"
  isautodeployed = "AUTODEPLOYED"
  comment        = "Testing"
  state          = "DISABLED"
}

resource "citrixadc_appfwprofile_jsondosurl_binding" "tf_binding" {
  name                        = citrixadc_appfwprofile.test_appfw_profile.name
  jsondosurl                  = ".*"
  state                       = "ENABLED"
  alertonly                   = "ON"
  isautodeployed              = "AUTODEPLOYED"
  jsonmaxarraylengthcheck     = "OFF"
  jsonmaxdocumentlengthcheck  = "ON"
  jsonmaxcontainerdepth       = 5
  jsonmaxobjectkeylengthcheck = "OFF"
  jsonmaxarraylength          = 100000
  jsonmaxdocumentlength       = 200000
  jsonmaxobjectkeycountcheck  = "ON"
  jsonmaxobjectkeylength      = 128
  jsonmaxobjectkeycount       = 1000
  jsonmaxstringlengthcheck    = "ON"
  jsonmaxcontainerdepthcheck  = "ON"
  jsonmaxstringlength         = 1000
  comment                     = "Testing"
}

resource "citrixadc_appfwprofile_xmlxss_binding" "tf_binding" {
  name                    = citrixadc_appfwprofile.test_appfw_profile.name
  xmlxss                  = "tf_xmlxss"
  as_scan_location_xmlxss = "ELEMENT"
  state                   = "ENABLED"
  alertonly               = "ON"
  isregex_xmlxss          = "NOTREGEX"
  isautodeployed          = "AUTODEPLOYED"
}

resource "citrixadc_appfwprofile_xmldosurl_binding" "tf_binding" {
  name                           = citrixadc_appfwprofile.test_appfw_profile.name
  xmldosurl                      = ".*"
  state                          = "ENABLED"
  xmlsoaparraycheck              = "ON"
  xmlmaxelementdepthcheck        = "ON"
  xmlmaxfilesize                 = 100000
  xmlmaxfilesizecheck            = "OFF"
  xmlmaxnamespaceurilength       = 200
  xmlmaxnamespaceurilengthcheck  = "ON"
  xmlmaxelementnamelength        = 300
  xmlmaxelementnamelengthcheck   = "ON"
  xmlmaxelements                 = 30
  xmlmaxelementscheck            = "ON"
  xmlmaxattributes               = 20
  xmlmaxattributescheck          = "ON"
  xmlmaxchardatalength           = 1000
  xmlmaxchardatalengthcheck      = "ON"
  xmlmaxnamespaces               = 30
  xmlmaxnamespacescheck          = "ON"
  xmlmaxattributenamelength      = 200
  xmlmaxattributenamelengthcheck = "ON"
}

resource "citrixadc_appfwprofile_xmlwsiurl_binding" "tf_binding" {
	name           = citrixadc_appfwprofile.test_appfw_profile.name
	xmlwsiurl      = ".*"
	state          = "DISABLED"
	xmlwsichecks   = "R1140"
	isautodeployed = "AUTODEPLOYED"
	comment        = "Testing"
	alertonly      = "ON"
}

resource "citrixadc_appfwprofile_xmlattachmenturl_binding" "tf_binding" {
  name                          = citrixadc_appfwprofile.test_appfw_profile.name
  xmlattachmenturl              = ".*"
  xmlattachmentcontenttype      = "abc*"
  alertonly                     = "ON"
  state                         = "ENABLED"
  isautodeployed                = "AUTODEPLOYED"
  comment                       = "Testing"
  xmlattachmentcontenttypecheck = "ON"
  xmlmaxattachmentsize          = "1000"
  xmlmaxattachmentsizecheck     = "ON"
}

`

const testAccAppfwprofile_add = `
	resource citrixadc_appfwprofile test_appfw {
		name = "test_appfw"
		bufferoverflowaction = ["none"]
		contenttypeaction = ["none"]
		cookieconsistencyaction = ["none"]
		creditcard = ["none"]
		creditcardaction = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction = ["none"]
		denyurlaction = ["none"]
		dynamiclearning = ["none"]
		fieldconsistencyaction = ["none"]
		fieldformataction = ["none"]
		fileuploadtypesaction = ["none"]
		inspectcontenttypes = ["none"]
		jsondosaction = ["none"]
		jsonsqlinjectionaction = ["none"]
		jsonxssaction = ["none"]
		multipleheaderaction = ["none"]
		sqlinjectionaction = ["none"]
		starturlaction = ["none"]
		type = ["HTML"]
		xmlattachmentaction = ["none"]
		xmldosaction = ["none"]
		xmlformataction = ["none"]
		xmlsoapfaultaction = ["none"]
		xmlsqlinjectionaction = ["none"]
		xmlvalidationaction = ["none"]
		xmlwsiaction = ["none"]
		xmlxssaction = ["none"]

	}
`

const testAccAppfwprofileDataSource_basic = `
	resource citrixadc_appfwprofile test_appfw {
		name = "test_appfw"
		bufferoverflowaction = ["none"]
		contenttypeaction = ["none"]
		cookieconsistencyaction = ["none"]
		creditcard = ["none"]
		creditcardaction = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction = ["none"]
		denyurlaction = ["none"]
		dynamiclearning = ["none"]
		fieldconsistencyaction = ["none"]
		fieldformataction = ["none"]
		fileuploadtypesaction = ["none"]
		inspectcontenttypes = ["none"]
		jsondosaction = ["none"]
		jsonsqlinjectionaction = ["none"]
		jsonxssaction = ["none"]
		multipleheaderaction = ["none"]
		sqlinjectionaction = ["none"]
		starturlaction = ["none"]
		type = ["HTML"]
		xmlattachmentaction = ["none"]
		xmldosaction = ["none"]
		xmlformataction = ["none"]
		xmlsoapfaultaction = ["none"]
		xmlsqlinjectionaction = ["none"]
		xmlvalidationaction = ["none"]
		xmlwsiaction = ["none"]
		xmlxssaction = ["none"]
	}

	data "citrixadc_appfwprofile" "test_appfw" {
		name = citrixadc_appfwprofile.test_appfw.name
		depends_on = [citrixadc_appfwprofile.test_appfw]
	}
`

const testAccAppfwprofile_update = `
	resource citrixadc_appfwprofile test_appfw {
		name = "test_appfw"
		responsecontenttype = "application/octet-stream"
		bufferoverflowaction = ["none"]
		contenttypeaction = ["none"]
		cookieconsistencyaction = ["none"]
		creditcard = ["none"]
		creditcardaction = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction = ["none"]
		denyurlaction = ["none"]
		dynamiclearning = ["none"]
		fieldconsistencyaction = ["none"]
		fieldformataction = ["none"]
		fileuploadtypesaction = ["none"]
		inspectcontenttypes = ["none"]
		jsondosaction = ["none"]
		jsonsqlinjectionaction = ["none"]
		jsonxssaction = ["none"]
		multipleheaderaction = ["none"]
		sqlinjectionaction = ["none"]
		starturlaction = ["none"]
		type = ["HTML"]
		xmlattachmentaction = ["none"]
		xmldosaction = ["none"]
		xmlformataction = ["none"]
		xmlsoapfaultaction = ["none"]
		xmlsqlinjectionaction = ["none"]
		xmlvalidationaction = ["none"]
		xmlwsiaction = ["none"]
		xmlxssaction = ["none"]

	}
`

func TestAccAppfwprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileExist("citrixadc_appfwprofile.test_appfw", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile.test_appfw", "name", "test_appfw"),
				),
			},
			{
				Config: testAccAppfwprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileExist("citrixadc_appfwprofile.test_appfw", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile.test_appfw", "name", "test_appfw"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile.test_appfw", "responsecontenttype", "application/octet-stream"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
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
		data, err := client.FindResource(service.Appfwprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAppfwprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile.test_appfw", "name", "test_appfw"),
				),
			},
		},
	})
}
