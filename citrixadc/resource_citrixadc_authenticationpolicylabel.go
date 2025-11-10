package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAuthenticationpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationpolicylabelFunc,
		ReadContext:   readAuthenticationpolicylabelFunc,
		DeleteContext: deleteAuthenticationpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"loginschema": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAuthenticationpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpolicylabelName := d.Get("labelname").(string)
	authenticationpolicylabel := authentication.Authenticationpolicylabel{
		Comment:     d.Get("comment").(string),
		Labelname:   d.Get("labelname").(string),
		Loginschema: d.Get("loginschema").(string),
		Type:        d.Get("type").(string),
	}

	_, err := client.AddResource(service.Authenticationpolicylabel.Type(), authenticationpolicylabelName, &authenticationpolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationpolicylabelName)

	return readAuthenticationpolicylabelFunc(ctx, d, meta)
}

func readAuthenticationpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationpolicylabel state %s", authenticationpolicylabelName)
	data, err := client.FindResource(service.Authenticationpolicylabel.Type(), authenticationpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationpolicylabel state %s", authenticationpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("comment", data["comment"])
	d.Set("labelname", data["labelname"])
	d.Set("loginschema", data["loginschema"])
	d.Set("type", data["type"])

	return nil

}

func deleteAuthenticationpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationpolicylabelName := d.Id()
	err := client.DeleteResource(service.Authenticationpolicylabel.Type(), authenticationpolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
