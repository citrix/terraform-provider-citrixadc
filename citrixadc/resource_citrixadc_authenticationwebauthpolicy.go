package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationwebauthpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationwebauthpolicyFunc,
		Read:          readAuthenticationwebauthpolicyFunc,
		Update:        updateAuthenticationwebauthpolicyFunc,
		Delete:        deleteAuthenticationwebauthpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createAuthenticationwebauthpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationwebauthpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthpolicyName := d.Get("name").(string)
	authenticationwebauthpolicy := authentication.Authenticationwebauthpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName, &authenticationwebauthpolicy)
	if err != nil {
		return err
	}

	d.SetId(authenticationwebauthpolicyName)

	err = readAuthenticationwebauthpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationwebauthpolicy but we can't read it ?? %s", authenticationwebauthpolicyName)
		return nil
	}
	return nil
}

func readAuthenticationwebauthpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationwebauthpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationwebauthpolicy state %s", authenticationwebauthpolicyName)
	data, err := client.FindResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationwebauthpolicy state %s", authenticationwebauthpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationwebauthpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationwebauthpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthpolicyName := d.Get("name").(string)

	authenticationwebauthpolicy := authentication.Authenticationwebauthpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for authenticationwebauthpolicy %s, starting update", authenticationwebauthpolicyName)
		authenticationwebauthpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationwebauthpolicy %s, starting update", authenticationwebauthpolicyName)
		authenticationwebauthpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName, &authenticationwebauthpolicy)
		if err != nil {
			return fmt.Errorf("Error updating authenticationwebauthpolicy %s", authenticationwebauthpolicyName)
		}
	}
	return readAuthenticationwebauthpolicyFunc(d, meta)
}

func deleteAuthenticationwebauthpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationwebauthpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthpolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
