package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cmp"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCmppolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCmppolicyFunc,
		Read:          readCmppolicyFunc,
		Update:        updateCmppolicyFunc,
		Delete:        deleteCmppolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resaction": &schema.Schema{
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

func createCmppolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCmppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicyName := d.Get("name").(string)
	cmppolicy := cmp.Cmppolicy{
		Name:      d.Get("name").(string),
		Resaction: d.Get("resaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Cmppolicy.Type(), cmppolicyName, &cmppolicy)
	if err != nil {
		return err
	}

	d.SetId(cmppolicyName)

	err = readCmppolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cmppolicy but we can't read it ?? %s", cmppolicyName)
		return nil
	}
	return nil
}

func readCmppolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCmppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cmppolicy state %s", cmppolicyName)
	data, err := client.FindResource(service.Cmppolicy.Type(), cmppolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cmppolicy state %s", cmppolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("name", data["name"])
	d.Set("resaction", data["resaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateCmppolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCmppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicyName := d.Get("name").(string)

	cmppolicy := cmp.Cmppolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for cmppolicy %s, starting update", cmppolicyName)
		cmppolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("resaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Resaction has changed for cmppolicy %s, starting update", cmppolicyName)
		cmppolicy.Resaction = d.Get("resaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for cmppolicy %s, starting update", cmppolicyName)
		cmppolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Cmppolicy.Type(), cmppolicyName, &cmppolicy)
		if err != nil {
			return fmt.Errorf("Error updating cmppolicy %s", cmppolicyName)
		}
	}
	return readCmppolicyFunc(d, meta)
}

func deleteCmppolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCmppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cmppolicyName := d.Id()
	err := client.DeleteResource(service.Cmppolicy.Type(), cmppolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
