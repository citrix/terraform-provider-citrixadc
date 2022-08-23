package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ica"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcIcapolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIcapolicyFunc,
		Read:          readIcapolicyFunc,
		Update:        updateIcapolicyFunc,
		Delete:        deleteIcapolicyFunc,
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
		},
	}
}

func createIcapolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIcapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	icapolicyName := d.Get("name").(string)
	icapolicy := ica.Icapolicy{
		Action:    d.Get("action").(string),
		Comment:   d.Get("comment").(string),
		Logaction: d.Get("logaction").(string),
		Name:      d.Get("name").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource("icapolicy", icapolicyName, &icapolicy)
	if err != nil {
		return err
	}

	d.SetId(icapolicyName)

	err = readIcapolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this icapolicy but we can't read it ?? %s", icapolicyName)
		return nil
	}
	return nil
}

func readIcapolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIcapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	icapolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading icapolicy state %s", icapolicyName)
	data, err := client.FindResource("icapolicy", icapolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing icapolicy state %s", icapolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateIcapolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIcapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	icapolicyName := d.Get("name").(string)

	icapolicy := ica.Icapolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for icapolicy %s, starting update", icapolicyName)
		icapolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for icapolicy %s, starting update", icapolicyName)
		icapolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for icapolicy %s, starting update", icapolicyName)
		icapolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for icapolicy %s, starting update", icapolicyName)
		icapolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("icapolicy", &icapolicy)
		if err != nil {
			return fmt.Errorf("Error updating icapolicy %s", icapolicyName)
		}
	}
	return readIcapolicyFunc(d, meta)
}

func deleteIcapolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIcapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	icapolicyName := d.Id()
	err := client.DeleteResource("icapolicy", icapolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
