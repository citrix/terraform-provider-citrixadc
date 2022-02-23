package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnglobal_vpneula_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnglobal_vpneula_bindingFunc,
		Read:          readVpnglobal_vpneula_bindingFunc,
		Delete:        deleteVpnglobal_vpneula_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"eula": &schema.Schema{
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

func createVpnglobal_vpneula_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_vpneula_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	eula := d.Get("eula").(string)
	vpnglobal_vpneula_binding := vpn.Vpnglobalvpneulabinding{
		Eula:                   d.Get("eula").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_vpneula_binding.Type(), &vpnglobal_vpneula_binding)
	if err != nil {
		return err
	}

	d.SetId(eula)

	err = readVpnglobal_vpneula_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnglobal_vpneula_binding but we can't read it ?? %s", eula)
		return nil
	}
	return nil
}

func readVpnglobal_vpneula_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_vpneula_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	eula := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_vpneula_binding state %s", eula)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_vpneula_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpneula_binding state %s", eula)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["eula"].(string) == eula {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_vpneula_binding state %s", eula)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("eula", data["eula"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])

	return nil

}

func deleteVpnglobal_vpneula_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_vpneula_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	eula := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("eula:%s", eula))

	err := client.DeleteResourceWithArgs(service.Vpnglobal_vpneula_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
