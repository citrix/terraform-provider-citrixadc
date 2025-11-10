package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func resourceCitrixAdcVlan_channel_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVlan_channel_bindingFunc,
		ReadContext:   readVlan_channel_bindingFunc,
		DeleteContext: deleteVlan_channel_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"vlanid": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ifnum": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tagged": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVlan_channel_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVlan_channel_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vlanid := strconv.Itoa(d.Get("vlanid").(int))
	ifnum := d.Get("ifnum").(string)
	bindingId := fmt.Sprintf("%s,%s", vlanid, ifnum)
	vlan_channel_binding := network.Vlanchannelbinding{
		Ifnum:      d.Get("ifnum").(string),
		Ownergroup: d.Get("ownergroup").(string),
		Tagged:     d.Get("tagged").(bool),
	}

	if raw := d.GetRawConfig().GetAttr("vlanid"); !raw.IsNull() {
		vlan_channel_binding.Id = intPtr(d.Get("vlanid").(int))
	}

	err := client.UpdateUnnamedResource(service.Vlan_channel_binding.Type(), &vlan_channel_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readVlan_channel_bindingFunc(ctx, d, meta)
}

func readVlan_channel_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVlan_channel_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vlanid := idSlice[0]
	ifnum := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vlan_channel_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vlan_channel_binding",
		ResourceName:             vlanid,
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
		log.Printf("[WARN] citrixadc-provider: Clearing vlan_channel_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["ifnum"].(string) == ifnum {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams ifnum not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vlan_channel_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	setToInt("vlanid", d, data["id"])
	d.Set("ifnum", data["ifnum"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("tagged", data["tagged"])

	return nil

}

func deleteVlan_channel_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVlan_channel_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vlan := idSlice[0]
	ifnum := idSlice[1]

	args := make([]string, 0)

	args = append(args, fmt.Sprintf("ifnum:%s", url.PathEscape(ifnum)))
	if v, ok := d.GetOk("tagged"); ok {
		tagged := v.(bool)
		args = append(args, fmt.Sprintf("tagged:%s", strconv.FormatBool(tagged)))
	}
	if v, ok := d.GetOk("ownergroup"); ok {
		ownergroup := v.(int)
		args = append(args, fmt.Sprintf("ownergroup:%s", strconv.Itoa(ownergroup)))
	}

	err := client.DeleteResourceWithArgs(service.Vlan_channel_binding.Type(), vlan, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
