package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/transform"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcTransformpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createTransformpolicylabelFunc,
		Read:          readTransformpolicylabelFunc,
		Delete:        deleteTransformpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createTransformpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createTransformpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	transformpolicylabelName := d.Get("labelname").(string)
	transformpolicylabel := transform.Transformpolicylabel{
		Labelname:       transformpolicylabelName,
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	_, err := client.AddResource(service.Transformpolicylabel.Type(), transformpolicylabelName, &transformpolicylabel)
	if err != nil {
		return err
	}

	d.SetId(transformpolicylabelName)

	err = readTransformpolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this transformpolicylabel but we can't read it ?? %s", transformpolicylabelName)
		return nil
	}
	return nil
}

func readTransformpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readTransformpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	transformpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading transformpolicylabel state %s", transformpolicylabelName)
	data, err := client.FindResource(service.Transformpolicylabel.Type(), transformpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing transformpolicylabel state %s", transformpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("policylabeltype", data["policylabeltype"])

	return nil

}


func deleteTransformpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTransformpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	transformpolicylabelName := d.Id()
	err := client.DeleteResource(service.Transformpolicylabel.Type(), transformpolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
