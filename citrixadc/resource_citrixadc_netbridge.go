package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNetbridge() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNetbridgeFunc,
		ReadContext:   readNetbridgeFunc,
		UpdateContext: updateNetbridgeFunc,
		DeleteContext: deleteNetbridgeFunc,
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
			"vxlanvlanmap": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNetbridgeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNetbridgeFunc")
	client := meta.(*NetScalerNitroClient).client
	netbridgeName := d.Get("name").(string)
	netbridge := network.Netbridge{
		Name:         d.Get("name").(string),
		Vxlanvlanmap: d.Get("vxlanvlanmap").(string),
	}

	_, err := client.AddResource(service.Netbridge.Type(), netbridgeName, &netbridge)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(netbridgeName)

	return readNetbridgeFunc(ctx, d, meta)
}

func readNetbridgeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNetbridgeFunc")
	client := meta.(*NetScalerNitroClient).client
	netbridgeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading netbridge state %s", netbridgeName)
	data, err := client.FindResource(service.Netbridge.Type(), netbridgeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing netbridge state %s", netbridgeName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("vxlanvlanmap", data["vxlanvlanmap"])

	return nil

}

func updateNetbridgeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNetbridgeFunc")
	client := meta.(*NetScalerNitroClient).client
	netbridgeName := d.Get("name").(string)

	netbridge := network.Netbridge{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("vxlanvlanmap") {
		log.Printf("[DEBUG]  citrixadc-provider: Vxlanvlanmap has changed for netbridge %s, starting update", netbridgeName)
		netbridge.Vxlanvlanmap = d.Get("vxlanvlanmap").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Netbridge.Type(), netbridgeName, &netbridge)
		if err != nil {
			return diag.Errorf("Error updating netbridge %s", netbridgeName)
		}
	}
	return readNetbridgeFunc(ctx, d, meta)
}

func deleteNetbridgeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNetbridgeFunc")
	client := meta.(*NetScalerNitroClient).client
	netbridgeName := d.Id()
	err := client.DeleteResource(service.Netbridge.Type(), netbridgeName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
