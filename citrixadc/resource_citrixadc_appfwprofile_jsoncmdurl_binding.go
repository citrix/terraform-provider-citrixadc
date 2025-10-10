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

func resourceCitrixAdcAppfwprofile_jsoncmdurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_jsoncmdurl_bindingFunc,
		Read:          readAppfwprofile_jsoncmdurl_bindingFunc,
		Delete:        deleteAppfwprofile_jsoncmdurl_bindingFunc,
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
			"jsoncmdurl": {
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
			"iskeyregex_json_cmd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"keyname_json_cmd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_type_json_cmd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_expr_json_cmd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isvalueregex_json_cmd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_jsoncmdurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_jsoncmdurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	jsoncmdurl := d.Get("jsoncmdurl").(string)
	bindingId := fmt.Sprintf("%s,%s", name, jsoncmdurl)
	keyname_json_cmd := d.Get("keyname_json_cmd").(string)
	as_value_type_json_cmd := d.Get("as_value_type_json_cmd").(string)
	as_value_expr_json_cmd := d.Get("as_value_expr_json_cmd").(string)

	if keyname_json_cmd != "" {
		bindingId = fmt.Sprintf("%s,%s", bindingId, keyname_json_cmd)
		if as_value_type_json_cmd != "" && as_value_expr_json_cmd != "" {
			bindingId = fmt.Sprintf("%s,%s,%s", bindingId, as_value_type_json_cmd, as_value_expr_json_cmd)
		}
	}
	appfwprofile_jsoncmdurl_binding := appfw.Appfwprofilejsoncmdurlbinding{
		Alertonly:           d.Get("alertonly").(string),
		Comment:             d.Get("comment").(string),
		Isautodeployed:      d.Get("isautodeployed").(string),
		Jsoncmdurl:          d.Get("jsoncmdurl").(string),
		Name:                d.Get("name").(string),
		Resourceid:          d.Get("resourceid").(string),
		Ruletype:            d.Get("ruletype").(string),
		State:               d.Get("state").(string),
		Iskeyregexjsoncmd:   d.Get("iskeyregex_json_cmd").(string),
		Keynamejsoncmd:      d.Get("keyname_json_cmd").(string),
		Asvaluetypejsoncmd:  d.Get("as_value_type_json_cmd").(string),
		Asvalueexprjsoncmd:  d.Get("as_value_expr_json_cmd").(string),
		Isvalueregexjsoncmd: d.Get("isvalueregex_json_cmd").(string),
	}

	err := client.UpdateUnnamedResource("appfwprofile_jsoncmdurl_binding", &appfwprofile_jsoncmdurl_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_jsoncmdurl_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_jsoncmdurl_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_jsoncmdurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_jsoncmdurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	name := idSlice[0]
	jsoncmdurl := idSlice[1]
	keyname_json_cmd := ""
	as_value_type_json_cmd := ""
	as_value_expr_json_cmd := ""
	if len(idSlice) > 2 {
		keyname_json_cmd = idSlice[2]
	} else {
		keyname_json_cmd = d.Get("keyname_json_cmd").(string)
		if keyname_json_cmd != "" {
			bindingId = fmt.Sprintf("%s,%s", bindingId, keyname_json_cmd)
		}
	}
	if len(idSlice) > 4 {
		as_value_type_json_cmd = idSlice[3]
		as_value_expr_json_cmd = idSlice[4]
	} else {
		as_value_type_json_cmd = d.Get("as_value_type_json_cmd").(string)
		as_value_expr_json_cmd = d.Get("as_value_expr_json_cmd").(string)
		if as_value_type_json_cmd != "" && as_value_expr_json_cmd != "" {
			bindingId = fmt.Sprintf("%s,%s,%s", bindingId, as_value_type_json_cmd, as_value_expr_json_cmd)
		}
	}
	d.SetId(bindingId)
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_jsoncmdurl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_jsoncmdurl_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_jsoncmdurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the matching components
	foundIndex := -1
	for i, v := range dataArr {
		if v["jsoncmdurl"] != nil && v["jsoncmdurl"].(string) == jsoncmdurl {
			vKeyname := ""
			if v["keyname_json_cmd"] != nil {
				vKeyname = v["keyname_json_cmd"].(string)
			}
			if keyname_json_cmd != "" {
				if vKeyname == keyname_json_cmd {
					vType := ""
					vExpr := ""
					if v["as_value_type_json_cmd"] != nil {
						vType = v["as_value_type_json_cmd"].(string)
					}
					if v["as_value_expr_json_cmd"] != nil {
						vExpr = v["as_value_expr_json_cmd"].(string)
					}
					if as_value_type_json_cmd != "" && as_value_expr_json_cmd != "" {
						if strings.EqualFold(vType, as_value_type_json_cmd) && vExpr == as_value_expr_json_cmd {
							foundIndex = i
							break
						}
					} else if v["as_value_type_json_cmd"] == nil && v["as_value_expr_json_cmd"] == nil {
						foundIndex = i
						break
					}
				}
			} else if v["keyname_json_cmd"] == nil {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_jsoncmdurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	// d.Set("isautodeployed", data["isautodeployed"])
	d.Set("jsoncmdurl", data["jsoncmdurl"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])
	d.Set("iskeyregex_json_cmd", data["iskeyregex_json_cmd"])
	d.Set("keyname_json_cmd", data["keyname_json_cmd"])
	d.Set("as_value_type_json_cmd", data["as_value_type_json_cmd"])
	d.Set("as_value_expr_json_cmd", data["as_value_expr_json_cmd"])
	d.Set("isvalueregex_json_cmd", data["isvalueregex_json_cmd"])

	return nil

}

func deleteAppfwprofile_jsoncmdurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_jsoncmdurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	name := idSlice[0]
	jsoncmdurl := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("jsoncmdurl:%s", url.QueryEscape(jsoncmdurl)))
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("keyname_json_cmd"); ok {
		args = append(args, fmt.Sprintf("keyname_json_cmd:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("as_value_type_json_cmd"); ok {
		args = append(args, fmt.Sprintf("as_value_type_json_cmd:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("as_value_expr_json_cmd"); ok {
		args = append(args, fmt.Sprintf("as_value_expr_json_cmd:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs("appfwprofile_jsoncmdurl_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
