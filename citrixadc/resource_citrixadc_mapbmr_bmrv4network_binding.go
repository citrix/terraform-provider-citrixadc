package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
	"net/url"
)

func resourceCitrixAdcMapbmr_bmrv4network_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createMapbmr_bmrv4network_bindingFunc,
		Read:          readMapbmr_bmrv4network_bindingFunc,
		Delete:        deleteMapbmr_bmrv4network_bindingFunc,
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
			"network": &schema.Schema{
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

func createMapbmr_bmrv4network_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createMapbmr_bmrv4network_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	networkid := d.Get("network")
	bindingId := fmt.Sprintf("%s,%s", name, networkid)
	mapbmr_bmrv4network_binding := network.Mapbmrbmrv4networkbinding {
		Name:    d.Get("name").(string),
		Netmask: d.Get("netmask").(string),
		Network: d.Get("network").(string),
	}

	err := client.UpdateUnnamedResource("mapbmr_bmrv4network_binding", &mapbmr_bmrv4network_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readMapbmr_bmrv4network_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this mapbmr_bmrv4network_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readMapbmr_bmrv4network_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readMapbmr_bmrv4network_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	networkid := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading mapbmr_bmrv4network_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "mapbmr_bmrv4network_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing mapbmr_bmrv4network_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["network"].(string) == networkid {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing mapbmr_bmrv4network_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("netmask", data["netmask"])
	d.Set("network", data["network"])

	return nil

}

func deleteMapbmr_bmrv4network_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteMapbmr_bmrv4network_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	networkid := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("network:%s", networkid))
	if val, ok := d.GetOk("netmask"); ok {
		args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(val.(string))))
	}


	err := client.DeleteResourceWithArgs("mapbmr_bmrv4network_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
