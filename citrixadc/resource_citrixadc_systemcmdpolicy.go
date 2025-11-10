package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSystemcmdpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSystemcmdpolicyFunc,
		ReadContext:   readSystemcmdpolicyFunc,
		UpdateContext: updateSystemcmdpolicyFunc,
		DeleteContext: deleteSystemcmdpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cmdspec": {
				Type:     schema.TypeString,
				Required: true,
			},
			"policyname": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createSystemcmdpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemcmdpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	systemcmdpolicyName := d.Get("policyname").(string)

	systemcmdpolicy := system.Systemcmdpolicy{
		Action:     d.Get("action").(string),
		Cmdspec:    d.Get("cmdspec").(string),
		Policyname: d.Get("policyname").(string),
	}

	_, err := client.AddResource(service.Systemcmdpolicy.Type(), systemcmdpolicyName, &systemcmdpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(systemcmdpolicyName)

	return readSystemcmdpolicyFunc(ctx, d, meta)
}

func readSystemcmdpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemcmdpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	systemcmdpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading systemcmdpolicy state %s", systemcmdpolicyName)
	data, err := client.FindResource(service.Systemcmdpolicy.Type(), systemcmdpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing systemcmdpolicy state %s", systemcmdpolicyName)
		d.SetId("")
		return nil
	}

	d.Set("action", data["action"])
	d.Set("cmdspec", data["cmdspec"])
	d.Set("policyname", data["policyname"])

	return nil

}

func updateSystemcmdpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSystemcmdpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	systemcmdpolicyName := d.Get("policyname").(string)

	systemcmdpolicy := system.Systemcmdpolicy{
		Policyname: d.Get("policyname").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for systemcmdpolicy %s, starting update", systemcmdpolicyName)
		systemcmdpolicy.Action = d.Get("action").(string)
		systemcmdpolicy.Cmdspec = d.Get("cmdspec").(string)
		hasChange = true
	}
	if d.HasChange("cmdspec") {
		log.Printf("[DEBUG]  citrixadc-provider: Cmdspec has changed for systemcmdpolicy %s, starting update", systemcmdpolicyName)
		systemcmdpolicy.Cmdspec = d.Get("cmdspec").(string)
		systemcmdpolicy.Action = d.Get("action").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Systemcmdpolicy.Type(), systemcmdpolicyName, &systemcmdpolicy)
		if err != nil {
			return diag.Errorf("Error updating systemcmdpolicy %s:%s", systemcmdpolicyName, err.Error())
		}
	}
	return readSystemcmdpolicyFunc(ctx, d, meta)
}

func deleteSystemcmdpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemcmdpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	systemcmdpolicyName := d.Id()
	err := client.DeleteResource(service.Systemcmdpolicy.Type(), systemcmdpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
