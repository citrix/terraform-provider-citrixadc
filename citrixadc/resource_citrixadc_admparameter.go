package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/adm"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAdmparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAdmparameterFunc,
		Read:          readAdmparameterFunc,
		Update:        updateAdmparameterFunc,
		Delete:        deleteAdmparameterFunc,
		Schema: map[string]*schema.Schema{
			"admserviceconnect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAdmparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAdmparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	admparameterName := resource.PrefixedUniqueId("tf-admparameter-")

	admparameter := adm.Admparameter{
		Admserviceconnect: d.Get("admserviceconnect").(string),
	}

	err := client.UpdateUnnamedResource("admparameter", &admparameter)
	if err != nil {
		return err
	}

	d.SetId(admparameterName)

	err = readAdmparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this admparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readAdmparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAdmparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading admparameter state")
	data, err := client.FindResource("admparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing admparameter state")
		d.SetId("")
		return nil
	}
	d.Set("admserviceconnect", data["admserviceconnect"])

	return nil

}

func updateAdmparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAdmparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	admparameter := adm.Admparameter{}
	hasChange := false
	if d.HasChange("admserviceconnect") {
		log.Printf("[DEBUG]  citrixadc-provider: Admserviceconnect has changed for admparameter, starting update")
		admparameter.Admserviceconnect = d.Get("admserviceconnect").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("admparameter", &admparameter)
		if err != nil {
			return fmt.Errorf("Error updating admparameter")
		}
	}
	return readAdmparameterFunc(d, meta)
}

func deleteAdmparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAdmparameterFunc")
	// admparameter does not support DELETE operation
	d.SetId("")

	return nil
}
