package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
	"net/url"
)

func resourceCitrixAdcDnsaddrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnsaddrecFunc,
		Read:          readDnsaddrecFunc,
		Delete:        deleteDnsaddrecFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ecssubnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func createDnsaddrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsaddrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsaddrecName := d.Get("hostname").(string)

	dnsaddrec := dns.Dnsaddrec{
		Ecssubnet: d.Get("ecssubnet").(string),
		Hostname:  d.Get("hostname").(string),
		Ipaddress: d.Get("ipaddress").(string),
		Nodeid:    d.Get("nodeid").(int),
		Ttl:       d.Get("ttl").(int),
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource(service.Dnsaddrec.Type(), dnsaddrecName, &dnsaddrec)
	if err != nil {
		return err
	}

	d.SetId(dnsaddrecName)

	err = readDnsaddrecFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnsaddrec but we can't read it ?? %s", dnsaddrecName)
		return nil
	}
	return nil
}

func readDnsaddrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsaddrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsaddrecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsaddrec state %s", dnsaddrecName)
	data, err := client.FindResource(service.Dnsaddrec.Type(), dnsaddrecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsaddrec state %s", dnsaddrecName)
		d.SetId("")
		return nil
	}
	d.Set("ecssubnet", data["ecssubnet"])
	d.Set("hostname", data["hostname"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("nodeid", data["nodeid"])
	d.Set("ttl", data["ttl"])
	d.Set("type", data["type"])

	return nil

}

func deleteDnsaddrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsaddrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsaddrecName := d.Id()
	argsMap := make(map[string]string) 
	if ecs,ok := d.GetOk("ecssubnet");ok{
		argsMap["ecssubnet"] = url.QueryEscape(ecs.(string))
	}
	argsMap["ipaddress"] = url.QueryEscape(d.Get("ipaddress").(string))

	err := client.DeleteResourceWithArgsMap(service.Dnsaddrec.Type(),dnsaddrecName, argsMap)
	if err != nil {
		return err
	}
	d.SetId("")

	return nil
}
