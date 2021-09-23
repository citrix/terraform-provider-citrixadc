package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcVpneula() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpneulaFunc,
		Read:          readVpneulaFunc,
		Delete:        deleteVpneulaFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createVpneulaFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId(vpneulaName)

	err = readVpneulaFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpneula but we can't read it ?? %s", vpneulaName)
		return nil
	}
	return nil
}

func readVpneulaFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("name", data["name"])

	return nil

}

func deleteVpneulaFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpneulaFunc")
	client := meta.(*NetScalerNitroClient).client
	vpneulaName := d.Id()
	err := client.DeleteResource(service.Vpneula.Type(), vpneulaName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
