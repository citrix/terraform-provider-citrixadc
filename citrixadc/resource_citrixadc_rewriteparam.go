package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/rewrite"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcRewriteparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRewriteparamFunc,
		Read:          readRewriteparamFunc,
		Update:        updateRewriteparamFunc,
		Delete:        deleteRewriteparamFunc, // Thought rewriteparam resource does not have a DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
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

func createRewriteparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRewriteparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var rewriteparamName string

	// there is no primary key in rewriteparam resource. Hence generate one for terraform state maintenance
	rewriteparamName = resource.PrefixedUniqueId("tf-rewriteparam-")

	rewriteparam := rewrite.Rewriteparam{
		Timeout:     d.Get("timeout").(int),
		Undefaction: d.Get("undefaction").(string),
	}

	err := client.UpdateUnnamedResource(service.Rewriteparam.Type(), &rewriteparam)
	if err != nil {
		return fmt.Errorf("Error updating rewriteparam")
	}

	d.SetId(rewriteparamName)

	err = readRewriteparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rewriteparam but we can't read it ??")
		return nil
	}
	return nil
}

func readRewriteparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRewriteparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading rewriteparam state")
	data, err := client.FindResource(service.Rewriteparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rewriteparam state")
		d.SetId("")
		return nil
	}
	d.Set("timeout", data["timeout"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateRewriteparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRewriteparamFunc")
	client := meta.(*NetScalerNitroClient).client

	rewriteparam := rewrite.Rewriteparam{}
	hasChange := false

	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for rewriteparam, starting update")
		rewriteparam.Timeout = d.Get("timeout").(int)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for rewriteparam, starting update")
		rewriteparam.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Rewriteparam.Type(), &rewriteparam)
		if err != nil {
			return fmt.Errorf("Error updating rewriteparam %s", err.Error())
		}
	}
	return readRewriteparamFunc(d, meta)
}

func deleteRewriteparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewriteparamFunc")
	// rewriteparam does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")
	return nil
}
