package citrixadc

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcVlan_interface_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVlan_interface_bindingFunc,
		Read:          readVlan_interface_bindingFunc,
		Delete:        deleteVlan_interface_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vlanid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ifnum": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ownergroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tagged": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVlan_interface_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVlan_interface_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vlanid := strconv.Itoa(d.Get("vlanid").(int))
	ifnum := d.Get("ifnum").(string)
	bindingId := fmt.Sprintf("%s,%s", vlanid, ifnum)
	vlan_interface_binding := network.Vlaninterfacebinding{
		Id:         d.Get("vlanid").(int),
		Ifnum:      d.Get("ifnum").(string),
		Ownergroup: d.Get("ownergroup").(string),
		Tagged:     d.Get("tagged").(bool),
	}

	err := client.UpdateUnnamedResource(service.Vlan_interface_binding.Type(), &vlan_interface_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readVlan_interface_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vlan_interface_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readVlan_interface_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVlan_interface_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vlan_interface_bindingName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vlan_interface_binding state %s", vlan_interface_bindingName)
	bindingId := d.Id()

	idSlice := strings.SplitN(bindingId, ",", 2)

	vlanid := idSlice[0]
	ifnum := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vlan_interface_bindingName state %s", bindingId)
	findParams := service.FindParams{
		ResourceType:             "vlan_interface_binding",
		ResourceName:             vlanid,
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
		log.Printf("[WARN] citrixadc-provider: Clearing vlan_interface_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right monitor name
	foundIndex := -1
	for i, v := range dataArr {
		if v["ifnum"].(string) == ifnum {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vlan_interface_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("vlanid", data["id"])
	d.Set("ifnum", data["ifnum"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("tagged", data["tagged"])

	return nil

}

func deleteVlan_interface_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVlan_interface_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	vlanid := idSlice[0]
	ifnum := idSlice[1]
	args := make([]string, 1)
	args[0] = fmt.Sprintf("ifnum:%s", url.QueryEscape(ifnum))
	err := client.DeleteResourceWithArgs(service.Vlan_interface_binding.Type(), vlanid, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
