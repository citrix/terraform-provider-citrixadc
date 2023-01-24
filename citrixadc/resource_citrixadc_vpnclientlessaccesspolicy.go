package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnclientlessaccesspolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnclientlessaccesspolicyFunc,
		Read:          readVpnclientlessaccesspolicyFunc,
		Update:        updateVpnclientlessaccesspolicyFunc,
		Delete:        deleteVpnclientlessaccesspolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"profilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVpnclientlessaccesspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnclientlessaccesspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	var vpnclientlessaccesspolicyName string
	if v, ok := d.GetOk("name"); ok {
		vpnclientlessaccesspolicyName = v.(string)
	} else {
		vpnclientlessaccesspolicyName = resource.PrefixedUniqueId("tf-vpnclientlessaccesspolicy-")
		d.Set("name", vpnclientlessaccesspolicyName)
	}
	vpnclientlessaccesspolicy := vpn.Vpnclientlessaccesspolicy{
		Name:        d.Get("name").(string),
		Profilename: d.Get("profilename").(string),
		Rule:        d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Vpnclientlessaccesspolicy.Type(), vpnclientlessaccesspolicyName, &vpnclientlessaccesspolicy)
	if err != nil {
		return err
	}

	d.SetId(vpnclientlessaccesspolicyName)

	err = readVpnclientlessaccesspolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnclientlessaccesspolicy but we can't read it ?? %s", vpnclientlessaccesspolicyName)
		return nil
	}
	return nil
}

func readVpnclientlessaccesspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnclientlessaccesspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnclientlessaccesspolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnclientlessaccesspolicy state %s", vpnclientlessaccesspolicyName)
	data, err := client.FindResource(service.Vpnclientlessaccesspolicy.Type(), vpnclientlessaccesspolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnclientlessaccesspolicy state %s", vpnclientlessaccesspolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("name", data["name"])
	d.Set("profilename", data["profilename"])
	d.Set("rule", data["rule"])

	return nil

}

func updateVpnclientlessaccesspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnclientlessaccesspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnclientlessaccesspolicyName := d.Get("name").(string)

	vpnclientlessaccesspolicy := vpn.Vpnclientlessaccesspolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for vpnclientlessaccesspolicy %s, starting update", vpnclientlessaccesspolicyName)
		vpnclientlessaccesspolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("profilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Profilename has changed for vpnclientlessaccesspolicy %s, starting update", vpnclientlessaccesspolicyName)
		vpnclientlessaccesspolicy.Profilename = d.Get("profilename").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for vpnclientlessaccesspolicy %s, starting update", vpnclientlessaccesspolicyName)
		vpnclientlessaccesspolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Vpnclientlessaccesspolicy.Type(), vpnclientlessaccesspolicyName, &vpnclientlessaccesspolicy)
		if err != nil {
			return fmt.Errorf("Error updating vpnclientlessaccesspolicy %s", vpnclientlessaccesspolicyName)
		}
	}
	return readVpnclientlessaccesspolicyFunc(d, meta)
}

func deleteVpnclientlessaccesspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnclientlessaccesspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnclientlessaccesspolicyName := d.Id()
	err := client.DeleteResource(service.Vpnclientlessaccesspolicy.Type(), vpnclientlessaccesspolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
