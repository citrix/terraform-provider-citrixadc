package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLsnappsattributes() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnappsattributesFunc,
		Read:          readLsnappsattributesFunc,
		Update:        updateLsnappsattributesFunc,
		Delete:        deleteLsnappsattributesFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"sessiontimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"transportprotocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnappsattributesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnappsattributesFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsattributesName := d.Get("name").(string)
	lsnappsattributes := lsn.Lsnappsattributes{
		Name:              d.Get("name").(string),
		Port:              d.Get("port").(string),
		Sessiontimeout:    d.Get("sessiontimeout").(int),
		Transportprotocol: d.Get("transportprotocol").(string),
	}

	_, err := client.AddResource("lsnappsattributes", lsnappsattributesName, &lsnappsattributes)
	if err != nil {
		return err
	}

	d.SetId(lsnappsattributesName)

	err = readLsnappsattributesFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnappsattributes but we can't read it ?? %s", lsnappsattributesName)
		return nil
	}
	return nil
}

func readLsnappsattributesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnappsattributesFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsattributesName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnappsattributes state %s", lsnappsattributesName)
	data, err := client.FindResource("lsnappsattributes", lsnappsattributesName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnappsattributes state %s", lsnappsattributesName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("port", data["port"])
	d.Set("sessiontimeout", data["sessiontimeout"])
	d.Set("transportprotocol", data["transportprotocol"])

	return nil

}

func updateLsnappsattributesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsnappsattributesFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsattributesName := d.Get("name").(string)

	lsnappsattributes := lsn.Lsnappsattributes{
		Name: d.Get("name").(string),
	}
	hasChange := false

	if d.HasChange("sessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessiontimeout has changed for lsnappsattributes %s, starting update", lsnappsattributesName)
		lsnappsattributes.Sessiontimeout = d.Get("sessiontimeout").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lsnappsattributes", &lsnappsattributes)
		if err != nil {
			return fmt.Errorf("Error updating lsnappsattributes %s", lsnappsattributesName)
		}
	}
	return readLsnappsattributesFunc(d, meta)
}

func deleteLsnappsattributesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnappsattributesFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsattributesName := d.Id()
	err := client.DeleteResource("lsnappsattributes", lsnappsattributesName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
