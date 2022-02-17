package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationpolicyFunc,
		Read:          readAuthenticationpolicyFunc,
		Update:        updateAuthenticationpolicyFunc,
		Delete:        deleteAuthenticationpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
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
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpolicyName := d.Get("name").(string)
	authenticationpolicy := authentication.Authenticationpolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Newname:     d.Get("newname").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource(service.Authenticationpolicy.Type(), authenticationpolicyName, &authenticationpolicy)
	if err != nil {
		return err
	}

	d.SetId(authenticationpolicyName)

	err = readAuthenticationpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationpolicy but we can't read it ?? %s", authenticationpolicyName)
		return nil
	}
	return nil
}

func readAuthenticationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationpolicy state %s", authenticationpolicyName)
	data, err := client.FindResource(service.Authenticationpolicy.Type(), authenticationpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationpolicy state %s", authenticationpolicyName)
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

func updateAuthenticationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpolicyName := d.Get("name").(string)

	authenticationpolicy := authentication.Authenticationpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for authenticationpolicy %s, starting update", authenticationpolicyName)
		authenticationpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for authenticationpolicy %s, starting update", authenticationpolicyName)
		authenticationpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for authenticationpolicy %s, starting update", authenticationpolicyName)
		authenticationpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for authenticationpolicy %s, starting update", authenticationpolicyName)
		authenticationpolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationpolicy %s, starting update", authenticationpolicyName)
		authenticationpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for authenticationpolicy %s, starting update", authenticationpolicyName)
		authenticationpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationpolicy.Type(), authenticationpolicyName, &authenticationpolicy)
		if err != nil {
			return fmt.Errorf("Error updating authenticationpolicy %s", authenticationpolicyName)
		}
	}
	return readAuthenticationpolicyFunc(d, meta)
}

func deleteAuthenticationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationpolicy.Type(), authenticationpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
