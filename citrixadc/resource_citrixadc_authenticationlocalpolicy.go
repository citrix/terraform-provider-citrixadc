package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationlocalpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationlocalpolicyFunc,
		Read:          readAuthenticationlocalpolicyFunc,
		Update:        updateAuthenticationlocalpolicyFunc,
		Delete:        deleteAuthenticationlocalpolicyFunc,
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
		},
	}
}

func createAuthenticationlocalpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationlocalpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationlocalpolicyName := d.Get("name").(string)
	authenticationlocalpolicy := authentication.Authenticationlocalpolicy{
		Name: d.Get("name").(string),
		Rule: d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationlocalpolicy.Type(), authenticationlocalpolicyName, &authenticationlocalpolicy)
	if err != nil {
		return err
	}

	d.SetId(authenticationlocalpolicyName)

	err = readAuthenticationlocalpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationlocalpolicy but we can't read it ?? %s", authenticationlocalpolicyName)
		return nil
	}
	return nil
}

func readAuthenticationlocalpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationlocalpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationlocalpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationlocalpolicy state %s", authenticationlocalpolicyName)
	data, err := client.FindResource(service.Authenticationlocalpolicy.Type(), authenticationlocalpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationlocalpolicy state %s", authenticationlocalpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationlocalpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationlocalpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationlocalpolicyName := d.Get("name").(string)

	authenticationlocalpolicy := authentication.Authenticationlocalpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationlocalpolicy %s, starting update", authenticationlocalpolicyName)
		authenticationlocalpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationlocalpolicy.Type(), authenticationlocalpolicyName, &authenticationlocalpolicy)
		if err != nil {
			return fmt.Errorf("Error updating authenticationlocalpolicy %s", authenticationlocalpolicyName)
		}
	}
	return readAuthenticationlocalpolicyFunc(d, meta)
}

func deleteAuthenticationlocalpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationlocalpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationlocalpolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationlocalpolicy.Type(), authenticationlocalpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
