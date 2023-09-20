package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcGslbservice_dnsview_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbservice_dnsview_bindingFunc,
		Read:          readGslbservice_dnsview_bindingFunc,
		Delete:        deleteGslbservice_dnsview_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"servicename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"viewip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"viewname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createGslbservice_dnsview_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbservice_dnsview_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicename := d.Get("servicename")
	viewname := d.Get("viewname")

	bindingId := fmt.Sprintf("%s,%s", servicename, viewname)
	gslbservice_dnsview_binding := gslb.Gslbservicednsviewbinding{
		Servicename: d.Get("servicename").(string),
		Viewip:      d.Get("viewip").(string),
		Viewname:    d.Get("viewname").(string),
	}

	err := client.UpdateUnnamedResource(service.Gslbservice_dnsview_binding.Type(), &gslbservice_dnsview_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readGslbservice_dnsview_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbservice_dnsview_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readGslbservice_dnsview_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbservice_dnsview_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicename := idSlice[0]
	viewname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading gslbservice_dnsview_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "gslbservice_dnsview_binding",
		ResourceName:             servicename,
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
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservice_dnsview_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["viewname"].(string) == viewname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservice_dnsview_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("servicename", data["servicename"])
	d.Set("viewip", data["viewip"])
	d.Set("viewname", data["viewname"])

	return nil

}

func deleteGslbservice_dnsview_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbservice_dnsview_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicename := idSlice[0]
	viewname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("viewname:%s", viewname))

	err := client.DeleteResourceWithArgs(service.Gslbservice_dnsview_binding.Type(), servicename, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
