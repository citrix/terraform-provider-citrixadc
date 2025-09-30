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

func resourceCitrixAdcAppfwprofile_xmlsqlinjection_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_xmlsqlinjection_bindingFunc,
		Read:          readAppfwprofile_xmlsqlinjection_bindingFunc,
		Delete:        deleteAppfwprofile_xmlsqlinjection_bindingFunc,
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
			"xmlsqlinjection": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"as_scan_location_xmlsql": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "ELEMENT",
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
			"isregex_xmlsql": {
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

func createAppfwprofile_xmlsqlinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_xmlsqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	xmlsqlinjection := d.Get("xmlsqlinjection")
	as_scan_location_xmlsql := d.Get("as_scan_location_xmlsql")
	bindingId := fmt.Sprintf("%s,%s,%s", name, xmlsqlinjection, as_scan_location_xmlsql)
	appfwprofile_xmlsqlinjection_binding := appfw.Appfwprofilexmlsqlinjectionbinding{
		Alertonly:            d.Get("alertonly").(string),
		Asscanlocationxmlsql: d.Get("as_scan_location_xmlsql").(string),
		Comment:              d.Get("comment").(string),
		Isautodeployed:       d.Get("isautodeployed").(string),
		Isregexxmlsql:        d.Get("isregex_xmlsql").(string),
		Name:                 d.Get("name").(string),
		Resourceid:           d.Get("resourceid").(string),
		Ruletype:             d.Get("ruletype").(string),
		State:                d.Get("state").(string),
		Xmlsqlinjection:      d.Get("xmlsqlinjection").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_xmlsqlinjection_binding.Type(), &appfwprofile_xmlsqlinjection_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_xmlsqlinjection_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_xmlsqlinjection_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_xmlsqlinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_xmlsqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	name := idSlice[0]
	xmlsqlinjection := idSlice[1]
	if len(idSlice) == 2 {
		as_scan_location_xmlsql := d.Get("as_scan_location_xmlsql").(string)
		bindingId = fmt.Sprintf("%s,%s,%s", name, xmlsqlinjection, as_scan_location_xmlsql)
	}

	d.SetId(bindingId)
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_xmlsqlinjection_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_xmlsqlinjection_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_xmlsqlinjection_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["xmlsqlinjection"].(string) == xmlsqlinjection {
			if v["as_scan_location_xmlsql"].(string) == d.Get("as_scan_location_xmlsql").(string) {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_xmlsqlinjection_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("as_scan_location_xmlsql", data["as_scan_location_xmlsql"])
	d.Set("comment", data["comment"])
	// d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregex_xmlsql", data["isregex_xmlsql"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])
	d.Set("xmlsqlinjection", data["xmlsqlinjection"])

	return nil

}

func deleteAppfwprofile_xmlsqlinjection_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_xmlsqlinjection_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	name := idSlice[0]
	xmlsqlinjection := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("xmlsqlinjection:%s", xmlsqlinjection))
	if val, ok := d.GetOk("as_scan_location_xmlsql"); ok {
		args = append(args, fmt.Sprintf("as_scan_location_xmlsql:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Appfwprofile_xmlsqlinjection_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
