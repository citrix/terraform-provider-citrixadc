package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpntrafficpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpntrafficpolicyFunc,
		Read:          readVpntrafficpolicyFunc,
		Update:        updateVpntrafficpolicyFunc,
		Delete:        deleteVpntrafficpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createVpntrafficpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpntrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpntrafficpolicyName := d.Get("name").(string)
	vpntrafficpolicy := vpn.Vpntrafficpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Vpntrafficpolicy.Type(), vpntrafficpolicyName, &vpntrafficpolicy)
	if err != nil {
		return err
	}

	d.SetId(vpntrafficpolicyName)

	err = readVpntrafficpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpntrafficpolicy but we can't read it ?? %s", vpntrafficpolicyName)
		return nil
	}
	return nil
}

func readVpntrafficpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpntrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpntrafficpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpntrafficpolicy state %s", vpntrafficpolicyName)
	data, err := client.FindResource(service.Vpntrafficpolicy.Type(), vpntrafficpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpntrafficpolicy state %s", vpntrafficpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateVpntrafficpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpntrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpntrafficpolicyName := d.Get("name").(string)

	vpntrafficpolicy := vpn.Vpntrafficpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for vpntrafficpolicy %s, starting update", vpntrafficpolicyName)
		vpntrafficpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for vpntrafficpolicy %s, starting update", vpntrafficpolicyName)
		vpntrafficpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if hasChange {
		_, err := client.UpdateResource(service.Vpntrafficpolicy.Type(), vpntrafficpolicyName, &vpntrafficpolicy)
		if err != nil {
			return fmt.Errorf("Error updating vpntrafficpolicy %s", vpntrafficpolicyName)
		}
	}
	return readVpntrafficpolicyFunc(d, meta)
}

func deleteVpntrafficpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpntrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpntrafficpolicyName := d.Id()
	err := client.DeleteResource(service.Vpntrafficpolicy.Type(), vpntrafficpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
