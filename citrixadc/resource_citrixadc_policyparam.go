package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcPolicyparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPolicyparamFunc,
		ReadContext:   readPolicyparamFunc,
		UpdateContext: updatePolicyparamFunc,
		DeleteContext: deletePolicyparamFunc, // Thought policyparam resource does not have a DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createPolicyparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicyparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var policyparamName string

	// there is no primary key in policyparam resource. Hence generate one for terraform state maintenance
	policyparamName = resource.PrefixedUniqueId("tf-policyparam-")

	policyparam := policy.Policyparam{
		Timeout: d.Get("timeout").(int),
	}

	err := client.UpdateUnnamedResource("policyparam", &policyparam)
	if err != nil {
		return diag.Errorf("Error updating policyparam")
	}

	d.SetId(policyparamName)

	return readPolicyparamFunc(ctx, d, meta)
}

func readPolicyparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicyparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading policyparam state")
	data, err := client.FindResource("policyparam", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policyparam state")
		d.SetId("")
		return nil
	}
	setToInt("timeout", d, data["timeout"])

	return nil

}

func updatePolicyparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePolicyparamFunc")
	client := meta.(*NetScalerNitroClient).client

	policyparam := policy.Policyparam{}
	hasChange := false

	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for policyparam, starting update")
		policyparam.Timeout = d.Get("timeout").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("policyparam", &policyparam)
		if err != nil {
			return diag.Errorf("Error updating policyparam %s", err.Error())
		}
	}
	return readPolicyparamFunc(ctx, d, meta)
}

func deletePolicyparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicyparamFunc")
	// policyparam does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")
	return nil
}
