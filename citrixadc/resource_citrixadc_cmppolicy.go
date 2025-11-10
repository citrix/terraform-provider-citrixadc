package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cmp"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcCmppolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCmppolicyFunc,
		ReadContext:   readCmppolicyFunc,
		UpdateContext: updateCmppolicyFunc,
		DeleteContext: deleteCmppolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resaction": {
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

func createCmppolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCmppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicyName := d.Get("name").(string)
	cmppolicy := cmp.Cmppolicy{
		Name:      d.Get("name").(string),
		Resaction: d.Get("resaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Cmppolicy.Type(), cmppolicyName, &cmppolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cmppolicyName)

	return readCmppolicyFunc(ctx, d, meta)
}

func readCmppolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCmppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cmppolicy state %s", cmppolicyName)
	data, err := client.FindResource(service.Cmppolicy.Type(), cmppolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cmppolicy state %s", cmppolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("name", data["name"])
	d.Set("resaction", data["resaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateCmppolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCmppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicyName := d.Get("name").(string)

	cmppolicy := cmp.Cmppolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for cmppolicy %s, starting update", cmppolicyName)
		cmppolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("resaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Resaction has changed for cmppolicy %s, starting update", cmppolicyName)
		cmppolicy.Resaction = d.Get("resaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for cmppolicy %s, starting update", cmppolicyName)
		cmppolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Cmppolicy.Type(), cmppolicyName, &cmppolicy)
		if err != nil {
			return diag.Errorf("Error updating cmppolicy %s", cmppolicyName)
		}
	}
	return readCmppolicyFunc(ctx, d, meta)
}

func deleteCmppolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCmppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicyName := d.Id()
	err := client.DeleteResource(service.Cmppolicy.Type(), cmppolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
