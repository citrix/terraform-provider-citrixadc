package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcDnsnaptrrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnsnaptrrecFunc,
		Read:          readDnsnaptrrecFunc,
		Delete:        deleteDnsnaptrrecFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ecssubnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"order": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"preference": &schema.Schema{
				Type:     schema.TypeInt,
				Required : true,
				ForceNew: true,
			},
			"recordid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"regexp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"replacement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"services": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createDnsnaptrrecFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId(dnsnaptrrecName)

	err = readDnsnaptrrecFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnsnaptrrec but we can't read it ?? %s", dnsnaptrrecName)
		return nil
	}
	return nil
}

func readDnsnaptrrecFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("nodeid", data["nodeid"])
	d.Set("order", data["order"])
	d.Set("preference", data["preference"])
	d.Set("recordid", data["recordid"])
	d.Set("regexp", data["regexp"])
	d.Set("replacement", data["replacement"])
	d.Set("services", data["services"])
	d.Set("ttl", data["ttl"])
	d.Set("type", data["type"])

	return nil

}



func deleteDnsnaptrrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsnaptrrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsnaptrrecName := d.Id()
	err := client.DeleteResource(service.Dnsnaptrrec.Type(), dnsnaptrrecName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
