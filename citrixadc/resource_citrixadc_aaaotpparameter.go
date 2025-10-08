package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAaaotpparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaaotpparameterFunc,
		ReadContext:   readAaaotpparameterFunc,
		UpdateContext: updateAaaotpparameterFunc,
		DeleteContext: deleteAaaotpparameterFunc,
		Schema: map[string]*schema.Schema{
			"encryption": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxotpdevices": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaaotpparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaaotpparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	aaaotpparameterName := resource.PrefixedUniqueId("tf-aaaotpparameter-")

	aaaotpparameter := aaa.Aaaotpparameter{
		Encryption:    d.Get("encryption").(string),
		Maxotpdevices: d.Get("maxotpdevices").(int),
	}

	err := client.UpdateUnnamedResource("aaaotpparameter", &aaaotpparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaaotpparameterName)

	return readAaaotpparameterFunc(ctx, d, meta)
}

func readAaaotpparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaaotpparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading aaaotpparameter state")
	data, err := client.FindResource("aaaotpparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaaotpparameter state")
		d.SetId("")
		return nil
	}
	d.Set("encryption", data["encryption"])
	setToInt("maxotpdevices", d, data["maxotpdevices"])

	return nil

}

func updateAaaotpparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaaotpparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	aaaotpparameter := aaa.Aaaotpparameter{}
	hasChange := false
	if d.HasChange("encryption") {
		log.Printf("[DEBUG]  citrixadc-provider: Encryption has changed for aaaotpparameter, starting update")
		aaaotpparameter.Encryption = d.Get("encryption").(string)
		hasChange = true
	}
	if d.HasChange("maxotpdevices") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxotpdevices has changed for aaaotpparameter, starting update")
		aaaotpparameter.Maxotpdevices = d.Get("maxotpdevices").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("aaaotpparameter", &aaaotpparameter)
		if err != nil {
			return diag.Errorf("Error updating aaaotpparameter")
		}
	}
	return readAaaotpparameterFunc(ctx, d, meta)
}

func deleteAaaotpparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaaotpparameterFunc")
	// aaaotpparameter does not support DELETE operation
	d.SetId("")

	return nil
}
