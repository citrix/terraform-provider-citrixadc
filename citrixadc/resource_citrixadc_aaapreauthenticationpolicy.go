package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAaapreauthenticationpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaapreauthenticationpolicyFunc,
		ReadContext:   readAaapreauthenticationpolicyFunc,
		UpdateContext: updateAaapreauthenticationpolicyFunc,
		DeleteContext: deleteAaapreauthenticationpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reqaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaapreauthenticationpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaapreauthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationpolicyName := d.Get("name").(string)

	aaapreauthenticationpolicy := aaa.Aaapreauthenticationpolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Aaapreauthenticationpolicy.Type(), aaapreauthenticationpolicyName, &aaapreauthenticationpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaapreauthenticationpolicyName)

	return readAaapreauthenticationpolicyFunc(ctx, d, meta)
}

func readAaapreauthenticationpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaapreauthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading aaapreauthenticationpolicy state %s", aaapreauthenticationpolicyName)
	data, err := client.FindResource(service.Aaapreauthenticationpolicy.Type(), aaapreauthenticationpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaapreauthenticationpolicy state %s", aaapreauthenticationpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAaapreauthenticationpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaapreauthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationpolicyName := d.Get("name").(string)

	aaapreauthenticationpolicy := aaa.Aaapreauthenticationpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for aaapreauthenticationpolicy %s, starting update", aaapreauthenticationpolicyName)
		aaapreauthenticationpolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for aaapreauthenticationpolicy %s, starting update", aaapreauthenticationpolicyName)
		aaapreauthenticationpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaapreauthenticationpolicy.Type(), &aaapreauthenticationpolicy)
		if err != nil {
			return diag.Errorf("Error updating aaapreauthenticationpolicy %s", aaapreauthenticationpolicyName)
		}
	}
	return readAaapreauthenticationpolicyFunc(ctx, d, meta)
}

func deleteAaapreauthenticationpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaapreauthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationpolicyName := d.Id()
	err := client.DeleteResource(service.Aaapreauthenticationpolicy.Type(), aaapreauthenticationpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
