package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcService_lbmonitor_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createService_lbmonitor_bindingFunc,
		Read:          readService_lbmonitor_bindingFunc,
		Delete:        deleteService_lbmonitor_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"monitor_name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"monstate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"passive": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createService_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createService_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	monitor_name := d.Get("monitor_name")
	bindingId := fmt.Sprintf("%s,%s", name, monitor_name)
	service_lbmonitor_binding := basic.Servicelbmonitorbinding{
		Monitorname: d.Get("monitor_name").(string),
		Monstate:    d.Get("monstate").(string),
		Name:        d.Get("name").(string),
		Passive:     d.Get("passive").(bool),
		Weight:      d.Get("weight").(int),
	}

	err := client.UpdateUnnamedResource(service.Service_lbmonitor_binding.Type(), &service_lbmonitor_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readService_lbmonitor_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this service_lbmonitor_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readService_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readService_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	monitor_name := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading service_lbmonitor_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "service_lbmonitor_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing service_lbmonitor_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing service_lbmonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("monitor_name", data["monitor_name"])
	// d.Set("monstate", data["monstate"])
	d.Set("name", data["name"])
	d.Set("passive", data["passive"])
	d.Set("weight", data["weight"])

	return nil

}

func deleteService_lbmonitor_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteService_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	monitor_name := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("monitor_name:%s", monitor_name))

	err := client.DeleteResourceWithArgs(service.Service_lbmonitor_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
