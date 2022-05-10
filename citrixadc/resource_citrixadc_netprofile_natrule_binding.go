package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
	"net/url"
)

func resourceCitrixAdcNetprofile_natrule_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNetprofile_natrule_bindingFunc,
		Read:          readNetprofile_natrule_bindingFunc,
		Delete:        deleteNetprofile_natrule_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"natrule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"rewriteip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNetprofile_natrule_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNetprofile_natrule_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	natrule := d.Get("natrule")
	bindingId := fmt.Sprintf("%s,%s", name, natrule)
	netprofile_natrule_binding := network.Netprofilenatrulebinding{
		Name:      d.Get("name").(string),
		Natrule:   d.Get("natrule").(string),
		Netmask:   d.Get("netmask").(string),
		Rewriteip: d.Get("rewriteip").(string),
	}

	err := client.UpdateUnnamedResource(service.Netprofile_natrule_binding.Type(), &netprofile_natrule_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readNetprofile_natrule_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this netprofile_natrule_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readNetprofile_natrule_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNetprofile_natrule_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	natrule := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading netprofile_natrule_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "netprofile_natrule_binding",
		ResourceName:             name,
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
		log.Printf("[WARN] citrixadc-provider: Clearing netprofile_natrule_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["natrule"].(string) == natrule {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing netprofile_natrule_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("natrule", data["natrule"])
	d.Set("netmask", data["netmask"])
	d.Set("rewriteip", data["rewriteip"])

	return nil

}

func deleteNetprofile_natrule_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNetprofile_natrule_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	natrule := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("natrule:%s", url.QueryEscape(natrule)))
	if val, ok := d.GetOk("netmask"); ok {
		args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Netprofile_natrule_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
