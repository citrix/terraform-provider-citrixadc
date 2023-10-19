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

func resourceCitrixAdcVpnvserver_authenticationldappolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnvserver_authenticationldappolicy_bindingFunc,
		Read:          readVpnvserver_authenticationldappolicy_bindingFunc,
		Delete:        deleteVpnvserver_authenticationldappolicy_bindingFunc,
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
				Optional: true,
				Computed: true,
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

func createVpnvserver_authenticationldappolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_authenticationldappolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	policy := d.Get("policy")
	bindingId := fmt.Sprintf("%s,%s", name, policy)
	vpnvserver_authenticationldappolicy_binding := vpn.Vpnvserverauthenticationldappolicybinding{
		Bindpoint:              d.Get("bindpoint").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Groupextraction:        d.Get("groupextraction").(bool),
		Name:                   d.Get("name").(string),
		Policy:                 d.Get("policy").(string),
		Priority:               d.Get("priority").(int),
		Secondary:              d.Get("secondary").(bool),
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_authenticationldappolicy_binding.Type(), &vpnvserver_authenticationldappolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readVpnvserver_authenticationldappolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnvserver_authenticationldappolicy_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readVpnvserver_authenticationldappolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_authenticationldappolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	policy := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_authenticationldappolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_authenticationldappolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_authenticationldappolicy_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_authenticationldappolicy_binding state %s", bindingId)
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
	setToInt("priority", d, data["priority"])
	d.Set("secondary", data["secondary"])

	return nil

}

func deleteVpnvserver_authenticationldappolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_authenticationldappolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	policy := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policy:%s", policy))
	if val, ok := d.GetOk("secondary"); ok {
		args = append(args, fmt.Sprintf("secondary:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("groupextraction"); ok {
		args = append(args, fmt.Sprintf("groupextraction:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("bindpoint"); ok {
		args = append(args, fmt.Sprintf("bindpoint:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Vpnvserver_authenticationldappolicy_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
