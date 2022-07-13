package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
    "strings"
)


func resourceCitrixAdcGslbservice_lbmonitor_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbservice_lbmonitor_bindingFunc,
		Read:          readGslbservice_lbmonitor_bindingFunc,
		Delete:        deleteGslbservice_lbmonitor_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"monitor_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
                ForceNew: true,
			},
			"monstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
                ForceNew: true,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func createGslbservice_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbservice_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
    servicename:= d.Get("servicename").(string)
    monitor_name := d.Get("monitor_name").(string)
	
	bindingId := fmt.Sprintf("%s,%s", servicename, monitor_name)
	gslbservice_lbmonitor_binding := gslb.Gslbservicelbmonitorbinding{
		Monitorname:           d.Get("monitor_name").(string),
		Monstate:           d.Get("monstate").(string),
		Servicename:           d.Get("servicename").(string),
		Weight:           d.Get("weight").(int),
		
	}

	err := client.UpdateUnnamedResource(service.Gslbservice_lbmonitor_binding.Type(), &gslbservice_lbmonitor_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readGslbservice_lbmonitor_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbservice_lbmonitor_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readGslbservice_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbservice_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId:= d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	if len(idSlice) < 2 {
		return fmt.Errorf("Cannot deduce monitor_name from id string")
	}

	if len(idSlice) > 2 {
		return fmt.Errorf("Too many separators \",\" in id string")
	}

    servicename:= idSlice[0]
    monitor_name := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading gslbservice_lbmonitor_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "gslbservice_lbmonitor_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservice_lbmonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id 
	foundIndex := -1
	for i, v := range dataArr {
		if v["monitor_name"].(string) == monitor_name {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservice_lbmonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]


	d.Set("monitor_name", data["monitor_name"])
	d.Set("monstate", data["monstate"])
	d.Set("servicename", data["servicename"])
	d.Set("weight", data["weight"])
	

	return nil

}

func deleteGslbservice_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbservice_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	monitor_name  := idSlice[1]

	args := make([]string, 0) 
	args = append(args, fmt.Sprintf("monitor_name:%s", monitor_name ))

	err := client.DeleteResourceWithArgs(service.Gslbservice_lbmonitor_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
