package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcLsnclient() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnclientFunc,
		Read:          readLsnclientFunc,
		Delete:        deleteLsnclientFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"clientname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnclientFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnclientFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnclientName := d.Get("clientname").(string)
	lsnclient := lsn.Lsnclient{
		Clientname: d.Get("clientname").(string),
	}

	_, err := client.AddResource("lsnclient", lsnclientName, &lsnclient)
	if err != nil {
		return err
	}

	d.SetId(lsnclientName)

	err = readLsnclientFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnclient but we can't read it ?? %s", lsnclientName)
		return nil
	}
	return nil
}

func readLsnclientFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnclientFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnclientName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnclient state %s", lsnclientName)
	data, err := client.FindResource("lsnclient", lsnclientName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnclient state %s", lsnclientName)
		d.SetId("")
		return nil
	}
	d.Set("clientname", data["clientname"])

	return nil

}

func deleteLsnclientFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnclientFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnclientName := d.Id()
	err := client.DeleteResource("lsnclient", lsnclientName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
