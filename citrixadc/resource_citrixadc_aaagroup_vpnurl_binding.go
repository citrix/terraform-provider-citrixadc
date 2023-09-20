package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcAaagroup_vpnurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaagroup_vpnurl_bindingFunc,
		Read:          readAaagroup_vpnurl_bindingFunc,
		Delete:        deleteAaagroup_vpnurl_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"urlname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAaagroup_vpnurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaagroup_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname")
	urlname := d.Get("urlname")
	bindingId := fmt.Sprintf("%s,%s", groupname, urlname)
	aaagroup_vpnurl_binding := aaa.Aaagroupvpnurlbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Groupname:              d.Get("groupname").(string),
		Urlname:                d.Get("urlname").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaagroup_vpnurl_binding.Type(), &aaagroup_vpnurl_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAaagroup_vpnurl_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaagroup_vpnurl_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAaagroup_vpnurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaagroup_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	urlname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading aaagroup_vpnurl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "aaagroup_vpnurl_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup_vpnurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["urlname"].(string) == urlname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams urlname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup_vpnurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("groupname", data["groupname"])
	d.Set("urlname", data["urlname"])

	return nil

}

func deleteAaagroup_vpnurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaagroup_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	urlname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("urlname:%s", urlname))

	err := client.DeleteResourceWithArgs(service.Aaagroup_vpnurl_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
