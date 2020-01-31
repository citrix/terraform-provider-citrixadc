package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/cs"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCspolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCspolicyFunc,
		Read:          readCspolicyFunc,
		Update:        updateCspolicyFunc,
		Delete:        deleteCspolicyFunc,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	var cspolicyName string
	if v, ok := d.GetOk("name"); ok {
		cspolicyName = v.(string)
	} else {
		cspolicyName = resource.PrefixedUniqueId("tf-cspolicy-")
		d.Set("name", cspolicyName)
	}
	cspolicy := cs.Cspolicy{
		Action:     d.Get("action").(string),
		Domain:     d.Get("domain").(string),
		Logaction:  d.Get("logaction").(string),
		Newname:    d.Get("newname").(string),
		Policyname: d.Get("policyname").(string),
		Rule:       d.Get("rule").(string),
		Url:        d.Get("url").(string),
	}

	_, err := client.AddResource(netscaler.Cspolicy.Type(), cspolicyName, &cspolicy)
	if err != nil {
		return err
	}

	d.SetId(cspolicyName)

	err = readCspolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cspolicy but we can't read it ?? %s", cspolicyName)
		return nil
	}
	return nil
}

func readCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cspolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cspolicy state %s", cspolicyName)
	data, err := client.FindResource(netscaler.Cspolicy.Type(), cspolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cspolicy state %s", cspolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("domain", data["domain"])
	d.Set("logaction", data["logaction"])
	d.Set("newname", data["newname"])
	d.Set("policyname", data["policyname"])
	d.Set("rule", data["rule"])
	d.Set("url", data["url"])

	return nil

}

func updateCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cspolicyName := d.Get("name").(string)

	cspolicy := cs.Cspolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG]  citrixadc-provider: Domain has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("policyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Policyname has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Policyname = d.Get("policyname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("url") {
		log.Printf("[DEBUG]  citrixadc-provider: Url has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Url = d.Get("url").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Cspolicy.Type(), cspolicyName, &cspolicy)
		if err != nil {
			return fmt.Errorf("Error updating cspolicy %s", cspolicyName)
		}
	}
	return readCspolicyFunc(d, meta)
}

func deleteCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cspolicyName := d.Id()
	err := client.DeleteResource(netscaler.Cspolicy.Type(), cspolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
