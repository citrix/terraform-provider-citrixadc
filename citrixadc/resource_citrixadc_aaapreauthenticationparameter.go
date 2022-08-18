package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAaapreauthenticationparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaapreauthenticationparameterFunc,
		Read:          readAaapreauthenticationparameterFunc,
		Update:        updateAaapreauthenticationparameterFunc,
		Delete:        deleteAaapreauthenticationparameterFunc,
		Schema: map[string]*schema.Schema{
			"deletefiles": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"killprocess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preauthenticationaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaapreauthenticationparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaapreauthenticationparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationparameterName := resource.PrefixedUniqueId("tf-aaapreauthenticationparameter-")
	
	aaapreauthenticationparameter := aaa.Aaapreauthenticationparameter{
		Deletefiles:             d.Get("deletefiles").(string),
		Killprocess:             d.Get("killprocess").(string),
		Preauthenticationaction: d.Get("preauthenticationaction").(string),
		Rule:                    d.Get("rule").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaapreauthenticationparameter.Type(), &aaapreauthenticationparameter)
	if err != nil {
		return err
	}

	d.SetId(aaapreauthenticationparameterName)

	err = readAaapreauthenticationparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaapreauthenticationparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readAaapreauthenticationparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaapreauthenticationparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading aaapreauthenticationparameter state")
	data, err := client.FindResource(service.Aaapreauthenticationparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaapreauthenticationparameter state")
		d.SetId("")
		return nil
	}
	d.Set("deletefiles", data["deletefiles"])
	d.Set("killprocess", data["killprocess"])
	d.Set("preauthenticationaction", data["preauthenticationaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAaapreauthenticationparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaapreauthenticationparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationparameter := aaa.Aaapreauthenticationparameter{}
	hasChange := false
	if d.HasChange("deletefiles") {
		log.Printf("[DEBUG]  citrixadc-provider: Deletefiles has changed for aaapreauthenticationparameter, starting update")
		aaapreauthenticationparameter.Deletefiles = d.Get("deletefiles").(string)
		hasChange = true
	}
	if d.HasChange("killprocess") {
		log.Printf("[DEBUG]  citrixadc-provider: Killprocess has changed for aaapreauthenticationparameter, starting update")
		aaapreauthenticationparameter.Killprocess = d.Get("killprocess").(string)
		hasChange = true
	}
	if d.HasChange("preauthenticationaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Preauthenticationaction has changed for aaapreauthenticationparameter, starting update")
		aaapreauthenticationparameter.Preauthenticationaction = d.Get("preauthenticationaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for aaapreauthenticationparameter, starting update")
		aaapreauthenticationparameter.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaapreauthenticationparameter.Type(), &aaapreauthenticationparameter)
		if err != nil {
			return fmt.Errorf("Error updating aaapreauthenticationparameter")
		}
	}
	return readAaapreauthenticationparameterFunc(d, meta)
}

func deleteAaapreauthenticationparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaapreauthenticationparameterFunc")
	// aaapreauthenticationparameter does not suppor DELETE operation
	d.SetId("")

	return nil
}
