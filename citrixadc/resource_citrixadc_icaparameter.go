package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcIcaparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIcaparameterFunc,
		Read:          readIcaparameterFunc,
		Update:        updateIcaparameterFunc,
		Delete:        deleteIcaparameterFunc,
		Schema: map[string]*schema.Schema{
			"enablesronhafailover": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hdxinsightnonnsap": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"edtpmtuddf": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"edtpmtuddftimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l7latencyfrequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIcaparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIcaparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	icaparameterName := resource.PrefixedUniqueId("tf-icaparameter-")

	icaparameter := ica.Icaparameter{
		Enablesronhafailover: d.Get("enablesronhafailover").(string),
		Hdxinsightnonnsap:    d.Get("hdxinsightnonnsap").(string),
		L7latencyfrequency:   d.Get("l7latencyfrequency").(int),
		Edtpmtuddf:           d.Get("edtpmtuddf").(string),
		Edtpmtuddftimeout:    d.Get("edtpmtuddftimeout").(int),
	}

	err := client.UpdateUnnamedResource("icaparameter", &icaparameter)
	if err != nil {
		return err
	}

	d.SetId(icaparameterName)

	err = readIcaparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this icaparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readIcaparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIcaparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading icaparameter state")
	data, err := client.FindResource("icaparameter", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing icaparameter state")
		d.SetId("")
		return nil
	}
	d.Set("enablesronhafailover", data["enablesronhafailover"])
	d.Set("hdxinsightnonnsap", data["hdxinsightnonnsap"])
	d.Set("l7latencyfrequency", data["l7latencyfrequency"])
	d.Set("edtpmtuddf", data["edtpmtuddf"])
	d.Set("edtpmtuddftimeout", data["edtpmtuddftimeout"])

	return nil

}

func updateIcaparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIcaparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	icaparameter := ica.Icaparameter{}
	hasChange := false
	if d.HasChange("enablesronhafailover") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablesronhafailover has changed for icaparameter, starting update")
		icaparameter.Enablesronhafailover = d.Get("enablesronhafailover").(string)
		hasChange = true
	}
	if d.HasChange("edtpmtuddf") {
		log.Printf("[DEBUG]  citrixadc-provider: Edtpmtuddf has changed for icaparameter, starting update")
		icaparameter.Edtpmtuddf = d.Get("edtpmtuddf").(string)
		hasChange = true
	}
	if d.HasChange("edtpmtuddftimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Edtpmtuddftimeout has changed for icaparameter, starting update")
		icaparameter.Edtpmtuddftimeout = d.Get("edtpmtuddftimeout").(int)
		hasChange = true
	}
	if d.HasChange("hdxinsightnonnsap") {
		log.Printf("[DEBUG]  citrixadc-provider: Hdxinsightnonnsap has changed for icaparameter, starting update")
		icaparameter.Hdxinsightnonnsap = d.Get("hdxinsightnonnsap").(string)
		hasChange = true
	}
	if d.HasChange("l7latencyfrequency") {
		log.Printf("[DEBUG]  citrixadc-provider: L7latencyfrequency has changed for icaparameter, starting update")
		icaparameter.L7latencyfrequency = d.Get("l7latencyfrequency").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("icaparameter", &icaparameter)
		if err != nil {
			return fmt.Errorf("Error updating icaparameter")
		}
	}
	return readIcaparameterFunc(d, meta)
}

func deleteIcaparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIcaparameterFunc")
	// icaparameter does not support DELETE operation
	d.SetId("")

	return nil
}
