package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcLsnip6profile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnip6profileFunc,
		Read:          readLsnip6profileFunc,
		Delete:        deleteLsnip6profileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"natprefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"network6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnip6profileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnip6profileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnip6profileName := d.Get("name").(string)
	lsnip6profile := lsn.Lsnip6profile{
		Name:      d.Get("name").(string),
		Natprefix: d.Get("natprefix").(string),
		Network6:  d.Get("network6").(string),
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource("lsnip6profile", lsnip6profileName, &lsnip6profile)
	if err != nil {
		return err
	}

	d.SetId(lsnip6profileName)

	err = readLsnip6profileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnip6profile but we can't read it ?? %s", lsnip6profileName)
		return nil
	}
	return nil
}

func readLsnip6profileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnip6profileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnip6profileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnip6profile state %s", lsnip6profileName)
	data, err := client.FindResource("lsnip6profile", lsnip6profileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnip6profile state %s", lsnip6profileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("network6", data["network6"])
	d.Set("type", data["type"])

	return nil

}

func deleteLsnip6profileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnip6profileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnip6profileName := d.Id()
	err := client.DeleteResource("lsnip6profile", lsnip6profileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
