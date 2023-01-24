package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceCitrixAdcDnsptrrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnsptrrecFunc,
		Read:          readDnsptrrecFunc,
		Delete:        deleteDnsptrrecFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"reversedomain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
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

func createDnsptrrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsptrrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsptrrecName:= d.Get("reversedomain").(string)
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
		return err
	}

	d.SetId(dnsptrrecName)

	err = readDnsptrrecFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnsptrrec but we can't read it ?? %s", dnsptrrecName)
		return nil
	}
	return nil
}

func readDnsptrrecFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("nodeid", data["nodeid"])
	d.Set("reversedomain", data["reversedomain"])
	d.Set("ttl", data["ttl"])
	d.Set("type", data["type"])

	return nil

}

func deleteDnsptrrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsptrrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsptrrecName := d.Id()
	err := client.DeleteResource(service.Dnsptrrec.Type(), dnsptrrecName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
