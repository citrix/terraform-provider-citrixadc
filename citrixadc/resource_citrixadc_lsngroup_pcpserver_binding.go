package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcLsngroup_pcpserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsngroup_pcpserver_bindingFunc,
		Read:          readLsngroup_pcpserver_bindingFunc,
		Delete:        deleteLsngroup_pcpserver_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pcpserver": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsngroup_pcpserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsngroup_pcpserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname")
	pcpserver := d.Get("pcpserver")
	bindingId := fmt.Sprintf("%s,%s", groupname, pcpserver)
	lsngroup_pcpserver_binding := lsn.Lsngrouppcpserverbinding{
		Groupname: d.Get("groupname").(string),
		Pcpserver: d.Get("pcpserver").(string),
	}

	err := client.UpdateUnnamedResource("lsngroup_pcpserver_binding", &lsngroup_pcpserver_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLsngroup_pcpserver_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsngroup_pcpserver_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLsngroup_pcpserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsngroup_pcpserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	pcpserver := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsngroup_pcpserver_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsngroup_pcpserver_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_pcpserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["pcpserver"].(string) == pcpserver {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams pcpserver not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_pcpserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("groupname", data["groupname"])
	d.Set("pcpserver", data["pcpserver"])

	return nil

}

func deleteLsngroup_pcpserver_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsngroup_pcpserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	pcpserver := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("pcpserver:%s", pcpserver))

	err := client.DeleteResourceWithArgs("lsngroup_pcpserver_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
