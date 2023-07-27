package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcNsservicepath_nsservicefunction_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsservicepath_nsservicefunction_bindingFunc,
		Read:          readNsservicepath_nsservicefunction_bindingFunc,
		Delete:        deleteNsservicepath_nsservicefunction_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"servicepathname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"servicefunction": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNsservicepath_nsservicefunction_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsservicepath_nsservicefunction_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicepathname := d.Get("servicepathname")
	servicefunction := d.Get("servicefunction")
	bindingId := fmt.Sprintf("%s,%s", servicepathname, servicefunction)
	nsservicepath_nsservicefunction_binding := ns.Nsservicepathnsservicefunctionbinding{
		Index:           d.Get("index").(int),
		Servicefunction: d.Get("servicefunction").(string),
		Servicepathname: d.Get("servicepathname").(string),
	}

	err := client.UpdateUnnamedResource(service.Nsservicepath_nsservicefunction_binding.Type(), &nsservicepath_nsservicefunction_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readNsservicepath_nsservicefunction_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsservicepath_nsservicefunction_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readNsservicepath_nsservicefunction_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsservicepath_nsservicefunction_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicepathname := idSlice[0]
	servicefunction := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading nsservicepath_nsservicefunction_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "nsservicepath_nsservicefunction_binding",
		ResourceName:             servicepathname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing nsservicepath_nsservicefunction_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["servicefunction"].(string) == servicefunction {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing nsservicepath_nsservicefunction_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("index", data["index"])
	d.Set("servicefunction", data["servicefunction"])
	d.Set("servicepathname", data["servicepathname"])

	return nil

}

func deleteNsservicepath_nsservicefunction_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsservicepath_nsservicefunction_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	servicefunction := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("servicefunction:%s", servicefunction))

	err := client.DeleteResourceWithArgs(service.Nsservicepath_nsservicefunction_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
