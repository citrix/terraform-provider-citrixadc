package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"
	"strconv"

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
	val, _ := strconv.Atoi(data["alert1gracetimeout"].(string))
	d.Set("alert1gracetimeout", val)
	val, _ = strconv.Atoi(data["alert2gracetimeout"].(string))
	d.Set("alert2gracetimeout", val)

	return nil

}

func updateNslicenseparametersFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNslicenseparametersFunc")
	client := meta.(*NetScalerNitroClient).client

	nslicenseparameters := ns.Nslicenseparameters{}

	if raw := d.GetRawConfig().GetAttr("alert1gracetimeout"); !raw.IsNull() {
		nslicenseparameters.Alert1gracetimeout = intPtr(d.Get("alert1gracetimeout").(int))
	}
	if d.HasChange("alert2gracetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Alert2gracetimeout has changed for nslicenseparameters, starting update")
		nslicenseparameters.Alert2gracetimeout = intPtr(d.Get("alert2gracetimeout").(int))
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
