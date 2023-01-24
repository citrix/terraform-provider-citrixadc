package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcIpset_nsip6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIpset_nsip6_bindingFunc,
		Read:          readIpset_nsip6_bindingFunc,
		Delete:        deleteIpset_nsip6_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createIpset_nsip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIpset_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	ipaddress := d.Get("ipaddress")
	bindingId := fmt.Sprintf("%s,%s", name, ipaddress)
	ipset_nsip6_binding := network.Ipsetnsip6binding{
		Ipaddress: d.Get("ipaddress").(string),
		Name:      d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource(service.Ipset_nsip6_binding.Type(), &ipset_nsip6_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readIpset_nsip6_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ipset_nsip6_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readIpset_nsip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIpset_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ipaddress := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading ipset_nsip6_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "ipset_nsip6_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing ipset_nsip6_binding state %s", bindingId)
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
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams ipaddress not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing ipset_nsip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ipaddress", data["ipaddress"])
	d.Set("name", data["name"])

	return nil

}

func deleteIpset_nsip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIpset_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ipaddress := idSlice[1]
	ipaddressEscaped := url.PathEscape((ipaddress))
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ipaddress:%s", ipaddressEscaped))

	err := client.DeleteResourceWithArgs(service.Ipset_nsip6_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
