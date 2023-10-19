package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
)

func resourceCitrixAdcRnatglobal_auditsyslogpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRnatglobal_auditsyslogpolicy_bindingFunc,
		Read:          readRnatglobal_auditsyslogpolicy_bindingFunc,
		Delete:        deleteRnatglobal_auditsyslogpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"policy": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"all": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createRnatglobal_auditsyslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRnatglobal_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policy := d.Get("policy").(string)
	rnatglobal_auditsyslogpolicy_binding := network.Rnatglobalauditsyslogpolicybinding{
		All:      d.Get("all").(bool),
		Policy:   d.Get("policy").(string),
		Priority: d.Get("priority").(int),
	}

	err := client.UpdateUnnamedResource(service.Rnatglobal_auditsyslogpolicy_binding.Type(), &rnatglobal_auditsyslogpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(policy)

	err = readRnatglobal_auditsyslogpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rnatglobal_auditsyslogpolicy_binding but we can't read it ?? %s", policy)
		return nil
	}
	return nil
}

func readRnatglobal_auditsyslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRnatglobal_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policy := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading rnatglobal_auditsyslogpolicy_binding state %s", policy)

	findParams := service.FindParams{
		ResourceType:             "rnatglobal_auditsyslogpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing rnatglobal_auditsyslogpolicy_binding state %s", policy)
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
		log.Printf("[WARN] citrixadc-provider: Clearing rnatglobal_auditsyslogpolicy_binding state %s", policy)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("all", data["all"])
	d.Set("policy", data["policy"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteRnatglobal_auditsyslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRnatglobal_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policy := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policy:%s", policy))
	if val, ok := d.GetOk("all"); ok {
		args = append(args, fmt.Sprintf("all:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Rnatglobal_auditsyslogpolicy_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
