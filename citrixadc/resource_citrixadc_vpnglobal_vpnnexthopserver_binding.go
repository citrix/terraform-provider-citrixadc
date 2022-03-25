package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnglobal_vpnnexthopserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnglobal_vpnnexthopserver_bindingFunc,
		Read:          readVpnglobal_vpnnexthopserver_bindingFunc,
		Delete:        deleteVpnglobal_vpnnexthopserver_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"nexthopserver": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnglobal_vpnnexthopserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_vpnnexthopserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	nexthopserver := d.Get("nexthopserver").(string)
	vpnglobal_vpnnexthopserver_binding := vpn.Vpnglobalvpnnexthopserverbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Nexthopserver:          d.Get("nexthopserver").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_vpnnexthopserver_binding.Type(), &vpnglobal_vpnnexthopserver_binding)
	if err != nil {
		return err
	}

	d.SetId(nexthopserver)

	err = readVpnglobal_vpnnexthopserver_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnglobal_vpnnexthopserver_binding but we can't read it ?? %s", nexthopserver)
		return nil
	}
	return nil
}

func readVpnglobal_vpnnexthopserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_vpnnexthopserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	nexthopserver := d.Id()
	
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_vpnnexthopserver_binding state %s", nexthopserver)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_vpnnexthopserver_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnnexthopserver_binding state %s", nexthopserver)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["nexthopserver"].(string) == nexthopserver {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnnexthopserver_binding state %s", nexthopserver)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("nexthopserver", data["nexthopserver"])

	return nil

}

func deleteVpnglobal_vpnnexthopserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_vpnnexthopserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	nexthopserver := d.Id()
	
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("nexthopserver:%s", nexthopserver))

	err := client.DeleteResourceWithArgs(service.Vpnglobal_vpnnexthopserver_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
