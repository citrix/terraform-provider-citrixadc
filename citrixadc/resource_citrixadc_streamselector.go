package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/stream"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcStreamselector() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createStreamselectorFunc,
		ReadContext:   readStreamselectorFunc,
		UpdateContext: updateStreamselectorFunc,
		DeleteContext: deleteStreamselectorFunc,
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
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func createStreamselectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createStreamselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	streamselectorName := d.Get("name").(string)
	streamselector := stream.Streamselector{
		Name: d.Get("name").(string),
		Rule: toStringList(d.Get("rule").([]interface{})),
	}

	_, err := client.AddResource(service.Streamselector.Type(), streamselectorName, &streamselector)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(streamselectorName)

	return readStreamselectorFunc(ctx, d, meta)
}

func readStreamselectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readStreamselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	streamselectorName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading streamselector state %s", streamselectorName)
	data, err := client.FindResource(service.Streamselector.Type(), streamselectorName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing streamselector state %s", streamselectorName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateStreamselectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateStreamselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	streamselectorName := d.Get("name").(string)

	streamselector := stream.Streamselector{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for streamselector %s, starting update", streamselectorName)
		streamselector.Rule = toStringList(d.Get("rule").([]interface{}))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Streamselector.Type(), &streamselector)
		if err != nil {
			return diag.Errorf("Error updating streamselector %s", streamselectorName)
		}
	}
	return readStreamselectorFunc(ctx, d, meta)
}

func deleteStreamselectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteStreamselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	streamselectorName := d.Id()
	err := client.DeleteResource(service.Streamselector.Type(), streamselectorName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
