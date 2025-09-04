package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAppfwprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofileFunc,
		Read:          readAppfwprofileFunc,
		Update:        updateAppfwprofileFunc,
		Delete:        deleteAppfwprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"addcookieflags": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"archivename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bufferoverflowmaxcookielength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bufferoverflowmaxheaderlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bufferoverflowmaxurllength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"canonicalizehtmlresponse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"checkrequestheaders": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookieencryption": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookieproxying": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookietransforms": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"creditcardmaxallowed": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"creditcardxout": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crosssitescriptingcheckcompleteurls": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crosssitescriptingtransformunsafehtml": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customsettings": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultcharset": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultfieldformatmaxlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultfieldformatminlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultfieldformattype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaults": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dosecurecreditcardlogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enableformtagging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"errorurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"excludefileuploadfromchecks": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exemptclosureurlsfromsecuritychecks": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fileuploadmaxnum": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"htmlerrorobject": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalidpercenthandling": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jsonerrorobject": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jsonsqlinjectiontype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logeverypolicyhit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"optimizepartialreqs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"percentdecoderecursively": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"postbodylimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"postbodylimitsignature": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"refererheadercheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requestcontenttype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"responsecontenttype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rfcprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"semicolonfieldseparator": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionlessfieldconsistency": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionlessurlclosure": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signatures": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectionchecksqlwildchars": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectiononlycheckfieldswithsqlchars": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectionparsecomments": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectiontransformspecialchars": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectiontype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"starturlclosure": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"postbodylimitaction": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"bufferoverflowmaxquerylength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cookiehijackingaction": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"infercontenttypexmlpayloadaction": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cmdinjectionaction": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"streaming": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stripcomments": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"striphtmlcomments": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stripxmlcomments": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"urldecoderequestcookies": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usehtmlerrorobject": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verboseloglevel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xmlerrorobject": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xmlsqlinjectionchecksqlwildchars": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xmlsqlinjectiononlycheckfieldswithsqlchars": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xmlsqlinjectionparsecomments": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xmlsqlinjectiontype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bufferoverflowaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"contenttypeaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cookieconsistencyaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"creditcard": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"creditcardaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"crosssitescriptingaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"csrftagaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"denyurlaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dynamiclearning": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fieldconsistencyaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fieldformataction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fileuploadtypesaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"inspectcontenttypes": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"jsondosaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"jsonsqlinjectionaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"jsonxssaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"multipleheaderaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sqlinjectionaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"starturlaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"type": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlattachmentaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmldosaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlformataction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlsoapfaultaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlsqlinjectionaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlvalidationaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlwsiaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlxssaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"bufferoverflowmaxtotalheaderlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cmdinjectiontype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"htmlerrorstatuscode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"htmlerrorstatusmessage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectiongrammar": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apispec": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"as_prof_bypass_list_enable": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"as_prof_deny_list_enable": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"augment": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"blockkeywordaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ceflogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientipexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmdinjectiongrammar": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookiesamesiteattribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultfieldformatmaxoccurrences": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fakeaccountdetection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fieldscan": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fieldscanlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"geolocationlogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"grpcaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"importprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertcookiesamesiteattribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inspectquerycontenttypes": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"jsonblockkeywordaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"jsoncmdinjectionaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"jsoncmdinjectiongrammar": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jsoncmdinjectiontype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jsonerrorstatuscode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jsonerrorstatusmessage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jsonfieldscan": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jsonfieldscanlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jsonmessagescan": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jsonmessagescanlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jsonsqlinjectiongrammar": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"matchurlstring": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"messagescan": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"messagescanlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"messagescanlimitcontenttypes": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"overwrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"protofileobject": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"relaxationrules": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"replaceurlstring": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sessioncookiename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectionruletype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xmlerrorstatuscode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"xmlerrorstatusmessage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppfwprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwprofileName := d.Get("name").(string)

	appfwprofile := appfw.Appfwprofile{
		Name:                                       appfwprofileName,
		Postbodylimitaction:                        toStringList(d.Get("postbodylimitaction").([]interface{})),
		Bufferoverflowmaxquerylength:               d.Get("bufferoverflowmaxquerylength").(int),
		Cookiehijackingaction:                      toStringList(d.Get("cookiehijackingaction").([]interface{})),
		Infercontenttypexmlpayloadaction:           toStringList(d.Get("infercontenttypexmlpayloadaction").([]interface{})),
		Cmdinjectionaction:                         toStringList(d.Get("cmdinjectionaction").([]interface{})),
		Addcookieflags:                             d.Get("addcookieflags").(string),
		Archivename:                                d.Get("archivename").(string),
		Bufferoverflowmaxcookielength:              d.Get("bufferoverflowmaxcookielength").(int),
		Bufferoverflowmaxheaderlength:              d.Get("bufferoverflowmaxheaderlength").(int),
		Bufferoverflowmaxurllength:                 d.Get("bufferoverflowmaxurllength").(int),
		Canonicalizehtmlresponse:                   d.Get("canonicalizehtmlresponse").(string),
		Checkrequestheaders:                        d.Get("checkrequestheaders").(string),
		Comment:                                    d.Get("comment").(string),
		Cookieencryption:                           d.Get("cookieencryption").(string),
		Cookieproxying:                             d.Get("cookieproxying").(string),
		Cookietransforms:                           d.Get("cookietransforms").(string),
		Creditcardmaxallowed:                       d.Get("creditcardmaxallowed").(int),
		Creditcardxout:                             d.Get("creditcardxout").(string),
		Crosssitescriptingcheckcompleteurls:        d.Get("crosssitescriptingcheckcompleteurls").(string),
		Crosssitescriptingtransformunsafehtml:      d.Get("crosssitescriptingtransformunsafehtml").(string),
		Customsettings:                             d.Get("customsettings").(string),
		Defaultcharset:                             d.Get("defaultcharset").(string),
		Defaultfieldformatmaxlength:                d.Get("defaultfieldformatmaxlength").(int),
		Defaultfieldformatminlength:                d.Get("defaultfieldformatminlength").(int),
		Defaultfieldformattype:                     d.Get("defaultfieldformattype").(string),
		Defaults:                                   d.Get("defaults").(string),
		Dosecurecreditcardlogging:                  d.Get("dosecurecreditcardlogging").(string),
		Enableformtagging:                          d.Get("enableformtagging").(string),
		Errorurl:                                   d.Get("errorurl").(string),
		Excludefileuploadfromchecks:                d.Get("excludefileuploadfromchecks").(string),
		Exemptclosureurlsfromsecuritychecks:        d.Get("exemptclosureurlsfromsecuritychecks").(string),
		Fileuploadmaxnum:                           d.Get("fileuploadmaxnum").(int),
		Htmlerrorobject:                            d.Get("htmlerrorobject").(string),
		Invalidpercenthandling:                     d.Get("invalidpercenthandling").(string),
		Jsonerrorobject:                            d.Get("jsonerrorobject").(string),
		Jsonsqlinjectiontype:                       d.Get("jsonsqlinjectiontype").(string),
		Logeverypolicyhit:                          d.Get("logeverypolicyhit").(string),
		Optimizepartialreqs:                        d.Get("optimizepartialreqs").(string),
		Percentdecoderecursively:                   d.Get("percentdecoderecursively").(string),
		Postbodylimit:                              d.Get("postbodylimit").(int),
		Postbodylimitsignature:                     d.Get("postbodylimitsignature").(int),
		Refererheadercheck:                         d.Get("refererheadercheck").(string),
		Requestcontenttype:                         d.Get("requestcontenttype").(string),
		Responsecontenttype:                        d.Get("responsecontenttype").(string),
		Rfcprofile:                                 d.Get("rfcprofile").(string),
		Semicolonfieldseparator:                    d.Get("semicolonfieldseparator").(string),
		Sessionlessfieldconsistency:                d.Get("sessionlessfieldconsistency").(string),
		Sessionlessurlclosure:                      d.Get("sessionlessurlclosure").(string),
		Signatures:                                 d.Get("signatures").(string),
		Sqlinjectionchecksqlwildchars:              d.Get("sqlinjectionchecksqlwildchars").(string),
		Sqlinjectiononlycheckfieldswithsqlchars:    d.Get("sqlinjectiononlycheckfieldswithsqlchars").(string),
		Sqlinjectionparsecomments:                  d.Get("sqlinjectionparsecomments").(string),
		Sqlinjectiontransformspecialchars:          d.Get("sqlinjectiontransformspecialchars").(string),
		Sqlinjectiontype:                           d.Get("sqlinjectiontype").(string),
		Starturlclosure:                            d.Get("starturlclosure").(string),
		Streaming:                                  d.Get("streaming").(string),
		Stripcomments:                              d.Get("stripcomments").(string),
		Striphtmlcomments:                          d.Get("striphtmlcomments").(string),
		Stripxmlcomments:                           d.Get("stripxmlcomments").(string),
		Trace:                                      d.Get("trace").(string),
		Urldecoderequestcookies:                    d.Get("urldecoderequestcookies").(string),
		Usehtmlerrorobject:                         d.Get("usehtmlerrorobject").(string),
		Verboseloglevel:                            d.Get("verboseloglevel").(string),
		Xmlerrorobject:                             d.Get("xmlerrorobject").(string),
		Xmlsqlinjectionchecksqlwildchars:           d.Get("xmlsqlinjectionchecksqlwildchars").(string),
		Xmlsqlinjectiononlycheckfieldswithsqlchars: d.Get("xmlsqlinjectiononlycheckfieldswithsqlchars").(string),
		Xmlsqlinjectionparsecomments:               d.Get("xmlsqlinjectionparsecomments").(string),
		Xmlsqlinjectiontype:                        d.Get("xmlsqlinjectiontype").(string),
		Bufferoverflowmaxtotalheaderlength:         d.Get("bufferoverflowmaxtotalheaderlength").(int),
		Cmdinjectiontype:                           d.Get("cmdinjectiontype").(string),
		Htmlerrorstatuscode:                        d.Get("htmlerrorstatuscode").(int),
		Htmlerrorstatusmessage:                     d.Get("htmlerrorstatusmessage").(string),
		Sqlinjectiongrammar:                        d.Get("sqlinjectiongrammar").(string),
		Apispec:                                    d.Get("apispec").(string),
		Asprofbypasslistenable:                     d.Get("as_prof_bypass_list_enable").(string),
		Asprofdenylistenable:                       d.Get("as_prof_deny_list_enable").(string),
		Augment:                                    d.Get("augment").(bool),
		Ceflogging:                                 d.Get("ceflogging").(string),
		Clientipexpression:                         d.Get("clientipexpression").(string),
		Cmdinjectiongrammar:                        d.Get("cmdinjectiongrammar").(string),
		Cookiesamesiteattribute:                    d.Get("cookiesamesiteattribute").(string),
		Defaultfieldformatmaxoccurrences:           d.Get("defaultfieldformatmaxoccurrences").(int),
		Fakeaccountdetection:                       d.Get("fakeaccountdetection").(string),
		Fieldscan:                                  d.Get("fieldscan").(string),
		Fieldscanlimit:                             d.Get("fieldscanlimit").(int),
		Geolocationlogging:                         d.Get("geolocationlogging").(string),
		Importprofilename:                          d.Get("importprofilename").(string),
		Insertcookiesamesiteattribute:              d.Get("insertcookiesamesiteattribute").(string),
		Jsoncmdinjectiongrammar:                    d.Get("jsoncmdinjectiongrammar").(string),
		Jsoncmdinjectiontype:                       d.Get("jsoncmdinjectiontype").(string),
		Jsonerrorstatuscode:                        d.Get("jsonerrorstatuscode").(int),
		Jsonerrorstatusmessage:                     d.Get("jsonerrorstatusmessage").(string),
		Jsonfieldscan:                              d.Get("jsonfieldscan").(string),
		Jsonfieldscanlimit:                         d.Get("jsonfieldscanlimit").(int),
		Jsonmessagescan:                            d.Get("jsonmessagescan").(string),
		Jsonmessagescanlimit:                       d.Get("jsonmessagescanlimit").(int),
		Jsonsqlinjectiongrammar:                    d.Get("jsonsqlinjectiongrammar").(string),
		Matchurlstring:                             d.Get("matchurlstring").(string),
		Messagescan:                                d.Get("messagescan").(string),
		Messagescanlimit:                           d.Get("messagescanlimit").(int),
		Overwrite:                                  d.Get("overwrite").(bool),
		Protofileobject:                            d.Get("protofileobject").(string),
		Relaxationrules:                            d.Get("relaxationrules").(bool),
		Replaceurlstring:                           d.Get("replaceurlstring").(string),
		Sessioncookiename:                          d.Get("sessioncookiename").(string),
		Sqlinjectionruletype:                       d.Get("sqlinjectionruletype").(string),
		Xmlerrorstatuscode:                         d.Get("xmlerrorstatuscode").(int),
		Xmlerrorstatusmessage:                      d.Get("xmlerrorstatusmessage").(string),
	}

	appfwprofile.Bufferoverflowaction = toStringList(d.Get("bufferoverflowaction").(*schema.Set).List())
	appfwprofile.Contenttypeaction = toStringList(d.Get("contenttypeaction").(*schema.Set).List())
	appfwprofile.Cookieconsistencyaction = toStringList(d.Get("cookieconsistencyaction").(*schema.Set).List())
	appfwprofile.Creditcard = toStringList(d.Get("creditcard").(*schema.Set).List())
	appfwprofile.Creditcardaction = toStringList(d.Get("creditcardaction").(*schema.Set).List())
	appfwprofile.Crosssitescriptingaction = toStringList(d.Get("crosssitescriptingaction").(*schema.Set).List())
	appfwprofile.Csrftagaction = toStringList(d.Get("csrftagaction").(*schema.Set).List())
	appfwprofile.Denyurlaction = toStringList(d.Get("denyurlaction").(*schema.Set).List())
	appfwprofile.Dynamiclearning = toStringList(d.Get("dynamiclearning").(*schema.Set).List())
	appfwprofile.Fieldconsistencyaction = toStringList(d.Get("fieldconsistencyaction").(*schema.Set).List())
	appfwprofile.Fieldformataction = toStringList(d.Get("fieldformataction").(*schema.Set).List())
	appfwprofile.Fileuploadtypesaction = toStringList(d.Get("fileuploadtypesaction").(*schema.Set).List())
	appfwprofile.Inspectcontenttypes = toStringList(d.Get("inspectcontenttypes").(*schema.Set).List())
	appfwprofile.Jsondosaction = toStringList(d.Get("jsondosaction").(*schema.Set).List())
	appfwprofile.Jsonsqlinjectionaction = toStringList(d.Get("jsonsqlinjectionaction").(*schema.Set).List())
	appfwprofile.Jsonxssaction = toStringList(d.Get("jsonxssaction").(*schema.Set).List())
	appfwprofile.Multipleheaderaction = toStringList(d.Get("multipleheaderaction").(*schema.Set).List())
	appfwprofile.Sqlinjectionaction = toStringList(d.Get("sqlinjectionaction").(*schema.Set).List())
	appfwprofile.Starturlaction = toStringList(d.Get("starturlaction").(*schema.Set).List())
	appfwprofile.Type = toStringList(d.Get("type").(*schema.Set).List())
	appfwprofile.Xmlattachmentaction = toStringList(d.Get("xmlattachmentaction").(*schema.Set).List())
	appfwprofile.Xmldosaction = toStringList(d.Get("xmldosaction").(*schema.Set).List())
	appfwprofile.Xmlformataction = toStringList(d.Get("xmlformataction").(*schema.Set).List())
	appfwprofile.Xmlsoapfaultaction = toStringList(d.Get("xmlsoapfaultaction").(*schema.Set).List())
	appfwprofile.Xmlsqlinjectionaction = toStringList(d.Get("xmlsqlinjectionaction").(*schema.Set).List())
	appfwprofile.Xmlvalidationaction = toStringList(d.Get("xmlvalidationaction").(*schema.Set).List())
	appfwprofile.Xmlwsiaction = toStringList(d.Get("xmlwsiaction").(*schema.Set).List())
	appfwprofile.Xmlxssaction = toStringList(d.Get("xmlxssaction").(*schema.Set).List())
	appfwprofile.Blockkeywordaction = toStringList(d.Get("blockkeywordaction").(*schema.Set).List())
	appfwprofile.Grpcaction = toStringList(d.Get("grpcaction").(*schema.Set).List())
	appfwprofile.Inspectquerycontenttypes = toStringList(d.Get("inspectquerycontenttypes").(*schema.Set).List())
	appfwprofile.Jsonblockkeywordaction = toStringList(d.Get("jsonblockkeywordaction").(*schema.Set).List())
	appfwprofile.Jsoncmdinjectionaction = toStringList(d.Get("jsoncmdinjectionaction").(*schema.Set).List())
	appfwprofile.Messagescanlimitcontenttypes = toStringList(d.Get("messagescanlimitcontenttypes").(*schema.Set).List())
	appfwprofile.Restaction = toStringList(d.Get("restaction").(*schema.Set).List())

	_, err := client.AddResource(service.Appfwprofile.Type(), appfwprofileName, &appfwprofile)
	if err != nil {
		return err
	}

	d.SetId(appfwprofileName)

	err = readAppfwprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile but we can't read it ?? %s", appfwprofileName)
		return nil
	}
	return nil
}

func readAppfwprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile state %s", appfwprofileName)
	data, err := client.FindResource(service.Appfwprofile.Type(), appfwprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile state %s", appfwprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("postbodylimitaction", data["postbodylimitaction"])
	setToInt("bufferoverflowmaxquerylength", d, data["bufferoverflowmaxquerylength"])
	d.Set("cookiehijackingaction", data["cookiehijackingaction"])
	d.Set("infercontenttypexmlpayloadaction", data["infercontenttypexmlpayloadaction"])
	d.Set("cmdinjectionaction", data["cmdinjectionaction"])
	d.Set("addcookieflags", data["addcookieflags"])
	d.Set("archivename", data["archivename"])
	d.Set("bufferoverflowmaxcookielength", data["bufferoverflowmaxcookielength"])
	d.Set("bufferoverflowmaxheaderlength", data["bufferoverflowmaxheaderlength"])
	d.Set("bufferoverflowmaxurllength", data["bufferoverflowmaxurllength"])
	d.Set("canonicalizehtmlresponse", data["canonicalizehtmlresponse"])
	d.Set("checkrequestheaders", data["checkrequestheaders"])
	d.Set("comment", data["comment"])
	d.Set("cookieencryption", data["cookieencryption"])
	d.Set("cookieproxying", data["cookieproxying"])
	d.Set("cookietransforms", data["cookietransforms"])
	d.Set("creditcardmaxallowed", data["creditcardmaxallowed"])
	d.Set("creditcardxout", data["creditcardxout"])
	d.Set("crosssitescriptingcheckcompleteurls", data["crosssitescriptingcheckcompleteurls"])
	d.Set("crosssitescriptingtransformunsafehtml", data["crosssitescriptingtransformunsafehtml"])
	d.Set("customsettings", data["customsettings"])
	d.Set("defaultcharset", data["defaultcharset"])
	d.Set("defaultfieldformatmaxlength", data["defaultfieldformatmaxlength"])
	d.Set("defaultfieldformatminlength", data["defaultfieldformatminlength"])
	d.Set("defaultfieldformattype", data["defaultfieldformattype"])
	d.Set("defaults", data["defaults"])
	d.Set("dosecurecreditcardlogging", data["dosecurecreditcardlogging"])
	d.Set("enableformtagging", data["enableformtagging"])
	d.Set("errorurl", data["errorurl"])
	d.Set("excludefileuploadfromchecks", data["excludefileuploadfromchecks"])
	d.Set("exemptclosureurlsfromsecuritychecks", data["exemptclosureurlsfromsecuritychecks"])
	d.Set("fileuploadmaxnum", data["fileuploadmaxnum"])
	d.Set("htmlerrorobject", data["htmlerrorobject"])
	d.Set("invalidpercenthandling", data["invalidpercenthandling"])
	d.Set("jsonerrorobject", data["jsonerrorobject"])
	d.Set("jsonsqlinjectiontype", data["jsonsqlinjectiontype"])
	d.Set("logeverypolicyhit", data["logeverypolicyhit"])
	d.Set("name", data["name"])
	d.Set("optimizepartialreqs", data["optimizepartialreqs"])
	d.Set("percentdecoderecursively", data["percentdecoderecursively"])
	d.Set("postbodylimit", data["postbodylimit"])
	d.Set("postbodylimitsignature", data["postbodylimitsignature"])
	d.Set("refererheadercheck", data["refererheadercheck"])
	d.Set("requestcontenttype", data["requestcontenttype"])
	d.Set("responsecontenttype", data["responsecontenttype"])
	d.Set("rfcprofile", data["rfcprofile"])
	d.Set("semicolonfieldseparator", data["semicolonfieldseparator"])
	d.Set("sessionlessfieldconsistency", data["sessionlessfieldconsistency"])
	d.Set("sessionlessurlclosure", data["sessionlessurlclosure"])
	d.Set("signatures", data["signatures"])
	d.Set("sqlinjectionchecksqlwildchars", data["sqlinjectionchecksqlwildchars"])
	// d.Set("sqlinjectiononlycheckfieldswithsqlchars", data["sqlinjectiononlycheckfieldswithsqlchars"])
	d.Set("sqlinjectionparsecomments", data["sqlinjectionparsecomments"])
	d.Set("sqlinjectiontransformspecialchars", data["sqlinjectiontransformspecialchars"])
	d.Set("sqlinjectiontype", data["sqlinjectiontype"])
	d.Set("starturlclosure", data["starturlclosure"])
	d.Set("streaming", data["streaming"])
	// d.Set("stripcomments", data["stripcomments"])
	d.Set("striphtmlcomments", data["striphtmlcomments"])
	d.Set("stripxmlcomments", data["stripxmlcomments"])
	d.Set("trace", data["trace"])
	d.Set("urldecoderequestcookies", data["urldecoderequestcookies"])
	d.Set("usehtmlerrorobject", data["usehtmlerrorobject"])
	d.Set("verboseloglevel", data["verboseloglevel"])
	d.Set("xmlerrorobject", data["xmlerrorobject"])
	d.Set("xmlsqlinjectionchecksqlwildchars", data["xmlsqlinjectionchecksqlwildchars"])
	// d.Set("xmlsqlinjectiononlycheckfieldswithsqlchars", data["xmlsqlinjectiononlycheckfieldswithsqlchars"])
	d.Set("xmlsqlinjectionparsecomments", data["xmlsqlinjectionparsecomments"])
	d.Set("xmlsqlinjectiontype", data["xmlsqlinjectiontype"])

	d.Set("bufferoverflowaction", data["bufferoverflowaction"])
	d.Set("contenttypeaction", data["contenttypeaction"])
	d.Set("cookieconsistencyaction", data["cookieconsistencyaction"])
	d.Set("creditcard", data["creditcard"])
	d.Set("creditcardaction", data["creditcardaction"])
	d.Set("crosssitescriptingaction", data["crosssitescriptingaction"])
	d.Set("csrftagaction", data["csrftagaction"])
	d.Set("denyurlaction", data["denyurlaction"])
	d.Set("dynamiclearning", data["dynamiclearning"])
	d.Set("fieldconsistencyaction", data["fieldconsistencyaction"])
	d.Set("fieldformataction", data["fieldformataction"])
	d.Set("fileuploadtypesaction", data["fileuploadtypesaction"])
	d.Set("inspectcontenttypes", data["inspectcontenttypes"])
	d.Set("jsondosaction", data["jsondosaction"])
	d.Set("jsonsqlinjectionaction", data["jsonsqlinjectionaction"])
	d.Set("jsonxssaction", data["jsonxssaction"])
	d.Set("multipleheaderaction", data["multipleheaderaction"])
	d.Set("sqlinjectionaction", data["sqlinjectionaction"])
	d.Set("starturlaction", data["starturlaction"])
	d.Set("type", data["type"])
	d.Set("xmlattachmentaction", data["xmlattachmentaction"])
	d.Set("xmldosaction", data["xmldosaction"])
	d.Set("xmlformataction", data["xmlformataction"])
	d.Set("xmlsoapfaultaction", data["xmlsoapfaultaction"])
	d.Set("xmlsqlinjectionaction", data["xmlsqlinjectionaction"])
	d.Set("xmlvalidationaction", data["xmlvalidationaction"])
	d.Set("xmlwsiaction", data["xmlwsiaction"])
	d.Set("xmlxssaction", data["xmlxssaction"])
	setToInt("bufferoverflowmaxtotalheaderlength", d, data["bufferoverflowmaxtotalheaderlength"])
	d.Set("cmdinjectiontype", data["cmdinjectiontype"])
	d.Set("htmlerrorstatusmessage", data["htmlerrorstatusmessage"])
	setToInt("htmlerrorstatuscode", d, data["htmlerrorstatuscode"])
	d.Set("sqlinjectiongrammar", data["sqlinjectiongrammar"])

	// Add new attributes
	d.Set("apispec", data["apispec"])
	d.Set("as_prof_bypass_list_enable", data["asprofbypasslistenable"])
	d.Set("as_prof_deny_list_enable", data["asprofdenylistenable"])
	d.Set("augment", data["augment"])
	d.Set("ceflogging", data["ceflogging"])
	d.Set("clientipexpression", data["clientipexpression"])
	d.Set("cmdinjectiongrammar", data["cmdinjectiongrammar"])
	d.Set("cookiesamesiteattribute", data["cookiesamesiteattribute"])
	d.Set("defaultfieldformatmaxoccurrences", data["defaultfieldformatmaxoccurrences"])
	d.Set("fakeaccountdetection", data["fakeaccountdetection"])
	d.Set("fieldscan", data["fieldscan"])
	d.Set("fieldscanlimit", data["fieldscanlimit"])
	d.Set("geolocationlogging", data["geolocationlogging"])
	d.Set("importprofilename", data["importprofilename"])
	d.Set("insertcookiesamesiteattribute", data["insertcookiesamesiteattribute"])
	d.Set("jsoncmdinjectiongrammar", data["jsoncmdinjectiongrammar"])
	d.Set("jsoncmdinjectiontype", data["jsoncmdinjectiontype"])
	d.Set("jsonerrorstatuscode", data["jsonerrorstatuscode"])
	d.Set("jsonerrorstatusmessage", data["jsonerrorstatusmessage"])
	d.Set("jsonfieldscan", data["jsonfieldscan"])
	d.Set("jsonfieldscanlimit", data["jsonfieldscanlimit"])
	d.Set("jsonmessagescan", data["jsonmessagescan"])
	d.Set("jsonmessagescanlimit", data["jsonmessagescanlimit"])
	d.Set("jsonsqlinjectiongrammar", data["jsonsqlinjectiongrammar"])
	d.Set("matchurlstring", data["matchurlstring"])
	d.Set("messagescan", data["messagescan"])
	d.Set("messagescanlimit", data["messagescanlimit"])
	d.Set("overwrite", data["overwrite"])
	d.Set("protofileobject", data["protofileobject"])
	d.Set("relaxationrules", data["relaxationrules"])
	d.Set("replaceurlstring", data["replaceurlstring"])
	d.Set("sessioncookiename", data["sessioncookiename"])
	d.Set("sqlinjectionruletype", data["sqlinjectionruletype"])
	d.Set("xmlerrorstatuscode", data["xmlerrorstatuscode"])
	d.Set("xmlerrorstatusmessage", data["xmlerrorstatusmessage"])

	// Add new array/set type attributes
	d.Set("blockkeywordaction", data["blockkeywordaction"])
	d.Set("grpcaction", data["grpcaction"])
	d.Set("inspectquerycontenttypes", data["inspectquerycontenttypes"])
	d.Set("jsonblockkeywordaction", data["jsonblockkeywordaction"])
	d.Set("jsoncmdinjectionaction", data["jsoncmdinjectionaction"])
	d.Set("messagescanlimitcontenttypes", data["messagescanlimitcontenttypes"])
	d.Set("restaction", data["restaction"])

	return nil

}

func updateAppfwprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppfwprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwprofileName := d.Get("name").(string)

	appfwprofile := appfw.Appfwprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("postbodylimitaction") {
		log.Printf("[DEBUG]  citrixadc-provider: postbodylimitaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Postbodylimitaction = toStringList(d.Get("postbodylimitaction").([]interface{}))
		hasChange = true
	}
	if d.HasChange("bufferoverflowmaxquerylength") {
		log.Printf("[DEBUG]  citrixadc-provider: bufferoverflowmaxquerylength has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Bufferoverflowmaxquerylength = d.Get("bufferoverflowmaxquerylength").(int)
		hasChange = true
	}
	if d.HasChange("cookiehijackingaction") {
		log.Printf("[DEBUG]  citrixadc-provider: cookiehijackingaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Cookiehijackingaction = toStringList(d.Get("cookiehijackingaction").([]interface{}))
		hasChange = true
	}
	if d.HasChange("infercontenttypexmlpayloadaction") {
		log.Printf("[DEBUG]  citrixadc-provider: infercontenttypexmlpayloadaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Infercontenttypexmlpayloadaction = toStringList(d.Get("infercontenttypexmlpayloadaction").([]interface{}))
		hasChange = true
	}
	if d.HasChange("cmdinjectionaction") {
		log.Printf("[DEBUG]  citrixadc-provider: cmdinjectionaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Cmdinjectionaction = toStringList(d.Get("cmdinjectionaction").([]interface{}))
		hasChange = true
	}
	if d.HasChange("addcookieflags") {
		log.Printf("[DEBUG]  citrixadc-provider: Addcookieflags has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Addcookieflags = d.Get("addcookieflags").(string)
		hasChange = true
	}
	if d.HasChange("archivename") {
		log.Printf("[DEBUG]  citrixadc-provider: Archivename has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Archivename = d.Get("archivename").(string)
		hasChange = true
	}
	if d.HasChange("bufferoverflowmaxcookielength") {
		log.Printf("[DEBUG]  citrixadc-provider: Bufferoverflowmaxcookielength has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Bufferoverflowmaxcookielength = d.Get("bufferoverflowmaxcookielength").(int)
		hasChange = true
	}
	if d.HasChange("bufferoverflowmaxheaderlength") {
		log.Printf("[DEBUG]  citrixadc-provider: Bufferoverflowmaxheaderlength has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Bufferoverflowmaxheaderlength = d.Get("bufferoverflowmaxheaderlength").(int)
		hasChange = true
	}
	if d.HasChange("bufferoverflowmaxurllength") {
		log.Printf("[DEBUG]  citrixadc-provider: Bufferoverflowmaxurllength has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Bufferoverflowmaxurllength = d.Get("bufferoverflowmaxurllength").(int)
		hasChange = true
	}
	if d.HasChange("canonicalizehtmlresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Canonicalizehtmlresponse has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Canonicalizehtmlresponse = d.Get("canonicalizehtmlresponse").(string)
		hasChange = true
	}
	if d.HasChange("checkrequestheaders") {
		log.Printf("[DEBUG]  citrixadc-provider: Checkrequestheaders has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Checkrequestheaders = d.Get("checkrequestheaders").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("cookieencryption") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookieencryption has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Cookieencryption = d.Get("cookieencryption").(string)
		hasChange = true
	}
	if d.HasChange("cookieproxying") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookieproxying has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Cookieproxying = d.Get("cookieproxying").(string)
		hasChange = true
	}
	if d.HasChange("cookietransforms") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookietransforms has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Cookietransforms = d.Get("cookietransforms").(string)
		hasChange = true
	}
	if d.HasChange("creditcardmaxallowed") {
		log.Printf("[DEBUG]  citrixadc-provider: Creditcardmaxallowed has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Creditcardmaxallowed = d.Get("creditcardmaxallowed").(int)
		hasChange = true
	}
	if d.HasChange("creditcardxout") {
		log.Printf("[DEBUG]  citrixadc-provider: Creditcardxout has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Creditcardxout = d.Get("creditcardxout").(string)
		hasChange = true
	}
	if d.HasChange("crosssitescriptingcheckcompleteurls") {
		log.Printf("[DEBUG]  citrixadc-provider: Crosssitescriptingcheckcompleteurls has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Crosssitescriptingcheckcompleteurls = d.Get("crosssitescriptingcheckcompleteurls").(string)
		hasChange = true
	}
	if d.HasChange("crosssitescriptingtransformunsafehtml") {
		log.Printf("[DEBUG]  citrixadc-provider: Crosssitescriptingtransformunsafehtml has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Crosssitescriptingtransformunsafehtml = d.Get("crosssitescriptingtransformunsafehtml").(string)
		hasChange = true
	}
	if d.HasChange("customsettings") {
		log.Printf("[DEBUG]  citrixadc-provider: Customsettings has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Customsettings = d.Get("customsettings").(string)
		hasChange = true
	}
	if d.HasChange("defaultcharset") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultcharset has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Defaultcharset = d.Get("defaultcharset").(string)
		hasChange = true
	}
	if d.HasChange("defaultfieldformatmaxlength") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultfieldformatmaxlength has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Defaultfieldformatmaxlength = d.Get("defaultfieldformatmaxlength").(int)
		hasChange = true
	}
	if d.HasChange("defaultfieldformatminlength") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultfieldformatminlength has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Defaultfieldformatminlength = d.Get("defaultfieldformatminlength").(int)
		hasChange = true
	}
	if d.HasChange("defaultfieldformattype") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultfieldformattype has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Defaultfieldformattype = d.Get("defaultfieldformattype").(string)
		hasChange = true
	}
	if d.HasChange("defaults") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaults has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Defaults = d.Get("defaults").(string)
		hasChange = true
	}
	if d.HasChange("dosecurecreditcardlogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Dosecurecreditcardlogging has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Dosecurecreditcardlogging = d.Get("dosecurecreditcardlogging").(string)
		hasChange = true
	}
	if d.HasChange("enableformtagging") {
		log.Printf("[DEBUG]  citrixadc-provider: Enableformtagging has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Enableformtagging = d.Get("enableformtagging").(string)
		hasChange = true
	}
	if d.HasChange("errorurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Errorurl has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Errorurl = d.Get("errorurl").(string)
		hasChange = true
	}
	if d.HasChange("excludefileuploadfromchecks") {
		log.Printf("[DEBUG]  citrixadc-provider: Excludefileuploadfromchecks has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Excludefileuploadfromchecks = d.Get("excludefileuploadfromchecks").(string)
		hasChange = true
	}
	if d.HasChange("exemptclosureurlsfromsecuritychecks") {
		log.Printf("[DEBUG]  citrixadc-provider: Exemptclosureurlsfromsecuritychecks has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Exemptclosureurlsfromsecuritychecks = d.Get("exemptclosureurlsfromsecuritychecks").(string)
		hasChange = true
	}
	if d.HasChange("fileuploadmaxnum") {
		log.Printf("[DEBUG]  citrixadc-provider: Fileuploadmaxnum has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Fileuploadmaxnum = d.Get("fileuploadmaxnum").(int)
		hasChange = true
	}
	if d.HasChange("htmlerrorobject") {
		log.Printf("[DEBUG]  citrixadc-provider: Htmlerrorobject has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Htmlerrorobject = d.Get("htmlerrorobject").(string)
		hasChange = true
	}
	if d.HasChange("invalidpercenthandling") {
		log.Printf("[DEBUG]  citrixadc-provider: Invalidpercenthandling has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Invalidpercenthandling = d.Get("invalidpercenthandling").(string)
		hasChange = true
	}
	if d.HasChange("jsonerrorobject") {
		log.Printf("[DEBUG]  citrixadc-provider: Jsonerrorobject has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonerrorobject = d.Get("jsonerrorobject").(string)
		hasChange = true
	}
	if d.HasChange("jsonsqlinjectiontype") {
		log.Printf("[DEBUG]  citrixadc-provider: Jsonsqlinjectiontype has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonsqlinjectiontype = d.Get("jsonsqlinjectiontype").(string)
		hasChange = true
	}
	if d.HasChange("logeverypolicyhit") {
		log.Printf("[DEBUG]  citrixadc-provider: Logeverypolicyhit has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Logeverypolicyhit = d.Get("logeverypolicyhit").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("optimizepartialreqs") {
		log.Printf("[DEBUG]  citrixadc-provider: Optimizepartialreqs has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Optimizepartialreqs = d.Get("optimizepartialreqs").(string)
		hasChange = true
	}
	if d.HasChange("percentdecoderecursively") {
		log.Printf("[DEBUG]  citrixadc-provider: Percentdecoderecursively has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Percentdecoderecursively = d.Get("percentdecoderecursively").(string)
		hasChange = true
	}
	if d.HasChange("postbodylimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Postbodylimit has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Postbodylimit = d.Get("postbodylimit").(int)
		hasChange = true
	}
	if d.HasChange("postbodylimitsignature") {
		log.Printf("[DEBUG]  citrixadc-provider: Postbodylimitsignature has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Postbodylimitsignature = d.Get("postbodylimitsignature").(int)
		hasChange = true
	}
	if d.HasChange("refererheadercheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Refererheadercheck has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Refererheadercheck = d.Get("refererheadercheck").(string)
		hasChange = true
	}
	if d.HasChange("requestcontenttype") {
		log.Printf("[DEBUG]  citrixadc-provider: Requestcontenttype has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Requestcontenttype = d.Get("requestcontenttype").(string)
		hasChange = true
	}
	if d.HasChange("responsecontenttype") {
		log.Printf("[DEBUG]  citrixadc-provider: Responsecontenttype has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Responsecontenttype = d.Get("responsecontenttype").(string)
		hasChange = true
	}
	if d.HasChange("rfcprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Rfcprofile has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Rfcprofile = d.Get("rfcprofile").(string)
		hasChange = true
	}
	if d.HasChange("semicolonfieldseparator") {
		log.Printf("[DEBUG]  citrixadc-provider: Semicolonfieldseparator has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Semicolonfieldseparator = d.Get("semicolonfieldseparator").(string)
		hasChange = true
	}
	if d.HasChange("sessionlessfieldconsistency") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionlessfieldconsistency has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Sessionlessfieldconsistency = d.Get("sessionlessfieldconsistency").(string)
		hasChange = true
	}
	if d.HasChange("sessionlessurlclosure") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionlessurlclosure has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Sessionlessurlclosure = d.Get("sessionlessurlclosure").(string)
		hasChange = true
	}
	if d.HasChange("signatures") {
		log.Printf("[DEBUG]  citrixadc-provider: Signatures has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Signatures = d.Get("signatures").(string)
		hasChange = true
	}
	if d.HasChange("sqlinjectionchecksqlwildchars") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectionchecksqlwildchars has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Sqlinjectionchecksqlwildchars = d.Get("sqlinjectionchecksqlwildchars").(string)
		hasChange = true
	}
	if d.HasChange("sqlinjectiononlycheckfieldswithsqlchars") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectiononlycheckfieldswithsqlchars has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Sqlinjectiononlycheckfieldswithsqlchars = d.Get("sqlinjectiononlycheckfieldswithsqlchars").(string)
		hasChange = true
	}
	if d.HasChange("sqlinjectionparsecomments") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectionparsecomments has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Sqlinjectionparsecomments = d.Get("sqlinjectionparsecomments").(string)
		hasChange = true
	}
	if d.HasChange("sqlinjectiontransformspecialchars") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectiontransformspecialchars has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Sqlinjectiontransformspecialchars = d.Get("sqlinjectiontransformspecialchars").(string)
		hasChange = true
	}
	if d.HasChange("sqlinjectiontype") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectiontype has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Sqlinjectiontype = d.Get("sqlinjectiontype").(string)
		hasChange = true
	}
	if d.HasChange("starturlclosure") {
		log.Printf("[DEBUG]  citrixadc-provider: Starturlclosure has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Starturlclosure = d.Get("starturlclosure").(string)
		hasChange = true
	}
	if d.HasChange("streaming") {
		log.Printf("[DEBUG]  citrixadc-provider: Streaming has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Streaming = d.Get("streaming").(string)
		hasChange = true
	}
	if d.HasChange("stripcomments") {
		log.Printf("[DEBUG]  citrixadc-provider: Stripcomments has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Stripcomments = d.Get("stripcomments").(string)
		hasChange = true
	}
	if d.HasChange("striphtmlcomments") {
		log.Printf("[DEBUG]  citrixadc-provider: Striphtmlcomments has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Striphtmlcomments = d.Get("striphtmlcomments").(string)
		hasChange = true
	}
	if d.HasChange("stripxmlcomments") {
		log.Printf("[DEBUG]  citrixadc-provider: Stripxmlcomments has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Stripxmlcomments = d.Get("stripxmlcomments").(string)
		hasChange = true
	}
	if d.HasChange("trace") {
		log.Printf("[DEBUG]  citrixadc-provider: Trace has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Trace = d.Get("trace").(string)
		hasChange = true
	}
	if d.HasChange("urldecoderequestcookies") {
		log.Printf("[DEBUG]  citrixadc-provider: Urldecoderequestcookies has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Urldecoderequestcookies = d.Get("urldecoderequestcookies").(string)
		hasChange = true
	}
	if d.HasChange("usehtmlerrorobject") {
		log.Printf("[DEBUG]  citrixadc-provider: Usehtmlerrorobject has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Usehtmlerrorobject = d.Get("usehtmlerrorobject").(string)
		hasChange = true
	}
	if d.HasChange("verboseloglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Verboseloglevel has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Verboseloglevel = d.Get("verboseloglevel").(string)
		hasChange = true
	}
	if d.HasChange("xmlerrorobject") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlerrorobject has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlerrorobject = d.Get("xmlerrorobject").(string)
		hasChange = true
	}
	if d.HasChange("xmlsqlinjectionchecksqlwildchars") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlsqlinjectionchecksqlwildchars has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlsqlinjectionchecksqlwildchars = d.Get("xmlsqlinjectionchecksqlwildchars").(string)
		hasChange = true
	}
	if d.HasChange("xmlsqlinjectiononlycheckfieldswithsqlchars") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlsqlinjectiononlycheckfieldswithsqlchars has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlsqlinjectiononlycheckfieldswithsqlchars = d.Get("xmlsqlinjectiononlycheckfieldswithsqlchars").(string)
		hasChange = true
	}
	if d.HasChange("xmlsqlinjectionparsecomments") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlsqlinjectionparsecomments has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlsqlinjectionparsecomments = d.Get("xmlsqlinjectionparsecomments").(string)
		hasChange = true
	}
	if d.HasChange("xmlsqlinjectiontype") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlsqlinjectiontype has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlsqlinjectiontype = d.Get("xmlsqlinjectiontype").(string)
		hasChange = true
	}

	if d.HasChange("bufferoverflowaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Bufferoverflowaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Bufferoverflowaction = toStringList(d.Get("bufferoverflowaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("contenttypeaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Contenttypeaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Contenttypeaction = toStringList(d.Get("contenttypeaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("cookieconsistencyaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookieconsistencyaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Cookieconsistencyaction = toStringList(d.Get("cookieconsistencyaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("creditcard") {
		log.Printf("[DEBUG]  citrixadc-provider: Creditcard has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Creditcard = toStringList(d.Get("creditcard").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("creditcardaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Creditcardaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Creditcardaction = toStringList(d.Get("creditcardaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("crosssitescriptingaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Crosssitescriptingaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Crosssitescriptingaction = toStringList(d.Get("crosssitescriptingaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("csrftagaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Csrftagaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Csrftagaction = toStringList(d.Get("csrftagaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("denyurlaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Denyurlaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Denyurlaction = toStringList(d.Get("denyurlaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("dynamiclearning") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamiclearning has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Dynamiclearning = toStringList(d.Get("dynamiclearning").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("fieldconsistencyaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldconsistencyaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Fieldconsistencyaction = toStringList(d.Get("fieldconsistencyaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("fieldformataction") {
		log.Printf("[DEBUG]  citrixadc-provider: Fieldformataction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Fieldformataction = toStringList(d.Get("fieldformataction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("fileuploadtypesaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Fileuploadtypesaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Fileuploadtypesaction = toStringList(d.Get("fileuploadtypesaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("inspectcontenttypes") {
		log.Printf("[DEBUG]  citrixadc-provider: Inspectcontenttypes has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Inspectcontenttypes = toStringList(d.Get("inspectcontenttypes").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("jsondosaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Jsondosaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsondosaction = toStringList(d.Get("jsondosaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("jsonsqlinjectionaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Jsonsqlinjectionaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonsqlinjectionaction = toStringList(d.Get("jsonsqlinjectionaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("jsonxssaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Jsonxssaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonxssaction = toStringList(d.Get("jsonxssaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("multipleheaderaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Multipleheaderaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Multipleheaderaction = toStringList(d.Get("multipleheaderaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("sqlinjectionaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Sqlinjectionaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Sqlinjectionaction = toStringList(d.Get("sqlinjectionaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("starturlaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Starturlaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Starturlaction = toStringList(d.Get("starturlaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Type = toStringList(d.Get("type").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("xmlattachmentaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlattachmentaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlattachmentaction = toStringList(d.Get("xmlattachmentaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("xmldosaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmldosaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmldosaction = toStringList(d.Get("xmldosaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("xmlformataction") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlformataction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlformataction = toStringList(d.Get("xmlformataction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("xmlsoapfaultaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlsoapfaultaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlsoapfaultaction = toStringList(d.Get("xmlsoapfaultaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("xmlsqlinjectionaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlsqlinjectionaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlsqlinjectionaction = toStringList(d.Get("xmlsqlinjectionaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("xmlvalidationaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlvalidationaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlvalidationaction = toStringList(d.Get("xmlvalidationaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("xmlwsiaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlwsiaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlwsiaction = toStringList(d.Get("xmlwsiaction").(*schema.Set).List())
		hasChange = true
	}

	if d.HasChange("xmlxssaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Xmlxssaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlxssaction = toStringList(d.Get("xmlxssaction").(*schema.Set).List())
		hasChange = true
	}
	if d.HasChange("bufferoverflowmaxtotalheaderlength") {
		log.Printf("[DEBUG]  citrixadc-provider: bufferoverflowmaxtotalheaderlength has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Bufferoverflowmaxtotalheaderlength = d.Get("bufferoverflowmaxtotalheaderlength").(int)
		hasChange = true
	}
	if d.HasChange("cmdinjectiontype") {
		log.Printf("[DEBUG]  citrixadc-provider: cmdinjectiontype has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Cmdinjectiontype = d.Get("cmdinjectiontype").(string)
		hasChange = true
	}
	if d.HasChange("htmlerrorstatuscode") {
		log.Printf("[DEBUG]  citrixadc-provider: Htmlerrorstatuscode has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Htmlerrorstatuscode = d.Get("htmlerrorstatuscode").(int)
		hasChange = true
	}
	if d.HasChange("htmlerrorstatusmessage") {
		log.Printf("[DEBUG]  citrixadc-provider: Htmlerrorstatusmessage has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Htmlerrorstatusmessage = d.Get("htmlerrorstatusmessage").(string)
		hasChange = true
	}
	if d.HasChange("sqlinjectiongrammar") {
		log.Printf("[DEBUG]  citrixadc-provider: sqlinjectiongrammar has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Sqlinjectiongrammar = d.Get("sqlinjectiongrammar").(string)
		hasChange = true
	}

	// Add new attributes to update function
	if d.HasChange("apispec") {
		log.Printf("[DEBUG]  citrixadc-provider: apispec has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Apispec = d.Get("apispec").(string)
		hasChange = true
	}
	if d.HasChange("as_prof_bypass_list_enable") {
		log.Printf("[DEBUG]  citrixadc-provider: as_prof_bypass_list_enable has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Asprofbypasslistenable = d.Get("as_prof_bypass_list_enable").(string)
		hasChange = true
	}
	if d.HasChange("as_prof_deny_list_enable") {
		log.Printf("[DEBUG]  citrixadc-provider: as_prof_deny_list_enable has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Asprofdenylistenable = d.Get("as_prof_deny_list_enable").(string)
		hasChange = true
	}
	if d.HasChange("augment") {
		log.Printf("[DEBUG]  citrixadc-provider: augment has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Augment = d.Get("augment").(bool)
		hasChange = true
	}
	if d.HasChange("ceflogging") {
		log.Printf("[DEBUG]  citrixadc-provider: ceflogging has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Ceflogging = d.Get("ceflogging").(string)
		hasChange = true
	}
	if d.HasChange("clientipexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: clientipexpression has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Clientipexpression = d.Get("clientipexpression").(string)
		hasChange = true
	}
	if d.HasChange("cmdinjectiongrammar") {
		log.Printf("[DEBUG]  citrixadc-provider: cmdinjectiongrammar has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Cmdinjectiongrammar = d.Get("cmdinjectiongrammar").(string)
		hasChange = true
	}
	if d.HasChange("cookiesamesiteattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: cookiesamesiteattribute has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Cookiesamesiteattribute = d.Get("cookiesamesiteattribute").(string)
		hasChange = true
	}
	if d.HasChange("defaultfieldformatmaxoccurrences") {
		log.Printf("[DEBUG]  citrixadc-provider: defaultfieldformatmaxoccurrences has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Defaultfieldformatmaxoccurrences = d.Get("defaultfieldformatmaxoccurrences").(int)
		hasChange = true
	}
	if d.HasChange("fakeaccountdetection") {
		log.Printf("[DEBUG]  citrixadc-provider: fakeaccountdetection has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Fakeaccountdetection = d.Get("fakeaccountdetection").(string)
		hasChange = true
	}
	if d.HasChange("fieldscan") {
		log.Printf("[DEBUG]  citrixadc-provider: fieldscan has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Fieldscan = d.Get("fieldscan").(string)
		hasChange = true
	}
	if d.HasChange("fieldscanlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: fieldscanlimit has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Fieldscanlimit = d.Get("fieldscanlimit").(int)
		hasChange = true
	}
	if d.HasChange("geolocationlogging") {
		log.Printf("[DEBUG]  citrixadc-provider: geolocationlogging has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Geolocationlogging = d.Get("geolocationlogging").(string)
		hasChange = true
	}
	if d.HasChange("importprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: importprofilename has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Importprofilename = d.Get("importprofilename").(string)
		hasChange = true
	}
	if d.HasChange("insertcookiesamesiteattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: insertcookiesamesiteattribute has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Insertcookiesamesiteattribute = d.Get("insertcookiesamesiteattribute").(string)
		hasChange = true
	}
	if d.HasChange("jsoncmdinjectiongrammar") {
		log.Printf("[DEBUG]  citrixadc-provider: jsoncmdinjectiongrammar has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsoncmdinjectiongrammar = d.Get("jsoncmdinjectiongrammar").(string)
		hasChange = true
	}
	if d.HasChange("jsoncmdinjectiontype") {
		log.Printf("[DEBUG]  citrixadc-provider: jsoncmdinjectiontype has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsoncmdinjectiontype = d.Get("jsoncmdinjectiontype").(string)
		hasChange = true
	}
	if d.HasChange("jsonerrorstatuscode") {
		log.Printf("[DEBUG]  citrixadc-provider: jsonerrorstatuscode has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonerrorstatuscode = d.Get("jsonerrorstatuscode").(int)
		hasChange = true
	}
	if d.HasChange("jsonerrorstatusmessage") {
		log.Printf("[DEBUG]  citrixadc-provider: jsonerrorstatusmessage has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonerrorstatusmessage = d.Get("jsonerrorstatusmessage").(string)
		hasChange = true
	}
	if d.HasChange("jsonfieldscan") {
		log.Printf("[DEBUG]  citrixadc-provider: jsonfieldscan has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonfieldscan = d.Get("jsonfieldscan").(string)
		hasChange = true
	}
	if d.HasChange("jsonfieldscanlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: jsonfieldscanlimit has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonfieldscanlimit = d.Get("jsonfieldscanlimit").(int)
		hasChange = true
	}
	if d.HasChange("jsonmessagescan") {
		log.Printf("[DEBUG]  citrixadc-provider: jsonmessagescan has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonmessagescan = d.Get("jsonmessagescan").(string)
		hasChange = true
	}
	if d.HasChange("jsonmessagescanlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: jsonmessagescanlimit has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonmessagescanlimit = d.Get("jsonmessagescanlimit").(int)
		hasChange = true
	}
	if d.HasChange("jsonsqlinjectiongrammar") {
		log.Printf("[DEBUG]  citrixadc-provider: jsonsqlinjectiongrammar has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonsqlinjectiongrammar = d.Get("jsonsqlinjectiongrammar").(string)
		hasChange = true
	}
	if d.HasChange("matchurlstring") {
		log.Printf("[DEBUG]  citrixadc-provider: matchurlstring has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Matchurlstring = d.Get("matchurlstring").(string)
		hasChange = true
	}
	if d.HasChange("messagescan") {
		log.Printf("[DEBUG]  citrixadc-provider: messagescan has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Messagescan = d.Get("messagescan").(string)
		hasChange = true
	}
	if d.HasChange("messagescanlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: messagescanlimit has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Messagescanlimit = d.Get("messagescanlimit").(int)
		hasChange = true
	}
	if d.HasChange("overwrite") {
		log.Printf("[DEBUG]  citrixadc-provider: overwrite has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Overwrite = d.Get("overwrite").(bool)
		hasChange = true
	}
	if d.HasChange("protofileobject") {
		log.Printf("[DEBUG]  citrixadc-provider: protofileobject has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Protofileobject = d.Get("protofileobject").(string)
		hasChange = true
	}
	if d.HasChange("relaxationrules") {
		log.Printf("[DEBUG]  citrixadc-provider: relaxationrules has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Relaxationrules = d.Get("relaxationrules").(bool)
		hasChange = true
	}
	if d.HasChange("replaceurlstring") {
		log.Printf("[DEBUG]  citrixadc-provider: replaceurlstring has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Replaceurlstring = d.Get("replaceurlstring").(string)
		hasChange = true
	}
	if d.HasChange("sessioncookiename") {
		log.Printf("[DEBUG]  citrixadc-provider: sessioncookiename has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Sessioncookiename = d.Get("sessioncookiename").(string)
		hasChange = true
	}
	if d.HasChange("sqlinjectionruletype") {
		log.Printf("[DEBUG]  citrixadc-provider: sqlinjectionruletype has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Sqlinjectionruletype = d.Get("sqlinjectionruletype").(string)
		hasChange = true
	}
	if d.HasChange("xmlerrorstatuscode") {
		log.Printf("[DEBUG]  citrixadc-provider: xmlerrorstatuscode has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlerrorstatuscode = d.Get("xmlerrorstatuscode").(int)
		hasChange = true
	}
	if d.HasChange("xmlerrorstatusmessage") {
		log.Printf("[DEBUG]  citrixadc-provider: xmlerrorstatusmessage has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Xmlerrorstatusmessage = d.Get("xmlerrorstatusmessage").(string)
		hasChange = true
	}

	// Add new array/set type attributes to update function
	if d.HasChange("blockkeywordaction") {
		log.Printf("[DEBUG]  citrixadc-provider: blockkeywordaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Blockkeywordaction = toStringList(d.Get("blockkeywordaction").(*schema.Set).List())
		hasChange = true
	}
	if d.HasChange("grpcaction") {
		log.Printf("[DEBUG]  citrixadc-provider: grpcaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Grpcaction = toStringList(d.Get("grpcaction").(*schema.Set).List())
		hasChange = true
	}
	if d.HasChange("inspectquerycontenttypes") {
		log.Printf("[DEBUG]  citrixadc-provider: inspectquerycontenttypes has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Inspectquerycontenttypes = toStringList(d.Get("inspectquerycontenttypes").(*schema.Set).List())
		hasChange = true
	}
	if d.HasChange("jsonblockkeywordaction") {
		log.Printf("[DEBUG]  citrixadc-provider: jsonblockkeywordaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsonblockkeywordaction = toStringList(d.Get("jsonblockkeywordaction").(*schema.Set).List())
		hasChange = true
	}
	if d.HasChange("jsoncmdinjectionaction") {
		log.Printf("[DEBUG]  citrixadc-provider: jsoncmdinjectionaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Jsoncmdinjectionaction = toStringList(d.Get("jsoncmdinjectionaction").(*schema.Set).List())
		hasChange = true
	}
	if d.HasChange("messagescanlimitcontenttypes") {
		log.Printf("[DEBUG]  citrixadc-provider: messagescanlimitcontenttypes has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Messagescanlimitcontenttypes = toStringList(d.Get("messagescanlimitcontenttypes").(*schema.Set).List())
		hasChange = true
	}
	if d.HasChange("restaction") {
		log.Printf("[DEBUG]  citrixadc-provider: restaction has changed for appfwprofile %s, starting update", appfwprofileName)
		appfwprofile.Restaction = toStringList(d.Get("restaction").(*schema.Set).List())
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Appfwprofile.Type(), appfwprofileName, &appfwprofile)
		if err != nil {
			return fmt.Errorf("Error updating appfwprofile %s", appfwprofileName)
		}
	}
	return readAppfwprofileFunc(d, meta)
}

func deleteAppfwprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwprofileName := d.Id()
	err := client.DeleteResource(service.Appfwprofile.Type(), appfwprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
