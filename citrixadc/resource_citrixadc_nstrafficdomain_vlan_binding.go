package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
	"strconv"
)

func resourceCitrixAdcNstrafficdomain_vlan_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNstrafficdomain_vlan_bindingFunc,
		Read:          readNstrafficdomain_vlan_bindingFunc,
		Delete:        deleteNstrafficdomain_vlan_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNstrafficdomain_vlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstrafficdomain_vlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	td := strconv.Itoa(d.Get("td").(int))
	vlan := strconv.Itoa(d.Get("vlan").(int))
	bindingId := fmt.Sprintf("%s,%s", td, vlan)
	nstrafficdomain_vlan_binding := ns.Nstrafficdomainvlanbinding{
		Td:   d.Get("td").(int),
		Vlan: d.Get("vlan").(int),
	}

	err := client.UpdateUnnamedResource(service.Nstrafficdomain_vlan_binding.Type(), &nstrafficdomain_vlan_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readNstrafficdomain_vlan_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nstrafficdomain_vlan_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readNstrafficdomain_vlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstrafficdomain_vlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	td := idSlice[0]
	vlan := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading nstrafficdomain_vlan_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "nstrafficdomain_vlan_binding",
		ResourceName:             td,
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
		log.Printf("[WARN] citrixadc-provider: Clearing nstrafficdomain_vlan_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing nstrafficdomain_vlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("td", data["td"])
	d.Set("vlan", data["vlan"])

	return nil

}

func deleteNstrafficdomain_vlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstrafficdomain_vlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	vlan := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("vlan:%s", vlan))

	err := client.DeleteResourceWithArgs(service.Nstrafficdomain_vlan_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
