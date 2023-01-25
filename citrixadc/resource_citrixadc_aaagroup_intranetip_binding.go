package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcAaagroup_intranetip_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaagroup_intranetip_bindingFunc,
		Read:          readAaagroup_intranetip_bindingFunc,
		Delete:        deleteAaagroup_intranetip_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"groupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"intranetip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAaagroup_intranetip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaagroup_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname").(string)
	intranetip := d.Get("intranetip").(string)
	bindingId := fmt.Sprintf("%s,%s", groupname, intranetip)
	aaagroup_intranetip_binding := aaa.Aaagroupintranetipbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Groupname:              d.Get("groupname").(string),
		Intranetip:             d.Get("intranetip").(string),
		Netmask:                d.Get("netmask").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaagroup_intranetip_binding.Type(), &aaagroup_intranetip_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAaagroup_intranetip_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaagroup_intranetip_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAaagroup_intranetip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaagroup_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	intranetip := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading aaagroup_intranetip_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "aaagroup_intranetip_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup_intranetip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetip"].(string) == intranetip {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams intranetip not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup_intranetip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("groupname", data["groupname"])
	d.Set("intranetip", data["intranetip"])
	d.Set("netmask", data["netmask"])

	return nil

}

func deleteAaagroup_intranetip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaagroup_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	intranetip := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetip:%s", intranetip))
	if v, ok := d.GetOk("netmask"); ok {
		netmask := v.(string)
		args = append(args, fmt.Sprintf("netmask:%s", netmask))
	}

	err := client.DeleteResourceWithArgs(service.Aaagroup_intranetip_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
