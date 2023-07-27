package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSslpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslpolicylabelFunc,
		Read:          readSslpolicylabelFunc,
		Delete:        deleteSslpolicylabelFunc,
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

func createSslpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslpolicylabelName = d.Get("labelname").(string)

	sslpolicylabel := ssl.Sslpolicylabel{
		Labelname: sslpolicylabelName,
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource(service.Sslpolicylabel.Type(), sslpolicylabelName, &sslpolicylabel)
	if err != nil {
		return err
	}

	d.SetId(sslpolicylabelName)

	err = readSslpolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslpolicylabel but we can't read it ?? %s", sslpolicylabelName)
		return nil
	}
	return nil
}

func readSslpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	sslpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslpolicylabel state %s", sslpolicylabelName)
	data, err := client.FindResource(service.Sslpolicylabel.Type(), sslpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslpolicylabel state %s", sslpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("labelname", data["labelname"])
	d.Set("type", data["type"])

	return nil

}

func deleteSslpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	sslpolicylabelName := d.Id()
	err := client.DeleteResource(service.Sslpolicylabel.Type(), sslpolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
