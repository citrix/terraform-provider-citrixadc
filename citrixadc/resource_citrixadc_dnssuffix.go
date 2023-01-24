package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceCitrixAdcDnssuffix() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnssuffixFunc,
		Read:          readDnssuffixFunc,
		Delete:        deleteDnssuffixFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"dnssuffix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDnssuffixFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnssuffixFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssuffixName := d.Get("dnssuffix").(string)
	dnssuffix := dns.Dnssuffix{
		Dnssuffix: dnssuffixName,
	}

	_, err := client.AddResource(service.Dnssuffix.Type(), dnssuffixName, &dnssuffix)
	if err != nil {
		return err
	}

	d.SetId(dnssuffixName)

	err = readDnssuffixFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnssuffix but we can't read it ?? %s", dnssuffixName)
		return nil
	}
	return nil
}

func readDnssuffixFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnssuffixFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssuffixName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnssuffix state %s", dnssuffixName)
	data, err := client.FindResource(service.Dnssuffix.Type(), dnssuffixName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnssuffix state %s", dnssuffixName)
		d.SetId("")
		return nil
	}
	d.Set("dnssuffix", data["Dnssuffix"])

	return nil

}

func deleteDnssuffixFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnssuffixFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssuffixName := d.Id()
	err := client.DeleteResource(service.Dnssuffix.Type(), dnssuffixName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
