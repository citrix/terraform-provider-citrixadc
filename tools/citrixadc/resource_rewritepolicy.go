package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/rewrite"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcRewritepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRewritepolicyFunc,
		Read:          readRewritepolicyFunc,
		Update:        updateRewritepolicyFunc,
		Delete:        deleteRewritepolicyFunc,
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
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRewritepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRewritepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	var rewritepolicyName string
	if v, ok := d.GetOk("name"); ok {
		rewritepolicyName = v.(string)
	} else {
		rewritepolicyName = resource.PrefixedUniqueId("tf-rewritepolicy-")
		d.Set("name", rewritepolicyName)
	}
	rewritepolicy := rewrite.Rewritepolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Newname:     d.Get("newname").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource(netscaler.Rewritepolicy.Type(), rewritepolicyName, &rewritepolicy)
	if err != nil {
		return err
	}

	d.SetId(rewritepolicyName)

	err = readRewritepolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rewritepolicy but we can't read it ?? %s", rewritepolicyName)
		return nil
	}
	return nil
}

func readRewritepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRewritepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rewritepolicy state %s", rewritepolicyName)
	data, err := client.FindResource(netscaler.Rewritepolicy.Type(), rewritepolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rewritepolicy state %s", rewritepolicyName)
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
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateRewritepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRewritepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicyName := d.Get("name").(string)

	rewritepolicy := rewrite.Rewritepolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Rewritepolicy.Type(), rewritepolicyName, &rewritepolicy)
		if err != nil {
			return fmt.Errorf("Error updating rewritepolicy %s", rewritepolicyName)
		}
	}
	return readRewritepolicyFunc(d, meta)
}

func deleteRewritepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewritepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicyName := d.Id()
	err := client.DeleteResource(netscaler.Rewritepolicy.Type(), rewritepolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
