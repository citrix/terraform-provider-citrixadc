package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcDnsptrrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnsptrrecFunc,
		ReadContext:   readDnsptrrecFunc,
		DeleteContext: deleteDnsptrrecFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"reversedomain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"nodeid": {
				Type:     schema.TypeInt,
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

func createDnsptrrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsptrrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsptrrecName := d.Get("reversedomain").(string)
	dnsptrrec := dns.Dnsptrrec{
		Domain:        d.Get("domain").(string),
		Ecssubnet:     d.Get("ecssubnet").(string),
		Nodeid:        d.Get("nodeid").(int),
		Reversedomain: d.Get("reversedomain").(string),
		Ttl:           d.Get("ttl").(int),
		Type:          d.Get("type").(string),
	}

	_, err := client.AddResource(service.Dnsptrrec.Type(), dnsptrrecName, &dnsptrrec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnsptrrecName)

	return readDnsptrrecFunc(ctx, d, meta)
}

func readDnsptrrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsptrrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsptrrecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsptrrec state %s", dnsptrrecName)
	data, err := client.FindResource(service.Dnsptrrec.Type(), dnsptrrecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsptrrec state %s", dnsptrrecName)
		d.SetId("")
		return nil
	}
	d.Set("reversedomain", data["reversedomain"])
	d.Set("domain", data["domain"])
	d.Set("ecssubnet", data["ecssubnet"])
	setToInt("nodeid", d, data["nodeid"])
	d.Set("reversedomain", data["reversedomain"])
	setToInt("ttl", d, data["ttl"])
	d.Set("type", data["type"])

	return nil

}

func deleteDnsptrrecFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsptrrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsptrrecName := d.Id()
	err := client.DeleteResource(service.Dnsptrrec.Type(), dnsptrrecName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
