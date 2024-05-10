package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcVlan() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVlanFunc,
		Read:          readVlanFunc,
		Update:        updateVlanFunc,
		Delete:        deleteVlanFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"aliasname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamicrouting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlanid": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ipv6dynamicrouting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sharing": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVlanFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vlanId := d.Get("vlanid").(int)

	vlan := network.Vlan{
		Aliasname:          d.Get("aliasname").(string),
		Dynamicrouting:     d.Get("dynamicrouting").(string),
		Id:                 d.Get("vlanid").(int),
		Ipv6dynamicrouting: d.Get("ipv6dynamicrouting").(string),
		Mtu:                d.Get("mtu").(int),
		Sharing:            d.Get("sharing").(string),
	}

	vlanIdStr := strconv.Itoa(vlanId)
	_, err := client.AddResource(service.Vlan.Type(), vlanIdStr, &vlan)
	if err != nil {
		return err
	}

	d.SetId(vlanIdStr)

	err = readVlanFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vlan but we can't read it ?? %d", vlanId)
		return nil
	}
	return nil
}

func readVlanFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vlanIdStr := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vlan state %s", vlanIdStr)
	data, err := client.FindResource(service.Vlan.Type(), vlanIdStr)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vlan state %s", vlanIdStr)
		d.SetId("")
		return nil
	}
	d.Set("aliasname", data["aliasname"])
	d.Set("dynamicrouting", data["dynamicrouting"])
	setToInt("vlanid", d, data["id"])
	d.Set("ipv6dynamicrouting", data["ipv6dynamicrouting"])
	setToInt("mtu", d, data["mtu"])
	d.Set("sharing", data["sharing"])

	return nil

}

func updateVlanFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vlanId := d.Get("vlanid").(int)
	vlan := network.Vlan{
		Id: d.Get("vlanid").(int),
	}
	hasChange := false
	if d.HasChange("aliasname") {
		log.Printf("[DEBUG]  citrixadc-provider: Aliasname has changed for vlan %d, starting update", vlanId)
		vlan.Aliasname = d.Get("aliasname").(string)
		hasChange = true
	}
	if d.HasChange("dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamicrouting has changed for vlan %d, starting update", vlanId)
		vlan.Dynamicrouting = d.Get("dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("ipv6dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv6dynamicrouting has changed for vlan %d, starting update", vlanId)
		vlan.Ipv6dynamicrouting = d.Get("ipv6dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("mtu") {
		log.Printf("[DEBUG]  citrixadc-provider: Mtu has changed for vlan %d, starting update", vlanId)
		vlan.Mtu = d.Get("mtu").(int)
		hasChange = true
	}
	if d.HasChange("sharing") {
		log.Printf("[DEBUG]  citrixadc-provider: Sharing has changed for vlan %d, starting update", vlanId)
		vlan.Sharing = d.Get("sharing").(string)
		hasChange = true
	}

	vlanIdStr := strconv.Itoa(vlanId)
	if hasChange {
		_, err := client.UpdateResource(service.Vlan.Type(), vlanIdStr, &vlan)
		if err != nil {
			return fmt.Errorf("Error updating vlan %d", vlanId)
		}
	}
	return readVlanFunc(d, meta)
}

func deleteVlanFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vlanName := d.Id()
	err := client.DeleteResource(service.Vlan.Type(), vlanName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
