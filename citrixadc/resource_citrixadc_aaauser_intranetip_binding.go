package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcAaauser_intranetip_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaauser_intranetip_bindingFunc,
		Read:          readAaauser_intranetip_bindingFunc,
		Delete:        deleteAaauser_intranetip_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"intranetip": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAaauser_intranetip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaauser_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	username := d.Get("username").(string)
	intranetip := d.Get("intranetip").(string)
	bindingId := fmt.Sprintf("%s,%s", username, intranetip)
	aaauser_intranetip_binding := aaa.Aaauserintranetipbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Intranetip:             d.Get("intranetip").(string),
		Netmask:                d.Get("netmask").(string),
		Username:               d.Get("username").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaauser_intranetip_binding.Type(), &aaauser_intranetip_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAaauser_intranetip_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaauser_intranetip_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAaauser_intranetip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaauser_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	username := idSlice[0]
	intranetip := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading aaauser_intranetip_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "aaauser_intranetip_binding",
		ResourceName:             username,
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaauser_intranetip_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaauser_intranetip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("intranetip", data["intranetip"])
	d.Set("netmask", data["netmask"])
	d.Set("username", data["username"])

	return nil

}

func deleteAaauser_intranetip_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaauser_intranetip_bindingFunc")
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

	err := client.DeleteResourceWithArgs(service.Aaauser_intranetip_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
