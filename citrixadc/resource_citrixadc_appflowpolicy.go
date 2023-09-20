package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appflow"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAppflowpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppflowpolicyFunc,
		Read:          readAppflowpolicyFunc,
		Update:        updateAppflowpolicyFunc,
		Delete:        deleteAppflowpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppflowpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppflowpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowpolicyName := d.Get("name").(string)

	appflowpolicy := appflow.Appflowpolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Name:        d.Get("name").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource(service.Appflowpolicy.Type(), appflowpolicyName, &appflowpolicy)
	if err != nil {
		return err
	}

	d.SetId(appflowpolicyName)

	err = readAppflowpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appflowpolicy but we can't read it ?? %s", appflowpolicyName)
		return nil
	}
	return nil
}

func readAppflowpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppflowpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appflowpolicy state %s", appflowpolicyName)
	data, err := client.FindResource(service.Appflowpolicy.Type(), appflowpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appflowpolicy state %s", appflowpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateAppflowpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppflowpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowpolicyName := d.Get("name").(string)

	appflowpolicy := appflow.Appflowpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for appflowpolicy %s, starting update", appflowpolicyName)
		appflowpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for appflowpolicy %s, starting update", appflowpolicyName)
		appflowpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for appflowpolicy %s, starting update", appflowpolicyName)
		appflowpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for appflowpolicy %s, starting update", appflowpolicyName)
		appflowpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Appflowpolicy.Type(), appflowpolicyName, &appflowpolicy)
		if err != nil {
			return fmt.Errorf("Error updating appflowpolicy %s", appflowpolicyName)
		}
	}
	return readAppflowpolicyFunc(d, meta)
}

func deleteAppflowpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppflowpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowpolicyName := d.Id()
	err := client.DeleteResource(service.Appflowpolicy.Type(), appflowpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
