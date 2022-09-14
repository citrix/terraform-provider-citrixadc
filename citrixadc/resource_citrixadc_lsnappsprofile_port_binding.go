package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcLsnappsprofile_port_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnappsprofile_port_bindingFunc,
		Read:          readLsnappsprofile_port_bindingFunc,
		Delete:        deleteLsnappsprofile_port_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"appsprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"lsnport": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnappsprofile_port_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnappsprofile_port_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appsprofilename := d.Get("appsprofilename")
	lsnport := d.Get("lsnport")
	bindingId := fmt.Sprintf("%s,%s", appsprofilename, lsnport)
	lsnappsprofile_port_binding := lsn.Lsnappsprofileportbinding{
		Appsprofilename: d.Get("appsprofilename").(string),
		Lsnport:         d.Get("lsnport").(string),
	}

	err := client.UpdateUnnamedResource("lsnappsprofile_port_binding", &lsnappsprofile_port_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLsnappsprofile_port_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnappsprofile_port_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLsnappsprofile_port_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnappsprofile_port_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	appsprofilename := idSlice[0]
	lsnport := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsnappsprofile_port_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsnappsprofile_port_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsnappsprofile_port_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["lsnport"].(string) == lsnport {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams lsnport not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsnappsprofile_port_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("appsprofilename", data["appsprofilename"])
	d.Set("lsnport", data["lsnport"])

	return nil

}

func deleteLsnappsprofile_port_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnappsprofile_port_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	lsnport := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("lsnport:%s", lsnport))

	err := client.DeleteResourceWithArgs("lsnappsprofile_port_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
