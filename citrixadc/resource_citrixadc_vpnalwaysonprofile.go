package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnalwaysonprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnalwaysonprofileFunc,
		Read:          readVpnalwaysonprofileFunc,
		Update:        updateVpnalwaysonprofileFunc,
		Delete:        deleteVpnalwaysonprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"clientcontrol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"locationbasedvpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"networkaccessonvpnfailure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVpnalwaysonprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId(vpnalwaysonprofileName)

	err = readVpnalwaysonprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnalwaysonprofile but we can't read it ?? %s", vpnalwaysonprofileName)
		return nil
	}
	return nil
}

func readVpnalwaysonprofileFunc(d *schema.ResourceData, meta interface{}) error {
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

func updateVpnalwaysonprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
			return fmt.Errorf("Error updating vpnalwaysonprofile %s", vpnalwaysonprofileName)
		}
	}
	return readVpnalwaysonprofileFunc(d, meta)
}

func deleteVpnalwaysonprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnalwaysonprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnalwaysonprofileName := d.Id()
	err := client.DeleteResource("vpnalwaysonprofile", vpnalwaysonprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
