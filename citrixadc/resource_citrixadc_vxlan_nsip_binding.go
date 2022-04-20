package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
	"strconv"
	"net/url"
)

func resourceCitrixAdcVxlan_nsip_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVxlan_nsip_bindingFunc,
		Read:          readVxlan_nsip_bindingFunc,
		Delete:        deleteVxlan_nsip_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vxlanid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVxlan_nsip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVxlan_nsip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanid := strconv.Itoa(d.Get("vxlanid").(int))
	ipaddress := d.Get("ipaddress").(string)
	bindingId := fmt.Sprintf("%s,%s", vxlanid, ipaddress)
	vxlan_nsip_binding := network.Vxlannsipbinding{
		Id:        d.Get("vxlanid").(int),
		Ipaddress: d.Get("ipaddress").(string),
		Netmask:   d.Get("netmask").(string),
	}

	err := client.UpdateUnnamedResource(service.Vxlan_nsip_binding.Type(), &vxlan_nsip_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readVxlan_nsip_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vxlan_nsip_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readVxlan_nsip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVxlan_nsip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vxlanid := idSlice[0]
	ipaddress := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vxlan_nsip_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vxlan_nsip_binding",
		ResourceName:             vxlanid,
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
		log.Printf("[WARN] citrixadc-provider: Clearing vxlan_nsip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["ipaddress"].(string) == ipaddress {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vxlan_nsip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("vxlanid", data["id"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("netmask", data["netmask"])

	return nil

}

func deleteVxlan_nsip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVxlan_nsip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vxlanid := idSlice[0]
	ipaddress := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ipaddress:%s", url.QueryEscape(ipaddress)))
	if val, ok := d.GetOk("netmask"); ok {
		args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Vxlan_nsip_binding.Type(), vxlanid, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
