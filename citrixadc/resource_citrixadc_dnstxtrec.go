package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcDnstxtrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnstxtrecFunc,
		Read:          readDnstxtrecFunc,
		Delete:        deleteDnstxtrecFunc,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"string": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"ecssubnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"recordid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnstxtrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnstxtrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnstxtrecName := d.Get("domain").(string)
	dnstxtrec := dns.Dnstxtrec{
		Domain:    d.Get("domain").(string),
		Ecssubnet: d.Get("ecssubnet").(string),
		Nodeid:    d.Get("nodeid").(int),
		Recordid:  d.Get("recordid").(int),
		String:    toStringList(d.Get("string").([]interface{})),
		Ttl:       d.Get("ttl").(int),
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource(service.Dnstxtrec.Type(), dnstxtrecName, &dnstxtrec)
	if err != nil {
		return err
	}

	d.SetId(dnstxtrecName)

	err = readDnstxtrecFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnstxtrec but we can't read it ?? %s", dnstxtrecName)
		return nil
	}
	return nil
}

func readDnstxtrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnstxtrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnstxtrecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnstxtrec state %s", dnstxtrecName)
	data, err := client.FindResource(service.Dnstxtrec.Type(), dnstxtrecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnstxtrec state %s", dnstxtrecName)
		d.SetId("")
		return nil
	}
	d.Set("domain", data["domain"])
	d.Set("domain", data["domain"])
	d.Set("ecssubnet", data["ecssubnet"])
	d.Set("nodeid", data["nodeid"])
	d.Set("recordid", data["recordid"])
	d.Set("string", data["string"])
	d.Set("ttl", data["ttl"])
	d.Set("type", data["type"])

	return nil

}

func deleteDnstxtrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnstxtrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnstxtrecName := d.Id()
	err := client.DeleteResource(service.Dnstxtrec.Type(), dnstxtrecName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
