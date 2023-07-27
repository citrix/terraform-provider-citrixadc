package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcVxlan() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVxlanFunc,
		Read:          readVxlanFunc,
		Update:        updateVxlanFunc,
		Delete:        deleteVxlanFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vxlanid": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"dynamicrouting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"innervlantagging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6dynamicrouting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVxlanFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVxlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanId := d.Get("vxlanid").(int)
	vxlan := network.Vxlan{
		Dynamicrouting:     d.Get("dynamicrouting").(string),
		Id:                 d.Get("vxlanid").(int),
		Innervlantagging:   d.Get("innervlantagging").(string),
		Ipv6dynamicrouting: d.Get("ipv6dynamicrouting").(string),
		Port:               d.Get("port").(int),
		Protocol:           d.Get("protocol").(string),
		Type:               d.Get("type").(string),
		Vlan:               d.Get("vlan").(int),
	}
	vxlanIdStr := strconv.Itoa(vxlanId)
	_, err := client.AddResource(service.Vxlan.Type(), vxlanIdStr, &vxlan)
	if err != nil {
		return err
	}

	d.SetId(vxlanIdStr)

	err = readVxlanFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vxlan but we can't read it ?? %d", vxlanId)
		return nil
	}
	return nil
}

func readVxlanFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVxlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanIdStr := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vxlan state %s", vxlanIdStr)
	data, err := client.FindResource(service.Vxlan.Type(), vxlanIdStr)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vxlan state %s", vxlanIdStr)
		d.SetId("")
		return nil
	}
	d.Set("dynamicrouting", data["dynamicrouting"])
	d.Set("vxlanid", data["id"])
	d.Set("innervlantagging", data["innervlantagging"])
	d.Set("ipv6dynamicrouting", data["ipv6dynamicrouting"])
	d.Set("port", data["port"])
	d.Set("protocol", data["protocol"])
	d.Set("type", data["type"])
	d.Set("vlan", data["vlan"])

	return nil

}

func updateVxlanFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVxlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanId := d.Get("vxlanid").(int)

	vxlan := network.Vxlan{
		Id: d.Get("vxlanid").(int),
	}
	hasChange := false
	if d.HasChange("dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamicrouting has changed for vxlan %d, starting update", vxlanId)
		vxlan.Dynamicrouting = d.Get("dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("innervlantagging") {
		log.Printf("[DEBUG]  citrixadc-provider: Innervlantagging has changed for vxlan %d, starting update", vxlanId)
		vxlan.Innervlantagging = d.Get("innervlantagging").(string)
		hasChange = true
	}
	if d.HasChange("ipv6dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv6dynamicrouting has changed for vxlan %d, starting update", vxlanId)
		vxlan.Ipv6dynamicrouting = d.Get("ipv6dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for vxlan %d, starting update", vxlanId)
		vxlan.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("protocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocol has changed for vxlan %d, starting update", vxlanId)
		vxlan.Protocol = d.Get("protocol").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for vxlan %d, starting update", vxlanId)
		vxlan.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for vxlan %d, starting update", vxlanId)
		vxlan.Vlan = d.Get("vlan").(int)
		hasChange = true
	}
	vxlanIdStr := strconv.Itoa(vxlanId)
	if hasChange {
		_, err := client.UpdateResource(service.Vxlan.Type(), vxlanIdStr, &vxlan)
		if err != nil {
			return fmt.Errorf("Error updating vxlan %d", vxlanId)
		}
	}
	return readVxlanFunc(d, meta)
}

func deleteVxlanFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVxlanFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanName := d.Id()
	err := client.DeleteResource(service.Vxlan.Type(), vxlanName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
