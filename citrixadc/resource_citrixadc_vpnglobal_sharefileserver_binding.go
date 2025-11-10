package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVpnglobal_sharefileserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnglobal_sharefileserver_bindingFunc,
		ReadContext:   readVpnglobal_sharefileserver_bindingFunc,
		DeleteContext: deleteVpnglobal_sharefileserver_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"sharefile": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
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

func createVpnglobal_sharefileserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_sharefileserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	sharefile := d.Get("sharefile").(string)
	vpnglobal_sharefileserver_binding := vpn.Vpnglobalsharefileserverbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Sharefile:              d.Get("sharefile").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_sharefileserver_binding.Type(), &vpnglobal_sharefileserver_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sharefile)

	return readVpnglobal_sharefileserver_bindingFunc(ctx, d, meta)
}

func readVpnglobal_sharefileserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_sharefileserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	sharefile := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_sharefileserver_binding state %s", sharefile)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_sharefileserver_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_sharefileserver_binding state %s", sharefile)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["sharefile"].(string) == sharefile {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_sharefileserver_binding state %s", sharefile)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("sharefile", data["sharefile"])

	return nil

}

func deleteVpnglobal_sharefileserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_sharefileserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	sharefile := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("sharefile:%s", sharefile))

	err := client.DeleteResourceWithArgs(service.Vpnglobal_sharefileserver_binding.Type(), "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
