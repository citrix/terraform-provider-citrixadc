package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnssoarec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnssoarecFunc,
		ReadContext:   readDnssoarecFunc,
		UpdateContext: updateDnssoarecFunc,
		DeleteContext: deleteDnssoarecFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"contact": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ecssubnet": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expire": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minimum": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"originserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retry": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"serial": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnssoarecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnssoarecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssoarecId := d.Get("domain").(string)
	dnssoarec := dns.Dnssoarec{
		Contact:      d.Get("contact").(string),
		Domain:       d.Get("domain").(string),
		Ecssubnet:    d.Get("ecssubnet").(string),
		Originserver: d.Get("originserver").(string),
	}

	if raw := d.GetRawConfig().GetAttr("expire"); !raw.IsNull() {
		dnssoarec.Expire = intPtr(d.Get("expire").(int))
	}
	if raw := d.GetRawConfig().GetAttr("minimum"); !raw.IsNull() {
		dnssoarec.Minimum = intPtr(d.Get("minimum").(int))
	}
	if raw := d.GetRawConfig().GetAttr("nodeid"); !raw.IsNull() {
		dnssoarec.Nodeid = intPtr(d.Get("nodeid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("refresh"); !raw.IsNull() {
		dnssoarec.Refresh = intPtr(d.Get("refresh").(int))
	}
	if raw := d.GetRawConfig().GetAttr("retry"); !raw.IsNull() {
		dnssoarec.Retry = intPtr(d.Get("retry").(int))
	}
	if raw := d.GetRawConfig().GetAttr("serial"); !raw.IsNull() {
		dnssoarec.Serial = intPtr(d.Get("serial").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ttl"); !raw.IsNull() {
		dnssoarec.Ttl = intPtr(d.Get("ttl").(int))
	}

	_, err := client.AddResource(service.Dnssoarec.Type(), dnssoarecId, &dnssoarec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnssoarecId)

	return readDnssoarecFunc(ctx, d, meta)
}

func readDnssoarecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnssoarecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssoarecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnssoarec state %s", dnssoarecName)
	data, err := client.FindResource(service.Dnssoarec.Type(), dnssoarecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnssoarec state %s", dnssoarecName)
		d.SetId("")
		return nil
	}
	d.Set("contact", data["contact"])
	d.Set("domain", data["domain"])
	d.Set("ecssubnet", data["ecssubnet"])
	setToInt("expire", d, data["expire"])
	setToInt("minimum", d, data["minimum"])
	setToInt("nodeid", d, data["nodeid"])
	d.Set("originserver", data["originserver"])
	setToInt("refresh", d, data["refresh"])
	setToInt("retry", d, data["retry"])
	setToInt("serial", d, data["serial"])
	setToInt("ttl", d, data["ttl"])

	return nil

}

func updateDnssoarecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnssoarecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssoarecId := d.Get("domain").(string)

	dnssoarec := dns.Dnssoarec{
		Domain: dnssoarecId,
	}
	hasChange := false
	if d.HasChange("contact") {
		log.Printf("[DEBUG]  citrixadc-provider: Contact has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Contact = d.Get("contact").(string)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG]  citrixadc-provider: Domain has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("ecssubnet") {
		log.Printf("[DEBUG]  citrixadc-provider: Ecssubnet has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Ecssubnet = d.Get("ecssubnet").(string)
		hasChange = true
	}
	if d.HasChange("expire") {
		log.Printf("[DEBUG]  citrixadc-provider: Expire has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Expire = intPtr(d.Get("expire").(int))
		hasChange = true
	}
	if d.HasChange("minimum") {
		log.Printf("[DEBUG]  citrixadc-provider: Minimum has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Minimum = intPtr(d.Get("minimum").(int))
		hasChange = true
	}
	if d.HasChange("nodeid") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodeid has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Nodeid = intPtr(d.Get("nodeid").(int))
		hasChange = true
	}
	if d.HasChange("originserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Originserver has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Originserver = d.Get("originserver").(string)
		hasChange = true
	}
	if d.HasChange("refresh") {
		log.Printf("[DEBUG]  citrixadc-provider: Refresh has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Refresh = intPtr(d.Get("refresh").(int))
		hasChange = true
	}
	if d.HasChange("retry") {
		log.Printf("[DEBUG]  citrixadc-provider: Retry has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Retry = intPtr(d.Get("retry").(int))
		hasChange = true
	}
	if d.HasChange("serial") {
		log.Printf("[DEBUG]  citrixadc-provider: Serial has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Serial = intPtr(d.Get("serial").(int))
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Ttl has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Ttl = intPtr(d.Get("ttl").(int))
		hasChange = true
	}
	if hasChange {
		_, err := client.UpdateResource(service.Dnssoarec.Type(), dnssoarecId, &dnssoarec)
		if err != nil {
			return diag.Errorf("Error updating dnssoarec %s. %s", dnssoarecId, err)
		}
	}
	return readDnssoarecFunc(ctx, d, meta)
}

func deleteDnssoarecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnssoarecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssoarecId := d.Id()
	err := client.DeleteResource(service.Dnssoarec.Type(), dnssoarecId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
