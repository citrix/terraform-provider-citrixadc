package citrixadc

import (
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"fmt"
	"log"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCitrixAdcAppfwprofile_crosssitescripting_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_crosssitescripting_bindingFunc,
		Read:          readAppfwprofile_crosssitescripting_bindingFunc,
		Delete:        deleteAppfwprofile_crosssitescripting_bindingFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"crosssitescripting": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"formactionurl_xss": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"as_scan_location_xss": {
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
			"as_value_expr_xss": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_type_xss": {
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
			"isregex_xss": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isvalueregex_xss": {
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
		},
	}
}

func createAppfwprofile_crosssitescripting_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_crosssitescripting_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appFwName := d.Get("name").(string)
	crosssitescripting := d.Get("crosssitescripting").(string)
	formactionurl_xss := d.Get("formactionurl_xss").(string)
	as_scan_location_xss := d.Get("as_scan_location_xss").(string)
	as_value_type_xss := d.Get("as_value_type_xss").(string)
	as_value_expr_xss := d.Get("as_value_expr_xss").(string)
	bindingId := fmt.Sprintf("%s,%s,%s,%s", appFwName, crosssitescripting, formactionurl_xss, as_scan_location_xss)
	if as_value_type_xss != "" && as_value_expr_xss != "" {
		bindingId = fmt.Sprintf("%s,%s,%s", bindingId, as_value_type_xss, as_value_expr_xss)
	}

	appfwprofile_crosssitescripting_binding := appfw.Appfwprofilecrosssitescriptingbinding{
		Alertonly:          d.Get("alertonly").(string),
		Asscanlocationxss:  d.Get("as_scan_location_xss").(string),
		Asvalueexprxss:     d.Get("as_value_expr_xss").(string),
		Asvaluetypexss:     d.Get("as_value_type_xss").(string),
		Comment:            d.Get("comment").(string),
		Crosssitescripting: crosssitescripting,
		Formactionurlxss:   d.Get("formactionurl_xss").(string),
		Isautodeployed:     d.Get("isautodeployed").(string),
		Isregexxss:         d.Get("isregex_xss").(string),
		Isvalueregexxss:    d.Get("isvalueregex_xss").(string),
		Name:               appFwName,
		State:              d.Get("state").(string),
	}

	_, err := client.AddResource(service.Appfwprofile_crosssitescripting_binding.Type(), appFwName, &appfwprofile_crosssitescripting_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_crosssitescripting_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_crosssitescripting_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_crosssitescripting_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_crosssitescripting_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: readAppfwprofile_crosssitescripting_bindingFunc: bindingId: %s", bindingId)
	idSlice := strings.Split(bindingId, ",")
	appFwName := idSlice[0]
	crosssitescripting := idSlice[1]
	formactionurl_xss := ""
	as_scan_location_xss := ""
	as_value_type_xss := ""
	as_value_expr_xss := ""
	if len(idSlice) > 2 {
		formactionurl_xss = idSlice[2]
		as_scan_location_xss = idSlice[3]
	} else {
		formactionurl_xss = d.Get("formactionurl_xss").(string)
		as_scan_location_xss = d.Get("as_scan_location_xss").(string)
		bindingId = fmt.Sprintf("%s,%s,%s", bindingId, formactionurl_xss, as_scan_location_xss)
	}
	if len(idSlice) > 4 {
		as_value_type_xss = idSlice[4]
		as_value_expr_xss = idSlice[5]
	} else {
		as_value_type_xss = d.Get("as_value_type_xss").(string)
		as_value_expr_xss = d.Get("as_value_expr_xss").(string)
		if as_value_type_xss != "" && as_value_expr_xss != "" {
			bindingId = fmt.Sprintf("%s,%s,%s", bindingId, as_value_type_xss, as_value_expr_xss)
		}
	}
	d.SetId(bindingId)
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_crosssitescripting_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_crosssitescripting_binding.Type(),
		ResourceName:             appFwName,
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_crosssitescripting_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if v["crosssitescripting"].(string) == crosssitescripting && v["formactionurl_xss"].(string) == formactionurl_xss && v["as_scan_location_xss"].(string) == as_scan_location_xss {
			if as_value_type_xss != "" && as_value_expr_xss != "" {
				if v["as_value_type_xss"] != nil && v["as_value_expr_xss"] != nil && v["as_value_type_xss"].(string) == as_value_type_xss && v["as_value_expr_xss"].(string) == as_value_expr_xss {
					foundIndex = i
					break
				}
			} else if v["as_value_type_xss"] == nil && v["as_value_expr_xss"] == nil {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_crosssitescripting_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough
	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("alertonly", data["alertonly"])
	d.Set("as_scan_location_xss", data["as_scan_location_xss"])
	d.Set("as_value_expr_xss", data["as_value_expr_xss"])
	d.Set("as_value_type_xss", data["as_value_type_xss"])
	d.Set("comment", data["comment"])
	d.Set("crosssitescripting", data["crosssitescripting"])
	d.Set("formactionurl_xss", data["formactionurl_xss"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregex_xss", data["isregex_xss"])
	d.Set("isvalueregex_xss", data["isvalueregex_xss"])
	d.Set("name", data["name"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_crosssitescripting_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_crosssitescripting_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")
	appFwName := idSlice[0]
	crosssitescripting := idSlice[1]

	args := make(map[string]string)
	args["crosssitescripting"] = crosssitescripting
	args["formactionurl_xss"] = url.QueryEscape(d.Get("formactionurl_xss").(string))
	args["as_scan_location_xss"] = d.Get("as_scan_location_xss").(string)

	if val, ok := d.GetOk("as_value_type_xss"); ok {
		args["as_value_type_xss"] = url.QueryEscape(val.(string))
	}
	if val, ok := d.GetOk("as_value_expr_xss"); ok {
		args["as_value_expr_xss"] = url.QueryEscape(val.(string))
	}
	if val, ok := d.GetOk("ruletype"); ok {
		args["ruletype"] = url.QueryEscape(val.(string))
	}

	err := client.DeleteResourceWithArgsMap(service.Appfwprofile_crosssitescripting_binding.Type(), appFwName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
