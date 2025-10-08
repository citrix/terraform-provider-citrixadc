package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcCachepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCachepolicyFunc,
		ReadContext:   readCachepolicyFunc,
		UpdateContext: updateCachepolicyFunc,
		DeleteContext: deleteCachepolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"policyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalgroups": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"invalobjects": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"storeingroup": {
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

func createCachepolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCachepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicyName := d.Get("policyname").(string)
	cachepolicy := cache.Cachepolicy{
		Action:       d.Get("action").(string),
		Invalgroups:  toStringList(d.Get("invalgroups").([]interface{})),
		Invalobjects: toStringList(d.Get("invalobjects").([]interface{})),
		Policyname:   d.Get("policyname").(string),
		Rule:         d.Get("rule").(string),
		Storeingroup: d.Get("storeingroup").(string),
		Undefaction:  d.Get("undefaction").(string),
	}

	_, err := client.AddResource(service.Cachepolicy.Type(), cachepolicyName, &cachepolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cachepolicyName)

	return readCachepolicyFunc(ctx, d, meta)
}

func readCachepolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCachepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cachepolicy state %s", cachepolicyName)
	data, err := client.FindResource(service.Cachepolicy.Type(), cachepolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cachepolicy state %s", cachepolicyName)
		d.SetId("")
		return nil
	}
	d.Set("policyname", data["policyname"])
	d.Set("action", data["action"])
	d.Set("invalgroups", data["invalgroups"])
	d.Set("invalobjects", data["invalobjects"])
	d.Set("policyname", data["policyname"])
	d.Set("rule", data["rule"])
	d.Set("storeingroup", data["storeingroup"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateCachepolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCachepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicyName := d.Get("policyname").(string)

	cachepolicy := cache.Cachepolicy{
		Policyname: d.Get("policyname").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("invalgroups") {
		log.Printf("[DEBUG]  citrixadc-provider: Invalgroups has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Invalgroups = toStringList(d.Get("invalgroups").([]interface{}))
		hasChange = true
	}
	if d.HasChange("invalobjects") {
		log.Printf("[DEBUG]  citrixadc-provider: Invalobjects has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Invalobjects = toStringList(d.Get("invalobjects").([]interface{}))
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("storeingroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Storeingroup has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Storeingroup = d.Get("storeingroup").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Cachepolicy.Type(), &cachepolicy)
		if err != nil {
			return diag.Errorf("Error updating cachepolicy %s", cachepolicyName)
		}
	}
	return readCachepolicyFunc(ctx, d, meta)
}

func deleteCachepolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCachepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicyName := d.Id()
	err := client.DeleteResource(service.Cachepolicy.Type(), cachepolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
