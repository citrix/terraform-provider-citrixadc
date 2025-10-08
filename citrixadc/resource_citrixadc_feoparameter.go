package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/feo"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcFeoparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createFeoparameterFunc,
		ReadContext:   readFeoparameterFunc,
		UpdateContext: updateFeoparameterFunc,
		DeleteContext: deleteFeoparameterFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"cssinlinethressize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"imginlinethressize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jpegqualitypercent": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jsinlinethressize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createFeoparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createFeoparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	feoparameterName := resource.PrefixedUniqueId("tf-feoparameter-")

	feoparameter := make(map[string]interface{})
	if v, ok := d.GetOk("jsinlinethressize"); ok {
		feoparameter["jsinlinethressize"] = v.(int)
	}
	if v, ok := d.GetOkExists("jpegqualitypercent"); ok {
		feoparameter["jpegqualitypercent"] = v.(int)
	}
	if v, ok := d.GetOk("imginlinethressize"); ok {
		feoparameter["imginlinethressize"] = v.(int)
	}
	if v, ok := d.GetOk("cssinlinethressize"); ok {
		feoparameter["cssinlinethressize"] = v.(int)
	}

	err := client.UpdateUnnamedResource("feoparameter", &feoparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(feoparameterName)

	return readFeoparameterFunc(ctx, d, meta)
}

func readFeoparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readFeoparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading feoparameter state")
	data, err := client.FindResource("feoparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing feoparameter state")
		d.SetId("")
		return nil
	}
	setToInt("cssinlinethressize", d, data["cssinlinethressize"])
	setToInt("imginlinethressize", d, data["imginlinethressize"])
	setToInt("jpegqualitypercent", d, data["jpegqualitypercent"])
	setToInt("jsinlinethressize", d, data["jsinlinethressize"])

	return nil

}

func updateFeoparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateFeoparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	feoparameter := feo.Feoparameter{}
	hasChange := false
	if d.HasChange("cssinlinethressize") {
		log.Printf("[DEBUG]  citrixadc-provider: Cssinlinethressize has changed for feoparameter, starting update")
		feoparameter.Cssinlinethressize = d.Get("cssinlinethressize").(int)
		hasChange = true
	}
	if d.HasChange("imginlinethressize") {
		log.Printf("[DEBUG]  citrixadc-provider: Imginlinethressize has changed for feoparameter, starting update")
		feoparameter.Imginlinethressize = d.Get("imginlinethressize").(int)
		hasChange = true
	}
	if d.HasChange("jpegqualitypercent") {
		log.Printf("[DEBUG]  citrixadc-provider: Jpegqualitypercent has changed for feoparameter, starting update")
		feoparameter.Jpegqualitypercent = d.Get("jpegqualitypercent").(int)
		hasChange = true
	}
	if d.HasChange("jsinlinethressize") {
		log.Printf("[DEBUG]  citrixadc-provider: Jsinlinethressize has changed for feoparameter, starting update")
		feoparameter.Jsinlinethressize = d.Get("jsinlinethressize").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("feoparameter", &feoparameter)
		if err != nil {
			return diag.Errorf("Error updating feoparameter")
		}
	}
	return readFeoparameterFunc(ctx, d, meta)
}

func deleteFeoparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFeoparameterFunc")
	// feoparameter does not support DELETE operation
	d.SetId("")

	return nil
}
