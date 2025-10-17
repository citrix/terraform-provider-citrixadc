package citrixadc

import (
	"context"
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcAppfwprofileDenyurlBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwprofileDenyurlBindingFunc,
		ReadContext:   readAppfwprofileDenyurlBindingFunc,
		DeleteContext: deleteAppfwprofileDenyurlBindingFunc,
		Schema: map[string]*schema.Schema{
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
			"denyurl": {
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
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
		},
	}
}

func createAppfwprofileDenyurlBindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofileDenyurlBindingFunc")
	client := meta.(*NetScalerNitroClient).client

	profileName := d.Get("name")
	denyURL := d.Get("denyurl")

	// Use `,` as the separator since it is invalid character for adc entity strings
	bindingID := fmt.Sprintf("%s,%s", profileName, denyURL)

	appfwprofileDenyurlBinding := appfw.Appfwprofiledenyurlbinding{
		Alertonly:      d.Get("alertonly").(string),
		Comment:        d.Get("comment").(string),
		Denyurl:        d.Get("denyurl").(string),
		Isautodeployed: d.Get("isautodeployed").(string),
		Name:           d.Get("name").(string),
		State:          d.Get("state").(string),
		Resourceid:     d.Get("resourceid").(string),
		Ruletype:       d.Get("ruletype").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_denyurl_binding.Type(), &appfwprofileDenyurlBinding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingID)

	return readAppfwprofileDenyurlBindingFunc(ctx, d, meta)
}

func readAppfwprofileDenyurlBindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofileDenyurlBindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingID := d.Id()
	idSlice := strings.SplitN(bindingID, ",", 2)

	if len(idSlice) < 2 {
		return diag.Errorf("Cannot deduce appfwprofile and denyurl from ID string")
	}

	profileName := idSlice[0]
	denyURL := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofileDenyurlBinding state %s", bindingID)

	findParams := service.FindParams{
		ResourceType: service.Appfwprofile_denyurl_binding.Type(),
		ResourceName: profileName,
	}
	findParams.FilterMap = make(map[string]string)
	findParams.FilterMap["denyurl"] = url.QueryEscape(denyURL)
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return diag.FromErr(err)
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_denyurl_binding state %s", bindingID)
		d.SetId("")
		return nil
	}

	data := dataArr[0]

	d.Set("name", data["name"])
	d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	d.Set("denyurl", data["denyurl"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("name", data["name"])
	d.Set("state", data["state"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])

	return nil

}

func deleteAppfwprofileDenyurlBindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofileDenyurlBindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingID := d.Id()
	idSlice := strings.SplitN(bindingID, ",", 2)

	if len(idSlice) < 2 {
		return diag.Errorf("Cannot deduce appfwprofile and denyurl from ID string")
	}

	profileName := idSlice[0]
	denyURL := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("denyurl:%v", url.QueryEscape(denyURL)))

	err := client.DeleteResourceWithArgs(service.Appfwprofile_denyurl_binding.Type(), profileName, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
