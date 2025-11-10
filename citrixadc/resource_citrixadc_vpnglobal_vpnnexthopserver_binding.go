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

func resourceCitrixAdcVpnglobal_vpnnexthopserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnglobal_vpnnexthopserver_bindingFunc,
		ReadContext:   readVpnglobal_vpnnexthopserver_bindingFunc,
		DeleteContext: deleteVpnglobal_vpnnexthopserver_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"nexthopserver": {
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

func createVpnglobal_vpnnexthopserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_vpnnexthopserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	nexthopserver := d.Get("nexthopserver").(string)
	vpnglobal_vpnnexthopserver_binding := vpn.Vpnglobalvpnnexthopserverbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Nexthopserver:          d.Get("nexthopserver").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_vpnnexthopserver_binding.Type(), &vpnglobal_vpnnexthopserver_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nexthopserver)

	return readVpnglobal_vpnnexthopserver_bindingFunc(ctx, d, meta)
}

func readVpnglobal_vpnnexthopserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_vpnnexthopserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	nexthopserver := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_vpnnexthopserver_binding state %s", nexthopserver)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_vpnnexthopserver_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnnexthopserver_binding state %s", nexthopserver)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["nexthopserver"].(string) == nexthopserver {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnnexthopserver_binding state %s", nexthopserver)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("nexthopserver", data["nexthopserver"])

	return nil

}

func deleteVpnglobal_vpnnexthopserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_vpnnexthopserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	nexthopserver := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("nexthopserver:%s", nexthopserver))

	err := client.DeleteResourceWithArgs(service.Vpnglobal_vpnnexthopserver_binding.Type(), "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
