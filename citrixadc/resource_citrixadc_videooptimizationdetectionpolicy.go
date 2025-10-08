package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVideooptimizationdetectionpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVideooptimizationdetectionpolicyFunc,
		ReadContext:   readVideooptimizationdetectionpolicyFunc,
		UpdateContext: updateVideooptimizationdetectionpolicyFunc,
		DeleteContext: deleteVideooptimizationdetectionpolicyFunc,
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

func createVideooptimizationdetectionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVideooptimizationdetectionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionpolicyName := d.Get("name").(string)
	videooptimizationdetectionpolicy := videooptimization.Videooptimizationdetectionpolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Newname:     d.Get("newname").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource("videooptimizationdetectionpolicy", videooptimizationdetectionpolicyName, &videooptimizationdetectionpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(videooptimizationdetectionpolicyName)

	return readVideooptimizationdetectionpolicyFunc(ctx, d, meta)
}

func readVideooptimizationdetectionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVideooptimizationdetectionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading videooptimizationdetectionpolicy state %s", videooptimizationdetectionpolicyName)
	data, err := client.FindResource("videooptimizationdetectionpolicy", videooptimizationdetectionpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing videooptimizationdetectionpolicy state %s", videooptimizationdetectionpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateVideooptimizationdetectionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVideooptimizationdetectionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionpolicyName := d.Get("name").(string)

	videooptimizationdetectionpolicy := videooptimization.Videooptimizationdetectionpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for videooptimizationdetectionpolicy %s, starting update", videooptimizationdetectionpolicyName)
		videooptimizationdetectionpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for videooptimizationdetectionpolicy %s, starting update", videooptimizationdetectionpolicyName)
		videooptimizationdetectionpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for videooptimizationdetectionpolicy %s, starting update", videooptimizationdetectionpolicyName)
		videooptimizationdetectionpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for videooptimizationdetectionpolicy %s, starting update", videooptimizationdetectionpolicyName)
		videooptimizationdetectionpolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for videooptimizationdetectionpolicy %s, starting update", videooptimizationdetectionpolicyName)
		videooptimizationdetectionpolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for videooptimizationdetectionpolicy %s, starting update", videooptimizationdetectionpolicyName)
		videooptimizationdetectionpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for videooptimizationdetectionpolicy %s, starting update", videooptimizationdetectionpolicyName)
		videooptimizationdetectionpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("videooptimizationdetectionpolicy", videooptimizationdetectionpolicyName, &videooptimizationdetectionpolicy)
		if err != nil {
			return diag.Errorf("Error updating videooptimizationdetectionpolicy %s", videooptimizationdetectionpolicyName)
		}
	}
	return readVideooptimizationdetectionpolicyFunc(ctx, d, meta)
}

func deleteVideooptimizationdetectionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVideooptimizationdetectionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	videooptimizationdetectionpolicyName := d.Id()
	err := client.DeleteResource("videooptimizationdetectionpolicy", videooptimizationdetectionpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
