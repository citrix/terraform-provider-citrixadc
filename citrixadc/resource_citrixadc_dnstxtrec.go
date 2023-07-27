package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
)

func resourceCitrixAdcDnstxtrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnstxtrecFunc,
		Read:          readDnstxtrecFunc,
		Delete:        deleteDnstxtrecFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"string": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createDnstxtrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnstxtrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnstxtrecName := d.Get("domain").(string)
	dnstxtrec := dns.Dnstxtrec{
		Domain: dnstxtrecName,
		String: toStringList(d.Get("string").([]interface{})),
		Ttl:    d.Get("ttl").(int),
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
	d.Set("string", data["String"])
	d.Set("ttl", data["ttl"])

	return nil

}

func deleteDnstxtrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnstxtrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnstxtrecName := d.Id()
	data, err := client.FindResource(service.Dnstxtrec.Type(), dnstxtrecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnstxtrec state %s", dnstxtrecName)
		d.SetId("")
		return nil
	}
	argsMap := make(map[string]string)
	argsMap["recordid"] = fmt.Sprintf("%v", data["recordid"])
	argsMap["domain"] = url.QueryEscape(d.Id())

	err = client.DeleteResourceWithArgsMap(service.Dnstxtrec.Type(), dnstxtrecName, argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
