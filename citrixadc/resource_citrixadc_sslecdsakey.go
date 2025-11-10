package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSslecdsakey() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslecdsakeyFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"curve": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"aes256": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"des": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"des3": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"keyfile": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"keyform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"pkcs8": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  false,
				ForceNew:  true,
				Sensitive: true,
			},
		},
	}
}

func createSslecdsakeyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslecdsakeyFunc")
	client := meta.(*NetScalerNitroClient).client

	sslecdsakeyName := resource.PrefixedUniqueId("tf-sslecdsakey-")
	sslecdsakey := ssl.Sslecdsakey{
		Curve:    d.Get("curve").(string),
		Des:      d.Get("des").(bool),
		Des3:     d.Get("des3").(bool),
		Aes256:   d.Get("aes256").(bool),
		Pkcs8:    d.Get("pkcs8").(bool),
		Password: d.Get("password").(string),
		Keyfile:  d.Get("keyfile").(string),
		Keyform:  d.Get("keyform").(string),
	}

	err := client.ActOnResource("sslecdsakey", &sslecdsakey, "create")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sslecdsakeyName)

	return nil
}
