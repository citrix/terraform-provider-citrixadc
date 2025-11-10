package citrixadc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	_ "fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcPasswordResetter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPasswordResetter,
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

func createPasswordResetter(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(id)

	return nil
}

type resetterPayload struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	New_password string `json:"new_password,omitempty"`
}
