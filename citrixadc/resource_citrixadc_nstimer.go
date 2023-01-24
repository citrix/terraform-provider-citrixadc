package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNstimer() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNstimerFunc,
		Read:          readNstimerFunc,
		Update:        updateNstimerFunc,
		Delete:        deleteNstimerFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNstimerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstimerFunc")
	client := meta.(*NetScalerNitroClient).client
	nstimerName := d.Get("name").(string)
	nstimer := ns.Nstimer{
		Comment:  d.Get("comment").(string),
		Interval: d.Get("interval").(int),
		Name:     d.Get("name").(string),
		Newname:  d.Get("newname").(string),
		Unit:     d.Get("unit").(string),
	}

	_, err := client.AddResource(service.Nstimer.Type(), nstimerName, &nstimer)
	if err != nil {
		return err
	}

	d.SetId(nstimerName)

	err = readNstimerFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nstimer but we can't read it ?? %s", nstimerName)
		return nil
	}
	return nil
}

func readNstimerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstimerFunc")
	client := meta.(*NetScalerNitroClient).client
	nstimerName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nstimer state %s", nstimerName)
	data, err := client.FindResource(service.Nstimer.Type(), nstimerName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nstimer state %s", nstimerName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("interval", data["interval"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("unit", data["unit"])

	return nil

}

func updateNstimerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNstimerFunc")
	client := meta.(*NetScalerNitroClient).client
	nstimerName := d.Get("name").(string)

	nstimer := ns.Nstimer{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for nstimer %s, starting update", nstimerName)
		nstimer.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("interval") {
		log.Printf("[DEBUG]  citrixadc-provider: Interval has changed for nstimer %s, starting update", nstimerName)
		nstimer.Interval = d.Get("interval").(int)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for nstimer %s, starting update", nstimerName)
		nstimer.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("unit") {
		log.Printf("[DEBUG]  citrixadc-provider: Unit has changed for nstimer %s, starting update", nstimerName)
		nstimer.Unit = d.Get("unit").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nstimer.Type(), nstimerName, &nstimer)
		if err != nil {
			return fmt.Errorf("Error updating nstimer %s", nstimerName)
		}
	}
	return readNstimerFunc(d, meta)
}

func deleteNstimerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstimerFunc")
	client := meta.(*NetScalerNitroClient).client
	nstimerName := d.Id()
	err := client.DeleteResource(service.Nstimer.Type(), nstimerName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
