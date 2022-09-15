package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ha"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"strconv"
	"log"
	"strings"
)

func resourceCitrixAdcHanode_routemonitor_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createHanode_routemonitor_bindingFunc,
		Read:          readHanode_routemonitor_bindingFunc,
		Delete:        deleteHanode_routemonitor_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"hanode_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"routemonitor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createHanode_routemonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createHanode_routemonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	id := strconv.Itoa(d.Get("hanode_id").(int))
	routemonitor := d.Get("routemonitor").(string)
	bindingId := fmt.Sprintf("%s,%s", id, routemonitor)
	hanode_routemonitor_binding := ha.Hanoderoutemonitorbinding{
		Id:           d.Get("hanode_id").(int),
		Netmask:      d.Get("netmask").(string),
		Routemonitor: d.Get("routemonitor").(string),
	}

	err := client.UpdateUnnamedResource(service.Hanode_routemonitor_binding.Type(), &hanode_routemonitor_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readHanode_routemonitor_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this hanode_routemonitor_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readHanode_routemonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readHanode_routemonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	id := idSlice[0]
	routemonitor := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading hanode_routemonitor_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "hanode_routemonitor_binding",
		ResourceName:             id,
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
		log.Printf("[WARN] citrixadc-provider: Clearing hanode_routemonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["routemonitor"].(string) == routemonitor {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams routemonitor not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing hanode_routemonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("id", data["id"])
	d.Set("netmask", data["netmask"])
	d.Set("routemonitor", data["routemonitor"])

	return nil

}

func deleteHanode_routemonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteHanode_routemonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	routemonitor := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("routemonitor:%s", routemonitor))
	args = append(args, fmt.Sprintf("netmask:%s", d.Get("netmask").(string)))

	err := client.DeleteResourceWithArgs(service.Hanode_routemonitor_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
