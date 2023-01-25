package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAaaotpparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaaotpparameterFunc,
		Read:          readAaaotpparameterFunc,
		Update:        updateAaaotpparameterFunc,
		Delete:        deleteAaaotpparameterFunc,
		Schema: map[string]*schema.Schema{
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxotpdevices": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaaotpparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaaotpparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	aaaotpparameterName := resource.PrefixedUniqueId("tf-aaaotpparameter-")
	
	aaaotpparameter := aaa.Aaaotpparameter{
		Encryption:    d.Get("encryption").(string),
		Maxotpdevices: d.Get("maxotpdevices").(int),
	}

	err := client.UpdateUnnamedResource("aaaotpparameter", &aaaotpparameter)
	if err != nil {
		return err
	}

	d.SetId(aaaotpparameterName)

	err = readAaaotpparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaaotpparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readAaaotpparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaaotpparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading aaaotpparameter state")
	data, err := client.FindResource("aaaotpparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaaotpparameter state")
		d.SetId("")
		return nil
	}
	d.Set("encryption", data["encryption"])
	d.Set("maxotpdevices", data["maxotpdevices"])

	return nil

}

func updateAaaotpparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaaotpparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	aaaotpparameter := aaa.Aaaotpparameter{}
	hasChange := false
	if d.HasChange("encryption") {
		log.Printf("[DEBUG]  citrixadc-provider: Encryption has changed for aaaotpparameter, starting update")
		aaaotpparameter.Encryption = d.Get("encryption").(string)
		hasChange = true
	}
	if d.HasChange("maxotpdevices") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxotpdevices has changed for aaaotpparameter, starting update")
		aaaotpparameter.Maxotpdevices = d.Get("maxotpdevices").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("aaaotpparameter", &aaaotpparameter)
		if err != nil {
			return fmt.Errorf("Error updating aaaotpparameter")
		}
	}
	return readAaaotpparameterFunc(d, meta)
}

func deleteAaaotpparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaaotpparameterFunc")
	// aaaotpparameter does not support DELETE operation
	d.SetId("")

	return nil
}
