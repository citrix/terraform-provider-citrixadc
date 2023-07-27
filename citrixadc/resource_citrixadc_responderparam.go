package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcResponderparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createResponderparamFunc,
		Read:          readResponderparamFunc,
		Update:        updateResponderparamFunc,
		Delete:        deleteResponderparamFunc, // Thought responderparam resource does not have a DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createResponderparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createResponderparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var responderparamName string

	// there is no primary key in rewriteparam resource. Hence generate one for terraform state maintenance
	responderparamName = resource.PrefixedUniqueId("tf-responderparam-")

	responderparam := responder.Responderparam{
		Timeout:     d.Get("timeout").(int),
		Undefaction: d.Get("undefaction").(string),
	}

	err := client.UpdateUnnamedResource(service.Responderparam.Type(), &responderparam)
	if err != nil {
		return fmt.Errorf("Error updating responderparam")
	}

	d.SetId(responderparamName)

	err = readResponderparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this responderparam but we can't read it ??")
		return nil
	}
	return nil
}

func readResponderparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readResponderparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading responderparam state")
	data, err := client.FindResource(service.Responderparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing responderparam state")
		d.SetId("")
		return nil
	}
	d.Set("timeout", data["timeout"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateResponderparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateResponderparamFunc")
	client := meta.(*NetScalerNitroClient).client

	responderparam := responder.Responderparam{}
	hasChange := false

	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for responderparam, starting update")
		responderparam.Timeout = d.Get("timeout").(int)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for responderparam, starting update")
		responderparam.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Responderparam.Type(), &responderparam)
		if err != nil {
			return fmt.Errorf("Error updating responderparam %s", err.Error())
		}
	}
	return readResponderparamFunc(d, meta)
}

func deleteResponderparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteResponderparamFunc")
	// responderparam does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")
	return nil
}
