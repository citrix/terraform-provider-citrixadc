package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cmp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcCmpparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCmpparameterFunc,
		ReadContext:   readCmpparameterFunc,
		UpdateContext: updateCmpparameterFunc,
		DeleteContext: deleteCmpparameterFunc,
		Schema: map[string]*schema.Schema{
			"addvaryheader": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmpbypasspct": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cmplevel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmponpush": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"externalcache": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"heurexpiry": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"heurexpiryhistwt": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"heurexpirythres": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minressize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"policytype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quantumsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servercmp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"varyheadervalue": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCmpparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCmpparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	cmpparameterName := resource.PrefixedUniqueId("tf-cmpparameter-")

	cmpparameter := cmp.Cmpparameter{
		Addvaryheader:   d.Get("addvaryheader").(string),
		Cmplevel:        d.Get("cmplevel").(string),
		Cmponpush:       d.Get("cmponpush").(string),
		Externalcache:   d.Get("externalcache").(string),
		Heurexpiry:      d.Get("heurexpiry").(string),
		Policytype:      d.Get("policytype").(string),
		Servercmp:       d.Get("servercmp").(string),
		Varyheadervalue: d.Get("varyheadervalue").(string),
	}

	if raw := d.GetRawConfig().GetAttr("cmpbypasspct"); !raw.IsNull() {
		cmpparameter.Cmpbypasspct = intPtr(d.Get("cmpbypasspct").(int))
	}
	if raw := d.GetRawConfig().GetAttr("heurexpiryhistwt"); !raw.IsNull() {
		cmpparameter.Heurexpiryhistwt = intPtr(d.Get("heurexpiryhistwt").(int))
	}
	if raw := d.GetRawConfig().GetAttr("heurexpirythres"); !raw.IsNull() {
		cmpparameter.Heurexpirythres = intPtr(d.Get("heurexpirythres").(int))
	}
	if raw := d.GetRawConfig().GetAttr("minressize"); !raw.IsNull() {
		cmpparameter.Minressize = intPtr(d.Get("minressize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("quantumsize"); !raw.IsNull() {
		cmpparameter.Quantumsize = intPtr(d.Get("quantumsize").(int))
	}

	err := client.UpdateUnnamedResource(service.Cmpparameter.Type(), &cmpparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cmpparameterName)

	return readCmpparameterFunc(ctx, d, meta)
}

func readCmpparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCmpparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading cmpparameter state")
	data, err := client.FindResource(service.Cmpparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cmpparameter state")
		d.SetId("")
		return nil
	}
	d.Set("addvaryheader", data["addvaryheader"])
	setToInt("cmpbypasspct", d, data["cmpbypasspct"])
	d.Set("cmplevel", data["cmplevel"])
	d.Set("cmponpush", data["cmponpush"])
	d.Set("externalcache", data["externalcache"])
	d.Set("heurexpiry", data["heurexpiry"])
	setToInt("heurexpiryhistwt", d, data["heurexpiryhistwt"])
	setToInt("heurexpirythres", d, data["heurexpirythres"])
	setToInt("minressize", d, data["minressize"])
	d.Set("policytype", data["policytype"])
	setToInt("quantumsize", d, data["quantumsize"])
	d.Set("servercmp", data["servercmp"])
	d.Set("varyheadervalue", data["varyheadervalue"])

	return nil

}

func updateCmpparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCmpparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	cmpparameter := cmp.Cmpparameter{}
	hasChange := false
	if d.HasChange("addvaryheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Addvaryheader has changed for cmpparameter, starting update")
		cmpparameter.Addvaryheader = d.Get("addvaryheader").(string)
		hasChange = true
	}
	if d.HasChange("cmpbypasspct") {
		log.Printf("[DEBUG]  citrixadc-provider: Cmpbypasspct has changed for cmpparameter, starting update")
		cmpparameter.Cmpbypasspct = intPtr(d.Get("cmpbypasspct").(int))
		hasChange = true
	}
	if d.HasChange("cmplevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Cmplevel has changed for cmpparameter, starting update")
		cmpparameter.Cmplevel = d.Get("cmplevel").(string)
		hasChange = true
	}
	if d.HasChange("cmponpush") {
		log.Printf("[DEBUG]  citrixadc-provider: Cmponpush has changed for cmpparameter, starting update")
		cmpparameter.Cmponpush = d.Get("cmponpush").(string)
		hasChange = true
	}
	if d.HasChange("externalcache") {
		log.Printf("[DEBUG]  citrixadc-provider: Externalcache has changed for cmpparameter, starting update")
		cmpparameter.Externalcache = d.Get("externalcache").(string)
		hasChange = true
	}
	if d.HasChange("heurexpiry") {
		log.Printf("[DEBUG]  citrixadc-provider: Heurexpiry has changed for cmpparameter, starting update")
		cmpparameter.Heurexpiry = d.Get("heurexpiry").(string)
		hasChange = true
	}
	if d.HasChange("heurexpiryhistwt") {
		log.Printf("[DEBUG]  citrixadc-provider: Heurexpiryhistwt has changed for cmpparameter, starting update")
		cmpparameter.Heurexpiryhistwt = intPtr(d.Get("heurexpiryhistwt").(int))
		hasChange = true
	}
	if d.HasChange("heurexpirythres") {
		log.Printf("[DEBUG]  citrixadc-provider: Heurexpirythres has changed for cmpparameter, starting update")
		cmpparameter.Heurexpirythres = intPtr(d.Get("heurexpirythres").(int))
		hasChange = true
	}
	if d.HasChange("minressize") {
		log.Printf("[DEBUG]  citrixadc-provider: Minressize has changed for cmpparameter, starting update")
		cmpparameter.Minressize = intPtr(d.Get("minressize").(int))
		hasChange = true
	}
	if d.HasChange("policytype") {
		log.Printf("[DEBUG]  citrixadc-provider: Policytype has changed for cmpparameter, starting update")
		cmpparameter.Policytype = d.Get("policytype").(string)
		hasChange = true
	}
	if d.HasChange("quantumsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Quantumsize has changed for cmpparameter, starting update")
		cmpparameter.Quantumsize = intPtr(d.Get("quantumsize").(int))
		hasChange = true
	}
	if d.HasChange("servercmp") {
		log.Printf("[DEBUG]  citrixadc-provider: Servercmp has changed for cmpparameter, starting update")
		cmpparameter.Servercmp = d.Get("servercmp").(string)
		hasChange = true
	}
	if d.HasChange("varyheadervalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Varyheadervalue has changed for cmpparameter, starting update")
		cmpparameter.Varyheadervalue = d.Get("varyheadervalue").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Cmpparameter.Type(), &cmpparameter)
		if err != nil {
			return diag.Errorf("Error updating cmpparameter")
		}
	}
	return readCmpparameterFunc(ctx, d, meta)
}

func deleteCmpparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCmpparameterFunc")
	// cmpparameter does not support DELETE operation
	d.SetId("")

	return nil
}
