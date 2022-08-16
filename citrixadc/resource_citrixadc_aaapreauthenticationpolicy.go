package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAaapreauthenticationpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaapreauthenticationpolicyFunc,
		Read:          readAaapreauthenticationpolicyFunc,
		Update:        updateAaapreauthenticationpolicyFunc,
		Delete:        deleteAaapreauthenticationpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reqaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaapreauthenticationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaapreauthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationpolicyName  := d.Get("name").(string)


	aaapreauthenticationpolicy := aaa.Aaapreauthenticationpolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Aaapreauthenticationpolicy.Type(), aaapreauthenticationpolicyName, &aaapreauthenticationpolicy)
	if err != nil {
		return err
	}

	d.SetId(aaapreauthenticationpolicyName)

	err = readAaapreauthenticationpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaapreauthenticationpolicy but we can't read it ?? %s", aaapreauthenticationpolicyName)
		return nil
	}
	return nil
}

func readAaapreauthenticationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaapreauthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading aaapreauthenticationpolicy state %s", aaapreauthenticationpolicyName)
	data, err := client.FindResource(service.Aaapreauthenticationpolicy.Type(), aaapreauthenticationpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaapreauthenticationpolicy state %s", aaapreauthenticationpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAaapreauthenticationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaapreauthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationpolicyName := d.Get("name").(string)

	aaapreauthenticationpolicy := aaa.Aaapreauthenticationpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for aaapreauthenticationpolicy %s, starting update", aaapreauthenticationpolicyName)
		aaapreauthenticationpolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for aaapreauthenticationpolicy %s, starting update", aaapreauthenticationpolicyName)
		aaapreauthenticationpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaapreauthenticationpolicy.Type(),&aaapreauthenticationpolicy)
		if err != nil {
			return fmt.Errorf("Error updating aaapreauthenticationpolicy %s", aaapreauthenticationpolicyName)
		}
	}
	return readAaapreauthenticationpolicyFunc(d, meta)
}

func deleteAaapreauthenticationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaapreauthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	aaapreauthenticationpolicyName := d.Id()
	err := client.DeleteResource(service.Aaapreauthenticationpolicy.Type(), aaapreauthenticationpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
