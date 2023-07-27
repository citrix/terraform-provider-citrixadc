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

func resourceCitrixAdcBridgegroup_nsip6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBridgegroup_nsip6_bindingFunc,
		Read:          readBridgegroup_nsip6_bindingFunc,
		Delete:        deleteBridgegroup_nsip6_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bridgegroup_id": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ownergroup": {
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

func createBridgegroup_nsip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createBridgegroup_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bridgegroup_id := strconv.Itoa(d.Get("bridgegroup_id").(int))
	ipaddress := d.Get("ipaddress")
	bindingId := fmt.Sprintf("%s,%s", bridgegroup_id, ipaddress)
	bridgegroup_nsip6_binding := network.Bridgegroupnsip6binding{
		Id:         d.Get("bridgegroup_id").(int),
		Ipaddress:  d.Get("ipaddress").(string),
		Netmask:    d.Get("netmask").(string),
		Ownergroup: d.Get("ownergroup").(string),
		Td:         d.Get("td").(int),
	}

	err := client.UpdateUnnamedResource(service.Bridgegroup_nsip6_binding.Type(), &bridgegroup_nsip6_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readBridgegroup_nsip6_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this bridgegroup_nsip6_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readBridgegroup_nsip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readBridgegroup_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	bridgegroup_id := idSlice[0]
	ipaddress := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading bridgegroup_nsip6_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "bridgegroup_nsip6_binding",
		ResourceName:             bridgegroup_id,
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
		log.Printf("[WARN] citrixadc-provider: Clearing bridgegroup_nsip6_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing bridgegroup_nsip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bridgegroup_id", data["id"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("netmask", data["netmask"])
	d.Set("ownergroup", data["ownergroup"])
	d.Set("td", data["td"])

	return nil

}

func deleteBridgegroup_nsip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBridgegroup_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	bridgegroup_id := idSlice[0]
	ipaddress := idSlice[1]
	ipaddressEscaped := url.PathEscape((ipaddress))
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ipaddress:%s", ipaddressEscaped))
	if val, ok := d.GetOk("netmask"); ok {
		args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("td"); ok {
		args = append(args, fmt.Sprintf("td:%d", val.(int)))
	}
	if val, ok := d.GetOk("ownergroup"); ok {
		args = append(args, fmt.Sprintf("ownergroup:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Bridgegroup_nsip6_binding.Type(), bridgegroup_id, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
