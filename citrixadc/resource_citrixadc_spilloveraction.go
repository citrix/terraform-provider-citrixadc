package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/spillover"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcSpilloveraction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSpilloveractionFunc,
		ReadContext:   readSpilloveractionFunc,
		DeleteContext: deleteSpilloveractionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createSpilloveractionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSpilloveractionFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloveractionName := d.Get("name").(string)
	spilloveraction := spillover.Spilloveraction{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
	}

	_, err := client.AddResource(service.Spilloveraction.Type(), spilloveractionName, &spilloveraction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(spilloveractionName)

	return readSpilloveractionFunc(ctx, d, meta)
}

func readSpilloveractionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSpilloveractionFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloveractionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading spilloveraction state %s", spilloveractionName)
	data, err := client.FindResource(service.Spilloveraction.Type(), spilloveractionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing spilloveraction state %s", spilloveractionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])

	return nil

}

func deleteSpilloveractionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSpilloveractionFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloveractionName := d.Id()
	err := client.DeleteResource(service.Spilloveraction.Type(), spilloveractionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
