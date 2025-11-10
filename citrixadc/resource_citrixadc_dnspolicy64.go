package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnspolicy64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnspolicy64Func,
		ReadContext:   readDnspolicy64Func,
		UpdateContext: updateDnspolicy64Func,
		DeleteContext: deleteDnspolicy64Func,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
		},
	}
}

func createDnspolicy64Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnspolicy64Func")
	client := meta.(*NetScalerNitroClient).client
	dnspolicy64Name := d.Get("name").(string)
	dnspolicy64 := dns.Dnspolicy64{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Dnspolicy64.Type(), dnspolicy64Name, &dnspolicy64)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnspolicy64Name)

	return readDnspolicy64Func(ctx, d, meta)
}

func readDnspolicy64Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnspolicy64Func")
	client := meta.(*NetScalerNitroClient).client
	dnspolicy64Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnspolicy64 state %s", dnspolicy64Name)
	data, err := client.FindResource(service.Dnspolicy64.Type(), dnspolicy64Name)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnspolicy64 state %s", dnspolicy64Name)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateDnspolicy64Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnspolicy64Func")
	client := meta.(*NetScalerNitroClient).client
	dnspolicy64Name := d.Get("name").(string)

	dnspolicy64 := dns.Dnspolicy64{
		Name: dnspolicy64Name,
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for dnspolicy64 %s, starting update", dnspolicy64Name)
		dnspolicy64.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for dnspolicy64 %s, starting update", dnspolicy64Name)
		dnspolicy64.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Dnspolicy64.Type(), dnspolicy64Name, &dnspolicy64)
		if err != nil {
			return diag.Errorf("Error updating dnspolicy64 %s", dnspolicy64Name)
		}
	}
	return readDnspolicy64Func(ctx, d, meta)
}

func deleteDnspolicy64Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnspolicy64Func")
	client := meta.(*NetScalerNitroClient).client
	dnspolicy64Name := d.Id()
	err := client.DeleteResource(service.Dnspolicy64.Type(), dnspolicy64Name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
