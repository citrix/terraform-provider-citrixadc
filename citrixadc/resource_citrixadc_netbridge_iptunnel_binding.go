package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcNetbridge_iptunnel_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNetbridge_iptunnel_bindingFunc,
		ReadContext:   readNetbridge_iptunnel_bindingFunc,
		DeleteContext: deleteNetbridge_iptunnel_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"tunnel": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNetbridge_iptunnel_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNetbridge_iptunnel_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	tunnel := d.Get("tunnel")
	bindingId := fmt.Sprintf("%s,%s", name, tunnel)
	netbridge_iptunnel_binding := network.Netbridgeiptunnelbinding{
		Name:   d.Get("name").(string),
		Tunnel: d.Get("tunnel").(string),
	}

	err := client.UpdateUnnamedResource(service.Netbridge_iptunnel_binding.Type(), &netbridge_iptunnel_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readNetbridge_iptunnel_bindingFunc(ctx, d, meta)
}

func readNetbridge_iptunnel_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNetbridge_iptunnel_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	tunnel := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading netbridge_iptunnel_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "netbridge_iptunnel_binding",
		ResourceName:             name,
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
		log.Printf("[WARN] citrixadc-provider: Clearing netbridge_iptunnel_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["tunnel"].(string) == tunnel {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing netbridge_iptunnel_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("tunnel", data["tunnel"])

	return nil

}

func deleteNetbridge_iptunnel_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNetbridge_iptunnel_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	tunnel := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("tunnel:%s", tunnel))

	err := client.DeleteResourceWithArgs(service.Netbridge_iptunnel_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
