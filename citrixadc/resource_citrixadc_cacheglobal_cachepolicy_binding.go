package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cache"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCacheglobal_cachepolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCacheglobal_cachepolicy_bindingFunc,
		Read:          readCacheglobal_cachepolicy_bindingFunc,
		Delete:        deleteCacheglobal_cachepolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"policy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"invoke": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"precededefrules": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createCacheglobal_cachepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCacheglobal_cachepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policy := d.Get("policy").(string)
	cacheglobal_cachepolicy_binding := cache.Cacheglobalcachepolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Policy:                 d.Get("policy").(string),
		Precededefrules:        d.Get("precededefrules").(string),
		Priority:               d.Get("priority").(int),
		Type:                   d.Get("type").(string),
	}

	err := client.UpdateUnnamedResource(service.Cacheglobal_cachepolicy_binding.Type(),&cacheglobal_cachepolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(policy)

	err = readCacheglobal_cachepolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cacheglobal_cachepolicy_binding but we can't read it ?? %s", policy)
		return nil
	}
	return nil
}

func readCacheglobal_cachepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCacheglobal_cachepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policy := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading cacheglobal_cachepolicy_binding state %s", policy)

	findParams := service.FindParams{
		ResourceType:             "cacheglobal_cachepolicy_binding",
		ArgsMap:                  map[string]string{"type": d.Get("type").(string)},
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
		log.Printf("[WARN] citrixadc-provider: Clearing cacheglobal_cachepolicy_binding state %s", policy)
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
		log.Printf("[WARN] citrixadc-provider: Clearing cacheglobal_cachepolicy_binding state %s", policy)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("invoke", data["invoke"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("policy", data["policy"])
	d.Set("precededefrules", data["precededefrules"])
	setToInt("priority", d, data["priority"])
	d.Set("type", data["type"])

	return nil

}

func deleteCacheglobal_cachepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCacheglobal_cachepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policy := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policy:%s", policy))
	args = append(args, fmt.Sprintf("priority:%v", d.Get("priority").(int)))
	args = append(args, fmt.Sprintf("type:%s", d.Get("type").(string)))

	err := client.DeleteResourceWithArgs(service.Cacheglobal_cachepolicy_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
