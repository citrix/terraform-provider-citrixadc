package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/subscriber"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSubscriberparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSubscriberparamFunc,
		Read:          readSubscriberparamFunc,
		Update:        updateSubscriberparamFunc,
		Delete:        deleteSubscriberparamFunc,
		Schema: map[string]*schema.Schema{
			"idleaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idlettl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interfacetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6prefixlookuplist": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"keytype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSubscriberparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSubscriberparamFunc")
	client := meta.(*NetScalerNitroClient).client
	subscriberparamName := resource.PrefixedUniqueId("tf-subscriberparam-")

	subscriberparam := subscriber.Subscriberparam{
		Idleaction:           d.Get("idleaction").(string),
		Idlettl:              d.Get("idlettl").(int),
		Interfacetype:        d.Get("interfacetype").(string),
		Ipv6prefixlookuplist: toIntegerList(d.Get("ipv6prefixlookuplist").([]interface{})),
		Keytype:              d.Get("keytype").(string),
	}

	err := client.UpdateUnnamedResource("subscriberparam", &subscriberparam)
	if err != nil {
		return err
	}

	d.SetId(subscriberparamName)

	err = readSubscriberparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this subscriberparam but we can't read it ??")
		return nil
	}
	return nil
}

func readSubscriberparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSubscriberparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading subscriberparam state")
	data, err := client.FindResource("subscriberparam", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing subscriberparam state")
		d.SetId("")
		return nil
	}
	d.Set("idleaction", data["idleaction"])
	d.Set("idlettl", data["idlettl"])
	d.Set("interfacetype", data["interfacetype"])
	//d.Set("ipv6prefixlookuplist", data["ipv6prefixlookuplist"])
	d.Set("keytype", data["keytype"])

	return nil

}

func updateSubscriberparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSubscriberparamFunc")
	client := meta.(*NetScalerNitroClient).client

	subscriberparam := subscriber.Subscriberparam{}
	hasChange := false
	if d.HasChange("idleaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Idleaction has changed for subscriberparam, starting update")
		subscriberparam.Idleaction = d.Get("idleaction").(string)
		hasChange = true
	}
	if d.HasChange("idlettl") {
		log.Printf("[DEBUG]  citrixadc-provider: Idlettl has changed for subscriberparam, starting update")
		subscriberparam.Idlettl = d.Get("idlettl").(int)
		hasChange = true
	}
	if d.HasChange("interfacetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Interfacetype has changed for subscriberparam, starting update")
		subscriberparam.Interfacetype = d.Get("interfacetype").(string)
		hasChange = true
	}
	if d.HasChange("ipv6prefixlookuplist") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv6prefixlookuplist has changed for subscriberparam, starting update")
		subscriberparam.Ipv6prefixlookuplist = toIntegerList(d.Get("ipv6prefixlookuplist").([]interface{}))
		hasChange = true
	}
	if d.HasChange("keytype") {
		log.Printf("[DEBUG]  citrixadc-provider: Keytype has changed for subscriberparam, starting update")
		subscriberparam.Keytype = d.Get("keytype").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("subscriberparam", &subscriberparam)
		if err != nil {
			return fmt.Errorf("Error updating subscriberparam")
		}
	}
	return readSubscriberparamFunc(d, meta)
}

func deleteSubscriberparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSubscriberparamFunc")
	//subscriberparam does not support DELETE operation
	d.SetId("")

	return nil
}
