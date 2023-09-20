package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnglobal_vpnintranetapplication_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnglobal_vpnintranetapplication_bindingFunc,
		Read:          readVpnglobal_vpnintranetapplication_bindingFunc,
		Delete:        deleteVpnglobal_vpnintranetapplication_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"intranetapplication": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnglobal_vpnintranetapplication_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_vpnintranetapplication_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	intranetapplication := d.Get("intranetapplication").(string)
	vpnglobal_vpnintranetapplication_binding := vpn.Vpnglobalvpnintranetapplicationbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Intranetapplication:    d.Get("intranetapplication").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_vpnintranetapplication_binding.Type(), &vpnglobal_vpnintranetapplication_binding)
	if err != nil {
		return err
	}

	d.SetId(intranetapplication)

	err = readVpnglobal_vpnintranetapplication_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnglobal_vpnintranetapplication_binding but we can't read it ?? %s", intranetapplication)
		return nil
	}
	return nil
}

func readVpnglobal_vpnintranetapplication_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_vpnintranetapplication_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	intranetapplication := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_vpnintranetapplication_binding state %s", intranetapplication)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_vpnintranetapplication_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnintranetapplication_binding state %s", intranetapplication)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetapplication"].(string) == intranetapplication {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpnintranetapplication_binding state %s", intranetapplication)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("intranetapplication", data["intranetapplication"])

	return nil

}

func deleteVpnglobal_vpnintranetapplication_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_vpnintranetapplication_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	intranetapplication := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetapplication:%s", intranetapplication))

	err := client.DeleteResourceWithArgs(service.Vpnglobal_vpnintranetapplication_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
