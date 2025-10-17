package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnsparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnsparameterFunc,
		ReadContext:   readDnsparameterFunc,
		UpdateContext: updateDnsparameterFunc,
		DeleteContext: deleteDnsparameterFunc,
		Schema: map[string]*schema.Schema{
			"cacheecszeroprefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachehitbypass": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachenoexpire": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacherecords": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns64timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dnsrootreferral": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnssec": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ecsmaxsubnets": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxcachesize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxnegativecachesize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxnegcachettl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxpipeline": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxudppacketsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"namelookuppriority": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nxdomainratelimitthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"recursion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resolutionorder": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"splitpktqueryprocessing": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnsparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	dnsparameterName := resource.PrefixedUniqueId("tf-dnsparameter-")

	dnsparameter := dns.Dnsparameter{
		Cacheecszeroprefix:      d.Get("cacheecszeroprefix").(string),
		Cachehitbypass:          d.Get("cachehitbypass").(string),
		Cachenoexpire:           d.Get("cachenoexpire").(string),
		Cacherecords:            d.Get("cacherecords").(string),
		Dnsrootreferral:         d.Get("dnsrootreferral").(string),
		Dnssec:                  d.Get("dnssec").(string),
		Namelookuppriority:      d.Get("namelookuppriority").(string),
		Recursion:               d.Get("recursion").(string),
		Resolutionorder:         d.Get("resolutionorder").(string),
		Splitpktqueryprocessing: d.Get("splitpktqueryprocessing").(string),
	}

	if raw := d.GetRawConfig().GetAttr("dns64timeout"); !raw.IsNull() {
		dnsparameter.Dns64timeout = intPtr(d.Get("dns64timeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ecsmaxsubnets"); !raw.IsNull() {
		dnsparameter.Ecsmaxsubnets = intPtr(d.Get("ecsmaxsubnets").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxcachesize"); !raw.IsNull() {
		dnsparameter.Maxcachesize = intPtr(d.Get("maxcachesize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxnegativecachesize"); !raw.IsNull() {
		dnsparameter.Maxnegativecachesize = intPtr(d.Get("maxnegativecachesize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxnegcachettl"); !raw.IsNull() {
		dnsparameter.Maxnegcachettl = intPtr(d.Get("maxnegcachettl").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxpipeline"); !raw.IsNull() {
		dnsparameter.Maxpipeline = intPtr(d.Get("maxpipeline").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxttl"); !raw.IsNull() {
		dnsparameter.Maxttl = intPtr(d.Get("maxttl").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxudppacketsize"); !raw.IsNull() {
		dnsparameter.Maxudppacketsize = intPtr(d.Get("maxudppacketsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("minttl"); !raw.IsNull() {
		dnsparameter.Minttl = intPtr(d.Get("minttl").(int))
	}
	if raw := d.GetRawConfig().GetAttr("nxdomainratelimitthreshold"); !raw.IsNull() {
		dnsparameter.Nxdomainratelimitthreshold = intPtr(d.Get("nxdomainratelimitthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("retries"); !raw.IsNull() {
		dnsparameter.Retries = intPtr(d.Get("retries").(int))
	}

	err := client.UpdateUnnamedResource(service.Dnsparameter.Type(), &dnsparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnsparameterName)

	return readDnsparameterFunc(ctx, d, meta)
}

func readDnsparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsparameterName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsparameter state %s", dnsparameterName)
	data, err := client.FindResource(service.Dnsparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsparameter state %s", dnsparameterName)
		d.SetId("")
		return nil
	}
	d.Set("cacheecszeroprefix", data["cacheecszeroprefix"])
	d.Set("cachehitbypass", data["cachehitbypass"])
	d.Set("cachenoexpire", data["cachenoexpire"])
	d.Set("cacherecords", data["cacherecords"])
	setToInt("dns64timeout", d, data["dns64timeout"])
	d.Set("dnsrootreferral", data["dnsrootreferral"])
	d.Set("dnssec", data["dnssec"])
	setToInt("ecsmaxsubnets", d, data["ecsmaxsubnets"])
	setToInt("maxcachesize", d, data["maxcachesize"])
	setToInt("maxnegativecachesize", d, data["maxnegativecachesize"])
	setToInt("maxnegcachettl", d, data["maxnegcachettl"])
	setToInt("maxpipeline", d, data["maxpipeline"])
	setToInt("maxttl", d, data["maxttl"])
	setToInt("maxudppacketsize", d, data["maxudppacketsize"])
	setToInt("minttl", d, data["minttl"])
	d.Set("namelookuppriority", data["namelookuppriority"])
	setToInt("nxdomainratelimitthreshold", d, data["nxdomainratelimitthreshold"])
	d.Set("recursion", d.Get("recursion").(string))
	d.Set("resolutionorder", data["resolutionorder"])
	setToInt("retries", d, data["retries"])
	d.Set("splitpktqueryprocessing", data["splitpktqueryprocessing"])

	return nil

}

func updateDnsparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnsparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	dnsparameter := dns.Dnsparameter{}
	hasChange := false
	if d.HasChange("cacheecszeroprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacheecszeroprefix has changed for dnsparameter, starting update")
		dnsparameter.Cacheecszeroprefix = d.Get("cacheecszeroprefix").(string)
		hasChange = true
	}
	if d.HasChange("cachehitbypass") {
		log.Printf("[DEBUG]  citrixadc-provider: Cachehitbypass has changed for dnsparameter, starting update")
		dnsparameter.Cachehitbypass = d.Get("cachehitbypass").(string)
		hasChange = true
	}
	if d.HasChange("cachenoexpire") {
		log.Printf("[DEBUG]  citrixadc-provider: Cachenoexpire has changed for dnsparameter, starting update")
		dnsparameter.Cachenoexpire = d.Get("cachenoexpire").(string)
		hasChange = true
	}
	if d.HasChange("cacherecords") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacherecords has changed for dnsparameter, starting update")
		dnsparameter.Cacherecords = d.Get("cacherecords").(string)
		hasChange = true
	}
	if d.HasChange("dns64timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Dns64timeout has changed for dnsparameter, starting update")
		dnsparameter.Dns64timeout = intPtr(d.Get("dns64timeout").(int))
		hasChange = true
	}
	if d.HasChange("dnsrootreferral") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsrootreferral has changed for dnsparameter, starting update")
		dnsparameter.Dnsrootreferral = d.Get("dnsrootreferral").(string)
		hasChange = true
	}
	if d.HasChange("dnssec") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnssec has changed for dnsparameter, starting update")
		dnsparameter.Dnssec = d.Get("dnssec").(string)
		hasChange = true
	}
	if d.HasChange("ecsmaxsubnets") {
		log.Printf("[DEBUG]  citrixadc-provider: Ecsmaxsubnets has changed for dnsparameter, starting update")
		dnsparameter.Ecsmaxsubnets = intPtr(d.Get("ecsmaxsubnets").(int))
		hasChange = true
	}
	if d.HasChange("maxcachesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxcachesize has changed for dnsparameter, starting update")
		dnsparameter.Maxcachesize = intPtr(d.Get("maxcachesize").(int))
		hasChange = true
	}
	if d.HasChange("maxnegativecachesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxnegativecachesize has changed for dnsparameter, starting update")
		dnsparameter.Maxnegativecachesize = intPtr(d.Get("maxnegativecachesize").(int))
		hasChange = true
	}
	if d.HasChange("maxnegcachettl") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxnegcachettl has changed for dnsparameter, starting update")
		dnsparameter.Maxnegcachettl = intPtr(d.Get("maxnegcachettl").(int))
		hasChange = true
	}
	if d.HasChange("maxpipeline") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxpipeline has changed for dnsparameter, starting update")
		dnsparameter.Maxpipeline = intPtr(d.Get("maxpipeline").(int))
		hasChange = true
	}
	if d.HasChange("maxttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxttl has changed for dnsparameter, starting update")
		dnsparameter.Maxttl = intPtr(d.Get("maxttl").(int))
		hasChange = true
	}
	if d.HasChange("maxudppacketsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxudppacketsize has changed for dnsparameter, starting update")
		dnsparameter.Maxudppacketsize = intPtr(d.Get("maxudppacketsize").(int))
		hasChange = true
	}
	if d.HasChange("minttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Minttl has changed for dnsparameter, starting update")
		dnsparameter.Minttl = intPtr(d.Get("minttl").(int))
		hasChange = true
	}
	if d.HasChange("namelookuppriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Namelookuppriority has changed for dnsparameter, starting update")
		dnsparameter.Namelookuppriority = d.Get("namelookuppriority").(string)
		hasChange = true
	}
	if d.HasChange("nxdomainratelimitthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Nxdomainratelimitthreshold has changed for dnsparameter, starting update")
		dnsparameter.Nxdomainratelimitthreshold = intPtr(d.Get("nxdomainratelimitthreshold").(int))
		hasChange = true
	}
	if d.HasChange("recursion") {
		log.Printf("[DEBUG]  citrixadc-provider: Recursion has changed for dnsparameter, starting update")
		dnsparameter.Recursion = d.Get("recursion").(string)
		hasChange = true
	}
	if d.HasChange("resolutionorder") {
		log.Printf("[DEBUG]  citrixadc-provider: Resolutionorder has changed for dnsparameter, starting update")
		dnsparameter.Resolutionorder = d.Get("resolutionorder").(string)
		hasChange = true
	}
	if d.HasChange("retries") {
		log.Printf("[DEBUG]  citrixadc-provider: Retries has changed for dnsparameter, starting update")
		dnsparameter.Retries = intPtr(d.Get("retries").(int))
		hasChange = true
	}
	if d.HasChange("splitpktqueryprocessing") {
		log.Printf("[DEBUG]  citrixadc-provider: Splitpktqueryprocessing has changed for dnsparameter, starting update")
		dnsparameter.Splitpktqueryprocessing = d.Get("splitpktqueryprocessing").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Dnsparameter.Type(), &dnsparameter)
		if err != nil {
			return diag.Errorf("Error updating dnsparameter %s", err.Error())
		}
	}
	return readDnsparameterFunc(ctx, d, meta)
}

func deleteDnsparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsparameterFunc")

	d.SetId("")

	return nil
}
