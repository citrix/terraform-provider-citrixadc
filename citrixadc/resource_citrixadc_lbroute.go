package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
)

func resourceCitrixAdcLbroute() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbrouteFunc,
		Read:          readLbrouteFunc,
		Delete:        deleteLbrouteFunc,
		Schema: map[string]*schema.Schema{
			"gatewayname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createLbrouteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbrouteFunc")
	client := meta.(*NetScalerNitroClient).client

	lbrouteName := fmt.Sprintf("%s,%s,%s", d.Get("network").(string), d.Get("netmask").(string), d.Get("gatewayname").(string))

	lbroute := lb.Lbroute{
		Gatewayname: d.Get("gatewayname").(string),
		Netmask:     d.Get("netmask").(string),
		Network:     d.Get("network").(string),
		Td:          d.Get("td").(int),
	}

	_, err := client.AddResource(service.Lbroute.Type(), lbrouteName, &lbroute)
	if err != nil {
		return err
	}

	d.SetId(lbrouteName)

	err = readLbrouteFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbroute but we can't read it ?? %s", lbrouteName)
		return nil
	}
	return nil
}

func readLbrouteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbrouteFunc")
	client := meta.(*NetScalerNitroClient).client
	lbrouteName := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading lbroute state %s", lbrouteName)
	findParams := service.FindParams{
		ResourceType: service.Lbroute.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lb route state %s", lbrouteName)
		d.SetId("")
		return nil
	}
	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: lb route does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, lbroute := range dataArray {
		match := true
		if lbroute["network"] != d.Get("network").(string) {
			match = false
		}
		if lbroute["netmask"] != d.Get("netmask").(string) {
			match = false
		}
		if lbroute["gatewayname"] != d.Get("gatewayname").(string) {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams route not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lb route state %s", lbrouteName)
		d.SetId("")
		return nil
	}

	data := dataArray[foundIndex]

	d.Set("gatewayname", data["gatewayname"])
	d.Set("netmask", data["netmask"])
	d.Set("network", data["network"])
	d.Set("td", data["td"])

	return nil

}

func deleteLbrouteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbrouteFunc")
	client := meta.(*NetScalerNitroClient).client

	argsMap := make(map[string]string)
	// Only the network and netmask properties are required for deletion - not gatewayname
	argsMap["network"] = url.QueryEscape(d.Get("network").(string))
	argsMap["netmask"] = url.QueryEscape(d.Get("netmask").(string))

	err := client.DeleteResourceWithArgsMap(service.Lbroute.Type(), "", argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
