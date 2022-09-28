package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcLsngroup_lsntransportprofile_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsngroup_lsntransportprofile_bindingFunc,
		Read:          readLsngroup_lsntransportprofile_bindingFunc,
		Delete:        deleteLsngroup_lsntransportprofile_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"groupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transportprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsngroup_lsntransportprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsngroup_lsntransportprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname")
	transportprofilename := d.Get("transportprofilename")
	bindingId := fmt.Sprintf("%s,%s", groupname, transportprofilename)
	lsngroup_lsntransportprofile_binding := lsn.Lsngrouplsntransportprofilebinding{
		Groupname:            d.Get("groupname").(string),
		Transportprofilename: d.Get("transportprofilename").(string),
	}

	err := client.UpdateUnnamedResource("lsngroup_lsntransportprofile_binding", &lsngroup_lsntransportprofile_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLsngroup_lsntransportprofile_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsngroup_lsntransportprofile_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLsngroup_lsntransportprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsngroup_lsntransportprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	transportprofilename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsngroup_lsntransportprofile_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsngroup_lsntransportprofile_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsntransportprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["transportprofilename"].(string) == transportprofilename {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams transportprofilename not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsntransportprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("groupname", data["groupname"])
	d.Set("transportprofilename", data["transportprofilename"])

	return nil

}

func deleteLsngroup_lsntransportprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsngroup_lsntransportprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	transportprofilename := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("transportprofilename:%s", transportprofilename))

	err := client.DeleteResourceWithArgs("lsngroup_lsntransportprofile_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
