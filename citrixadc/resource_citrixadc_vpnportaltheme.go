package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVpnportaltheme() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnportalthemeFunc,
		ReadContext:   readVpnportalthemeFunc,
		DeleteContext: deleteVpnportalthemeFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"basetheme": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createVpnportalthemeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnportalthemeFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnportalthemeName := d.Get("name").(string)
	vpnportaltheme := vpn.Vpnportaltheme{
		Basetheme: d.Get("basetheme").(string),
		Name:      d.Get("name").(string),
	}

	_, err := client.AddResource(service.Vpnportaltheme.Type(), vpnportalthemeName, &vpnportaltheme)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vpnportalthemeName)

	return readVpnportalthemeFunc(ctx, d, meta)
}

func readVpnportalthemeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnportalthemeFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnportalthemeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnportaltheme state %s", vpnportalthemeName)
	data, err := client.FindResource(service.Vpnportaltheme.Type(), vpnportalthemeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnportaltheme state %s", vpnportalthemeName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("basetheme", data["basetheme"])

	return nil

}

func deleteVpnportalthemeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnportalthemeFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnportalthemeName := d.Id()
	err := client.DeleteResource(service.Vpnportaltheme.Type(), vpnportalthemeName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
