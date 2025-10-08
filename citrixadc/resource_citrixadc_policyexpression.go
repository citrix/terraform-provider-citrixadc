package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcPolicyexpression() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPolicyexpressionFunc,
		ReadContext:   readPolicyexpressionFunc,
		UpdateContext: updatePolicyexpressionFunc,
		DeleteContext: deletePolicyexpressionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"clientsecuritymessage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createPolicyexpressionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicyexpressionFunc")
	client := meta.(*NetScalerNitroClient).client
	policyexpressionName := d.Get("name").(string)
	policyexpression := policy.Policyexpression{
		Clientsecuritymessage: d.Get("clientsecuritymessage").(string),
		Comment:               d.Get("comment").(string),
		Name:                  d.Get("name").(string),
		Value:                 d.Get("value").(string),
	}

	_, err := client.AddResource(service.Policyexpression.Type(), policyexpressionName, &policyexpression)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policyexpressionName)

	return readPolicyexpressionFunc(ctx, d, meta)
}

func readPolicyexpressionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicyexpressionFunc")
	client := meta.(*NetScalerNitroClient).client
	policyexpressionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading policyexpression state %s", policyexpressionName)
	data, err := client.FindResource(service.Policyexpression.Type(), policyexpressionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policyexpression state %s", policyexpressionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("clientsecuritymessage", data["clientsecuritymessage"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("value", data["value"])

	return nil

}

func updatePolicyexpressionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePolicyexpressionFunc")
	client := meta.(*NetScalerNitroClient).client
	policyexpressionName := d.Get("name").(string)

	policyexpression := policy.Policyexpression{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("clientsecuritymessage") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecuritymessage has changed for policyexpression %s, starting update", policyexpressionName)
		policyexpression.Clientsecuritymessage = d.Get("clientsecuritymessage").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for policyexpression %s, starting update", policyexpressionName)
		policyexpression.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for policyexpression %s, starting update", policyexpressionName)
		policyexpression.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("value") {
		log.Printf("[DEBUG]  citrixadc-provider: Value has changed for policyexpression %s, starting update", policyexpressionName)
		policyexpression.Value = d.Get("value").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Policyexpression.Type(), policyexpressionName, &policyexpression)
		if err != nil {
			return diag.Errorf("Error updating policyexpression %s: %s", policyexpressionName, err.Error())
		}
	}
	return readPolicyexpressionFunc(ctx, d, meta)
}

func deletePolicyexpressionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicyexpressionFunc")
	client := meta.(*NetScalerNitroClient).client
	policyexpressionName := d.Id()
	err := client.DeleteResource(service.Policyexpression.Type(), policyexpressionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
