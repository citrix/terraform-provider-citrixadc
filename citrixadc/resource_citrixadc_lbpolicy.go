package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLbpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLbpolicyFunc,
		ReadContext:   readLbpolicyFunc,
		UpdateContext: updateLbpolicyFunc,
		DeleteContext: deleteLbpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLbpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbpolicyFunc")
	client := meta.(*NetScalerNitroClient).client

	lbpolicyName := d.Get("name").(string)

	lbpolicy := lb.Lbpolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Newname:     d.Get("newname").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource("lbpolicy", lbpolicyName, &lbpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lbpolicyName)

	return readLbpolicyFunc(ctx, d, meta)
}

func readLbpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	lbpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbpolicy state %s", lbpolicyName)
	data, err := client.FindResource("lbpolicy", lbpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbpolicy state %s", lbpolicyName)
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

func updateLbpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	lbpolicyName := d.Get("name").(string)

	lbpolicy := lb.Lbpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for lbpolicy %s, starting update", lbpolicyName)
		lbpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("lbpolicy", lbpolicyName, &lbpolicy)
		if err != nil {
			return diag.Errorf("Error updating lbpolicy %s", lbpolicyName)
		}
	}
	return readLbpolicyFunc(ctx, d, meta)
}

func deleteLbpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	lbpolicyName := d.Id()
	err := client.DeleteResource("lbpolicy", lbpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
