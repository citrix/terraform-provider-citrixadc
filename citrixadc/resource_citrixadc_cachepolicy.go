package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCachepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCachepolicyFunc,
		Read:          readCachepolicyFunc,
		Update:        updateCachepolicyFunc,
		Delete:        deleteCachepolicyFunc,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"invalgroups": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"invalobjects": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"storeingroup": &schema.Schema{
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

func createCachepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCachepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	var cachepolicyName string
	cachepolicyName = d.Get("policyname").(string)
	cachepolicy := cache.Cachepolicy{
		Action:       d.Get("action").(string),
		Invalgroups:  toStringList(d.Get("invalgroups").([]interface{})),
		Invalobjects: toStringList(d.Get("invalobjects").([]interface{})),
		Newname:      d.Get("newname").(string),
		Policyname:   d.Get("policyname").(string),
		Rule:         d.Get("rule").(string),
		Storeingroup: d.Get("storeingroup").(string),
		Undefaction:  d.Get("undefaction").(string),
	}

	_, err := client.AddResource(service.Cachepolicy.Type(), cachepolicyName, &cachepolicy)
	if err != nil {
		return err
	}

	d.SetId(cachepolicyName)

	err = readCachepolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cachepolicy but we can't read it ?? %s", cachepolicyName)
		return nil
	}
	return nil
}

func readCachepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCachepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cachepolicy state %s", cachepolicyName)
	data, err := client.FindResource(service.Cachepolicy.Type(), cachepolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cachepolicy state %s", cachepolicyName)
		d.SetId("")
		return nil
	}
	d.Set("policyname", data["policyname"])
	d.Set("action", data["action"])
	d.Set("invalgroups", data["invalgroups"])
	d.Set("invalobjects", data["invalobjects"])
	d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])
	d.Set("storeingroup", data["storeingroup"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateCachepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCachepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicyName := d.Get("policyname").(string)

	cachepolicy := cache.Cachepolicy{
		Policyname: d.Get("policyname").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("invalgroups") {
		log.Printf("[DEBUG]  citrixadc-provider: Invalgroups has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Invalgroups = toStringList(d.Get("invalgroups").([]interface{}))
		hasChange = true
	}
	if d.HasChange("invalobjects") {
		log.Printf("[DEBUG]  citrixadc-provider: Invalobjects has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Invalobjects = toStringList(d.Get("invalobjects").([]interface{}))
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("storeingroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Storeingroup has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Storeingroup = d.Get("storeingroup").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for cachepolicy %s, starting update", cachepolicyName)
		cachepolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Cachepolicy.Type(), cachepolicyName, &cachepolicy)
		if err != nil {
			return fmt.Errorf("Error updating cachepolicy %s", cachepolicyName)
		}
	}
	return readCachepolicyFunc(d, meta)
}

func deleteCachepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCachepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cachepolicyName := d.Id()
	err := client.DeleteResource(service.Cachepolicy.Type(), cachepolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
