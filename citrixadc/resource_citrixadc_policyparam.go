package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcPolicyparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPolicyparamFunc,
		Read:          readPolicyparamFunc,
		Update:        updatePolicyparamFunc,
		Delete:        deletePolicyparamFunc, // Thought policyparam resource does not have a DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createPolicyparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicyparamFunc")
	client := meta.(*NetScalerNitroClient).client
	var policyparamName string

	// there is no primary key in policyparam resource. Hence generate one for terraform state maintenance
	policyparamName = resource.PrefixedUniqueId("tf-policyparam-")

	policyparam := policy.Policyparam{
		Timeout: d.Get("timeout").(int),
	}

	err := client.UpdateUnnamedResource("policyparam", &policyparam)
	if err != nil {
		return fmt.Errorf("Error updating policyparam")
	}

	d.SetId(policyparamName)

	err = readPolicyparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this policyparam but we can't read it ??")
		return nil
	}
	return nil
}

func readPolicyparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicyparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading policyparam state")
	data, err := client.FindResource("policyparam", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policyparam state")
		d.SetId("")
		return nil
	}
	d.Set("timeout", data["timeout"])

	return nil

}

func updatePolicyparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePolicyparamFunc")
	client := meta.(*NetScalerNitroClient).client

	policyparam := policy.Policyparam{}
	hasChange := false

	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for policyparam, starting update")
		policyparam.Timeout = d.Get("timeout").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("policyparam", &policyparam)
		if err != nil {
			return fmt.Errorf("Error updating policyparam %s", err.Error())
		}
	}
	return readPolicyparamFunc(d, meta)
}

func deletePolicyparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicyparamFunc")
	// policyparam does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")
	return nil
}
