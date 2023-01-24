package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnformssoaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnformssoactionFunc,
		Read:          readVpnformssoactionFunc,
		Update:        updateVpnformssoactionFunc,
		Delete:        deleteVpnformssoactionFunc,
		Schema: map[string]*schema.Schema{
			"actionurl": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namevaluepair": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nvtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passwdfield": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"responsesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssosuccessrule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"submitmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"userfield": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnformssoactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnformssoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	var vpnformssoactionName string
	if v, ok := d.GetOk("name"); ok {
		vpnformssoactionName = v.(string)
	} else {
		vpnformssoactionName = resource.PrefixedUniqueId("tf-vpnformssoaction-")
		d.Set("name", vpnformssoactionName)
	}
	vpnformssoaction := vpn.Vpnformssoaction{
		Actionurl:      d.Get("actionurl").(string),
		Name:           d.Get("name").(string),
		Namevaluepair:  d.Get("namevaluepair").(string),
		Nvtype:         d.Get("nvtype").(string),
		Passwdfield:    d.Get("passwdfield").(string),
		Responsesize:   d.Get("responsesize").(int),
		Ssosuccessrule: d.Get("ssosuccessrule").(string),
		Submitmethod:   d.Get("submitmethod").(string),
		Userfield:      d.Get("userfield").(string),
	}

	_, err := client.AddResource(service.Vpnformssoaction.Type(), vpnformssoactionName, &vpnformssoaction)
	if err != nil {
		return err
	}

	d.SetId(vpnformssoactionName)

	err = readVpnformssoactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnformssoaction but we can't read it ?? %s", vpnformssoactionName)
		return nil
	}
	return nil
}

func readVpnformssoactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnformssoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnformssoactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnformssoaction state %s", vpnformssoactionName)
	data, err := client.FindResource(service.Vpnformssoaction.Type(), vpnformssoactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnformssoaction state %s", vpnformssoactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("actionurl", data["actionurl"])
	d.Set("name", data["name"])
	d.Set("namevaluepair", data["namevaluepair"])
	d.Set("nvtype", data["nvtype"])
	d.Set("passwdfield", data["passwdfield"])
	d.Set("responsesize", data["responsesize"])
	d.Set("ssosuccessrule", data["ssosuccessrule"])
	d.Set("submitmethod", data["submitmethod"])
	d.Set("userfield", data["userfield"])

	return nil

}

func updateVpnformssoactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnformssoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnformssoactionName := d.Get("name").(string)

	vpnformssoaction := vpn.Vpnformssoaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("actionurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Actionurl has changed for vpnformssoaction %s, starting update", vpnformssoactionName)
		vpnformssoaction.Actionurl = d.Get("actionurl").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for vpnformssoaction %s, starting update", vpnformssoactionName)
		vpnformssoaction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("namevaluepair") {
		log.Printf("[DEBUG]  citrixadc-provider: Namevaluepair has changed for vpnformssoaction %s, starting update", vpnformssoactionName)
		vpnformssoaction.Namevaluepair = d.Get("namevaluepair").(string)
		hasChange = true
	}
	if d.HasChange("nvtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Nvtype has changed for vpnformssoaction %s, starting update", vpnformssoactionName)
		vpnformssoaction.Nvtype = d.Get("nvtype").(string)
		hasChange = true
	}
	if d.HasChange("passwdfield") {
		log.Printf("[DEBUG]  citrixadc-provider: Passwdfield has changed for vpnformssoaction %s, starting update", vpnformssoactionName)
		vpnformssoaction.Passwdfield = d.Get("passwdfield").(string)
		hasChange = true
	}
	if d.HasChange("responsesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Responsesize has changed for vpnformssoaction %s, starting update", vpnformssoactionName)
		vpnformssoaction.Responsesize = d.Get("responsesize").(int)
		hasChange = true
	}
	if d.HasChange("ssosuccessrule") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssosuccessrule has changed for vpnformssoaction %s, starting update", vpnformssoactionName)
		vpnformssoaction.Ssosuccessrule = d.Get("ssosuccessrule").(string)
		hasChange = true
	}
	if d.HasChange("submitmethod") {
		log.Printf("[DEBUG]  citrixadc-provider: Submitmethod has changed for vpnformssoaction %s, starting update", vpnformssoactionName)
		vpnformssoaction.Submitmethod = d.Get("submitmethod").(string)
		hasChange = true
	}
	if d.HasChange("userfield") {
		log.Printf("[DEBUG]  citrixadc-provider: Userfield has changed for vpnformssoaction %s, starting update", vpnformssoactionName)
		vpnformssoaction.Userfield = d.Get("userfield").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Vpnformssoaction.Type(), vpnformssoactionName, &vpnformssoaction)
		if err != nil {
			return fmt.Errorf("Error updating vpnformssoaction %s", vpnformssoactionName)
		}
	}
	return readVpnformssoactionFunc(d, meta)
}

func deleteVpnformssoactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnformssoactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnformssoactionName := d.Id()
	err := client.DeleteResource(service.Vpnformssoaction.Type(), vpnformssoactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
