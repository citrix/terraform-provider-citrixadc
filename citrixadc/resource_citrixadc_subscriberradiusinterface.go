package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/subscriber"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSubscriberradiusinterface() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSubscriberradiusinterfaceFunc,
		Read:          readSubscriberradiusinterfaceFunc,
		Update:        updateSubscriberradiusinterfaceFunc,
		Delete:        deleteSubscriberradiusinterfaceFunc,
		Schema: map[string]*schema.Schema{
			"listeningservice": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radiusinterimasstart": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSubscriberradiusinterfaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSubscriberradiusinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client
	subscriberradiusinterfaceName := resource.PrefixedUniqueId("tf-subscriberradiusinterface-")

	subscriberradiusinterface := subscriber.Subscriberradiusinterface{
		Listeningservice:     d.Get("listeningservice").(string),
		Radiusinterimasstart: d.Get("radiusinterimasstart").(string),
	}

	err := client.UpdateUnnamedResource("subscriberradiusinterface", &subscriberradiusinterface)
	if err != nil {
		return err
	}

	d.SetId(subscriberradiusinterfaceName)

	err = readSubscriberradiusinterfaceFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this subscriberradiusinterface but we can't read it ??")
		return nil
	}
	return nil
}

func readSubscriberradiusinterfaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSubscriberradiusinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading subscriberradiusinterface state")
	data, err := client.FindResource("subscriberradiusinterface", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing subscriberradiusinterface state")
		d.SetId("")
		return nil
	}
	d.Set("listeningservice", data["listeningservice"])
	d.Set("radiusinterimasstart", data["radiusinterimasstart"])

	return nil

}

func updateSubscriberradiusinterfaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSubscriberradiusinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client

	subscriberradiusinterface := subscriber.Subscriberradiusinterface{}
	hasChange := false
	if d.HasChange("listeningservice") {
		log.Printf("[DEBUG]  citrixadc-provider: Listeningservice has changed for subscriberradiusinterface, starting update")
		subscriberradiusinterface.Listeningservice = d.Get("listeningservice").(string)
		hasChange = true
	}
	if d.HasChange("radiusinterimasstart") {
		log.Printf("[DEBUG]  citrixadc-provider: Radiusinterimasstart has changed for subscriberradiusinterface, starting update")
		subscriberradiusinterface.Radiusinterimasstart = d.Get("radiusinterimasstart").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("subscriberradiusinterface", &subscriberradiusinterface)
		if err != nil {
			return fmt.Errorf("Error updating subscriberradiusinterface")
		}
	}
	return readSubscriberradiusinterfaceFunc(d, meta)
}

func deleteSubscriberradiusinterfaceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSubscriberradiusinterfaceFunc")
	//subscriberradiusinterface does not support DELETE operation
	d.SetId("")

	return nil
}
