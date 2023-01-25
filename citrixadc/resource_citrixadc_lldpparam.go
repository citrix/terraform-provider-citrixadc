package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lldp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLldpparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLldpparamFunc,
		Read:          readLldpparamFunc,
		Update:        updateLldpparamFunc,
		Delete:        deleteLldpparamFunc,
		Schema: map[string]*schema.Schema{
			"holdtimetxmult": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLldpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLldpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	lldpparamName := resource.PrefixedUniqueId("tf-lldpparam-")
	
	lldpparam := lldp.Lldpparam{
		Holdtimetxmult: d.Get("holdtimetxmult").(int),
		Mode:           d.Get("mode").(string),
		Timer:          d.Get("timer").(int),
	}

	err := client.UpdateUnnamedResource("lldpparam", &lldpparam)
	if err != nil {
		return err
	}

	d.SetId(lldpparamName)

	err = readLldpparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lldpparam but we can't read it ??")
		return nil
	}
	return nil
}

func readLldpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLldpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading lldpparam state")
	data, err := client.FindResource("lldpparam", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lldpparam state")
		d.SetId("")
		return nil
	}
	d.Set("holdtimetxmult", data["holdtimetxmult"])
	d.Set("mode", data["mode"])
	d.Set("timer", data["timer"])

	return nil

}

func updateLldpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLldpparamFunc")
	client := meta.(*NetScalerNitroClient).client

	lldpparam := lldp.Lldpparam{}
	hasChange := false
	if d.HasChange("holdtimetxmult") {
		log.Printf("[DEBUG]  citrixadc-provider: Holdtimetxmult has changed for lldpparam, starting update")
		lldpparam.Holdtimetxmult = d.Get("holdtimetxmult").(int)
		hasChange = true
	}
	if d.HasChange("mode") {
		log.Printf("[DEBUG]  citrixadc-provider: Mode has changed for lldpparam, starting update")
		lldpparam.Mode = d.Get("mode").(string)
		hasChange = true
	}
	if d.HasChange("timer") {
		log.Printf("[DEBUG]  citrixadc-provider: Timer has changed for lldpparam, starting update")
		lldpparam.Timer = d.Get("timer").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lldpparam", &lldpparam)
		if err != nil {
			return fmt.Errorf("Error updating lldpparam")
		}
	}
	return readLldpparamFunc(d, meta)
}

func deleteLldpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLldpparamFunc")
	// lldpparam does not support DELETE operation
	d.SetId("")

	return nil
}
