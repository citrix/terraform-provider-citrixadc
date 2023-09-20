package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcAppfwprofile_xmlvalidationurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_xmlvalidationurl_bindingFunc,
		Read:          readAppfwprofile_xmlvalidationurl_bindingFunc,
		Delete:        deleteAppfwprofile_xmlvalidationurl_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"xmlvalidationurl": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"alertonly": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isautodeployed": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resourceid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ruletype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmladditionalsoapheaders": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlendpointcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlrequestschema": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlresponseschema": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlvalidateresponse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlvalidatesoapenvelope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlwsdl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_xmlvalidationurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_xmlvalidationurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	xmlvalidationurl := d.Get("xmlvalidationurl")
	bindingId := fmt.Sprintf("%s,%s", name, xmlvalidationurl)
	appfwprofile_xmlvalidationurl_binding := appfw.Appfwprofilexmlvalidationurlbinding{
		Alertonly:                d.Get("alertonly").(string),
		Comment:                  d.Get("comment").(string),
		Isautodeployed:           d.Get("isautodeployed").(string),
		Name:                     d.Get("name").(string),
		Resourceid:               d.Get("resourceid").(string),
		Ruletype:                 d.Get("ruletype").(string),
		State:                    d.Get("state").(string),
		Xmladditionalsoapheaders: d.Get("xmladditionalsoapheaders").(string),
		Xmlendpointcheck:         d.Get("xmlendpointcheck").(string),
		Xmlrequestschema:         d.Get("xmlrequestschema").(string),
		Xmlresponseschema:        d.Get("xmlresponseschema").(string),
		Xmlvalidateresponse:      d.Get("xmlvalidateresponse").(string),
		Xmlvalidatesoapenvelope:  d.Get("xmlvalidatesoapenvelope").(string),
		Xmlvalidationurl:         d.Get("xmlvalidationurl").(string),
		Xmlwsdl:                  d.Get("xmlwsdl").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_xmlvalidationurl_binding.Type(), &appfwprofile_xmlvalidationurl_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_xmlvalidationurl_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_xmlvalidationurl_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_xmlvalidationurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_xmlvalidationurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	xmlvalidationurl := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_xmlvalidationurl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_xmlvalidationurl_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_xmlvalidationurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["xmlvalidationurl"].(string) == xmlvalidationurl {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_xmlvalidationurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	// d.Set("isautodeployed", data["isautodeployed"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])
	d.Set("xmladditionalsoapheaders", data["xmladditionalsoapheaders"])
	d.Set("xmlendpointcheck", data["xmlendpointcheck"])
	d.Set("xmlrequestschema", data["xmlrequestschema"])
	d.Set("xmlresponseschema", data["xmlresponseschema"])
	d.Set("xmlvalidateresponse", data["xmlvalidateresponse"])
	d.Set("xmlvalidatesoapenvelope", data["xmlvalidatesoapenvelope"])
	d.Set("xmlvalidationurl", data["xmlvalidationurl"])
	d.Set("xmlwsdl", data["xmlwsdl"])

	return nil

}

func deleteAppfwprofile_xmlvalidationurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_xmlvalidationurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	xmlvalidationurl := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("xmlvalidationurl:%s", url.QueryEscape(xmlvalidationurl)))
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Appfwprofile_xmlvalidationurl_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
