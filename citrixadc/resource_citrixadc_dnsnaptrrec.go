package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDnsnaptrrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnsnaptrrecFunc,
		ReadContext:   readDnsnaptrrecFunc,
		DeleteContext: deleteDnsnaptrrecFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ecssubnet": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"flags": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"nodeid": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"preference": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"recordid": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"regexp": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"replacement": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"services": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createDnsnaptrrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsnaptrrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsnaptrrecName := d.Get("domain").(string)
	dnsnaptrrec := dns.Dnsnaptrrec{
		Domain:      d.Get("domain").(string),
		Ecssubnet:   d.Get("ecssubnet").(string),
		Flags:       d.Get("flags").(string),
		Nodeid:      d.Get("nodeid").(int),
		Order:       d.Get("order").(int),
		Preference:  d.Get("preference").(int),
		Recordid:    d.Get("recordid").(int),
		Regexp:      d.Get("regexp").(string),
		Replacement: d.Get("replacement").(string),
		Services:    d.Get("services").(string),
		Ttl:         d.Get("ttl").(int),
		Type:        d.Get("type").(string),
	}

	_, err := client.AddResource(service.Dnsnaptrrec.Type(), dnsnaptrrecName, &dnsnaptrrec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnsnaptrrecName)

	return readDnsnaptrrecFunc(ctx, d, meta)
}

func readDnsnaptrrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsnaptrrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsnaptrrecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsnaptrrec state %s", dnsnaptrrecName)
	data, err := client.FindResource(service.Dnsnaptrrec.Type(), dnsnaptrrecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsnaptrrec state %s", dnsnaptrrecName)
		d.SetId("")
		return nil
	}
	d.Set("domain", data["domain"])
	d.Set("ecssubnet", data["ecssubnet"])
	d.Set("flags", data["flags"])
	setToInt("nodeid", d, data["nodeid"])
	setToInt("order", d, data["order"])
	setToInt("preference", d, data["preference"])
	setToInt("recordid", d, data["recordid"])
	d.Set("regexp", data["regexp"])
	d.Set("replacement", data["replacement"])
	d.Set("services", data["services"])
	setToInt("ttl", d, data["ttl"])
	d.Set("type", data["type"])

	return nil

}

func deleteDnsnaptrrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsnaptrrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsnaptrrecName := d.Id()
	data, err := client.FindResource(service.Dnsnaptrrec.Type(), dnsnaptrrecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsnaptrrec state %s", dnsnaptrrecName)
		d.SetId("")
		return nil
	}
	argsMap := make(map[string]string)

	if _, ok := data["recordid"]; ok {
		argsMap["recordid"] = fmt.Sprintf("%v", data["recordid"])
	}

	err = client.DeleteResourceWithArgsMap(service.Dnsnaptrrec.Type(), dnsnaptrrecName, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return nil
}
