package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcLsngroup_lsnpool_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsngroup_lsnpool_bindingFunc,
		Read:          readLsngroup_lsnpool_bindingFunc,
		Delete:        deleteLsngroup_lsnpool_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"poolname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsngroup_lsnpool_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsngroup_lsnpool_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname")
	poolname := d.Get("poolname")
	bindingId := fmt.Sprintf("%s,%s", groupname, poolname)
	lsngroup_lsnpool_binding := lsn.Lsngrouplsnpoolbinding{
		Groupname: d.Get("groupname").(string),
		Poolname:  d.Get("poolname").(string),
	}

	err := client.UpdateUnnamedResource("lsngroup_lsnpool_binding", &lsngroup_lsnpool_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLsngroup_lsnpool_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsngroup_lsnpool_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLsngroup_lsnpool_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsngroup_lsnpool_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	poolname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsngroup_lsnpool_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsngroup_lsnpool_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsnpool_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["poolname"].(string) == poolname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams poolname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsnpool_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("groupname", data["groupname"])
	d.Set("poolname", data["poolname"])

	return nil

}

func deleteLsngroup_lsnpool_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsngroup_lsnpool_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	poolname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("poolname:%s", poolname))

	err := client.DeleteResourceWithArgs("lsngroup_lsnpool_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
