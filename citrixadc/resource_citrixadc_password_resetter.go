package citrixadc

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	_ "fmt"
	"log"
)

func resourceCitrixAdcPasswordResetter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPasswordResetter,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"new_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
			},
		},
	}
}

func createPasswordResetter(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPasswordResetter")
	client := meta.(*NetScalerNitroClient).client
	id := resource.PrefixedUniqueId("tf-password-resetter-")

	payload := resetterPayload{
		Username:     d.Get("username").(string),
		Password:     d.Get("password").(string),
		New_password: d.Get("new_password").(string),
	}

	_, err := client.AddResource("login", "", &payload)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

type resetterPayload struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	New_password string `json:"new_password,omitempty"`
}
