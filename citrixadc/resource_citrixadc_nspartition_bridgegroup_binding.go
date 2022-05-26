package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
	"strconv"
)

func resourceCitrixAdcNspartition_bridgegroup_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNspartition_bridgegroup_bindingFunc,
		Read:          readNspartition_bridgegroup_bindingFunc,
		Delete:        deleteNspartition_bridgegroup_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"partitionname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"bridgegroup": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNspartition_bridgegroup_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNspartition_bridgegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	partitionname := d.Get("partitionname")
	bridgegroup := strconv.Itoa(d.Get("bridgegroup").(int))
	bindingId := fmt.Sprintf("%s,%s", partitionname, bridgegroup)
	nspartition_bridgegroup_binding := ns.Nspartitionbridgegroupbinding{
		Bridgegroup:   d.Get("bridgegroup").(int),
		Partitionname: d.Get("partitionname").(string),
	}

	err := client.UpdateUnnamedResource(service.Nspartition_bridgegroup_binding.Type(), &nspartition_bridgegroup_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readNspartition_bridgegroup_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nspartition_bridgegroup_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readNspartition_bridgegroup_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNspartition_bridgegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	partitionname := idSlice[0]
	bridgegroup := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading nspartition_bridgegroup_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "nspartition_bridgegroup_binding",
		ResourceName:             partitionname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing nspartition_bridgegroup_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["bridgegroup"].(string) == bridgegroup {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing nspartition_bridgegroup_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bridgegroup", data["bridgegroup"])
	d.Set("partitionname", data["partitionname"])

	return nil

}

func deleteNspartition_bridgegroup_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNspartition_bridgegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bridgegroup := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("bridgegroup:%s", bridgegroup))

	err := client.DeleteResourceWithArgs(service.Nspartition_bridgegroup_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
