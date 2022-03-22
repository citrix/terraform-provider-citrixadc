package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationcitrixauthaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationcitrixauthactionFunc,
		Read:          readAuthenticationcitrixauthactionFunc,
		Update:        updateAuthenticationcitrixauthactionFunc,
		Delete:        deleteAuthenticationcitrixauthactionFunc,
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
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authenticationtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationcitrixauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationcitrixauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcitrixauthactionName := d.Get("name").(string)
	authenticationcitrixauthaction := authentication.Authenticationcitrixauthaction{
		Authentication:     d.Get("authentication").(string),
		Authenticationtype: d.Get("authenticationtype").(string),
		Name:               d.Get("name").(string),
	}

	_, err := client.AddResource("authenticationcitrixauthaction", authenticationcitrixauthactionName, &authenticationcitrixauthaction)
	if err != nil {
		return err
	}

	d.SetId(authenticationcitrixauthactionName)

	err = readAuthenticationcitrixauthactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationcitrixauthaction but we can't read it ?? %s", authenticationcitrixauthactionName)
		return nil
	}
	return nil
}

func readAuthenticationcitrixauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationcitrixauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcitrixauthactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationcitrixauthaction state %s", authenticationcitrixauthactionName)
	data, err := client.FindResource("authenticationcitrixauthaction", authenticationcitrixauthactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationcitrixauthaction state %s", authenticationcitrixauthactionName)
		d.SetId("")
		return nil
	}
	d.Set("authentication", data["authentication"])
	d.Set("authenticationtype", data["authenticationtype"])
	d.Set("name", data["name"])

	return nil

}

func updateAuthenticationcitrixauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationcitrixauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcitrixauthactionName := d.Get("name").(string)

	authenticationcitrixauthaction := authentication.Authenticationcitrixauthaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for authenticationcitrixauthaction %s, starting update", authenticationcitrixauthactionName)
		authenticationcitrixauthaction.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authenticationtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Authenticationtype has changed for authenticationcitrixauthaction %s, starting update", authenticationcitrixauthactionName)
		authenticationcitrixauthaction.Authenticationtype = d.Get("authenticationtype").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationcitrixauthaction", authenticationcitrixauthactionName, &authenticationcitrixauthaction)
		if err != nil {
			return fmt.Errorf("Error updating authenticationcitrixauthaction %s", authenticationcitrixauthactionName)
		}
	}
	return readAuthenticationcitrixauthactionFunc(d, meta)
}

func deleteAuthenticationcitrixauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationcitrixauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcitrixauthactionName := d.Id()
	err := client.DeleteResource("authenticationcitrixauthaction", authenticationcitrixauthactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
