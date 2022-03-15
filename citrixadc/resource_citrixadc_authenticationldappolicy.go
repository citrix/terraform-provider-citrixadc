package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationldappolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationldappolicyFunc,
		Read:          readAuthenticationldappolicyFunc,
		Update:        updateAuthenticationldappolicyFunc,
		Delete:        deleteAuthenticationldappolicyFunc,
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

func createAuthenticationldappolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationldappolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldappolicyName := d.Get("name").(string)
	authenticationldappolicy := authentication.Authenticationldappolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationldappolicy.Type(), authenticationldappolicyName, &authenticationldappolicy)
	if err != nil {
		return err
	}

	d.SetId(authenticationldappolicyName)

	err = readAuthenticationldappolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationldappolicy but we can't read it ?? %s", authenticationldappolicyName)
		return nil
	}
	return nil
}

func readAuthenticationldappolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationldappolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldappolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationldappolicy state %s", authenticationldappolicyName)
	data, err := client.FindResource(service.Authenticationldappolicy.Type(), authenticationldappolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationldappolicy state %s", authenticationldappolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationldappolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationldappolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldappolicyName := d.Get("name").(string)

	authenticationldappolicy := authentication.Authenticationldappolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for authenticationldappolicy %s, starting update", authenticationldappolicyName)
		authenticationldappolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationldappolicy %s, starting update", authenticationldappolicyName)
		authenticationldappolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationldappolicy.Type(), authenticationldappolicyName, &authenticationldappolicy)
		if err != nil {
			return fmt.Errorf("Error updating authenticationldappolicy %s", authenticationldappolicyName)
		}
	}
	return readAuthenticationldappolicyFunc(d, meta)
}

func deleteAuthenticationldappolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationldappolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldappolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationldappolicy.Type(), authenticationldappolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}