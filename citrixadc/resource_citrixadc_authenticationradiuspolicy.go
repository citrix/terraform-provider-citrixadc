package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationradiuspolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationradiuspolicyFunc,
		Read:          readAuthenticationradiuspolicyFunc,
		Update:        updateAuthenticationradiuspolicyFunc,
		Delete:        deleteAuthenticationradiuspolicyFunc,
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
			"reqaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationradiuspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationradiuspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiuspolicyName := d.Get("name").(string)
	authenticationradiuspolicy := authentication.Authenticationradiuspolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName, &authenticationradiuspolicy)
	if err != nil {
		return err
	}

	d.SetId(authenticationradiuspolicyName)

	err = readAuthenticationradiuspolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationradiuspolicy but we can't read it ?? %s", authenticationradiuspolicyName)
		return nil
	}
	return nil
}

func readAuthenticationradiuspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationradiuspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiuspolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationradiuspolicy state %s", authenticationradiuspolicyName)
	data, err := client.FindResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationradiuspolicy state %s", authenticationradiuspolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationradiuspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationradiuspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiuspolicyName := d.Get("name").(string)

	authenticationradiuspolicy := authentication.Authenticationradiuspolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for authenticationradiuspolicy %s, starting update", authenticationradiuspolicyName)
		authenticationradiuspolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationradiuspolicy %s, starting update", authenticationradiuspolicyName)
		authenticationradiuspolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName, &authenticationradiuspolicy)
		if err != nil {
			return fmt.Errorf("Error updating authenticationradiuspolicy %s", authenticationradiuspolicyName)
		}
	}
	return readAuthenticationradiuspolicyFunc(d, meta)
}

func deleteAuthenticationradiuspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationradiuspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiuspolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
