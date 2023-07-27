package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLsnparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnparameterFunc,
		Read:          readLsnparameterFunc,
		Update:        updateLsnparameterFunc,
		Delete:        deleteLsnparameterFunc,
		Schema: map[string]*schema.Schema{
			"sessionsync": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscrsessionremoval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsnparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnparameterName := resource.PrefixedUniqueId("tf-lsnparameter-")

	lsnparameter := lsn.Lsnparameter{
		Sessionsync:          d.Get("sessionsync").(string),
		Subscrsessionremoval: d.Get("subscrsessionremoval").(string),
	}

	err := client.UpdateUnnamedResource("lsnparameter", &lsnparameter)
	if err != nil {
		return err
	}

	d.SetId(lsnparameterName)

	err = readLsnparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readLsnparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnparameter state")
	data, err := client.FindResource("lsnparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnparameter state")
		d.SetId("")
		return nil
	}
	d.Set("memlimit", data["memlimit"])
	d.Set("sessionsync", data["sessionsync"])
	d.Set("subscrsessionremoval", data["subscrsessionremoval"])

	return nil

}

func updateLsnparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsnparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	lsnparameter := lsn.Lsnparameter{}
	hasChange := false
	if d.HasChange("memlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Memlimit has changed for lsnparameter, starting update")
		lsnparameter.Memlimit = d.Get("memlimit").(int)
		hasChange = true
	}
	if d.HasChange("sessionsync") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionsync has changed for lsnparameter, starting update")
		lsnparameter.Sessionsync = d.Get("sessionsync").(string)
		hasChange = true
	}
	if d.HasChange("subscrsessionremoval") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscrsessionremoval has changed for lsnparameter, starting update")
		lsnparameter.Subscrsessionremoval = d.Get("subscrsessionremoval").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lsnparameter", &lsnparameter)
		if err != nil {
			return fmt.Errorf("Error updating lsnparameter")
		}
	}
	return readLsnparameterFunc(d, meta)
}

func deleteLsnparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnparameterFunc")
	//lsnparameter does not support DELETE operation
	d.SetId("")

	return nil
}
