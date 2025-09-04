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

func resourceCitrixAdcAppfwprofile_jsonxssurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_jsonxssurl_bindingFunc,
		Read:          readAppfwprofile_jsonxssurl_bindingFunc,
		Delete:        deleteAppfwprofile_jsonxssurl_bindingFunc,
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
			"jsonxssurl": {
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
			"iskeyregex_json_xss": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"keyname_json_xss": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_type_json_xss": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_expr_json_xss": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isvalueregex_json_xss": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_jsonxssurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_jsonxssurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	jsonxssurl := d.Get("jsonxssurl")
	bindingId := fmt.Sprintf("%s,%s", name, jsonxssurl)
	appfwprofile_jsonxssurl_binding := appfw.Appfwprofilejsonxssurlbinding{
		Alertonly:           d.Get("alertonly").(string),
		Comment:             d.Get("comment").(string),
		Isautodeployed:      d.Get("isautodeployed").(string),
		Jsonxssurl:          d.Get("jsonxssurl").(string),
		Name:                d.Get("name").(string),
		Resourceid:          d.Get("resourceid").(string),
		Ruletype:            d.Get("ruletype").(string),
		State:               d.Get("state").(string),
		Iskeyregexjsonxss:   d.Get("iskeyregex_json_xss").(string),
		Keynamejsonxss:      d.Get("keyname_json_xss").(string),
		Asvaluetypejsonxss:  d.Get("as_value_type_json_xss").(string),
		Asvalueexprjsonxss:  d.Get("as_value_expr_json_xss").(string),
		Isvalueregexjsonxss: d.Get("isvalueregex_json_xss").(string),
	}

	err := client.UpdateUnnamedResource("appfwprofile_jsonxssurl_binding", &appfwprofile_jsonxssurl_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_jsonxssurl_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_jsonxssurl_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_jsonxssurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_jsonxssurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	jsonxssurl := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_jsonxssurl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_jsonxssurl_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_jsonxssurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["jsonxssurl"].(string) == jsonxssurl {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_jsonxssurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	// d.Set("isautodeployed", data["isautodeployed"])
	d.Set("jsonxssurl", data["jsonxssurl"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])
	d.Set("iskeyregex_json_xss", data["iskeyregex_jsonxss"])
	d.Set("keyname_json_xss", data["keynamejsonxss"])
	d.Set("as_value_type_json_xss", data["asvaluetypejsonxss"])
	d.Set("as_value_expr_json_xss", data["asvalueexprjsonxss"])
	d.Set("isvalueregex_json_xss", data["isvalueregexjsonxss"])

	return nil

}

func deleteAppfwprofile_jsonxssurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_jsonxssurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	jsonxssurl := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("jsonxssurl:%s", url.QueryEscape(jsonxssurl)))
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs("appfwprofile_jsonxssurl_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
