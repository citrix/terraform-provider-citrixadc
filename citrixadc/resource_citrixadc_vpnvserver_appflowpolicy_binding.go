package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcVpnvserver_appflowpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnvserver_appflowpolicy_bindingFunc,
		Read:          readVpnvserver_appflowpolicy_bindingFunc,
		Delete:        deleteVpnvserver_appflowpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"policy": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"bindpoint": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"groupextraction": {
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
			"secondary": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnvserver_appflowpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_appflowpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	policy := d.Get("policy")
	bindpoint := d.Get("bindpoint")
	bindingId := fmt.Sprintf("%s,%s,%s", name, policy, bindpoint)
	vpnvserver_appflowpolicy_binding := vpn.Vpnvserverappflowpolicybinding{
		Bindpoint:              d.Get("bindpoint").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Groupextraction:        d.Get("groupextraction").(bool),
		Name:                   d.Get("name").(string),
		Policy:                 d.Get("policy").(string),
		Priority:               d.Get("priority").(int),
		Secondary:              d.Get("secondary").(bool),
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_appflowpolicy_binding.Type(), &vpnvserver_appflowpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readVpnvserver_appflowpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnvserver_appflowpolicy_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readVpnvserver_appflowpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_appflowpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()

	// To make the resource backward compatible, in the prev state file user will have ID with 2 values, but we have updated Id. So here we are changing the code to make it backward compatible
	// here we are checking for id, if it has 2 elements then we are appending the 3rd attribute to the old Id.
	oldIdSlice := strings.Split(bindingId, ",")

	if len(oldIdSlice) == 2 {
		if _, ok := d.GetOk("bindpoint"); ok {
			bindingId = bindingId + "," + d.Get("bindpoint").(string)
		} else {
			return fmt.Errorf("bindpoint should be given, as it is the part of Id. The id of the vpnvserver_appflowpolicy_binding is the concatenation of the `name`, `policy` and `bindpoint` attributes separated by a comma")
		}
		d.SetId(bindingId)
	}

	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	policy := idSlice[1]
	bindpoint := idSlice[2]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_appflowpolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_appflowpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_appflowpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["policy"].(string) == policy && v["bindpoint"].(string) == bindpoint {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_appflowpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bindpoint", data["bindpoint"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("groupextraction", data["groupextraction"])
	d.Set("name", data["name"])
	d.Set("policy", data["policy"])
	d.Set("priority", data["priority"])
	d.Set("secondary", data["secondary"])

	return nil

}

func deleteVpnvserver_appflowpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_appflowpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	policy := idSlice[1]
	bindpoint := idSlice[2]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policy:%s", policy))
	args = append(args, fmt.Sprintf("bindpoint:%s", bindpoint))
	if val, ok := d.GetOk("secondary"); ok {
		args = append(args, fmt.Sprintf("secondary:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("groupextraction"); ok {
		args = append(args, fmt.Sprintf("groupextraction:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Vpnvserver_appflowpolicy_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
