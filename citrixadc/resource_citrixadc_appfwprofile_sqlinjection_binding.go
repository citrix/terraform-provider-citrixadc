package citrixadc

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppfwprofile_sqlinjection_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwprofile_sqlinjection_bindingFunc,
		ReadContext:   readAppfwprofile_sqlinjection_bindingFunc,
		DeleteContext: deleteAppfwprofile_sqlinjection_bindingFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sqlinjection": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"as_scan_location_sql": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"formactionurl_sql": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alertonly": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_expr_sql": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_type_sql": {
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
			"isregex_sql": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isvalueregex_sql": {
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
			"ruletype": {
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
		},
	}
}

func createAppfwprofile_sqlinjection_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_sqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appFwName := d.Get("name").(string)
	sqlinjection := d.Get("sqlinjection").(string)
	formactionurl_sql := d.Get("formactionurl_sql").(string)
	as_scan_location_sql := d.Get("as_scan_location_sql").(string)
	as_value_type_sql := d.Get("as_value_type_sql").(string)
	as_value_expr_sql := d.Get("as_value_expr_sql").(string)
	bindingId := fmt.Sprintf("%s,%s,%s,%s", appFwName, sqlinjection, formactionurl_sql, as_scan_location_sql)
	if as_value_type_sql != "" && as_value_expr_sql != "" {
		rule_type := d.Get("ruletype").(string)
		if rule_type == "" {
			rule_type = "ALLOW"
		}
		bindingId = fmt.Sprintf("%s,%s,%s,%s", bindingId, as_value_type_sql, as_value_expr_sql, rule_type)
	}

	appfwprofile_sqlinjection_binding := appfw.Appfwprofilesqlinjectionbinding{
		Alertonly:         d.Get("alertonly").(string),
		Asscanlocationsql: d.Get("as_scan_location_sql").(string),
		Asvalueexprsql:    d.Get("as_value_expr_sql").(string),
		Asvaluetypesql:    d.Get("as_value_type_sql").(string),
		Comment:           d.Get("comment").(string),
		Formactionurlsql:  d.Get("formactionurl_sql").(string),
		Isautodeployed:    d.Get("isautodeployed").(string),
		Isregexsql:        d.Get("isregex_sql").(string),
		Isvalueregexsql:   d.Get("isvalueregex_sql").(string),
		Name:              appFwName,
		Sqlinjection:      sqlinjection,
		State:             d.Get("state").(string),
		Ruletype:          d.Get("ruletype").(string),
		Resourceid:        d.Get("resourceid").(string),
	}

	_, err := client.AddResource(service.Appfwprofile_sqlinjection_binding.Type(), appFwName, &appfwprofile_sqlinjection_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAppfwprofile_sqlinjection_bindingFunc(ctx, d, meta)
}

func readAppfwprofile_sqlinjection_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_sqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: readAppfwprofile_sqlinjection_bindingFunc: bindingId: %s", bindingId)
	idSlice := strings.Split(bindingId, ",")
	appFwName := idSlice[0]
	sqlinjection := idSlice[1]
	formactionurl_sql := ""
	as_scan_location_sql := ""
	as_value_type_sql := ""
	as_value_expr_sql := ""
	rule_type := ""
	if len(idSlice) > 2 {
		formactionurl_sql = idSlice[2]
		as_scan_location_sql = idSlice[3]
	} else {
		formactionurl_sql = d.Get("formactionurl_sql").(string)
		as_scan_location_sql = d.Get("as_scan_location_sql").(string)
		bindingId = fmt.Sprintf("%s,%s,%s", bindingId, formactionurl_sql, as_scan_location_sql)
	}
	if len(idSlice) > 4 {
		as_value_type_sql = idSlice[4]
		as_value_expr_sql = idSlice[5]
		rule_type = idSlice[6]
	} else {
		as_value_type_sql = d.Get("as_value_type_sql").(string)
		as_value_expr_sql = d.Get("as_value_expr_sql").(string)
		if as_value_type_sql != "" && as_value_expr_sql != "" {
			rule_type = d.Get("ruletype").(string)
			if rule_type == "" {
				rule_type = "ALLOW"
			}
			bindingId = fmt.Sprintf("%s,%s,%s,%s", bindingId, as_value_type_sql, as_value_expr_sql, rule_type)
		}
	}
	d.SetId(bindingId)
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_sqlinjection_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_sqlinjection_binding.Type(),
		ResourceName:             appFwName,
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_sqlinjection_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if v["sqlinjection"].(string) == sqlinjection && v["formactionurl_sql"].(string) == formactionurl_sql && v["as_scan_location_sql"].(string) == as_scan_location_sql {
			if as_value_type_sql != "" && as_value_expr_sql != "" {
				if v["as_value_type_sql"] != nil && v["as_value_expr_sql"] != nil && v["as_value_type_sql"].(string) == as_value_type_sql && v["as_value_expr_sql"].(string) == as_value_expr_sql && v["ruletype"].(string) == rule_type {
					foundIndex = i
					break
				}
			} else if v["as_value_type_sql"] == nil && v["as_value_expr_sql"] == nil {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams appfwprofile_sqlinjection_binding not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_sqlinjection_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough
	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("alertonly", data["alertonly"])
	d.Set("as_scan_location_sql", data["as_scan_location_sql"])
	d.Set("as_value_expr_sql", data["as_value_expr_sql"])
	d.Set("as_value_type_sql", data["as_value_type_sql"])
	d.Set("comment", data["comment"])
	d.Set("formactionurl_sql", data["formactionurl_sql"].(string))
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregex_sql", data["isregex_sql"])
	d.Set("isvalueregex_sql", data["isvalueregex_sql"])
	d.Set("name", data["name"])
	d.Set("sqlinjection", data["sqlinjection"])
	d.Set("state", data["state"])
	d.Set("ruletype", data["ruletype"])
	d.Set("resourceid", data["resourceid"])

	return nil

}

func deleteAppfwprofile_sqlinjection_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_sqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")
	appFwName := idSlice[0]
	sqlinjection := idSlice[1]

	args := make(map[string]string)
	args["sqlinjection"] = sqlinjection
	args["formactionurl_sql"] = url.QueryEscape(d.Get("formactionurl_sql").(string))
	args["as_scan_location_sql"] = d.Get("as_scan_location_sql").(string)

	if val, ok := d.GetOk("as_value_type_sql"); ok {
		args["as_value_type_sql"] = url.QueryEscape(val.(string))
	}
	if val, ok := d.GetOk("as_value_expr_sql"); ok {
		args["as_value_expr_sql"] = url.QueryEscape(val.(string))
	}
	if val, ok := d.GetOk("ruletype"); ok {
		args["ruletype"] = url.QueryEscape(val.(string))
	}

	err := client.DeleteResourceWithArgsMap(service.Appfwprofile_sqlinjection_binding.Type(), appFwName, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
