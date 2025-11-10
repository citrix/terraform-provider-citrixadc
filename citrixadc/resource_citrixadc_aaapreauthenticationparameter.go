package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAaapreauthenticationparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaapreauthenticationparameterFunc,
		ReadContext:   readAaapreauthenticationparameterFunc,
		UpdateContext: updateAaapreauthenticationparameterFunc,
		DeleteContext: deleteAaapreauthenticationparameterFunc,
		Schema: map[string]*schema.Schema{
			"deletefiles": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"killprocess": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preauthenticationaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaapreauthenticationparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaapreauthenticationparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationparameterName := resource.PrefixedUniqueId("tf-aaapreauthenticationparameter-")

	aaapreauthenticationparameter := aaa.Aaapreauthenticationparameter{
		Deletefiles:             d.Get("deletefiles").(string),
		Killprocess:             d.Get("killprocess").(string),
		Preauthenticationaction: d.Get("preauthenticationaction").(string),
		Rule:                    d.Get("rule").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaapreauthenticationparameter.Type(), &aaapreauthenticationparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaapreauthenticationparameterName)

	return readAaapreauthenticationparameterFunc(ctx, d, meta)
}

func readAaapreauthenticationparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaapreauthenticationparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading aaapreauthenticationparameter state")
	data, err := client.FindResource(service.Aaapreauthenticationparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaapreauthenticationparameter state")
		d.SetId("")
		return nil
	}
	d.Set("deletefiles", data["deletefiles"])
	d.Set("killprocess", data["killprocess"])
	d.Set("preauthenticationaction", data["preauthenticationaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAaapreauthenticationparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaapreauthenticationparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationparameter := aaa.Aaapreauthenticationparameter{}
	hasChange := false
	if d.HasChange("deletefiles") {
		log.Printf("[DEBUG]  citrixadc-provider: Deletefiles has changed for aaapreauthenticationparameter, starting update")
		aaapreauthenticationparameter.Deletefiles = d.Get("deletefiles").(string)
		hasChange = true
	}
	if d.HasChange("killprocess") {
		log.Printf("[DEBUG]  citrixadc-provider: Killprocess has changed for aaapreauthenticationparameter, starting update")
		aaapreauthenticationparameter.Killprocess = d.Get("killprocess").(string)
		hasChange = true
	}
	if d.HasChange("preauthenticationaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Preauthenticationaction has changed for aaapreauthenticationparameter, starting update")
		aaapreauthenticationparameter.Preauthenticationaction = d.Get("preauthenticationaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for aaapreauthenticationparameter, starting update")
		aaapreauthenticationparameter.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaapreauthenticationparameter.Type(), &aaapreauthenticationparameter)
		if err != nil {
			return diag.Errorf("Error updating aaapreauthenticationparameter")
		}
	}
	return readAaapreauthenticationparameterFunc(ctx, d, meta)
}

func deleteAaapreauthenticationparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaapreauthenticationparameterFunc")
	// aaapreauthenticationparameter does not suppor DELETE operation
	d.SetId("")

	return nil
}
