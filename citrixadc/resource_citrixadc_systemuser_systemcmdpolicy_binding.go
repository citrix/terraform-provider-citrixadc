package citrixadc

import (
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSystemuser_systemcmdpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystemuser_systemcmdpolicy_bindingFunc,
		Read:          readSystemuser_systemcmdpolicy_bindingFunc,
		Delete:        deleteSystemuser_systemcmdpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"policyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSystemuser_systemcmdpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemuser_systemcmdpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	username := d.Get("username")
	policyname := d.Get("policyname")
	bindingId := fmt.Sprintf("%s,%s", username, policyname)

	systemuser_systemcmdpolicy_binding := system.Systemusersystemcmdpolicybinding{
		Policyname: d.Get("policyname").(string),
		Priority:   d.Get("priority").(int),
		Username:   d.Get("username").(string),
	}

	_, err := client.AddResource(service.Systemuser_systemcmdpolicy_binding.Type(), bindingId, &systemuser_systemcmdpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSystemuser_systemcmdpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this systemuser_systemcmdpolicy_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSystemuser_systemcmdpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemuser_systemcmdpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	username := idSlice[0]
	policyname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading systemuser_systemcmdpolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "systemuser_systemcmdpolicy_binding",
		ResourceName:             username,
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
		log.Printf("[WARN] citrixadc-provider: Clearing systemuser_systemcmdpolicy_binding state %s", bindingId)
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
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing systemuser_systemcmdpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])
	d.Set("username", data["username"])

	return nil

}

func deleteSystemuser_systemcmdpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemuser_systemcmdpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	policyname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", url.QueryEscape(policyname)))

	err := client.DeleteResourceWithArgs(service.Systemuser_systemcmdpolicy_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
