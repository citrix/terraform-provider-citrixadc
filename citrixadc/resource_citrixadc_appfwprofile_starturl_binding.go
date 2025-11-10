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

func resourceCitrixAdcAppfwprofileStarturlBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwprofileStarturlBindingFunc,
		ReadContext:   readAppfwprofileStarturlBindingFunc,
		DeleteContext: deleteAppfwprofileStarturlBindingFunc,
		Schema: map[string]*schema.Schema{
			"alertonly": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"starturl": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"isautodeployed": {
				Type:     schema.TypeString,
				Optional: true,
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
				ForceNew: true,
			},
			"ruletype": {
				Type:     schema.TypeString,
				Optional: true,
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

func createAppfwprofileStarturlBindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofileStarturlBindingFunc")
	client := meta.(*NetScalerNitroClient).client

	profileName := d.Get("name")
	startURL := d.Get("starturl")

	// Use `,` as the separator since it is invalid character for adc entity strings
	bindingID := fmt.Sprintf("%s,%s", profileName, startURL)

	appfwprofileStarturlBinding := appfw.Appfwprofilestarturlbinding{
		Alertonly:      d.Get("alertonly").(string),
		Comment:        d.Get("comment").(string),
		Starturl:       d.Get("starturl").(string),
		Isautodeployed: d.Get("isautodeployed").(string),
		Name:           d.Get("name").(string),
		State:          d.Get("state").(string),
		Ruletype:       d.Get("ruletype").(string),
		Resourceid:     d.Get("resourceid").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_starturl_binding.Type(), &appfwprofileStarturlBinding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingID)

	return readAppfwprofileStarturlBindingFunc(ctx, d, meta)
}

func readAppfwprofileStarturlBindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofileStarturlBindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingID := d.Id()
	idSlice := strings.SplitN(bindingID, ",", 2)

	if len(idSlice) < 2 {
		return diag.Errorf("Cannot deduce appfwprofile and starturl from ID string")
	}

	profileName := idSlice[0]
	startURL := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofileStarturlBinding state %s", bindingID)

	findParams := service.FindParams{
		ResourceType: service.Appfwprofile_starturl_binding.Type(),
		ResourceName: profileName,
	}
	findParams.FilterMap = make(map[string]string)
	findParams.FilterMap["starturl"] = url.QueryEscape(startURL)
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return diag.FromErr(err)
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_starturl_binding state %s", bindingID)
		d.SetId("")
		return nil
	}

	data := dataArr[0]

	d.Set("name", data["name"])
	d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	d.Set("starturl", data["starturl"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("name", data["name"])
	d.Set("state", data["state"])
	d.Set("ruletype", data["ruletype"])
	d.Set("resourceid", data["resourceid"])

	return nil

}

func deleteAppfwprofileStarturlBindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofileStarturlBindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingID := d.Id()
	idSlice := strings.SplitN(bindingID, ",", 2)

	if len(idSlice) < 2 {
		return diag.Errorf("Cannot deduce appfwprofile and starturl from ID string")
	}

	profileName := idSlice[0]
	startURL := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("starturl:%v", url.QueryEscape(startURL)))

	err := client.DeleteResourceWithArgs(service.Appfwprofile_starturl_binding.Type(), profileName, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
