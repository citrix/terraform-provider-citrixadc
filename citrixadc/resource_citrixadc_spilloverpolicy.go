package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/spillover"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSpilloverpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSpilloverpolicyFunc,
		Read:          readSpilloverpolicyFunc,
		Update:        updateSpilloverpolicyFunc,
		Delete:        deleteSpilloverpolicyFunc,
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
				Computed: false,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSpilloverpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSpilloverpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloverpolicyName:= d.Get("name").(string)
	spilloverpolicy := spillover.Spilloverpolicy{
		Action:  d.Get("action").(string),
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Newname: d.Get("newname").(string),
		Rule:    d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Spilloverpolicy.Type(), spilloverpolicyName, &spilloverpolicy)
	if err != nil {
		return err
	}

	d.SetId(spilloverpolicyName)

	err = readSpilloverpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this spilloverpolicy but we can't read it ?? %s", spilloverpolicyName)
		return nil
	}
	return nil
}

func readSpilloverpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSpilloverpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloverpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading spilloverpolicy state %s", spilloverpolicyName)
	data, err := client.FindResource(service.Spilloverpolicy.Type(), spilloverpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing spilloverpolicy state %s", spilloverpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])

	return nil

}

func updateSpilloverpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSpilloverpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloverpolicyName := d.Get("name").(string)

	spilloverpolicy := spillover.Spilloverpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for spilloverpolicy %s, starting update", spilloverpolicyName)
		spilloverpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for spilloverpolicy %s, starting update", spilloverpolicyName)
		spilloverpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for spilloverpolicy %s, starting update", spilloverpolicyName)
		spilloverpolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for spilloverpolicy %s, starting update", spilloverpolicyName)
		spilloverpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Spilloverpolicy.Type(), spilloverpolicyName, &spilloverpolicy)
		if err != nil {
			return fmt.Errorf("Error updating spilloverpolicy %s", spilloverpolicyName)
		}
	}
	return readSpilloverpolicyFunc(d, meta)
}

func deleteSpilloverpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSpilloverpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloverpolicyName := d.Id()
	err := client.DeleteResource(service.Spilloverpolicy.Type(), spilloverpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
