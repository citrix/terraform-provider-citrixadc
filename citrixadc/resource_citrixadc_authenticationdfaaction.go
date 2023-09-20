package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationdfaaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationdfaactionFunc,
		Read:          readAuthenticationdfaactionFunc,
		Update:        updateAuthenticationdfaactionFunc,
		Delete:        deleteAuthenticationdfaactionFunc,
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
			"passphrase": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"serverurl": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"clientid": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationdfaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationdfaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationdfaactionName := d.Get("name").(string)
	authenticationdfaaction := authentication.Authenticationdfaaction{
		Clientid:                   d.Get("clientid").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Name:                       d.Get("name").(string),
		Passphrase:                 d.Get("passphrase").(string),
		Serverurl:                  d.Get("serverurl").(string),
	}

	_, err := client.AddResource(service.Authenticationdfaaction.Type(), authenticationdfaactionName, &authenticationdfaaction)
	if err != nil {
		return err
	}

	d.SetId(authenticationdfaactionName)

	err = readAuthenticationdfaactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationdfaaction but we can't read it ?? %s", authenticationdfaactionName)
		return nil
	}
	return nil
}

func readAuthenticationdfaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationdfaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationdfaactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationdfaaction state %s", authenticationdfaactionName)
	data, err := client.FindResource(service.Authenticationdfaaction.Type(), authenticationdfaactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationdfaaction state %s", authenticationdfaactionName)
		d.SetId("")
		return nil
	}
	d.Set("clientid", data["clientid"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("name", data["name"])
	// d.Set("passphrase", data["passphrase"]) Encrypted value is returned each time
	d.Set("serverurl", data["serverurl"])

	return nil

}

func updateAuthenticationdfaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationdfaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationdfaactionName := d.Get("name").(string)

	authenticationdfaaction := authentication.Authenticationdfaaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("clientid") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientid has changed for authenticationdfaaction %s, starting update", authenticationdfaactionName)
		authenticationdfaaction.Clientid = d.Get("clientid").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationdfaaction %s, starting update", authenticationdfaactionName)
		authenticationdfaaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("passphrase") {
		log.Printf("[DEBUG]  citrixadc-provider: Passphrase has changed for authenticationdfaaction %s, starting update", authenticationdfaactionName)
		authenticationdfaaction.Passphrase = d.Get("passphrase").(string)
		hasChange = true
	}
	if d.HasChange("serverurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverurl has changed for authenticationdfaaction %s, starting update", authenticationdfaactionName)
		authenticationdfaaction.Serverurl = d.Get("serverurl").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationdfaaction.Type(), authenticationdfaactionName, &authenticationdfaaction)
		if err != nil {
			return fmt.Errorf("Error updating authenticationdfaaction %s", authenticationdfaactionName)
		}
	}
	return readAuthenticationdfaactionFunc(d, meta)
}

func deleteAuthenticationdfaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationdfaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationdfaactionName := d.Id()
	err := client.DeleteResource(service.Authenticationdfaaction.Type(), authenticationdfaactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
