package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/adm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAdmparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAdmparameterFunc,
		ReadContext:   readAdmparameterFunc,
		UpdateContext: updateAdmparameterFunc,
		DeleteContext: deleteAdmparameterFunc,
		Schema: map[string]*schema.Schema{
			"admserviceconnect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAdmparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAdmparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	admparameterName := resource.PrefixedUniqueId("tf-admparameter-")

	admparameter := adm.Admparameter{
		Admserviceconnect: d.Get("admserviceconnect").(string),
	}

	err := client.UpdateUnnamedResource("admparameter", &admparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(admparameterName)

	return readAdmparameterFunc(ctx, d, meta)
}

func readAdmparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAdmparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading admparameter state")
	data, err := client.FindResource("admparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing admparameter state")
		d.SetId("")
		return nil
	}
	d.Set("admserviceconnect", data["admserviceconnect"])

	return nil

}

func updateAdmparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAdmparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	admparameter := adm.Admparameter{}
	hasChange := false
	if d.HasChange("admserviceconnect") {
		log.Printf("[DEBUG]  citrixadc-provider: Admserviceconnect has changed for admparameter, starting update")
		admparameter.Admserviceconnect = d.Get("admserviceconnect").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("admparameter", &admparameter)
		if err != nil {
			return diag.Errorf("Error updating admparameter")
		}
	}
	return readAdmparameterFunc(ctx, d, meta)
}

func deleteAdmparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAdmparameterFunc")
	// admparameter does not support DELETE operation
	d.SetId("")

	return nil
}
