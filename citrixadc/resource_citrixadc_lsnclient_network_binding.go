package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcLsnclient_network_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnclient_network_bindingFunc,
		Read:          readLsnclient_network_bindingFunc,
		Delete:        deleteLsnclient_network_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"clientname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnclient_network_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnclient_network_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	clientname := d.Get("clientname")
	network := d.Get("network")
	bindingId := fmt.Sprintf("%s,%s", clientname, network)
	lsnclient_network_binding := lsn.Lsnclientnetworkbinding{
		Clientname: d.Get("clientname").(string),
		Netmask:    d.Get("netmask").(string),
		Network:    d.Get("network").(string),
		Td:         d.Get("td").(int),
	}

	err := client.UpdateUnnamedResource("lsnclient_network_binding", &lsnclient_network_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLsnclient_network_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnclient_network_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLsnclient_network_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnclient_network_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	clientname := idSlice[0]
	network := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsnclient_network_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsnclient_network_binding",
		ResourceName:             clientname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsnclient_network_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["network"].(string) == network {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams network not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsnclient_network_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("clientname", data["clientname"])
	d.Set("netmask", data["netmask"])
	d.Set("network", data["network"])
	d.Set("td", data["td"])

	return nil

}

func deleteLsnclient_network_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnclient_network_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	network := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("network:%s", network))
	if v, ok := d.GetOk("netmask"); ok {
		args = append(args, fmt.Sprintf("netmask:%s", v.(string)))
	}
	if v, ok := d.GetOk("td"); ok {
		args = append(args, fmt.Sprintf("td:%v", v.(int)))
	}
	err := client.DeleteResourceWithArgs("lsnclient_network_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
