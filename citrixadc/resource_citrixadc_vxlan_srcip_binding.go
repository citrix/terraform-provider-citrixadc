package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcVxlan_srcip_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVxlan_srcip_bindingFunc,
		ReadContext:   readVxlan_srcip_bindingFunc,
		DeleteContext: deleteVxlan_srcip_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"vxlanid": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"srcip": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createVxlan_srcip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVxlan_srcip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanid := strconv.Itoa(d.Get("vxlanid").(int))
	srcip := d.Get("srcip").(string)
	bindingId := fmt.Sprintf("%s,%s", vxlanid, srcip)
	vxlan_srcip_binding := network.Vxlansrcipbinding{
		Srcip: d.Get("srcip").(string),
	}

	if raw := d.GetRawConfig().GetAttr("vxlanid"); !raw.IsNull() {
		vxlan_srcip_binding.Id = intPtr(d.Get("vxlanid").(int))
	}

	err := client.UpdateUnnamedResource("vxlan_srcip_binding", &vxlan_srcip_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readVxlan_srcip_bindingFunc(ctx, d, meta)
}

func readVxlan_srcip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVxlan_srcip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vxlanid := idSlice[0]
	srcip := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vxlan_srcip_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vxlan_srcip_binding",
		ResourceName:             vxlanid,
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
		log.Printf("[WARN] citrixadc-provider: Clearing vxlan_srcip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["srcip"].(string) == srcip {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vxlan_srcip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	setToInt("vxlanid", d, data["id"])
	d.Set("srcip", data["srcip"])

	return nil

}

func deleteVxlan_srcip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVxlan_srcip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vxlanid := idSlice[0]
	srcip := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("srcip:%s", srcip))

	err := client.DeleteResourceWithArgs("vxlan_srcip_binding", vxlanid, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
