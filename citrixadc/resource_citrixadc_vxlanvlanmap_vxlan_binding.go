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

func resourceCitrixAdcVxlanvlanmap_vxlan_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVxlanvlanmap_vxlan_bindingFunc,
		Read:          readVxlanvlanmap_vxlan_bindingFunc,
		Delete:        deleteVxlanvlanmap_vxlan_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"vxlan": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"vlan": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVxlanvlanmap_vxlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVxlanvlanmap_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	vxlan := strconv.Itoa(d.Get("vxlan").(int))
	bindingId := fmt.Sprintf("%s,%s", name, vxlan)
	vxlanvlanmap_vxlan_binding := network.Vxlanvlanmapvxlanbinding{
		Name:  d.Get("name").(string),
		Vlan:  toStringList(d.Get("vlan").([]interface{})),
		Vxlan: d.Get("vxlan").(int),
	}

	err := client.UpdateUnnamedResource("vxlanvlanmap_vxlan_binding", &vxlanvlanmap_vxlan_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readVxlanvlanmap_vxlan_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vxlanvlanmap_vxlan_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readVxlanvlanmap_vxlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVxlanvlanmap_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	vxlan := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vxlanvlanmap_vxlan_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vxlanvlanmap_vxlan_binding",
		ResourceName:             name,
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
		log.Printf("[WARN] citrixadc-provider: Clearing vxlanvlanmap_vxlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["vxlan"].(string) == vxlan {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vxlanvlanmap_vxlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("vlan", data["vlan"])
	d.Set("vxlan", data["vxlan"])

	return nil

}

func deleteVxlanvlanmap_vxlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVxlanvlanmap_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	vxlan := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("vxlan:%s", vxlan))

	err := client.DeleteResourceWithArgs("vxlanvlanmap_vxlan_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
