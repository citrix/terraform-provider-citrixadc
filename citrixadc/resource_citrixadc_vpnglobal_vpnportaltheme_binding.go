package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnglobal_vpnportaltheme_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnglobal_vpnportaltheme_bindingFunc,
		Read:          readVpnglobal_vpnportaltheme_bindingFunc,
		Delete:        deleteVpnglobal_vpnportaltheme_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"portaltheme": &schema.Schema{
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

func createVpnglobal_vpnportaltheme_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_vpnportaltheme_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	portaltheme := d.Get("portaltheme").(string)
	vpnglobal_vpnportaltheme_binding := vpn.Vpnglobalvpnportalthemebinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Portaltheme:            d.Get("portaltheme").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_vpnportaltheme_binding.Type(), &vpnglobal_vpnportaltheme_binding)
	if err != nil {
		return err
	}

	d.SetId(portaltheme)

	err = readVpnglobal_vpnportaltheme_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnglobal_vpnportaltheme_binding but we can't read it ?? %s", portaltheme)
		return nil
	}
	return nil
}

func readVpnglobal_vpnportaltheme_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_vpnportaltheme_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	portaltheme := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_vpnportaltheme_binding state %s", portaltheme)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_vpnportaltheme_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnportaltheme_binding state %s", portaltheme)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["portaltheme"].(string) == portaltheme {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnportaltheme_binding state %s", portaltheme)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("portaltheme", data["portaltheme"])

	return nil

}

func deleteVpnglobal_vpnportaltheme_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_vpnportaltheme_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	portaltheme := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("portaltheme:%s", portaltheme))

	err := client.DeleteResourceWithArgs(service.Vpnglobal_vpnportaltheme_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
