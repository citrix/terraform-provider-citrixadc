package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cmp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcCmppolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCmppolicylabelFunc,
		Read:          readCmppolicylabelFunc,
		Delete:        deleteCmppolicylabelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createCmppolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCmppolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicylabelName := d.Get("labelname").(string)
	cmppolicylabel := cmp.Cmppolicylabel{
		Labelname: d.Get("labelname").(string),
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource(service.Cmppolicylabel.Type(), cmppolicylabelName, &cmppolicylabel)
	if err != nil {
		return err
	}

	d.SetId(cmppolicylabelName)

	err = readCmppolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cmppolicylabel but we can't read it ?? %s", cmppolicylabelName)
		return nil
	}
	return nil
}

func readCmppolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCmppolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cmppolicylabel state %s", cmppolicylabelName)
	data, err := client.FindResource(service.Cmppolicylabel.Type(), cmppolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cmppolicylabel state %s", cmppolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("type", data["type"])

	return nil

}

func deleteCmppolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCmppolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicylabelName := d.Id()
	err := client.DeleteResource(service.Cmppolicylabel.Type(), cmppolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
