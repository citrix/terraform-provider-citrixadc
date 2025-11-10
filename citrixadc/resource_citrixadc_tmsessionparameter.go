package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcTmsessionparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createTmsessionparameterFunc,
		ReadContext:   readTmsessionparameterFunc,
		UpdateContext: updateTmsessionparameterFunc,
		DeleteContext: deleteTmsessionparameterFunc,
		Schema: map[string]*schema.Schema{
			"defaultauthorizationaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"homepage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httponlycookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdaccount": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistentcookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistentcookievalidity": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sesstimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sso": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssocredential": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssodomain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createTmsessionparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createTmsessionparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionparameterName := resource.PrefixedUniqueId("tf-tmsessionparameter-")

	tmsessionparameter := tm.Tmsessionparameter{
		Defaultauthorizationaction: d.Get("defaultauthorizationaction").(string),
		Homepage:                   d.Get("homepage").(string),
		Httponlycookie:             d.Get("httponlycookie").(string),
		Kcdaccount:                 d.Get("kcdaccount").(string),
		Persistentcookie:           d.Get("persistentcookie").(string),
		Sso:                        d.Get("sso").(string),
		Ssocredential:              d.Get("ssocredential").(string),
		Ssodomain:                  d.Get("ssodomain").(string),
	}

	if raw := d.GetRawConfig().GetAttr("persistentcookievalidity"); !raw.IsNull() {
		tmsessionparameter.Persistentcookievalidity = intPtr(d.Get("persistentcookievalidity").(int))
	}
	if raw := d.GetRawConfig().GetAttr("sesstimeout"); !raw.IsNull() {
		tmsessionparameter.Sesstimeout = intPtr(d.Get("sesstimeout").(int))
	}

	err := client.UpdateUnnamedResource(service.Tmsessionparameter.Type(), &tmsessionparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(tmsessionparameterName)

	return readTmsessionparameterFunc(ctx, d, meta)
}

func readTmsessionparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readTmsessionparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading tmsessionparameter state")
	data, err := client.FindResource(service.Tmsessionparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing tmsessionparameter state")
		d.SetId("")
		return nil
	}
	d.Set("defaultauthorizationaction", data["defaultauthorizationaction"])
	d.Set("homepage", data["homepage"])
	d.Set("httponlycookie", data["httponlycookie"])
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("persistentcookie", data["persistentcookie"])
	setToInt("persistentcookievalidity", d, data["persistentcookievalidity"])
	setToInt("sesstimeout", d, data["sesstimeout"])
	d.Set("sso", data["sso"])
	d.Set("ssocredential", data["ssocredential"])
	d.Set("ssodomain", data["ssodomain"])

	return nil

}

func updateTmsessionparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTmsessionparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	tmsessionparameter := tm.Tmsessionparameter{}
	hasChange := false
	if d.HasChange("defaultauthorizationaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthorizationaction has changed for tmsessionparameter, starting update")
		tmsessionparameter.Defaultauthorizationaction = d.Get("defaultauthorizationaction").(string)
		hasChange = true
	}
	if d.HasChange("homepage") {
		log.Printf("[DEBUG]  citrixadc-provider: Homepage has changed for tmsessionparameter, starting update")
		tmsessionparameter.Homepage = d.Get("homepage").(string)
		hasChange = true
	}
	if d.HasChange("httponlycookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Httponlycookie has changed for tmsessionparameter, starting update")
		tmsessionparameter.Httponlycookie = d.Get("httponlycookie").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG]  citrixadc-provider: Kcdaccount has changed for tmsessionparameter, starting update")
		tmsessionparameter.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("persistentcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistentcookie has changed for tmsessionparameter, starting update")
		tmsessionparameter.Persistentcookie = d.Get("persistentcookie").(string)
		hasChange = true
	}
	if d.HasChange("persistentcookievalidity") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistentcookievalidity has changed for tmsessionparameter, starting update")
		tmsessionparameter.Persistentcookievalidity = intPtr(d.Get("persistentcookievalidity").(int))
		hasChange = true
	}
	if d.HasChange("sesstimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sesstimeout has changed for tmsessionparameter, starting update")
		tmsessionparameter.Sesstimeout = intPtr(d.Get("sesstimeout").(int))
		hasChange = true
	}
	if d.HasChange("sso") {
		log.Printf("[DEBUG]  citrixadc-provider: Sso has changed for tmsessionparameter, starting update")
		tmsessionparameter.Sso = d.Get("sso").(string)
		hasChange = true
	}
	if d.HasChange("ssocredential") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssocredential has changed for tmsessionparameter, starting update")
		tmsessionparameter.Ssocredential = d.Get("ssocredential").(string)
		hasChange = true
	}
	if d.HasChange("ssodomain") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssodomain has changed for tmsessionparameter, starting update")
		tmsessionparameter.Ssodomain = d.Get("ssodomain").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Tmsessionparameter.Type(), &tmsessionparameter)
		if err != nil {
			return diag.Errorf("Error updating tmsessionparameter")
		}
	}
	return readTmsessionparameterFunc(ctx, d, meta)
}

func deleteTmsessionparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTmsessionparameterFunc")
	// tmsessionparameter does not support DELETE operation
	d.SetId("")

	return nil
}
