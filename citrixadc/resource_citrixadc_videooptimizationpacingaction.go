package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVideooptimizationpacingaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVideooptimizationpacingactionFunc,
		ReadContext:   readVideooptimizationpacingactionFunc,
		UpdateContext: updateVideooptimizationpacingactionFunc,
		DeleteContext: deleteVideooptimizationpacingactionFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rate": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVideooptimizationpacingactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVideooptimizationpacingactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingactionName := d.Get("name").(string)

	videooptimizationpacingaction := videooptimization.Videooptimizationpacingaction{
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Newname: d.Get("newname").(string),
	}

	if raw := d.GetRawConfig().GetAttr("rate"); !raw.IsNull() {
		videooptimizationpacingaction.Rate = intPtr(d.Get("rate").(int))
	}

	_, err := client.AddResource("videooptimizationpacingaction", videooptimizationpacingactionName, &videooptimizationpacingaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(videooptimizationpacingactionName)

	return readVideooptimizationpacingactionFunc(ctx, d, meta)
}

func readVideooptimizationpacingactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVideooptimizationpacingactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading videooptimizationpacingaction state %s", videooptimizationpacingactionName)
	data, err := client.FindResource("videooptimizationpacingaction", videooptimizationpacingactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing videooptimizationpacingaction state %s", videooptimizationpacingactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("newname", data["newname"])
	setToInt("rate", d, data["rate"])

	return nil

}

func updateVideooptimizationpacingactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVideooptimizationpacingactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingactionName := d.Get("name").(string)

	videooptimizationpacingaction := videooptimization.Videooptimizationpacingaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for videooptimizationpacingaction %s, starting update", videooptimizationpacingactionName)
		videooptimizationpacingaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for videooptimizationpacingaction %s, starting update", videooptimizationpacingactionName)
		videooptimizationpacingaction.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rate") {
		log.Printf("[DEBUG]  citrixadc-provider: Rate has changed for videooptimizationpacingaction %s, starting update", videooptimizationpacingactionName)
		videooptimizationpacingaction.Rate = intPtr(d.Get("rate").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("videooptimizationpacingaction", videooptimizationpacingactionName, &videooptimizationpacingaction)
		if err != nil {
			return diag.Errorf("Error updating videooptimizationpacingaction %s", videooptimizationpacingactionName)
		}
	}
	return readVideooptimizationpacingactionFunc(ctx, d, meta)
}

func deleteVideooptimizationpacingactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVideooptimizationpacingactionFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingactionName := d.Id()
	err := client.DeleteResource("videooptimizationpacingaction", videooptimizationpacingactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
