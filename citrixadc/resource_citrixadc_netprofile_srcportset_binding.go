package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcNetprofile_srcportset_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNetprofile_srcportset_bindingFunc,
		Read:          readNetprofile_srcportset_bindingFunc,
		Delete:        deleteNetprofile_srcportset_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"srcportrange": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNetprofile_srcportset_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNetprofile_srcportset_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	srcportrange := d.Get("srcportrange")
	bindingId := fmt.Sprintf("%s,%s", name, srcportrange)
	netprofile_srcportset_binding := network.Netprofilesrcportsetbinding{
		Name:         d.Get("name").(string),
		Srcportrange: d.Get("srcportrange").(string),
	}

	err := client.UpdateUnnamedResource(service.Netprofile_srcportset_binding.Type(), &netprofile_srcportset_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readNetprofile_srcportset_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this netprofile_srcportset_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readNetprofile_srcportset_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNetprofile_srcportset_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	srcportrange := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading netprofile_srcportset_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "netprofile_srcportset_binding",
		ResourceName:             name,
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
		log.Printf("[WARN] citrixadc-provider: Clearing netprofile_srcportset_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["srcportrange"].(string) == srcportrange {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing netprofile_srcportset_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("srcportrange", data["srcportrange"])

	return nil

}

func deleteNetprofile_srcportset_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNetprofile_srcportset_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	srcportrange := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("srcportrange:%s", srcportrange))

	err := client.DeleteResourceWithArgs(service.Netprofile_srcportset_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
