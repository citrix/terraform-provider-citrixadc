package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsconsoleloginprompt() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsconsoleloginpromptFunc,
		Read:          readNsconsoleloginpromptFunc,
		Update:        updateNsconsoleloginpromptFunc,
		Delete:        deleteNsconsoleloginpromptFunc,
		Schema: map[string]*schema.Schema{
			"promptstring": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createNsconsoleloginpromptFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsconsoleloginpromptFunc")
	client := meta.(*NetScalerNitroClient).client
	var nsconsoleloginpromptName string
	// there is no primary key in nsconsoleloginprompt resource. Hence generate one for terraform state maintenance
	nsconsoleloginpromptName = resource.PrefixedUniqueId("tf-nsconsoleloginprompt-")
	nsconsoleloginprompt := ns.Nsconsoleloginprompt{
		Promptstring: d.Get("promptstring").(string),
	}

	err := client.UpdateUnnamedResource(service.Nsconsoleloginprompt.Type(), &nsconsoleloginprompt)
	if err != nil {
		return err
	}

	d.SetId(nsconsoleloginpromptName)

	err = readNsconsoleloginpromptFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsconsoleloginprompt but we can't read it ??")
		return nil
	}
	return nil
}

func readNsconsoleloginpromptFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsconsoleloginpromptFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nsconsoleloginprompt state")
	data, err := client.FindResource(service.Nsconsoleloginprompt.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsconsoleloginprompt state")
		d.SetId("")
		return nil
	}
	d.Set("promptstring", data["promptstring"])

	return nil

}

func updateNsconsoleloginpromptFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsconsoleloginpromptFunc")
	client := meta.(*NetScalerNitroClient).client

	nsconsoleloginprompt := ns.Nsconsoleloginprompt{}
	hasChange := false
	if d.HasChange("promptstring") {
		log.Printf("[DEBUG]  citrixadc-provider: Promptstring has changed for nsconsoleloginprompt, starting update")
		nsconsoleloginprompt.Promptstring = d.Get("promptstring").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nsconsoleloginprompt.Type(), &nsconsoleloginprompt)
		if err != nil {
			return fmt.Errorf("Error updating nsconsoleloginprompt")
		}
	}
	return readNsconsoleloginpromptFunc(d, meta)
}

func deleteNsconsoleloginpromptFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsconsoleloginpromptFunc")
	
	// nsconsoleloginprompt do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
