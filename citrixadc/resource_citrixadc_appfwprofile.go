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
		Schema: map[string]*schema.Schema{
			"addcookieflags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"archivename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bufferoverflowmaxcookielength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bufferoverflowmaxheaderlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bufferoverflowmaxurllength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"canonicalizehtmlresponse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"checkrequestheaders": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookieencryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookieproxying": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookietransforms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"creditcardmaxallowed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"creditcardxout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crosssitescriptingcheckcompleteurls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crosssitescriptingtransformunsafehtml": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customsettings": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultcharset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultfieldformatmaxlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultfieldformatminlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultfieldformattype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaults": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dosecurecreditcardlogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enableformtagging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"errorurl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"excludefileuploadfromchecks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exemptclosureurlsfromsecuritychecks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fileuploadmaxnum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"htmlerrorobject": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalidpercenthandling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jsonerrorobject": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jsonsqlinjectiontype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logeverypolicyhit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"optimizepartialreqs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"percentdecoderecursively": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"postbodylimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"postbodylimitsignature": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"refererheadercheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requestcontenttype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"responsecontenttype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rfcprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"semicolonfieldseparator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionlessfieldconsistency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionlessurlclosure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signatures": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectionchecksqlwildchars": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectiononlycheckfieldswithsqlchars": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectionparsecomments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectiontransformspecialchars": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sqlinjectiontype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"starturlclosure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"postbodylimitaction": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"bufferoverflowmaxquerylength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cookiehijackingaction": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"infercontenttypexmlpayloadaction": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cmdinjectionaction": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"streaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stripcomments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"striphtmlcomments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stripxmlcomments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"urldecoderequestcookies": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usehtmlerrorobject": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verboseloglevel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xmlerrorobject": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xmlsqlinjectionchecksqlwildchars": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xmlsqlinjectiononlycheckfieldswithsqlchars": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xmlsqlinjectionparsecomments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xmlsqlinjectiontype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bufferoverflowaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"contenttypeaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cookieconsistencyaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"creditcard": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"creditcardaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"crosssitescriptingaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"csrftagaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"denyurlaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dynamiclearning": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fieldconsistencyaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fieldformataction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fileuploadtypesaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"inspectcontenttypes": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"jsondosaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"jsonsqlinjectionaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"jsonxssaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"multipleheaderaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sqlinjectionaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"starturlaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"type": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlattachmentaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmldosaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlformataction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlsoapfaultaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlsqlinjectionaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlvalidationaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlwsiaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"xmlxssaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
		Cookiehijackingaction:    					toStringList(d.Get("cookiehijackingaction").([]interface{})),
		Infercontenttypexmlpayloadaction:   		toStringList(d.Get("infercontenttypexmlpayloadaction").([]interface{})),
		Cmdinjectionaction: 						toStringList(d.Get("cmdinjectionaction").([]interface{})),
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
	d.Set("bufferoverflowmaxquerylength", data["bufferoverflowmaxquerylength"])
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
	d.Set("sqlinjectiononlycheckfieldswithsqlchars", data["sqlinjectiononlycheckfieldswithsqlchars"])
	d.Set("sqlinjectionparsecomments", data["sqlinjectionparsecomments"])
	d.Set("sqlinjectiontransformspecialchars", data["sqlinjectiontransformspecialchars"])
	d.Set("sqlinjectiontype", data["sqlinjectiontype"])
	d.Set("starturlclosure", data["starturlclosure"])
	d.Set("streaming", data["streaming"])
	d.Set("stripcomments", data["stripcomments"])
	d.Set("striphtmlcomments", data["striphtmlcomments"])
	d.Set("stripxmlcomments", data["stripxmlcomments"])
	d.Set("trace", data["trace"])
	d.Set("urldecoderequestcookies", data["urldecoderequestcookies"])
	d.Set("usehtmlerrorobject", data["usehtmlerrorobject"])
	d.Set("verboseloglevel", data["verboseloglevel"])
	d.Set("xmlerrorobject", data["xmlerrorobject"])
	d.Set("xmlsqlinjectionchecksqlwildchars", data["xmlsqlinjectionchecksqlwildchars"])
	d.Set("xmlsqlinjectiononlycheckfieldswithsqlchars", data["xmlsqlinjectiononlycheckfieldswithsqlchars"])
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
