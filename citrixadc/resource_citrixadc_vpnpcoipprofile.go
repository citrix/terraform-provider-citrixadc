package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnpcoipprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnpcoipprofileFunc,
		Read:          readVpnpcoipprofileFunc,
		Update:        updateVpnpcoipprofileFunc,
		Delete:        deleteVpnpcoipprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"conserverurl": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"icvverification": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionidletimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVpnpcoipprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnpcoipprofileFunc")
	client := meta.(*NetScalerNitroClient).client

	vpnpcoipprofileName := d.Get("name").(string)
	vpnpcoipprofile := vpn.Vpnpcoipprofile{
		Conserverurl:       d.Get("conserverurl").(string),
		Icvverification:    d.Get("icvverification").(string),
		Name:               d.Get("name").(string),
		Sessionidletimeout: d.Get("sessionidletimeout").(int),
	}

	_, err := client.AddResource("vpnpcoipprofile", vpnpcoipprofileName, &vpnpcoipprofile)
	if err != nil {
		return err
	}

	d.SetId(vpnpcoipprofileName)

	err = readVpnpcoipprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnpcoipprofile but we can't read it ?? %s", vpnpcoipprofileName)
		return nil
	}
	return nil
}

func readVpnpcoipprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnpcoipprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnpcoipprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnpcoipprofile state %s", vpnpcoipprofileName)
	data, err := client.FindResource("vpnpcoipprofile", vpnpcoipprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnpcoipprofile state %s", vpnpcoipprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("conserverurl", data["conserverurl"])
	d.Set("icvverification", data["icvverification"])
	d.Set("name", data["name"])
	d.Set("sessionidletimeout", data["sessionidletimeout"])

	return nil

}

func updateVpnpcoipprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnpcoipprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnpcoipprofileName := d.Get("name").(string)

	vpnpcoipprofile := vpn.Vpnpcoipprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("conserverurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Conserverurl has changed for vpnpcoipprofile %s, starting update", vpnpcoipprofileName)
		vpnpcoipprofile.Conserverurl = d.Get("conserverurl").(string)
		hasChange = true
	}
	if d.HasChange("icvverification") {
		log.Printf("[DEBUG]  citrixadc-provider: Icvverification has changed for vpnpcoipprofile %s, starting update", vpnpcoipprofileName)
		vpnpcoipprofile.Icvverification = d.Get("icvverification").(string)
		hasChange = true
	}
	if d.HasChange("sessionidletimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionidletimeout has changed for vpnpcoipprofile %s, starting update", vpnpcoipprofileName)
		vpnpcoipprofile.Sessionidletimeout = d.Get("sessionidletimeout").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("vpnpcoipprofile", vpnpcoipprofileName, &vpnpcoipprofile)
		if err != nil {
			return fmt.Errorf("Error updating vpnpcoipprofile %s", vpnpcoipprofileName)
		}
	}
	return readVpnpcoipprofileFunc(d, meta)
}

func deleteVpnpcoipprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnpcoipprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnpcoipprofileName := d.Id()
	err := client.DeleteResource("vpnpcoipprofile", vpnpcoipprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
