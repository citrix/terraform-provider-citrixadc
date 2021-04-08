package citrixadc

import (
	"net/url"

	"github.com/chiradeep/go-nitro/config/appfw"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwprofile_sqlinjection_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_sqlinjection_bindingFunc,
		Read:          readAppfwprofile_sqlinjection_bindingFunc,
		Delete:        deleteAppfwprofile_sqlinjection_bindingFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sqlinjection": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"as_scan_location_sql": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"formactionurl_sql": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alertonly": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"as_value_expr_sql": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"as_value_type_sql": &schema.Schema{
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
			"isregex_sql": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isvalueregex_sql": &schema.Schema{
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

func createAppfwprofile_sqlinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_sqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwprofile_sqlinjection_bindingName := d.Get("name").(string)
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
		Name:              d.Get("name").(string),
		Sqlinjection:      d.Get("sqlinjection").(string),
		State:             d.Get("state").(string),
	}

	_, err := client.AddResource(netscaler.Appfwprofile_sqlinjection_binding.Type(), appfwprofile_sqlinjection_bindingName, &appfwprofile_sqlinjection_binding)
	if err != nil {
		return err
	}

	d.SetId(appfwprofile_sqlinjection_bindingName)

	err = readAppfwprofile_sqlinjection_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_sqlinjection_binding but we can't read it ?? %s", appfwprofile_sqlinjection_bindingName)
		return nil
	}
	return nil
}

func readAppfwprofile_sqlinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_sqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwprofile_sqlinjection_bindingName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_sqlinjection_binding state %s", appfwprofile_sqlinjection_bindingName)
	data, err := client.FindResource(netscaler.Appfwprofile_sqlinjection_binding.Type(), appfwprofile_sqlinjection_bindingName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_sqlinjection_binding state %s", appfwprofile_sqlinjection_bindingName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("alertonly", data["alertonly"])
	d.Set("as_scan_location_sql", data["as_scan_location_sql"])
	d.Set("as_value_expr_sql", data["as_value_expr_sql"])
	d.Set("as_value_type_sql", data["as_value_type_sql"])
	d.Set("comment", data["comment"])
	d.Set("formactionurl_sql", data["formactionurl_sql"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregex_sql", data["isregex_sql"])
	d.Set("isvalueregex_sql", data["isvalueregex_sql"])
	d.Set("name", data["name"])
	d.Set("sqlinjection", data["sqlinjection"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_sqlinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_sqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	args := make(map[string]string)
	args["sqlinjection"] = d.Get("sqlinjection").(string)
	args["formactionurl_sql"] = url.QueryEscape(d.Get("formactionurl_sql").(string))
	args["as_scan_location_sql"] = d.Get("as_scan_location_sql").(string)
	err := client.DeleteResourceWithArgsMap(netscaler.Appfwprofile_sqlinjection_binding.Type(), d.Get("name").(string), args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
