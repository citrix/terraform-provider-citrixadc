package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
	"strings"
)

func resourceCitrixAdcNspartition_vxlan_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNspartition_vxlan_bindingFunc,
		Read:          readNspartition_vxlan_bindingFunc,
		Delete:        deleteNspartition_vxlan_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"partitionname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"vxlan": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNspartition_vxlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNspartition_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	partitionname := d.Get("partitionname")
	vxlan := strconv.Itoa(d.Get("vxlan").(int))
	bindingId := fmt.Sprintf("%s,%s", partitionname, vxlan)
	nspartition_vxlan_binding := ns.Nspartitionvxlanbinding{
		Partitionname: d.Get("partitionname").(string),
		Vxlan:         d.Get("vxlan").(int),
	}

	err := client.UpdateUnnamedResource("nspartition_vxlan_binding", &nspartition_vxlan_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readNspartition_vxlan_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nspartition_vxlan_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readNspartition_vxlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNspartition_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	partitionname := idSlice[0]
	vxlan := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading nspartition_vxlan_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "nspartition_vxlan_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing nspartition_vxlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["vxlan"].(string) == vxlan {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing nspartition_vxlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("partitionname", data["partitionname"])
	d.Set("vxlan", data["vxlan"])

	return nil

}

func deleteNspartition_vxlan_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNspartition_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	vxlan := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("vxlan:%s", vxlan))

	err := client.DeleteResourceWithArgs("nspartition_vxlan_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
