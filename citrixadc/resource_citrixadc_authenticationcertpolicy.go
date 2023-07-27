package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationcertpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationcertpolicyFunc,
		Read:          readAuthenticationcertpolicyFunc,
		Update:        updateAuthenticationcertpolicyFunc,
		Delete:        deleteAuthenticationcertpolicyFunc,
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
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"reqaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationcertpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationcertpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertpolicyName := d.Get("name").(string)
	authenticationcertpolicy := authentication.Authenticationcertpolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName, &authenticationcertpolicy)
	if err != nil {
		return err
	}

	d.SetId(authenticationcertpolicyName)

	err = readAuthenticationcertpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationcertpolicy but we can't read it ?? %s", authenticationcertpolicyName)
		return nil
	}
	return nil
}

func readAuthenticationcertpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationcertpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationcertpolicy state %s", authenticationcertpolicyName)
	data, err := client.FindResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationcertpolicy state %s", authenticationcertpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationcertpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationcertpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertpolicyName := d.Get("name").(string)

	authenticationcertpolicy := authentication.Authenticationcertpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for authenticationcertpolicy %s, starting update", authenticationcertpolicyName)
		authenticationcertpolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationcertpolicy %s, starting update", authenticationcertpolicyName)
		authenticationcertpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName, &authenticationcertpolicy)
		if err != nil {
			return fmt.Errorf("Error updating authenticationcertpolicy %s", authenticationcertpolicyName)
		}
	}
	return readAuthenticationcertpolicyFunc(d, meta)
}

func deleteAuthenticationcertpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationcertpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertpolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
