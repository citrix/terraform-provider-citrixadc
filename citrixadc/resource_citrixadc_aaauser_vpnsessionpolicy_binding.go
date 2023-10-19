package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcAaauser_vpnsessionpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaauser_vpnsessionpolicy_bindingFunc,
		Read:          readAaauser_vpnsessionpolicy_bindingFunc,
		Delete:        deleteAaauser_vpnsessionpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAaauser_vpnsessionpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaauser_vpnsessionpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	username := d.Get("username").(string)
	policy := d.Get("policy").(string)
	bindingId := fmt.Sprintf("%s,%s", username, policy)
	aaauser_vpnsessionpolicy_binding := aaa.Aaauservpnsessionpolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Policy:                 d.Get("policy").(string),
		Priority:               d.Get("priority").(int),
		Type:                   d.Get("type").(string),
		Username:               d.Get("username").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaauser_vpnsessionpolicy_binding.Type(), &aaauser_vpnsessionpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAaauser_vpnsessionpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaauser_vpnsessionpolicy_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAaauser_vpnsessionpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaauser_vpnsessionpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	username := idSlice[0]
	policy := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading aaauser_vpnsessionpolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "aaauser_vpnsessionpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaauser_vpnsessionpolicy_binding state %s", bindingId)
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
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams policy not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing aaauser_vpnsessionpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("policy", data["policy"])
	setToInt("priority", d, data["priority"])
	d.Set("type", data["type"])
	d.Set("username", data["username"])

	return nil

}

func deleteAaauser_vpnsessionpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaauser_vpnsessionpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	policy := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policy:%s", policy))
	if v, ok := d.GetOk("type"); ok {
		type_val := v.(string)
		args = append(args, fmt.Sprintf("type:%s", type_val))
	}

	err := client.DeleteResourceWithArgs(service.Aaauser_vpnsessionpolicy_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
