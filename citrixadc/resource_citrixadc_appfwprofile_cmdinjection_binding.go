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

func resourceCitrixAdcAppfwprofile_cmdinjection_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_cmdinjection_bindingFunc,
		Read:          readAppfwprofile_cmdinjection_bindingFunc,
		Delete:        deleteAppfwprofile_cmdinjection_bindingFunc,
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
			"cmdinjection": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"formactionurl_cmd": {
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
			"as_scan_location_cmd": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "FORMFIELD",
			},
			"as_value_expr_cmd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"as_value_type_cmd": {
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
			"isregex_cmd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isvalueregex_cmd": {
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
		},
	}
}

func createAppfwprofile_cmdinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_cmdinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appFwName := d.Get("name").(string)
	cmdinjection := d.Get("cmdinjection").(string)
	formactionurl_cmd := d.Get("formactionurl_cmd").(string)
	as_scan_location_cmd := d.Get("as_scan_location_cmd").(string)
	as_value_type_cmd := d.Get("as_value_type_cmd").(string)
	as_value_expr_cmd := d.Get("as_value_expr_cmd").(string)
	bindingId := fmt.Sprintf("%s,%s,%s,%s", appFwName, cmdinjection, formactionurl_cmd, as_scan_location_cmd)
	if as_value_type_cmd != "" && as_value_expr_cmd != "" {
		bindingId = fmt.Sprintf("%s,%s,%s", bindingId, as_value_type_cmd, as_value_expr_cmd)
	}
	appfwprofile_cmdinjection_binding := appfw.Appfwprofilecmdinjectionbinding{
		Alertonly:         d.Get("alertonly").(string),
		Asscanlocationcmd: d.Get("as_scan_location_cmd").(string),
		Asvalueexprcmd:    d.Get("as_value_expr_cmd").(string),
		Asvaluetypecmd:    d.Get("as_value_type_cmd").(string),
		Cmdinjection:      d.Get("cmdinjection").(string),
		Comment:           d.Get("comment").(string),
		Formactionurlcmd:  d.Get("formactionurl_cmd").(string),
		Isautodeployed:    d.Get("isautodeployed").(string),
		Isregexcmd:        d.Get("isregex_cmd").(string),
		Isvalueregexcmd:   d.Get("isvalueregex_cmd").(string),
		Name:              appFwName,
		Resourceid:        d.Get("resourceid").(string),
		Ruletype:          d.Get("ruletype").(string),
		State:             d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource("appfwprofile_cmdinjection_binding", &appfwprofile_cmdinjection_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_cmdinjection_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_cmdinjection_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_cmdinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_cmdinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: readAppfwprofile_cmdinjection_bindingFunc: bindingId: %s", bindingId)
	idSlice := strings.Split(bindingId, ",")
	appFwName := idSlice[0]
	cmdinjection := idSlice[1]
	formactionurl_cmd := idSlice[2]
	as_scan_location_cmd := ""
	as_value_type_cmd := ""
	as_value_expr_cmd := ""
	if len(idSlice) > 3 {
		as_scan_location_cmd = idSlice[3]
	} else {
		as_scan_location_cmd = d.Get("as_scan_location_cmd").(string)
		bindingId = fmt.Sprintf("%s,%s", bindingId, as_scan_location_cmd)
	}
	if len(idSlice) > 4 {
		as_value_type_cmd = idSlice[4]
		as_value_expr_cmd = idSlice[5]
	} else {
		as_value_type_cmd = d.Get("as_value_type_cmd").(string)
		as_value_expr_cmd = d.Get("as_value_expr_cmd").(string)
		if as_value_type_cmd != "" && as_value_expr_cmd != "" {
			bindingId = fmt.Sprintf("%s,%s,%s", bindingId, as_value_type_cmd, as_value_expr_cmd)
		}
	}
	d.SetId(bindingId)
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_cmdinjection_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_cmdinjection_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_cmdinjection_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if v["cmdinjection"].(string) == cmdinjection && v["formactionurl_cmd"].(string) == formactionurl_cmd && v["as_scan_location_cmd"].(string) == as_scan_location_cmd {
			if as_value_type_cmd != "" && as_value_expr_cmd != "" {
				if v["as_value_type_cmd"] != nil && v["as_value_expr_cmd"] != nil && v["as_value_type_cmd"].(string) == as_value_type_cmd && v["as_value_expr_cmd"].(string) == as_value_expr_cmd {
					foundIndex = i
					break
				}
			} else if v["as_value_type_cmd"] == nil && v["as_value_expr_cmd"] == nil {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams appfwprofile_cmdinjection_binding not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_cmdinjection_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("alertonly", data["alertonly"])
	d.Set("as_scan_location_cmd", data["as_scan_location_cmd"])
	d.Set("as_value_expr_cmd", data["as_value_expr_cmd"])
	d.Set("as_value_type_cmd", data["as_value_type_cmd"])
	d.Set("cmdinjection", data["cmdinjection"])
	d.Set("comment", data["comment"])
	d.Set("formactionurl_cmd", data["formactionurl_cmd"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregex_cmd", data["isregex_cmd"])
	d.Set("isvalueregex_cmd", data["isvalueregex_cmd"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_cmdinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_cmdinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")
	appFwName := idSlice[0]
	cmdinjection := idSlice[1]

	args := make(map[string]string)
	args["cmdinjection"] = cmdinjection
	args["formactionurl_cmd"] = url.QueryEscape(d.Get("formactionurl_cmd").(string))
	args["as_scan_location_cmd"] = d.Get("as_scan_location_cmd").(string)

	if val, ok := d.GetOk("as_value_type_cmd"); ok {
		args["as_value_type_cmd"] = url.QueryEscape(val.(string))
	}
	if val, ok := d.GetOk("as_value_expr_cmd"); ok {
		args["as_value_expr_cmd"] = url.QueryEscape(val.(string))
	}
	if val, ok := d.GetOk("ruletype"); ok {
		args["ruletype"] = url.QueryEscape(val.(string))
	}

	err := client.DeleteResourceWithArgsMap("appfwprofile_cmdinjection_binding", appFwName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
