package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcLsngroup_lsnlogprofile_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsngroup_lsnlogprofile_bindingFunc,
		Read:          readLsngroup_lsnlogprofile_bindingFunc,
		Delete:        deleteLsngroup_lsnlogprofile_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"groupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"logprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsngroup_lsnlogprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsngroup_lsnlogprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname")
	logprofilename := d.Get("logprofilename")
	bindingId := fmt.Sprintf("%s,%s", groupname, logprofilename)
	lsngroup_lsnlogprofile_binding := lsn.Lsngrouplsnlogprofilebinding{
		Groupname:      d.Get("groupname").(string),
		Logprofilename: d.Get("logprofilename").(string),
	}

	err := client.UpdateUnnamedResource("lsngroup_lsnlogprofile_binding", &lsngroup_lsnlogprofile_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLsngroup_lsnlogprofile_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsngroup_lsnlogprofile_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLsngroup_lsnlogprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsngroup_lsnlogprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	logprofilename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsngroup_lsnlogprofile_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsngroup_lsnlogprofile_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsnlogprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["logprofilename"].(string) == logprofilename {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams logprofilename not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsnlogprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("groupname", data["groupname"])
	d.Set("logprofilename", data["logprofilename"])

	return nil

}

func deleteLsngroup_lsnlogprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsngroup_lsnlogprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	logprofilename := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("logprofilename:%s", logprofilename))

	err := client.DeleteResourceWithArgs("lsngroup_lsnlogprofile_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
