package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSystemgroup_nspartition_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystemgroup_nspartition_bindingFunc,
		Read:          readSystemgroup_nspartition_bindingFunc,
		Delete:        deleteSystemgroup_nspartition_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"groupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"partitionname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSystemgroup_nspartition_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemgroup_nspartition_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname")
	partitionname := d.Get("partitionname")
	bindingId := fmt.Sprintf("%s,%s", groupname, partitionname)
	systemgroup_nspartition_binding := system.Systemgroupnspartitionbinding{
		Groupname:     d.Get("groupname").(string),
		Partitionname: d.Get("partitionname").(string),
	}

	err := client.UpdateUnnamedResource("systemgroup_nspartition_binding", &systemgroup_nspartition_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSystemgroup_nspartition_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this systemgroup_nspartition_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSystemgroup_nspartition_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemgroup_nspartition_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	partitionname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading systemgroup_nspartition_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "systemgroup_nspartition_binding",
		ResourceName:             groupname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing systemgroup_nspartition_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["partitionname"].(string) == partitionname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams partitionname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing systemgroup_nspartition_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("groupname", data["groupname"])
	d.Set("partitionname", data["partitionname"])

	return nil

}

func deleteSystemgroup_nspartition_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemgroup_nspartition_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	partitionname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("partitionname:%s", partitionname))

	err := client.DeleteResourceWithArgs("systemgroup_nspartition_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
