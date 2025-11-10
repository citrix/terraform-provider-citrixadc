package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
)

func resourceCitrixAdcVpnglobal_appcontroller_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnglobal_appcontroller_bindingFunc,
		ReadContext:   readVpnglobal_appcontroller_bindingFunc,
		DeleteContext: deleteVpnglobal_appcontroller_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"appcontroller": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnglobal_appcontroller_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_appcontroller_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appcontroller := d.Get("appcontroller").(string)
	vpnglobal_appcontroller_binding := vpn.Vpnglobalappcontrollerbinding{
		Appcontroller:          d.Get("appcontroller").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
	}

	_, err := client.AddResource(service.Vpnglobal_appcontroller_binding.Type(), appcontroller, &vpnglobal_appcontroller_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appcontroller)

	return readVpnglobal_appcontroller_bindingFunc(ctx, d, meta)
}

func readVpnglobal_appcontroller_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_appcontroller_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appcontroller, _ := url.QueryUnescape(d.Id())

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_appcontroller_binding state %s", appcontroller)

	findParams := service.FindParams{
		ResourceType: "vpnglobal_appcontroller_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_appcontroller_binding state %s", appcontroller)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["appcontroller"].(string) == appcontroller {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_appcontroller_binding state %s", appcontroller)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("appcontroller", data["appcontroller"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])

	return nil

}

func deleteVpnglobal_appcontroller_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_appcontroller_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	appcontroller := d.Id()

	argsMap := make(map[string]string)
	argsMap["appcontroller"] = appcontroller

	err := client.DeleteResourceWithArgsMap(service.Vpnglobal_appcontroller_binding.Type(), "", argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
