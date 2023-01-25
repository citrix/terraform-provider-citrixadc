package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationsamlpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationsamlpolicyFunc,
		Read:          readAuthenticationsamlpolicyFunc,
		Update:        updateAuthenticationsamlpolicyFunc,
		Delete:        deleteAuthenticationsamlpolicyFunc,
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
			"reqaction": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createAuthenticationsamlpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationsamlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlpolicyName := d.Get("name").(string)
	authenticationsamlpolicy := authentication.Authenticationsamlpolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName, &authenticationsamlpolicy)
	if err != nil {
		return err
	}

	d.SetId(authenticationsamlpolicyName)

	err = readAuthenticationsamlpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationsamlpolicy but we can't read it ?? %s", authenticationsamlpolicyName)
		return nil
	}
	return nil
}

func readAuthenticationsamlpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationsamlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationsamlpolicy state %s", authenticationsamlpolicyName)
	data, err := client.FindResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationsamlpolicy state %s", authenticationsamlpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationsamlpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationsamlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlpolicyName := d.Get("name").(string)

	authenticationsamlpolicy := authentication.Authenticationsamlpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for authenticationsamlpolicy %s, starting update", authenticationsamlpolicyName)
		authenticationsamlpolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationsamlpolicy %s, starting update", authenticationsamlpolicyName)
		authenticationsamlpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName, &authenticationsamlpolicy)
		if err != nil {
			return fmt.Errorf("Error updating authenticationsamlpolicy %s", authenticationsamlpolicyName)
		}
	}
	return readAuthenticationsamlpolicyFunc(d, meta)
}

func deleteAuthenticationsamlpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationsamlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlpolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
