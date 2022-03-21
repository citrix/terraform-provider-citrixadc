package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationemailaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationemailactionFunc,
		Read:          readAuthenticationemailactionFunc,
		Update:        updateAuthenticationemailactionFunc,
		Delete:        deleteAuthenticationemailactionFunc,
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
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"serverurl": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationemailactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationemailactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationemailactionName := d.Get("name").(string)
	authenticationemailaction := authentication.Authenticationemailaction{
		Content:                    d.Get("content").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Emailaddress:               d.Get("emailaddress").(string),
		Name:                       d.Get("name").(string),
		Password:                   d.Get("password").(string),
		Serverurl:                  d.Get("serverurl").(string),
		Timeout:                    d.Get("timeout").(int),
		Type:                       d.Get("type").(string),
		Username:                   d.Get("username").(string),
	}

	_, err := client.AddResource("authenticationemailaction", authenticationemailactionName, &authenticationemailaction)
	if err != nil {
		return err
	}

	d.SetId(authenticationemailactionName)

	err = readAuthenticationemailactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationemailaction but we can't read it ?? %s", authenticationemailactionName)
		return nil
	}
	return nil
}

func readAuthenticationemailactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationemailactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationemailactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationemailaction state %s", authenticationemailactionName)
	data, err := client.FindResource("authenticationemailaction", authenticationemailactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationemailaction state %s", authenticationemailactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("content", data["content"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("emailaddress", data["emailaddress"])
	d.Set("name", data["name"])
	//d.Set("password", data["password"]) encrypted value recieved from ADC
	d.Set("serverurl", data["serverurl"])
	//d.Set("timeout", data["timeout"]) not recieved from ADC
	d.Set("type", data["type"])
	d.Set("username", data["username"])

	return nil

}

func updateAuthenticationemailactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationemailactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationemailactionName := d.Get("name").(string)

	authenticationemailaction := authentication.Authenticationemailaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("content") {
		log.Printf("[DEBUG]  citrixadc-provider: Content has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Content = d.Get("content").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("emailaddress") {
		log.Printf("[DEBUG]  citrixadc-provider: Emailaddress has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Emailaddress = d.Get("emailaddress").(string)
		hasChange = true
	}
	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("serverurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverurl has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Serverurl = d.Get("serverurl").(string)
		hasChange = true
	}
	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Timeout = d.Get("timeout").(int)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("username") {
		log.Printf("[DEBUG]  citrixadc-provider: Username has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Username = d.Get("username").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationemailaction", authenticationemailactionName, &authenticationemailaction)
		if err != nil {
			return fmt.Errorf("Error updating authenticationemailaction %s", authenticationemailactionName)
		}
	}
	return readAuthenticationemailactionFunc(d, meta)
}

func deleteAuthenticationemailactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationemailactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationemailactionName := d.Id()
	err := client.DeleteResource("authenticationemailaction", authenticationemailactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
