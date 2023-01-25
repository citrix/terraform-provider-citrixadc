package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
	"net/url"
)

func resourceCitrixAdcVpnglobal_appcontroller_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnglobal_appcontroller_bindingFunc,
		Read:          readVpnglobal_appcontroller_bindingFunc,
		Delete:        deleteVpnglobal_appcontroller_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"appcontroller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func createVpnglobal_appcontroller_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_appcontroller_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appcontroller := url.QueryEscape(d.Get("appcontroller").(string))
	vpnglobal_appcontroller_binding := vpn.Vpnglobalappcontrollerbinding{
		Appcontroller:          d.Get("appcontroller").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
	}

	_, err := client.AddResource(service.Vpnglobal_appcontroller_binding.Type(), appcontroller, &vpnglobal_appcontroller_binding)
	if err != nil {
		return err
	}

	d.SetId(appcontroller)

	err = readVpnglobal_appcontroller_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnglobal_appcontroller_binding but we can't read it ?? %s", appcontroller)
		return nil
	}
	return nil
}

func readVpnglobal_appcontroller_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_appcontroller_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	appcontroller, _ := url.QueryUnescape(d.Id())

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_appcontroller_binding state %s", appcontroller)

	findParams := service.FindParams{
		ResourceType: "vpnglobal_appcontroller_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_appcontroller_binding state %s", appcontroller)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["appcontroller"].(string) == appcontroller {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_appcontroller_binding state %s", appcontroller)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("appcontroller", data["appcontroller"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])

	return nil

}

func deleteVpnglobal_appcontroller_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_appcontroller_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	appcontroller := d.Id()

	argsMap := make(map[string]string)
	argsMap["appcontroller"] = appcontroller

	err := client.DeleteResourceWithArgsMap(service.Vpnglobal_appcontroller_binding.Type(), "", argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
