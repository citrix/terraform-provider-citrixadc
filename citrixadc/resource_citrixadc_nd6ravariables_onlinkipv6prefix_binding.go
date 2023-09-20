package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func resourceCitrixAdcNd6ravariables_onlinkipv6prefix_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNd6ravariables_onlinkipv6prefix_bindingFunc,
		Read:          readNd6ravariables_onlinkipv6prefix_bindingFunc,
		Delete:        deleteNd6ravariables_onlinkipv6prefix_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ipv6prefix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createNd6ravariables_onlinkipv6prefix_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNd6ravariables_onlinkipv6prefix_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vlan := strconv.Itoa(d.Get("vlan").(int))
	ipv6prefix := d.Get("ipv6prefix")
	bindingId := fmt.Sprintf("%s,%s", vlan, ipv6prefix)
	nd6ravariables_onlinkipv6prefix_binding := network.Nd6ravariablesonlinkipv6prefixbinding{
		Ipv6prefix: d.Get("ipv6prefix").(string),
		Vlan:       d.Get("vlan").(int),
	}

	err := client.UpdateUnnamedResource(service.Nd6ravariables_onlinkipv6prefix_binding.Type(), &nd6ravariables_onlinkipv6prefix_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readNd6ravariables_onlinkipv6prefix_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nd6ravariables_onlinkipv6prefix_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readNd6ravariables_onlinkipv6prefix_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNd6ravariables_onlinkipv6prefix_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vlan := idSlice[0]
	ipv6prefix := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading nd6ravariables_onlinkipv6prefix_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "nd6ravariables_onlinkipv6prefix_binding",
		ResourceName:             vlan,
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
		log.Printf("[WARN] citrixadc-provider: Clearing nd6ravariables_onlinkipv6prefix_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["ipv6prefix"].(string) == ipv6prefix {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams ipv6prefix not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing nd6ravariables_onlinkipv6prefix_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ipv6prefix", data["ipv6prefix"])
	d.Set("vlan", data["vlan"])

	return nil

}

func deleteNd6ravariables_onlinkipv6prefix_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNd6ravariables_onlinkipv6prefix_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ipv6prefix := idSlice[1]

	args := make([]string, 0)
	ipv6prefixEscaped := url.PathEscape(ipv6prefix)
	args = append(args, fmt.Sprintf("ipv6prefix:%s", ipv6prefixEscaped))

	err := client.DeleteResourceWithArgs(service.Nd6ravariables_onlinkipv6prefix_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
