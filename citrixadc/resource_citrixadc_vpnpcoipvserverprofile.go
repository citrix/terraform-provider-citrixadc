package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVpnpcoipvserverprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnpcoipvserverprofileFunc,
		ReadContext:   readVpnpcoipvserverprofileFunc,
		UpdateContext: updateVpnpcoipvserverprofileFunc,
		DeleteContext: deleteVpnpcoipvserverprofileFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"logindomain": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"udpport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVpnpcoipvserverprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnpcoipvserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnpcoipvserverprofileName := d.Get("name").(string)
	vpnpcoipvserverprofile := vpn.Vpnpcoipvserverprofile{
		Logindomain: d.Get("logindomain").(string),
		Name:        d.Get("name").(string),
	}

	if raw := d.GetRawConfig().GetAttr("udpport"); !raw.IsNull() {
		vpnpcoipvserverprofile.Udpport = intPtr(d.Get("udpport").(int))
	}

	_, err := client.AddResource("vpnpcoipvserverprofile", vpnpcoipvserverprofileName, &vpnpcoipvserverprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vpnpcoipvserverprofileName)

	return readVpnpcoipvserverprofileFunc(ctx, d, meta)
}

func readVpnpcoipvserverprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnpcoipvserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnpcoipvserverprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnpcoipvserverprofile state %s", vpnpcoipvserverprofileName)
	data, err := client.FindResource("vpnpcoipvserverprofile", vpnpcoipvserverprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnpcoipvserverprofile state %s", vpnpcoipvserverprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("logindomain", data["logindomain"])
	d.Set("name", data["name"])
	setToInt("udpport", d, data["udpport"])

	return nil

}

func updateVpnpcoipvserverprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnpcoipvserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnpcoipvserverprofileName := d.Get("name").(string)

	vpnpcoipvserverprofile := vpn.Vpnpcoipvserverprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("logindomain") {
		log.Printf("[DEBUG]  citrixadc-provider: Logindomain has changed for vpnpcoipvserverprofile %s, starting update", vpnpcoipvserverprofileName)
		vpnpcoipvserverprofile.Logindomain = d.Get("logindomain").(string)
		hasChange = true
	}
	if d.HasChange("udpport") {
		log.Printf("[DEBUG]  citrixadc-provider: Udpport has changed for vpnpcoipvserverprofile %s, starting update", vpnpcoipvserverprofileName)
		vpnpcoipvserverprofile.Udpport = intPtr(d.Get("udpport").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("vpnpcoipvserverprofile", vpnpcoipvserverprofileName, &vpnpcoipvserverprofile)
		if err != nil {
			return diag.Errorf("Error updating vpnpcoipvserverprofile %s", vpnpcoipvserverprofileName)
		}
	}
	return readVpnpcoipvserverprofileFunc(ctx, d, meta)
}

func deleteVpnpcoipvserverprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnpcoipvserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnpcoipvserverprofileName := d.Id()
	err := client.DeleteResource("vpnpcoipvserverprofile", vpnpcoipvserverprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
