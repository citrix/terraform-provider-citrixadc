package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationauthnprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationauthnprofileFunc,
		Read:          readAuthenticationauthnprofileFunc,
		Update:        updateAuthenticationauthnprofileFunc,
		Delete:        deleteAuthenticationauthnprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"authnvsname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"authenticationdomain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authenticationhost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authenticationlevel": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationauthnprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationauthnprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationauthnprofileName := d.Get("name").(string)
	authenticationauthnprofile := authentication.Authenticationauthnprofile{
		Authenticationdomain: d.Get("authenticationdomain").(string),
		Authenticationhost:   d.Get("authenticationhost").(string),
		Authenticationlevel:  d.Get("authenticationlevel").(int),
		Authnvsname:          d.Get("authnvsname").(string),
		Name:                 d.Get("name").(string),
	}

	_, err := client.AddResource(service.Authenticationauthnprofile.Type(), authenticationauthnprofileName, &authenticationauthnprofile)
	if err != nil {
		return err
	}

	d.SetId(authenticationauthnprofileName)

	err = readAuthenticationauthnprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationauthnprofile but we can't read it ?? %s", authenticationauthnprofileName)
		return nil
	}
	return nil
}

func readAuthenticationauthnprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationauthnprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationauthnprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationauthnprofile state %s", authenticationauthnprofileName)
	data, err := client.FindResource(service.Authenticationauthnprofile.Type(), authenticationauthnprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationauthnprofile state %s", authenticationauthnprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("authenticationdomain", data["authenticationdomain"])
	d.Set("authenticationhost", data["authenticationhost"])
	d.Set("authenticationlevel", data["authenticationlevel"])
	d.Set("authnvsname", data["authnvsname"])
	d.Set("name", data["name"])

	return nil

}

func updateAuthenticationauthnprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationauthnprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationauthnprofileName := d.Get("name").(string)

	authenticationauthnprofile := authentication.Authenticationauthnprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("authenticationdomain") {
		log.Printf("[DEBUG]  citrixadc-provider: Authenticationdomain has changed for authenticationauthnprofile %s, starting update", authenticationauthnprofileName)
		authenticationauthnprofile.Authenticationdomain = d.Get("authenticationdomain").(string)
		hasChange = true
	}
	if d.HasChange("authenticationhost") {
		log.Printf("[DEBUG]  citrixadc-provider: Authenticationhost has changed for authenticationauthnprofile %s, starting update", authenticationauthnprofileName)
		authenticationauthnprofile.Authenticationhost = d.Get("authenticationhost").(string)
		hasChange = true
	}
	if d.HasChange("authenticationlevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Authenticationlevel has changed for authenticationauthnprofile %s, starting update", authenticationauthnprofileName)
		authenticationauthnprofile.Authenticationlevel = d.Get("authenticationlevel").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationauthnprofile.Type(), authenticationauthnprofileName, &authenticationauthnprofile)
		if err != nil {
			return fmt.Errorf("Error updating authenticationauthnprofile %s", authenticationauthnprofileName)
		}
	}
	return readAuthenticationauthnprofileFunc(d, meta)
}

func deleteAuthenticationauthnprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationauthnprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationauthnprofileName := d.Id()
	err := client.DeleteResource(service.Authenticationauthnprofile.Type(), authenticationauthnprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
