package citrixadc

import (
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/filter"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcFilterglobal_filterpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createFilterglobal_filterpolicy_bindingFunc,
		Read:          readFilterglobal_filterpolicy_bindingFunc,
		Delete:        deleteFilterglobal_filterpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createFilterglobal_filterpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createFilterglobal_filterpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	filterglobal_filterpolicy_binding := filter.Filterglobalfilterpolicybinding{
		Policyname: d.Get("policyname").(string),
		Priority:   d.Get("priority").(int),
		State:      d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource(service.Filterglobal_filterpolicy_binding.Type(), &filterglobal_filterpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(policyname)

	err = readFilterglobal_filterpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this filterglobal_filterpolicy_binding but we can't read it ?? %s", policyname)
		return nil
	}
	return nil
}

func readFilterglobal_filterpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readFilterglobal_filterpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading filterglobal_filterpolicy_binding state %s", policyname)

	findParams := service.FindParams{
		ResourceType: "filterglobal_filterpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing filterglobal_filterpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["policyname"].(string) == policyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams policyname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing filterglobal_filterpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("policyname", data["policyname"])
	d.Set("priority", data["priority"])
	d.Set("state", data["state"])

	return nil

}

func deleteFilterglobal_filterpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFilterglobal_filterpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", url.QueryEscape(policyname)))

	err := client.DeleteResourceWithArgs(service.Filterglobal_filterpolicy_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
