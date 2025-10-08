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

func resourceCitrixAdcAppfwprofile_xmlxss_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwprofile_xmlxss_bindingFunc,
		ReadContext:   readAppfwprofile_xmlxss_bindingFunc,
		DeleteContext: deleteAppfwprofile_xmlxss_bindingFunc,
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
			"xmlxss": {
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
			"as_scan_location_xmlxss": {
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
			"isregex_xmlxss": {
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

func createAppfwprofile_xmlxss_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_xmlxss_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	xmlxss := d.Get("xmlxss")
	as_scan_location_xmlxss := d.Get("as_scan_location_xmlxss")
	bindingId := fmt.Sprintf("%s,%s,%s", name, xmlxss, as_scan_location_xmlxss)
	appfwprofile_xmlxss_binding := appfw.Appfwprofilexmlxssbinding{
		Alertonly:            d.Get("alertonly").(string),
		Asscanlocationxmlxss: d.Get("as_scan_location_xmlxss").(string),
		Comment:              d.Get("comment").(string),
		Isautodeployed:       d.Get("isautodeployed").(string),
		Isregexxmlxss:        d.Get("isregex_xmlxss").(string),
		Name:                 d.Get("name").(string),
		Resourceid:           d.Get("resourceid").(string),
		Ruletype:             d.Get("ruletype").(string),
		State:                d.Get("state").(string),
		Xmlxss:               d.Get("xmlxss").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_xmlxss_binding.Type(), &appfwprofile_xmlxss_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAppfwprofile_xmlxss_bindingFunc(ctx, d, meta)
}

func readAppfwprofile_xmlxss_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_xmlxss_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	xmlxss := idSlice[1]
	as_scan_location_xmlxss := idSlice[2]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_xmlxss_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_xmlxss_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_xmlxss_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true
		if v["xmlxss"].(string) != xmlxss {
			match = false
		}
		if v["as_scan_location_xmlxss"].(string) != as_scan_location_xmlxss {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_xmlxss_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("as_scan_location_xmlxss", data["as_scan_location_xmlxss"])
	d.Set("comment", data["comment"])
	// d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregex_xmlxss", data["isregex_xmlxss"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])
	d.Set("xmlxss", data["xmlxss"])

	return nil

}

func deleteAppfwprofile_xmlxss_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_xmlxss_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	xmlxss := idSlice[1]
	as_scan_location_xmlxss := idSlice[2]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("xmlxss:%s", xmlxss))
	args = append(args, fmt.Sprintf("as_scan_location_xmlxss:%s", as_scan_location_xmlxss))
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Appfwprofile_xmlxss_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
