package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationnegotiatepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationnegotiatepolicyFunc,
		Read:          readAuthenticationnegotiatepolicyFunc,
		Update:        updateAuthenticationnegotiatepolicyFunc,
		Delete:        deleteAuthenticationnegotiatepolicyFunc,
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

func createAuthenticationnegotiatepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationnegotiatepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnegotiatepolicyName := d.Get("name").(string)
	authenticationnegotiatepolicy := authentication.Authenticationnegotiatepolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationnegotiatepolicy.Type(), authenticationnegotiatepolicyName, &authenticationnegotiatepolicy)
	if err != nil {
		return err
	}

	d.SetId(authenticationnegotiatepolicyName)

	err = readAuthenticationnegotiatepolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationnegotiatepolicy but we can't read it ?? %s", authenticationnegotiatepolicyName)
		return nil
	}
	return nil
}

func readAuthenticationnegotiatepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationnegotiatepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnegotiatepolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationnegotiatepolicy state %s", authenticationnegotiatepolicyName)
	data, err := client.FindResource(service.Authenticationnegotiatepolicy.Type(), authenticationnegotiatepolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationnegotiatepolicy state %s", authenticationnegotiatepolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationnegotiatepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationnegotiatepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnegotiatepolicyName := d.Get("name").(string)

	authenticationnegotiatepolicy := authentication.Authenticationnegotiatepolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for authenticationnegotiatepolicy %s, starting update", authenticationnegotiatepolicyName)
		authenticationnegotiatepolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationnegotiatepolicy %s, starting update", authenticationnegotiatepolicyName)
		authenticationnegotiatepolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationnegotiatepolicy.Type(), authenticationnegotiatepolicyName, &authenticationnegotiatepolicy)
		if err != nil {
			return fmt.Errorf("Error updating authenticationnegotiatepolicy %s", authenticationnegotiatepolicyName)
		}
	}
	return readAuthenticationnegotiatepolicyFunc(d, meta)
}

func deleteAuthenticationnegotiatepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationnegotiatepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnegotiatepolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationnegotiatepolicy.Type(), authenticationnegotiatepolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
