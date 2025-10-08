package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVideooptimizationpacingpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVideooptimizationpacingpolicyFunc,
		ReadContext:   readVideooptimizationpacingpolicyFunc,
		UpdateContext: updateVideooptimizationpacingpolicyFunc,
		DeleteContext: deleteVideooptimizationpacingpolicyFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVideooptimizationpacingpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVideooptimizationpacingpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingpolicyName := d.Get("name").(string)

	videooptimizationpacingpolicy := videooptimization.Videooptimizationpacingpolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Newname:     d.Get("newname").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource("videooptimizationpacingpolicy", videooptimizationpacingpolicyName, &videooptimizationpacingpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(videooptimizationpacingpolicyName)

	return readVideooptimizationpacingpolicyFunc(ctx, d, meta)
}

func readVideooptimizationpacingpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVideooptimizationpacingpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading videooptimizationpacingpolicy state %s", videooptimizationpacingpolicyName)
	data, err := client.FindResource("videooptimizationpacingpolicy", videooptimizationpacingpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing videooptimizationpacingpolicy state %s", videooptimizationpacingpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateVideooptimizationpacingpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVideooptimizationpacingpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingpolicyName := d.Get("name").(string)

	videooptimizationpacingpolicy := videooptimization.Videooptimizationpacingpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for videooptimizationpacingpolicy %s, starting update", videooptimizationpacingpolicyName)
		videooptimizationpacingpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("videooptimizationpacingpolicy", videooptimizationpacingpolicyName, &videooptimizationpacingpolicy)
		if err != nil {
			return diag.Errorf("Error updating videooptimizationpacingpolicy %s", videooptimizationpacingpolicyName)
		}
	}
	return readVideooptimizationpacingpolicyFunc(ctx, d, meta)
}

func deleteVideooptimizationpacingpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVideooptimizationpacingpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationpacingpolicyName := d.Id()
	err := client.DeleteResource("videooptimizationpacingpolicy", videooptimizationpacingpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
