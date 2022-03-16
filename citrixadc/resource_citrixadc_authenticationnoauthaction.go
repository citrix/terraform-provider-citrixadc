package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationnoauthaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationnoauthactionFunc,
		Read:          readAuthenticationnoauthactionFunc,
		Update:        updateAuthenticationnoauthactionFunc,
		Delete:        deleteAuthenticationnoauthactionFunc,
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
			"defaultauthenticationgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationnoauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationnoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnoauthactionName := d.Get("name").(string)
	authenticationnoauthaction := authentication.Authenticationnoauthaction{
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Name:                       d.Get("name").(string),
	}

	_, err := client.AddResource("authenticationnoauthaction", authenticationnoauthactionName, &authenticationnoauthaction)
	if err != nil {
		return err
	}

	d.SetId(authenticationnoauthactionName)

	err = readAuthenticationnoauthactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationnoauthaction but we can't read it ?? %s", authenticationnoauthactionName)
		return nil
	}
	return nil
}

func readAuthenticationnoauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationnoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnoauthactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationnoauthaction state %s", authenticationnoauthactionName)
	data, err := client.FindResource("authenticationnoauthaction", authenticationnoauthactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationnoauthaction state %s", authenticationnoauthactionName)
		d.SetId("")
		return nil
	}
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("name", data["name"])

	return nil

}

func updateAuthenticationnoauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationnoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnoauthactionName := d.Get("name").(string)

	authenticationnoauthaction := authentication.Authenticationnoauthaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationnoauthaction %s, starting update", authenticationnoauthactionName)
		authenticationnoauthaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationnoauthaction", authenticationnoauthactionName, &authenticationnoauthaction)
		if err != nil {
			return fmt.Errorf("Error updating authenticationnoauthaction %s", authenticationnoauthactionName)
		}
	}
	return readAuthenticationnoauthactionFunc(d, meta)
}

func deleteAuthenticationnoauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationnoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnoauthactionName := d.Id()
	err := client.DeleteResource("authenticationnoauthaction", authenticationnoauthactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
