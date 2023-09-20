package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnpcoipvserverprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnpcoipvserverprofileFunc,
		Read:          readVpnpcoipvserverprofileFunc,
		Update:        updateVpnpcoipvserverprofileFunc,
		Delete:        deleteVpnpcoipvserverprofileFunc,
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

func createVpnpcoipvserverprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnpcoipvserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnpcoipvserverprofileName := d.Get("name").(string)
	vpnpcoipvserverprofile := vpn.Vpnpcoipvserverprofile{
		Logindomain: d.Get("logindomain").(string),
		Name:        d.Get("name").(string),
		Udpport:     d.Get("udpport").(int),
	}

	_, err := client.AddResource("vpnpcoipvserverprofile", vpnpcoipvserverprofileName, &vpnpcoipvserverprofile)
	if err != nil {
		return err
	}

	d.SetId(vpnpcoipvserverprofileName)

	err = readVpnpcoipvserverprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnpcoipvserverprofile but we can't read it ?? %s", vpnpcoipvserverprofileName)
		return nil
	}
	return nil
}

func readVpnpcoipvserverprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("udpport", data["udpport"])

	return nil

}

func updateVpnpcoipvserverprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
		vpnpcoipvserverprofile.Udpport = d.Get("udpport").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("vpnpcoipvserverprofile", vpnpcoipvserverprofileName, &vpnpcoipvserverprofile)
		if err != nil {
			return fmt.Errorf("Error updating vpnpcoipvserverprofile %s", vpnpcoipvserverprofileName)
		}
	}
	return readVpnpcoipvserverprofileFunc(d, meta)
}

func deleteVpnpcoipvserverprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnpcoipvserverprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnpcoipvserverprofileName := d.Id()
	err := client.DeleteResource("vpnpcoipvserverprofile", vpnpcoipvserverprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
