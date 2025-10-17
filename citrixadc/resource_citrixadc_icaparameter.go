package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcIcaparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIcaparameterFunc,
		ReadContext:   readIcaparameterFunc,
		UpdateContext: updateIcaparameterFunc,
		DeleteContext: deleteIcaparameterFunc,
		Schema: map[string]*schema.Schema{
			"enablesronhafailover": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hdxinsightnonnsap": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"edtpmtuddf": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"edtpmtuddftimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l7latencyfrequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIcaparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIcaparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	icaparameterName := resource.PrefixedUniqueId("tf-icaparameter-")

	icaparameter := ica.Icaparameter{
		Enablesronhafailover: d.Get("enablesronhafailover").(string),
		Hdxinsightnonnsap:    d.Get("hdxinsightnonnsap").(string),
		Edtpmtuddf:           d.Get("edtpmtuddf").(string),
	}

	if raw := d.GetRawConfig().GetAttr("l7latencyfrequency"); !raw.IsNull() {
		icaparameter.L7latencyfrequency = intPtr(d.Get("l7latencyfrequency").(int))
	}
	if raw := d.GetRawConfig().GetAttr("edtpmtuddftimeout"); !raw.IsNull() {
		icaparameter.Edtpmtuddftimeout = intPtr(d.Get("edtpmtuddftimeout").(int))
	}

	err := client.UpdateUnnamedResource("icaparameter", &icaparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(icaparameterName)

	return readIcaparameterFunc(ctx, d, meta)
}

func readIcaparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIcaparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading icaparameter state")
	data, err := client.FindResource("icaparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing icaparameter state")
		d.SetId("")
		return nil
	}
	d.Set("enablesronhafailover", data["enablesronhafailover"])
	d.Set("hdxinsightnonnsap", data["hdxinsightnonnsap"])
	setToInt("l7latencyfrequency", d, data["l7latencyfrequency"])
	d.Set("edtpmtuddf", data["edtpmtuddf"])
	setToInt("edtpmtuddftimeout", d, data["edtpmtuddftimeout"])

	return nil

}

func updateIcaparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIcaparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	icaparameter := ica.Icaparameter{}
	hasChange := false
	if d.HasChange("enablesronhafailover") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablesronhafailover has changed for icaparameter, starting update")
		icaparameter.Enablesronhafailover = d.Get("enablesronhafailover").(string)
		hasChange = true
	}
	if d.HasChange("edtpmtuddf") {
		log.Printf("[DEBUG]  citrixadc-provider: Edtpmtuddf has changed for icaparameter, starting update")
		icaparameter.Edtpmtuddf = d.Get("edtpmtuddf").(string)
		hasChange = true
	}
	if d.HasChange("edtpmtuddftimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Edtpmtuddftimeout has changed for icaparameter, starting update")
		icaparameter.Edtpmtuddftimeout = intPtr(d.Get("edtpmtuddftimeout").(int))
		hasChange = true
	}
	if d.HasChange("hdxinsightnonnsap") {
		log.Printf("[DEBUG]  citrixadc-provider: Hdxinsightnonnsap has changed for icaparameter, starting update")
		icaparameter.Hdxinsightnonnsap = d.Get("hdxinsightnonnsap").(string)
		hasChange = true
	}
	if d.HasChange("l7latencyfrequency") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencyfrequency has changed for icaparameter, starting update")
		icaparameter.L7latencyfrequency = intPtr(d.Get("l7latencyfrequency").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("icaparameter", &icaparameter)
		if err != nil {
			return diag.Errorf("Error updating icaparameter")
		}
	}
	return readIcaparameterFunc(ctx, d, meta)
}

func deleteIcaparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIcaparameterFunc")
	// icaparameter does not support DELETE operation
	d.SetId("")

	return nil
}
