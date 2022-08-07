package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"strconv"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcClusternode_routemonitor_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createClusternode_routemonitor_bindingFunc,
		Read:          readClusternode_routemonitor_bindingFunc,
		Delete:        deleteClusternode_routemonitor_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"routemonitor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createClusternode_routemonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternode_routemonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	nodeid := strconv.Itoa(d.Get("nodeid").(int))
	routemonitor := d.Get("routemonitor").(string)
	bindingId := fmt.Sprintf("%s,%s", nodeid, routemonitor)
	clusternode_routemonitor_binding := cluster.Clusternoderoutemonitorbinding{
		Netmask:      d.Get("netmask").(string),
		Nodeid:       d.Get("nodeid").(int),
		Routemonitor: d.Get("routemonitor").(string),
	}

	err := client.UpdateUnnamedResource("clusternode_routemonitor_binding", &clusternode_routemonitor_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readClusternode_routemonitor_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this clusternode_routemonitor_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readClusternode_routemonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternode_routemonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	nodeid := idSlice[0]
	routemonitor := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading clusternode_routemonitor_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "clusternode_routemonitor_binding",
		ResourceName:             nodeid,
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
		log.Printf("[WARN] citrixadc-provider: Clearing clusternode_routemonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["routemonitor"].(string) == routemonitor {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams routemonitor not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing clusternode_routemonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]
	nodeid_int, err := strconv.Atoi(data["nodeid"].(string))
	if err !=nil {
		return err
	}
	d.Set("netmask", data["netmask"])
	d.Set("nodeid", nodeid_int)
	d.Set("routemonitor", data["routemonitor"])

	return nil

}

func deleteClusternode_routemonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternode_routemonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	routemonitor := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("routemonitor:%s",url.QueryEscape(routemonitor)))
	args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(d.Get("netmask").(string))))


	err := client.DeleteResourceWithArgs("clusternode_routemonitor_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
