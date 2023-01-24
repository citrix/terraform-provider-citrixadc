package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/tunnel"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcTunnelglobal_tunneltrafficpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createTunnelglobal_tunneltrafficpolicy_bindingFunc,
		Read:          readTunnelglobal_tunneltrafficpolicy_bindingFunc,
		Delete:        deleteTunnelglobal_tunneltrafficpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"feature": &schema.Schema{
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

func createTunnelglobal_tunneltrafficpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createTunnelglobal_tunneltrafficpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	tunnelglobal_tunneltrafficpolicy_binding := tunnel.Tunnelglobaltunneltrafficpolicybinding{
		Feature:                d.Get("feature").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
		State:                  d.Get("state").(string),
		Type:                   d.Get("type").(string),
	}

	err := client.UpdateUnnamedResource(service.Tunnelglobal_tunneltrafficpolicy_binding.Type(), &tunnelglobal_tunneltrafficpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(policyname)

	err = readTunnelglobal_tunneltrafficpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this tunnelglobal_tunneltrafficpolicy_binding but we can't read it ?? %s", policyname)
		return nil
	}
	return nil
}

func readTunnelglobal_tunneltrafficpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readTunnelglobal_tunneltrafficpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading tunnelglobal_tunneltrafficpolicy_binding state %s", policyname)
	
	findParams := service.FindParams{
		ResourceType:             "tunnelglobal_tunneltrafficpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing tunnelglobal_tunneltrafficpolicy_binding state %s", policyname)
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
		log.Printf("[WARN] citrixadc-provider: Clearing tunnelglobal_tunneltrafficpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("feature", data["feature"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("policyname", data["policyname"])
	d.Set("priority", data["priority"])
	d.Set("state", data["state"])
	d.Set("type", data["type"])

	return nil

}

func deleteTunnelglobal_tunneltrafficpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTunnelglobal_tunneltrafficpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))
	if v, ok := d.GetOk("type"); ok {
		args = append(args, fmt.Sprintf("type:%s", v.(string)))
	}
	if v, ok := d.GetOk("priority"); ok {
		args = append(args, fmt.Sprintf("priority:%v", v.(int)))
	}
	err := client.DeleteResourceWithArgs(service.Tunnelglobal_tunneltrafficpolicy_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
