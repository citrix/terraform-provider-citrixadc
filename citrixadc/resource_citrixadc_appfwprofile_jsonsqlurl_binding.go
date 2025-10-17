package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcAppfwprofile_jsonsqlurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwprofile_jsonsqlurl_bindingFunc,
		ReadContext:   readAppfwprofile_jsonsqlurl_bindingFunc,
		DeleteContext: deleteAppfwprofile_jsonsqlurl_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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

func createAppfwprofile_jsonsqlurl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_jsonsqlurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	jsonsqlurl := d.Get("jsonsqlurl").(string)
	bindingId := fmt.Sprintf("%s,%s", name, jsonsqlurl)
	keyname_json_sql := d.Get("keyname_json_sql").(string)
	as_value_type_json_sql := d.Get("as_value_type_json_sql").(string)
	as_value_expr_json_sql := d.Get("as_value_expr_json_sql").(string)

	if keyname_json_sql != "" {
		bindingId = fmt.Sprintf("%s,%s", bindingId, keyname_json_sql)
		if as_value_type_json_sql != "" && as_value_expr_json_sql != "" {
			bindingId = fmt.Sprintf("%s,%s,%s", bindingId, as_value_type_json_sql, url.QueryEscape(as_value_expr_json_sql))
		}
	}
	appfwprofile_jsonsqlurl_binding := appfw.Appfwprofilejsonsqlurlbinding{
		Alertonly:           d.Get("alertonly").(string),
		Comment:             d.Get("comment").(string),
		Isautodeployed:      d.Get("isautodeployed").(string),
		Jsonsqlurl:          d.Get("jsonsqlurl").(string),
		Name:                d.Get("name").(string),
		Resourceid:          d.Get("resourceid").(string),
		Ruletype:            d.Get("ruletype").(string),
		State:               d.Get("state").(string),
		Keynamejsonsql:      d.Get("keyname_json_sql").(string),
		Asvalueexprjsonsql:  d.Get("as_value_expr_json_sql").(string),
		Asvaluetypejsonsql:  d.Get("as_value_type_json_sql").(string),
		Iskeyregexjsonsql:   d.Get("iskeyregex_json_sql").(string),
		Isvalueregexjsonsql: d.Get("isvalueregex_json_sql").(string),
	}

	err := client.UpdateUnnamedResource("appfwprofile_jsonsqlurl_binding", &appfwprofile_jsonsqlurl_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAppfwprofile_jsonsqlurl_bindingFunc(ctx, d, meta)
}

func readAppfwprofile_jsonsqlurl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_jsonsqlurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	name := idSlice[0]
	jsonsqlurl := idSlice[1]
	keyname_json_sql := ""
	as_value_type_json_sql := ""
	as_value_expr_json_sql := ""
	if len(idSlice) > 2 {
		keyname_json_sql = idSlice[2]
	} else {
		keyname_json_sql = d.Get("keyname_json_sql").(string)
		if keyname_json_sql != "" {
			bindingId = fmt.Sprintf("%s,%s", bindingId, keyname_json_sql)
		}
	}
	if len(idSlice) > 4 {
		as_value_type_json_sql = idSlice[3]
		as_value_expr_json_sql = idSlice[4]
	} else {
		as_value_type_json_sql = d.Get("as_value_type_json_sql").(string)
		as_value_expr_json_sql = d.Get("as_value_expr_json_sql").(string)
		if as_value_type_json_sql != "" && as_value_expr_json_sql != "" {
			bindingId = fmt.Sprintf("%s,%s,%s", bindingId, as_value_type_json_sql, url.QueryEscape(as_value_expr_json_sql))
		}
	}
	d.SetId(bindingId)
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
		return diag.FromErr(err)
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_jsonsqlurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the matching components
	foundIndex := -1
	for i, v := range dataArr {
		if v["jsonsqlurl"] != nil && v["jsonsqlurl"].(string) == jsonsqlurl {
			vKeyname := ""
			if v["keyname_json_sql"] != nil {
				vKeyname = v["keyname_json_sql"].(string)
			}
			if keyname_json_sql != "" {
				if vKeyname == keyname_json_sql {
					vType := ""
					vExpr := ""
					if v["as_value_type_json_sql"] != nil {
						vType = v["as_value_type_json_sql"].(string)
					}
					if v["as_value_expr_json_sql"] != nil {
						vExpr = v["as_value_expr_json_sql"].(string)
					}
					if as_value_type_json_sql != "" && as_value_expr_json_sql != "" {
						if strings.EqualFold(vType, as_value_type_json_sql) && vExpr == as_value_expr_json_sql {
							foundIndex = i
							break
						}
					} else if v["as_value_type_json_sql"] == nil && v["as_value_expr_json_sql"] == nil {
						foundIndex = i
						break
					}
				}
			} else if v["keyname_json_sql"] == nil {
				foundIndex = i
				break
			}
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

func deleteAppfwprofile_jsonsqlurl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_jsonsqlurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

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
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
