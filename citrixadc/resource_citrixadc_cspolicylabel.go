package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cs"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcCspolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCspolicylabelFunc,
		Read:          readCspolicylabelFunc,
		Delete:        deleteCspolicylabelFunc,
		Schema: map[string]*schema.Schema{
			"cspolicylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createCspolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	var cspolicylabelName string
	if v, ok := d.GetOk("labelname"); ok {
		cspolicylabelName = v.(string)
	} else {
		cspolicylabelName = resource.PrefixedUniqueId("tf-cspolicylabel-")
		d.Set("labelname", cspolicylabelName)
	}
	cspolicylabel := cs.Cspolicylabel{
		Cspolicylabeltype: d.Get("cspolicylabeltype").(string),
		Labelname:         d.Get("labelname").(string),
	}

	_, err := client.AddResource(service.Cspolicylabel.Type(), cspolicylabelName, &cspolicylabel)
	if err != nil {
		return err
	}

	d.SetId(cspolicylabelName)

	err = readCspolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cspolicylabel but we can't read it ?? %s", cspolicylabelName)
		return nil
	}
	return nil
}

func readCspolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cspolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cspolicylabel state %s", cspolicylabelName)
	data, err := client.FindResource(service.Cspolicylabel.Type(), cspolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cspolicylabel state %s", cspolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("cspolicylabeltype", data["cspolicylabeltype"])

	return nil

}

func deleteCspolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCspolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	cspolicylabelName := d.Id()
	err := client.DeleteResource(service.Cspolicylabel.Type(), cspolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
