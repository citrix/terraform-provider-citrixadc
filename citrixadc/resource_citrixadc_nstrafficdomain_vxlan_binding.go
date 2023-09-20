package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
	"strings"
)

func resourceCitrixAdcNstrafficdomain_vxlan_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNstrafficdomain_vxlan_bindingFunc,
		Read:          readNstrafficdomain_vxlan_bindingFunc,
		Delete:        deleteNstrafficdomain_vxlan_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"td": {
				Type:     schema.TypeInt,
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
		},
	}
}

func createNstrafficdomain_vxlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstrafficdomain_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	td := strconv.Itoa(d.Get("td").(int))
	vxlan := strconv.Itoa(d.Get("vxlan").(int))
	bindingId := fmt.Sprintf("%s,%s", td, vxlan)
	nstrafficdomain_vxlan_binding := ns.Nstrafficdomainvxlanbinding{
		Td:    d.Get("td").(int),
		Vxlan: d.Get("vxlan").(int),
	}

	err := client.UpdateUnnamedResource(service.Nstrafficdomain_vxlan_binding.Type(), &nstrafficdomain_vxlan_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readNstrafficdomain_vxlan_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nstrafficdomain_vxlan_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readNstrafficdomain_vxlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstrafficdomain_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	td := idSlice[0]
	vxlan := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading nstrafficdomain_vxlan_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "nstrafficdomain_vxlan_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing nstrafficdomain_vxlan_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing nstrafficdomain_vxlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("td", data["td"])
	d.Set("vxlan", data["vxlan"])

	return nil

}

func deleteNstrafficdomain_vxlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstrafficdomain_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	vxlan := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("vxlan:%s", vxlan))

	err := client.DeleteResourceWithArgs(service.Nstrafficdomain_vxlan_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
