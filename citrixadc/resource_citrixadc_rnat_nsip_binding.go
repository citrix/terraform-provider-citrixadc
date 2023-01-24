package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcRnat_nsip_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRnat_nsip_bindingFunc,
		Read:          readRnat_nsip_bindingFunc,
		Delete:        deleteRnat_nsip_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"natip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createRnat_nsip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRnat_nsip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	natip := d.Get("natip")
	bindingId := fmt.Sprintf("%s,%s", name, natip)
	rnat_nsip_binding := network.Rnatnsipbinding{
		Name:  d.Get("name").(string),
		Natip: d.Get("natip").(string),
	}

	err := client.UpdateUnnamedResource("rnat_nsip_binding", &rnat_nsip_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readRnat_nsip_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rnat_nsip_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readRnat_nsip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRnat_nsip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	natip := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading rnat_nsip_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "rnat_nsip_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing rnat_nsip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["natip"].(string) == natip {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams natip not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing rnat_nsip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("natip", data["natip"])

	return nil

}

func deleteRnat_nsip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRnat_nsip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	natip := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("natip:%s", natip))

	err := client.DeleteResourceWithArgs("rnat_nsip_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
