package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAppfwpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwpolicyFunc,
		Read:          readAppfwpolicyFunc,
		Update:        updateAppfwpolicyFunc,
		Delete:        deleteAppfwpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
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
				Required: true,
				ForceNew: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profilename": &schema.Schema{
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

func createAppfwpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwpolicyName := d.Get("name").(string)

	appfwpolicy := appfw.Appfwpolicy{
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        appfwpolicyName,
		Newname:     d.Get("newname").(string),
		Profilename: d.Get("profilename").(string),
		Rule:        d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Appfwpolicy.Type(), appfwpolicyName, &appfwpolicy)
	if err != nil {
		return err
	}

	d.SetId(appfwpolicyName)

	err = readAppfwpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwpolicy but we can't read it ?? %s", appfwpolicyName)
		return nil
	}
	return nil
}

func readAppfwpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwpolicy state %s", appfwpolicyName)
	data, err := client.FindResource(service.Appfwpolicy.Type(), appfwpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwpolicy state %s", appfwpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("profilename", data["profilename"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAppfwpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppfwpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwpolicyName := d.Get("name").(string)

	appfwpolicy := appfw.Appfwpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for appfwpolicy %s, starting update", appfwpolicyName)
		appfwpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for appfwpolicy %s, starting update", appfwpolicyName)
		appfwpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for appfwpolicy %s, starting update", appfwpolicyName)
		appfwpolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for appfwpolicy %s, starting update", appfwpolicyName)
		appfwpolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("profilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Profilename has changed for appfwpolicy %s, starting update", appfwpolicyName)
		appfwpolicy.Profilename = d.Get("profilename").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for appfwpolicy %s, starting update", appfwpolicyName)
		appfwpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Appfwpolicy.Type(), appfwpolicyName, &appfwpolicy)
		if err != nil {
			return fmt.Errorf("Error updating appfwpolicy %s", appfwpolicyName)
		}
	}
	return readAppfwpolicyFunc(d, meta)
}

func deleteAppfwpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwpolicyName := d.Id()
	err := client.DeleteResource(service.Appfwpolicy.Type(), appfwpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
