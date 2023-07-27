package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcLsngroup_lsnappsprofile_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsngroup_lsnappsprofile_bindingFunc,
		Read:          readLsngroup_lsnappsprofile_bindingFunc,
		Delete:        deleteLsngroup_lsnappsprofile_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"appsprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsngroup_lsnappsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsngroup_lsnappsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname")
	appsprofilename := d.Get("appsprofilename")
	bindingId := fmt.Sprintf("%s,%s", groupname, appsprofilename)
	lsngroup_lsnappsprofile_binding := lsn.Lsngrouplsnappsprofilebinding{
		Appsprofilename: d.Get("appsprofilename").(string),
		Groupname:       d.Get("groupname").(string),
	}

	err := client.UpdateUnnamedResource("lsngroup_lsnappsprofile_binding", &lsngroup_lsnappsprofile_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLsngroup_lsnappsprofile_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsngroup_lsnappsprofile_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLsngroup_lsnappsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsngroup_lsnappsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	appsprofilename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsngroup_lsnappsprofile_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsngroup_lsnappsprofile_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsnappsprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["appsprofilename"].(string) == appsprofilename {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams appsprofilename not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsnappsprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("appsprofilename", data["appsprofilename"])
	d.Set("groupname", data["groupname"])

	return nil

}

func deleteLsngroup_lsnappsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsngroup_lsnappsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	appsprofilename := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("appsprofilename:%s", appsprofilename))

	err := client.DeleteResourceWithArgs("lsngroup_lsnappsprofile_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
