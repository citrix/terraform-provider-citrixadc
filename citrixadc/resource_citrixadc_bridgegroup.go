package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcBridgegroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBridgegroupFunc,
		Read:          readBridgegroupFunc,
		Update:        updateBridgegroupFunc,
		Delete:        deleteBridgegroupFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bridgegroup_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"dynamicrouting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6dynamicrouting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createBridgegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createBridgegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgegroup_Id := d.Get("bridgegroup_id").(int)
	bridgegroup := network.Bridgegroup{
		Dynamicrouting:     d.Get("dynamicrouting").(string),
		Id:                 d.Get("bridgegroup_id").(int),
		Ipv6dynamicrouting: d.Get("ipv6dynamicrouting").(string),
	}
	bridgegroup_IdStr := strconv.Itoa(bridgegroup_Id)
	_, err := client.AddResource(service.Bridgegroup.Type(), bridgegroup_IdStr, &bridgegroup)
	if err != nil {
		return err
	}

	d.SetId(bridgegroup_IdStr)

	err = readBridgegroupFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this bridgegroup but we can't read it ?? %s", bridgegroup_IdStr)
		return nil
	}
	return nil
}

func readBridgegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readBridgegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgegroup_IdStr := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading bridgegroup state %s", bridgegroup_IdStr)
	data, err := client.FindResource(service.Bridgegroup.Type(), bridgegroup_IdStr)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing bridgegroup state %s", bridgegroup_IdStr)
		d.SetId("")
		return nil
	}
	d.Set("dynamicrouting", data["dynamicrouting"])
	d.Set("bridgegroup_id", data["id"])
	d.Set("ipv6dynamicrouting", data["ipv6dynamicrouting"])

	return nil

}

func updateBridgegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateBridgegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgegroup_Id := d.Get("bridgegroup_id").(int)

	bridgegroup := network.Bridgegroup{
		Id: d.Get("bridgegroup_id").(int),
	}
	hasChange := false
	if d.HasChange("dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamicrouting has changed for bridgegroup %d, starting update", bridgegroup_Id)
		bridgegroup.Dynamicrouting = d.Get("dynamicrouting").(string)
		hasChange = true
	}
	if d.HasChange("ipv6dynamicrouting") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv6dynamicrouting has changed for bridgegroup %d, starting update", bridgegroup_Id)
		bridgegroup.Ipv6dynamicrouting = d.Get("ipv6dynamicrouting").(string)
		hasChange = true
	}
	bridgegroup_IdStr := strconv.Itoa(bridgegroup_Id)
	if hasChange {
		_, err := client.UpdateResource(service.Bridgegroup.Type(), bridgegroup_IdStr, &bridgegroup)
		if err != nil {
			return fmt.Errorf("Error updating bridgegroup %s", bridgegroup_IdStr)
		}
	}
	return readBridgegroupFunc(d, meta)
}

func deleteBridgegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBridgegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgegroupName := d.Id()
	err := client.DeleteResource(service.Bridgegroup.Type(), bridgegroupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
