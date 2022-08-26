package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/autoscale"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAutoscalepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAutoscalepolicyFunc,
		Read:          readAutoscalepolicyFunc,
		Update:        updateAutoscalepolicyFunc,
		Delete:        deleteAutoscalepolicyFunc,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
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

func createAutoscalepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAutoscalepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	var autoscalepolicyName string
	if v, ok := d.GetOk("name"); ok {
		autoscalepolicyName = v.(string)
	} else {
		autoscalepolicyName = resource.PrefixedUniqueId("tf-autoscalepolicy-")
		d.Set("name", autoscalepolicyName)
	}
	autoscalepolicy := autoscale.Autoscalepolicy{
		Action:    d.Get("action").(string),
		Comment:   d.Get("comment").(string),
		Logaction: d.Get("logaction").(string),
		Name:      d.Get("name").(string),
		Newname:   d.Get("newname").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Autoscalepolicy.Type(), autoscalepolicyName, &autoscalepolicy)
	if err != nil {
		return err
	}

	d.SetId(autoscalepolicyName)

	err = readAutoscalepolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this autoscalepolicy but we can't read it ?? %s", autoscalepolicyName)
		return nil
	}
	return nil
}

func readAutoscalepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAutoscalepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscalepolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading autoscalepolicy state %s", autoscalepolicyName)
	data, err := client.FindResource(service.Autoscalepolicy.Type(), autoscalepolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing autoscalepolicy state %s", autoscalepolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAutoscalepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAutoscalepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscalepolicyName := d.Get("name").(string)

	autoscalepolicy := autoscale.Autoscalepolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for autoscalepolicy %s, starting update", autoscalepolicyName)
		autoscalepolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for autoscalepolicy %s, starting update", autoscalepolicyName)
		autoscalepolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for autoscalepolicy %s, starting update", autoscalepolicyName)
		autoscalepolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for autoscalepolicy %s, starting update", autoscalepolicyName)
		autoscalepolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for autoscalepolicy %s, starting update", autoscalepolicyName)
		autoscalepolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for autoscalepolicy %s, starting update", autoscalepolicyName)
		autoscalepolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Autoscalepolicy.Type(), autoscalepolicyName, &autoscalepolicy)
		if err != nil {
			return fmt.Errorf("Error updating autoscalepolicy %s", autoscalepolicyName)
		}
	}
	return readAutoscalepolicyFunc(d, meta)
}

func deleteAutoscalepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAutoscalepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscalepolicyName := d.Id()
	err := client.DeleteResource(service.Autoscalepolicy.Type(), autoscalepolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
