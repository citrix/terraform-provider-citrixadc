package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnsessionpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnsessionpolicyFunc,
		Read:          readVpnsessionpolicyFunc,
		Update:        updateVpnsessionpolicyFunc,
		Delete:        deleteVpnsessionpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createVpnsessionpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnsessionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnsessionpolicyName := d.Get("name").(string)
	vpnsessionpolicy := vpn.Vpnsessionpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Vpnsessionpolicy.Type(), vpnsessionpolicyName, &vpnsessionpolicy)
	if err != nil {
		return err
	}

	d.SetId(vpnsessionpolicyName)

	err = readVpnsessionpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnsessionpolicy but we can't read it ?? %s", vpnsessionpolicyName)
		return nil
	}
	return nil
}

func readVpnsessionpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnsessionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnsessionpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnsessionpolicy state %s", vpnsessionpolicyName)
	data, err := client.FindResource(service.Vpnsessionpolicy.Type(), vpnsessionpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnsessionpolicy state %s", vpnsessionpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateVpnsessionpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnsessionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnsessionpolicyName := d.Get("name").(string)

	vpnsessionpolicy := vpn.Vpnsessionpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for vpnsessionpolicy %s, starting update", vpnsessionpolicyName)
		vpnsessionpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for vpnsessionpolicy %s, starting update", vpnsessionpolicyName)
		vpnsessionpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Vpnsessionpolicy.Type(), vpnsessionpolicyName, &vpnsessionpolicy)
		if err != nil {
			return fmt.Errorf("Error updating vpnsessionpolicy %s", vpnsessionpolicyName)
		}
	}
	return readVpnsessionpolicyFunc(d, meta)
}

func deleteVpnsessionpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnsessionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnsessionpolicyName := d.Id()
	err := client.DeleteResource(service.Vpnsessionpolicy.Type(), vpnsessionpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
