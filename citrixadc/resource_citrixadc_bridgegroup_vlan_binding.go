package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
	"strings"
)

func resourceCitrixAdcBridgegroup_vlan_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBridgegroup_vlan_bindingFunc,
		Read:          readBridgegroup_vlan_bindingFunc,
		Delete:        deleteBridgegroup_vlan_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bridgegroup_id": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createBridgegroup_vlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createBridgegroup_vlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgegroup_id := strconv.Itoa(d.Get("bridgegroup_id").(int))
	vlan := strconv.Itoa(d.Get("vlan").(int))
	bindingId := fmt.Sprintf("%s,%s", bridgegroup_id, vlan)
	bridgegroup_vlan_binding := network.Bridgegroupvlanbinding{
		Id:   d.Get("bridgegroup_id").(int),
		Vlan: d.Get("vlan").(int),
	}

	err := client.UpdateUnnamedResource(service.Bridgegroup_vlan_binding.Type(), &bridgegroup_vlan_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readBridgegroup_vlan_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this bridgegroup_vlan_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readBridgegroup_vlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readBridgegroup_vlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	bridgegroup_id := idSlice[0]
	vlan := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading bridgegroup_vlan_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "bridgegroup_vlan_binding",
		ResourceName:             bridgegroup_id,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing bridgegroup_vlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["vlan"].(string) == vlan {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing bridgegroup_vlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bridgegroup_id", data["id"])
	d.Set("vlan", data["vlan"])

	return nil

}

func deleteBridgegroup_vlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBridgegroup_vlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	bridgegroup_id := idSlice[0]
	vlan := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("vlan:%s", vlan))

	err := client.DeleteResourceWithArgs(service.Bridgegroup_vlan_binding.Type(), bridgegroup_id, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
