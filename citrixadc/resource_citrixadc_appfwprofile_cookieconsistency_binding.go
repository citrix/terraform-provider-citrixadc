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

func resourceCitrixAdcAppfwprofile_cookieconsistency_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwprofile_cookieconsistency_bindingFunc,
		ReadContext:   readAppfwprofile_cookieconsistency_bindingFunc,
		DeleteContext: deleteAppfwprofile_cookieconsistency_bindingFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cookieconsistency": {
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
			"isregex": {
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

func createAppfwprofile_cookieconsistency_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_cookieconsistency_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appFwName := d.Get("name").(string)
	cookieconsistency := d.Get("cookieconsistency").(string)
	bindingId := fmt.Sprintf("%s,%s", appFwName, cookieconsistency)
	appfwprofile_cookieconsistency_binding := appfw.Appfwprofilecookieconsistencybinding{
		Alertonly:         d.Get("alertonly").(string),
		Comment:           d.Get("comment").(string),
		Cookieconsistency: cookieconsistency,
		Isautodeployed:    d.Get("isautodeployed").(string),
		Isregex:           d.Get("isregex").(string),
		Name:              appFwName,
		State:             d.Get("state").(string),
		Resourceid:        d.Get("resourceid").(string),
		Ruletype:          d.Get("ruletype").(string),
	}

	_, err := client.AddResource(service.Appfwprofile_cookieconsistency_binding.Type(), appFwName, &appfwprofile_cookieconsistency_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAppfwprofile_cookieconsistency_bindingFunc(ctx, d, meta)
}
func readAppfwprofile_cookieconsistency_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_cookieconsistency_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: readAppfwprofile_cookieconsistency_bindingFunc: bindingId: %s", bindingId)
	idSlice := strings.SplitN(bindingId, ",", 2)
	appFwName := idSlice[0]
	cookieconsistency := idSlice[1]
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_cookieconsistency_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_cookieconsistency_binding.Type(),
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_cookieconsistency_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if v["cookieconsistency"].(string) == cookieconsistency {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_cookieconsistency_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough
	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	d.Set("cookieconsistency", data["cookieconsistency"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregex", data["isregex"])
	d.Set("name", data["name"])
	d.Set("state", data["state"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])

	return nil

}

func deleteAppfwprofile_cookieconsistency_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_cookieconsistency_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	appFwName := idSlice[0]
	cookieConsistencyString := idSlice[1]

	args := make(map[string]string)
	args["cookieconsistency"] = url.QueryEscape(cookieConsistencyString)

	err := client.DeleteResourceWithArgsMap(service.Appfwprofile_cookieconsistency_binding.Type(), appFwName, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
