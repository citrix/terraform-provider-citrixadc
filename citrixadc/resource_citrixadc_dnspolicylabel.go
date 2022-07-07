package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcDnspolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnspolicylabelFunc,
		Read:          readDnspolicylabelFunc,
		Delete:        deleteDnspolicylabelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transform": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDnspolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	dnspolicylabelName := d.Get("labelname").(string)
	dnspolicylabel := dns.Dnspolicylabel{
		Labelname: d.Get("labelname").(string),
		Transform: d.Get("transform").(string),
	}

	_, err := client.AddResource(service.Dnspolicylabel.Type(), dnspolicylabelName, &dnspolicylabel)
	if err != nil {
		return err
	}

	d.SetId(dnspolicylabelName)

	err = readDnspolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnspolicylabel but we can't read it ?? %s", dnspolicylabelName)
		return nil
	}
	return nil
}

func readDnspolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	dnspolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnspolicylabel state %s", dnspolicylabelName)
	data, err := client.FindResource(service.Dnspolicylabel.Type(), dnspolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnspolicylabel state %s", dnspolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("transform", data["transform"])

	return nil

}

func deleteDnspolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	dnspolicylabelName := d.Id()
	err := client.DeleteResource(service.Dnspolicylabel.Type(), dnspolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
