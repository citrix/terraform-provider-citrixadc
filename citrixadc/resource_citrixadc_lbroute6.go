package citrixadc

import (
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcLbroute6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbroute6Func,
		Read:          readLbroute6Func,
		Delete:        deleteLbroute6Func,
		Schema: map[string]*schema.Schema{
			"gatewayname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLbroute6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbroute6Func")
	client := meta.(*NetScalerNitroClient).client
	var network = d.Get("network").(string)
	lbroute6 := lb.Lbroute6{
		Gatewayname: d.Get("gatewayname").(string),
		Network:     d.Get("network").(string),
		Td:          d.Get("td").(int),
	}

	_, err := client.AddResource(service.Lbroute6.Type(), "", &lbroute6)
	if err != nil {
		return err
	}

	d.SetId(network)

	err = readLbroute6Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbroute6 but we can't read it ?? %s", network)
		return nil
	}
	return nil
}

func readLbroute6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbroute6Func")
	client := meta.(*NetScalerNitroClient).client
	network := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading lbroute6 state %s", network)
	findParams := service.FindParams{
		ResourceType: service.Lbroute6.Type(),
	}

	dataArray, err := client.FindResourceArrayWithParams(findParams)

	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lb route6 state %s", network)
		d.SetId("")
		return nil
	}

	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: lb route6 does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, lbroute6 := range dataArray {
		if lbroute6["network"] == network {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams route6 not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lb route6 state %s", network)
		d.SetId("")
		return nil
	}

	data := dataArray[foundIndex]

	d.Set("gatewayname", data["gatewayname"])
	d.Set("network", data["network"])
	d.Set("td", data["td"])

	return nil

}

func deleteLbroute6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbroute6Func")
	client := meta.(*NetScalerNitroClient).client

	argsMap := make(map[string]string)
	argsMap["network"] = url.QueryEscape(d.Get("network").(string))

	err := client.DeleteResourceWithArgsMap(service.Lbroute6.Type(), "", argsMap)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
