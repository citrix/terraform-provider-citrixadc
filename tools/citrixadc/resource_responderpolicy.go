package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/responder"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcResponderpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createResponderpolicyFunc,
		Read:          readResponderpolicyFunc,
		Update:        updateResponderpolicyFunc,
		Delete:        deleteResponderpolicyFunc,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"appflowaction": &schema.Schema{
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
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createResponderpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createResponderpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	var responderpolicyName string
	if v, ok := d.GetOk("name"); ok {
		responderpolicyName = v.(string)
	} else {
		responderpolicyName = resource.PrefixedUniqueId("tf-responderpolicy-")
		d.Set("name", responderpolicyName)
	}
	responderpolicy := responder.Responderpolicy{
		Action:        d.Get("action").(string),
		Appflowaction: d.Get("appflowaction").(string),
		Comment:       d.Get("comment").(string),
		Logaction:     d.Get("logaction").(string),
		Name:          d.Get("name").(string),
		Newname:       d.Get("newname").(string),
		Rule:          d.Get("rule").(string),
		Undefaction:   d.Get("undefaction").(string),
	}

	_, err := client.AddResource(netscaler.Responderpolicy.Type(), responderpolicyName, &responderpolicy)
	if err != nil {
		return err
	}

	d.SetId(responderpolicyName)

	err = readResponderpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this responderpolicy but we can't read it ?? %s", responderpolicyName)
		return nil
	}
	return nil
}

func readResponderpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readResponderpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading responderpolicy state %s", responderpolicyName)
	data, err := client.FindResource(netscaler.Responderpolicy.Type(), responderpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing responderpolicy state %s", responderpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("appflowaction", data["appflowaction"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateResponderpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateResponderpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicyName := d.Get("name").(string)

	responderpolicy := responder.Responderpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("appflowaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowaction has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Appflowaction = d.Get("appflowaction").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Responderpolicy.Type(), responderpolicyName, &responderpolicy)
		if err != nil {
			return fmt.Errorf("Error updating responderpolicy %s", responderpolicyName)
		}
	}
	return readResponderpolicyFunc(d, meta)
}

func deleteResponderpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteResponderpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicyName := d.Id()
	err := client.DeleteResource(netscaler.Responderpolicy.Type(), responderpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
