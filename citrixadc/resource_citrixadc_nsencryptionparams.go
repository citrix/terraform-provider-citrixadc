package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsencryptionparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsencryptionparamsFunc,
		Read:          readNsencryptionparamsFunc,
		Update:        updateNsencryptionparamsFunc,
		Delete:        deleteNsencryptionparamsFunc,
		Schema: map[string]*schema.Schema{
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"keyvalue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createNsencryptionparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsencryptionparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	nsencryptionparamsName := resource.PrefixedUniqueId("tf-nsencryptionparams-")
	nsencryptionparams := ns.Nsencryptionparams{
		Keyvalue: d.Get("keyvalue").(string),
		Method:   d.Get("method").(string),
	}

	err := client.UpdateUnnamedResource(service.Nsencryptionparams.Type(), &nsencryptionparams)
	if err != nil {
		return err
	}

	d.SetId(nsencryptionparamsName)

	err = readNsencryptionparamsFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsencryptionparams but we can't read it ??", )
		return nil
	}
	return nil
}

func readNsencryptionparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsencryptionparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nsencryptionparams state")
	data, err := client.FindResource(service.Nsencryptionparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsencryptionparams state")
		d.SetId("")
		return nil
	}
	d.Set("keyvalue", data["keyvalue"])
	d.Set("method", data["method"])

	return nil

}

func updateNsencryptionparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsencryptionparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	nsencryptionparams := ns.Nsencryptionparams{}
	hasChange := false
	
	if d.HasChange("keyvalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Keyvalue has changed for nsencryptionparams, starting update")
		nsencryptionparams.Keyvalue = d.Get("keyvalue").(string)
		hasChange = true
	}
	if d.HasChange("method") {
		log.Printf("[DEBUG]  citrixadc-provider: Method has changed for nsencryptionparams, starting update")
		nsencryptionparams.Method = d.Get("method").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nsencryptionparams.Type(), &nsencryptionparams)
		if err != nil {
			return fmt.Errorf("Error updating nsencryptionparams")
		}
	}
	return readNsencryptionparamsFunc(d, meta)
}

func deleteNsencryptionparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsencryptionparamsFunc")
	// nsencryption does not support delete operation
	d.SetId("")

	return nil
}
