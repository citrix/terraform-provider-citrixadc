package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcArpparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createArpparamFunc,
		Read:          readArpparamFunc,
		Update:        updateArpparamFunc,
		Delete:        deleteArpparamFunc,
		Schema: map[string]*schema.Schema{
			"spoofvalidation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createArpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createArpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var arpparamName string
	// there is no primary key in arpparam resource. Hence generate one for terraform state maintenance
	arpparamName = resource.PrefixedUniqueId("tf-arpparam-")
	arpparam := network.Arpparam{
		Spoofvalidation: d.Get("spoofvalidation").(string),
		Timeout:         d.Get("timeout").(int),
	}

	err := client.UpdateUnnamedResource(service.Arpparam.Type(), &arpparam)
	if err != nil {
		return err
	}

	d.SetId(arpparamName)

	err = readArpparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this arpparam but we can't read it ?? %s", arpparamName)
		return nil
	}
	return nil
}

func readArpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readArpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading arpparam state")
	data, err := client.FindResource(service.Arpparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing arpparam state")
		d.SetId("")
		return nil
	}
	d.Set("spoofvalidation", data["spoofvalidation"])
	val,_ := strconv.Atoi(data["timeout"].(string))
	d.Set("timeout", val)

	return nil

}

func updateArpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateArpparamFunc")
	client := meta.(*NetScalerNitroClient).client

	arpparam := network.Arpparam{}
	hasChange := false
	if d.HasChange("spoofvalidation") {
		log.Printf("[DEBUG]  citrixadc-provider: Spoofvalidation has changed for arpparam, starting update")
		arpparam.Spoofvalidation = d.Get("spoofvalidation").(string)
		hasChange = true
	}
	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for arpparam, starting update")
		arpparam.Timeout = d.Get("timeout").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Arpparam.Type(), &arpparam)
		if err != nil {
			return fmt.Errorf("Error updating arpparam")
		}
	}
	return readArpparamFunc(d, meta)
}

func deleteArpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteArpparamFunc")

	d.SetId("")

	return nil
}
