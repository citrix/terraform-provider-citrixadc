package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcVpnvserver_vpnnexthopserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnvserver_vpnnexthopserver_bindingFunc,
		Read:          readVpnvserver_vpnnexthopserver_bindingFunc,
		Delete:        deleteVpnvserver_vpnnexthopserver_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"nexthopserver": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createVpnvserver_vpnnexthopserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_vpnnexthopserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	nexthopserver := d.Get("nexthopserver")
	bindingId := fmt.Sprintf("%s,%s", name, nexthopserver)
	vpnvserver_vpnnexthopserver_binding := vpn.Vpnvservervpnnexthopserverbinding{
		Name:          d.Get("name").(string),
		Nexthopserver: d.Get("nexthopserver").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_vpnnexthopserver_binding.Type(), &vpnvserver_vpnnexthopserver_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readVpnvserver_vpnnexthopserver_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnvserver_vpnnexthopserver_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readVpnvserver_vpnnexthopserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_vpnnexthopserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	nexthopserver := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_vpnnexthopserver_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_vpnnexthopserver_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_vpnnexthopserver_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_vpnnexthopserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("nexthopserver", data["nexthopserver"])

	return nil

}

func deleteVpnvserver_vpnnexthopserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_vpnnexthopserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	nexthopserver := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("nexthopserver:%s", nexthopserver))

	err := client.DeleteResourceWithArgs(service.Vpnvserver_vpnnexthopserver_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
