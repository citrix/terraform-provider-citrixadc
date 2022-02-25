package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcVpnvserver_vpneula_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnvserver_vpneula_bindingFunc,
		Read:          readVpnvserver_vpneula_bindingFunc,
		Delete:        deleteVpnvserver_vpneula_bindingFunc,
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
			"eula": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createVpnvserver_vpneula_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_vpneula_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	eula := d.Get("eula")
	bindingId := fmt.Sprintf("%s,%s", name, eula)
	vpnvserver_vpneula_binding := vpn.Vpnvservervpneulabinding{
		Eula: d.Get("eula").(string),
		Name: d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_vpneula_binding.Type(), &vpnvserver_vpneula_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readVpnvserver_vpneula_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnvserver_vpneula_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readVpnvserver_vpneula_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_vpneula_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	eula := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_vpneula_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_vpneula_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_vpneula_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["eula"].(string) == eula {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_vpneula_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("eula", data["eula"])
	d.Set("name", data["name"])

	return nil

}

func deleteVpnvserver_vpneula_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_vpneula_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	eula := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("eula:%s", eula))

	err := client.DeleteResourceWithArgs(service.Vpnvserver_vpneula_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
