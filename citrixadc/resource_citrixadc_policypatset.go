package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/policy"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcPolicypatset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPolicypatsetFunc,
		ReadContext:   readPolicypatsetFunc,
		UpdateContext: updatePolicypatsetFunc,
		DeleteContext: deletePolicypatsetFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"patsetfile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"dynamic": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createPolicypatsetFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicypatsetFunc")
	client := meta.(*NetScalerNitroClient).client
	policypatsetName := d.Get("name").(string)
	policypatset := policy.Policypatset{
		Comment:    d.Get("comment").(string),
		Name:       d.Get("name").(string),
		Dynamic:    d.Get("dynamic").(string),
		Patsetfile: d.Get("patsetfile").(string),
	}

	_, err := client.AddResource(service.Policypatset.Type(), policypatsetName, &policypatset)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policypatsetName)

	return readPolicypatsetFunc(ctx, d, meta)
}

func readPolicypatsetFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicypatsetFunc")
	client := meta.(*NetScalerNitroClient).client
	policypatsetName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading policypatset state %s", policypatsetName)
	data, err := client.FindResource(service.Policypatset.Type(), policypatsetName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policypatset state %s", policypatsetName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("patsetfile", data["patsetfile"])
	d.Set("dynamic", data["dynamic"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])

	return nil

}

func updatePolicypatsetFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePolicypatsetFunc")
	client := meta.(*NetScalerNitroClient).client
	policypatsetName := d.Id()
	policypatset := policy.Policypatset{
		Name: policypatsetName,
	}
	hasChange := false
	if d.HasChange("dynamic") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamic has changed for policypatset %s, starting update", policypatsetName)
		policypatset.Dynamic = d.Get("dynamic").(string)
		hasChange = true
	}
	if hasChange {
		_, err := client.UpdateResource(service.Policypatset.Type(), policypatsetName, &policypatset)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return readPolicypatsetFunc(ctx, d, meta)
}

func deletePolicypatsetFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicypatsetFunc")
	client := meta.(*NetScalerNitroClient).client
	policypatsetName := d.Id()
	err := client.DeleteResource(service.Policypatset.Type(), policypatsetName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
