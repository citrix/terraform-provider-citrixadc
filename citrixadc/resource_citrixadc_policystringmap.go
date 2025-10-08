package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcPolicystringmap() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPolicystringmapFunc,
		ReadContext:   readPolicystringmapFunc,
		UpdateContext: updatePolicystringmapFunc,
		DeleteContext: deletePolicystringmapFunc,
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createPolicystringmapFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicystringmapFunc")
	client := meta.(*NetScalerNitroClient).client

	policystringmapName := d.Get("name").(string)

	policystringmap := policy.Policystringmap{
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
	}

	_, err := client.AddResource(service.Policystringmap.Type(), policystringmapName, &policystringmap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policystringmapName)

	return readPolicystringmapFunc(ctx, d, meta)
}

func readPolicystringmapFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicystringmapFunc")
	client := meta.(*NetScalerNitroClient).client
	policystringmapName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading policystringmap state %s", policystringmapName)
	data, err := client.FindResource(service.Policystringmap.Type(), policystringmapName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policystringmap state %s", policystringmapName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])

	return nil

}

func updatePolicystringmapFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePolicystringmapFunc")
	client := meta.(*NetScalerNitroClient).client
	policystringmapName := d.Get("name").(string)

	policystringmap := policy.Policystringmap{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for policystringmap %s, starting update", policystringmapName)
		policystringmap.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for policystringmap %s, starting update", policystringmapName)
		policystringmap.Name = d.Get("name").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Policystringmap.Type(), policystringmapName, &policystringmap)
		if err != nil {
			return diag.Errorf("Error updating policystringmap %s", policystringmapName)
		}
	}
	return readPolicystringmapFunc(ctx, d, meta)
}

func deletePolicystringmapFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicystringmapFunc")
	client := meta.(*NetScalerNitroClient).client
	policystringmapName := d.Id()
	err := client.DeleteResource(service.Policystringmap.Type(), policystringmapName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
