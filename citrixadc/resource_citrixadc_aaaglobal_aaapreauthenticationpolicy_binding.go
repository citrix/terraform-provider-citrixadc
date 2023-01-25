package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAaaglobal_aaapreauthenticationpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaaglobal_aaapreauthenticationpolicy_bindingFunc,
		Read:          readAaaglobal_aaapreauthenticationpolicy_bindingFunc,
		Delete:        deleteAaaglobal_aaapreauthenticationpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"builtin": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"policy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAaaglobal_aaapreauthenticationpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaaglobal_aaapreauthenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policy := d.Get("policy").(string)
	aaaglobal_aaapreauthenticationpolicy_binding := aaa.Aaaglobalaaapreauthenticationpolicybinding{
		Builtin:  toStringList(d.Get("builtin").([]interface{})),
		Policy:   d.Get("policy").(string),
		Priority: d.Get("priority").(int),
	}

	err := client.UpdateUnnamedResource(service.Aaaglobal_aaapreauthenticationpolicy_binding.Type(), &aaaglobal_aaapreauthenticationpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(policy)

	err = readAaaglobal_aaapreauthenticationpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaaglobal_aaapreauthenticationpolicy_binding but we can't read it ?? %s", policy)
		return nil
	}
	return nil
}

func readAaaglobal_aaapreauthenticationpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaaglobal_aaapreauthenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policy := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading aaaglobal_aaapreauthenticationpolicy_binding state %s", policy)

	findParams := service.FindParams{
		ResourceType:             "aaaglobal_aaapreauthenticationpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaaglobal_aaapreauthenticationpolicy_binding state %s", policy)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["policy"].(string) == policy {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing aaaglobal_aaapreauthenticationpolicy_binding state %s", policy)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("builtin", data["builtin"])
	d.Set("policy", data["policy"])
	d.Set("priority", data["priority"])

	return nil

}

func deleteAaaglobal_aaapreauthenticationpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaaglobal_aaapreauthenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policy := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policy:%s", policy))

	err := client.DeleteResourceWithArgs(service.Aaaglobal_aaapreauthenticationpolicy_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
