package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcAaauser_vpnurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaauser_vpnurl_bindingFunc,
		Read:          readAaauser_vpnurl_bindingFunc,
		Delete:        deleteAaauser_vpnurl_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"urlname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": &schema.Schema{
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

func createAaauser_vpnurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaauser_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	username := d.Get("username").(string)
	urlname := d.Get("urlname").(string)
	bindingId := fmt.Sprintf("%s,%s", username, urlname)
	aaauser_vpnurl_binding := aaa.Aaauservpnurlbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Urlname:                d.Get("urlname").(string),
		Username:               d.Get("username").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaauser_vpnurl_binding.Type(), &aaauser_vpnurl_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAaauser_vpnurl_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaauser_vpnurl_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAaauser_vpnurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaauser_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	username := idSlice[0]
	urlname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading aaauser_vpnurl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "aaauser_vpnurl_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaauser_vpnurl_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaauser_vpnurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("urlname", data["urlname"])
	d.Set("username", data["username"])

	return nil

}

func deleteAaauser_vpnurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaauser_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	urlname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("urlname:%s", urlname))

	err := client.DeleteResourceWithArgs(service.Aaauser_vpnurl_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
