package citrixadc

import (
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLbvserver_appfwpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbvserver_appfwpolicy_bindingFunc,
		Read:          readLbvserver_appfwpolicy_bindingFunc,
		Delete:        deleteLbvserver_appfwpolicy_bindingFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bindpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLbvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserverName := d.Get("name").(string)
	appfwPolicyName := d.Get("policyname").(string)
	bindingId := fmt.Sprintf("%s,%s", lbvserverName, appfwPolicyName)

	lbvserver_appfwpolicy_binding := lb.Lbvserverpolicybinding{
		Bindpoint:              d.Get("bindpoint").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Name:                   lbvserverName,
		Policyname:             appfwPolicyName,
		Priority:               uint32(d.Get("priority").(int)),
	}

	_, err := client.AddResource(service.Lbvserver_appfwpolicy_binding.Type(), lbvserverName, &lbvserver_appfwpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLbvserver_appfwpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbvserver_appfwpolicy_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLbvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: readLbvserver_appfwpolicy_bindingFunc: bindingId: %s", bindingId)
	idSlice := strings.SplitN(bindingId, ",", 2)
	lbvserverName := idSlice[0]
	appfwPolicyName := idSlice[1]
	log.Printf("[DEBUG] citrixadc-provider: Reading lbvserver_appfwpolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             service.Lbvserver_appfwpolicy_binding.Type(),
		ResourceName:             lbvserverName,
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
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_appfwpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if v["policyname"].(string) == appfwPolicyName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_appfwpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough
	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("bindpoint", data["bindpoint"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("invoke", data["invoke"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("name", data["name"])
	d.Set("policyname", data["policyname"])
	d.Set("priority", data["priority"])

	return nil

}

func deleteLbvserver_appfwpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbvserver_appfwpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	lbvserverName := idSlice[0]
	appfwPolicyName := idSlice[1]

	args := make(map[string]string)
	args["policyname"] = appfwPolicyName
	err := client.DeleteResourceWithArgsMap(service.Lbvserver_appfwpolicy_binding.Type(), lbvserverName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
