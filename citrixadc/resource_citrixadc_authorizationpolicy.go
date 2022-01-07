package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authorization"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthorizationpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthorizationpolicyFunc,
		Read:          readAuthorizationpolicyFunc,
		Update:        updateAuthorizationpolicyFunc,
		Delete:        deleteAuthorizationpolicyFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Computed: false,
			},
			// "newname": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Computed: true,
			// },
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createAuthorizationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthorizationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicyName := d.Get("name").(string)
	authorizationpolicy := authorization.Authorizationpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		// Newname: d.Get("newname").(string),
		Rule: d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authorizationpolicy.Type(), authorizationpolicyName, &authorizationpolicy)
	if err != nil {
		return err
	}

	d.SetId(authorizationpolicyName)

	err = readAuthorizationpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authorizationpolicy but we can't read it ?? %s", authorizationpolicyName)
		return nil
	}
	return nil
}

func readAuthorizationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthorizationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authorizationpolicy state %s", authorizationpolicyName)
	data, err := client.FindResource(service.Authorizationpolicy.Type(), authorizationpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authorizationpolicy state %s", authorizationpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	// d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthorizationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthorizationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicyName := d.Get("name").(string)

	authorizationpolicy := authorization.Authorizationpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for authorizationpolicy %s, starting update", authorizationpolicyName)
		authorizationpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	// if d.HasChange("name") {
	// 	log.Printf("[DEBUG]  citrixadc-provider: Name has changed for authorizationpolicy %s, starting update", authorizationpolicyName)
	// 	authorizationpolicy.Name = d.Get("name").(string)
	// 	hasChange = true
	// }
	// if d.HasChange("newname") {
	// 	log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for authorizationpolicy %s, starting update", authorizationpolicyName)
	// 	authorizationpolicy.Newname = d.Get("newname").(string)
	// 	hasChange = true
	// }
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authorizationpolicy %s, starting update", authorizationpolicyName)
		authorizationpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authorizationpolicy.Type(), authorizationpolicyName, &authorizationpolicy)
		if err != nil {
			return fmt.Errorf("error updating authorizationpolicy %s", authorizationpolicyName)
		}
	}
	return readAuthorizationpolicyFunc(d, meta)
}

func deleteAuthorizationpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthorizationpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicyName := d.Id()
	err := client.DeleteResource(service.Authorizationpolicy.Type(), authorizationpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
