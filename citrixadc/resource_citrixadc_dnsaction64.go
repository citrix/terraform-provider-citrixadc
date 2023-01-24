package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcDnsaction64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnsaction64Func,
		Read:          readDnsaction64Func,
		Update:        updateDnsaction64Func,
		Delete:        deleteDnsaction64Func,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"actionname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"excluderule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mappedrule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createDnsaction64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsaction64Func")
	client := meta.(*NetScalerNitroClient).client
	dnsaction64Name := d.Get("actionname").(string)
	dnsaction64 := dns.Dnsaction64{
		Actionname:  dnsaction64Name,
		Excluderule: d.Get("excluderule").(string),
		Mappedrule:  d.Get("mappedrule").(string),
		Prefix:      d.Get("prefix").(string),
	}

	_, err := client.AddResource(service.Dnsaction64.Type(), dnsaction64Name, &dnsaction64)
	if err != nil {
		return err
	}

	d.SetId(dnsaction64Name)

	err = readDnsaction64Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnsaction64 but we can't read it ?? %s", dnsaction64Name)
		return nil
	}
	return nil
}

func readDnsaction64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsaction64Func")
	client := meta.(*NetScalerNitroClient).client
	dnsaction64Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsaction64 state %s", dnsaction64Name)
	data, err := client.FindResource(service.Dnsaction64.Type(), dnsaction64Name)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsaction64 state %s", dnsaction64Name)
		d.SetId("")
		return nil
	}
	d.Set("actionname", data["actionname"])
	d.Set("excluderule", data["excluderule"])
	d.Set("mappedrule", data["mappedrule"])
	d.Set("prefix", data["prefix"])

	return nil

}

func updateDnsaction64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnsaction64Func")
	client := meta.(*NetScalerNitroClient).client
	dnsaction64Name := d.Get("actionname").(string)

	dnsaction64 := dns.Dnsaction64{
		Actionname: dnsaction64Name,
	}
	hasChange := false
	if d.HasChange("excluderule") {
		log.Printf("[DEBUG]  citrixadc-provider: Excluderule has changed for dnsaction64 %s, starting update", dnsaction64Name)
		dnsaction64.Excluderule = d.Get("excluderule").(string)
		hasChange = true
	}
	if d.HasChange("mappedrule") {
		log.Printf("[DEBUG]  citrixadc-provider: Mappedrule has changed for dnsaction64 %s, starting update", dnsaction64Name)
		dnsaction64.Mappedrule = d.Get("mappedrule").(string)
		hasChange = true
	}
	if d.HasChange("prefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Prefix has changed for dnsaction64 %s, starting update", dnsaction64Name)
		dnsaction64.Prefix = d.Get("prefix").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Dnsaction64.Type(), dnsaction64Name, &dnsaction64)
		if err != nil {
			return fmt.Errorf("Error updating dnsaction64 %s", dnsaction64Name)
		}
	}
	return readDnsaction64Func(d, meta)
}

func deleteDnsaction64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsaction64Func")
	client := meta.(*NetScalerNitroClient).client
	dnsaction64Name := d.Id()
	err := client.DeleteResource(service.Dnsaction64.Type(), dnsaction64Name)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
