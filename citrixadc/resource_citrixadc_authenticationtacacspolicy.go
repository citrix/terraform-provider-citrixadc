package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationtacacspolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationtacacspolicyFunc,
		Read:          readAuthenticationtacacspolicyFunc,
		Update:        updateAuthenticationtacacspolicyFunc,
		Delete:        deleteAuthenticationtacacspolicyFunc,
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

func createAuthenticationtacacspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationtacacspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationtacacspolicyName := d.Get("name").(string)
	authenticationtacacspolicy := authentication.Authenticationtacacspolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationtacacspolicy.Type(), authenticationtacacspolicyName, &authenticationtacacspolicy)
	if err != nil {
		return err
	}

	d.SetId(authenticationtacacspolicyName)

	err = readAuthenticationtacacspolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationtacacspolicy but we can't read it ?? %s", authenticationtacacspolicyName)
		return nil
	}
	return nil
}

func readAuthenticationtacacspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationtacacspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationtacacspolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationtacacspolicy state %s", authenticationtacacspolicyName)
	data, err := client.FindResource(service.Authenticationtacacspolicy.Type(), authenticationtacacspolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationtacacspolicy state %s", authenticationtacacspolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationtacacspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationtacacspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationtacacspolicyName := d.Get("name").(string)

	authenticationtacacspolicy := authentication.Authenticationtacacspolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for authenticationtacacspolicy %s, starting update", authenticationtacacspolicyName)
		authenticationtacacspolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationtacacspolicy %s, starting update", authenticationtacacspolicyName)
		authenticationtacacspolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationtacacspolicy.Type(), authenticationtacacspolicyName, &authenticationtacacspolicy)
		if err != nil {
			return fmt.Errorf("Error updating authenticationtacacspolicy %s", authenticationtacacspolicyName)
		}
	}
	return readAuthenticationtacacspolicyFunc(d, meta)
}

func deleteAuthenticationtacacspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationtacacspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationtacacspolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationtacacspolicy.Type(), authenticationtacacspolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
