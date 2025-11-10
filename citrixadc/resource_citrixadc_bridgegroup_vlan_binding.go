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

func resourceCitrixAdcBridgegroup_vlan_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createBridgegroup_vlan_bindingFunc,
		ReadContext:   readBridgegroup_vlan_bindingFunc,
		DeleteContext: deleteBridgegroup_vlan_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"bridgegroup_id": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createBridgegroup_vlan_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createBridgegroup_vlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgegroup_id := strconv.Itoa(d.Get("bridgegroup_id").(int))
	vlan := strconv.Itoa(d.Get("vlan").(int))
	bindingId := fmt.Sprintf("%s,%s", bridgegroup_id, vlan)
	bridgegroup_vlan_binding := network.Bridgegroupvlanbinding{}

	if raw := d.GetRawConfig().GetAttr("bridgegroup_id"); !raw.IsNull() {
		bridgegroup_vlan_binding.Id = intPtr(d.Get("bridgegroup_id").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vlan"); !raw.IsNull() {
		bridgegroup_vlan_binding.Vlan = intPtr(d.Get("vlan").(int))
	}

	err := client.UpdateUnnamedResource(service.Bridgegroup_vlan_binding.Type(), &bridgegroup_vlan_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readBridgegroup_vlan_bindingFunc(ctx, d, meta)
}

func readBridgegroup_vlan_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readBridgegroup_vlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	bridgegroup_id := idSlice[0]
	vlan := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading bridgegroup_vlan_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "bridgegroup_vlan_binding",
		ResourceName:             bridgegroup_id,
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
		log.Printf("[WARN] citrixadc-provider: Clearing bridgegroup_vlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["vlan"].(string) == vlan {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing bridgegroup_vlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	setToInt("bridgegroup_id", d, data["id"])
	setToInt("vlan", d, data["vlan"])

	return nil

}

func deleteBridgegroup_vlan_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBridgegroup_vlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	bridgegroup_id := idSlice[0]
	vlan := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("vlan:%s", vlan))

	err := client.DeleteResourceWithArgs(service.Bridgegroup_vlan_binding.Type(), bridgegroup_id, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
