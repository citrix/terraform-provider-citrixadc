package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcDnsview() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnsviewFunc,
		Read:          readDnsviewFunc,
		Delete:        deleteDnsviewFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"viewname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDnsviewFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsviewFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsviewName := d.Get("viewname").(string)
	dnsview := dns.Dnsview{
		Viewname: d.Get("viewname").(string),
	}

	_, err := client.AddResource(service.Dnsview.Type(), dnsviewName, &dnsview)
	if err != nil {
		return err
	}

	d.SetId(dnsviewName)

	err = readDnsviewFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnsview but we can't read it ?? %s", dnsviewName)
		return nil
	}
	return nil
}

func readDnsviewFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsviewFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsviewName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsview state %s", dnsviewName)
	data, err := client.FindResource(service.Dnsview.Type(), dnsviewName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsview state %s", dnsviewName)
		d.SetId("")
		return nil
	}
	d.Set("viewname", data["viewname"])

	return nil

}

func deleteDnsviewFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsviewFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsviewName := d.Id()
	err := client.DeleteResource(service.Dnsview.Type(), dnsviewName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
