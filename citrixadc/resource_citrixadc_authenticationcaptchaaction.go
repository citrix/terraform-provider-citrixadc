package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationcaptchaaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationcaptchaactionFunc,
		Read:          readAuthenticationcaptchaactionFunc,
		Update:        updateAuthenticationcaptchaactionFunc,
		Delete:        deleteAuthenticationcaptchaactionFunc,
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
			"secretkey": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"sitekey": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationcaptchaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationcaptchaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcaptchaactionName := d.Get("name").(string)
	authenticationcaptchaaction := authentication.Authenticationcaptchaaction{
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Name:                       d.Get("name").(string),
		Secretkey:                  d.Get("secretkey").(string),
		Serverurl:                  d.Get("serverurl").(string),
		Sitekey:                    d.Get("sitekey").(string),
	}

	_, err := client.AddResource("authenticationcaptchaaction", authenticationcaptchaactionName, &authenticationcaptchaaction)
	if err != nil {
		return err
	}

	d.SetId(authenticationcaptchaactionName)

	err = readAuthenticationcaptchaactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationcaptchaaction but we can't read it ?? %s", authenticationcaptchaactionName)
		return nil
	}
	return nil
}

func readAuthenticationcaptchaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationcaptchaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcaptchaactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationcaptchaaction state %s", authenticationcaptchaactionName)
	data, err := client.FindResource("authenticationcaptchaaction", authenticationcaptchaactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationcaptchaaction state %s", authenticationcaptchaactionName)
		d.SetId("")
		return nil
	}
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("name", data["name"])
	// d.Set("secretkey", data["secretkey"])
	d.Set("serverurl", data["serverurl"])
	// d.Set("sitekey", data["sitekey"])

	return nil

}

func updateAuthenticationcaptchaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationcaptchaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcaptchaactionName := d.Get("name").(string)

	authenticationcaptchaaction := authentication.Authenticationcaptchaaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationcaptchaaction %s, starting update", authenticationcaptchaactionName)
		authenticationcaptchaaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("secretkey") {
		log.Printf("[DEBUG]  citrixadc-provider: Secretkey has changed for authenticationcaptchaaction %s, starting update", authenticationcaptchaactionName)
		authenticationcaptchaaction.Secretkey = d.Get("secretkey").(string)
		hasChange = true
	}
	if d.HasChange("serverurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverurl has changed for authenticationcaptchaaction %s, starting update", authenticationcaptchaactionName)
		authenticationcaptchaaction.Serverurl = d.Get("serverurl").(string)
		hasChange = true
	}
	if d.HasChange("sitekey") {
		log.Printf("[DEBUG]  citrixadc-provider: Sitekey has changed for authenticationcaptchaaction %s, starting update", authenticationcaptchaactionName)
		authenticationcaptchaaction.Sitekey = d.Get("sitekey").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationcaptchaaction", authenticationcaptchaactionName, &authenticationcaptchaaction)
		if err != nil {
			return fmt.Errorf("Error updating authenticationcaptchaaction %s", authenticationcaptchaactionName)
		}
	}
	return readAuthenticationcaptchaactionFunc(d, meta)
}

func deleteAuthenticationcaptchaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationcaptchaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcaptchaactionName := d.Id()
	err := client.DeleteResource("authenticationcaptchaaction", authenticationcaptchaactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
