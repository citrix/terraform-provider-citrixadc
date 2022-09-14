package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/feo"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcFeoparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createFeoparameterFunc,
		Read:          readFeoparameterFunc,
		Update:        updateFeoparameterFunc,
		Delete:        deleteFeoparameterFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cssinlinethressize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"imginlinethressize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jpegqualitypercent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jsinlinethressize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createFeoparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createFeoparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	feoparameterName := resource.PrefixedUniqueId("tf-feoparameter-")

	feoparameter := make(map[string]interface{})
	if v, ok := d.GetOk("jsinlinethressize"); ok {
		feoparameter["jsinlinethressize"] = v.(int)
	}
	if v, ok := d.GetOkExists("jpegqualitypercent"); ok {
		feoparameter["jpegqualitypercent"] = v.(int)
	} 
	if v, ok := d.GetOk("imginlinethressize"); ok {
		feoparameter["imginlinethressize"] = v.(int)
	}
	if v, ok := d.GetOk("cssinlinethressize"); ok {
		feoparameter["cssinlinethressize"] = v.(int)
	}

	err := client.UpdateUnnamedResource("feoparameter", &feoparameter)
	if err != nil {
		return err
	}

	d.SetId(feoparameterName)

	err = readFeoparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this feoparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readFeoparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readFeoparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading feoparameter state")
	data, err := client.FindResource("feoparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing feoparameter state")
		d.SetId("")
		return nil
	}
	d.Set("cssinlinethressize", data["cssinlinethressize"])
	d.Set("imginlinethressize", data["imginlinethressize"])
	d.Set("jpegqualitypercent", data["jpegqualitypercent"])
	d.Set("jsinlinethressize", data["jsinlinethressize"])

	return nil

}

func updateFeoparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateFeoparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	feoparameter := feo.Feoparameter{}
	hasChange := false
	if d.HasChange("cssinlinethressize") {
		log.Printf("[DEBUG]  citrixadc-provider: Cssinlinethressize has changed for feoparameter, starting update")
		feoparameter.Cssinlinethressize = d.Get("cssinlinethressize").(int)
		hasChange = true
	}
	if d.HasChange("imginlinethressize") {
		log.Printf("[DEBUG]  citrixadc-provider: Imginlinethressize has changed for feoparameter, starting update")
		feoparameter.Imginlinethressize = d.Get("imginlinethressize").(int)
		hasChange = true
	}
	if d.HasChange("jpegqualitypercent") {
		log.Printf("[DEBUG]  citrixadc-provider: Jpegqualitypercent has changed for feoparameter, starting update")
		feoparameter.Jpegqualitypercent = d.Get("jpegqualitypercent").(int)
		hasChange = true
	}
	if d.HasChange("jsinlinethressize") {
		log.Printf("[DEBUG]  citrixadc-provider: Jsinlinethressize has changed for feoparameter, starting update")
		feoparameter.Jsinlinethressize = d.Get("jsinlinethressize").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("feoparameter", &feoparameter)
		if err != nil {
			return fmt.Errorf("Error updating feoparameter")
		}
	}
	return readFeoparameterFunc(d, meta)
}

func deleteFeoparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFeoparameterFunc")
	// feoparameter does not support DELETE operation
	d.SetId("")

	return nil
}
