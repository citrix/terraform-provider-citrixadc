package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/reputation"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcReputationsettings() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createReputationsettingsFunc,
		Read:          readReputationsettingsFunc,
		Update:        updateReputationsettingsFunc,
		Delete:        deleteReputationsettingsFunc,
		Schema: map[string]*schema.Schema{
			"proxyport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxyserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createReputationsettingsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createReputationsettingsFunc")
	client := meta.(*NetScalerNitroClient).client
	reputationsettingsName := resource.PrefixedUniqueId("tf-reputationsettings-")

	reputationsettings := reputation.Reputationsettings{
		Proxyport:   d.Get("proxyport").(int),
		Proxyserver: d.Get("proxyserver").(string),
	}

	err := client.UpdateUnnamedResource("reputationsettings", &reputationsettings)
	if err != nil {
		return err
	}

	d.SetId(reputationsettingsName)

	err = readReputationsettingsFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this reputationsettings but we can't read it ??")
		return nil
	}
	return nil
}

func readReputationsettingsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readReputationsettingsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading reputationsettings state")
	data, err := client.FindResource("reputationsettings", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing reputationsettings state")
		d.SetId("")
		return nil
	}
	d.Set("proxyport", data["proxyport"])
	d.Set("proxyserver", data["proxyserver"])

	return nil

}

func updateReputationsettingsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateReputationsettingsFunc")
	client := meta.(*NetScalerNitroClient).client

	reputationsettings := reputation.Reputationsettings{}
	hasChange := false
	if d.HasChange("proxyport") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyport has changed for reputationsettings, starting update")
		reputationsettings.Proxyport = d.Get("proxyport").(int)
		hasChange = true
	}
	if d.HasChange("proxyserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyserver has changed for reputationsettings, starting update")
		reputationsettings.Proxyserver = d.Get("proxyserver").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("reputationsettings",  &reputationsettings)
		if err != nil {
			return fmt.Errorf("Error updating reputationsettings")
		}
	}
	return readReputationsettingsFunc(d, meta)
}

func deleteReputationsettingsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteReputationsettingsFunc")
	//reputationsettings does not support DELETE operation
	d.SetId("")

	return nil
}
