package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
	"strings"
)

func resourceCitrixAdcClusternodegroup_clusternode_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createClusternodegroup_clusternode_bindingFunc,
		Read:          readClusternodegroup_clusternode_bindingFunc,
		Delete:        deleteClusternodegroup_clusternode_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"node": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createClusternodegroup_clusternode_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternodegroup_clusternode_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	node := strconv.Itoa(d.Get("node").(int))
	bindingId := fmt.Sprintf("%s,%s", name, node)
	clusternodegroup_clusternode_binding := cluster.Clusternodegroupclusternodebinding{
		Name: d.Get("name").(string),
		Node: d.Get("node").(int),
	}

	err := client.UpdateUnnamedResource(service.Clusternodegroup_clusternode_binding.Type(), &clusternodegroup_clusternode_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readClusternodegroup_clusternode_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this clusternodegroup_clusternode_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readClusternodegroup_clusternode_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternodegroup_clusternode_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	node := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading clusternodegroup_clusternode_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "clusternodegroup_clusternode_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_clusternode_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["node"].(string) == node {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams node not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_clusternode_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("node", data["node"])

	return nil

}

func deleteClusternodegroup_clusternode_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternodegroup_clusternode_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	node := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("node:%s", node))

	err := client.DeleteResourceWithArgs(service.Clusternodegroup_clusternode_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
