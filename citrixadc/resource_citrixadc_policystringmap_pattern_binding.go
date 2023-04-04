package citrixadc

import (
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/policy"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcPolicystringmap_pattern_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPolicystringmap_pattern_bindingFunc,
		Read:          readPolicystringmap_pattern_bindingFunc,
		Delete:        deletePolicystringmap_pattern_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createPolicystringmap_pattern_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicystringmap_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	key := d.Get("key").(string)
	bindingId := fmt.Sprintf("%s,%s", name, key)
	policystringmap_pattern_binding := policy.Policystringmappatternbinding{
		Key:   d.Get("key").(string),
		Name:  d.Get("name").(string),
		Value: d.Get("value").(string),
	}

	_, err := client.AddResource(service.Policystringmap_pattern_binding.Type(), name, &policystringmap_pattern_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readPolicystringmap_pattern_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this policystringmap_pattern_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readPolicystringmap_pattern_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicystringmap_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	name := idSlice[0]
	key := idSlice[1]
	log.Printf("[DEBUG] citrixadc-provider: Reading policystringmap_pattern_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "policystringmap_pattern_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing policystringmap_pattern_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right monitor name
	foundIndex := -1
	for i, v := range dataArr {
		if v["key"].(string) == key {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing policystringmap_pattern_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("key", data["key"])
	d.Set("value", data["value"])

	return nil

}

func deletePolicystringmap_pattern_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicystringmap_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	key := url.PathEscape(idSlice[1])

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("key:%s", key))

	err := client.DeleteResourceWithArgs("policystringmap_pattern_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
