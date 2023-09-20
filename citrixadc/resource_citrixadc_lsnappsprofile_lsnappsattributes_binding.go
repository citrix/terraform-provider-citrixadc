package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcLsnappsprofile_lsnappsattributes_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnappsprofile_lsnappsattributes_bindingFunc,
		Read:          readLsnappsprofile_lsnappsattributes_bindingFunc,
		Delete:        deleteLsnappsprofile_lsnappsattributes_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"appsattributesname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"appsprofilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnappsprofile_lsnappsattributes_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnappsprofile_lsnappsattributes_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appsprofilename := d.Get("appsprofilename")
	appsattributesname := d.Get("appsattributesname")
	bindingId := fmt.Sprintf("%s,%s", appsprofilename, appsattributesname)
	lsnappsprofile_lsnappsattributes_binding := lsn.Lsnappsprofilelsnappsattributesbinding{
		Appsattributesname: d.Get("appsattributesname").(string),
		Appsprofilename:    d.Get("appsprofilename").(string),
	}

	err := client.UpdateUnnamedResource("lsnappsprofile_lsnappsattributes_binding", &lsnappsprofile_lsnappsattributes_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLsnappsprofile_lsnappsattributes_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnappsprofile_lsnappsattributes_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLsnappsprofile_lsnappsattributes_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnappsprofile_lsnappsattributes_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	appsprofilename := idSlice[0]
	appsattributesname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsnappsprofile_lsnappsattributes_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsnappsprofile_lsnappsattributes_binding",
		ResourceName:             appsprofilename,
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsnappsprofile_lsnappsattributes_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["appsattributesname"].(string) == appsattributesname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams appsattributesname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsnappsprofile_lsnappsattributes_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("appsattributesname", data["appsattributesname"])
	d.Set("appsprofilename", data["appsprofilename"])

	return nil

}

func deleteLsnappsprofile_lsnappsattributes_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnappsprofile_lsnappsattributes_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	appsattributesname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("appsattributesname:%s", appsattributesname))

	err := client.DeleteResourceWithArgs("lsnappsprofile_lsnappsattributes_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
