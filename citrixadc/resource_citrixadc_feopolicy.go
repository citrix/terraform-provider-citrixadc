package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/feo"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcFeopolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createFeopolicyFunc,
		Read:          readFeopolicyFunc,
		Update:        updateFeopolicyFunc,
		Delete:        deleteFeopolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createFeopolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createFeopolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	feopolicyName := d.Get("name").(string)
	feopolicy := feo.Feopolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource("feopolicy", feopolicyName, &feopolicy)
	if err != nil {
		return err
	}

	d.SetId(feopolicyName)

	err = readFeopolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this feopolicy but we can't read it ?? %s", feopolicyName)
		return nil
	}
	return nil
}

func readFeopolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readFeopolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	feopolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading feopolicy state %s", feopolicyName)
	data, err := client.FindResource("feopolicy", feopolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing feopolicy state %s", feopolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateFeopolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateFeopolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	feopolicyName := d.Get("name").(string)

	feopolicy := feo.Feopolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for feopolicy %s, starting update", feopolicyName)
		feopolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for feopolicy %s, starting update", feopolicyName)
		feopolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("feopolicy", &feopolicy)
		if err != nil {
			return fmt.Errorf("Error updating feopolicy %s", feopolicyName)
		}
	}
	return readFeopolicyFunc(d, meta)
}

func deleteFeopolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFeopolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	feopolicyName := d.Id()
	err := client.DeleteResource("feopolicy", feopolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
