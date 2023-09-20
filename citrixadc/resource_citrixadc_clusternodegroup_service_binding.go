package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcClusternodegroup_service_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createClusternodegroup_service_bindingFunc,
		Read:          readClusternodegroup_service_bindingFunc,
		Delete:        deleteClusternodegroup_service_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createClusternodegroup_service_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternodegroup_service_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	serviceName := d.Get("service")
	bindingId := fmt.Sprintf("%s,%s", name, serviceName)
	clusternodegroup_service_binding := cluster.Clusternodegroupservicebinding{
		Name:    d.Get("name").(string),
		Service: d.Get("service").(string),
	}

	err := client.UpdateUnnamedResource("clusternodegroup_service_binding", &clusternodegroup_service_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readClusternodegroup_service_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this clusternodegroup_service_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readClusternodegroup_service_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternodegroup_service_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	serviceName := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading clusternodegroup_service_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "clusternodegroup_service_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_service_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["service"].(string) == serviceName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams service not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_service_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("service", data["service"])

	return nil

}

func deleteClusternodegroup_service_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternodegroup_service_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	serviceName := idSlice[1]

	args := make([]string, 0)
	//args = append(args, fmt.Sprintf("name:%s", url.QueryEscape(name)))
	args = append(args, fmt.Sprintf("service:%s", url.QueryEscape(serviceName)))

	err := client.DeleteResourceWithArgs("clusternodegroup_service_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
