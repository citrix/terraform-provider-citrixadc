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

func resourceCitrixAdcAppfwprofile_csrftag_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwprofile_csrftag_bindingFunc,
		ReadContext:   readAppfwprofile_csrftag_bindingFunc,
		DeleteContext: deleteAppfwprofile_csrftag_bindingFunc,
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
			"csrftag": {
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
			"csrfformactionurl": {
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func createAppfwprofile_csrftag_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_csrftag_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appFwName := d.Get("name").(string)
	csrftag := d.Get("csrftag").(string)
	csrfformactionurl := d.Get("csrfformactionurl").(string)
	bindingId := fmt.Sprintf("%s,%s,%s", appFwName, csrftag, csrfformactionurl)
	appfwprofile_csrftag_binding := appfw.Appfwprofilecsrftagbinding{
		Alertonly:         d.Get("alertonly").(string),
		Comment:           d.Get("comment").(string),
		Csrfformactionurl: d.Get("csrfformactionurl").(string),
		Csrftag:           d.Get("csrftag").(string),
		Isautodeployed:    d.Get("isautodeployed").(string),
		Name:              appFwName,
		Resourceid:        d.Get("resourceid").(string),
		Ruletype:          d.Get("ruletype").(string),
		State:             d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_csrftag_binding.Type(), &appfwprofile_csrftag_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAppfwprofile_csrftag_bindingFunc(ctx, d, meta)
}

func readAppfwprofile_csrftag_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_csrftag_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: readAppfwprofile_csrftag_bindingFunc: bindingId: %s", bindingId)
	idSlice := strings.Split(bindingId, ",")
	appFwName := idSlice[0]
	csrftag := idSlice[1]
	csrfformactionurl := ""
	if len(idSlice) > 2 {
		csrfformactionurl = idSlice[2]
	} else {
		csrfformactionurl = d.Get("csrfformactionurl").(string)
		bindingId = fmt.Sprintf("%s,%s", bindingId, csrfformactionurl)
	}
	d.SetId(bindingId)
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_csrftag_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_csrftag_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_csrftag_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right binding
	foundIndex := -1
	for i, v := range dataArr {
		if v["csrftag"].(string) == csrftag && v["csrfformactionurl"].(string) == csrfformactionurl {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams appfwprofile_csrftag_binding not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_csrftag_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	d.Set("csrfformactionurl", data["csrfformactionurl"])
	d.Set("csrftag", data["csrftag"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_csrftag_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_csrftag_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")
	appFwName := idSlice[0]
	csrftag := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("csrftag:%s", url.QueryEscape(csrftag)))
	if val, ok := d.GetOk("csrfformactionurl"); ok {
		args = append(args, fmt.Sprintf("csrfformactionurl:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Appfwprofile_csrftag_binding.Type(), appFwName, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
