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

func resourceCitrixAdcVpnvserver_appcontroller_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnvserver_appcontroller_bindingFunc,
		Read:          readVpnvserver_appcontroller_bindingFunc,
		Delete:        deleteVpnvserver_appcontroller_bindingFunc,
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
			"appcontroller": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createVpnvserver_appcontroller_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_appcontroller_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	appcontroller := d.Get("appcontroller")
	bindingId := fmt.Sprintf("%s,%s", name, appcontroller)
	vpnvserver_appcontroller_binding := vpn.Vpnvserverappcontrollerbinding{
		Appcontroller: d.Get("appcontroller").(string),
		Name:          d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_appcontroller_binding.Type(), &vpnvserver_appcontroller_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readVpnvserver_appcontroller_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnvserver_appcontroller_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readVpnvserver_appcontroller_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_appcontroller_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	appcontroller := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_appcontroller_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_appcontroller_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_appcontroller_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["appcontroller"].(string) == appcontroller {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_appcontroller_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("appcontroller", data["appcontroller"])
	d.Set("name", data["name"])

	return nil

}

func deleteVpnvserver_appcontroller_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_appcontroller_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	appcontroller := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("appcontroller:%v", url.QueryEscape(appcontroller)))
	err := client.DeleteResourceWithArgs(service.Vpnvserver_appcontroller_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
