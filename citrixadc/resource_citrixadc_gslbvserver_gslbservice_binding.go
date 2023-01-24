package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcGslbvserver_gslbservice_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbvserver_gslbservice_bindingFunc,
		Read:          readGslbvserver_gslbservice_bindingFunc,
		Delete:        deleteGslbvserver_gslbservice_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domainname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createGslbvserver_gslbservice_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbvserver_gslbservice_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	servicename := d.Get("servicename")
	
	bindingId := fmt.Sprintf("%s,%s", name, servicename)
	gslbvserver_gslbservice_binding := gslb.Gslbvservergslbservicebinding{
		Domainname:  d.Get("domainname").(string),
		Name:        d.Get("name").(string),
		Servicename: d.Get("servicename").(string),
		Weight:      d.Get("weight").(int),
	}

    err := client.UpdateUnnamedResource(service.Gslbvserver_gslbservice_binding.Type(), &gslbvserver_gslbservice_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readGslbvserver_gslbservice_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbvserver_gslbservice_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readGslbvserver_gslbservice_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbvserver_gslbservice_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	servicename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading gslbvserver_gslbservice_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "gslbvserver_gslbservice_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing gslbvserver_gslbservice_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["servicename"].(string) == servicename {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbvserver_gslbservice_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("domainname", data["domainname"])
	d.Set("name", data["name"])
	d.Set("servicename", data["servicename"])
	d.Set("weight", data["weight"])

	return nil

}

func deleteGslbvserver_gslbservice_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbvserver_gslbservice_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	servicename := idSlice[1]
	
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("servicename:%s", servicename))

	err := client.DeleteResourceWithArgs(service.Gslbvserver_gslbservice_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
