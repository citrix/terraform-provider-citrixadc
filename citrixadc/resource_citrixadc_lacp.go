package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLacp() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLacpFunc,
		ReadContext:   readLacpFunc,
		UpdateContext: updateLacpFunc,
		DeleteContext: deleteLacpFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"syspriority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ownernode": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  255,
			},
		},
	}
}

func createLacpFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLacpFunc")
	client := meta.(*NetScalerNitroClient).client
	lacpId := strconv.Itoa(d.Get("ownernode").(int))

	lacp := network.Lacp{}

	if raw := d.GetRawConfig().GetAttr("syspriority"); !raw.IsNull() {
		lacp.Syspriority = intPtr(d.Get("syspriority").(int))
	}
	if _, ok := d.GetOk("ownernode"); ok {
		lacp.Ownernode = intPtr(d.Get("ownernode").(int))
	}

	err := client.UpdateUnnamedResource(service.Lacp.Type(), &lacp)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lacpId)

	return readLacpFunc(ctx, d, meta)
}

func readLacpFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLacpFunc")
	client := meta.(*NetScalerNitroClient).client
	lacpId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lacp state")
	data, err := client.FindResource(service.Lacp.Type(), strconv.Itoa(d.Get("ownernode").(int)))
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lacp state %s", lacpId)
		d.SetId("")
		return nil
	}

	setToInt("ownernode", d, data["ownernode"])
	setToInt("syspriority", d, data["syspriority"])

	return nil

}

func updateLacpFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLacpFunc")
	client := meta.(*NetScalerNitroClient).client
	lacpId := d.Id()

	lacp := network.Lacp{}

	if raw := d.GetRawConfig().GetAttr("ownernode"); !raw.IsNull() {
		lacp.Ownernode = intPtr(d.Get("ownernode").(int))
	}
	hasChange := false

	if d.HasChange("syspriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Syspriority has changed for lacp, starting update %s", lacpId)
		lacp.Syspriority = intPtr(d.Get("syspriority").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Lacp.Type(), &lacp)
		if err != nil {
			return diag.Errorf("Error updating lacp %s", lacpId)
		}
	}
	return readLacpFunc(ctx, d, meta)
}

func deleteLacpFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLacpFunc")
	//lacp does not support delete operation
	d.SetId("")

	return nil
}
