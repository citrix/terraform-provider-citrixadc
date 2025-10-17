package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLbsipparameters() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLbsipparametersFunc,
		ReadContext:   readLbsipparametersFunc,
		UpdateContext: updateLbsipparametersFunc,
		DeleteContext: deleteLbsipparametersFunc, // Thought lbsipparameters resource does not have a DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"addrportvip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retrydur": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rnatdstport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rnatsecuredstport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rnatsecuresrcport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rnatsrcport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sip503ratethreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLbsipparametersFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbsipparametersFunc")
	client := meta.(*NetScalerNitroClient).client
	var lbsipparametersName string

	// there is no primary key in lbsipparameters resource. Hence generate one for terraform state maintenance
	lbsipparametersName = resource.PrefixedUniqueId("tf-lbsipparameters-")

	lbsipparameters := lb.Lbsipparameters{
		Addrportvip: d.Get("addrportvip").(string),
	}

	if raw := d.GetRawConfig().GetAttr("retrydur"); !raw.IsNull() {
		lbsipparameters.Retrydur = intPtr(d.Get("retrydur").(int))
	}
	if raw := d.GetRawConfig().GetAttr("rnatdstport"); !raw.IsNull() {
		lbsipparameters.Rnatdstport = intPtr(d.Get("rnatdstport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("rnatsecuredstport"); !raw.IsNull() {
		lbsipparameters.Rnatsecuredstport = intPtr(d.Get("rnatsecuredstport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("rnatsecuresrcport"); !raw.IsNull() {
		lbsipparameters.Rnatsecuresrcport = intPtr(d.Get("rnatsecuresrcport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("rnatsrcport"); !raw.IsNull() {
		lbsipparameters.Rnatsrcport = intPtr(d.Get("rnatsrcport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("sip503ratethreshold"); !raw.IsNull() {
		lbsipparameters.Sip503ratethreshold = intPtr(d.Get("sip503ratethreshold").(int))
	}

	err := client.UpdateUnnamedResource(service.Lbsipparameters.Type(), &lbsipparameters)
	if err != nil {
		return diag.Errorf("Error updating lbsipparameters")
	}

	d.SetId(lbsipparametersName)

	return readLbsipparametersFunc(ctx, d, meta)
}

func readLbsipparametersFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbsipparametersFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading lbsipparameters state")
	data, err := client.FindResource(service.Lbsipparameters.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbsipparameters state")
		d.SetId("")
		return nil
	}
	d.Set("addrportvip", data["addrportvip"])
	setToInt("retrydur", d, data["retrydur"])
	setToInt("rnatdstport", d, data["rnatdstport"])
	setToInt("rnatsecuredstport", d, data["rnatsecuredstport"])
	setToInt("rnatsecuresrcport", d, data["rnatsecuresrcport"])
	setToInt("rnatsrcport", d, data["rnatsrcport"])
	setToInt("sip503ratethreshold", d, data["sip503ratethreshold"])

	return nil
}

func updateLbsipparametersFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbsipparametersFunc")
	client := meta.(*NetScalerNitroClient).client

	lbsipparameters := lb.Lbsipparameters{}
	hasChange := false

	if d.HasChange("addrportvip") {
		log.Printf("[DEBUG]  citrixadc-provider: Addrportvip has changed for lbsipparameters, starting update")
		lbsipparameters.Addrportvip = d.Get("addrportvip").(string)
		hasChange = true
	}
	if d.HasChange("retrydur") {
		log.Printf("[DEBUG]  citrixadc-provider: Retrydur has changed for lbsipparameters, starting update")
		lbsipparameters.Retrydur = intPtr(d.Get("retrydur").(int))
		hasChange = true
	}
	if d.HasChange("rnatdstport") {
		log.Printf("[DEBUG]  citrixadc-provider: Rnatdstport has changed for lbsipparameters, starting update")
		lbsipparameters.Rnatdstport = intPtr(d.Get("rnatdstport").(int))
		hasChange = true
	}
	if d.HasChange("rnatsecuredstport") {
		log.Printf("[DEBUG]  citrixadc-provider: Rnatsecuredstport has changed for lbsipparameters, starting update")
		lbsipparameters.Rnatsecuredstport = intPtr(d.Get("rnatsecuredstport").(int))
		hasChange = true
	}
	if d.HasChange("rnatsecuresrcport") {
		log.Printf("[DEBUG]  citrixadc-provider: Rnatsecuresrcport has changed for lbsipparameters, starting update")
		lbsipparameters.Rnatsecuresrcport = intPtr(d.Get("rnatsecuresrcport").(int))
		hasChange = true
	}
	if d.HasChange("rnatsrcport") {
		log.Printf("[DEBUG]  citrixadc-provider: Rnatsrcport has changed for lbsipparameters, starting update")
		lbsipparameters.Rnatsrcport = intPtr(d.Get("rnatsrcport").(int))
		hasChange = true
	}
	if d.HasChange("sip503ratethreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Sip503ratethreshold has changed for lbsipparameters, starting update")
		lbsipparameters.Sip503ratethreshold = intPtr(d.Get("sip503ratethreshold").(int))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Lbsipparameters.Type(), &lbsipparameters)
		if err != nil {
			return diag.Errorf("Error updating lbsipparameters: %s", err.Error())
		}
	}
	return readLbsipparametersFunc(ctx, d, meta)
}

func deleteLbsipparametersFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbsipparametersFunc")
	// lbsipparameters do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")
	return nil
}
