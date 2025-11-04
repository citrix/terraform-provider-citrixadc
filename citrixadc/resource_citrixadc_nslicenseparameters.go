package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNslicenseparameters() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNslicenseparametersFunc,
		ReadContext:   readNslicenseparametersFunc,
		UpdateContext: updateNslicenseparametersFunc,
		DeleteContext: deleteNslicenseparametersFunc,
		Schema: map[string]*schema.Schema{
			"licenseexpiryalerttime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"inventoryrefreshinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"heartbeatinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"alert1gracetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"alert2gracetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNslicenseparametersFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNslicenseparametersFunc")
	client := meta.(*NetScalerNitroClient).client
	var nslicenseparametersName string
	// there is no primary key in nslicenseparameters resource. Hence generate one for terraform state maintenance
	nslicenseparametersName = resource.PrefixedUniqueId("tf-nslicenseparameters-")
	nslicenseparameters := ns.Nslicenseparameters{}
	if raw := d.GetRawConfig().GetAttr("licenseexpiryalerttime"); !raw.IsNull() {
		nslicenseparameters.Licenseexpiryalerttime = intPtr(d.Get("licenseexpiryalerttime").(int))
	}
	if raw := d.GetRawConfig().GetAttr("inventoryrefreshinterval"); !raw.IsNull() {
		nslicenseparameters.Inventoryrefreshinterval = intPtr(d.Get("inventoryrefreshinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("heartbeatinterval"); !raw.IsNull() {
		nslicenseparameters.Heartbeatinterval = intPtr(d.Get("heartbeatinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("alert1gracetimeout"); !raw.IsNull() {
		nslicenseparameters.Alert1gracetimeout = intPtr(d.Get("alert1gracetimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("alert2gracetimeout"); !raw.IsNull() {
		nslicenseparameters.Alert2gracetimeout = intPtr(d.Get("alert2gracetimeout").(int))
	}

	err := client.UpdateUnnamedResource("nslicenseparameters", &nslicenseparameters)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nslicenseparametersName)

	return readNslicenseparametersFunc(ctx, d, meta)
}

func readNslicenseparametersFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNslicenseparametersFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nslicenseparameters state")
	data, err := client.FindResource("nslicenseparameters", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nslicenseparameters state")
		d.SetId("")
		return nil
	}
	log.Println(data)
	setToInt("alert1gracetimeout", d, data["alert1gracetimeout"])
	setToInt("licenseexpiryalerttime", d, data["licenseexpiryalerttime"])
	setToInt("inventoryrefreshinterval", d, data["inventoryrefreshinterval"])
	setToInt("heartbeatinterval", d, data["heartbeatinterval"])
	setToInt("alert2gracetimeout", d, data["alert2gracetimeout"])

	return nil

}

func updateNslicenseparametersFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNslicenseparametersFunc")
	client := meta.(*NetScalerNitroClient).client

	nslicenseparameters := ns.Nslicenseparameters{}

	if d.HasChange("alert1gracetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Alert1gracetimeout has changed for nslicenseparameters, starting update")
		nslicenseparameters.Alert1gracetimeout = intPtr(d.Get("alert1gracetimeout").(int))
	}
	if d.HasChange("alert2gracetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Alert2gracetimeout has changed for nslicenseparameters, starting update")
		nslicenseparameters.Alert2gracetimeout = intPtr(d.Get("alert2gracetimeout").(int))
	}
	if d.HasChange("heartbeatinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Heartbeatinterval has changed for nslicenseparameters, starting update")
		nslicenseparameters.Heartbeatinterval = intPtr(d.Get("heartbeatinterval").(int))
	}
	if d.HasChange("inventoryrefreshinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Inventoryrefreshinterval has changed for nslicenseparameters, starting update")
		nslicenseparameters.Inventoryrefreshinterval = intPtr(d.Get("inventoryrefreshinterval").(int))
	}
	if d.HasChange("licenseexpiryalerttime") {
		log.Printf("[DEBUG]  citrixadc-provider: Licenseexpiryalerttime has changed for nslicenseparameters, starting update")
		nslicenseparameters.Licenseexpiryalerttime = intPtr(d.Get("licenseexpiryalerttime").(int))
	}

	err := client.UpdateUnnamedResource("nslicenseparameters", &nslicenseparameters)
	if err != nil {
		return diag.Errorf("Error updating nslicenseparameters")
	}

	return readNslicenseparametersFunc(ctx, d, meta)
}

func deleteNslicenseparametersFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNslicenseparametersFunc")

	d.SetId("")

	return nil
}
