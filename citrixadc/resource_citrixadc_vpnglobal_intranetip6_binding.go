package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"
	//"net/url"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcVpnglobal_intranetip6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnglobal_intranetip6_bindingFunc,
		ReadContext:   readVpnglobal_intranetip6_bindingFunc,
		DeleteContext: deleteVpnglobal_intranetip6_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"intranetip6": {
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
			"numaddr": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnglobal_intranetip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	intranetip6 := d.Get("intranetip6").(string)

	vpnglobal_intranetip6_binding := vpn.Vpnglobalintranetip6binding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Intranetip6:            d.Get("intranetip6").(string),
	}

	if raw := d.GetRawConfig().GetAttr("numaddr"); !raw.IsNull() {
		vpnglobal_intranetip6_binding.Numaddr = intPtr(d.Get("numaddr").(int))
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_intranetip6_binding.Type(), &vpnglobal_intranetip6_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(intranetip6)

	return readVpnglobal_intranetip6_bindingFunc(ctx, d, meta)
}

func readVpnglobal_intranetip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	intranetip6 := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_intranetip6_binding state %s", intranetip6)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_intranetip6_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_intranetip6_binding state %s", intranetip6)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetip6"].(string) == intranetip6 {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_intranetip6_binding state %s", intranetip6)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("intranetip6", data["intranetip6"])
	setToInt("numaddr", d, data["numaddr"])

	return nil

}

func deleteVpnglobal_intranetip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	intranetip6 := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetip6:%s", intranetip6))
	if val, ok := d.GetOk("numaddr"); ok {
		args = append(args, fmt.Sprintf("numaddr:%d", (val.(int))))
	}

	err := client.DeleteResourceWithArgs(service.Vpnglobal_intranetip6_binding.Type(), "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
