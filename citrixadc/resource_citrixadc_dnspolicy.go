package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcDnspolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnspolicyFunc,
		Read:          readDnspolicyFunc,
		Update:        updateDnspolicyFunc,
		Delete:        deleteDnspolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"actionname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachebypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drop": &schema.Schema{
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
				Required: true,
				ForceNew: true,
			},
			"preferredlocation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preferredloclist": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"viewname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	dnspolicyName := d.Get("name").(string)
	dnspolicy := dns.Dnspolicy{
		Actionname:        d.Get("actionname").(string),
		Cachebypass:       d.Get("cachebypass").(string),
		Drop:              d.Get("drop").(string),
		Logaction:         d.Get("logaction").(string),
		Name:              dnspolicyName,
		Preferredlocation: d.Get("preferredlocation").(string),
		Preferredloclist:  toStringList(d.Get("preferredloclist").([]interface{})),
		Rule:              d.Get("rule").(string),
		Viewname:          d.Get("viewname").(string),
	}

	_, err := client.AddResource(service.Dnspolicy.Type(), dnspolicyName, &dnspolicy)
	if err != nil {
		return err
	}

	d.SetId(dnspolicyName)

	err = readDnspolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnspolicy but we can't read it ?? %s", dnspolicyName)
		return nil
	}
	return nil
}

func readDnspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	dnspolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnspolicy state %s", dnspolicyName)
	data, err := client.FindResource(service.Dnspolicy.Type(), dnspolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnspolicy state %s", dnspolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("actionname", data["actionname"])
	d.Set("cachebypass", data["cachebypass"])
	d.Set("drop", data["drop"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("preferredlocation", data["preferredlocation"])
	d.Set("preferredloclist", data["preferredloclist"])
	d.Set("rule", data["rule"])
	d.Set("viewname", data["viewname"])

	return nil

}

func updateDnspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	dnspolicyName := d.Get("name").(string)

	dnspolicy := dns.Dnspolicy{
		Name: dnspolicyName,
	}
	hasChange := false
	if d.HasChange("actionname") {
		log.Printf("[DEBUG]  citrixadc-provider: Actionname has changed for dnspolicy %s, starting update", dnspolicyName)
		dnspolicy.Actionname = d.Get("actionname").(string)
		hasChange = true
	}
	if d.HasChange("cachebypass") {
		log.Printf("[DEBUG]  citrixadc-provider: Cachebypass has changed for dnspolicy %s, starting update", dnspolicyName)
		dnspolicy.Cachebypass = d.Get("cachebypass").(string)
		hasChange = true
	}
	if d.HasChange("drop") {
		log.Printf("[DEBUG]  citrixadc-provider: Drop has changed for dnspolicy %s, starting update", dnspolicyName)
		dnspolicy.Drop = d.Get("drop").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for dnspolicy %s, starting update", dnspolicyName)
		dnspolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("preferredlocation") {
		log.Printf("[DEBUG]  citrixadc-provider: Preferredlocation has changed for dnspolicy %s, starting update", dnspolicyName)
		dnspolicy.Preferredlocation = d.Get("preferredlocation").(string)
		hasChange = true
	}
	if d.HasChange("preferredloclist") {
		log.Printf("[DEBUG]  citrixadc-provider: Preferredloclist has changed for dnspolicy %s, starting update", dnspolicyName)
		dnspolicy.Preferredloclist = toStringList(d.Get("preferredloclist").([]interface{}))
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for dnspolicy %s, starting update", dnspolicyName)
		dnspolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("viewname") {
		log.Printf("[DEBUG]  citrixadc-provider: Viewname has changed for dnspolicy %s, starting update", dnspolicyName)
		dnspolicy.Viewname = d.Get("viewname").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Dnspolicy.Type(), dnspolicyName, &dnspolicy)
		if err != nil {
			return fmt.Errorf("Error updating dnspolicy %s", dnspolicyName)
		}
	}
	return readDnspolicyFunc(d, meta)
}

func deleteDnspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	dnspolicyName := d.Id()
	err := client.DeleteResource(service.Dnspolicy.Type(), dnspolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
