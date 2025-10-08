package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
)

func resourceCitrixAdcVpnglobal_intranetip_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnglobal_intranetip_bindingFunc,
		ReadContext:   readVpnglobal_intranetip_bindingFunc,
		DeleteContext: deleteVpnglobal_intranetip_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"intranetip": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"netmask": {
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

func createVpnglobal_intranetip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	intranetip := d.Get("intranetip").(string)
	vpnglobal_intranetip_binding := vpn.Vpnglobalintranetipbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Intranetip:             d.Get("intranetip").(string),
		Netmask:                d.Get("netmask").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_intranetip_binding.Type(), &vpnglobal_intranetip_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(intranetip)

	return readVpnglobal_intranetip_bindingFunc(ctx, d, meta)
}

func readVpnglobal_intranetip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	intranetip := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_intranetip_binding state %s", intranetip)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_intranetip_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_intranetip_binding state %s", intranetip)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetip"].(string) == intranetip {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_intranetip_binding state %s", intranetip)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("intranetip", data["intranetip"])
	d.Set("netmask", data["netmask"])

	return nil

}

func deleteVpnglobal_intranetip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	intranetip := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetip:%s", intranetip))
	if val, ok := d.GetOk("netmask"); ok {
		args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Vpnglobal_intranetip_binding.Type(), "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
