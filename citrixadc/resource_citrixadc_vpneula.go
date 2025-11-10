package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcVpneula() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpneulaFunc,
		ReadContext:   readVpneulaFunc,
		DeleteContext: deleteVpneulaFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createVpneulaFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpneulaFunc")
	client := meta.(*NetScalerNitroClient).client
	var vpneulaName string
	if v, ok := d.GetOk("name"); ok {
		vpneulaName = v.(string)
	} else {
		vpneulaName = resource.PrefixedUniqueId("tf-vpneula-")
		d.Set("name", vpneulaName)
	}
	vpneula := vpn.Vpneula{
		Name: d.Get("name").(string),
	}

	_, err := client.AddResource(service.Vpneula.Type(), vpneulaName, &vpneula)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vpneulaName)

	return readVpneulaFunc(ctx, d, meta)
}

func readVpneulaFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpneulaFunc")
	client := meta.(*NetScalerNitroClient).client
	vpneulaName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpneula state %s", vpneulaName)
	data, err := client.FindResource(service.Vpneula.Type(), vpneulaName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpneula state %s", vpneulaName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteVpneulaFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpneulaFunc")
	client := meta.(*NetScalerNitroClient).client
	vpneulaName := d.Id()
	err := client.DeleteResource(service.Vpneula.Type(), vpneulaName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
