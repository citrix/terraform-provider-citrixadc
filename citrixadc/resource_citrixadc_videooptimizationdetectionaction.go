package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVideooptimizationdetectionaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVideooptimizationdetectionactionFunc,
		ReadContext:   readVideooptimizationdetectionactionFunc,
		UpdateContext: updateVideooptimizationdetectionactionFunc,
		DeleteContext: deleteVideooptimizationdetectionactionFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVideooptimizationdetectionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVideooptimizationdetectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionactionName := d.Get("name").(string)
	videooptimizationdetectionaction := videooptimization.Videooptimizationdetectionaction{
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Type:    d.Get("type").(string),
	}

	_, err := client.AddResource("videooptimizationdetectionaction", videooptimizationdetectionactionName, &videooptimizationdetectionaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(videooptimizationdetectionactionName)

	return readVideooptimizationdetectionactionFunc(ctx, d, meta)
}

func readVideooptimizationdetectionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVideooptimizationdetectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading videooptimizationdetectionaction state %s", videooptimizationdetectionactionName)
	data, err := client.FindResource("videooptimizationdetectionaction", videooptimizationdetectionactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing videooptimizationdetectionaction state %s", videooptimizationdetectionactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("type", data["type"])

	return nil

}

func updateVideooptimizationdetectionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVideooptimizationdetectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionactionName := d.Get("name").(string)

	videooptimizationdetectionaction := videooptimization.Videooptimizationdetectionaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for videooptimizationdetectionaction %s, starting update", videooptimizationdetectionactionName)
		videooptimizationdetectionaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for videooptimizationdetectionaction %s, starting update", videooptimizationdetectionactionName)
		videooptimizationdetectionaction.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("videooptimizationdetectionaction", videooptimizationdetectionactionName, &videooptimizationdetectionaction)
		if err != nil {
			return diag.Errorf("Error updating videooptimizationdetectionaction %s", videooptimizationdetectionactionName)
		}
	}
	return readVideooptimizationdetectionactionFunc(ctx, d, meta)
}

func deleteVideooptimizationdetectionactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVideooptimizationdetectionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionactionName := d.Id()
	err := client.DeleteResource("videooptimizationdetectionaction", videooptimizationdetectionactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
