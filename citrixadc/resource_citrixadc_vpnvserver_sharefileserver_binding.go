package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcVpnvserver_sharefileserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnvserver_sharefileserver_bindingFunc,
		Read:          readVpnvserver_sharefileserver_bindingFunc,
		Delete:        deleteVpnvserver_sharefileserver_bindingFunc,
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
			"sharefile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createVpnvserver_sharefileserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_sharefileserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	sharefile := d.Get("sharefile")
	bindingId := fmt.Sprintf("%s,%s", name, sharefile)
	vpnvserver_sharefileserver_binding := vpn.Vpnvserversharefileserverbinding{
		Name:      d.Get("name").(string),
		Sharefile: d.Get("sharefile").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_sharefileserver_binding.Type(), &vpnvserver_sharefileserver_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readVpnvserver_sharefileserver_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnvserver_sharefileserver_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readVpnvserver_sharefileserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_sharefileserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	sharefile := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_sharefileserver_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_sharefileserver_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_sharefileserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["sharefile"].(string) == sharefile {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_sharefileserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("sharefile", data["sharefile"])

	return nil

}

func deleteVpnvserver_sharefileserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_sharefileserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	sharefile := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("sharefile:%s", sharefile))

	err := client.DeleteResourceWithArgs(service.Vpnvserver_sharefileserver_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
