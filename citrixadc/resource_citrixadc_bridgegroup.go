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

func resourceCitrixAdcBridgegroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createBridgegroupFunc,
		ReadContext:   readBridgegroupFunc,
		UpdateContext: updateBridgegroupFunc,
		DeleteContext: deleteBridgegroupFunc,
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
			"dynamicrouting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6dynamicrouting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createBridgegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createBridgegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgegroup_Id := d.Get("bridgegroup_id").(int)
	bridgegroup := network.Bridgegroup{
		Dynamicrouting:     d.Get("dynamicrouting").(string),
		Ipv6dynamicrouting: d.Get("ipv6dynamicrouting").(string),
	}

	if raw := d.GetRawConfig().GetAttr("bridgegroup_id"); !raw.IsNull() {
		bridgegroup.Id = intPtr(d.Get("bridgegroup_id").(int))
	}
	bridgegroup_IdStr := strconv.Itoa(bridgegroup_Id)
	_, err := client.AddResource(service.Bridgegroup.Type(), bridgegroup_IdStr, &bridgegroup)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bridgegroup_IdStr)

	return readBridgegroupFunc(ctx, d, meta)
}

func readBridgegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readBridgegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgegroup_IdStr := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading bridgegroup state %s", bridgegroup_IdStr)
	data, err := client.FindResource(service.Bridgegroup.Type(), bridgegroup_IdStr)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing bridgegroup state %s", bridgegroup_IdStr)
		d.SetId("")
		return nil
	}
	d.Set("dynamicrouting", data["dynamicrouting"])
	setToInt("bridgegroup_id", d, data["id"])
	d.Set("ipv6dynamicrouting", data["ipv6dynamicrouting"])

	return nil

}

func updateBridgegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateBridgegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgegroup_Id := d.Get("bridgegroup_id").(int)

	bridgegroup := network.Bridgegroup{}

	if raw := d.GetRawConfig().GetAttr("bridgegroup_id"); !raw.IsNull() {
		bridgegroup.Id = intPtr(d.Get("bridgegroup_id").(int))
	}
	hasChange := false
	if d.HasChange("dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamicrouting has changed for bridgegroup %d, starting update", bridgegroup_Id)
		bridgegroup.Dynamicrouting = d.Get("dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("ipv6dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv6dynamicrouting has changed for bridgegroup %d, starting update", bridgegroup_Id)
		bridgegroup.Ipv6dynamicrouting = d.Get("ipv6dynamicrouting").(string)
		hasChange = true
	}
	bridgegroup_IdStr := strconv.Itoa(bridgegroup_Id)
	if hasChange {
		_, err := client.UpdateResource(service.Bridgegroup.Type(), bridgegroup_IdStr, &bridgegroup)
		if err != nil {
			return diag.Errorf("Error updating bridgegroup %s", bridgegroup_IdStr)
		}
	}
	return readBridgegroupFunc(ctx, d, meta)
}

func deleteBridgegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBridgegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgegroupName := d.Id()
	err := client.DeleteResource(service.Bridgegroup.Type(), bridgegroupName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
