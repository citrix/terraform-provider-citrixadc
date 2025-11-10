package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcCsparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCsparameterFunc,
		ReadContext:   readCsparameterFunc,
		UpdateContext: updateCsparameterFunc,
		DeleteContext: deleteCsparameterFunc,
		Schema: map[string]*schema.Schema{
			"stateupdate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCsparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCsparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	// there is no primary key in csparameter resource. Hence generate one for terraform state maintenance
	csparameterName := resource.PrefixedUniqueId("tf-csparameter-")

	csparameter := cs.Csparameter{
		Stateupdate: d.Get("stateupdate").(string),
	}

	err := client.UpdateUnnamedResource(service.Csparameter.Type(), &csparameter)
	if err != nil {
		return diag.Errorf("Error updating csparameter")
	}

	d.SetId(csparameterName)

	return readCsparameterFunc(ctx, d, meta)
}

func readCsparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCsparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading csparameter state")
	data, err := client.FindResource(service.Csparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing csparameter state")
		d.SetId("")
		return nil
	}
	d.Set("stateupdate", data["stateupdate"])

	return nil

}

func updateCsparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCsparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	csparameter := cs.Csparameter{}
	hasChange := false

	if d.HasChange("stateupdate") {
		log.Printf("[DEBUG]  citrixadc-provider: Stateupdate has changed for csparameter, starting update")
		csparameter.Stateupdate = d.Get("stateupdate").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Csparameter.Type(), &csparameter)
		if err != nil {
			return diag.Errorf("Error updating csparameter %s", err.Error())
		}
	}
	return readCsparameterFunc(ctx, d, meta)
}

func deleteCsparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCsparameterFunc")
	// csparameter does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
