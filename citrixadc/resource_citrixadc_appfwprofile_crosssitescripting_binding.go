package citrixadc

import (
	"net/url"

	"github.com/chiradeep/go-nitro/config/appfw"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwprofile_crosssitescripting_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_crosssitescripting_bindingFunc,
		Read:          readAppfwprofile_crosssitescripting_bindingFunc,
		Delete:        deleteAppfwprofile_crosssitescripting_bindingFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"crosssitescripting": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"formactionurl_xss": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"as_scan_location_xss": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alertonly": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"as_value_expr_xss": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"as_value_type_xss": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isautodeployed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isregex_xss": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isvalueregex_xss": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppfwprofile_crosssitescripting_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_crosssitescripting_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwprofile_crosssitescripting_bindingName := d.Get("name").(string)
	appfwprofile_crosssitescripting_binding := appfw.Appfwprofilecrosssitescriptingbinding{
		Alertonly:          d.Get("alertonly").(string),
		Asscanlocationxss:  d.Get("as_scan_location_xss").(string),
		Asvalueexprxss:     d.Get("as_value_expr_xss").(string),
		Asvaluetypexss:     d.Get("as_value_type_xss").(string),
		Comment:            d.Get("comment").(string),
		Crosssitescripting: d.Get("crosssitescripting").(string),
		Formactionurlxss:   d.Get("formactionurl_xss").(string),
		Isautodeployed:     d.Get("isautodeployed").(string),
		Isregexxss:         d.Get("isregex_xss").(string),
		Isvalueregexxss:    d.Get("isvalueregex_xss").(string),
		Name:               d.Get("name").(string),
		State:              d.Get("state").(string),
	}

	_, err := client.AddResource(netscaler.Appfwprofile_crosssitescripting_binding.Type(), appfwprofile_crosssitescripting_bindingName, &appfwprofile_crosssitescripting_binding)
	if err != nil {
		return err
	}

	d.SetId(appfwprofile_crosssitescripting_bindingName)

	err = readAppfwprofile_crosssitescripting_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_crosssitescripting_binding but we can't read it ?? %s", appfwprofile_crosssitescripting_bindingName)
		return nil
	}
	return nil
}

func readAppfwprofile_crosssitescripting_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_crosssitescripting_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwprofile_crosssitescripting_bindingName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_crosssitescripting_binding state %s", appfwprofile_crosssitescripting_bindingName)
	data, err := client.FindResource(netscaler.Appfwprofile_crosssitescripting_binding.Type(), appfwprofile_crosssitescripting_bindingName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_crosssitescripting_binding state %s", appfwprofile_crosssitescripting_bindingName)
		d.SetId("")
		return nil
	}
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
	args := make(map[string]string)
	args["crosssitescripting"] = d.Get("crosssitescripting").(string)
	args["formactionurl_xss"] = url.QueryEscape(d.Get("formactionurl_xss").(string))
	args["as_scan_location_xss"] = d.Get("as_scan_location_xss").(string)
	err := client.DeleteResourceWithArgsMap(netscaler.Appfwprofile_crosssitescripting_binding.Type(), d.Get("name").(string), args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
