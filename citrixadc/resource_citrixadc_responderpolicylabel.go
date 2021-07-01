package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/responder"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcResponderpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createResponderpolicylabelFunc,
		Read:          readResponderpolicylabelFunc,
		Delete:        deleteResponderpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"policylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createResponderpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createResponderpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicylabelName := d.Get("labelname").(string)
	responderpolicylabel := responder.Responderpolicylabel{
		Comment:         d.Get("comment").(string),
		Labelname:       d.Get("labelname").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	_, err := client.AddResource(service.Responderpolicylabel.Type(), responderpolicylabelName, &responderpolicylabel)
	if err != nil {
		return err
	}

	d.SetId(responderpolicylabelName)

	err = readResponderpolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this responderpolicylabel but we can't read it ?? %s", responderpolicylabelName)
		return nil
	}
	return nil
}

func readResponderpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readResponderpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading responderpolicylabel state %s", responderpolicylabelName)
	data, err := client.FindResource(service.Responderpolicylabel.Type(), responderpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing responderpolicylabel state %s", responderpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("labelname", data["labelname"])
	d.Set("policylabeltype", data["policylabeltype"])

	return nil

}

func deleteResponderpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteResponderpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicylabelName := d.Id()
	err := client.DeleteResource(service.Responderpolicylabel.Type(), responderpolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
