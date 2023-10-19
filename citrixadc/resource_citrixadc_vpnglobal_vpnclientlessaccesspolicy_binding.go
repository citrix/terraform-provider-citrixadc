package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
)

func resourceCitrixAdcVpnglobal_vpnclientlessaccesspolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnglobal_vpnclientlessaccesspolicy_bindingFunc,
		Read:          readVpnglobal_vpnclientlessaccesspolicy_bindingFunc,
		Delete:        deleteVpnglobal_vpnclientlessaccesspolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"policyname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"builtin": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"feature": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"globalbindtype": {
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
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnglobal_vpnclientlessaccesspolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_vpnclientlessaccesspolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	vpnglobal_vpnclientlessaccesspolicy_binding := vpn.Vpnglobalvpnclientlessaccesspolicybinding{
		Builtin:                toStringList(d.Get("builtin").([]interface{})),
		Feature:                d.Get("feature").(string),
		Globalbindtype:         d.Get("globalbindtype").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Groupextraction:        d.Get("groupextraction").(bool),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
		Secondary:              d.Get("secondary").(bool),
		Type:                   d.Get("type").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_vpnclientlessaccesspolicy_binding.Type(), &vpnglobal_vpnclientlessaccesspolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(policyname)

	err = readVpnglobal_vpnclientlessaccesspolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnglobal_vpnclientlessaccesspolicy_binding but we can't read it ?? %s", policyname)
		return nil
	}
	return nil
}

func readVpnglobal_vpnclientlessaccesspolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_vpnclientlessaccesspolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_vpnclientlessaccesspolicy_binding state %s", policyname)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_vpnclientlessaccesspolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnclientlessaccesspolicy_binding state %s", policyname)
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnclientlessaccesspolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("builtin", data["builtin"])
	d.Set("feature", data["feature"])
	// d.Set("globalbindtype", data["globalbindtype"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("groupextraction", data["groupextraction"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])
	d.Set("secondary", data["secondary"])
	// d.Set("type", data["type"])

	return nil

}

func deleteVpnglobal_vpnclientlessaccesspolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_vpnclientlessaccesspolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))
	if val, ok := d.GetOk("secondary"); ok {
		args = append(args, fmt.Sprintf("secondary:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("groupextraction"); ok {
		args = append(args, fmt.Sprintf("groupextraction:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Vpnglobal_vpnclientlessaccesspolicy_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
