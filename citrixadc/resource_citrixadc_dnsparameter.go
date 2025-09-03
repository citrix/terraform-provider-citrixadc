package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcDnsparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnsparameterFunc,
		Read:          readDnsparameterFunc,
		Update:        updateDnsparameterFunc,
		Delete:        deleteDnsparameterFunc,
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

func createDnsparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	dnsparameterName := resource.PrefixedUniqueId("tf-dnsparameter-")

	dnsparameter := dns.Dnsparameter{
		Cacheecszeroprefix:         d.Get("cacheecszeroprefix").(string),
		Cachehitbypass:             d.Get("cachehitbypass").(string),
		Cachenoexpire:              d.Get("cachenoexpire").(string),
		Cacherecords:               d.Get("cacherecords").(string),
		Dns64timeout:               d.Get("dns64timeout").(int),
		Dnsrootreferral:            d.Get("dnsrootreferral").(string),
		Dnssec:                     d.Get("dnssec").(string),
		Ecsmaxsubnets:              d.Get("ecsmaxsubnets").(int),
		Maxcachesize:               d.Get("maxcachesize").(int),
		Maxnegativecachesize:       d.Get("maxnegativecachesize").(int),
		Maxnegcachettl:             d.Get("maxnegcachettl").(int),
		Maxpipeline:                d.Get("maxpipeline").(int),
		Maxttl:                     d.Get("maxttl").(int),
		Maxudppacketsize:           d.Get("maxudppacketsize").(int),
		Minttl:                     d.Get("minttl").(int),
		Namelookuppriority:         d.Get("namelookuppriority").(string),
		Nxdomainratelimitthreshold: d.Get("nxdomainratelimitthreshold").(int),
		Recursion:                  d.Get("recursion").(string),
		Resolutionorder:            d.Get("resolutionorder").(string),
		Retries:                    d.Get("retries").(int),
		Splitpktqueryprocessing:    d.Get("splitpktqueryprocessing").(string),
	}

	err := client.UpdateUnnamedResource(service.Dnsparameter.Type(), &dnsparameter)
	if err != nil {
		return err
	}

	d.SetId(dnsparameterName)

	err = readDnsparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnsparameter but we can't read it ?? %s", dnsparameterName)
		return nil
	}
	return nil
}

func readDnsparameterFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("dns64timeout", data["dns64timeout"])
	d.Set("dnsrootreferral", data["dnsrootreferral"])
	d.Set("dnssec", data["dnssec"])
	d.Set("ecsmaxsubnets", data["ecsmaxsubnets"])
	d.Set("maxcachesize", data["maxcachesize"])
	d.Set("maxnegativecachesize", data["maxnegativecachesize"])
	d.Set("maxnegcachettl", data["maxnegcachettl"])
	d.Set("maxpipeline", data["maxpipeline"])
	d.Set("maxttl", data["maxttl"])
	d.Set("maxudppacketsize", data["maxudppacketsize"])
	d.Set("minttl", data["minttl"])
	d.Set("namelookuppriority", data["namelookuppriority"])
	d.Set("nxdomainratelimitthreshold", data["nxdomainratelimitthreshold"])
	d.Set("recursion", d.Get("recursion").(string))
	d.Set("resolutionorder", data["resolutionorder"])
	d.Set("retries", data["retries"])
	d.Set("splitpktqueryprocessing", data["splitpktqueryprocessing"])

	return nil

}

func updateDnsparameterFunc(d *schema.ResourceData, meta interface{}) error {
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
		dnsparameter.Dns64timeout = d.Get("dns64timeout").(int)
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
		dnsparameter.Ecsmaxsubnets = d.Get("ecsmaxsubnets").(int)
		hasChange = true
	}
	if d.HasChange("maxcachesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxcachesize has changed for dnsparameter, starting update")
		dnsparameter.Maxcachesize = d.Get("maxcachesize").(int)
		hasChange = true
	}
	if d.HasChange("maxnegativecachesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxnegativecachesize has changed for dnsparameter, starting update")
		dnsparameter.Maxnegativecachesize = d.Get("maxnegativecachesize").(int)
		hasChange = true
	}
	if d.HasChange("maxnegcachettl") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxnegcachettl has changed for dnsparameter, starting update")
		dnsparameter.Maxnegcachettl = d.Get("maxnegcachettl").(int)
		hasChange = true
	}
	if d.HasChange("maxpipeline") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxpipeline has changed for dnsparameter, starting update")
		dnsparameter.Maxpipeline = d.Get("maxpipeline").(int)
		hasChange = true
	}
	if d.HasChange("maxttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxttl has changed for dnsparameter, starting update")
		dnsparameter.Maxttl = d.Get("maxttl").(int)
		hasChange = true
	}
	if d.HasChange("maxudppacketsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxudppacketsize has changed for dnsparameter, starting update")
		dnsparameter.Maxudppacketsize = d.Get("maxudppacketsize").(int)
		hasChange = true
	}
	if d.HasChange("minttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Minttl has changed for dnsparameter, starting update")
		dnsparameter.Minttl = d.Get("minttl").(int)
		hasChange = true
	}
	if d.HasChange("namelookuppriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Namelookuppriority has changed for dnsparameter, starting update")
		dnsparameter.Namelookuppriority = d.Get("namelookuppriority").(string)
		hasChange = true
	}
	if d.HasChange("nxdomainratelimitthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Nxdomainratelimitthreshold has changed for dnsparameter, starting update")
		dnsparameter.Nxdomainratelimitthreshold = d.Get("nxdomainratelimitthreshold").(int)
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
		dnsparameter.Retries = d.Get("retries").(int)
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
			return fmt.Errorf("Error updating dnsparameter %s", err.Error())
		}
	}
	return readDnsparameterFunc(d, meta)
}

func deleteDnsparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsparameterFunc")

	d.SetId("")

	return nil
}
