package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/reputation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcReputationsettings() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createReputationsettingsFunc,
		ReadContext:   readReputationsettingsFunc,
		UpdateContext: updateReputationsettingsFunc,
		DeleteContext: deleteReputationsettingsFunc,
		Schema: map[string]*schema.Schema{
			"proxyport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxyserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createReputationsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createReputationsettingsFunc")
	client := meta.(*NetScalerNitroClient).client
	reputationsettingsName := resource.PrefixedUniqueId("tf-reputationsettings-")

	reputationsettings := reputation.Reputationsettings{
		Proxyserver: d.Get("proxyserver").(string),
	}

	if raw := d.GetRawConfig().GetAttr("proxyport"); !raw.IsNull() {
		reputationsettings.Proxyport = intPtr(d.Get("proxyport").(int))
	}

	err := client.UpdateUnnamedResource("reputationsettings", &reputationsettings)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(reputationsettingsName)

	return readReputationsettingsFunc(ctx, d, meta)
}

func readReputationsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readReputationsettingsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading reputationsettings state")
	data, err := client.FindResource("reputationsettings", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing reputationsettings state")
		d.SetId("")
		return nil
	}
	setToInt("proxyport", d, data["proxyport"])
	d.Set("proxyserver", data["proxyserver"])

	return nil

}

func updateReputationsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateReputationsettingsFunc")
	client := meta.(*NetScalerNitroClient).client

	reputationsettings := reputation.Reputationsettings{}
	hasChange := false
	if d.HasChange("proxyport") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyport has changed for reputationsettings, starting update")
		reputationsettings.Proxyport = intPtr(d.Get("proxyport").(int))
		hasChange = true
	}
	if d.HasChange("proxyserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyserver has changed for reputationsettings, starting update")
		reputationsettings.Proxyserver = d.Get("proxyserver").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("reputationsettings", &reputationsettings)
		if err != nil {
			return diag.Errorf("Error updating reputationsettings")
		}
	}
	return readReputationsettingsFunc(ctx, d, meta)
}

func deleteReputationsettingsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteReputationsettingsFunc")
	//reputationsettings does not support DELETE operation
	d.SetId("")

	return nil
}
