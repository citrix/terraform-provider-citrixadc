package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcCacheselector() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCacheselectorFunc,
		ReadContext:   readCacheselectorFunc,
		UpdateContext: updateCacheselectorFunc,
		DeleteContext: deleteCacheselectorFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"selectorname": {
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

func createCacheselectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCacheselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	var cacheselectorName string
	cacheselectorName = d.Get("selectorname").(string)
	cacheselector := cache.Cacheselector{
		Rule:         toStringList(d.Get("rule").([]interface{})),
		Selectorname: d.Get("selectorname").(string),
	}

	_, err := client.AddResource(service.Cacheselector.Type(), cacheselectorName, &cacheselector)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cacheselectorName)

	return readCacheselectorFunc(ctx, d, meta)
}

func readCacheselectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCacheselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheselectorName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cacheselector state %s", cacheselectorName)
	data, err := client.FindResource(service.Cacheselector.Type(), cacheselectorName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cacheselector state %s", cacheselectorName)
		d.SetId("")
		return nil
	}
	d.Set("selectorname", data["selectorname"])
	d.Set("rule", data["rule"])

	return nil

}

func updateCacheselectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCacheselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheselectorName := d.Get("selectorname").(string)

	cacheselector := cache.Cacheselector{
		Selectorname: d.Get("selectorname").(string),
	}
	hasChange := false
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for cacheselector %s, starting update", cacheselectorName)
		cacheselector.Rule = toStringList(d.Get("rule").([]interface{}))
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Cacheselector.Type(), &cacheselector)
		if err != nil {
			return diag.Errorf("Error updating cacheselector %s", cacheselectorName)
		}
	}
	return readCacheselectorFunc(ctx, d, meta)
}

func deleteCacheselectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCacheselectorFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheselectorName := d.Id()
	err := client.DeleteResource(service.Cacheselector.Type(), cacheselectorName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
