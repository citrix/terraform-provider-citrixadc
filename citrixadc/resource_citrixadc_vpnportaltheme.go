package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcVpnportaltheme() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnportalthemeFunc,
		Read:          readVpnportalthemeFunc,
		Delete:        deleteVpnportalthemeFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"basetheme": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createVpnportalthemeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnportalthemeFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnportalthemeName := d.Get("name").(string)
	vpnportaltheme := vpn.Vpnportaltheme{
		Basetheme: d.Get("basetheme").(string),
		Name:      d.Get("name").(string),
	}

	_, err := client.AddResource(service.Vpnportaltheme.Type(), vpnportalthemeName, &vpnportaltheme)
	if err != nil {
		return err
	}

	d.SetId(vpnportalthemeName)

	err = readVpnportalthemeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnportaltheme but we can't read it ?? %s", vpnportalthemeName)
		return nil
	}
	return nil
}

func readVpnportalthemeFunc(d *schema.ResourceData, meta interface{}) error {
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

func deleteVpnportalthemeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnportalthemeFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnportalthemeName := d.Id()
	err := client.DeleteResource(service.Vpnportaltheme.Type(), vpnportalthemeName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
