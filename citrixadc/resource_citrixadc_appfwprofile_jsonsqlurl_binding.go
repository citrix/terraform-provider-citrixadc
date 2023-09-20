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

func resourceCitrixAdcAppfwprofile_jsonsqlurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_jsonsqlurl_bindingFunc,
		Read:          readAppfwprofile_jsonsqlurl_bindingFunc,
		Delete:        deleteAppfwprofile_jsonsqlurl_bindingFunc,
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
			"jsonsqlurl": {
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
			"keyname_json_sql": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"iskeyregex_json_sql": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_type_json_sql": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_expr_json_sql": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isvalueregex_json_sql": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_jsonsqlurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_jsonsqlurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	jsonsqlurl := d.Get("jsonsqlurl")
	bindingId := fmt.Sprintf("%s,%s", name, jsonsqlurl)
	appfwprofile_jsonsqlurl_binding := appfw.Appfwprofilejsonsqlurlbinding{
		Alertonly:              d.Get("alertonly").(string),
		Comment:                d.Get("comment").(string),
		Isautodeployed:         d.Get("isautodeployed").(string),
		Jsonsqlurl:             d.Get("jsonsqlurl").(string),
		Name:                   d.Get("name").(string),
		Resourceid:             d.Get("resourceid").(string),
		Ruletype:               d.Get("ruletype").(string),
		State:                  d.Get("state").(string),
		Keyname_json_sql:       d.Get("keyname_json_sql").(string),
		As_value_expr_json_sql: d.Get("as_value_expr_json_sql").(string),
		As_value_type_json_sql: d.Get("as_value_type_json_sql").(string),
		Iskeyregex_json_sql:    d.Get("iskeyregex_json_sql").(string),
		Isvalueregex_json_sql:  d.Get("isvalueregex_json_sql").(string),
	}

	err := client.UpdateUnnamedResource("appfwprofile_jsonsqlurl_binding", &appfwprofile_jsonsqlurl_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_jsonsqlurl_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_jsonsqlurl_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_jsonsqlurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_jsonsqlurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	jsonsqlurl := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_jsonsqlurl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_jsonsqlurl_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_jsonsqlurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["jsonsqlurl"].(string) == jsonsqlurl {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_jsonsqlurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	// d.Set("isautodeployed", data["isautodeployed"])
	d.Set("jsonsqlurl", data["jsonsqlurl"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])
	d.Set("keyname_json_sql", data["keyname_json_sql"])
	d.Set("as_value_expr_json_sql", data["as_value_expr_json_sql"])
	d.Set("as_value_type_json_sql", data["as_value_type_json_sql"])
	d.Set("iskeyregex_json_sql", data["iskeyregex_json_sql"])
	d.Set("isvalueregex_json_sql", data["isvalueregex_json_sql"])

	return nil

}

func deleteAppfwprofile_jsonsqlurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_jsonsqlurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	jsonsqlurl := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("jsonsqlurl:%s", url.QueryEscape(jsonsqlurl)))
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("keyname_json_sql"); ok {
		args = append(args, fmt.Sprintf("keyname_json_sql:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("as_value_type_json_sql"); ok {
		args = append(args, fmt.Sprintf("as_value_type_json_sql:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("as_value_expr_json_sql"); ok {
		args = append(args, fmt.Sprintf("as_value_expr_json_sql:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs("appfwprofile_jsonsqlurl_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
