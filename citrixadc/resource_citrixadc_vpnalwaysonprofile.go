package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcVpnalwaysonprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnalwaysonprofileFunc,
		ReadContext:   readVpnalwaysonprofileFunc,
		UpdateContext: updateVpnalwaysonprofileFunc,
		DeleteContext: deleteVpnalwaysonprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"clientcontrol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"locationbasedvpn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"networkaccessonvpnfailure": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVpnalwaysonprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnalwaysonprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	var vpnalwaysonprofileName string
	if v, ok := d.GetOk("name"); ok {
		vpnalwaysonprofileName = v.(string)
	} else {
		vpnalwaysonprofileName = resource.PrefixedUniqueId("tf-vpnalwaysonprofile-")
		d.Set("name", vpnalwaysonprofileName)
	}
	vpnalwaysonprofile := vpn.Vpnalwaysonprofile{
		Clientcontrol:             d.Get("clientcontrol").(string),
		Locationbasedvpn:          d.Get("locationbasedvpn").(string),
		Name:                      d.Get("name").(string),
		Networkaccessonvpnfailure: d.Get("networkaccessonvpnfailure").(string),
	}

	_, err := client.AddResource("vpnalwaysonprofile", vpnalwaysonprofileName, &vpnalwaysonprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vpnalwaysonprofileName)

	return readVpnalwaysonprofileFunc(ctx, d, meta)
}

func readVpnalwaysonprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnalwaysonprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnalwaysonprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnalwaysonprofile state %s", vpnalwaysonprofileName)
	data, err := client.FindResource("vpnalwaysonprofile", vpnalwaysonprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnalwaysonprofile state %s", vpnalwaysonprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("clientcontrol", data["clientcontrol"])
	d.Set("locationbasedvpn", data["locationbasedvpn"])
	d.Set("networkaccessonvpnfailure", data["networkaccessonvpnfailure"])

	return nil

}

func updateVpnalwaysonprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnalwaysonprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnalwaysonprofileName := d.Get("name").(string)

	vpnalwaysonprofile := vpn.Vpnalwaysonprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("clientcontrol") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientcontrol has changed for vpnalwaysonprofile %s, starting update", vpnalwaysonprofileName)
		vpnalwaysonprofile.Clientcontrol = d.Get("clientcontrol").(string)
		hasChange = true
	}
	if d.HasChange("locationbasedvpn") {
		log.Printf("[DEBUG]  citrixadc-provider: Locationbasedvpn has changed for vpnalwaysonprofile %s, starting update", vpnalwaysonprofileName)
		vpnalwaysonprofile.Locationbasedvpn = d.Get("locationbasedvpn").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for vpnalwaysonprofile %s, starting update", vpnalwaysonprofileName)
		vpnalwaysonprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("networkaccessonvpnfailure") {
		log.Printf("[DEBUG]  citrixadc-provider: Networkaccessonvpnfailure has changed for vpnalwaysonprofile %s, starting update", vpnalwaysonprofileName)
		vpnalwaysonprofile.Networkaccessonvpnfailure = d.Get("networkaccessonvpnfailure").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("vpnalwaysonprofile", vpnalwaysonprofileName, &vpnalwaysonprofile)
		if err != nil {
			return diag.Errorf("Error updating vpnalwaysonprofile %s", vpnalwaysonprofileName)
		}
	}
	return readVpnalwaysonprofileFunc(ctx, d, meta)
}

func deleteVpnalwaysonprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnalwaysonprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnalwaysonprofileName := d.Id()
	err := client.DeleteResource("vpnalwaysonprofile", vpnalwaysonprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
