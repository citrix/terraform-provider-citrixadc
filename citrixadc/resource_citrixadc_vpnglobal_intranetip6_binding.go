package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	//"net/url"
)

func resourceCitrixAdcVpnglobal_intranetip6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnglobal_intranetip6_bindingFunc,
		Read:          readVpnglobal_intranetip6_bindingFunc,
		Delete:        deleteVpnglobal_intranetip6_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"intranetip6": &schema.Schema{
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
			"numaddr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnglobal_intranetip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	intranetip6 := d.Get("intranetip6").(string)
	
	vpnglobal_intranetip6_binding := vpn.Vpnglobalintranetip6binding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Intranetip6:            d.Get("intranetip6").(string),
		Numaddr:                d.Get("numaddr").(int),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_intranetip6_binding.Type(), &vpnglobal_intranetip6_binding)
	if err != nil {
		return err
	}

	d.SetId(intranetip6)

	err = readVpnglobal_intranetip6_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnglobal_intranetip6_binding but we can't read it ?? %s", intranetip6)
		return nil
	}
	return nil
}

func readVpnglobal_intranetip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	intranetip6 := d.Id()
	
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_intranetip6_binding state %s", intranetip6)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_intranetip6_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_intranetip6_binding state %s", intranetip6)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetip6"].(string) == intranetip6 {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_intranetip6_binding state %s", intranetip6)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("intranetip6", data["intranetip6"])
	d.Set("numaddr", data["numaddr"])

	return nil

}

func deleteVpnglobal_intranetip6_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	intranetip6 := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetip6:%s", intranetip6))
	if val, ok := d.GetOk("numaddr"); ok {
		args = append(args, fmt.Sprintf("numaddr:%d", (val.(int))))
	}

	err := client.DeleteResourceWithArgs(service.Vpnglobal_intranetip6_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
