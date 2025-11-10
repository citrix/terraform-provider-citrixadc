package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcContentinspectionparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createContentinspectionparameterFunc,
		ReadContext:   readContentinspectionparameterFunc,
		UpdateContext: updateContentinspectionparameterFunc,
		DeleteContext: deleteContentinspectionparameterFunc,
		Schema: map[string]*schema.Schema{
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createContentinspectionparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectionparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionparameterName := resource.PrefixedUniqueId("tf-contentinspectionparameter-")

	contentinspectionparameter := contentinspection.Contentinspectionparameter{
		Undefaction: d.Get("undefaction").(string),
	}

	err := client.UpdateUnnamedResource("contentinspectionparameter", &contentinspectionparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(contentinspectionparameterName)

	return readContentinspectionparameterFunc(ctx, d, meta)
}

func readContentinspectionparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectionparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectionparameter state")
	data, err := client.FindResource("contentinspectionparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectionparameter state")
		d.SetId("")
		return nil
	}
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateContentinspectionparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateContentinspectionparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	contentinspectionparameter := contentinspection.Contentinspectionparameter{}
	hasChange := false
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for contentinspectionparameter, starting update")
		contentinspectionparameter.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("contentinspectionparameter", &contentinspectionparameter)
		if err != nil {
			return diag.Errorf("Error updating contentinspectionparameter")
		}
	}
	return readContentinspectionparameterFunc(ctx, d, meta)
}

func deleteContentinspectionparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectionparameterFunc")
	//contentinspectionparameter does not support DELETE operation
	d.SetId("")

	return nil
}
