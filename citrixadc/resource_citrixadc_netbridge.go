package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNetbridge() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNetbridgeFunc,
		Read:          readNetbridgeFunc,
		Update:        updateNetbridgeFunc,
		Delete:        deleteNetbridgeFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"vxlanvlanmap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNetbridgeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNetbridgeFunc")
	client := meta.(*NetScalerNitroClient).client
	netbridgeName := d.Get("name").(string)
	netbridge := network.Netbridge{
		Name:         d.Get("name").(string),
		Vxlanvlanmap: d.Get("vxlanvlanmap").(string),
	}

	_, err := client.AddResource(service.Netbridge.Type(), netbridgeName, &netbridge)
	if err != nil {
		return err
	}

	d.SetId(netbridgeName)

	err = readNetbridgeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this netbridge but we can't read it ?? %s", netbridgeName)
		return nil
	}
	return nil
}

func readNetbridgeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNetbridgeFunc")
	client := meta.(*NetScalerNitroClient).client
	netbridgeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading netbridge state %s", netbridgeName)
	data, err := client.FindResource(service.Netbridge.Type(), netbridgeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing netbridge state %s", netbridgeName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("vxlanvlanmap", data["vxlanvlanmap"])

	return nil

}

func updateNetbridgeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNetbridgeFunc")
	client := meta.(*NetScalerNitroClient).client
	netbridgeName := d.Get("name").(string)

	netbridge := network.Netbridge{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("vxlanvlanmap") {
		log.Printf("[DEBUG]  citrixadc-provider: Vxlanvlanmap has changed for netbridge %s, starting update", netbridgeName)
		netbridge.Vxlanvlanmap = d.Get("vxlanvlanmap").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Netbridge.Type(), netbridgeName, &netbridge)
		if err != nil {
			return fmt.Errorf("Error updating netbridge %s", netbridgeName)
		}
	}
	return readNetbridgeFunc(d, meta)
}

func deleteNetbridgeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNetbridgeFunc")
	client := meta.(*NetScalerNitroClient).client
	netbridgeName := d.Id()
	err := client.DeleteResource(service.Netbridge.Type(), netbridgeName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
