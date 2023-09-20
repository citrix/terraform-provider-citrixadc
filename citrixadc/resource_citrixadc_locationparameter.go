package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLocationparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLocationparameterFunc,
		Read:          readLocationparameterFunc,
		Update:        updateLocationparameterFunc,
		Delete:        deleteLocationparameterFunc,
		Schema: map[string]*schema.Schema{
			"context": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"matchwildcardtoany": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"q1label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"q2label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"q3label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"q4label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"q5label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"q6label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLocationparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLocationparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	var locationparameterName string

	// there is no primary key in LOCATIONPARAMETER resource. Hence generate one for terraform state maintenance
	locationparameterName = resource.PrefixedUniqueId("tf-locationparameter-")
	locationparameter := basic.Locationparameter{
		Context:            d.Get("context").(string),
		Matchwildcardtoany: d.Get("matchwildcardtoany").(string),
		Q1label:            d.Get("q1label").(string),
		Q2label:            d.Get("q2label").(string),
		Q3label:            d.Get("q3label").(string),
		Q4label:            d.Get("q4label").(string),
		Q5label:            d.Get("q5label").(string),
		Q6label:            d.Get("q6label").(string),
	}

	err := client.UpdateUnnamedResource(service.Locationparameter.Type(), &locationparameter)
	if err != nil {
		return fmt.Errorf("Error updating locationparameter")
	}

	d.SetId(locationparameterName)

	err = readLocationparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this locationparameter but we can't read it ?? %s", locationparameterName)
		return nil
	}
	return nil
}

func readLocationparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLocationparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading locationparameter state")
	data, err := client.FindResource(service.Locationparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing locationparameter state ")
		d.SetId("")
		return nil
	}
	d.Set("context", data["context"])
	d.Set("matchwildcardtoany", data["matchwildcardtoany"])
	d.Set("q1label", data["q1label"])
	d.Set("q2label", data["q2label"])
	d.Set("q3label", data["q3label"])
	d.Set("q4label", data["q4label"])
	d.Set("q5label", data["q5label"])
	d.Set("q6label", data["q6label"])

	return nil

}

func updateLocationparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLocationparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	locationparameter := basic.Locationparameter{}

	hasChange := false
	if d.HasChange("context") {
		log.Printf("[DEBUG]  citrixadc-provider: Context has changed for locationparameter , starting update")
		locationparameter.Context = d.Get("context").(string)
		hasChange = true
	}
	if d.HasChange("matchwildcardtoany") {
		log.Printf("[DEBUG]  citrixadc-provider: Matchwildcardtoany has changed for locationparameter , starting update")
		locationparameter.Matchwildcardtoany = d.Get("matchwildcardtoany").(string)
		hasChange = true
	}
	if d.HasChange("q1label") {
		log.Printf("[DEBUG]  citrixadc-provider: Q1label has changed for locationparameter , starting update")
		locationparameter.Q1label = d.Get("q1label").(string)
		hasChange = true
	}
	if d.HasChange("q2label") {
		log.Printf("[DEBUG]  citrixadc-provider: Q2label has changed for locationparameter , starting update")
		locationparameter.Q2label = d.Get("q2label").(string)
		hasChange = true
	}
	if d.HasChange("q3label") {
		log.Printf("[DEBUG]  citrixadc-provider: Q3label has changed for locationparameter , starting update")
		locationparameter.Q3label = d.Get("q3label").(string)
		hasChange = true
	}
	if d.HasChange("q4label") {
		log.Printf("[DEBUG]  citrixadc-provider: Q4label has changed for locationparameter , starting update")
		locationparameter.Q4label = d.Get("q4label").(string)
		hasChange = true
	}
	if d.HasChange("q5label") {
		log.Printf("[DEBUG]  citrixadc-provider: Q5label has changed for locationparameter , starting update")
		locationparameter.Q5label = d.Get("q5label").(string)
		hasChange = true
	}
	if d.HasChange("q6label") {
		log.Printf("[DEBUG]  citrixadc-provider: Q6label has changed for locationparameter , starting update")
		locationparameter.Q6label = d.Get("q6label").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Locationparameter.Type(), &locationparameter)
		if err != nil {
			return fmt.Errorf("Error updating locationparameter: %s", err.Error())
		}
	}
	return readLocationparameterFunc(d, meta)
}

func deleteLocationparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLocationparameterFunc")
	// locationparameter do not have DELETE operation, but this function is required to set the ID to ""

	d.SetId("")

	return nil
}
