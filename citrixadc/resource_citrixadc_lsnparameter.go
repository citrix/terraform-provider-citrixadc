package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLsnparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnparameterFunc,
		ReadContext:   readLsnparameterFunc,
		UpdateContext: updateLsnparameterFunc,
		DeleteContext: deleteLsnparameterFunc,
		Schema: map[string]*schema.Schema{
			"sessionsync": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscrsessionremoval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsnparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnparameterName := resource.PrefixedUniqueId("tf-lsnparameter-")

	lsnparameter := lsn.Lsnparameter{
		Sessionsync:          d.Get("sessionsync").(string),
		Subscrsessionremoval: d.Get("subscrsessionremoval").(string),
	}

	if raw := d.GetRawConfig().GetAttr("memlimit"); !raw.IsNull() {
		lsnparameter.Memlimit = intPtr(d.Get("memlimit").(int))
	}

	err := client.UpdateUnnamedResource("lsnparameter", &lsnparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lsnparameterName)

	return readLsnparameterFunc(ctx, d, meta)
}

func readLsnparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnparameter state")
	data, err := client.FindResource("lsnparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnparameter state")
		d.SetId("")
		return nil
	}
	d.Set("memlimit", data["memlimit"])
	d.Set("sessionsync", data["sessionsync"])
	d.Set("subscrsessionremoval", data["subscrsessionremoval"])

	return nil

}

func updateLsnparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsnparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	lsnparameter := lsn.Lsnparameter{}
	hasChange := false
	if d.HasChange("memlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Memlimit has changed for lsnparameter, starting update")
		lsnparameter.Memlimit = intPtr(d.Get("memlimit").(int))
		hasChange = true
	}
	if d.HasChange("sessionsync") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionsync has changed for lsnparameter, starting update")
		lsnparameter.Sessionsync = d.Get("sessionsync").(string)
		hasChange = true
	}
	if d.HasChange("subscrsessionremoval") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscrsessionremoval has changed for lsnparameter, starting update")
		lsnparameter.Subscrsessionremoval = d.Get("subscrsessionremoval").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lsnparameter", &lsnparameter)
		if err != nil {
			return diag.Errorf("Error updating lsnparameter")
		}
	}
	return readLsnparameterFunc(ctx, d, meta)
}

func deleteLsnparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnparameterFunc")
	//lsnparameter does not support DELETE operation
	d.SetId("")

	return nil
}
