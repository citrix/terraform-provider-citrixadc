package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnglobal_vpnurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnglobal_vpnurl_bindingFunc,
		Read:          readVpnglobal_vpnurl_bindingFunc,
		Delete:        deleteVpnglobal_vpnurl_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"urlname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
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

func createVpnglobal_vpnurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	urlname := d.Get("urlname").(string)
	vpnglobal_vpnurl_binding := vpn.Vpnglobalvpnurlbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Urlname:                d.Get("urlname").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_vpnurl_binding.Type(), &vpnglobal_vpnurl_binding)
	if err != nil {
		return err
	}

	d.SetId(urlname)

	err = readVpnglobal_vpnurl_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnglobal_vpnurl_binding but we can't read it ?? %s", urlname)
		return nil
	}
	return nil
}

func readVpnglobal_vpnurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	urlname := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_vpnurl_binding state %s", urlname)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_vpnurl_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnurl_binding state %s", urlname)
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
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnurl_binding state %s", urlname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("urlname", data["urlname"])

	return nil

}

func deleteVpnglobal_vpnurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	urlname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("urlname:%s", urlname))

	err := client.DeleteResourceWithArgs(service.Vpnglobal_vpnurl_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
