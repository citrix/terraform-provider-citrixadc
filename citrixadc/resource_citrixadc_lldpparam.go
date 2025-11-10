package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lldp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLldpparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLldpparamFunc,
		ReadContext:   readLldpparamFunc,
		UpdateContext: updateLldpparamFunc,
		DeleteContext: deleteLldpparamFunc,
		Schema: map[string]*schema.Schema{
			"holdtimetxmult": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timer": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLldpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLldpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	lldpparamName := resource.PrefixedUniqueId("tf-lldpparam-")

	lldpparam := lldp.Lldpparam{
		Mode: d.Get("mode").(string),
	}

	if raw := d.GetRawConfig().GetAttr("holdtimetxmult"); !raw.IsNull() {
		lldpparam.Holdtimetxmult = intPtr(d.Get("holdtimetxmult").(int))
	}
	if raw := d.GetRawConfig().GetAttr("timer"); !raw.IsNull() {
		lldpparam.Timer = intPtr(d.Get("timer").(int))
	}

	err := client.UpdateUnnamedResource("lldpparam", &lldpparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lldpparamName)

	return readLldpparamFunc(ctx, d, meta)
}

func readLldpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLldpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading lldpparam state")
	data, err := client.FindResource("lldpparam", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lldpparam state")
		d.SetId("")
		return nil
	}
	setToInt("holdtimetxmult", d, data["holdtimetxmult"])
	d.Set("mode", data["mode"])
	setToInt("timer", d, data["timer"])

	return nil

}

func updateLldpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLldpparamFunc")
	client := meta.(*NetScalerNitroClient).client

	lldpparam := lldp.Lldpparam{}
	hasChange := false
	if d.HasChange("holdtimetxmult") {
		log.Printf("[DEBUG]  citrixadc-provider: Holdtimetxmult has changed for lldpparam, starting update")
		lldpparam.Holdtimetxmult = intPtr(d.Get("holdtimetxmult").(int))
		hasChange = true
	}
	if d.HasChange("mode") {
		log.Printf("[DEBUG]  citrixadc-provider: Mode has changed for lldpparam, starting update")
		lldpparam.Mode = d.Get("mode").(string)
		hasChange = true
	}
	if d.HasChange("timer") {
		log.Printf("[DEBUG]  citrixadc-provider: Timer has changed for lldpparam, starting update")
		lldpparam.Timer = intPtr(d.Get("timer").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lldpparam", &lldpparam)
		if err != nil {
			return diag.Errorf("Error updating lldpparam")
		}
	}
	return readLldpparamFunc(ctx, d, meta)
}

func deleteLldpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLldpparamFunc")
	// lldpparam does not support DELETE operation
	d.SetId("")

	return nil
}
