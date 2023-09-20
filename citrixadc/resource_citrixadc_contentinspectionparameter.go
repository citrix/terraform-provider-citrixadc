package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcContentinspectionparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createContentinspectionparameterFunc,
		Read:          readContentinspectionparameterFunc,
		Update:        updateContentinspectionparameterFunc,
		Delete:        deleteContentinspectionparameterFunc,
		Schema: map[string]*schema.Schema{
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createContentinspectionparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectionparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionparameterName := resource.PrefixedUniqueId("tf-contentinspectionparameter-")

	contentinspectionparameter := contentinspection.Contentinspectionparameter{
		Undefaction: d.Get("undefaction").(string),
	}

	err := client.UpdateUnnamedResource("contentinspectionparameter", &contentinspectionparameter)
	if err != nil {
		return err
	}

	d.SetId(contentinspectionparameterName)

	err = readContentinspectionparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this contentinspectionparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readContentinspectionparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectionparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectionparameter state")
	data, err := client.FindResource("contentinspectionparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectionparameter state")
		d.SetId("")
		return nil
	}
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateContentinspectionparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateContentinspectionparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	contentinspectionparameter := contentinspection.Contentinspectionparameter{}
	hasChange := false
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for contentinspectionparameter, starting update")
		contentinspectionparameter.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("contentinspectionparameter", &contentinspectionparameter)
		if err != nil {
			return fmt.Errorf("Error updating contentinspectionparameter")
		}
	}
	return readContentinspectionparameterFunc(d, meta)
}

func deleteContentinspectionparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectionparameterFunc")
	//contentinspectionparameter does not support DELETE operation
	d.SetId("")

	return nil
}
