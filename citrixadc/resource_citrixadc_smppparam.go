package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/smpp"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSmppparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSmppparamFunc,
		Read:          readSmppparamFunc,
		Update:        updateSmppparamFunc,
		Delete:        deleteSmppparamFunc,
		Schema: map[string]*schema.Schema{
			"addrnpi": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"addrrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"addrton": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"clientmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"msgqueue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"msgqueuesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSmppparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSmppparamFunc")
	client := meta.(*NetScalerNitroClient).client
	smppparamName := resource.PrefixedUniqueId("tf-smppparam-")

	smppparam := smpp.Smppparam{
		Addrnpi:      d.Get("addrnpi").(int),
		Addrrange:    d.Get("addrrange").(string),
		Addrton:      d.Get("addrton").(int),
		Clientmode:   d.Get("clientmode").(string),
		Msgqueue:     d.Get("msgqueue").(string),
		Msgqueuesize: d.Get("msgqueuesize").(int),
	}

	err := client.UpdateUnnamedResource("smppparam", &smppparam)
	if err != nil {
		return err
	}

	d.SetId(smppparamName)

	err = readSmppparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this smppparam but we can't read it ??")
		return nil
	}
	return nil
}

func readSmppparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSmppparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading smppparam state")
	data, err := client.FindResource("smppparam", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing smppparam state")
		d.SetId("")
		return nil
	}
	d.Set("addrnpi", data["addrnpi"])
	d.Set("addrrange", data["addrrange"])
	d.Set("addrton", data["addrton"])
	d.Set("clientmode", data["clientmode"])
	d.Set("msgqueue", data["msgqueue"])
	d.Set("msgqueuesize", data["msgqueuesize"])

	return nil

}

func updateSmppparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSmppparamFunc")
	client := meta.(*NetScalerNitroClient).client

	smppparam := smpp.Smppparam{}
	hasChange := false
	if d.HasChange("addrnpi") {
		log.Printf("[DEBUG]  citrixadc-provider: Addrnpi has changed for smppparam, starting update")
		smppparam.Addrnpi = d.Get("addrnpi").(int)
		hasChange = true
	}
	if d.HasChange("addrrange") {
		log.Printf("[DEBUG]  citrixadc-provider: Addrrange has changed for smppparam, starting update")
		smppparam.Addrrange = d.Get("addrrange").(string)
		hasChange = true
	}
	if d.HasChange("addrton") {
		log.Printf("[DEBUG]  citrixadc-provider: Addrton has changed for smppparam, starting update")
		smppparam.Addrton = d.Get("addrton").(int)
		hasChange = true
	}
	if d.HasChange("clientmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientmode has changed for smppparam, starting update")
		smppparam.Clientmode = d.Get("clientmode").(string)
		hasChange = true
	}
	if d.HasChange("msgqueue") {
		log.Printf("[DEBUG]  citrixadc-provider: Msgqueue has changed for smppparam, starting update")
		smppparam.Msgqueue = d.Get("msgqueue").(string)
		hasChange = true
	}
	if d.HasChange("msgqueuesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Msgqueuesize has changed for smppparam, starting update")
		smppparam.Msgqueuesize = d.Get("msgqueuesize").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("smppparam", &smppparam)
		if err != nil {
			return fmt.Errorf("Error updating smppparam")
		}
	}
	return readSmppparamFunc(d, meta)
}

func deleteSmppparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSmppparamFunc")
	//smppparam does not support DELETE operation
	d.SetId("")

	return nil
}
